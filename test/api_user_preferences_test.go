/*
OpenProject API V3 (Stable)

Testing UserPreferencesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openproject

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func Test_openproject_UserPreferencesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserPreferencesAPIService ShowMyPreferences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserPreferencesAPI.ShowMyPreferences(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserPreferencesAPIService UpdateUserPreferences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserPreferencesAPI.UpdateUserPreferences(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}