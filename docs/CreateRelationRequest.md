# CreateRelationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Links** | [**CreateRelationRequestLinks**](CreateRelationRequestLinks.md) |  | 

## Methods

### NewCreateRelationRequest

`func NewCreateRelationRequest(type_ string, links CreateRelationRequestLinks, ) *CreateRelationRequest`

NewCreateRelationRequest instantiates a new CreateRelationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRelationRequestWithDefaults

`func NewCreateRelationRequestWithDefaults() *CreateRelationRequest`

NewCreateRelationRequestWithDefaults instantiates a new CreateRelationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateRelationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRelationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRelationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *CreateRelationRequest) GetLinks() CreateRelationRequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateRelationRequest) GetLinksOk() (*CreateRelationRequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateRelationRequest) SetLinks(v CreateRelationRequestLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


