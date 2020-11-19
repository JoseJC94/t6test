package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type FlightRef struct {
	FlightId int    `json:"flight_id"`
	Title  string `json:"title"`
}

type Traveler struct {
	Id        int       `json:"_id"`
	Traveler string    `json:"traveler"`
	Country   string    `json:"country"`
	Founded   int       `json:"founded"`
	Genere    string    `json:"genere"`
	Flights     []FlightRef `json:"flights"`
}

var items []Traveler

var jsonData string = `[
	{
		"_id": 1,
		"traveler": "John Wiley & Sons",
		"country": "United States",
		"founded": 1807,
		"genere": "Academic",
		"flights": [
			{
				"flight_id": 1,
				"title": "Operating System Concepts"
			},
			{
				"flight_id": 2,
				"title": "Database System Concepts"
			}
		]
	},
	{
		"_id": 2,
		"traveler": "Pearson Education",
		"country": "United Kingdom",
		"founded": 1844,
		"genere": "Education",
		"flights": [
			{
				"flight_id": 3,
				"title": "Computer Networks"
			},
			{
				"flight_id": 4,
				"title": "Modern Operating Systems"
			}
		]
	}
]`

func FindItem(id int) *Traveler {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
