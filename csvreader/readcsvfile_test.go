package csvreader

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

