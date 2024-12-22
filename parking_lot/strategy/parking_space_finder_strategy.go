package strategy

import (
	"container/heap"
	"errors"
	"solid_design/parking_lot/entity"
	"solid_design/parking_lot/helper"
)


type NearElevatorStrategy struct {
	nearElevatorHeap *helper.NearElevatorHeap
}

func NewNearElevatorStrategy(parkingSpaces []entity.ParkingSpace) NearElevatorStrategy {

	nearElevatorHeap := &helper.NearElevatorHeap{}
	heap.Init(nearElevatorHeap)

	for _, val := range parkingSpaces {
		heap.Push(nearElevatorHeap, val)
	}

	return NearElevatorStrategy{
		nearElevatorHeap: nearElevatorHeap,
	}
}

func (n NearElevatorStrategy) Find() (entity.ParkingSpace, error) {

	if n.nearElevatorHeap.Len() <= 0 {
		return nil, errors.New("no parking space left")
	}
	return (*n.nearElevatorHeap)[0], nil
}

func (n NearElevatorStrategy) Fill(ps entity.ParkingSpace) error {
	x := heap.Pop(n.nearElevatorHeap).(entity.ParkingSpace)
	if x.GetID() != ps.GetID() {
		return errors.New("parking space cannot be filled")
	}
	return nil
}

func (n NearElevatorStrategy) Empty(ps entity.ParkingSpace) {
	heap.Push(n.nearElevatorHeap, ps)
}

// TODO: Below implementation is incomplete. Please complete the implementation
type NearElevatorAndNearEscalatorStrategy struct {
	nearELevatorAndNearEscalatorHeap *helper.NearELevatorAndNearEscalatorHeap
}

func NewNearElevatorAndNearEscalatorStrategy(parkingSpaces []entity.ParkingSpace) NearElevatorAndNearEscalatorStrategy {
	nearELevatorAndNearEscalatorHeap := &helper.NearELevatorAndNearEscalatorHeap{}
	heap.Init(nearELevatorAndNearEscalatorHeap)

	for _, val := range parkingSpaces {
		heap.Push(nearELevatorAndNearEscalatorHeap, val)
	}

	return NearElevatorAndNearEscalatorStrategy{
		nearELevatorAndNearEscalatorHeap: nearELevatorAndNearEscalatorHeap,
	}

}
func (n NearElevatorAndNearEscalatorStrategy) Find() (entity.ParkingSpace, error) {

	return heap.Pop(n.nearELevatorAndNearEscalatorHeap).(entity.ParkingSpace), nil
}

func (n NearElevatorAndNearEscalatorStrategy) Fill(ps entity.ParkingSpace) error {
	return nil
}

func (n NearElevatorAndNearEscalatorStrategy) Empty(ps entity.ParkingSpace) {
	return
}

type Default struct {
}

func NewDefault() Default {
	return Default{}
}

func (d Default) Find() (entity.ParkingSpace, error) {
	return nil, nil
}

func (d Default) Fill(ps entity.ParkingSpace) error {
	return nil

}

func (d Default) Empty(ps entity.ParkingSpace) {
	return
}
