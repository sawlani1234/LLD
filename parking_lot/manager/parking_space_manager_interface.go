package manager

import (
	"errors"
	"math/rand"
	"solid_design/parking_lot/entity"
	"solid_design/parking_lot/strategy"
)





type ParkingSpaceManagerBase struct {
    parkingSpaces  []entity.ParkingSpace
	psFinderStrategy entity.ParkingSpaceFinderStrategy
}

func (p *ParkingSpaceManagerBase) AddParkingSpace(n int) error {
     if n <= 0 {
		return errors.New("number of parking spaces should be greater than 0")
	}

	p.parkingSpaces = append(p.parkingSpaces, make([]entity.ParkingSpace, n)...)
	return nil
}

func (p *ParkingSpaceManagerBase) RemoveParkingSpace(n int) error {
	if n <= 0 {
		return errors.New("number of parking spaces should be greater than 0")
	}

	if n > len(p.parkingSpaces) {
		return errors.New("number of parking spaces to remove is greater than available parking spaces")
	}

	p.parkingSpaces = p.parkingSpaces[:len(p.parkingSpaces)-n]
	return nil
}

func (p *ParkingSpaceManagerBase) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.psFinderStrategy.Find()
}

func (p *ParkingSpaceManagerBase) BookParkingSpace(ps entity.ParkingSpace) error {

	if ps.IsOccupied() {
		return errors.New("parking space is already occupied")
	}

	ps.SetOccupied()

	for i, val := range p.parkingSpaces {
		if val.GetID() == ps.GetID() {
			p.parkingSpaces[i] = ps
			break
		}
	}
	return p.psFinderStrategy.Fill(ps)

}

func (p *ParkingSpaceManagerBase) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaces
}

func (p *ParkingSpaceManagerBase) EmptyParkingSpace(ps entity.ParkingSpace) {
	ps.Empty()
	p.parkingSpaces[ps.GetID()] = ps
	p.psFinderStrategy.Empty(ps)
}



type TwoWheelerParkingSpaceManager struct {
	parkingSpaceManagerBase ParkingSpaceManagerBase
}

func NewTwoWheelerParkingSpaceManager() TwoWheelerParkingSpaceManager {

	ps := make([]*entity.TwoWheelerParkingSpace, 10)
	spaces := []entity.ParkingSpace{}
	for i := 0; i < 10; i++ {

		ps[i] = &entity.TwoWheelerParkingSpace{
			   ParkingSpace: entity.ParkingSpaceStruct{
				Id:                               i,
				DistanceFromElevatorAndEscalator: rand.Int() % 97,
				DistanceFromElavator:             rand.Int() % 93,
				PriceCalculationStrategy:         strategy.NewPriceCalculationStrategy("hourly"),
			},
		}
		spaces = append(spaces, ps[i])
	}

	return TwoWheelerParkingSpaceManager{
		parkingSpaceManagerBase: ParkingSpaceManagerBase{
			parkingSpaces: spaces,
			psFinderStrategy: strategy.NewNearElevatorStrategy(spaces),
		},
	}
}

func (p *TwoWheelerParkingSpaceManager) AddParkingSpace(n int) error {
     return p.parkingSpaceManagerBase.AddParkingSpace(n)
}

func (p *TwoWheelerParkingSpaceManager) RemoveParkingSpace(n int) error {
	return p.parkingSpaceManagerBase.RemoveParkingSpace(n)
}

func (p *TwoWheelerParkingSpaceManager) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.parkingSpaceManagerBase.FindParkingSpace()
}

func (p *TwoWheelerParkingSpaceManager) BookParkingSpace(ps entity.ParkingSpace) error {
	return p.parkingSpaceManagerBase.BookParkingSpace(ps)
}

func (p *TwoWheelerParkingSpaceManager) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaceManagerBase.GetParkingSpaces()
}

func (p *TwoWheelerParkingSpaceManager) EmptyParkingSpace(ps entity.ParkingSpace) {
	p.parkingSpaceManagerBase.EmptyParkingSpace(ps)
}



// TODO : complete implementation
type FourWheelerParkingSpaceManager struct {
	parkingSpaceManagerBase ParkingSpaceManagerBase
}

func NewFourWheelerParkingSpaceManager() FourWheelerParkingSpaceManager {
	return FourWheelerParkingSpaceManager{}
}

func (p *FourWheelerParkingSpaceManager) AddParkingSpace(n int) error {
     return p.parkingSpaceManagerBase.AddParkingSpace(n)
}

func (p *FourWheelerParkingSpaceManager) RemoveParkingSpace(n int) error {
	return p.parkingSpaceManagerBase.RemoveParkingSpace(n)
}

func (p *FourWheelerParkingSpaceManager) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.parkingSpaceManagerBase.FindParkingSpace()
}

func (p *FourWheelerParkingSpaceManager) BookParkingSpace(ps entity.ParkingSpace) error {
	return p.parkingSpaceManagerBase.BookParkingSpace(ps)
}

func (p *FourWheelerParkingSpaceManager) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaceManagerBase.GetParkingSpaces()
}

func (p *FourWheelerParkingSpaceManager) EmptyParkingSpace(ps entity.ParkingSpace) {
	p.parkingSpaceManagerBase.EmptyParkingSpace(ps)
}