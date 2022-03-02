/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleResponseRelationships Relationships of the role object returned by the API.
type RoleResponseRelationships struct {
	Permissions *RelationshipToPermissions `json:"permissions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewRoleResponseRelationships instantiates a new RoleResponseRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResponseRelationships() *RoleResponseRelationships {
	this := RoleResponseRelationships{}
	return &this
}

// NewRoleResponseRelationshipsWithDefaults instantiates a new RoleResponseRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResponseRelationshipsWithDefaults() *RoleResponseRelationships {
	this := RoleResponseRelationships{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleResponseRelationships) GetPermissions() RelationshipToPermissions {
	if o == nil || o.Permissions == nil {
		var ret RelationshipToPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponseRelationships) GetPermissionsOk() (*RelationshipToPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleResponseRelationships) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given RelationshipToPermissions and assigns it to the Permissions field.
func (o *RoleResponseRelationships) SetPermissions(v RelationshipToPermissions) {
	o.Permissions = &v
}

func (o RoleResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

func (o *RoleResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Permissions *RelationshipToPermissions `json:"permissions,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Permissions = all.Permissions
	return nil
}

type NullableRoleResponseRelationships struct {
	value *RoleResponseRelationships
	isSet bool
}

func (v NullableRoleResponseRelationships) Get() *RoleResponseRelationships {
	return v.value
}

func (v *NullableRoleResponseRelationships) Set(val *RoleResponseRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleResponseRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleResponseRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleResponseRelationships(val *RoleResponseRelationships) *NullableRoleResponseRelationships {
	return &NullableRoleResponseRelationships{value: val, isSet: true}
}

func (v NullableRoleResponseRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleResponseRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}