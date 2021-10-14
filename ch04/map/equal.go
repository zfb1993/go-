package main

func equalMap(x,y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for i,xv := range x {
		if xv != y[i] {
			return false
		}
	}
	return true
}