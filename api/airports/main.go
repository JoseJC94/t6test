package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type FlightRef struct {
	FlightId int    `json:"flight_id"`
	Traveler  string `json:"traveler"`
	Airport  string `json:"airport"`
}

type Airport struct {
	Id          int       `json:"_id"`
	Name      string    `json:"name"`
	Country string    `json:"country"`
	City      string    `json:"city"`
	Flights       []FlightRef `json:"flights"`
}

var items []Airport

var jsonData string = `[
	{
		"_id": 1,
		"name": "Los Angeles",
		"country": "United Unites",
		"city": "Los Angeles",
		"flights": [
			{
				"flight_id": 3,
				"date": "11-12-2012",
				"boardingTime": "10:00",
				"gate": "D1",
				"traveler": "Tony Hawk",
				"traveler_id": 1,
				"airport": "Los Angeles",
				"airport_id": 1
			}
		]
	},
	{
		"_id": 2,
		"name": "Hartsfield Jackson",
		"country": "United Unites",
		"city": "Atlanta",
		"flights": [
			{
				"flight_id": 2,
				"date": "11-29-2012",
				"boardingTime": "20:00",
				"gate": "H1",
				"traveler": "Tony Hawk",
				"traveler_id": 1,
				"airport": "Hartsfield Jackson",
				"airport_id": 2
			}
		]
	},
	{
		"_id": 3,
		"name": "El Dorado",
		"country": "Colombia",
		"city": "Bogotá",
		"flights": [
			{
				"flight_id": 1,
				"date": "11-22-2012",
				"boardingTime": "20:00",
				"gate": "D1",
				"traveler": "Tony Hawk",
				"traveler_id": 1,
				"airport": "El Dorado",
				"airport_id": 3
	},
			{
				"flight_id": 6,
				"date": "11-02-2012",
				"boardingTime": "14:00",
				"gate": "E1",
				"traveler": "María Arias",
				"traveler_id": 3,
				"airport": "El Dorado",
				"airport_id": 3
			}
		]
	},
	{
		"_id": 4,
		"name": "Juan Santamaría",
		"country": "Costa Rica",
		"city": "San José",
		"flights": [
			{
				"flight_id": 4,
				"date": "11-03-2012",
				"boardingTime": "16:00",
				"gate": "J1",
				"traveler": "Rebeca Sauruer",
				"traveler_id": 4,
				"airport": "Juan Santamaría",
				"airport_id": 4
			},
			{
				"flight_id": 5,
				"date": "11-05-2012",
				"boardingTime": "18:00",
				"gate": "J1",
				"traveler": "Rebeca Sauruer",
				"traveler_id": 4,
				"airport": "Juan Santamaría",
				"airport_id": 4
			}
		]
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
