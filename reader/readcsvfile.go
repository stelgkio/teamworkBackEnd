package reader

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

func CsvReadFileAllLines() []entities.Customer {
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

func NewCSVReader(csvFile *os.File)  ( *csv.Reader , error)  {

	reader := csv.NewReader(bufio.NewReader(csvFile))
	if _, err := reader.Read(); err != nil {
		return reader,err
	}
	return reader, nil

}

func OpenFile(fileName string) (*os.File ,error){
	csvFile, err := os.Open(fileName)
	return csvFile,err
}

var Done = make(chan bool)
var MGS  = make( chan entities.Customer)

func CsvReadFileByLine(fileName string)  {

	csvFile, errFile := OpenFile(fileName)
	if errFile != nil {
		panic(errFile)
	}
 	reader ,err := NewCSVReader(csvFile)
	 if err != nil {
		 panic(err)
	 }
	defer csvFile.Close()
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
		customer := generateCustomerStruct(line)
		if validateData(customer) {
			continue
		}
		MGS <- customer
	}
	Done <- true

}
func generateCustomerStruct(line []string) entities.Customer{
	return entities.Customer{
		Name: line[0],
		LastName:  line[1],
		Email:  line[2],
		Gender: line[3],
		IPAndress: line[4],
		Domain:  findDomain(line[2]),
		NumberOfCustomers : returnNumberOfCustomersForEachDomain(findDomain(line[2])),

	}
}

func validateData(customer entities.Customer) bool  {
	if customer.Email == "email" && customer.IPAndress == "ip_address" {
		return true
	}
	return false
}

var domainMap = make(map[string]int)
func returnNumberOfCustomersForEachDomain(domain string) int {
	if val, ok := domainMap[domain]; ok {
		domainMap[domain] = val+1
	}else{
		domainMap[domain] = 1
	}
	return domainMap[domain]
}

func findDomain(email string ) string {
	if strings.Count(email, "@")>1 {
		return "Invalid Email Format"
	}
	if  len(email) > 0  && strings.Contains(email,"@"){
		return strings.Split(email, "@")[1]
	}
	return ""
}
