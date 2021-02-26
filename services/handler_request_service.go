package services

import (
	"fmt"
	"math"
	"sync"

	"lgracia.com/ip-analyzer/models"
	"lgracia.com/ip-analyzer/repository"
)

const (
	PI float64 = 3.1415926535897
	ARGLAT float64 = -38.416097
	ARGLON float64 = -63.616672
)

func handleRequest(c *models.Country, ip string) {
	geolocationRequest(c, ip)
}

func geolocationRequest(c *models.Country, ip string) {
	var wg sync.WaitGroup
	wg.Add(2)

	res := get("https://api.ip2country.info/ip?" + ip)

	c.ISOCode = fmt.Sprintf("%v", res["countryCode3"])
	c.Name = fmt.Sprintf("%v", res["countryName"])

	_ = repository.GetCountry(c)
    if c.Distance > 0 {
        return
    }

	go func() {
		countryInfoRequest(c)
		wg.Done()
	}()

	go func() {
    	currencyInfoRequest(c)
		wg.Done()
	}()

	wg.Wait()

	repository.SetCountry(c)
}

func countryInfoRequest(c *models.Country) {
	var wg sync.WaitGroup
	wg.Add(4)

	res := get("https://restcountries.eu/rest/v2/alpha/" + c.ISOCode)

	go func() {
		for _, language := range res["languages"].([] interface{}) {
			l, _ := language.(map[string]interface{})
			c.Languages = append(c.Languages, l["name"].(string))
		}

		wg.Done()
	}()

	go func() {
		for _, currency := range res["currencies"].([] interface{}) {
			cu, _ := currency.(map[string]interface{})
			c.Currency = cu["code"].(string)
		}

		wg.Done()
	}()

	go func() {
		for _, timezone := range res["timezones"].([] interface{}) {
			c.Timezones = append(c.Timezones, timezone.(string))
		}

		wg.Done()
	}()

	go func() {
		latlng := res["latlng"].([] interface{})
		lat := latlng[0].(float64)
		lon := latlng[1].(float64)
		d := distance(lat, lon)
		c.DistanceMessage = fmt.Sprintf("%f Kms (%f, %f) a (%f, %f)", d, lat, lon, ARGLAT, ARGLON)
		c.Distance = d
		wg.Done()
	}()

	wg.Wait()
}

func currencyInfoRequest(c *models.Country) {
	res := get("https://api.currencyfreaks.com/latest?apikey=d5c083956b164c99916a6b76f510f101")

	if val, ok := res["rates"].(map[string]interface{})[c.Currency]; ok {
		c.CurrencyInDollars = fmt.Sprintf("%s (1 %s = %s USD)", c.Currency, c.Currency, val.(string))
	}
}

func distance(lat float64, lon float64) float64 {
	radlat := float64(PI * lat / 180)
	radArg := float64(PI * ARGLAT / 180)

	theta := float64(lon - ARGLON)
	radtheta := float64(PI * theta / 180)

	d := math.Sin(radlat) * math.Sin(radArg) + math.Cos(radlat) * math.Cos(radArg) * math.Cos(radtheta)

	if d > 1 {
		d = 1
	}

	d = math.Acos(d)
	d = d * 180 / PI
	d = d * 60 / 1.1515

	d = d * 1.609344

	return d
}
