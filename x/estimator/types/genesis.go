package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ApiHitsList:     []ApiHits{},
		ApiDataList:     []ApiData{},
		ApiCountMapList: []ApiCountMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in apiHits
	apiHitsIdMap := make(map[uint64]bool)
	apiHitsCount := gs.GetApiHitsCount()
	for _, elem := range gs.ApiHitsList {
		if _, ok := apiHitsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for apiHits")
		}
		if elem.Id >= apiHitsCount {
			return fmt.Errorf("apiHits id should be lower or equal than the last id")
		}
		apiHitsIdMap[elem.Id] = true
	}
	// Check for duplicated ID in apiData
	apiDataIdMap := make(map[uint64]bool)
	apiDataCount := gs.GetApiDataCount()
	for _, elem := range gs.ApiDataList {
		if _, ok := apiDataIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for apiData")
		}
		if elem.Id >= apiDataCount {
			return fmt.Errorf("apiData id should be lower or equal than the last id")
		}
		apiDataIdMap[elem.Id] = true
	}
	// Check for duplicated index in apiCountMap
	apiCountMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.ApiCountMapList {
		index := string(ApiCountMapKey(elem.Creator))
		if _, ok := apiCountMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for apiCountMap")
		}
		apiCountMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
