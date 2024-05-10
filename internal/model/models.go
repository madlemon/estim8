package model

import (
	"sort"
	"strconv"
)

type GlobalInMemoryState struct {
	Rooms map[string]*Room
}

func InitGlobalState() GlobalInMemoryState {
	state := GlobalInMemoryState{}
	state.Rooms = make(map[string]*Room)
	return state
}

var Global = InitGlobalState()

type Estimate int

type User string

func (e Estimate) String() string {
	return strconv.Itoa(int(e))
}

func OrderEstimates(estimateMap map[User]*Estimate, resultsVisible bool) []EstimateViewModel {
	var estimatesList []EstimateSortModel
	for user, estimate := range estimateMap {
		var sortModel = EstimateSortModel{}
		sortModel.User = user
		if estimate != nil {
			int := int(*estimate)
			sortModel.EstimateValue = &int
		}
		estimatesList = append(estimatesList, sortModel)
	}

	if resultsVisible {
		sort.Slice(estimatesList, func(a, b int) bool {
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
			string := strconv.Itoa(*e.EstimateValue)
			viewModel.EstimateString = &string
		}
		resultList = append(resultList, viewModel)
	}

	return resultList
}
