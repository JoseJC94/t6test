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

type Airport struct {
	Id          int       `json:"_id"`
	Name      string    `json:"name"`
	Country string    `json:"country"`
	City      string    `json:"city"`
}

var items []Airport

var jsonData string = `[
	{
		"_id": 1,
		"name": "Los Angeles",
		"country": "United Unites",
		"city": "Los Angeles"
	},
	{
		"_id": 2,
		"name": "Hartsfield Jackson",
		"country": "United Unites",
		"city": "Atlanta"
	},
	{
		"_id": 3,
		"name": "El Dorado",
		"country": "Colombia",
		"city": "Los Angeles"
	},
	{
		"_id": 4,
		"name": "Juan Santamaría",
		"country": "Costa Rica",
		"city": "San José"
	}
]
`

func FindItem(id int) *Airport {
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
