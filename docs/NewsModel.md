# NewsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | News&#39; id | [optional] [readonly] 
**Title** | Pointer to **string** | The headline of the news | [optional] [readonly] 
**Summary** | Pointer to **string** | A short summary | [optional] [readonly] 
**Description** | Pointer to **string** | The main body of the news with all the details | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The time the news was created at | [optional] [readonly] 
**Links** | Pointer to [**NewsModelLinks**](NewsModelLinks.md) |  | [optional] 

## Methods

### NewNewsModel

`func NewNewsModel() *NewsModel`

NewNewsModel instantiates a new NewsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsModelWithDefaults

`func NewNewsModelWithDefaults() *NewsModel`

NewNewsModelWithDefaults instantiates a new NewsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewsModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewsModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewsModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NewsModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *NewsModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewsModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewsModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewsModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSummary

`func (o *NewsModel) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NewsModel) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NewsModel) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NewsModel) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *NewsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewsModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewsModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NewsModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NewsModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NewsModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NewsModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *NewsModel) GetLinks() NewsModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NewsModel) GetLinksOk() (*NewsModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NewsModel) SetLinks(v NewsModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NewsModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


