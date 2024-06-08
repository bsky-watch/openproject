# SchemaPropertyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource type for this property. | 
**Name** | **string** | The name of the property. | 
**Required** | **bool** | Indicates, if the property is required for submitting a request of this schema. | 
**HasDefault** | **bool** | Indicates, if the property has a default. | 
**Writable** | **bool** | Indicates, if the property is writable when sending a request of this schema. | 
**Object** | Pointer to **map[string]interface{}** | Additional options for the property. | [optional] 
**Location** | Pointer to **string** | Defines the json path where the property is located in the payload. | [optional] [default to ""]
**Links** | Pointer to **map[string]interface{}** | Useful links for this property (e.g. an endpoint to fetch allowed values) | [optional] 

## Methods

### NewSchemaPropertyModel

`func NewSchemaPropertyModel(type_ string, name string, required bool, hasDefault bool, writable bool, ) *SchemaPropertyModel`

NewSchemaPropertyModel instantiates a new SchemaPropertyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaPropertyModelWithDefaults

`func NewSchemaPropertyModelWithDefaults() *SchemaPropertyModel`

NewSchemaPropertyModelWithDefaults instantiates a new SchemaPropertyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaPropertyModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaPropertyModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaPropertyModel) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SchemaPropertyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaPropertyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaPropertyModel) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *SchemaPropertyModel) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SchemaPropertyModel) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SchemaPropertyModel) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetHasDefault

`func (o *SchemaPropertyModel) GetHasDefault() bool`

GetHasDefault returns the HasDefault field if non-nil, zero value otherwise.

### GetHasDefaultOk

`func (o *SchemaPropertyModel) GetHasDefaultOk() (*bool, bool)`

GetHasDefaultOk returns a tuple with the HasDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefault

`func (o *SchemaPropertyModel) SetHasDefault(v bool)`

SetHasDefault sets HasDefault field to given value.


### GetWritable

`func (o *SchemaPropertyModel) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *SchemaPropertyModel) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *SchemaPropertyModel) SetWritable(v bool)`

SetWritable sets Writable field to given value.


### GetObject

`func (o *SchemaPropertyModel) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SchemaPropertyModel) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SchemaPropertyModel) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *SchemaPropertyModel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetLocation

`func (o *SchemaPropertyModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SchemaPropertyModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SchemaPropertyModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SchemaPropertyModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLinks

`func (o *SchemaPropertyModel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchemaPropertyModel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchemaPropertyModel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SchemaPropertyModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


