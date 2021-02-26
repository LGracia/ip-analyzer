package models

import (
	"fmt"
	"strings"
)

type Country struct {
	ISOCode    		  string
	Name       		  string
	Languages  		  []string
	Timezones  		  []string
	Currency 		  string
	CurrencyInDollars string
	Distance          float64
	DistanceMessage   string
}

func (c *Country) Show() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Pais: " + c.Name)
	fmt.Println("ISO Code: " + c.ISOCode)
	fmt.Println("Idiomas: " + strings.Join(c.Languages[:], ","))
	fmt.Println("Moneda: " + c.CurrencyInDollars)
	fmt.Println("Hora: ")
	fmt.Println("Distancia estimada: " + c.DistanceMessage)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
}
