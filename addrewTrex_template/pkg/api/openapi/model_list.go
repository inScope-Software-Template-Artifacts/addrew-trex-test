/*
addrewTrex API

addrewTrex API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the List type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &List{}

// List struct for List
type List struct {
	Kind  string `json:"kind"`
	Page  int32  `json:"page"`
	Size  int32  `json:"size"`
	Total int32  `json:"total"`
}

// NewList instantiates a new List object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList(kind string, page int32, size int32, total int32) *List {
	this := List{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	return &this
}

// NewListWithDefaults instantiates a new List object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWithDefaults() *List {
	this := List{}
	return &this
}

// GetKind returns the Kind field value
func (o *List) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *List) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *List) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *List) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *List) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *List) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *List) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *List) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *List) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *List) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *List) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *List) SetTotal(v int32) {
	o.Total = v
}

func (o List) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o List) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableList struct {
	value *List
	isSet bool
}

func (v NullableList) Get() *List {
	return v.value
}

func (v *NullableList) Set(val *List) {
	v.value = val
	v.isSet = true
}

func (v NullableList) IsSet() bool {
	return v.isSet
}

func (v *NullableList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList(val *List) *NullableList {
	return &NullableList{value: val, isSet: true}
}

func (v NullableList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
