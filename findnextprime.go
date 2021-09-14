package piscine

func FindNextPrime(nb int) int {
	if primeCheck(nb) {
		return nb
	}
	for i := nb + 1; ; i++ {
		if primeCheck(i) {
			return i
		}
	}
}

func primeCheck(n int) bool {
	isPrime := true
	if n <= 1 || (n%2 == 0 && n != 2) {
		return false
	}
	for i := 3; i <= n/2 && isPrime; i += 2 {
		if n%i == 0 {
			isPrime = false
		}
	}
	return isPrime
}
