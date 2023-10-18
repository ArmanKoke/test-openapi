package oapi

import (
	"context"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/releaseband/test-openapi/oapi/gen"
)

func StartClient() {
	client, err := gen.NewClientWithResponses(
		":8080",
	)
	if err != nil {
		panic(err)
	}

	req := gen.BetRequest{
		Amount:           0,
		CloseRound:       nil,
		Currency:         "USD",
		PlayMode:         gen.Anonymous,
		PlayerId:         "",
		ProviderGameId:   "",
		RoundId:          "",
		SecondaryRoundId: nil,
		Secret:           "",
		SecurityToken:    "",
		SessionId:        "",
		TransactionId:    "",
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		panic(err)
	}

	httpResp, err := client.BetWithResponse(context.Background(), req)
	if err != nil {
		panic(err)
	}

	resp := new(gen.BetAPIResponse)
	err = jsoniter.Unmarshal(httpResp.Body, resp)
	if err != nil {
		panic(err)
	}

	if *resp.Success == false {
		errMsg := new(gen.IntegrationError) // could be inside Bet response
		err = jsoniter.Unmarshal(httpResp.Body, errMsg)
		if err != nil {
			panic(err)
		}
	}

	// process
}
