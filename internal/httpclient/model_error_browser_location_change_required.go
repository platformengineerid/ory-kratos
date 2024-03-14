/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.

API version:
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ErrorBrowserLocationChangeRequired type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorBrowserLocationChangeRequired{}

// ErrorBrowserLocationChangeRequired struct for ErrorBrowserLocationChangeRequired
type ErrorBrowserLocationChangeRequired struct {
	Error *ErrorGeneric `json:"error,omitempty"`
	// Points to where to redirect the user to next.
	RedirectBrowserTo *string `json:"redirect_browser_to,omitempty"`
}

// NewErrorBrowserLocationChangeRequired instantiates a new ErrorBrowserLocationChangeRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorBrowserLocationChangeRequired() *ErrorBrowserLocationChangeRequired {
	this := ErrorBrowserLocationChangeRequired{}
	return &this
}

// NewErrorBrowserLocationChangeRequiredWithDefaults instantiates a new ErrorBrowserLocationChangeRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorBrowserLocationChangeRequiredWithDefaults() *ErrorBrowserLocationChangeRequired {
	this := ErrorBrowserLocationChangeRequired{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorBrowserLocationChangeRequired) GetError() ErrorGeneric {
	if o == nil || IsNil(o.Error) {
		var ret ErrorGeneric
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBrowserLocationChangeRequired) GetErrorOk() (*ErrorGeneric, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorBrowserLocationChangeRequired) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorGeneric and assigns it to the Error field.
func (o *ErrorBrowserLocationChangeRequired) SetError(v ErrorGeneric) {
	o.Error = &v
}

// GetRedirectBrowserTo returns the RedirectBrowserTo field value if set, zero value otherwise.
func (o *ErrorBrowserLocationChangeRequired) GetRedirectBrowserTo() string {
	if o == nil || IsNil(o.RedirectBrowserTo) {
		var ret string
		return ret
	}
	return *o.RedirectBrowserTo
}

// GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBrowserLocationChangeRequired) GetRedirectBrowserToOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectBrowserTo) {
		return nil, false
	}
	return o.RedirectBrowserTo, true
}

// HasRedirectBrowserTo returns a boolean if a field has been set.
func (o *ErrorBrowserLocationChangeRequired) HasRedirectBrowserTo() bool {
	if o != nil && !IsNil(o.RedirectBrowserTo) {
		return true
	}

	return false
}

// SetRedirectBrowserTo gets a reference to the given string and assigns it to the RedirectBrowserTo field.
func (o *ErrorBrowserLocationChangeRequired) SetRedirectBrowserTo(v string) {
	o.RedirectBrowserTo = &v
}

func (o ErrorBrowserLocationChangeRequired) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorBrowserLocationChangeRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.RedirectBrowserTo) {
		toSerialize["redirect_browser_to"] = o.RedirectBrowserTo
	}
	return toSerialize, nil
}

type NullableErrorBrowserLocationChangeRequired struct {
	value *ErrorBrowserLocationChangeRequired
	isSet bool
}

func (v NullableErrorBrowserLocationChangeRequired) Get() *ErrorBrowserLocationChangeRequired {
	return v.value
}

func (v *NullableErrorBrowserLocationChangeRequired) Set(val *ErrorBrowserLocationChangeRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorBrowserLocationChangeRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorBrowserLocationChangeRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorBrowserLocationChangeRequired(val *ErrorBrowserLocationChangeRequired) *NullableErrorBrowserLocationChangeRequired {
	return &NullableErrorBrowserLocationChangeRequired{value: val, isSet: true}
}

func (v NullableErrorBrowserLocationChangeRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorBrowserLocationChangeRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
