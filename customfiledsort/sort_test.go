package customfiledsort

import (
	"teamworkBackEnd/entities"
	"testing"
)

func TestSortByDomainAndIpAddress(t *testing.T) {
	var listCustomer [] entities.Customer
	customer := entities.Customer{
		Domain : "github.io",
		LastName  : "Hernandez",
		IPAndress :"38.194.51.128",
		Name :"Mildred",
		Email : "mhernandez0@github.io",
		Gender : "Female",
	}
	listCustomer =append(listCustomer, customer)
	customer2 := entities.Customer{
		Domain : "aithub.io",
		LastName  : "Hernandez",
		IPAndress :"48.194.51.128",
		Name :"Mildred",
		Email : "mhernandez0@github.io",
		Gender : "Female",
	}
	listCustomer =append(listCustomer, customer2)

	var expectedListCustomer [] entities.Customer
	expcustomer1 := entities.Customer{
		Domain : "aithub.io",
		LastName  : "Hernandez",
		IPAndress :"48.194.51.128",
		Name :"Mildred",
		Email : "mhernandez0@github.io",
		Gender : "Female",
	}
	expectedListCustomer =append(expectedListCustomer, expcustomer1)
	expcustomer2 := entities.Customer{
		Domain : "github.io",
		LastName  : "Hernandez",
		IPAndress :"38.194.51.128",
		Name :"Mildred",
		Email : "mhernandez0@github.io",
		Gender : "Female",
	}
	expectedListCustomer =append(expectedListCustomer, expcustomer2)

	SortByDomainAndIpAddress(listCustomer)

	if expectedListCustomer[0].Domain != listCustomer[0] .Domain &&
		expectedListCustomer[0].IPAndress != listCustomer[0] .IPAndress {
		t.Fail()
	}

}

