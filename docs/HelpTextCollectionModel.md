# HelpTextCollectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**HelpTextCollectionModelAllOfLinks**](HelpTextCollectionModelAllOfLinks.md) |  | 
**Embedded** | [**HelpTextCollectionModelAllOfEmbedded**](HelpTextCollectionModelAllOfEmbedded.md) |  | 

## Methods

### NewHelpTextCollectionModel

`func NewHelpTextCollectionModel(type_ string, total int64, count int64, links HelpTextCollectionModelAllOfLinks, embedded HelpTextCollectionModelAllOfEmbedded, ) *HelpTextCollectionModel`

NewHelpTextCollectionModel instantiates a new HelpTextCollectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpTextCollectionModelWithDefaults

`func NewHelpTextCollectionModelWithDefaults() *HelpTextCollectionModel`

NewHelpTextCollectionModelWithDefaults instantiates a new HelpTextCollectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HelpTextCollectionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HelpTextCollectionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HelpTextCollectionModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *HelpTextCollectionModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HelpTextCollectionModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HelpTextCollectionModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *HelpTextCollectionModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HelpTextCollectionModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HelpTextCollectionModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *HelpTextCollectionModel) GetLinks() HelpTextCollectionModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HelpTextCollectionModel) GetLinksOk() (*HelpTextCollectionModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HelpTextCollectionModel) SetLinks(v HelpTextCollectionModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetEmbedded

`func (o *HelpTextCollectionModel) GetEmbedded() HelpTextCollectionModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *HelpTextCollectionModel) GetEmbeddedOk() (*HelpTextCollectionModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *HelpTextCollectionModel) SetEmbedded(v HelpTextCollectionModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


