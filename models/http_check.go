// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPCheck HTTP Check
//
// swagger:model http_check
type HTTPCheck struct {

	// check headers
	CheckHeaders []*ReturnHeader `json:"headers"`

	// addr
	// Pattern: ^[^\s]+$
	Addr string `json:"addr,omitempty"`

	// alpn
	// Pattern: ^[^\s]+$
	Alpn string `json:"alpn,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// body log format
	BodyLogFormat string `json:"body_log_format,omitempty"`

	// check comment
	CheckComment string `json:"check_comment,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// error status
	// Enum: [L7OKC L7RSP L7STS L6RSP L4CON]
	ErrorStatus string `json:"error_status,omitempty"`

	// exclamation mark
	ExclamationMark bool `json:"exclamation_mark,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// linger
	Linger bool `json:"linger,omitempty"`

	// match
	// Pattern: ^[^\s]+$
	// Enum: [status rstatus hdr fhdr string rstring]
	Match string `json:"match,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// min recv
	MinRecv *int64 `json:"min_recv,omitempty"`

	// ok status
	// Enum: [L7OK L7OKC L6OK L4OK]
	OkStatus string `json:"ok_status,omitempty"`

	// on error
	OnError string `json:"on_error,omitempty"`

	// on success
	OnSuccess string `json:"on_success,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// port string
	PortString string `json:"port_string,omitempty"`

	// proto
	Proto string `json:"proto,omitempty"`

	// send proxy
	SendProxy bool `json:"send_proxy,omitempty"`

	// sni
	Sni string `json:"sni,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// status code
	StatusCode string `json:"status-code,omitempty"`

	// tout status
	// Enum: [L7TOUT L6TOUT L4TOUT]
	ToutStatus string `json:"tout_status,omitempty"`

	// type
	// Required: true
	// Enum: [comment connect disable-on-404 expect send send-state set-var set-var-fmt unset-var]
	Type string `json:"type"`

	// uri
	URI string `json:"uri,omitempty"`

	// uri log format
	URILogFormat string `json:"uri_log_format,omitempty"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var format
	VarFormat string `json:"var_format,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	VarScope string `json:"var_scope,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// via socks4
	ViaSocks4 bool `json:"via_socks4,omitempty"`
}

// Validate validates this http check
func (m *HTTPCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlpn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOkStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToutStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPCheck) validateCheckHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.CheckHeaders); i++ {
		if swag.IsZero(m.CheckHeaders[i]) { // not required
			continue
		}

		if m.CheckHeaders[i] != nil {
			if err := m.CheckHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HTTPCheck) validateAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.Addr) { // not required
		return nil
	}

	if err := validate.Pattern("addr", "body", string(m.Addr), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validateAlpn(formats strfmt.Registry) error {

	if swag.IsZero(m.Alpn) { // not required
		return nil
	}

	if err := validate.Pattern("alpn", "body", string(m.Alpn), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeErrorStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7OKC","L7RSP","L7STS","L6RSP","L4CON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeErrorStatusPropEnum = append(httpCheckTypeErrorStatusPropEnum, v)
	}
}

const (

	// HTTPCheckErrorStatusL7OKC captures enum value "L7OKC"
	HTTPCheckErrorStatusL7OKC string = "L7OKC"

	// HTTPCheckErrorStatusL7RSP captures enum value "L7RSP"
	HTTPCheckErrorStatusL7RSP string = "L7RSP"

	// HTTPCheckErrorStatusL7STS captures enum value "L7STS"
	HTTPCheckErrorStatusL7STS string = "L7STS"

	// HTTPCheckErrorStatusL6RSP captures enum value "L6RSP"
	HTTPCheckErrorStatusL6RSP string = "L6RSP"

	// HTTPCheckErrorStatusL4CON captures enum value "L4CON"
	HTTPCheckErrorStatusL4CON string = "L4CON"
)

// prop value enum
func (m *HTTPCheck) validateErrorStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpCheckTypeErrorStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateErrorStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorStatusEnum("error_status", "body", m.ErrorStatus); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["status","rstatus","hdr","fhdr","string","rstring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeMatchPropEnum = append(httpCheckTypeMatchPropEnum, v)
	}
}

