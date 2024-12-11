package Ackermann

func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if m > 0 && n == 0 {
		return Ack(m-1, 1)
	}

	return Ack(m-1, Ack(m, n-1))

}
