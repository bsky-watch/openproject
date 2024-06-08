# RootModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InstanceName** | **string** | The name of the OpenProject instance | 
**CoreVersion** | Pointer to **string** | The OpenProject core version number for the instance  # Conditions  **Permission** requires admin privileges | [optional] 
**Links** | [**RootModelLinks**](RootModelLinks.md) |  | 

## Methods

### NewRootModel

`func NewRootModel(type_ string, instanceName string, links RootModelLinks, ) *RootModel`

NewRootModel instantiates a new RootModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootModelWithDefaults

`func NewRootModelWithDefaults() *RootModel`

NewRootModelWithDefaults instantiates a new RootModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RootModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RootModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RootModel) SetType(v string)`

SetType sets Type field to given value.


### GetInstanceName

`func (o *RootModel) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *RootModel) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *RootModel) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetCoreVersion

`func (o *RootModel) GetCoreVersion() string`

GetCoreVersion returns the CoreVersion field if non-nil, zero value otherwise.

### GetCoreVersionOk

`func (o *RootModel) GetCoreVersionOk() (*string, bool)`

GetCoreVersionOk returns a tuple with the CoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreVersion

`func (o *RootModel) SetCoreVersion(v string)`

SetCoreVersion sets CoreVersion field to given value.

### HasCoreVersion

`func (o *RootModel) HasCoreVersion() bool`

HasCoreVersion returns a boolean if a field has been set.

### GetLinks

`func (o *RootModel) GetLinks() RootModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RootModel) GetLinksOk() (*RootModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RootModel) SetLinks(v RootModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


