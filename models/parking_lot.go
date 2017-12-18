package models

import (
	"errors"
)

// ParkingLot represents a parking lot
type ParkingLot struct {
	capacity int
	slots    []*Slot
}

// Park allocates an empty slot for parking
func (lot *ParkingLot) Park(regNo string, colour string) (*Slot, error) {
	for _, slot := range lot.slots {
		if slot.IsEmpty() {
			slot.Park(regNo, colour)
			return slot, nil
		}
	}
	return &Slot{}, errors.New("Parking Lot Full")
}

// Leave empties the given slot
func (lot *ParkingLot) Leave(slotNumber int) (*Slot, error) {
	slot := lot.slots[slotNumber-1]
	slot.Leave()
	return slot, nil
}

// Status reports the status of allocated slots
func (lot *ParkingLot) Status() [][]string {
	var statuses [][]string
	for _, slot := range lot.slots {
		if slot.IsFull() {
			statuses = append(statuses, slot.Status())
		}
	}
	return statuses
}

// RegistrationNumbersForCarsWithColour returns registration numbers of parked
// vehicles with a specific colour
func (lot *ParkingLot) RegistrationNumbersForCarsWithColour(colour string) ([]string, error) {
	var numbers []string
	for _, slot := range lot.slots {
		if slot.Vehicle.Colour == colour {
			numbers = append(numbers, slot.Vehicle.RegNo)
		}
	}
	if len(numbers) == 0 {
		return numbers, errors.New("Slot Not Found")
	}
	return numbers, nil
}

// SlotNumbersForCarsWithColour returns slot numbers of slots with vehicles of
// a specific colour
func (lot *ParkingLot) SlotNumbersForCarsWithColour(colour string) ([]int, error) {
	var numbers []int
	for _, slot := range lot.slots {
		if slot.Vehicle.Colour == colour {
			numbers = append(numbers, slot.Number)
		}
	}
	if len(numbers) == 0 {
		return numbers, errors.New("Slot Not Found")
	}
	return numbers, nil
}

// SlotNumberForRegistrationNumber returns the slot number of a slot with
// a vehicle of a given registration number
func (lot *ParkingLot) SlotNumberForRegistrationNumber(regNo string) (int, error) {
	for _, slot := range lot.slots {
		if slot.Vehicle.RegNo == regNo {
			return slot.Number, nil
		}
	}
	return 0, errors.New("Slot Not Found")
}

// CreateParkingLot instantiates a parking lot of a given capacity
func CreateParkingLot(capacity int) *ParkingLot {
	var slots []*Slot
	for i := 0; i < capacity; i++ {
		slots = append(slots, &Slot{Number: i + 1, State: "empty"})
	}
	return &ParkingLot{capacity: capacity, slots: slots}
}
