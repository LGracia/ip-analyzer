package services

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

func get(url string) map[string]interface{} {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal( err )
	}

	data, _ := ioutil.ReadAll( res.Body )

	var jsonData map[string]interface{}
    err = json.Unmarshal([]byte(data), &jsonData)
    if err != nil {
        panic(err)
    }

	res.Body.Close()

	return jsonData
}
