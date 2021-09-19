// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package main

import (
	"fmt"
	"log"
	"teamworkBackEnd/csvreader"
	"teamworkBackEnd/customfiledsort"
	"teamworkBackEnd/entities"
)

var arrCustomer  [] entities.Customer
func sorting() {
	for{
		msg := <-csvreader.MGS
		arrCustomer =append(arrCustomer,msg)
		customfiledsort.SortByDomainAndIpAddress(arrCustomer)
	}
}

func main() {
	go csvreader.CsvReadFileByLine()
	go sorting()
	<-csvreader.Done
	fmt.Println("After calling DONE")
	log.Println(arrCustomer)
}
