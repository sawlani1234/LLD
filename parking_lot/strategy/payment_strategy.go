package strategy

import "solid_design/parking_lot/entity"

type CreditCard struct {
}

func (c CreditCard) Pay(amount float64) error {
	return nil
}

type Cash struct {
}

func (c Cash) Pay(amount float64) error {
	return nil
}

type UPI struct {
}

func (u UPI) Pay(amount float64) error {
	return nil
}

func NewPaymentStrategy(paymentType string) entity.PaymentStrategy {
	switch paymentType {
	case "credit_card":
		return CreditCard{}
	case "cash":
		return Cash{}
	case "upi":
		return UPI{}
	default:
		return nil
	}
}
