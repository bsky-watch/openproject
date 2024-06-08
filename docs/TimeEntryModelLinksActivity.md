# TimeEntryModelLinksActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | URL to the referenced resource (might be relative) | 
**Title** | Pointer to **string** | Representative label for the resource | [optional] 
**Templated** | Pointer to **bool** | If true the href contains parts that need to be replaced by the client | [optional] [default to false]
**Method** | Pointer to **string** | The HTTP verb to use when requesting the resource | [optional] [default to "GET"]
**Payload** | Pointer to **map[string]interface{}** | The payload to send in the request to achieve the desired result | [optional] 
**Identifier** | Pointer to **string** | An optional unique identifier to the link object | [optional] 
**Type** | Pointer to **string** | The MIME-Type of the returned resource. | [optional] 

## Methods

### NewTimeEntryModelLinksActivity

`func NewTimeEntryModelLinksActivity(href string, ) *TimeEntryModelLinksActivity`

NewTimeEntryModelLinksActivity instantiates a new TimeEntryModelLinksActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryModelLinksActivityWithDefaults

`func NewTimeEntryModelLinksActivityWithDefaults() *TimeEntryModelLinksActivity`

NewTimeEntryModelLinksActivityWithDefaults instantiates a new TimeEntryModelLinksActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *TimeEntryModelLinksActivity) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TimeEntryModelLinksActivity) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TimeEntryModelLinksActivity) SetHref(v string)`

SetHref sets Href field to given value.


### GetTitle

`func (o *TimeEntryModelLinksActivity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimeEntryModelLinksActivity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimeEntryModelLinksActivity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TimeEntryModelLinksActivity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTemplated

`func (o *TimeEntryModelLinksActivity) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *TimeEntryModelLinksActivity) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *TimeEntryModelLinksActivity) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *TimeEntryModelLinksActivity) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetMethod

`func (o *TimeEntryModelLinksActivity) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TimeEntryModelLinksActivity) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TimeEntryModelLinksActivity) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TimeEntryModelLinksActivity) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPayload

`func (o *TimeEntryModelLinksActivity) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TimeEntryModelLinksActivity) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TimeEntryModelLinksActivity) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TimeEntryModelLinksActivity) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetIdentifier

`func (o *TimeEntryModelLinksActivity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TimeEntryModelLinksActivity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TimeEntryModelLinksActivity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TimeEntryModelLinksActivity) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetType

`func (o *TimeEntryModelLinksActivity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeEntryModelLinksActivity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeEntryModelLinksActivity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeEntryModelLinksActivity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


