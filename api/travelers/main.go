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
		"sex": "M",
		"flights": [
			{
				"_id": 1,
				"date": "11-22-2012",
				"boardingTime": "20:00",
				"gate": "D1",
				"traveler": "Tony Hawk",
				"traveler_id": 1,
				"airport": "El Dorado",
				"airport_id": 3
			},
			{
				"_id": 2,
				"date": "11-29-2012",
				"boardingTime": "20:00",
				"gate": "H1",
				"traveler": "Tony Hawk",
				"traveler_id": 1,
				"airport": "Hartsfield Jackson",
				"airport_id": 2
			},
			{
				"_id": 3,
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
		"name": "Charles Negreanu",
		"country": "United Kingdom",
		"sex": "M",
		"flights": [
			{
				"_id": -1,
				"desc": "no flights"
			}
		]
	},
	{
		"_id": 3,
		"name": "María Arias",
		"country": "Costa Rica",
		"sex": "F",
		"flights": [
			{
				"_id": 6,
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
		"name": "Rebeca Sauruer",
		"country": "United States",
		"sex": "F",
		"flights": [
			{
				"_id": 4,
				"date": "11-03-2012",
				"boardingTime": "16:00",
				"gate": "J1",
				"traveler": "Rebeca Sauruer",
				"traveler_id": 4,
				"airport": "Juan Santamaría",
				"airport_id": 4
			},
			{
				"_id": 5,
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
