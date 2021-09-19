package customfiledsort

import (
	"sort"
	"teamworkBackEnd/csvreader"
	"teamworkBackEnd/entities"
)


func SortByDomainAndIpAddress(customer [] entities.Customer) {
	sort.SliceStable(customer, func(i, j int) bool {
		mi, mj := customer[i], customer[j]
		switch {
		case mi.Domain != mj.Domain:
			return mi.Domain < mj.Domain
		default:
			return mi.IPAndress < mj.IPAndress
		}
	})
}

var ArrCustomer  [] entities.Customer

func Sorting() {
	for{
		msg := <-csvreader.MGS
		ArrCustomer =append(ArrCustomer,msg)
		SortByDomainAndIpAddress(ArrCustomer)
	}
}

