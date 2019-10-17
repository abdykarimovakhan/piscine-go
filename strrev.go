package piscine

func StrRev(s string) string {

	var asstoface = ""
	for _, a := range s {
		asstoface = string(a) + asstoface
	}
	return asstoface
}