package api

import (
	db "backend-intern/db/sqlc"
	"backend-intern/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type createAdsRequest struct {
	Title    string            `json:"title"`
	StartAt  time.Time         `json:"start_at"`
	EndAt    time.Time         `json:"end_at"`
	Age      *int32            `json:"age"`
	Gender   []db.GenderEnum   `json:"gender"`
	Country  []db.CountryEnum  `json:"country"`
	Platform []db.PlatformEnum `json:"platform"`
}

type listAdsRequest struct {
	AgeStart int32    `form:"age_start"`
	AgeEnd   int32    `form:"age_end"`
	Country  []string `form:"country[]"`
	Platform []string `form:"platform[]"`
	Gender   []string `form:"gender[]"`
	Limit    int32    `form:"limit"`
	Offset   int32    `form:"offset"`
}

func GenerateCacheKey(req listAdsRequest) string {

	hasher := fnv.New32a()

	hasher.Write([]byte(fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v",
		req.AgeStart, req.AgeEnd, req.Country, req.Platform, req.Gender, req.Limit, req.Offset)))
	return fmt.Sprintf("ads_list_%d", hasher.Sum32())
}

func (server *Server) CreateRandomAds(ctx *gin.Context) {

	// generate a key for the cache
	date := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("create_ads_%s", date)

	// check the api request limit
	count, err := server.redis.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if count >= 3000 {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "add ads call limit reached for today"})
		return
	}

	// add the count
	newCount, err := server.redis.Incr(ctx, key).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if newCount == 1 {
		// set the expiration time for the key
		_, err := server.redis.Expire(ctx, key, 24*time.Hour).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	activeAdsCount, err := server.query.GetActiveAds(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if activeAdsCount[0] >= 1000 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Cannot create more ads. The limit of active ads has been reached."})
		return
	}

	age := util.RandomInt(1, 100)
	arg := db.CreateAdsParams{
		Title:    util.RandomString(5),
		StartAt:  time.Now().UTC(),        // Convert to UTC
		EndAt:    util.RandomTime().UTC(), // Convert to UTC
		Age:      sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:   []db.GenderEnum{db.GenderEnum(util.RandomGender())},
		Country:  []db.CountryEnum{db.CountryEnum(util.RandomCountry())},
		Platform: []db.PlatformEnum{db.PlatformEnum(util.RandomPlatform())},
	}
	ad, err := server.query.CreateAds(ctx, arg)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, ad)
}

func (server *Server) CreateAds(ctx *gin.Context) {

	// generate a key for the cache
	date := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("create_ads_%s", date)

	// check the api request limit
	count, err := server.redis.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if count >= 3000 {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "add ads call limit reached for today"})
		return
	}

	// add the count
	newCount, err := server.redis.Incr(ctx, key).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if newCount == 1 {
		// set the expiration time for the key
		_, err := server.redis.Expire(ctx, key, 24*time.Hour).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	activeAdsCount, err := server.query.GetActiveAds(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if activeAdsCount[0] >= 1000 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Cannot create more ads. The limit of active ads has been reached."})
		return
	}

	var req createAdsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}
	var age sql.NullInt32
	if req.Age != nil {
		age = sql.NullInt32{Int32: *req.Age, Valid: true}
	} else {
		age = sql.NullInt32{Valid: false}
	}

	arg := db.CreateAdsParams{
		Title:    req.Title,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		Age:      age,
		Gender:   req.Gender,
		Country:  req.Country,
		Platform: req.Platform,
	}
	ad, err := server.query.CreateAds(ctx, arg)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}
	ctx.JSON(200, ad)

}

func (server *Server) ListAds(ctx *gin.Context) {

	var req listAdsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 轉換 country, platform, gender 從 string 到 Enum 類型
	countryEnums := make([]db.CountryEnum, len(req.Country))
	for i, c := range req.Country {
		countryEnums[i] = db.CountryEnum(c)
	}

	platformEnums := make([]db.PlatformEnum, len(req.Platform))
	for i, p := range req.Platform {
		platformEnums[i] = db.PlatformEnum(p)
	}

	genderEnums := make([]db.GenderEnum, len(req.Gender))
	for i, g := range req.Gender {
		genderEnums[i] = db.GenderEnum(g)
	}

	arg := db.ListAdsParams{
		Column1: req.AgeStart,
		Column2: req.AgeEnd,
		Column3: countryEnums,
		Column4: platformEnums,
		Column5: genderEnums,
		Column6: req.Limit,
		Offset:  req.Offset,
	}

	cachekey := GenerateCacheKey(req)
	// Get data from Redis cache
	val, err := server.redis.Get(ctx, cachekey).Result() // Get data from Redis cache
	if err == redis.Nil {
		// data does not exist in Redis cache
		ads, err := server.query.ListAds(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		// Set data to Redis cache
		jsonData, err := json.Marshal(ads) // type invert to json
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		server.redis.Set(ctx, cachekey, jsonData, 30*time.Minute)
		ctx.JSON(http.StatusOK, ads)
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		// data exist in Redis cache
		var ads []db.ListAdsRow
		err := json.Unmarshal([]byte(val), &ads) // type invert to struct
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, ads)
	}

}
