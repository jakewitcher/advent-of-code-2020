package shuttle

import (
	"math"
	"strconv"
)

type Bus struct {
	Id       int
	Position int
}

func FindEarliestTimestamp(input []string) (int, error) {
	buses, err := extractBuses(input)
	if err != nil {
		return 0, err
	}
	origin := buses[0].Id
	timestamp := origin
	var found bool

	for !found {
		for i, bus := range buses {
			if bus.Position == 0 {
				continue
			}

			currentBusIdTimestamp := (timestamp/bus.Id + 1) * bus.Id
			if currentBusIdTimestamp == timestamp+bus.Position {
				if i+1 == len(buses) {
					found = true
				}
				continue
			}

			timestamp += origin
			break
		}
	}

	return timestamp, nil
}

func FindMultipleOfXGreaterThanZWhereXPlusDiffEqualsMultipleOfY(x, y, z, diff int) int {
	multipleX, multipleY := x, y

	for multipleX < z {
		multipleX += x
	}

	for multipleX + diff != multipleY {
		if multipleY <= multipleX {
			multipleY += y
		} else {
			multipleX += x
		}
	}

	return multipleX
}

func FindNextAvailableBus(timestamp int, input []string) (int, error) {
	buses, err := extractBuses(input)
	if err != nil {
		return 0, err
	}

	availableBusId, timeUntilArrival := 0, math.MaxInt64
	for _, bus := range buses {
		newTimeUntilArrival := (timestamp/bus.Id+1)*bus.Id - timestamp
		if newTimeUntilArrival < timeUntilArrival {
			timeUntilArrival = newTimeUntilArrival
			availableBusId = bus.Id
		}
	}

	return availableBusId * timeUntilArrival, nil
}

func extractBuses(busIds []string) ([]Bus, error) {
	buses := make([]Bus, 0)
	for i, busIdStr := range busIds {
		if busIdStr == "x" {
			continue
		}

		busId, err := strconv.Atoi(busIdStr)
		if err != nil {
			return nil, err
		}

		buses = append(buses, Bus{Id: busId, Position: i})
	}
	return buses, nil
}
