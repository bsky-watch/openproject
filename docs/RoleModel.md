# RoleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** | Role id | [optional] [readonly] 
**Name** | **string** | Role name | 
**Links** | Pointer to [**RoleModelLinks**](RoleModelLinks.md) |  | [optional] 

## Methods

### NewRoleModel

`func NewRoleModel(name string, ) *RoleModel`

NewRoleModel instantiates a new RoleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleModelWithDefaults

`func NewRoleModelWithDefaults() *RoleModel`

NewRoleModelWithDefaults instantiates a new RoleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoleModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *RoleModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RoleModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleModel) SetName(v string)`

SetName sets Name field to given value.


### GetLinks

`func (o *RoleModel) GetLinks() RoleModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleModel) GetLinksOk() (*RoleModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleModel) SetLinks(v RoleModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


