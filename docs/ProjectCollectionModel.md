# ProjectCollectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**ProjectCollectionModelAllOfLinks**](ProjectCollectionModelAllOfLinks.md) |  | 
**Embedded** | [**ProjectCollectionModelAllOfEmbedded**](ProjectCollectionModelAllOfEmbedded.md) |  | 

## Methods

### NewProjectCollectionModel

`func NewProjectCollectionModel(type_ string, total int64, count int64, links ProjectCollectionModelAllOfLinks, embedded ProjectCollectionModelAllOfEmbedded, ) *ProjectCollectionModel`

NewProjectCollectionModel instantiates a new ProjectCollectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCollectionModelWithDefaults

`func NewProjectCollectionModelWithDefaults() *ProjectCollectionModel`

NewProjectCollectionModelWithDefaults instantiates a new ProjectCollectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProjectCollectionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectCollectionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectCollectionModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *ProjectCollectionModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProjectCollectionModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProjectCollectionModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *ProjectCollectionModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProjectCollectionModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProjectCollectionModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *ProjectCollectionModel) GetLinks() ProjectCollectionModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectCollectionModel) GetLinksOk() (*ProjectCollectionModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectCollectionModel) SetLinks(v ProjectCollectionModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetEmbedded

`func (o *ProjectCollectionModel) GetEmbedded() ProjectCollectionModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ProjectCollectionModel) GetEmbeddedOk() (*ProjectCollectionModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ProjectCollectionModel) SetEmbedded(v ProjectCollectionModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


