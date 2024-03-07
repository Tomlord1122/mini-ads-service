package util

import (
	"math/rand"
	"strings"
	"time"
)

type GenderEnum string

const (
	GenderEnumM GenderEnum = "M"
	GenderEnumF GenderEnum = "F"
)

type CountryEnum string

const (
	CountryEnumTW CountryEnum = "TW"
	CountryEnumJP CountryEnum = "JP"
	CountryEnumUS CountryEnum = "US"
)

type PlatformEnum string

const (
	PlatformEnumAndroid PlatformEnum = "android"
	PlatformEnumIOS     PlatformEnum = "ios"
	PlatformEnumWeb     PlatformEnum = "web"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomTime
func RandomTime() time.Time {

	currentTime := time.Now()

	secondsInTenYears := int64(1 * 24 * 3600)
	randomSeconds := rand.Int63n(secondsInTenYears)

	randomTime := currentTime.Add(time.Second * time.Duration(randomSeconds))

	return randomTime
}

// RandomGender returns a random gender value ('M' or 'F')
func RandomGender() GenderEnum {
	genders := []GenderEnum{GenderEnumM, GenderEnumF}
	return genders[rand.Intn(len(genders))]
}

// RandomCountry returns a random country value ('TW', 'JP', 'US')
func RandomCountry() CountryEnum {
	countries := []CountryEnum{CountryEnumTW, CountryEnumJP, CountryEnumUS}
	return countries[rand.Intn(len(countries))]
}

// RandomPlatform returns a random platform value ('android', 'ios', 'web')
func RandomPlatform() PlatformEnum {
	platforms := []PlatformEnum{PlatformEnumAndroid, PlatformEnumIOS, PlatformEnumWeb}
	return platforms[rand.Intn(len(platforms))]
}
