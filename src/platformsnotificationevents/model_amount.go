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

// Amount struct for Amount
type Amount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency string `json:"currency"`
	// The amount of the transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}

// NewAmount instantiates a new Amount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmount(currency string, value int64, ) *Amount {
	this := Amount{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewAmountWithDefaults instantiates a new Amount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountWithDefaults() *Amount {
	this := Amount{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Amount) GetCurrency() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Amount) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Amount) SetCurrency(v string) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *Amount) GetValue() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Amount) GetValueOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Amount) SetValue(v int64) {
	o.Value = v
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAmount struct {
	value *Amount
	isSet bool
}

func (v NullableAmount) Get() *Amount {
	return v.value
}

func (v *NullableAmount) Set(val *Amount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmount(val *Amount) *NullableAmount {
	return &NullableAmount{value: val, isSet: true}
}

func (v NullableAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


