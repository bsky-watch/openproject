# WikiPageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identifier of this wiki page | [optional] [readonly] 
**Title** | **string** | The wiki page&#39;s title | 
**Links** | Pointer to [**WikiPageModelLinks**](WikiPageModelLinks.md) |  | [optional] 

## Methods

### NewWikiPageModel

`func NewWikiPageModel(title string, ) *WikiPageModel`

NewWikiPageModel instantiates a new WikiPageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiPageModelWithDefaults

`func NewWikiPageModelWithDefaults() *WikiPageModel`

NewWikiPageModelWithDefaults instantiates a new WikiPageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WikiPageModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiPageModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiPageModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WikiPageModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *WikiPageModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiPageModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiPageModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLinks

`func (o *WikiPageModel) GetLinks() WikiPageModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WikiPageModel) GetLinksOk() (*WikiPageModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WikiPageModel) SetLinks(v WikiPageModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WikiPageModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


