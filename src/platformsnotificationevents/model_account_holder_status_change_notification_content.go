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

// AccountHolderStatusChangeNotificationContent struct for AccountHolderStatusChangeNotificationContent
type AccountHolderStatusChangeNotificationContent struct {
	// The code of the account holder.
	AccountHolderCode string `json:"accountHolderCode"`
	// in case the account holder has not been updated, contains account holder fields, that did not pass the validation.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	NewStatus AccountHolderStatus `json:"newStatus"`
	OldStatus AccountHolderStatus `json:"oldStatus"`
	// The reason for the status change.
	Reason *string `json:"reason,omitempty"`
}

// NewAccountHolderStatusChangeNotificationContent instantiates a new AccountHolderStatusChangeNotificationContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderStatusChangeNotificationContent(accountHolderCode string, newStatus AccountHolderStatus, oldStatus AccountHolderStatus, ) *AccountHolderStatusChangeNotificationContent {
	this := AccountHolderStatusChangeNotificationContent{}
	this.AccountHolderCode = accountHolderCode
	this.NewStatus = newStatus
	this.OldStatus = oldStatus
	return &this
}

// NewAccountHolderStatusChangeNotificationContentWithDefaults instantiates a new AccountHolderStatusChangeNotificationContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderStatusChangeNotificationContentWithDefaults() *AccountHolderStatusChangeNotificationContent {
	this := AccountHolderStatusChangeNotificationContent{}
	return &this
}

// GetAccountHolderCode returns the AccountHolderCode field value
func (o *AccountHolderStatusChangeNotificationContent) GetAccountHolderCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountHolderCode
}

// GetAccountHolderCodeOk returns a tuple with the AccountHolderCode field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotificationContent) GetAccountHolderCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountHolderCode, true
}

// SetAccountHolderCode sets field value
func (o *AccountHolderStatusChangeNotificationContent) SetAccountHolderCode(v string) {
	o.AccountHolderCode = v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *AccountHolderStatusChangeNotificationContent) GetInvalidFields() []ErrorFieldType {
	if o == nil || o.InvalidFields == nil {
		var ret []ErrorFieldType
		return ret
	}
	return *o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotificationContent) GetInvalidFieldsOk() (*[]ErrorFieldType, bool) {
	if o == nil || o.InvalidFields == nil {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *AccountHolderStatusChangeNotificationContent) HasInvalidFields() bool {
	if o != nil && o.InvalidFields != nil {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []ErrorFieldType and assigns it to the InvalidFields field.
func (o *AccountHolderStatusChangeNotificationContent) SetInvalidFields(v []ErrorFieldType) {
	o.InvalidFields = &v
}

// GetNewStatus returns the NewStatus field value
func (o *AccountHolderStatusChangeNotificationContent) GetNewStatus() AccountHolderStatus {
	if o == nil  {
		var ret AccountHolderStatus
		return ret
	}

	return o.NewStatus
}

// GetNewStatusOk returns a tuple with the NewStatus field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotificationContent) GetNewStatusOk() (*AccountHolderStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewStatus, true
}

// SetNewStatus sets field value
func (o *AccountHolderStatusChangeNotificationContent) SetNewStatus(v AccountHolderStatus) {
	o.NewStatus = v
}

// GetOldStatus returns the OldStatus field value
func (o *AccountHolderStatusChangeNotificationContent) GetOldStatus() AccountHolderStatus {
	if o == nil  {
		var ret AccountHolderStatus
		return ret
	}

	return o.OldStatus
}

// GetOldStatusOk returns a tuple with the OldStatus field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotificationContent) GetOldStatusOk() (*AccountHolderStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OldStatus, true
}

// SetOldStatus sets field value
func (o *AccountHolderStatusChangeNotificationContent) SetOldStatus(v AccountHolderStatus) {
	o.OldStatus = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AccountHolderStatusChangeNotificationContent) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotificationContent) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AccountHolderStatusChangeNotificationContent) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AccountHolderStatusChangeNotificationContent) SetReason(v string) {
	o.Reason = &v
}

func (o AccountHolderStatusChangeNotificationContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountHolderCode"] = o.AccountHolderCode
	}
	if o.InvalidFields != nil {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if true {
		toSerialize["newStatus"] = o.NewStatus
	}
	if true {
		toSerialize["oldStatus"] = o.OldStatus
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableAccountHolderStatusChangeNotificationContent struct {
	value *AccountHolderStatusChangeNotificationContent
	isSet bool
}

func (v NullableAccountHolderStatusChangeNotificationContent) Get() *AccountHolderStatusChangeNotificationContent {
	return v.value
}

func (v *NullableAccountHolderStatusChangeNotificationContent) Set(val *AccountHolderStatusChangeNotificationContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderStatusChangeNotificationContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderStatusChangeNotificationContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderStatusChangeNotificationContent(val *AccountHolderStatusChangeNotificationContent) *NullableAccountHolderStatusChangeNotificationContent {
	return &NullableAccountHolderStatusChangeNotificationContent{value: val, isSet: true}
}

func (v NullableAccountHolderStatusChangeNotificationContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderStatusChangeNotificationContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


