package tests

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/adyen/adyen-go-api-library/src/payment"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Payment(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	card := payment.Card{
		Cvc:         "737",
		ExpiryMonth: "03",
		ExpiryYear:  "2030",
		HolderName:  "John Smith",
		Number:      "4111111111111111",
	}

	amount := payment.Amount{
		Currency: "EUR",
		Value:    1500,
	}

	// Used for 3D tests (Skipped for now)
	//browserInfo := &payment.BrowserInfo{
	//    AcceptHeader:      "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	//    ColorDepth:        24,
	//    JavaEnabled:       true,
	//    JavaScriptEnabled: true,
	//    Language:          "en",
	//    ScreenHeight:      1080,
	//    ScreenWidth:       1920,
	//    TimeZoneOffset:    1,
	//    UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
	//}

	client := api.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	// client.GetConfig().Debug = true

	assertForSuccessResponse := func(res interface{}, httpRes *http.Response, err error) {
		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "200 OK", httpRes.Status)
		require.NotNil(t, res)
	}

	authorisePost := func() (payment.PaymentResult, *http.Response, error) {
		return client.Payment.Authorise(&payment.PaymentRequest{
			Card:            &card,
			Amount:          amount,
			Reference:       time.Now().String(),
			MerchantAccount: MerchantAccount,
		})
	}

	t.Run("General", func(t *testing.T) {
		t.Run("Authorise3d", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			res, httpRes, err := client.Payment.Authorise3d(&payment.PaymentRequest3d{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise3ds2", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			res, httpRes, err := client.Payment.Authorise3ds2(&payment.PaymentRequest3ds2{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise", func(t *testing.T) {
			res, httpRes, err := authorisePost()

			assertForSuccessResponse(res, httpRes, err)
			assert.NotNil(t, res.PspReference)
			assert.Equal(t, res.ResultCode, "Authorised")
		})

		t.Run("GetAuthenticationResult", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			res, httpRes, err := client.Payment.GetAuthenticationResult(&payment.AuthenticationResultRequest{})
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Retrieve3ds2Result", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			res, httpRes, err := client.Payment.Retrieve3ds2Result(&payment.ThreeDS2ResultRequest{})
			assertForSuccessResponse(res, httpRes, err)
		})
	})

	t.Run("Modifications", func(t *testing.T) {
		t.Run("AdjustAuthorisation", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.AdjustAuthorisation(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payment.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[adjustAuthorisation-received]", res.Response)
		})

		t.Run("CancelOrRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.CancelOrRefund(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancelOrRefund-received]", res.Response)
		})

		t.Run("Cancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.Cancel(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancel-received]", res.Response)
		})

		t.Run("Capture", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.Capture(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payment.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[capture-received]", res.Response)
		})

		t.Run("Refund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.Refund(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				ModificationAmount: &payment.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       time.Now().String(),
				MerchantAccount: MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[refund-received]", res.Response)
		})

		t.Run("TechnicalCancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.TechnicalCancel(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[technical-cancel-received]", res.Response)
		})

		t.Run("VoidPendingRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			res, httpRes, err := client.Payment.VoidPendingRefund(&payment.ModificationRequest{
				OriginalReference: authRes.PspReference,
				Reference:         time.Now().String(),
				MerchantAccount:   MerchantAccount,
			})
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[voidPendingRefund-received]", res.Response)
		})
	})
}
