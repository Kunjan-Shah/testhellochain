package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/estimator/types"
	// "github.com/tendermint/tendermint/libs/log"
	"encoding/json"

	"net/http"

	"time"
	"os"
	"fmt"
	// "log"
)

var client *http.Client

func helper(url string, target *[]int) error {

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)

}

func Get() ([]int, error) {
	// get API_KEY from secret.env
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		return []int{0, 0}, fmt.Errorf("API_KEY does not exist in secret.env")
	}
	url := fmt.Sprintf("http://www.randomnumberapi.com/api/v1.0/randomredditnumber?min=%s&max=%s&count=2", apiKey, apiKey)

	var resp []int
	err := helper(url, &resp)
	if err != nil {
		return []int{0, 0}, err
	}
	return resp, nil
}

func (k msgServer) Estimate(goCtx context.Context, msg *types.MsgEstimate) (*types.MsgEstimateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	client = &http.Client{Timeout: 10 * time.Second}
	_ = ctx
	//var resp =

	var resp, err = Get()
	if err != nil {
		return nil, err
	}
	

	var apiData = types.ApiData{
		Creator:  msg.Creator,
		Start:    msg.Start,
		End:      msg.End,
		Time:     uint64(resp[0]),
		Distance: uint64(resp[1]),
	}
	k.AppendApiData(ctx, apiData)

	var apiCountMap = types.ApiCountMap{
		Creator:  msg.Creator,
		Count:    1,
	}

	valFound, isFound := k.GetApiCountMap(
		ctx,
		msg.Creator,
	)

	

	if !isFound {
		k.SetApiCountMap(ctx, apiCountMap)
	} else {
		apiCountMap.Count = valFound.Count+1
		k.SetApiCountMap(ctx, apiCountMap)
	}

	return &types.MsgEstimateResponse{Time: uint64(resp[0]), Distance: uint64(resp[1])}, nil
}

