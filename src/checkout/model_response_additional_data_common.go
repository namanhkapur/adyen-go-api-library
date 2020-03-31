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
// ResponseAdditionalDataCommon struct for ResponseAdditionalDataCommon
type ResponseAdditionalDataCommon struct {
	// The name of the Adyen acquirer account.  Example: PayPalSandbox_TestAcquirer > Only relevant for PayPal transactions.
	AcquirerAccountCode string `json:"acquirerAccountCode,omitempty"`
	// The name of the acquirer processing the payment request.  Example: TestPmmAcquirer
	AcquirerCode string `json:"acquirerCode,omitempty"`
	// The reference number that can be used for reconciliation in case a non-Adyen acquirer is used for settlement.  Example: 7C9N3FNBKT9
	AcquirerReference string `json:"acquirerReference,omitempty"`
	// The Adyen alias of the card.  Example: H167852639363479
	Alias string `json:"alias,omitempty"`
	// The type of the card alias.  Example: Default
	AliasType string `json:"aliasType,omitempty"`
	// Authorisation code: * When the payment is authorised successfully, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty.  Example: 58747
	AuthCode string `json:"authCode,omitempty"`
	// The currency of the authorised amount, as a three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	AuthorisedAmountCurrency string `json:"authorisedAmountCurrency,omitempty"`
	// Value of the amount authorised.  This amount is represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes).
	AuthorisedAmountValue string `json:"authorisedAmountValue,omitempty"`
	// The AVS result code of the payment, which provides information about the outcome of the AVS check.  For possible values, see [AVS](https://docs.adyen.com/risk-management/configure-standard-risk-rules/consistency-rules#billing-address-does-not-match-cardholder-address-avs).
	AvsResult string `json:"avsResult,omitempty"`
	// Raw AVS result received from the acquirer, where available.  Example: D
	AvsResultRaw string `json:"avsResultRaw,omitempty"`
	// BIC of a bank account.  Example: TESTNL01 > Only relevant for SEPA Direct Debit transactions.
	Bic string `json:"bic,omitempty"`
	// Supported for 3D Secure 2. The unique transaction identifier assigned by the DS to identify a single transaction.
	DsTransID string `json:"dsTransID,omitempty"`
	// The Electronic Commerce Indicator returned from the schemes for the 3DS payment session.  Example: 02
	Eci string `json:"eci,omitempty"`
	// The expiry date on the card.  Example: 6/2016 > Returned only in case of a card payment.
	ExpiryDate string `json:"expiryDate,omitempty"`
	// The currency of the extra amount charged due to additional amounts set in the skin used in the HPP payment request.  Example: EUR
	ExtraCostsCurrency string `json:"extraCostsCurrency,omitempty"`
	// The value of the extra amount charged due to additional amounts set in the skin used in the HPP payment request. The amount is in minor units.
	ExtraCostsValue string `json:"extraCostsValue,omitempty"`
	// The fraud score due to a particular fraud check. The fraud check name is found in the key of the key-value pair.
	FraudCheckItemNrFraudCheckname string `json:"fraudCheck-[itemNr]-[FraudCheckname],omitempty"`
	// Information regarding the funding type of the card. The possible return values are: * CHARGE * CREDIT * DEBIT * PREPAID * PREPAID_RELOADABLE * PREPAID_NONRELOADABLE * DEFFERED_DEBIT > This functionality requires additional configuration on Adyen's end. To enable it, contact the Support Team.  For receiving this field in the notification, enable **Include Funding Source** in **Notifications** > **Additional settings**.
	FundingSource string `json:"fundingSource,omitempty"`
	// Indicates availability of funds.  Visa: * \"I\" (fast funds are supported) * \"N\" (otherwise)  Mastercard: * \"I\" (product type is Prepaid or Debit, or issuing country is in CEE/HGEM list) * \"N\" (otherwise)  > Returned when you verify a card BIN or estimate costs, and only if payoutEligible is \"Y\" or \"D\".
	FundsAvailability string `json:"fundsAvailability,omitempty"`
	// Provides the more granular indication of why a transaction was refused. When a transaction fails with either \"Refused\", \"Restricted Card\", \"Transaction Not Permitted\", \"Not supported\" or \"DeclinedNon Generic\" refusalReason from the issuer, Adyen cross references its PSP-wide data for extra insight into the refusal reason. If an inferred refusal reason is available, the `inferredRefusalReason`, field is populated and the `refusalReason`, is set to \"Not Supported\".  Possible values: * 3D Secure Mandated * Closed Account * ContAuth Not Supported * CVC Mandated * Ecommerce Not Allowed * Crossborder Not Supported * Card Updated * Low Authrate Bin * Non-reloadable prepaid card
	InferredRefusalReason string `json:"inferredRefusalReason,omitempty"`
	// The issuing country of the card based on the BIN list that Adyen maintains.  Example: JP
	IssuerCountry string `json:"issuerCountry,omitempty"`
	// The `mcBankNetReferenceNumber`, is a minimum of six characters and a maximum of nine characters long. > Contact Support Team to enable this field.
	McBankNetReferenceNumber string `json:"mcBankNetReferenceNumber,omitempty"`
	// Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID.
	NetworkTxReference string `json:"networkTxReference,omitempty"`
	// The owner name of a bank account.  Only relevant for SEPA Direct Debit transactions.
	OwnerName string `json:"ownerName,omitempty"`
	// The Payment Account Reference (PAR) value links a network token with the underlying primary account number (PAN). The PAR value consists of 29 uppercase alphanumeric characters.
	PaymentAccountReference string `json:"paymentAccountReference,omitempty"`
	// The Adyen sub-variant of the payment method used for the payment request.  For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant).  Example: mcpro
	PaymentMethodVariant string `json:"paymentMethodVariant,omitempty"`
	// Indicates whether a payout is eligible or not for this card.  Visa: * \"Y\" * \"N\"  Mastercard: * \"Y\" (domestic and cross-border) * \"D\" (only domestic) * \"N\" (no MoneySend) * \"U\" (unknown)
	PayoutEligible string `json:"payoutEligible,omitempty"`
	// The response code from the Real Time Account Updater service.  Possible return values are: * CardChanged * CardExpiryChanged * CloseAccount * ContactCardAccountHolder
	RealtimeAccountUpdaterStatus string `json:"realtimeAccountUpdaterStatus,omitempty"`
	// Message to be displayed on the terminal.
	ReceiptFreeText string `json:"receiptFreeText,omitempty"`
	// The `pspReference`, of the first recurring payment that created the recurring detail.  This functionality requires additional configuration on Adyen's end. To enable it, contact the Support Team.
	RecurringFirstPspReference string `json:"recurring.firstPspReference,omitempty"`
	// The reference that uniquely identifies the recurring transaction.
	RecurringRecurringDetailReference string `json:"recurring.recurringDetailReference,omitempty"`
	// If the payment is referred, this field is set to true.  This field is unavailable if the payment is referred and is usually not returned with ecommerce transactions.  Example: true
	Referred string `json:"referred,omitempty"`
	// Raw refusal reason received from the acquirer, where available.  Example: AUTHORISED
	RefusalReasonRaw string `json:"refusalReasonRaw,omitempty"`
	// The shopper interaction type of the payment request.  Example: Ecommerce
	ShopperInteraction string `json:"shopperInteraction,omitempty"`
	// The shopperReference passed in the payment request.  Example: AdyenTestShopperXX
	ShopperReference string `json:"shopperReference,omitempty"`
	// The terminal ID used in a point-of-sale payment.  Example: 06022622
	TerminalId string `json:"terminalId,omitempty"`
	// A Boolean value indicating whether 3DS authentication was completed on this payment.  Example: true
	ThreeDAuthenticated string `json:"threeDAuthenticated,omitempty"`
	// The raw 3DS authentication result from the card issuer.  Example: N
	ThreeDAuthenticatedResponse string `json:"threeDAuthenticatedResponse,omitempty"`
	// A Boolean value indicating whether 3DS was offered for this payment.  Example: true
	ThreeDOffered string `json:"threeDOffered,omitempty"`
	// The raw enrollment result from the 3DS directory services of the card schemes.  Example: Y
	ThreeDOfferedResponse string `json:"threeDOfferedResponse,omitempty"`
	// The 3D Secure 2 version.
	ThreeDSVersion string `json:"threeDSVersion,omitempty"`
	// The `visaTransactionId`, has a fixed length of 15 numeric characters. > Contact Support Team to enable this field.
	VisaTransactionId string `json:"visaTransactionId,omitempty"`
	// The 3DS transaction ID of the 3DS session sent in notifications. The value is Base64-encoded and is returned for transactions with directoryResponse 'N' or 'Y'. If you want to submit the xid in your 3D Secure 1 request, use the `mpiData.xid`, field.  Example: ODgxNDc2MDg2MDExODk5MAAAAAA=
	Xid string `json:"xid,omitempty"`
}
