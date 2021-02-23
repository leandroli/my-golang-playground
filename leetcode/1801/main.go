package main

func replaceSpace(s string) string {
	var result string
	for _, v := range s {
		if v == ' ' {
			result += "%20"
		} else {
			result += string(v)
		}
	}
	return result
}
