package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type Flight struct {
	Id           int    `json:"_id"`
	Date        string `json:"date"`
	BoardingTime      string `json:"boardingTime"`
	Gate     string `json:"gate"`
	Airport       string `json:"airport"`
	Airport_Id    int    `json:"airport_id"`
	Traveler    string `json:"traveler"`
	Traveler_Id int    `json:"traveler_id"`
}

var flights []Flight

var jsonData string = `[
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
	},
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
	},
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
