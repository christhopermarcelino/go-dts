package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

type WeatherRequest struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Weather struct {
	Id    int `json:"id"`
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// var (
// 	water int = h8HelperRand.RandomInt(1, 100)
// 	wind  int = h8HelperRand.RandomInt(1, 100)
// )

// func randData() {
// for range time.Tick(15 * time.Second) {
// water = h8HelperRand.RandomInt(1, 100)
// wind = h8HelperRand.RandomInt(1, 100)
// }
// }

func main() {
	for range time.Tick(15 * time.Second) {
		// randData()
		data := CreateData()
		resData := SendRequest(data)
		PrintData(resData)
	}
}

func getWaterStatus(data int) string {
	if data < 5 {
		return "aman"
	}
	if data >= 6 && data <= 8 {
		return "siaga"
	}
	return "bahaya"
}

func getWindStatus(data int) string {
	if data < 6 {
		return "aman"
	}
	if data >= 7 && data <= 15 {
		return "siaga"
	}
	return "bahaya"
}

func PrintData(data Weather) {
	fmt.Printf("%+v\n", data)
	fmt.Println("status water :", getWaterStatus(data.Water))
	fmt.Println("status wind :", getWindStatus(data.Wind))
}

func CreateData() WeatherRequest {
	data := WeatherRequest{
		Water: h8HelperRand.RandomInt(1, 100),
		Wind:  h8HelperRand.RandomInt(1, 100),
	}

	return data
}

func SendRequest(data WeatherRequest) Weather {
	reqJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var resBody Weather
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		log.Fatalln(err)
	}

	return resBody
}
