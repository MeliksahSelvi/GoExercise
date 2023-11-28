package basket

import (
	asrt "github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {

	x, y := 3, 3

	c := Calculate{}

	actual := c.Add(x, y)
	expected := 6

	if actual != expected {
		t.Errorf("Calculate.Add(%d,%d) failed.Expected: %d, Actual: %d", x, y, expected, actual)
	}
}

func TestSubstract(t *testing.T) {
	c := Calculate{}

	//anonymous struct
	tables := []struct {
		x, y, expected int
	}{
		{2, 2, 0},
		{3, 2, 1},
	}

	for _, v := range tables {
		actual := c.Substract(v.x, v.y)
		if actual != v.expected {
			t.Errorf("Calculate.Substract(%d,%d) failed.Expected: %d, Actual: %d", v.x, v.y, v.expected, actual)
		}
	}

}

func TestMultiply(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}

	//anonymous struct
	tables := []struct {
		x, y, expected int
	}{
		{2, 2, 4},
		{3, 2, 6},
	}

	for _, v := range tables {
		actual := c.Multiply(v.x, v.y)
		assert.Equal(v.expected, actual)
	}
}

func TestDivide(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}

	//anonymous struct
	tables := []struct {
		x, y, expected int
	}{
		{4, 2, 2},
		{6, 2, 3},
	}

	for _, v := range tables {
		actual := c.Divide(v.x, v.y)
		assert.Equal(v.expected, int(actual))
	}
}
