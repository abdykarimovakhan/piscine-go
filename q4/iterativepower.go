package piscine

func IterativePower(nb int, power int) int {
	if nb < -20 || nb > 20 {
		return 0
	}
	ohn := nb
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	if power > 1 {
		for i := 2; i <= power; i++ {
			ohn = ohn * nb
		}
		return ohn
	}
	return 0
}
