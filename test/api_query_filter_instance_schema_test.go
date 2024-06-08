/*
OpenProject API V3 (Stable)

Testing QueryFilterInstanceSchemaAPIService

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

func Test_openproject_QueryFilterInstanceSchemaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QueryFilterInstanceSchemaAPIService ListQueryFilterInstanceSchemas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryFilterInstanceSchemaAPIService ListQueryFilterInstanceSchemasForProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemasForProject(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QueryFilterInstanceSchemaAPIService ViewQueryFilterInstanceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.QueryFilterInstanceSchemaAPI.ViewQueryFilterInstanceSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}