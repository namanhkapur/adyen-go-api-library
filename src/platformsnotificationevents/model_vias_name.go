/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// ViasName struct for ViasName
type ViasName struct {
	// The first name.
	FirstName string `json:"firstName"`
	// The gender. >The following values are permitted: `MALE`, `FEMALE`, `UNKNOWN`.
	Gender string `json:"gender"`
	// The name's infix, if applicable. >A maximum length of twenty (20) characters is imposed.
	Infix *string `json:"infix,omitempty"`
	// The last name.
	LastName string `json:"lastName"`
}

// NewViasName instantiates a new ViasName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViasName(firstName string, gender string, lastName string, ) *ViasName {
	this := ViasName{}
	this.FirstName = firstName
	this.Gender = gender
	this.LastName = lastName
	return &this
}

// NewViasNameWithDefaults instantiates a new ViasName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViasNameWithDefaults() *ViasName {
	this := ViasName{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *ViasName) GetFirstName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ViasName) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ViasName) SetFirstName(v string) {
	o.FirstName = v
}

// GetGender returns the Gender field value
func (o *ViasName) GetGender() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *ViasName) GetGenderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *ViasName) SetGender(v string) {
	o.Gender = v
}

// GetInfix returns the Infix field value if set, zero value otherwise.
func (o *ViasName) GetInfix() string {
	if o == nil || o.Infix == nil {
		var ret string
		return ret
	}
	return *o.Infix
}

// GetInfixOk returns a tuple with the Infix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViasName) GetInfixOk() (*string, bool) {
	if o == nil || o.Infix == nil {
		return nil, false
	}
	return o.Infix, true
}

// HasInfix returns a boolean if a field has been set.
func (o *ViasName) HasInfix() bool {
	if o != nil && o.Infix != nil {
		return true
	}

	return false
}

// SetInfix gets a reference to the given string and assigns it to the Infix field.
func (o *ViasName) SetInfix(v string) {
	o.Infix = &v
}

// GetLastName returns the LastName field value
func (o *ViasName) GetLastName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ViasName) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ViasName) SetLastName(v string) {
	o.LastName = v
}

func (o ViasName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["gender"] = o.Gender
	}
	if o.Infix != nil {
		toSerialize["infix"] = o.Infix
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	return json.Marshal(toSerialize)
}

type NullableViasName struct {
	value *ViasName
	isSet bool
}

func (v NullableViasName) Get() *ViasName {
	return v.value
}

func (v *NullableViasName) Set(val *ViasName) {
	v.value = val
	v.isSet = true
}

func (v NullableViasName) IsSet() bool {
	return v.isSet
}

func (v *NullableViasName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViasName(val *ViasName) *NullableViasName {
	return &NullableViasName{value: val, isSet: true}
}

func (v NullableViasName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViasName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


