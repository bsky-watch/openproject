# GridCollectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Total** | **int64** | The total amount of elements available in the collection. | 
**Count** | **int64** | Actual amount of elements in this response. | 
**Links** | [**PaginatedCollectionModelAllOfLinks**](PaginatedCollectionModelAllOfLinks.md) |  | 
**PageSize** | **int64** | Amount of elements that a response will hold. | 
**Offset** | **int64** | The page number that is requested from paginated collection. | 
**Embedded** | [**GridCollectionModelAllOfEmbedded**](GridCollectionModelAllOfEmbedded.md) |  | 

## Methods

### NewGridCollectionModel

`func NewGridCollectionModel(type_ string, total int64, count int64, links PaginatedCollectionModelAllOfLinks, pageSize int64, offset int64, embedded GridCollectionModelAllOfEmbedded, ) *GridCollectionModel`

NewGridCollectionModel instantiates a new GridCollectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCollectionModelWithDefaults

`func NewGridCollectionModelWithDefaults() *GridCollectionModel`

NewGridCollectionModelWithDefaults instantiates a new GridCollectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GridCollectionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridCollectionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridCollectionModel) SetType(v string)`

SetType sets Type field to given value.


### GetTotal

`func (o *GridCollectionModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GridCollectionModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GridCollectionModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCount

`func (o *GridCollectionModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GridCollectionModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GridCollectionModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLinks

`func (o *GridCollectionModel) GetLinks() PaginatedCollectionModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GridCollectionModel) GetLinksOk() (*PaginatedCollectionModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GridCollectionModel) SetLinks(v PaginatedCollectionModelAllOfLinks)`

SetLinks sets Links field to given value.


### GetPageSize

`func (o *GridCollectionModel) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GridCollectionModel) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GridCollectionModel) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetOffset

`func (o *GridCollectionModel) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GridCollectionModel) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GridCollectionModel) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetEmbedded

`func (o *GridCollectionModel) GetEmbedded() GridCollectionModelAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *GridCollectionModel) GetEmbeddedOk() (*GridCollectionModelAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *GridCollectionModel) SetEmbedded(v GridCollectionModelAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


