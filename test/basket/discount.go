package basket

type Discount interface {
	Amount(amount, discount float64) float64
	Percentage(amount, percentage float64) float64
}

type MinPriceDiscount struct {
	calculator Calculator
	minAmount  float64
}

func (m MinPriceDiscount) Amount(amount, discount float64) float64 {
	if amount < m.minAmount {
		return amount
	}

	if amount < discount {
		return amount
	}

	return float64(m.calculator.Substract(int(amount), int(discount)))
}

func (m MinPriceDiscount) Percentage(amount, percentage float64) float64 {
	if amount < m.minAmount {
		return amount
	}

	discount := m.calculator.Divide(m.calculator.Multiply(int(amount), int(percentage)), 100)

	if amount <= discount {
		return amount
	}
	return float64(m.calculator.Substract(int(amount), int(discount)))
}

func NewMinPriceDiscount(minAmount float64, calculator Calculator) Discount {
	return MinPriceDiscount{calculator: calculator, minAmount: minAmount}
}
