package keeper

import (
	"context"

	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"net/http"
	"testhellochain/x/testhellochain/types"
	"time"
)

var client *http.Client

type Venue struct {
	Id                  uint64  `json:"id"`
	Lat                 float64 `json:"lat"`
	Lon                 float64 `json:"lon"`
	Category            string  `json:"category"`
	Name                string  `json:"name"`
	Created_on          uint64  `json:"created_on"`
	Geolocation_degrees string  `json:"geolocation_degrees"`
}

type Venues struct {
	Venueslist []Venue `json:"venues"`
}

func GetCryptoAtmVenues() [2]string {
	url := "https://coinmap.org/api/v1/venues/"
	min := 0
	max := 10
	var index int = rand.Intn(max-min) + min

	var response Venues
	var arr [2]string

	err := GetJson(url, &response)
	if err != nil {
		arr[0] = err.Error()
		arr[1] = err.Error()
	} else {
		arr[0] = response.Venueslist[index].Category
		arr[1] = response.Venueslist[index].Name
	}
	return arr
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func (k Keeper) Callapi(goCtx context.Context, req *types.QueryCallapiRequest) (*types.QueryCallapiResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	client = &http.Client{Timeout: 10 * time.Second}
	var response_array = GetCryptoAtmVenues()
	_ = ctx

	return &types.QueryCallapiResponse{Category: response_array[0], Name: response_array[1]}, nil
}
