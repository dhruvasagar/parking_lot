package models

import (
	"strconv"
)

// Slot represents a slot in the Parking Lot
type Slot struct {
	Number  int
	State   string
	Vehicle *Vehicle
}

// IsEmpty tells if the slot is empty
func (slot *Slot) IsEmpty() bool {
	return slot.State == "empty"
}

// IsFull tells if the slot is full
func (slot *Slot) IsFull() bool {
	return slot.State == "full"
}

// Park parks a vehicle in the slot
func (slot *Slot) Park(regNo string, colour string) {
	slot.Vehicle = &Vehicle{Colour: colour, RegNo: regNo}
	slot.State = "full"
}

// Leave empties the slot
func (slot *Slot) Leave() {
	slot.Vehicle = &Vehicle{}
	slot.State = "empty"
}

// Status reports the Slot & occupying vehicle information
// It's meant to be used for full slots
func (slot *Slot) Status() []string {
	return []string{strconv.Itoa(slot.Number), slot.Vehicle.RegNo, slot.Vehicle.Colour}
}
