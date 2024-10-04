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

// checks if the DinosaurList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DinosaurList{}

// DinosaurList struct for DinosaurList
type DinosaurList struct {
	Kind  string     `json:"kind"`
	Page  int32      `json:"page"`
	Size  int32      `json:"size"`
	Total int32      `json:"total"`
	Items []Dinosaur `json:"items"`
}

// NewDinosaurList instantiates a new DinosaurList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDinosaurList(kind string, page int32, size int32, total int32, items []Dinosaur) *DinosaurList {
	this := DinosaurList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewDinosaurListWithDefaults instantiates a new DinosaurList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDinosaurListWithDefaults() *DinosaurList {
	this := DinosaurList{}
	return &this
}

// GetKind returns the Kind field value
func (o *DinosaurList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DinosaurList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DinosaurList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *DinosaurList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DinosaurList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DinosaurList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *DinosaurList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *DinosaurList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *DinosaurList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *DinosaurList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *DinosaurList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *DinosaurList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *DinosaurList) GetItems() []Dinosaur {
	if o == nil {
		var ret []Dinosaur
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DinosaurList) GetItemsOk() ([]Dinosaur, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DinosaurList) SetItems(v []Dinosaur) {
	o.Items = v
}

func (o DinosaurList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DinosaurList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableDinosaurList struct {
	value *DinosaurList
	isSet bool
}

func (v NullableDinosaurList) Get() *DinosaurList {
	return v.value
}

func (v *NullableDinosaurList) Set(val *DinosaurList) {
	v.value = val
	v.isSet = true
}

func (v NullableDinosaurList) IsSet() bool {
	return v.isSet
}

func (v *NullableDinosaurList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDinosaurList(val *DinosaurList) *NullableDinosaurList {
	return &NullableDinosaurList{value: val, isSet: true}
}

func (v NullableDinosaurList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDinosaurList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
