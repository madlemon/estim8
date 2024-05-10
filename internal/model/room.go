package model

import (
	"fmt"
	petname "github.com/dustinkirkland/golang-petname"
	"log"
	"regexp"
	"sort"
	"strconv"
)

type Room struct {
	Estimates      map[User]*Estimate
	ResultsVisible bool
	CurrentMode    EstimationMode
}

func NewRoom(mode EstimationMode) (string, *Room) {
	newRoom := new(Room)
	newRoom.CurrentMode = mode
	newRoom.Estimates = make(map[User]*Estimate)

	roomId := petname.Generate(2, "-")
	Global.Rooms[roomId] = newRoom
	return roomId, newRoom
}

func (room Room) AddUser(user User) {
	_, userPresent := room.Estimates[user]
	if !userPresent {
		room.Estimates[user] = nil
	}
}

func (room Room) AddPointEstimate(user User, estimateValue int) {
	estimate := Estimate(estimateValue)
	room.Estimates[user] = &estimate
	log.Printf("New Estimates from %q: %d\n", user, estimate)
}

func (room Room) AddTimeEstimate(user User, estimateString string) error {
	estimateValue, err := parseWorkTime(estimateString)
	if err != nil {
		return err
	}
	estimate := Estimate(estimateValue)
	room.Estimates[user] = &estimate
	log.Printf("New Estimates from %q: %d\n", user, estimate)
	return nil
}

func (room Room) ClearEstimates() {
	for user := range room.Estimates {
		room.Estimates[user] = nil
	}
	room.ResultsVisible = false
}

func (room Room) GetEstimates() []EstimateViewModel {
	var estimatesList []EstimateSortModel
	for user, estimate := range room.Estimates {
		var sortModel = EstimateSortModel{}
		sortModel.User = user
		if estimate != nil {
			intValue := int(*estimate)
			sortModel.EstimateValue = &intValue
		}
		estimatesList = append(estimatesList, sortModel)
	}

	if room.ResultsVisible {
		sort.Slice(estimatesList, func(a, b int) bool {
			if estimatesList[a].EstimateValue == nil {
				return false
			}
			if estimatesList[b].EstimateValue == nil {
				return true
			}
			valueA := *estimatesList[a].EstimateValue
			valueB := *estimatesList[b].EstimateValue
			return valueA > valueB
		})
	} else {
		sort.Slice(estimatesList, func(a, b int) bool {
			return estimatesList[a].User < estimatesList[b].User
		})
	}

	var resultList []EstimateViewModel
	for _, e := range estimatesList {
		var viewModel = EstimateViewModel{}
		viewModel.User = e.User
		if e.EstimateValue != nil {
			if room.CurrentMode == StoryPointMode {
				displayString := strconv.Itoa(*e.EstimateValue)
				viewModel.EstimateString = &displayString
			} else if room.CurrentMode == TimeMode {
				displayString := timeEstimateToString(*e.EstimateValue)
				viewModel.EstimateString = &displayString
			}

		}
		resultList = append(resultList, viewModel)
	}

	return resultList
}

func (room Room) GetAvgEstimate() string {
	var validEstimates []Estimate
	for _, estimate := range room.Estimates {
		if estimate != nil {
			validEstimates = append(validEstimates, *estimate)
		}
	}

	if len(validEstimates) == 0 {
		return ""
	}

	total := 0
	for _, estimate := range validEstimates {
		total += int(estimate)
	}
	avg := total / len(validEstimates)
	if room.CurrentMode == StoryPointMode {
		return strconv.Itoa(avg)
	} else if room.CurrentMode == TimeMode {
		return timeEstimateToString(avg)
	}
	return ""
}

type EstimateSortModel struct {
	User          User
	EstimateValue *int
}

type EstimateViewModel struct {
	User           User
	EstimateString *string
}

type ParsingError struct {
	s string
}

func (e *ParsingError) Error() string {
	return e.s
}

func parseWorkTime(input string) (int, error) {
	re := regexp.MustCompile(`(\d+)([wdhm])`)

	matches := re.FindAllStringSubmatch(input, -1)

	var totalMinutes int
	for _, match := range matches {
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		unit := match[2]

		switch unit {
		case "w":
			totalMinutes += value * 60 * 8 * 5
		case "d":
			totalMinutes += value * 60 * 8
		case "h":
			totalMinutes += value * 60
		case "m":
			totalMinutes += value
		default:
			return 0, fmt.Errorf("unsupported unit: %s", unit)
		}
	}
	if totalMinutes == 0 {
		return 0, &ParsingError{"Could not parse a value"}
	}
	return totalMinutes, nil
}

func timeEstimateToString(estimateInMinutes int) string {
	weeks := estimateInMinutes / (60 * 8 * 5)
	days := (estimateInMinutes % (60 * 8 * 5)) / (60 * 8)
	hours := (estimateInMinutes % (60 * 8)) / 60
	mins := estimateInMinutes % 60

	result := ""
	if weeks > 0 {
		result += fmt.Sprintf("%dw ", weeks)
	}
	if days > 0 {
		result += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dh ", hours)
	}
	if mins > 0 {
		result += fmt.Sprintf("%dm", mins)
	}

	return result
}
