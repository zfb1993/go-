package main

func Rotate(s []int, position int) []int {
	r := s[position:]
	for i := position-1 ; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}