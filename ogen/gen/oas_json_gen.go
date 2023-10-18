// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode encodes BetOK as json.
func (s BetOK) Encode(e *jx.Encoder) {
	switch s.Type {
	case BetResponseBetOK:
		s.BetResponse.Encode(e)
	case IntegrationErrorBetOK:
		s.IntegrationError.Encode(e)
	}
}

func (s BetOK) encodeFields(e *jx.Encoder) {
	switch s.Type {
	case BetResponseBetOK:
		s.BetResponse.encodeFields(e)
	case IntegrationErrorBetOK:
		s.IntegrationError.encodeFields(e)
	}
}

// Decode decodes BetOK from json.
func (s *BetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BetOK to nil")
	}
	// Sum type fields.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			switch string(key) {
			case "referenceId":
				match := BetResponseBetOK
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			case "currency":
				match := BetResponseBetOK
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			case "errorCode":
				match := IntegrationErrorBetOK
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case BetResponseBetOK:
		if err := s.BetResponse.Decode(d); err != nil {
			return err
		}
	case IntegrationErrorBetOK:
		if err := s.IntegrationError.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s BetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BetRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BetRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Secret.Set {
			e.FieldStart("secret")
			s.Secret.Encode(e)
		}
	}
	{
		if s.SessionId.Set {
			e.FieldStart("sessionId")
			s.SessionId.Encode(e)
		}
	}
	{
		if s.SecurityToken.Set {
			e.FieldStart("securityToken")
			s.SecurityToken.Encode(e)
		}
	}
	{
		if s.PlayerId.Set {
			e.FieldStart("playerId")
			s.PlayerId.Encode(e)
		}
	}
	{
		if s.PlayMode.Set {
			e.FieldStart("playMode")
			s.PlayMode.Encode(e)
		}
	}
	{
		if s.ProviderGameId.Set {
			e.FieldStart("providerGameId")
			s.ProviderGameId.Encode(e)
		}
	}
	{
		if s.RoundId.Set {
			e.FieldStart("roundId")
			s.RoundId.Encode(e)
		}
	}
	{
		if s.SecondaryRoundId.Set {
			e.FieldStart("secondaryRoundId")
			s.SecondaryRoundId.Encode(e)
		}
	}
	{
		if s.TransactionId.Set {
			e.FieldStart("transactionId")
			s.TransactionId.Encode(e)
		}
	}
	{
		if s.Currency.Set {
			e.FieldStart("currency")
			s.Currency.Encode(e)
		}
	}
	{
		if s.Amount.Set {
			e.FieldStart("amount")
			s.Amount.Encode(e)
		}
	}
	{
		if s.CloseRound.Set {
			e.FieldStart("closeRound")
			s.CloseRound.Encode(e)
		}
	}
}

var jsonFieldsNameOfBetRequest = [12]string{
	0:  "secret",
	1:  "sessionId",
	2:  "securityToken",
	3:  "playerId",
	4:  "playMode",
	5:  "providerGameId",
	6:  "roundId",
	7:  "secondaryRoundId",
	8:  "transactionId",
	9:  "currency",
	10: "amount",
	11: "closeRound",
}

// Decode decodes BetRequest from json.
func (s *BetRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BetRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "secret":
			if err := func() error {
				s.Secret.Reset()
				if err := s.Secret.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"secret\"")
			}
		case "sessionId":
			if err := func() error {
				s.SessionId.Reset()
				if err := s.SessionId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sessionId\"")
			}
		case "securityToken":
			if err := func() error {
				s.SecurityToken.Reset()
				if err := s.SecurityToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"securityToken\"")
			}
		case "playerId":
			if err := func() error {
				s.PlayerId.Reset()
				if err := s.PlayerId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"playerId\"")
			}
		case "playMode":
			if err := func() error {
				s.PlayMode.Reset()
				if err := s.PlayMode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"playMode\"")
			}
		case "providerGameId":
			if err := func() error {
				s.ProviderGameId.Reset()
				if err := s.ProviderGameId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"providerGameId\"")
			}
		case "roundId":
			if err := func() error {
				s.RoundId.Reset()
				if err := s.RoundId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"roundId\"")
			}
		case "secondaryRoundId":
			if err := func() error {
				s.SecondaryRoundId.Reset()
				if err := s.SecondaryRoundId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"secondaryRoundId\"")
			}
		case "transactionId":
			if err := func() error {
				s.TransactionId.Reset()
				if err := s.TransactionId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"transactionId\"")
			}
		case "currency":
			if err := func() error {
				s.Currency.Reset()
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "amount":
			if err := func() error {
				s.Amount.Reset()
				if err := s.Amount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"amount\"")
			}
		case "closeRound":
			if err := func() error {
				s.CloseRound.Reset()
				if err := s.CloseRound.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"closeRound\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BetRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BetRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BetRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BetResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BetResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Success.Set {
			e.FieldStart("success")
			s.Success.Encode(e)
		}
	}
	{
		if s.ReferenceId.Set {
			e.FieldStart("referenceId")
			s.ReferenceId.Encode(e)
		}
	}
	{
		if s.Currency.Set {
			e.FieldStart("currency")
			s.Currency.Encode(e)
		}
	}
	{
		if s.Balance.Set {
			e.FieldStart("balance")
			s.Balance.Encode(e)
		}
	}
	{
		if s.CashBalance.Set {
			e.FieldStart("cashBalance")
			s.CashBalance.Encode(e)
		}
	}
	{
		if s.BonusBalance.Set {
			e.FieldStart("bonusBalance")
			s.BonusBalance.Encode(e)
		}
	}
}

