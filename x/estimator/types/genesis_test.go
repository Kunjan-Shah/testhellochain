package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"testhellochain/x/estimator/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ApiHitsList: []types.ApiHits{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ApiHitsCount: 2,
				ApiDataList: []types.ApiData{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ApiDataCount: 2,
				ApiCountMapList: []types.ApiCountMap{
					{
						Creator: "0",
					},
					{
						Creator: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated apiHits",
			genState: &types.GenesisState{
				ApiHitsList: []types.ApiHits{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid apiHits count",
			genState: &types.GenesisState{
				ApiHitsList: []types.ApiHits{
					{
						Id: 1,
					},
				},
				ApiHitsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated apiData",
			genState: &types.GenesisState{
				ApiDataList: []types.ApiData{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid apiData count",
			genState: &types.GenesisState{
				ApiDataList: []types.ApiData{
					{
						Id: 1,
					},
				},
				ApiDataCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated apiCountMap",
			genState: &types.GenesisState{
				ApiCountMapList: []types.ApiCountMap{
					{
						Creator: "0",
					},
					{
						Creator: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
