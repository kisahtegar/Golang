package interfaces

// customerSlice need to implement 3 method for sort.Interface (Len, Less, Swap)
type CustomerSlice []Customer

func (v CustomerSlice) Len() int {
	return len(v)
}
func (v CustomerSlice) Less(i, j int) bool {
	return v[i].Name < v[j].Name
}
func (v CustomerSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// SortKeluarga need to implement 3 method for sort.Interface (Len, Less, Swap)
type SortKeluarga []Keluarga

func (v SortKeluarga) Len() int {
	return len(v)
}
func (v SortKeluarga) Less(i, j int) bool {
	// sorting dari umur tertua.
	return v[i].GetUmur() > v[j].GetUmur()
}
func (v SortKeluarga) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
