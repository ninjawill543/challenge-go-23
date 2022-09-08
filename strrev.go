package piscine

func StrRev(s string) string {
	liste := ""
	for _, let := range s {
		liste = string(let) + liste
	}
	return liste
}
