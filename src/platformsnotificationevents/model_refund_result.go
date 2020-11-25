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

// RefundResult struct for RefundResult
type RefundResult struct {
	OriginalTransaction Transaction `json:"originalTransaction"`
	// The reference of the refund.
	PspReference string `json:"pspReference"`
	// The response indicating if the refund has been received for processing.
	Response *string `json:"response,omitempty"`
}

// NewRefundResult instantiates a new RefundResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundResult(originalTransaction Transaction, pspReference string, ) *RefundResult {
	this := RefundResult{}
	this.OriginalTransaction = originalTransaction
	this.PspReference = pspReference
	return &this
}

// NewRefundResultWithDefaults instantiates a new RefundResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundResultWithDefaults() *RefundResult {
	this := RefundResult{}
	return &this
}

// GetOriginalTransaction returns the OriginalTransaction field value
func (o *RefundResult) GetOriginalTransaction() Transaction {
	if o == nil  {
		var ret Transaction
		return ret
	}

	return o.OriginalTransaction
}

// GetOriginalTransactionOk returns a tuple with the OriginalTransaction field value
// and a boolean to check if the value has been set.
func (o *RefundResult) GetOriginalTransactionOk() (*Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalTransaction, true
}

// SetOriginalTransaction sets field value
func (o *RefundResult) SetOriginalTransaction(v Transaction) {
	o.OriginalTransaction = v
}

// GetPspReference returns the PspReference field value
func (o *RefundResult) GetPspReference() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *RefundResult) GetPspReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *RefundResult) SetPspReference(v string) {
	o.PspReference = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *RefundResult) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundResult) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *RefundResult) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *RefundResult) SetResponse(v string) {
	o.Response = &v
}

func (o RefundResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["originalTransaction"] = o.OriginalTransaction
	}
	if true {
		toSerialize["pspReference"] = o.PspReference
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableRefundResult struct {
	value *RefundResult
	isSet bool
}

func (v NullableRefundResult) Get() *RefundResult {
	return v.value
}

func (v *NullableRefundResult) Set(val *RefundResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundResult(val *RefundResult) *NullableRefundResult {
	return &NullableRefundResult{value: val, isSet: true}
}

func (v NullableRefundResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


