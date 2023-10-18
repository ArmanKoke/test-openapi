package ogen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	api "github.com/releaseband/test-openapi/ogen/gen"
)

func StartClient() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client, err := api.NewClient(":8080")
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	res, err := client.Bet(context.Background(), api.NewOptBetRequest(
		api.BetRequest{
			Secret:           api.NewOptString("secret"),
			SessionId:        api.OptString{},
			SecurityToken:    api.OptString{},
			PlayerId:         api.OptString{},
			PlayMode:         api.NewOptPlayMode(api.PlayMode3),
			ProviderGameId:   api.OptString{},
			RoundId:          api.OptString{},
			SecondaryRoundId: api.OptString{},
			TransactionId:    api.OptString{},
			Currency:         api.OptString{},
			Amount:           api.OptInt64{},
			CloseRound:       api.OptBool{},
		},
	))
	if err != nil {
		return fmt.Errorf("bet: %w", err)
	}

	// error
	if res.Type == api.IntegrationErrorBetOK {
		data, err := res.IntegrationError.MarshalJSON()
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err := json.Indent(&out, data, "", "  "); err != nil {
			return err
		}

		color.New(color.FgRed).Println(out.String())
		// success
	} else if res.Type == api.BetResponseBetOK {
		data, err := res.BetResponse.MarshalJSON()
		if err != nil {
			return err
		}

		var out bytes.Buffer
		if err := json.Indent(&out, data, "", "  "); err != nil {
			return err
		}

		color.New(color.FgGreen).Println(out.String())
	}

	return nil
}
