package common

func Contains(strs *[]string, str *string) bool {
	if strs == nil || str == nil {
		return false
	}
	for _, v := range *strs {
		if v == *str {
			return true
		}
	}
	return false
}
