package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Presenter is an abstraction for capturing the logic of translating parking
// lot api into meaningful messages for presentation
type Presenter struct {
	Mode string
}

// IsFileMode says if the presenter is in file mode
func (presenter *Presenter) IsFileMode() bool {
	return presenter.Mode == "file"
}

// IsInteractiveMode says if the presenter is in interactive mode
func (presenter *Presenter) IsInteractiveMode() bool {
	return presenter.Mode == "interactive"
}

// SetMode allows to change the presenters mode
func (presenter *Presenter) SetMode(mode string) {
	presenter.Mode = mode
}

// CreateParkingLot reports about creating a parking lot successfully
func (presenter *Presenter) CreateParkingLot(capacity int) {
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
}

// Park reports about parking a vehicle
func (presenter *Presenter) Park(slot *Slot, err error) {
	if err != nil {
		fmt.Println("Sorry, parking lot is full")
	} else {
		fmt.Printf("Allocated slot number: %d\n", slot.Number)
	}
}

// Leave reports about leaving a vehicle from the parking lot
func (presenter *Presenter) Leave(slot *Slot, err error) {
	if err != nil {
		fmt.Println("Sorry, slot is empty")
	} else {
		fmt.Printf("Slot number %d is free\n", slot.Number)
	}
}

// Status reports the status of the parking lot
func (presenter *Presenter) Status(statuses [][]string) {
	var header [][]string
	header = append(header, []string{"Slot No.", "Registration No", "Colour"})
	statuses = append(header, statuses...)
	for _, status := range statuses {
		fmt.Println(strings.Join(status, "\t\t"))
	}
}

// RegistrationNumbersForCarsWithColour rerports the registration numbers from
// the parking lot api to get registration numbers by vehicle colour
func (presenter *Presenter) RegistrationNumbersForCarsWithColour(regNos []string, err error) {
	if err != nil {
		fmt.Println("Not found")
	} else {
		fmt.Println(strings.Join(regNos, ", "))
	}
}

// SlotNumbersForCarsWithColour reports the slot numbers from the parking lot
// api to get slot numbers by vehicle colour
func (presenter *Presenter) SlotNumbersForCarsWithColour(slotNumbers []int, err error) {
	if err != nil {
		fmt.Println("Not found")
	} else {
		var slotNumbersStrings []string
		for i := 0; i < len(slotNumbers); i++ {
			slotNumbersStrings = append(slotNumbersStrings, strconv.Itoa(slotNumbers[i]))
		}
		fmt.Println(strings.Join(slotNumbersStrings, ", "))
	}
}

// SlotNUmberForRegistrationNumber reports the slot number from the parking lot
// api to get the slot number by vehicle registration number
func (presenter *Presenter) SlotNumberForRegistrationNumber(slotNumber int, err error) {
	if err != nil {
		fmt.Println("Not found")
	} else {
		fmt.Println(slotNumber)
	}
}

// CreatePresenter intializes a presenter in file mode
func CreatePresenter() *Presenter {
	return &Presenter{Mode: "file"}
}