const (

	// HTTPCheckMatchStatus captures enum value "status"
	HTTPCheckMatchStatus string = "status"

	// HTTPCheckMatchRstatus captures enum value "rstatus"
	HTTPCheckMatchRstatus string = "rstatus"

	// HTTPCheckMatchHdr captures enum value "hdr"
	HTTPCheckMatchHdr string = "hdr"

	// HTTPCheckMatchFhdr captures enum value "fhdr"
	HTTPCheckMatchFhdr string = "fhdr"

	// HTTPCheckMatchString captures enum value "string"
	HTTPCheckMatchString string = "string"

	// HTTPCheckMatchRstring captures enum value "rstring"
	HTTPCheckMatchRstring string = "rstring"
)

// prop value enum
func (m *HTTPCheck) validateMatchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpCheckTypeMatchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if err := validate.Pattern("match", "body", string(m.Match), `^[^\s]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateMatchEnum("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeOkStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7OK","L7OKC","L6OK","L4OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeOkStatusPropEnum = append(httpCheckTypeOkStatusPropEnum, v)
	}
}

const (

	// HTTPCheckOkStatusL7OK captures enum value "L7OK"
	HTTPCheckOkStatusL7OK string = "L7OK"

	// HTTPCheckOkStatusL7OKC captures enum value "L7OKC"
	HTTPCheckOkStatusL7OKC string = "L7OKC"

	// HTTPCheckOkStatusL6OK captures enum value "L6OK"
	HTTPCheckOkStatusL6OK string = "L6OK"

	// HTTPCheckOkStatusL4OK captures enum value "L4OK"
	HTTPCheckOkStatusL4OK string = "L4OK"
)

// prop value enum
func (m *HTTPCheck) validateOkStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpCheckTypeOkStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateOkStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OkStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOkStatusEnum("ok_status", "body", m.OkStatus); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeToutStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7TOUT","L6TOUT","L4TOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeToutStatusPropEnum = append(httpCheckTypeToutStatusPropEnum, v)
	}
}

const (

	// HTTPCheckToutStatusL7TOUT captures enum value "L7TOUT"
	HTTPCheckToutStatusL7TOUT string = "L7TOUT"

	// HTTPCheckToutStatusL6TOUT captures enum value "L6TOUT"
	HTTPCheckToutStatusL6TOUT string = "L6TOUT"

	// HTTPCheckToutStatusL4TOUT captures enum value "L4TOUT"
	HTTPCheckToutStatusL4TOUT string = "L4TOUT"
)

// prop value enum
func (m *HTTPCheck) validateToutStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpCheckTypeToutStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateToutStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ToutStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateToutStatusEnum("tout_status", "body", m.ToutStatus); err != nil {
		return err
	}

	return nil
}

var httpCheckTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["comment","connect","disable-on-404","expect","send","send-state","set-var","set-var-fmt","unset-var"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeTypePropEnum = append(httpCheckTypeTypePropEnum, v)
	}
}

const (

	// HTTPCheckTypeComment captures enum value "comment"
	HTTPCheckTypeComment string = "comment"

	// HTTPCheckTypeConnect captures enum value "connect"
	HTTPCheckTypeConnect string = "connect"

	// HTTPCheckTypeDisableOn404 captures enum value "disable-on-404"
	HTTPCheckTypeDisableOn404 string = "disable-on-404"

	// HTTPCheckTypeExpect captures enum value "expect"
	HTTPCheckTypeExpect string = "expect"

	// HTTPCheckTypeSend captures enum value "send"
	HTTPCheckTypeSend string = "send"

	// HTTPCheckTypeSendState captures enum value "send-state"
	HTTPCheckTypeSendState string = "send-state"

	// HTTPCheckTypeSetVar captures enum value "set-var"
	HTTPCheckTypeSetVar string = "set-var"

	// HTTPCheckTypeSetVarFmt captures enum value "set-var-fmt"
	HTTPCheckTypeSetVarFmt string = "set-var-fmt"

	// HTTPCheckTypeUnsetVar captures enum value "unset-var"
	HTTPCheckTypeUnsetVar string = "unset-var"
)

// prop value enum
func (m *HTTPCheck) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpCheckTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validateVarScope(formats strfmt.Registry) error {

	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", string(m.VarScope), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPCheck) UnmarshalBinary(b []byte) error {
	var res HTTPCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
