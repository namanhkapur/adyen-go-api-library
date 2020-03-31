/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including One-Click and 3D Secure), mobile wallets, and local payment methods (e.g. iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v51/payments ```
 *
 * API version: 51
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// AdditionalDataCarRental struct for AdditionalDataCarRental
type AdditionalDataCarRental struct {
	// The rental agreement number associated with this car rental. * Format: Alphanumeric * maxLength: 14
	CarRentalRentalAgreementNumber string `json:"carRental.rentalAgreementNumber,omitempty"`
	// The name of the person renting the car. * Format: Alphanumeric * maxLength: 26
	CarRentalRenterName string `json:"carRental.renterName,omitempty"`
	// The city where the car must be returned. * Format: Alphanumeric * maxLength: 18
	CarRentalReturnCity string `json:"carRental.returnCity,omitempty"`
	// The state or province where the car must be returned. * Format: Alphanumeric * maxLength: 3
	CarRentalReturnStateProvince string `json:"carRental.returnStateProvince,omitempty"`
	// The country where the car must be returned. * Format: Alphanumeric * maxLength: 2
	CarRentalReturnCountry string `json:"carRental.returnCountry,omitempty"`
	// Agency code, phone number, or address abbreviation * Format: Alphanumeric * maxLength: 10
	CarRentalReturnLocationId string `json:"carRental.returnLocationId,omitempty"`
	// The last date to return the car by. * Date format: `yyyyMMdd`
	CarRentalReturnDate string `json:"carRental.returnDate,omitempty"`
	// Pick-up date. * Date format: `yyyyMMdd`
	CarRentalCheckOutDate string `json:"carRental.checkOutDate,omitempty"`
	// The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17
	CarRentalCustomerServiceTollFreeNumber string `json:"carRental.customerServiceTollFreeNumber,omitempty"`
	// Daily rental rate. * Format: Alphanumeric * maxLength: 12
	CarRentalRate string `json:"carRental.rate,omitempty"`
	// Specifies whether the given rate is applied daily or weekly. * D - Daily rate. * W - Weekly rate.
	CarRentalRateIndicator string `json:"carRental.rateIndicator,omitempty"`
	// The location from which the car is rented. * Format: Alphanumeric * maxLength: 18
	CarRentalLocationCity string `json:"carRental.locationCity,omitempty"`
	// Pick-up date. * Date format: `yyyyMMdd`
	CarRentalLocationStateProvince string `json:"carRental.locationStateProvince,omitempty"`
	// The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17
	CarRentalLocationCountry string `json:"carRental.locationCountry,omitempty"`
	// Daily rental rate. * Format: Alphanumeric * maxLength: 12
	CarRentalRentalClassId string `json:"carRental.rentalClassId,omitempty"`
	// Specifies whether the given rate is applied daily or weekly. * D - Daily rate. * W - Weekly rate.
	CarRentalDaysRented string `json:"carRental.daysRented,omitempty"`
	// Indicates whether the goods or services were tax-exempt, or tax was not collected.  Values: * 0 - Tax was not collected * 1 - Goods or services were tax exempt
	CarRentalTaxExemptIndicator string `json:"carRental.taxExemptIndicator,omitempty"`
	// Indicates what market-specific dataset will be submitted or is being submitted. Value should be \"A\" for Car rental. This should be included in the auth message. * Format: Alphanumeric * maxLength: 1
	TravelEntertainmentAuthDataMarket string `json:"travelEntertainmentAuthData.market,omitempty"`
	// Number of nights.  This should be included in the auth message. * Format: Numeric * maxLength: 2
	TravelEntertainmentAuthDataDuration string `json:"travelEntertainmentAuthData.duration,omitempty"`
	// Any fuel charges associated with the rental. * Format: Numeric * maxLength: 12
	CarRentalFuelCharges string `json:"carRental.fuelCharges,omitempty"`
	// Any insurance charges associated with the rental. * Format: Numeric * maxLength: 12
	CarRentalInsuranceCharges string `json:"carRental.insuranceCharges,omitempty"`
	// Indicates if the customer was a \"no-show\" (neither keeps nor cancels their booking). * 0 - Not applicable. * 1 - Customer was a no show.
	CarRentalNoShowIndicator string `json:"carRental.noShowIndicator,omitempty"`
	// Charge associated with not returning a vehicle to the original rental location.
	CarRentalOneWayDropOffCharges string `json:"carRental.oneWayDropOffCharges,omitempty"`
}
