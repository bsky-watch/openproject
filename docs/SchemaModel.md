# SchemaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Dependencies** | Pointer to **[]string** | A list of dependencies between one property&#39;s value and another property | [optional] 
**Links** | [**SchemaModelLinks**](SchemaModelLinks.md) |  | 

## Methods

### NewSchemaModel

`func NewSchemaModel(type_ string, links SchemaModelLinks, ) *SchemaModel`

NewSchemaModel instantiates a new SchemaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaModelWithDefaults

`func NewSchemaModelWithDefaults() *SchemaModel`

NewSchemaModelWithDefaults instantiates a new SchemaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaModel) SetType(v string)`

SetType sets Type field to given value.


### GetDependencies

`func (o *SchemaModel) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *SchemaModel) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *SchemaModel) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *SchemaModel) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetLinks

`func (o *SchemaModel) GetLinks() SchemaModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchemaModel) GetLinksOk() (*SchemaModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchemaModel) SetLinks(v SchemaModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


