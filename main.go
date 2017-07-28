package main

import (
	"fmt"
	"net/http"
	"os"

	q "google.golang.org/api/qpxexpress/v1"
)

func main() {
	fmt.Println("vim-go")
	httpClient := &http.Client{}
	service, err := q.New(httpClient)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tripSearchCall := service.Trips.Search(&q.TripsSearchRequest{
		Request: &q.TripOptionsRequest{
			Passengers: &q.PassengerCounts{
				AdultCount: 1,
			},
			SaleCountry: "US",
			Slice: []*q.SliceInput{
				&q.SliceInput{
					Date:        "2017/08/20",
					Origin:      "NYC",
					Destination: "CLT",
				}, &q.SliceInput{
					Date:        "2017/08/22",
					Origin:      "CLT",
					Destination: "NYC",
				},
			},
		},
	})

	response, err := tripSearchCall.Do()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%#v\n", response)
}
