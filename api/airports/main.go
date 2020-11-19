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
	Airport      string    `json:"airport"`
	Nationality string    `json:"nationality"`
	BirthYear   int       `json:"birth_year"`
	Fields      string    `json:"fields"`
	Flights       []FlightRef `json:"flights"`
}

var items []Airport

var jsonData string = `[
	{
		"_id": 1,
		"airport": "Abraham Silberschatz",
		"nationality": "Israelis / American",
		"birth_year": 1952,
		"fields": "Database Systems, Operating Systems",
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
		"airport": "Andrew S. Tanenbaum",
		"nationality": "Dutch / American",
		"birth_year": 1944,
		"fields": "Distributed computing, Operating Systems",
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
