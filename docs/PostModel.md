# PostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identifier of this post | [optional] [readonly] 
**Subject** | **string** | The post&#39;s subject | 
**Links** | Pointer to [**PostModelLinks**](PostModelLinks.md) |  | [optional] 

## Methods

### NewPostModel

`func NewPostModel(subject string, ) *PostModel`

NewPostModel instantiates a new PostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostModelWithDefaults

`func NewPostModelWithDefaults() *PostModel`

NewPostModelWithDefaults instantiates a new PostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *PostModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PostModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PostModel) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetLinks

`func (o *PostModel) GetLinks() PostModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostModel) GetLinksOk() (*PostModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostModel) SetLinks(v PostModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PostModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


