package main

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"strconv"
	"time"
)

type State struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryID   int    `json:"country_id"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	StateCode   string `json:"state_code"`
	Type        any    `json:"type"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type States []State

type StateInDB struct {
	ID          uint32    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	CountryID   uint32    `json:"country_id" db:"country_id"`
	CountryCode string    `json:"country_code" db:"country_code"`
	Type        string    `json:"type" db:"type"`
	Latitude    float64   `json:"latitude" db:"latitude"`
	Longitude   float64   `json:"longitude" db:"longitude"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Flag        bool      `json:"flag" db:"flag"`
	WikiDataID  string    `json:"wikiDataId" db:"wiki_data_id"`
}
type StatesInDB []StateInDB

type CountryTimezone struct {
	ID            int    `json:"id" db:"id"`
	ZoneName      string `json:"zoneName" db:"zone_name"`
	GmtOffset     int    `json:"gmtOffset" db:"gmt_offset"`
	GmtOffsetName string `json:"gmtOffsetName" db:"gmt_offset_name"`
	Abbreviation  string `json:"abbreviation" db:"abbreviation"`
	TzName        string `json:"tzName" db:"tz_name"`
}

type CountryTimezones []CountryTimezone

type CountryInDB struct {
	ID             uint32    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	ISO3           string    `json:"iso3,omitempty" db:"iso3"`
	NumericCode    string    `json:"numeric_code,omitempty" db:"numeric_code"`
	ISO2           string    `json:"iso2,omitempty" db:"iso2"`
	PhoneCode      string    `json:"phonecode,omitempty" db:"phonecode"`
	Capital        string    `json:"capital,omitempty" db:"capital"`
	Currency       string    `json:"currency,omitempty" db:"currency"`
	CurrencyName   string    `json:"currency_name,omitempty" db:"currency_name"`
	CurrencySymbol string    `json:"currency_symbol,omitempty" db:"currency_symbol"`
	TLD            string    `json:"tld,omitempty" db:"tld"`
	Native         string    `json:"native,omitempty" db:"native"`
	Region         string    `json:"region,omitempty" db:"region"`
	Subregion      string    `json:"subregion,omitempty" db:"subregion"`
	TimezoneID     []uint32  `json:"timezone_id,omitempty" db:"timezone_id"`
	Translations   string    `json:"translations,omitempty" db:"translations"`
	Latitude       float64   `json:"latitude,omitempty" db:"latitude"`
	Longitude      float64   `json:"longitude,omitempty" db:"longitude"`
	Emoji          string    `json:"emoji,omitempty" db:"emoji"`
	EmojiU         string    `json:"emojiU,omitempty" db:"emojiU"`
	CreatedAt      time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type CountriesInDB []CountryInDB

type Country struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Iso3           string `json:"iso3"`
	Iso2           string `json:"iso2"`
	NumericCode    string `json:"numeric_code"`
	PhoneCode      string `json:"phone_code"`
	Capital        string `json:"capital"`
	Currency       string `json:"currency"`
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	Tld            string `json:"tld"`
	Native         string `json:"native"`
	Region         string `json:"region"`
	Subregion      string `json:"subregion"`
	Timezones      CountryTimezones
	Translations   struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Hr   string `json:"hr"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		Ja   string `json:"ja"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Emoji         string `json:"emoji"`
	EmojiU        string `json:"emojiU"`
	Translations0 struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Fr   string `json:"fr"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
	Translations1 struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Fr   string `json:"fr"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
	Translations2 struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Hr   string `json:"hr"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		Ja   string `json:"ja"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
	Translations3 struct {
		Kr string `json:"kr"`
		Cn string `json:"cn"`
		Tr string `json:"tr"`
	} `json:"translations,omitempty"`
	Translations4 struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Fr   string `json:"fr"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
	Translations5 struct {
		Kr   string `json:"kr"`
		PtBR string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		Ja   string `json:"ja"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations,omitempty"`
}
type Countries []Country
type City struct {
	ID          uint32    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	StateID     uint32    `json:"state_id" db:"state_id"`
	StateCode   string    `json:"state_code" db:"state_code"`
	CountryID   uint32    `json:"country_id" db:"country_id"`
	CountryCode string    `json:"country_code" db:"country_code"`
	Latitude    float64   `json:"latitude" db:"latitude"`
	Longitude   float64   `json:"longitude" db:"longitude"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	WikiDataID  string    `json:"wikiDataId,omitempty" db:"wiki_data_id"`
}
type Cities []City

type CitiesToUnmarshal []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	StateID     int    `json:"state_id"`
	StateCode   string `json:"state_code"`
	StateName   string `json:"state_name"`
	CountryID   int    `json:"country_id"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	WikiDataID  string `json:"wikiDataId"`
}

func CreateWorldTables(statesJSONFileName string, countriesJSONFileName string, citiesJSONFileName string, db *pgxpool.Pool) {
	jsonFileData, err := os.ReadFile(citiesJSONFileName)
	if err != nil {
		panic(err)
	}
	var cityData CitiesToUnmarshal
	err = json.Unmarshal(jsonFileData, &cityData)
	if err != nil {
		panic(err)
	}
	// dismiss error to replicate create database if not exists
	var cityDataFixed Cities
	for _, city := range cityData {
		var dumpCity City
		dumpCity.StateCode = city.StateCode
		dumpCity.CountryCode = city.CountryCode
		lat, err := strconv.ParseFloat(city.Latitude, 64)
		if err != nil {
			panic(err)
		}
		lon, err := strconv.ParseFloat(city.Longitude, 64)
		if err != nil {
			panic(err)
		}
		dumpCity.Latitude = lat
		dumpCity.Longitude = lon
		dumpCity.ID = uint32(city.ID)
		dumpCity.WikiDataID = city.WikiDataID
		dumpCity.CountryID = uint32(city.CountryID)
		dumpCity.Name = city.Name
		dumpCity.CreatedAt = time.Now()
		dumpCity.UpdatedAt = time.Now()
		dumpCity.StateID = uint32(city.StateID)
		cityDataFixed = append(cityDataFixed, dumpCity)
	}
	var countryData Countries
	var countryDataFixed CountriesInDB
	var timezones CountryTimezones
	jsonFileData, err = os.ReadFile(countriesJSONFileName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFileData, &countryData)
	if err != nil {
		panic(err)
	}
	for _, country := range countryData {
		var dumpCountry CountryInDB
		dumpCountry.ID = uint32(country.ID)
		dumpCountry.Name = country.Name
		dumpCountry.ISO2 = country.Iso2
		dumpCountry.ISO3 = country.Iso3
		dumpCountry.NumericCode = country.NumericCode
		dumpCountry.Capital = country.Capital
		dumpCountry.CurrencySymbol = country.CurrencySymbol
		dumpCountry.TLD = country.Tld
		dumpCountry.CurrencyName = country.CurrencyName
		dumpCountry.Native = country.Native
		dumpCountry.CreatedAt = time.Now()
		dumpCountry.UpdatedAt = time.Now()
		dumpCountry.Region = country.Region
		dumpCountry.Subregion = country.Subregion
		// marshal translations into json and put it into dumpCountry.Translations as a string
		translations, err := json.Marshal(country.Translations)
		if err != nil {
			panic(err)
		}
		dumpCountry.Translations = string(translations)
		dumpCountry.EmojiU = country.EmojiU
		dumpCountry.PhoneCode = country.PhoneCode
		dumpCountry.Currency = country.Currency
		dumpCountry.Emoji = country.Emoji
		for _, tz := range country.Timezones {
			tz.ID = len(timezones) + 1
			timezones = append(timezones, tz)
			dumpCountry.TimezoneID = append(dumpCountry.TimezoneID, uint32(tz.ID))
		}
		lat, err := strconv.ParseFloat(country.Latitude, 64)
		if err != nil {
			panic(err)
		}
		lon, err := strconv.ParseFloat(country.Longitude, 64)
		if err != nil {
			panic(err)
		}
		dumpCountry.Latitude = lat
		dumpCountry.Longitude = lon
		countryDataFixed = append(countryDataFixed, dumpCountry)
	}
	var stateData States
	var stateDataFixed StatesInDB
	jsonFileData, err = os.ReadFile(statesJSONFileName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFileData, &stateData)
	if err != nil {
		panic(err)
	}
	for _, state := range stateData {
		var dumpState StateInDB
		dumpState.ID = uint32(state.ID)
		dumpState.Name = state.Name
		dumpState.CountryID = uint32(state.CountryID)
		dumpState.CountryCode = state.CountryCode
		if state.Type != nil {
			dumpState.Type = state.Type.(string)
		}
		dumpState.CreatedAt = time.Now()
		dumpState.UpdatedAt = time.Now()
		lat, err := strconv.ParseFloat(state.Latitude, 64)
		if err != nil {
			lat = 0
		}
		lon, err := strconv.ParseFloat(state.Longitude, 64)
		if err != nil {
			lon = 0
		}
		dumpState.Latitude = lat
		dumpState.Longitude = lon
		stateDataFixed = append(stateDataFixed, dumpState)
	}
	if err = DropAndInsertTimezones(timezones, db); err != nil {
		panic(err)
	}
	if err = DropAndInsertCountry(countryDataFixed, db); err != nil {
		panic(err)
	}
	if err = DropAndInsertStates(stateDataFixed, db); err != nil {
		panic(err)
	}
	if err = DropAndInsertCities(cityDataFixed, db); err != nil {
		panic(err)
	}
	sqlQ := `ALTER TABLE "users"
    ADD FOREIGN KEY ("place_id") REFERENCES "places" ("id");
ALTER TABLE "users"
    ADD FOREIGN KEY ("city") REFERENCES "cities" ("id");
ALTER TABLE "users"
    ADD FOREIGN KEY ("country") REFERENCES "countries" ("id");
ALTER TABLE "users"
    ADD FOREIGN KEY ("state") REFERENCES "states" ("id");`
	to, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Exec(to, sqlQ)
	if err != nil {
		panic(err)
	}
}
func DropAndInsertStates(states StatesInDB, db *pgxpool.Pool) error {
	var rows [][]interface{}
	for _, state := range states {
		rows = append(rows, []interface{}{state.ID, state.Name, state.CountryID, state.CountryCode, state.Type, state.Latitude, state.Longitude, state.CreatedAt, state.UpdatedAt})
	}
	cols := []string{"id", "name", "country_id", "country_code", "type", "latitude", "longitude", "created_at", "updated_at"}
	_, err := db.CopyFrom(
		context.Background(),
		pgx.Identifier{"states"},
		cols,
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}
	return nil
}
func DropAndInsertTimezones(timezones CountryTimezones, db *pgxpool.Pool) error {
	cols := []string{"zone_name", "gmt_offset", "gmt_offset_name", "abbreviation", "tz_name"}
	var rows [][]interface{}
	for _, tz := range timezones {
		rows = append(rows, []interface{}{tz.ZoneName, tz.GmtOffset, tz.GmtOffsetName, tz.Abbreviation, tz.TzName})
	}
	to, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.CopyFrom(to, pgx.Identifier{"timezones"}, cols, pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}
	return nil

}
func DropAndInsertCountry(countries CountriesInDB, db *pgxpool.Pool) error {
	cols := []string{"name", "iso3", "numeric_code", "iso2", "phonecode", "capital", "currency", "currency_name", "currency_symbol", "tld", "native", "region", "subregion", "timezone_id", "translations", "latitude", "longitude", "emoji", "emojiu", "created_at", "updated_at"}
	var rows [][]interface{}
	for _, country := range countries {
		rows = append(rows, []interface{}{country.Name, country.ISO3, country.NumericCode, country.ISO2, country.PhoneCode, country.Capital, country.Currency, country.CurrencyName, country.CurrencySymbol, country.TLD, country.Native, country.Region, country.Subregion, country.TimezoneID, country.Translations, country.Latitude, country.Longitude, country.Emoji, country.EmojiU, time.Now(), time.Now()})
	}
	to, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.CopyFrom(to, pgx.Identifier{"countries"}, cols, pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}
	return nil
}

func DropAndInsertCities(cities []City, db *pgxpool.Pool) error {
	// Prepare the data for bulk insert
	cols := []string{"name", "state_id", "state_code", "country_id", "country_code", "latitude", "longitude", "created_at", "updated_at", "wiki_data_id"}
	var rows [][]interface{}
	for _, city := range cities {
		rows = append(rows, []interface{}{city.Name, city.StateID, city.StateCode, city.CountryID, city.CountryCode, city.Latitude, city.Longitude, city.CreatedAt, city.UpdatedAt, city.WikiDataID})
	}
	to, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.CopyFrom(to, pgx.Identifier{"cities"}, cols, pgx.CopyFromRows(rows))
	if err != nil {
		panic(err)
	}
	return nil
}
