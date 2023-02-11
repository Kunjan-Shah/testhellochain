package keeper

import (
	"context"

	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"testhellochain/x/testhellochain/types"
	"time"
	//"math/rand"
	"bytes"
	//"log"
	//"fmt"
	"strconv"
)

func helper(url string, arr []int, target *[]int) error {
	json_data, err := json.Marshal(arr)

	if err != nil {
		return err
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if error != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)

}

func GetDriverId(arr []int) []int {
	// change if gitpod workspace changed
	url := "https://8000-yogeshgupta-testhelloch-o4z062ecd0j.ws-us72.gitpod.io/"

	var resp []int
	err := helper(url, arr, &resp)
	if err != nil {
		return []int{-1}
	}
	//fmt.Print(resp[0])
	return resp

}

func (k Keeper) Getgeolocations(goCtx context.Context, req *types.QueryGetgeolocationsRequest) (*types.QueryGetgeolocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	client = &http.Client{Timeout: 10 * time.Second}
	_ = ctx
	//var resp =
	var arr []int
	str := strings.Split(req.Did, ",")
	for _, value := range str {
		i, err := strconv.Atoi(value)
		if err != nil {
			// ... handle error
		}
		arr = append(arr, i)
	}
	var resp = GetDriverId(arr)

	//res := strings.Split(req.Did, ",")
	var res []string
	for _, value := range resp {
		res = append(res, strconv.Itoa(value))
	}
	return &types.QueryGetgeolocationsResponse{Geolocations: res}, nil
}
