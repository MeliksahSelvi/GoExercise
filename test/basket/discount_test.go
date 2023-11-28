package basket

import (
	"fmt"
	asrt "github.com/stretchr/testify/assert"
	"testing"
)

type TestCalculator struct {
	add       map[string]int
	substract map[string]int
	multiply  map[string]int
	divide    map[string]float64
}

func (t TestCalculator) On(f string, x, y int, rt float64) {
	switch f {
	case "Add":
		t.add[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Substract":
		t.substract[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Multiply":
		t.multiply[fmt.Sprintf("%d%d", x, y)] = int(rt)
	case "Divide":
		t.divide[fmt.Sprintf("%d%d", x, y)] = rt

	}
}

func (t TestCalculator) Add(x, y int) int {
	return t.add[fmt.Sprintf("%d%d", x, y)]
}

func (t TestCalculator) Substract(x, y int) int {
	return t.substract[fmt.Sprintf("%d%d", x, y)]
}

func (t TestCalculator) Multiply(x, y int) int {
	return t.multiply[fmt.Sprintf("%d%d", x, y)]
}

func (t TestCalculator) Divide(x, y int) float64 {
	return t.divide[fmt.Sprintf("%d%d", x, y)]
}

func TestAmount(t *testing.T) {
	assert := asrt.New(t)
	testCalculator := TestCalculator{
		add:       map[string]int{},
		substract: map[string]int{},
		multiply:  map[string]int{},
		divide:    map[string]float64{},
	}
	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount     float64
		percentage float64
		expected   float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 120, 100},
	}

	testCalculator.On("Substract", 100, 20, 80)

	for _, v := range tables {
		actual := d.Amount(v.amount, v.percentage)
		assert.Equal(v.expected, actual)
	}
}

func TestPercentage(t *testing.T) {
	assert := asrt.New(t)
	testCalculator := TestCalculator{
		add:       map[string]int{},
		substract: map[string]int{},
		multiply:  map[string]int{},
		divide:    map[string]float64{},
	}
	d := NewMinPriceDiscount(11, testCalculator)

	tables := []struct {
		amount     float64
		percentage float64
		expected   float64
	}{
		{100, 20, 80},
		{10, 20, 10},
		{100, 101, 100},
	}

	testCalculator.On("Multiply", 100, 20, 2000)
	testCalculator.On("Multiply", 100, 101, 10100)
	testCalculator.On("Divide", 2000, 100, 20)
	testCalculator.On("Divide", 10100, 100, 101)
	testCalculator.On("Substract", 100, 20, 80)

	for _, v := range tables {
		actual := d.Percentage(v.amount, v.percentage)
		assert.Equal(v.expected, actual)
	}
}
