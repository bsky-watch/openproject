# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ErrorResponseEmbedded**](ErrorResponseEmbedded.md) |  | [optional] 
**Type** | **string** |  | 
**ErrorIdentifier** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewErrorResponse

`func NewErrorResponse(type_ string, errorIdentifier string, message string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ErrorResponse) GetEmbedded() ErrorResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ErrorResponse) GetEmbeddedOk() (*ErrorResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ErrorResponse) SetEmbedded(v ErrorResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ErrorResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetType

`func (o *ErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetErrorIdentifier

`func (o *ErrorResponse) GetErrorIdentifier() string`

GetErrorIdentifier returns the ErrorIdentifier field if non-nil, zero value otherwise.

### GetErrorIdentifierOk

`func (o *ErrorResponse) GetErrorIdentifierOk() (*string, bool)`

GetErrorIdentifierOk returns a tuple with the ErrorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIdentifier

`func (o *ErrorResponse) SetErrorIdentifier(v string)`

SetErrorIdentifier sets ErrorIdentifier field to given value.


### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


