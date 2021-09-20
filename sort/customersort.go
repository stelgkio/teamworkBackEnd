package sort

import (
	"sort"
	"teamworkBackEnd/entities"
	"teamworkBackEnd/reader"
)


func SortByDomainAndNumberOfCustomers(customer [] entities.Customer) {
	sort.SliceStable(customer, func(i, j int) bool {
		mi, mj := customer[i], customer[j]
		switch {
		case mi.Domain != mj.Domain:
			return mi.Domain < mj.Domain
		default:
			return mi.NumberOfCustomers < mj.NumberOfCustomers
		}
	})
}

var ArrCustomer  [] entities.Customer

func Sorting() {
	for{
		msg := <-reader.MGS
		ArrCustomer =append(ArrCustomer,msg)
		SortByDomainAndNumberOfCustomers(ArrCustomer)
	}
}

