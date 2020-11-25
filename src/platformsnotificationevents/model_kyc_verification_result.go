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

// KYCVerificationResult struct for KYCVerificationResult
type KYCVerificationResult struct {
	AccountHolder *KYCCheckResult `json:"accountHolder,omitempty"`
	// The result(s) of the checks on the payout method(s).
	PayoutMethods *[]KYCPayoutMethodCheckResult `json:"payoutMethods,omitempty"`
	// The result(s) of the checks on the shareholder(s).
	Shareholders *[]KYCShareholderCheckResult `json:"shareholders,omitempty"`
}

// NewKYCVerificationResult instantiates a new KYCVerificationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCVerificationResult() *KYCVerificationResult {
	this := KYCVerificationResult{}
	return &this
}

// NewKYCVerificationResultWithDefaults instantiates a new KYCVerificationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCVerificationResultWithDefaults() *KYCVerificationResult {
	this := KYCVerificationResult{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *KYCVerificationResult) GetAccountHolder() KYCCheckResult {
	if o == nil || o.AccountHolder == nil {
		var ret KYCCheckResult
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCVerificationResult) GetAccountHolderOk() (*KYCCheckResult, bool) {
	if o == nil || o.AccountHolder == nil {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *KYCVerificationResult) HasAccountHolder() bool {
	if o != nil && o.AccountHolder != nil {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given KYCCheckResult and assigns it to the AccountHolder field.
func (o *KYCVerificationResult) SetAccountHolder(v KYCCheckResult) {
	o.AccountHolder = &v
}

// GetPayoutMethods returns the PayoutMethods field value if set, zero value otherwise.
func (o *KYCVerificationResult) GetPayoutMethods() []KYCPayoutMethodCheckResult {
	if o == nil || o.PayoutMethods == nil {
		var ret []KYCPayoutMethodCheckResult
		return ret
	}
	return *o.PayoutMethods
}

// GetPayoutMethodsOk returns a tuple with the PayoutMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCVerificationResult) GetPayoutMethodsOk() (*[]KYCPayoutMethodCheckResult, bool) {
	if o == nil || o.PayoutMethods == nil {
		return nil, false
	}
	return o.PayoutMethods, true
}

// HasPayoutMethods returns a boolean if a field has been set.
func (o *KYCVerificationResult) HasPayoutMethods() bool {
	if o != nil && o.PayoutMethods != nil {
		return true
	}

	return false
}

// SetPayoutMethods gets a reference to the given []KYCPayoutMethodCheckResult and assigns it to the PayoutMethods field.
func (o *KYCVerificationResult) SetPayoutMethods(v []KYCPayoutMethodCheckResult) {
	o.PayoutMethods = &v
}

// GetShareholders returns the Shareholders field value if set, zero value otherwise.
func (o *KYCVerificationResult) GetShareholders() []KYCShareholderCheckResult {
	if o == nil || o.Shareholders == nil {
		var ret []KYCShareholderCheckResult
		return ret
	}
	return *o.Shareholders
}

// GetShareholdersOk returns a tuple with the Shareholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCVerificationResult) GetShareholdersOk() (*[]KYCShareholderCheckResult, bool) {
	if o == nil || o.Shareholders == nil {
		return nil, false
	}
	return o.Shareholders, true
}

// HasShareholders returns a boolean if a field has been set.
func (o *KYCVerificationResult) HasShareholders() bool {
	if o != nil && o.Shareholders != nil {
		return true
	}

	return false
}

// SetShareholders gets a reference to the given []KYCShareholderCheckResult and assigns it to the Shareholders field.
func (o *KYCVerificationResult) SetShareholders(v []KYCShareholderCheckResult) {
	o.Shareholders = &v
}

func (o KYCVerificationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountHolder != nil {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if o.PayoutMethods != nil {
		toSerialize["payoutMethods"] = o.PayoutMethods
	}
	if o.Shareholders != nil {
		toSerialize["shareholders"] = o.Shareholders
	}
	return json.Marshal(toSerialize)
}

type NullableKYCVerificationResult struct {
	value *KYCVerificationResult
	isSet bool
}

func (v NullableKYCVerificationResult) Get() *KYCVerificationResult {
	return v.value
}

func (v *NullableKYCVerificationResult) Set(val *KYCVerificationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCVerificationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCVerificationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCVerificationResult(val *KYCVerificationResult) *NullableKYCVerificationResult {
	return &NullableKYCVerificationResult{value: val, isSet: true}
}

func (v NullableKYCVerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCVerificationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


