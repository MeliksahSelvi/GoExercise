package basket

type Calculator interface {
	Add(x, y int) int
	Substract(x, y int) int
	Multiply(x, y int) int
	Divide(x, y int) float64
}

type Calculate struct {
}

func (c Calculate) Add(x, y int) int {
	return x + y
}

func (c Calculate) Substract(x, y int) int {
	return x - y
}

func (c Calculate) Multiply(x, y int) int {
	return x * y
}

func (c Calculate) Divide(x, y int) float64 {
	return float64(x / y)
}
