package helper

import (
	"fmt"
	"os"
	"teamworkBackEnd/entities"

	"github.com/gocarina/gocsv"
)

func readCSVFile() {
	in, err := os.Open("customerTest.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	Customers := []*entities.Customer{}

	if err := gocsv.UnmarshalFile(in, &Customers); err != nil {
		panic(err)
	}
	for _, customer := range Customers {
		fmt.Println("Hello, ", customer.Name)
	}
}
