package customfiledsort

import (
	"sort"
	"teamworkBackEnd/entities"
)

type By func(p1, p2 *entities.Customer) bool

func (by By) Sort(customers []entities.Customer) {
	ps := &customerSorter{
		customers: customers,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type customerSorter struct {
	customers []entities.Customer
	by      func(p1, p2 *entities.Customer) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *customerSorter) Len() int {
	return len(s.customers)
}

// Swap is part of sort.Interface.
func (s *customerSorter) Swap(i, j int) {
	s.customers[i], s.customers[j] = s.customers[j], s.customers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *customerSorter) Less(i, j int) bool {
	return s.by(&s.customers[i], &s.customers[j])
}

func SortByDomainAndIpAddress(customer [] entities.Customer) {
	sort.SliceStable(customer, func(i, j int) bool {
		mi, mj := customer[i], customer[j]
		switch {
		case mi.LastName != mj.LastName:
			return mi.Domain < mj.Domain
		default:
			return mi.IPAndress < mj.IPAndress
		}
	})
}
