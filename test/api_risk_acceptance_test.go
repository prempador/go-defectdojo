/*
Defect Dojo API v2

Testing RiskAcceptanceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/prempador/go-defectdojo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_RiskAcceptanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceDeletePreviewList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDeletePreviewList(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDestroy(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceDownloadProofRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDownloadProofRetrieve(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptancePartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptancePartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceRetrieve(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAcceptanceAPIService RiskAcceptanceUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
