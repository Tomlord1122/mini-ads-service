package db

import (
	"backend-intern/util"
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomAds(t *testing.T) Ad {

	age := util.RandomInt(1, 100)
	arg := CreateAdsParams{
		Title:    util.RandomString(5),
		StartAt:  time.Now().UTC(),        // Convert to UTC
		EndAt:    util.RandomTime().UTC(), // Convert to UTC
		Age:      sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:   []GenderEnum{GenderEnum(util.RandomGender())},
		Country:  []CountryEnum{CountryEnum(util.RandomCountry())},
		Platform: []PlatformEnum{PlatformEnum(util.RandomPlatform())},
	}

	ad, err := testQueries.CreateAds(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	require.NotEmpty(t, ad)
	require.Equal(t, arg.Title, ad.Title)
	require.WithinDuration(t, arg.StartAt, ad.StartAt, time.Second)
	require.WithinDuration(t, arg.EndAt, ad.EndAt, time.Second)
	require.Equal(t, arg.Age, ad.Age)
	require.Equal(t, arg.Gender, ad.Gender)
	require.Equal(t, arg.Country, ad.Country)
	require.Equal(t, arg.Platform, ad.Platform)

	return ad
}

func TestCreateAds(t *testing.T) {
	for i := 0; i < 300; i++ {
		CreateRandomAds(t)
	}
}

func TestListAds(t *testing.T) {
	ad := CreateRandomAds(t)
	arg := ListAdsParams{
		Column1: ad.Age.Int32,
		Column2: ad.Age.Int32,
		Column3: ad.Country,
		Column4: ad.Platform,
		Column5: ad.Gender,
		Column6: 1,
		Offset:  0,
	}

	ads, err := testQueries.ListAds(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ads)

}
