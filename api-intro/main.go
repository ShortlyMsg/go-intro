package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ApiResponse struct {
	Mins      int
	Price     string
	CloseTime string
}

func main() {
	response, err := http.Get("https://api1.binance.com/api/v3/avgPrice?symbol=BTCUSDT")
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == 200 {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		var apiResponse ApiResponse
		json.Unmarshal(data, &apiResponse)

		fmt.Println(apiResponse)
	} else {
		fmt.Println("Status code is ", response.StatusCode)
	}
}
