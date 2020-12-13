package shuttle

import (
	"math"
	"sort"
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
	currentTimeStamp := origin
	seen := make(map[int]int)

	buses = sortBusesDescending(buses)

	for i := 1; i < len(buses); {
		if prevTimeStamp, ok := seen[buses[i].Id]; ok && prevTimeStamp == currentTimeStamp {
			i++
			continue
		}

		nextTimeStamp := FindMultipleOfXGreaterThanZWhereXPlusOffsetEqualsMultipleOfY(
			origin,
			buses[i].Id,
			currentTimeStamp,
			buses[i].Position,
		)

		seen[buses[i].Id] = nextTimeStamp

		if nextTimeStamp != currentTimeStamp {
			currentTimeStamp = nextTimeStamp
			i = 0
		} else {
			i++
		}
	}

	return currentTimeStamp, nil
}

func sortBusesDescending(buses []Bus) []Bus {
	sort.Slice(buses, func(i, j int) bool {
		return buses[i].Id > buses[j].Id
	})

	return buses
}

func FindMultipleOfXGreaterThanZWhereXPlusOffsetEqualsMultipleOfY(x, y, z, offset int) int {
	multipleX, multipleY := z, z/y*y

	for multipleX+offset != multipleY {
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
