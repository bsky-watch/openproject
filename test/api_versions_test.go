/*
OpenProject API V3 (Stable)

Testing VersionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openproject

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "bsky.watch/openproject"
)

func Test_openproject_VersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VersionsAPIService AvailableProjectsForVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VersionsAPI.AvailableProjectsForVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService CreateVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VersionsAPI.CreateVersion(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService DeleteVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.VersionsAPI.DeleteVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService ListVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VersionsAPI.ListVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService ListVersionsAvailableInAProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.VersionsAPI.ListVersionsAvailableInAProject(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService UpdateVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.VersionsAPI.UpdateVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService VersionCreateForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VersionsAPI.VersionCreateForm(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService VersionUpdateForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.VersionsAPI.VersionUpdateForm(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService ViewVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.VersionsAPI.ViewVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VersionsAPIService ViewVersionSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VersionsAPI.ViewVersionSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
