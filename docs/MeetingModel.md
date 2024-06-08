# MeetingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identifier of this meeting | [optional] [readonly] 
**Title** | **string** | The meeting&#39;s title | 
**Links** | Pointer to [**MeetingModelLinks**](MeetingModelLinks.md) |  | [optional] 

## Methods

### NewMeetingModel

`func NewMeetingModel(title string, ) *MeetingModel`

NewMeetingModel instantiates a new MeetingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingModelWithDefaults

`func NewMeetingModelWithDefaults() *MeetingModel`

NewMeetingModelWithDefaults instantiates a new MeetingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeetingModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeetingModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeetingModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MeetingModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *MeetingModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MeetingModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MeetingModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLinks

`func (o *MeetingModel) GetLinks() MeetingModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeetingModel) GetLinksOk() (*MeetingModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeetingModel) SetLinks(v MeetingModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MeetingModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


