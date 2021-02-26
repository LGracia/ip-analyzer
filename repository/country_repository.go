package repository

import (
	"encoding/json"
    "context"

	"lgracia.com/ip-analyzer/models"
)

var ctx = context.Background()

func SetCountry(c *models.Country) {
    database, _ := NewDatabase()
    p, _ := json.Marshal(c)

    database.Client.Set(ctx, c.ISOCode, p, 0).Err()
}

func GetCountry(dest *models.Country) error {
	database, _ := NewDatabase()
    p, _ := database.Client.Get(ctx, dest.ISOCode).Result()

    json.Unmarshal([]byte(p), dest)

    return nil
}
