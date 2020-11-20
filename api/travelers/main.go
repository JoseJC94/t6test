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
	Name string    `json:"name"`
	Country   string    `json:"country"`
	Sex    string    `json:"sex"`
}

var items []Traveler

var jsonData string = `[
	{
		"_id": 1,
		"name": "Tony Hawk",
		"country": "United States",
		"sex": "M"
	},
	{
		"_id": 2,
		"name": "Charles Negreanu",
		"country": "United Kingdom",
		"sex": "M"
	},
	{
		"_id": 3,
		"name": "María Arias",
		"country": "Costa Rica",
		"sex": "F"
	},
	{
		"_id": 4,
		"name": "Rebeca Sauruer",
		"country": "United States",
		"sex": "F"
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
