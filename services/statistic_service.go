package services

import (
	"fmt"

	"lgracia.com/ip-analyzer/models"
	"lgracia.com/ip-analyzer/repository"
)

func saveStatistic(s *models.Statistic, c *models.Country) {
	repository.GetStatistic(s)

	if (models.Statistic{}) == *s {
		s.Sum = c.Distance
		s.Count = 1
		s.Closest = c.Distance
		s.Farthest = c.Distance
	} else {
		s.Sum = s.Sum + c.Distance
		s.Count = s.Count + 1
		if s.Closest > c.Distance {
			s.Closest = c.Distance
		}

		if s.Farthest < c.Distance {
			s.Farthest = c.Distance
		}
	}

	repository.SetStatistic(s)
}

func showStatistic(s *models.Statistic) {
	repository.GetStatistic(s)

	fmt.Println("---------STATISTICS---------")
	fmt.Println(fmt.Sprintf("Mas lejano: %f Kms", s.Farthest))
	fmt.Println(fmt.Sprintf("Mas cercano: %f Kms", s.Closest))
	avg := s.Sum / float64(s.Count)
	fmt.Println(fmt.Sprintf("Promedio: %f Kms", avg))
	fmt.Println("----------------------------")
}
