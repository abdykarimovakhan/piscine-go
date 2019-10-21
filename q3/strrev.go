package piscine

func StrRev(s string) string {

	var akhan = ""
	for _, a := range s {
		akhan = string(a) + akhan
	}
	return akhan
}
