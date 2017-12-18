package main

import (
	"bufio"
	"log"
	"os"
	"parking_lot/models"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) == 2 {
		// Process file
		file := os.Args[1]
		fr, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fr.Close()
		scanner = bufio.NewScanner(fr)
	} else {
		// Interactive
		scanner = bufio.NewScanner(os.Stdin)
	}
	for scanner.Scan() {
		line := scanner.Text()
		cmd := strings.Split(line, " ")
		runCmd(cmd)
	}
}

var parkingLot *models.ParkingLot

func runCmd(cmdWithArgs []string) {
	cmd := cmdWithArgs[0]
	args := cmdWithArgs[1:]
	presenter := models.CreatePresenter()
	switch cmd {
	case "create_parking_lot":
		size, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		parkingLot = models.CreateParkingLot(size)
		presenter.CreateParkingLot(size)
	case "park":
		presenter.Park(parkingLot.Park(args[0], args[1]))
	case "leave":
		slotNumber, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		presenter.Leave(parkingLot.Leave(slotNumber))
	case "status":
		presenter.Status(parkingLot.Status())
	case "registration_numbers_for_cars_with_colour":
		regNos, err := parkingLot.RegistrationNumbersForCarsWithColour(args[0])
		presenter.RegistrationNumbersForCarsWithColour(regNos, err)
	case "slot_numbers_for_cars_with_colour":
		slotNos, err := parkingLot.SlotNumbersForCarsWithColour(args[0])
		presenter.SlotNumbersForCarsWithColour(slotNos, err)
	case "slot_number_for_registration_number":
		slotNo, err := parkingLot.SlotNumberForRegistrationNumber(args[0])
		presenter.SlotNumberForRegistrationNumber(slotNo, err)
	}
}
