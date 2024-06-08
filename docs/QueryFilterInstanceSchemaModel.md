# QueryFilterInstanceSchemaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Describes the name attribute | [readonly] 
**Filter** | **string** | QuerySortBy name | 
**Links** | Pointer to [**QueryFilterInstanceSchemaModelLinks**](QueryFilterInstanceSchemaModelLinks.md) |  | [optional] 

## Methods

### NewQueryFilterInstanceSchemaModel

`func NewQueryFilterInstanceSchemaModel(name string, filter string, ) *QueryFilterInstanceSchemaModel`

NewQueryFilterInstanceSchemaModel instantiates a new QueryFilterInstanceSchemaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFilterInstanceSchemaModelWithDefaults

`func NewQueryFilterInstanceSchemaModelWithDefaults() *QueryFilterInstanceSchemaModel`

NewQueryFilterInstanceSchemaModelWithDefaults instantiates a new QueryFilterInstanceSchemaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryFilterInstanceSchemaModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryFilterInstanceSchemaModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryFilterInstanceSchemaModel) SetName(v string)`

SetName sets Name field to given value.


### GetFilter

`func (o *QueryFilterInstanceSchemaModel) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryFilterInstanceSchemaModel) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryFilterInstanceSchemaModel) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetLinks

`func (o *QueryFilterInstanceSchemaModel) GetLinks() QueryFilterInstanceSchemaModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *QueryFilterInstanceSchemaModel) GetLinksOk() (*QueryFilterInstanceSchemaModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *QueryFilterInstanceSchemaModel) SetLinks(v QueryFilterInstanceSchemaModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *QueryFilterInstanceSchemaModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


