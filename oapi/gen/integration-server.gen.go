// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Defines values for AudioMode.
const (
	AskPlayer     AudioMode = 3
	Off           AudioMode = 2
	OnDefaultOnPC AudioMode = 1
)

// Defines values for PlayMode.
const (
	Anonymous PlayMode = 3
	RealMoney PlayMode = 1
)

// AudioMode The game audio mode which can be one of:
// 1 – on (default on PC)
// 2 – off
// 3 – ask player
// Audio mode 3 is only relevant for Mobile. On PC only
// modes 1 and 2 are valid.
// Player preferences made during the session will
// override this.
type AudioMode int8

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// PlayMode The play mode which can be one of:
// 1 – real money
// 3 – anonymous
type PlayMode int8

// LaunchGameParams defines parameters for LaunchGame.
type LaunchGameParams struct {
	// CustomerId The customer ID.
	// Will be provided to you by your account manager.
	CustomerId string `form:"customerId" json:"customerId"`

	// BrandId The ID of the customer’s brand.
	// Will be provided to you by your account manager.
	BrandId string `form:"brandId" json:"brandId"`

	// PlayMode The play mode which can be one of:
	// 1 – real money
	// 3 – anonymous
	PlayMode PlayMode `form:"playMode" json:"playMode"`

	// SecurityToken A token which will later be sent by the game server to
	// the game provider’s system to uniquely identify the
	// player’s session.
	// Mandatory parameter except for anonymous play mode.
	SecurityToken string `form:"securityToken" json:"securityToken"`

	// PlayerId The unique ID of the player’s account in the game
	// provider’s system.
	// Mandatory parameter except for anonymous play mode.
	PlayerId string `form:"playerId" json:"playerId"`

	// Nickname The player’s nickname.
	// Mandatory parameter except for anonymous play mode
	Nickname string `form:"nickname" json:"nickname"`

	// Balance The current player’s balance in cents.
	// In anonymous play mode the balance parameter will set
	// the initial balance, and when the player tries to make a
	// wager greater than his/her available balance, his/her
	// balance will be reloaded automatically to the initial
	// balance.
	Balance int64 `form:"balance" json:"balance"`

	// Currency ISO 4217 currency code.
	// For example: “EUR”, “GBP”.
	Currency string `form:"currency" json:"currency"`

	// Language ISO 639-1 language code.
	// For example: “en”, “nl”.
	Language string `form:"language" json:"language"`

	// Country ISO 3166 country code.
	// For example: “GB”, “DE”.
	Country string `form:"country" json:"country"`

	// ProviderGameId The unique game ID.
	ProviderGameId string `form:"providerGameId" json:"providerGameId"`

	// LobbyURL URL to the game aggregator’s lobby.
	LobbyURL string `form:"lobbyURL" json:"lobbyURL"`

	// Jurisdiction ISO 3166 country code.
	// Specifies the jurisdiction under which this game is
	// being played. The game will enforce specific
	// jurisdiction requirements based on this parameter.
	// For example: use “GB” to conform with UKGC
	// requirements.
	// If not provided, no specific jurisdiction requirements are
	// applied.
	Jurisdiction string `form:"jurisdiction" json:"jurisdiction"`

	// RealityCheckInterval Specifies the time interval in seconds at which the
	// reality check will appear on the screen. A zero value
	// indicates no time interval which means that the reality
	// check will never be presented (or presented just once if
	// realityCheckStartTime is provided).
	// Mandatory in case jurisdiction parameters are set to one
	// which requires reality checks.
	// If the specified jurisdiction does not require reality
	// checks, or is not provided, but this parameter is
	// provided, an error will be returned.
	RealityCheckInterval *int32 `form:"realityCheckInterval,omitempty" json:"realityCheckInterval,omitempty"`

	// RealityCheckStartTime Indicates the seconds to wait until the 1st reality check
	// should be presented. A zero value means that the reality
	// check should be presented immediately w/o delay.
	// If you don’t want to present a reality check at all, just
	// set realityCheckInterval to zero, and don’t set
	// realityCheckStartTime at all.
	// This parameter provides a way to seamlessly continue
	// previous reality checks presented during the player’s
	// game session.
	// If not provided, realityCheckInterval is used instead.
	RealityCheckStartTime *int32 `form:"realityCheckStartTime,omitempty" json:"realityCheckStartTime,omitempty"`

	// Audio The game audio mode which can be one of:
	// 1 – on (default on PC)
	// 2 – off
	// 3 – ask player
	// Audio mode 3 is only relevant for Mobile. On PC only
	// modes 1 and 2 are valid.
	// Player preferences made during the session will
	// override this.
	Audio *AudioMode `form:"audio,omitempty" json:"audio,omitempty"`

	// OriginUrl The origin url of the lobby opening the game
	OriginUrl string `form:"originUrl" json:"originUrl"`

	// MinBet The minimum bet allowed (for supported games only)
	MinBet *int64 `form:"minBet,omitempty" json:"minBet,omitempty"`

	// MaxTotalBet The max total bet allowed (for supported games only)
	MaxTotalBet *int64 `form:"maxTotalBet,omitempty" json:"maxTotalBet,omitempty"`

	// DefaultBet The default bet (for supported games only)
	DefaultBet *int64 `form:"defaultBet,omitempty" json:"defaultBet,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Launch game on provider side
	// (GET /launchGame)
	LaunchGame(w http.ResponseWriter, r *http.Request, params LaunchGameParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Launch game on provider side
// (GET /launchGame)
func (_ Unimplemented) LaunchGame(w http.ResponseWriter, r *http.Request, params LaunchGameParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// LaunchGame operation middleware
func (siw *ServerInterfaceWrapper) LaunchGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LaunchGameParams

	// ------------- Required query parameter "customerId" -------------

	if paramValue := r.URL.Query().Get("customerId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "customerId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "customerId", r.URL.Query(), &params.CustomerId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "customerId", Err: err})
		return
	}

	// ------------- Required query parameter "brandId" -------------

	if paramValue := r.URL.Query().Get("brandId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "brandId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "brandId", r.URL.Query(), &params.BrandId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "brandId", Err: err})
		return
	}

	// ------------- Required query parameter "playMode" -------------

	if paramValue := r.URL.Query().Get("playMode"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "playMode"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "playMode", r.URL.Query(), &params.PlayMode)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "playMode", Err: err})
		return
	}

	// ------------- Required query parameter "securityToken" -------------

	if paramValue := r.URL.Query().Get("securityToken"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "securityToken"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "securityToken", r.URL.Query(), &params.SecurityToken)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "securityToken", Err: err})
		return
	}

	// ------------- Required query parameter "playerId" -------------

	if paramValue := r.URL.Query().Get("playerId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "playerId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "playerId", r.URL.Query(), &params.PlayerId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "playerId", Err: err})
		return
	}

	// ------------- Required query parameter "nickname" -------------

	if paramValue := r.URL.Query().Get("nickname"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "nickname"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "nickname", r.URL.Query(), &params.Nickname)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nickname", Err: err})
		return
	}

	// ------------- Required query parameter "balance" -------------

	if paramValue := r.URL.Query().Get("balance"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "balance"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "balance", r.URL.Query(), &params.Balance)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "balance", Err: err})
		return
	}

	// ------------- Required query parameter "currency" -------------

	if paramValue := r.URL.Query().Get("currency"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "currency"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "currency", r.URL.Query(), &params.Currency)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "currency", Err: err})
		return
	}

	// ------------- Required query parameter "language" -------------

	if paramValue := r.URL.Query().Get("language"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "language"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "language", r.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "language", Err: err})
		return
	}

	// ------------- Required query parameter "country" -------------

	if paramValue := r.URL.Query().Get("country"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "country"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "country", r.URL.Query(), &params.Country)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "country", Err: err})
		return
	}

	// ------------- Required query parameter "providerGameId" -------------

	if paramValue := r.URL.Query().Get("providerGameId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "providerGameId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "providerGameId", r.URL.Query(), &params.ProviderGameId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "providerGameId", Err: err})
		return
	}

	// ------------- Required query parameter "lobbyURL" -------------

	if paramValue := r.URL.Query().Get("lobbyURL"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "lobbyURL"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "lobbyURL", r.URL.Query(), &params.LobbyURL)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "lobbyURL", Err: err})
		return
	}

	// ------------- Required query parameter "jurisdiction" -------------

	if paramValue := r.URL.Query().Get("jurisdiction"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "jurisdiction"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "jurisdiction", r.URL.Query(), &params.Jurisdiction)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "jurisdiction", Err: err})
		return
	}

	// ------------- Optional query parameter "realityCheckInterval" -------------

	err = runtime.BindQueryParameter("form", true, false, "realityCheckInterval", r.URL.Query(), &params.RealityCheckInterval)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "realityCheckInterval", Err: err})
		return
	}

	// ------------- Optional query parameter "realityCheckStartTime" -------------

	err = runtime.BindQueryParameter("form", true, false, "realityCheckStartTime", r.URL.Query(), &params.RealityCheckStartTime)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "realityCheckStartTime", Err: err})
		return
	}

	// ------------- Optional query parameter "audio" -------------

	err = runtime.BindQueryParameter("form", true, false, "audio", r.URL.Query(), &params.Audio)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "audio", Err: err})
		return
	}

	// ------------- Required query parameter "originUrl" -------------

	if paramValue := r.URL.Query().Get("originUrl"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "originUrl"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "originUrl", r.URL.Query(), &params.OriginUrl)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "originUrl", Err: err})
		return
	}

	// ------------- Optional query parameter "minBet" -------------

	err = runtime.BindQueryParameter("form", true, false, "minBet", r.URL.Query(), &params.MinBet)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "minBet", Err: err})
		return
	}

	// ------------- Optional query parameter "maxTotalBet" -------------

	err = runtime.BindQueryParameter("form", true, false, "maxTotalBet", r.URL.Query(), &params.MaxTotalBet)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "maxTotalBet", Err: err})
		return
	}

	// ------------- Optional query parameter "defaultBet" -------------

	err = runtime.BindQueryParameter("form", true, false, "defaultBet", r.URL.Query(), &params.DefaultBet)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "defaultBet", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LaunchGame(w, r, params)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/launchGame", wrapper.LaunchGame)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}