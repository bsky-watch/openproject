# CategoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Category id | [optional] [readonly] 
**Name** | Pointer to **string** | Category name | [optional] 
**Links** | Pointer to [**CategoryModelLinks**](CategoryModelLinks.md) |  | [optional] 

## Methods

### NewCategoryModel

`func NewCategoryModel() *CategoryModel`

NewCategoryModel instantiates a new CategoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryModelWithDefaults

`func NewCategoryModelWithDefaults() *CategoryModel`

NewCategoryModelWithDefaults instantiates a new CategoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CategoryModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CategoryModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CategoryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *CategoryModel) GetLinks() CategoryModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CategoryModel) GetLinksOk() (*CategoryModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CategoryModel) SetLinks(v CategoryModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CategoryModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


