/*
OpenProject API V3 (Stable)

Testing WorkScheduleAPIService

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

func Test_openproject_WorkScheduleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkScheduleAPIService CreateNonWorkingDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkScheduleAPI.CreateNonWorkingDay(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService DeleteNonWorkingDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var date string

		httpRes, err := apiClient.WorkScheduleAPI.DeleteNonWorkingDay(context.Background(), date).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ListDays", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkScheduleAPI.ListDays(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ListNonWorkingDays", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkScheduleAPI.ListNonWorkingDays(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ListWeekDays", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkScheduleAPI.ListWeekDays(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService UpdateNonWorkingDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var date string

		resp, httpRes, err := apiClient.WorkScheduleAPI.UpdateNonWorkingDay(context.Background(), date).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService UpdateWeekDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var day int64

		resp, httpRes, err := apiClient.WorkScheduleAPI.UpdateWeekDay(context.Background(), day).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService UpdateWeekDays", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkScheduleAPI.UpdateWeekDays(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ViewDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var date string

		resp, httpRes, err := apiClient.WorkScheduleAPI.ViewDay(context.Background(), date).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ViewNonWorkingDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var date string

		resp, httpRes, err := apiClient.WorkScheduleAPI.ViewNonWorkingDay(context.Background(), date).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkScheduleAPIService ViewWeekDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var day int64

		resp, httpRes, err := apiClient.WorkScheduleAPI.ViewWeekDay(context.Background(), day).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}