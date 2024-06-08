# CustomActionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the custom action | [optional] 
**Description** | Pointer to **string** | The description for the custom action | [optional] 
**Links** | Pointer to [**CustomActionModelLinks**](CustomActionModelLinks.md) |  | [optional] 

## Methods

### NewCustomActionModel

`func NewCustomActionModel() *CustomActionModel`

NewCustomActionModel instantiates a new CustomActionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomActionModelWithDefaults

`func NewCustomActionModelWithDefaults() *CustomActionModel`

NewCustomActionModelWithDefaults instantiates a new CustomActionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomActionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomActionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomActionModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomActionModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CustomActionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomActionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomActionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomActionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CustomActionModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomActionModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomActionModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomActionModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *CustomActionModel) GetLinks() CustomActionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomActionModel) GetLinksOk() (*CustomActionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomActionModel) SetLinks(v CustomActionModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomActionModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


