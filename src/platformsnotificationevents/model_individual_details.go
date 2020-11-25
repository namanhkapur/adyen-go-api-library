/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsNotifications

import (
	"encoding/json"
)

// IndividualDetails struct for IndividualDetails
type IndividualDetails struct {
	Name *ViasName `json:"name,omitempty"`
	PersonalData *ViasPersonalData `json:"personalData,omitempty"`
}

// NewIndividualDetails instantiates a new IndividualDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualDetails() *IndividualDetails {
	this := IndividualDetails{}
	return &this
}

// NewIndividualDetailsWithDefaults instantiates a new IndividualDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualDetailsWithDefaults() *IndividualDetails {
	this := IndividualDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IndividualDetails) GetName() ViasName {
	if o == nil || o.Name == nil {
		var ret ViasName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDetails) GetNameOk() (*ViasName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndividualDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given ViasName and assigns it to the Name field.
func (o *IndividualDetails) SetName(v ViasName) {
	o.Name = &v
}

// GetPersonalData returns the PersonalData field value if set, zero value otherwise.
func (o *IndividualDetails) GetPersonalData() ViasPersonalData {
	if o == nil || o.PersonalData == nil {
		var ret ViasPersonalData
		return ret
	}
	return *o.PersonalData
}

// GetPersonalDataOk returns a tuple with the PersonalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDetails) GetPersonalDataOk() (*ViasPersonalData, bool) {
	if o == nil || o.PersonalData == nil {
		return nil, false
	}
	return o.PersonalData, true
}

// HasPersonalData returns a boolean if a field has been set.
func (o *IndividualDetails) HasPersonalData() bool {
	if o != nil && o.PersonalData != nil {
		return true
	}

	return false
}

// SetPersonalData gets a reference to the given ViasPersonalData and assigns it to the PersonalData field.
func (o *IndividualDetails) SetPersonalData(v ViasPersonalData) {
	o.PersonalData = &v
}

func (o IndividualDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PersonalData != nil {
		toSerialize["personalData"] = o.PersonalData
	}
	return json.Marshal(toSerialize)
}

type NullableIndividualDetails struct {
	value *IndividualDetails
	isSet bool
}

func (v NullableIndividualDetails) Get() *IndividualDetails {
	return v.value
}

func (v *NullableIndividualDetails) Set(val *IndividualDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualDetails(val *IndividualDetails) *NullableIndividualDetails {
	return &NullableIndividualDetails{value: val, isSet: true}
}

func (v NullableIndividualDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


