package shuttle

import (
	"fmt"
	"math"
	"strconv"
)

type Bus struct {
	Id     int
	Offset int
}

func FindEarliestTimestamp(input []string) (int, error) {
	buses, err := extractBuses(input)
	if err != nil {
		return 0, err
	}

	product := 1

	for _, bus := range buses {
		product *= bus.Id
	}

	findInv := func(num, pp int) (int, error) {
		for n := 1; true; n++ {
			if (pp*n)%num == 1 {
				return n, nil
			}
		}

		return 0, fmt.Errorf("no inv found for num %d, pp %d", num, pp)
	}

	var sum int
	for _, bus := range buses {
		rem := bus.Id - (bus.Offset % bus.Id)
		pp := product / bus.Id
		inv, err := findInv(bus.Id, pp)
		if err != nil {
			return sum, err
		}

		sum += rem * pp * inv
	}

	return sum % product, nil
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

		buses = append(buses, Bus{Id: busId, Offset: i})
	}
	return buses, nil
}