var jsonFieldsNameOfBetResponse = [6]string{
	0: "success",
	1: "referenceId",
	2: "currency",
	3: "balance",
	4: "cashBalance",
	5: "bonusBalance",
}

// Decode decodes BetResponse from json.
func (s *BetResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BetResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "success":
			if err := func() error {
				s.Success.Reset()
				if err := s.Success.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"success\"")
			}
		case "referenceId":
			if err := func() error {
				s.ReferenceId.Reset()
				if err := s.ReferenceId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"referenceId\"")
			}
		case "currency":
			if err := func() error {
				s.Currency.Reset()
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "balance":
			if err := func() error {
				s.Balance.Reset()
				if err := s.Balance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"balance\"")
			}
		case "cashBalance":
			if err := func() error {
				s.CashBalance.Reset()
				if err := s.CashBalance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cashBalance\"")
			}
		case "bonusBalance":
			if err := func() error {
				s.BonusBalance.Reset()
				if err := s.BonusBalance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bonusBalance\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BetResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BetResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BetResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		if s.Code.Set {
			e.FieldStart("code")
			s.Code.Encode(e)
		}
	}
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "code",
	1: "message",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			if err := func() error {
				s.Code.Reset()
				if err := s.Code.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *IntegrationError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IntegrationError) encodeFields(e *jx.Encoder) {
	{
		if s.Success.Set {
			e.FieldStart("success")
			s.Success.Encode(e)
		}
	}
	{
		if s.ErrorCode.Set {
			e.FieldStart("errorCode")
			s.ErrorCode.Encode(e)
		}
	}
	{
		if s.Balance.Set {
			e.FieldStart("balance")
			s.Balance.Encode(e)
		}
	}
	{
		if s.CashBalance.Set {
			e.FieldStart("cashBalance")
			s.CashBalance.Encode(e)
		}
	}
	{
		if s.BonusBalance.Set {
			e.FieldStart("bonusBalance")
			s.BonusBalance.Encode(e)
		}
	}
}

var jsonFieldsNameOfIntegrationError = [5]string{
	0: "success",
	1: "errorCode",
	2: "balance",
	3: "cashBalance",
	4: "bonusBalance",
}

// Decode decodes IntegrationError from json.
func (s *IntegrationError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IntegrationError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "success":
			if err := func() error {
				s.Success.Reset()
				if err := s.Success.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"success\"")
			}
		case "errorCode":
			if err := func() error {
				s.ErrorCode.Reset()
				if err := s.ErrorCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errorCode\"")
			}
		case "balance":
			if err := func() error {
				s.Balance.Reset()
				if err := s.Balance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"balance\"")
			}
		case "cashBalance":
			if err := func() error {
				s.CashBalance.Reset()
				if err := s.CashBalance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cashBalance\"")
			}
		case "bonusBalance":
			if err := func() error {
				s.BonusBalance.Reset()
				if err := s.BonusBalance.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bonusBalance\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IntegrationError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IntegrationError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IntegrationError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes IntegrationErrorCode as json.
func (s IntegrationErrorCode) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes IntegrationErrorCode from json.
func (s *IntegrationErrorCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IntegrationErrorCode to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = IntegrationErrorCode(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s IntegrationErrorCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IntegrationErrorCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes BetRequest as json.
func (o OptBetRequest) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes BetRequest from json.
func (o *OptBetRequest) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBetRequest to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBetRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBetRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int32 as json.
func (o OptInt32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int32(int32(o.Value))
}

// Decode decodes int32 from json.
func (o *OptInt32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt32 to nil")
	}
	o.Set = true
	v, err := d.Int32()
	if err != nil {
		return err
	}
	o.Value = int32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes IntegrationErrorCode as json.
func (o OptIntegrationErrorCode) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes IntegrationErrorCode from json.
func (o *OptIntegrationErrorCode) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptIntegrationErrorCode to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptIntegrationErrorCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptIntegrationErrorCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PlayMode as json.
func (o OptPlayMode) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int8(int8(o.Value))
}

// Decode decodes PlayMode from json.
func (o *OptPlayMode) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPlayMode to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPlayMode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPlayMode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PlayMode as json.
func (s PlayMode) Encode(e *jx.Encoder) {
	e.Int8(int8(s))
}

// Decode decodes PlayMode from json.
func (s *PlayMode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PlayMode to nil")
	}
	v, err := d.Int8()
	if err != nil {
		return err
	}
	*s = PlayMode(v)

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s PlayMode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PlayMode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}