/*
Defect Dojo API v2

Testing ImportScanAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package defectdojo

import (
	"context"
	"testing"

	openapiclient "github.com/prempador/go-defectdojo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ImportScanAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportScanAPIService ImportScanCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImportScanAPI.ImportScanCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
