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

// CloseAccountResponse struct for CloseAccountResponse
type CloseAccountResponse struct {
	// The account code of the account that is closed.
	AccountCode *string `json:"accountCode,omitempty"`
	// Contains field validation errors that would prevent requests from being processed.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	// The reference of a request. Can be used to uniquely identify the request.
	PspReference *string `json:"pspReference,omitempty"`
	// The result code.
	ResultCode *string `json:"resultCode,omitempty"`
	// The new status of the account. >Permitted values: `Active`, `Inactive`, `Suspended`, `Closed`.
	Status string `json:"status"`
}

// NewCloseAccountResponse instantiates a new CloseAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseAccountResponse(status string, ) *CloseAccountResponse {
	this := CloseAccountResponse{}
	this.Status = status
	return &this
}

// NewCloseAccountResponseWithDefaults instantiates a new CloseAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseAccountResponseWithDefaults() *CloseAccountResponse {
	this := CloseAccountResponse{}
	return &this
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *CloseAccountResponse) GetAccountCode() string {
	if o == nil || o.AccountCode == nil {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccountResponse) GetAccountCodeOk() (*string, bool) {
	if o == nil || o.AccountCode == nil {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCode returns a boolean if a field has been set.
func (o *CloseAccountResponse) HasAccountCode() bool {
	if o != nil && o.AccountCode != nil {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *CloseAccountResponse) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *CloseAccountResponse) GetInvalidFields() []ErrorFieldType {
	if o == nil || o.InvalidFields == nil {
		var ret []ErrorFieldType
		return ret
	}
	return *o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccountResponse) GetInvalidFieldsOk() (*[]ErrorFieldType, bool) {
	if o == nil || o.InvalidFields == nil {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *CloseAccountResponse) HasInvalidFields() bool {
	if o != nil && o.InvalidFields != nil {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []ErrorFieldType and assigns it to the InvalidFields field.
func (o *CloseAccountResponse) SetInvalidFields(v []ErrorFieldType) {
	o.InvalidFields = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *CloseAccountResponse) GetPspReference() string {
	if o == nil || o.PspReference == nil {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccountResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || o.PspReference == nil {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *CloseAccountResponse) HasPspReference() bool {
	if o != nil && o.PspReference != nil {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *CloseAccountResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *CloseAccountResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloseAccountResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *CloseAccountResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *CloseAccountResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetStatus returns the Status field value
func (o *CloseAccountResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloseAccountResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloseAccountResponse) SetStatus(v string) {
	o.Status = v
}

func (o CloseAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountCode != nil {
		toSerialize["accountCode"] = o.AccountCode
	}
	if o.InvalidFields != nil {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if o.PspReference != nil {
		toSerialize["pspReference"] = o.PspReference
	}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCloseAccountResponse struct {
	value *CloseAccountResponse
	isSet bool
}

func (v NullableCloseAccountResponse) Get() *CloseAccountResponse {
	return v.value
}

func (v *NullableCloseAccountResponse) Set(val *CloseAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseAccountResponse(val *CloseAccountResponse) *NullableCloseAccountResponse {
	return &NullableCloseAccountResponse{value: val, isSet: true}
}

func (v NullableCloseAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

