/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// HTTPLogError Invalid query performed.
type HTTPLogError struct {
	// Error code.
	Code int32 `json:"code"`
	// Error message.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewHTTPLogError instantiates a new HTTPLogError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPLogError(code int32, message string) *HTTPLogError {
	this := HTTPLogError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewHTTPLogErrorWithDefaults instantiates a new HTTPLogError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPLogErrorWithDefaults() *HTTPLogError {
	this := HTTPLogError{}
	return &this
}

// GetCode returns the Code field value
func (o *HTTPLogError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *HTTPLogError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *HTTPLogError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *HTTPLogError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *HTTPLogError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *HTTPLogError) SetMessage(v string) {
	o.Message = v
}

func (o HTTPLogError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

func (o *HTTPLogError) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Code    *int32  `json:"code"`
		Message *string `json:"message"`
	}{}
	all := struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Code == nil {
		return fmt.Errorf("Required field code missing")
	}
	if required.Message == nil {
		return fmt.Errorf("Required field message missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Code = all.Code
	o.Message = all.Message
	return nil
}

type NullableHTTPLogError struct {
	value *HTTPLogError
	isSet bool
}

func (v NullableHTTPLogError) Get() *HTTPLogError {
	return v.value
}

func (v *NullableHTTPLogError) Set(val *HTTPLogError) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPLogError) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTPLogError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPLogError(val *HTTPLogError) *NullableHTTPLogError {
	return &NullableHTTPLogError{value: val, isSet: true}
}

func (v NullableHTTPLogError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPLogError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}