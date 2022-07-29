package common

func UniqueStringSlice(stringSlice []string) []string {
	var (
		unq []string
		ent = map[string]bool{}
	)
	for _, v := range stringSlice {
		if _, ok := ent[v]; !ok {
			ent[v] = true
			unq = append(unq, v)
		}
	}
	return unq
}
