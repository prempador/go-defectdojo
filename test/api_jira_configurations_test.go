/*
Defect Dojo API v2

Testing JiraConfigurationsAPIService

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

func Test_openapi_JiraConfigurationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsDeletePreviewList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsDeletePreviewList(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsDestroy(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsPartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsRetrieve(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JiraConfigurationsAPIService JiraConfigurationsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
