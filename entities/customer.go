package entities

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Customer struct {
	Name      string `csv:"first_name"` // .csv column headers
	LastName  string `csv:"last_name"`
	Email     string `csv:"email"`
	Gender    string `csv:"gender"`
	IPAndress string `csv:"ip_address"`
}

type CustomerList []Customer

func (e CustomerList) Len() int {
	return len(e)
}

func (e CustomerList) Less(i, j int) bool {
	return e[i].Name > e[j].Name
}

func (e CustomerList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
