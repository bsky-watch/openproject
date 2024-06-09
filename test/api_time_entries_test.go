/*
OpenProject API V3 (Stable)

Testing TimeEntriesAPIService

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

func Test_openproject_TimeEntriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TimeEntriesAPIService AvailableProjectsForTimeEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TimeEntriesAPI.AvailableProjectsForTimeEntries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService CreateTimeEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TimeEntriesAPI.CreateTimeEntry(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService DeleteTimeEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.TimeEntriesAPI.DeleteTimeEntry(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService GetTimeEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.TimeEntriesAPI.GetTimeEntry(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService ListTimeEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TimeEntriesAPI.ListTimeEntries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService TimeEntryCreateForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TimeEntriesAPI.TimeEntryCreateForm(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService TimeEntryUpdateForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		httpRes, err := apiClient.TimeEntriesAPI.TimeEntryUpdateForm(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService UpdateTimeEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int

		resp, httpRes, err := apiClient.TimeEntriesAPI.UpdateTimeEntry(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeEntriesAPIService ViewTimeEntrySchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TimeEntriesAPI.ViewTimeEntrySchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
