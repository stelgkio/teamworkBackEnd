package entities

type Customer struct {
	Name      string `csv:"first_name"` // .csv column headers
	LastName  string `csv:"last_name"`
	Email     string `csv:"email"`
	Gender    string `csv:"gender"`
	IPAndress string `csv:"ip_address"`
	Domain    string
}

type ByName []Customer

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }


type ByDomain []Customer

func (a ByDomain) Len() int           { return len(a) }
func (a ByDomain) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDomain) Less(i, j int) bool { return a[i].Domain < a[j].Domain }

type ByLastName []Customer

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].LastName < a[j].LastName }
