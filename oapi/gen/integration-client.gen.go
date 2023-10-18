// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package gen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BetRequest defines model for BetRequest.
type BetRequest struct {
	// Amount The bet amount made in cents
	Amount int64 `json:"amount"`

	// CloseRound closeRound = true means
	// no further request for the
	// given roundId should be
	// handled.
	CloseRound *bool `json:"closeRound,omitempty"`

	// Currency The ISO 4217 code e.g. EUR
	Currency string `json:"currency"`

	// PlayMode The play mode which can be one of:
	// 1 – real money
	// 3 – anonymous
	PlayMode PlayMode `json:"playMode"`

	// PlayerId The ID of the player’s
	// account in the game
	// provider’s system
	//
	// Uniquely identifies a player
	// account per brand.
	PlayerId string `json:"playerId"`

	// ProviderGameId The game ID
	//
	// Each game has a unique ID.
	ProviderGameId string `json:"providerGameId"`

	// RoundId The unique round ID
	RoundId string `json:"roundId"`

	// SecondaryRoundId The secondary round
	// identifier which was
	// provided for the
	// original round (used
	// for games free round /
	// bonus round feature \
	// multiple credit\debit)
	SecondaryRoundId *string `json:"secondaryRoundId,omitempty"`

	// Secret This is the shared
	// secret between the
	// game aggregator and
	// the game provider
	//
	// The secret is configured by
	// the Game aggregator admin
	// for each game provider
	// separately. See Security
	// section above.
	Secret string `json:"secret"`

	// SecurityToken The security token is
	// provided so that the
	// wallet platform can
	// uniquely identify the
	// player’s session
	//
	// The security token was
	// previously provided to the
	// Game aggregator upon
	// launching the game or
	// replaced with serverToken
	// on initGame call.
	SecurityToken string `json:"securityToken"`

	// SessionId The unique session id
	// for the new session.
	//
	// Each game launching will
	// create a new session.
	SessionId string `json:"sessionId"`

	// TransactionId The unique transaction ID of the debit
	TransactionId string `json:"transactionId"`
}

// BetAPIResponse defines model for BetResponse.
type BetAPIResponse struct {
	Balance *int64 `json:"balance,omitempty"`

	// BonusBalance The player’s account
	// bonus balance in
	// cents
	//
	// Used in regulated markets
	BonusBalance *int64 `json:"bonusBalance,omitempty"`

	// CashBalance The player’s account
	// cash balance in cents
	//
	// Used in regulated markets
	CashBalance *int64 `json:"cashBalance,omitempty"`

	// Currency The ISO 4217 code
	Currency *string `json:"currency,omitempty"`

	// ReferenceId Unique ID of the record in the game provider’s system.
	ReferenceId *string `json:"referenceId,omitempty"`

	// Success true in case the request succeeded, false otherwise
	Success *bool `json:"success,omitempty"`
}

// IntegrationError defines model for IntegrationError.
type IntegrationError struct {
	Balance int64 `json:"balance"`

	// BonusBalance The player’s account
	// bonus balance in
	// cents
	//
	// Used in regulated markets
	BonusBalance int64 `json:"bonusBalance"`

	// CashBalance The player’s account
	// cash balance in cents
	//
	// Used in regulated markets
	CashBalance int64 `json:"cashBalance"`

	// ErrorCode An error code
	// describing the reason
	// for the failure
	//
	// See the table of possible
	// error codes below.
	ErrorCode IntegrationErrorCode `json:"errorCode"`

	// Success Will be set to false in case of failure
	Success bool `json:"success"`
}

// IntegrationErrorCode An error code
// describing the reason
// for the failure
//
// See the table of possible
// error codes below.
type IntegrationErrorCode = string

// BetJSONRequestBody defines body for Bet for application/json ContentType.
type BetJSONRequestBody = BetRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// BetWithBody request with any body
	BetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Bet(ctx context.Context, body BetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) BetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Bet(ctx context.Context, body BetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewBetRequest calls the generic Bet builder with application/json body
func NewBetRequest(server string, body BetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBetRequestWithBody(server, "application/json", bodyReader)
}

// NewBetRequestWithBody generates requests for Bet with any type of body
func NewBetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/debit")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// BetWithBodyWithResponse request with any body
	BetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BetResponse, error)

	BetWithResponse(ctx context.Context, body BetJSONRequestBody, reqEditors ...RequestEditorFn) (*BetResponse, error)
}

type BetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		union json.RawMessage
	}
}

// Status returns HTTPResponse.Status
func (r BetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BetWithBodyWithResponse request with arbitrary body returning *BetResponse
func (c *ClientWithResponses) BetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BetResponse, error) {
	rsp, err := c.BetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBetResponse(rsp)
}

func (c *ClientWithResponses) BetWithResponse(ctx context.Context, body BetJSONRequestBody, reqEditors ...RequestEditorFn) (*BetResponse, error) {
	rsp, err := c.Bet(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBetResponse(rsp)
}

// ParseBetResponse parses an HTTP response from a BetWithResponse call
func ParseBetResponse(rsp *http.Response) (*BetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			union json.RawMessage
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}