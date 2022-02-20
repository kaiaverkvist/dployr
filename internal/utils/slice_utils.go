package utils

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	var us []string
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}
