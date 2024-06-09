/*
OpenProject API V3 (Stable)

Testing FileLinksAPIService

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

func Test_openproject_FileLinksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FileLinksAPIService CreateStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FileLinksAPI.CreateStorage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService CreateStorageOauthCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.CreateStorageOauthCredentials(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService DeleteFileLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.DeleteFileLink(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService DeleteStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.DeleteStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService DownloadFileLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.DownloadFileLink(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService GetProjectStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.GetProjectStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService GetStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.GetStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService GetStorageFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.GetStorageFiles(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService ListProjectStorages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FileLinksAPI.ListProjectStorages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService ListStorages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FileLinksAPI.ListStorages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService OpenFileLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.OpenFileLink(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService OpenProjectStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.OpenProjectStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService OpenStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.FileLinksAPI.OpenStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService PrepareStorageFileUpload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.PrepareStorageFileUpload(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService UpdateStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.UpdateStorage(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FileLinksAPIService ViewFileLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.FileLinksAPI.ViewFileLink(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
