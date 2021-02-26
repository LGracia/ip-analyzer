package repository

import (
	"encoding/json"

	"lgracia.com/ip-analyzer/models"
)

func SetStatistic(s *models.Statistic) {
    database, _ := NewDatabase()
    p, _ := json.Marshal(s)

    database.Client.Set(ctx, "statistic", p, 0).Err()
}

func GetStatistic(s *models.Statistic) {
	database, _ := NewDatabase()
    p, _ := database.Client.Get(ctx, "statistic").Result()

    json.Unmarshal([]byte(p), s)
}
