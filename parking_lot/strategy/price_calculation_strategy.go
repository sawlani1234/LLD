package strategy

import (
	"solid_design/parking_lot/entity"
	"time"
)



type FixedPrice struct {
	price float64
}

func (f FixedPrice) CalculatePrice(ticket entity.Ticket) float64 {
	return f.price
}

type HourlyPrice struct {
}

func NewHourlyPrice() HourlyPrice {
	return HourlyPrice{}
}

func (h HourlyPrice) CalculatePrice(ticket entity.Ticket) float64 {
	return time.Now().Sub(ticket.GetEntryTime()).Seconds() * 10
}

type MinutePrice struct {
}

func (m MinutePrice) CalculatePrice(ticket entity.Ticket) float64 {
	return time.Now().Sub(ticket.GetEntryTime()).Minutes() * 0.2
}

func NewPriceCalculationStrategy(priceType string) entity.PriceCalculationStrategy {
	switch priceType {
	case "fixed":
		return FixedPrice{price: 100}
	case "hourly":
		return HourlyPrice{}
	case "minute":
		return MinutePrice{}
	default:
		return nil
	}
}
