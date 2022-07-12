package common

func UniqueStringSlice(str []string) []string {
	var (
		unq []string
		ent = map[string]bool{}
	)
	for _, v := range str {
		if _, ok := ent[v]; !ok {
			enc[v] = true
			unq = append(unq, v)
		}
	}
	return unq
}
