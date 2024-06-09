/*
OpenProject API V3 (Stable)

Testing UsersAPIService

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

func Test_openproject_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService CreateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService LockUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.UsersAPI.LockUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UnlockUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.UsersAPI.UnlockUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.UsersAPI.UpdateUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UserUpdateForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.UsersAPI.UserUpdateForm(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ViewUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.UsersAPI.ViewUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ViewUserSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ViewUserSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
