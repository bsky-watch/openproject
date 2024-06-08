# UpdateActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**UpdateActivityRequestLinks**](UpdateActivityRequestLinks.md) |  | [optional] 

## Methods

### NewUpdateActivityRequest

`func NewUpdateActivityRequest() *UpdateActivityRequest`

NewUpdateActivityRequest instantiates a new UpdateActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActivityRequestWithDefaults

`func NewUpdateActivityRequestWithDefaults() *UpdateActivityRequest`

NewUpdateActivityRequestWithDefaults instantiates a new UpdateActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateActivityRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateActivityRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateActivityRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateActivityRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLinks

`func (o *UpdateActivityRequest) GetLinks() UpdateActivityRequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateActivityRequest) GetLinksOk() (*UpdateActivityRequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateActivityRequest) SetLinks(v UpdateActivityRequestLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateActivityRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


