package shuttle

import (
	"fmt"
	"math"
	"sort"
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
	origin := buses[0].Id
	seen := make(map[int]int)

	buses = sortBusesDescending(buses[1:])
	currentTimeStamp := buses[0].Id / origin * origin

	for i := 0; i < len(buses); {
		if prevTimeStamp, ok := seen[buses[i].Id]; ok && prevTimeStamp == currentTimeStamp {
			i++
			continue
		}
		fmt.Printf("timestamp: %20d, bus id: %3d\n", currentTimeStamp, buses[i].Id)
		nextTimeStamp := FindMultipleOfXGreaterThanZWhereXPlusOffsetEqualsMultipleOfY(
			origin,
			buses[i].Id,
			currentTimeStamp,
			buses[i].Offset,
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
	xMult := z / x
	getYMult := func(xMultiplier int) float64 {
		return float64(x*xMultiplier+offset) / float64(y)
	}

	for yMult := getYMult(xMult); yMult != math.Floor(yMult); yMult = getYMult(xMult) {
		xMult++
	}

	return x * xMult
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
