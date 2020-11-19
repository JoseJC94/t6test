package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type Flight struct {
	Id           int    `json:"_id"`
	Title        string `json:"title"`
	Edition      string `json:"edition"`
	Copyright    int    `json:"copyright"`
	Language     string `json:"language"`
	Pages        int    `json:"pages"`
	Airport       string `json:"airport"`
	Airport_Id    int    `json:"airport_id"`
	Traveler    string `json:"traveler"`
	Traveler_Id int    `json:"traveler_id"`
}

var flights []Flight

var jsonData string = `[
	{
		"_id": 1,
		"title": "Operating System Concepts",
		"edition": "9th",
		"copyright": 2012,
		"language": "ENGLISH",
		"pages": 976,
		"airport": "Abraham Silberschatz",
		"airport_id": 1,
		"traveler": "John Wiley & Sons",
		"traveler_id": 1
	},
	{
		"_id": 2,
		"title": "Database System Concepts",
		"edition": "6th",
		"copyright": 2010,
		"language": "ENGLISH",
		"pages": 1376,
		"airport": "Abraham Silberschatz",
		"airport_id": 1,
		"traveler": "John Wiley & Sons",
		"traveler_id": 1
	},
	{
		"_id": 3,
		"title": "Computer Networks",
		"edition": "5th",
		"copyright": 2010,
		"language": "ENGLISH",
		"pages": 960,
		"airport": "Andrew S. Tanenbaum",
		"airport_id": 2,
		"traveler": "Pearson Education",
		"traveler_id": 2
	},
	{
		"_id": 4,
		"title": "Modern Operating Systems",
		"edition": "4th",
		"copyright": 2014,
		"language": "ENGLISH",
		"pages": 1136,
		"airport": "Andrew S. Tanenbaum",
		"airport_id": 2,
		"traveler": "Pearson Education",
		"traveler_id": 2
	}
]`

func FindFlight(id int) *Flight {
	for _, flight := range flights {
		if flight.Id == id {
			return &flight
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(flights)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			flight := FindFlight(param)
			if flight != nil {
				data, _ = json.Marshal(*flight)
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
	_ = json.Unmarshal([]byte(jsonData), &flights)
	lambda.Start(handler)
}
