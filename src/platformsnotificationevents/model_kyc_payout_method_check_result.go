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

// KYCPayoutMethodCheckResult struct for KYCPayoutMethodCheckResult
type KYCPayoutMethodCheckResult struct {
	// A list of the checks and their statuses.
	Checks *[]KYCCheckStatusData `json:"checks,omitempty"`
	// The unique ID of the payoput method to which the check applies.
	PayoutMethodCode *string `json:"payoutMethodCode,omitempty"`
}

// NewKYCPayoutMethodCheckResult instantiates a new KYCPayoutMethodCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCPayoutMethodCheckResult() *KYCPayoutMethodCheckResult {
	this := KYCPayoutMethodCheckResult{}
	return &this
}

// NewKYCPayoutMethodCheckResultWithDefaults instantiates a new KYCPayoutMethodCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCPayoutMethodCheckResultWithDefaults() *KYCPayoutMethodCheckResult {
	this := KYCPayoutMethodCheckResult{}
	return &this
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *KYCPayoutMethodCheckResult) GetChecks() []KYCCheckStatusData {
	if o == nil || o.Checks == nil {
		var ret []KYCCheckStatusData
		return ret
	}
	return *o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCPayoutMethodCheckResult) GetChecksOk() (*[]KYCCheckStatusData, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *KYCPayoutMethodCheckResult) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []KYCCheckStatusData and assigns it to the Checks field.
func (o *KYCPayoutMethodCheckResult) SetChecks(v []KYCCheckStatusData) {
	o.Checks = &v
}

// GetPayoutMethodCode returns the PayoutMethodCode field value if set, zero value otherwise.
func (o *KYCPayoutMethodCheckResult) GetPayoutMethodCode() string {
	if o == nil || o.PayoutMethodCode == nil {
		var ret string
		return ret
	}
	return *o.PayoutMethodCode
}

// GetPayoutMethodCodeOk returns a tuple with the PayoutMethodCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCPayoutMethodCheckResult) GetPayoutMethodCodeOk() (*string, bool) {
	if o == nil || o.PayoutMethodCode == nil {
		return nil, false
	}
	return o.PayoutMethodCode, true
}

// HasPayoutMethodCode returns a boolean if a field has been set.
func (o *KYCPayoutMethodCheckResult) HasPayoutMethodCode() bool {
	if o != nil && o.PayoutMethodCode != nil {
		return true
	}

	return false
}

// SetPayoutMethodCode gets a reference to the given string and assigns it to the PayoutMethodCode field.
func (o *KYCPayoutMethodCheckResult) SetPayoutMethodCode(v string) {
	o.PayoutMethodCode = &v
}

func (o KYCPayoutMethodCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	if o.PayoutMethodCode != nil {
		toSerialize["payoutMethodCode"] = o.PayoutMethodCode
	}
	return json.Marshal(toSerialize)
}

type NullableKYCPayoutMethodCheckResult struct {
	value *KYCPayoutMethodCheckResult
	isSet bool
}

func (v NullableKYCPayoutMethodCheckResult) Get() *KYCPayoutMethodCheckResult {
	return v.value
}

func (v *NullableKYCPayoutMethodCheckResult) Set(val *KYCPayoutMethodCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCPayoutMethodCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCPayoutMethodCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCPayoutMethodCheckResult(val *KYCPayoutMethodCheckResult) *NullableKYCPayoutMethodCheckResult {
	return &NullableKYCPayoutMethodCheckResult{value: val, isSet: true}
}

func (v NullableKYCPayoutMethodCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCPayoutMethodCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


