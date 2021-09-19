package csvreader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"teamworkBackEnd/entities"

	"github.com/gocarina/gocsv"
)

func CsvReadfileAllLines() []entities.Customer {
	in, err := os.Open("customer.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	Customers := []entities.Customer{}

	if err := gocsv.UnmarshalFile(in, &Customers); err != nil {
		panic(err)
	}
	return Customers
}

var Done = make(chan bool)
var MGS  = make( chan entities.Customer)

func CsvReadFileByLine()  {

	csvFile, _ := os.Open("customers.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
			break
		}else if line == nil {
			break
		}
		customer := entities.Customer{
			Name: line[0],
			LastName:  line[1],
			Email:  line[2],
			Gender: line[3],
			IPAndress: line[4],
			Domain:  findDomain(line[2]),

		}
		MGS <- customer
	}
	Done <- true


}

func findDomain(email string ) string {
	if  len(email) > 0  && strings.Contains(email,"@"){
		return strings.Split(email, "@")[1]
	}
	return ""
}
