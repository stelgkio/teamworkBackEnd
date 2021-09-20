package reader

import (
	"teamworkBackEnd/entities"
	"testing"
)

func TestFindDomain(t *testing.T){

	result := findDomain("test@test.com")
	want := "test.com"
	if result != want {
		t.Errorf("result %q, wanted %q", result, want)
	}
}

func TestFindDomainEmpty(t *testing.T){

	result := findDomain("testtest.com")
	want := ""
	if result != want {
		t.Errorf("result %q, wanted %q", result, want)
	}
}

func TestOpenFile(t *testing.T) {
	_, err := OpenFile("customerTest.csv")
	if  err != nil {
		t.Fail()
	}
}

func TestNewCSVReader(t *testing.T){
	result, err := OpenFile("customerTest.csv")
	if  err != nil {
		t.Fail()
	}
	reader ,_ := NewCSVReader(result)
	if  reader == nil {
		t.Fail()
	}
}

func TestCsvReadFileByLine (t *testing.T) {
	expected := 2
	count :=0
	go CsvReadFileByLine("customerTest.csv")
	go func() {
		for{
			emptyCustomer := entities.Customer{}
			testmsg := <- MGS
			if testmsg != emptyCustomer{
				t.Fail()
			}
			count++

		}
		if expected != count{
			t.Fail()
		}
		complete := <-Done
		if complete != true {
			t.Fail()
		}
	}()

}

func TestValidateData(t *testing.T){
	t.Run("validate the data of the customer", func(t *testing.T) {
		expected := true
		customer := entities.Customer{
			Name: "first_name" ,
			LastName:  "last_name",
			Email: "email" ,
			Gender: "gender" ,
			IPAndress: "ip_address",
			Domain:  findDomain("email"),

		}
		if expected != validateData(customer) {
			t.Fail()
		}
	})

}

func TestReturnNumberOfCustomersForEachDomain(t *testing.T){

}

