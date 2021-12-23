package zsort

import "sort"

func SSIndex(ss sort.StringSlice, s string) int {
	i := ss.Search(s)
	if i >= ss.Len() || ss[i] != s {
		return -1
	}
	return i
}

type StringSlice sort.StringSlice

func (ss StringSlice) Index(s string) int {
	i := sort.StringSlice(ss).Search(s)
	if i >= sort.StringSlice(ss).Len() || sort.StringSlice(ss)[i] != s {
		return -1
	}
	return i
}
