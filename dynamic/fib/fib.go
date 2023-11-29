package main

func main() {
	n := 15
	memo := make([]int, n)
	println(Fib(n))
	println(FibNaive(n, memo))
	println(FibBottomUp(n, make([]int, n+1)))
}

// 1 1 2 3 5 8 13 21 34 55
func Fib(n int) int {
	if n <= 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

/*
FibNaive Fib1'den daha hızlı çalışması için Naive Algorithm ile tasarlandı.Tekrar eden verileri bir slice içerisinde tutarak
recursive ile tekrar hesaplamasın diye bu verileri depoladık.İhtiyaç duyulduğunda ise hesaplama yapmadan slice'deki daha önceden
hesaplanan değeri döndük.Bu da gereksiz işlemleri azalttı.
*/
func FibNaive(n int, memo []int) int {
	if n <= 1 {
		return 1
	}

	if memo[n-1] != 0 {
		return memo[n-1]
	}

	res := FibNaive(n-1, memo) + FibNaive(n-2, memo)

	memo[n-1] = res

	return res
}

/*
FibBottomUp
*/
func FibBottomUp(n int, k []int) int {
	if n <= 1 {
		return 1
	}

	k[0] = 1
	k[1] = 1

	for i := 2; i <= n; i++ {
		k[i] = k[i-1] + k[i-2]
	}

	return k[n]
}
