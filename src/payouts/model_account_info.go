/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

import (
	"time"
)

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	// Indicator for the length of time since this shopper account was created in the merchant's environment. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountAgeIndicator string `json:"accountAgeIndicator,omitempty"`
	// Date when the shopper's account was last changed.
	AccountChangeDate *time.Time `json:"accountChangeDate,omitempty"`
	// Indicator for the length of time since the shopper's account was last updated. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountChangeIndicator string `json:"accountChangeIndicator,omitempty"`
	// Date when the shopper's account was created.
	AccountCreationDate *time.Time `json:"accountCreationDate,omitempty"`
	// Indicates the type of account. For example, for a multi-account card product. Allowed values: * notApplicable * credit * debit
	AccountType string `json:"accountType,omitempty"`
	// Number of attempts the shopper tried to add a card to their account in the last day.
	AddCardAttemptsDay int32 `json:"addCardAttemptsDay,omitempty"`
	// Date the selected delivery address was first used.
	DeliveryAddressUsageDate *time.Time `json:"deliveryAddressUsageDate,omitempty"`
	// Indicator for the length of time since this delivery address was first used. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	DeliveryAddressUsageIndicator string `json:"deliveryAddressUsageIndicator,omitempty"`
	// Shopper's home phone number (including the country code).
	HomePhone string `json:"homePhone,omitempty"`
	// Shopper's mobile phone number (including the country code).
	MobilePhone string `json:"mobilePhone,omitempty"`
	// Date when the shopper last changed their password.
	PasswordChangeDate *time.Time `json:"passwordChangeDate,omitempty"`
	// Indicator when the shopper has changed their password. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PasswordChangeIndicator string `json:"passwordChangeIndicator,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past 24 hours.
	PastTransactionsDay int32 `json:"pastTransactionsDay,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past year.
	PastTransactionsYear int32 `json:"pastTransactionsYear,omitempty"`
	// Date this payment method was added to the shopper's account.
	PaymentAccountAge *time.Time `json:"paymentAccountAge,omitempty"`
	// Indicator for the length of time since this payment method was added to this shopper's account. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PaymentAccountIndicator string `json:"paymentAccountIndicator,omitempty"`
	// Number of successful purchases in the last six months.
	PurchasesLast6Months int32 `json:"purchasesLast6Months,omitempty"`
	// Whether suspicious activity was recorded on this account.
	SuspiciousActivity bool `json:"suspiciousActivity,omitempty"`
	// Shopper's work phone number (including the country code).
	WorkPhone string `json:"workPhone,omitempty"`
}
