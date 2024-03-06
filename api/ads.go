package api

import (
	db "backend-intern/db/sqlc"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func (server *Server) CreateAds(ctx *gin.Context) {
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

	ads, err := server.query.ListAds(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ads)
}
