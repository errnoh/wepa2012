package main

// Create an interface that can be sorted with "sort" package.
type sortable interface {
	ID() int
}

type sortableSlice []sortable

func (s sortableSlice) Len() int {
	return len(s)
}

func (s sortableSlice) Less(i, j int) bool {
	return s[i].ID() < s[j].ID()
}

func (s sortableSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
