# FileLinkReadModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**FileLinkReadModelLinksSelf**](FileLinkReadModelLinksSelf.md) |  | 
**Storage** | [**FileLinkReadModelLinksStorage**](FileLinkReadModelLinksStorage.md) |  | 
**Container** | [**FileLinkReadModelLinksContainer**](FileLinkReadModelLinksContainer.md) |  | 
**Creator** | [**FileLinkReadModelLinksCreator**](FileLinkReadModelLinksCreator.md) |  | 
**Delete** | Pointer to [**FileLinkReadModelLinksDelete**](FileLinkReadModelLinksDelete.md) |  | [optional] 
**Status** | Pointer to [**FileLinkReadModelLinksStatus**](FileLinkReadModelLinksStatus.md) |  | [optional] 
**OriginOpen** | [**FileLinkReadModelLinksOriginOpen**](FileLinkReadModelLinksOriginOpen.md) |  | 
**StaticOriginOpen** | [**FileLinkReadModelLinksStaticOriginOpen**](FileLinkReadModelLinksStaticOriginOpen.md) |  | 
**OriginOpenLocation** | [**FileLinkReadModelLinksOriginOpenLocation**](FileLinkReadModelLinksOriginOpenLocation.md) |  | 
**StaticOriginOpenLocation** | [**FileLinkReadModelLinksStaticOriginOpenLocation**](FileLinkReadModelLinksStaticOriginOpenLocation.md) |  | 
**StaticOriginDownload** | [**FileLinkReadModelLinksStaticOriginDownload**](FileLinkReadModelLinksStaticOriginDownload.md) |  | 

## Methods

### NewFileLinkReadModelLinks

`func NewFileLinkReadModelLinks(self FileLinkReadModelLinksSelf, storage FileLinkReadModelLinksStorage, container FileLinkReadModelLinksContainer, creator FileLinkReadModelLinksCreator, originOpen FileLinkReadModelLinksOriginOpen, staticOriginOpen FileLinkReadModelLinksStaticOriginOpen, originOpenLocation FileLinkReadModelLinksOriginOpenLocation, staticOriginOpenLocation FileLinkReadModelLinksStaticOriginOpenLocation, staticOriginDownload FileLinkReadModelLinksStaticOriginDownload, ) *FileLinkReadModelLinks`

NewFileLinkReadModelLinks instantiates a new FileLinkReadModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLinkReadModelLinksWithDefaults

`func NewFileLinkReadModelLinksWithDefaults() *FileLinkReadModelLinks`

NewFileLinkReadModelLinksWithDefaults instantiates a new FileLinkReadModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *FileLinkReadModelLinks) GetSelf() FileLinkReadModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FileLinkReadModelLinks) GetSelfOk() (*FileLinkReadModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FileLinkReadModelLinks) SetSelf(v FileLinkReadModelLinksSelf)`

SetSelf sets Self field to given value.


### GetStorage

`func (o *FileLinkReadModelLinks) GetStorage() FileLinkReadModelLinksStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *FileLinkReadModelLinks) GetStorageOk() (*FileLinkReadModelLinksStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *FileLinkReadModelLinks) SetStorage(v FileLinkReadModelLinksStorage)`

SetStorage sets Storage field to given value.


### GetContainer

`func (o *FileLinkReadModelLinks) GetContainer() FileLinkReadModelLinksContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *FileLinkReadModelLinks) GetContainerOk() (*FileLinkReadModelLinksContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *FileLinkReadModelLinks) SetContainer(v FileLinkReadModelLinksContainer)`

SetContainer sets Container field to given value.


### GetCreator

`func (o *FileLinkReadModelLinks) GetCreator() FileLinkReadModelLinksCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *FileLinkReadModelLinks) GetCreatorOk() (*FileLinkReadModelLinksCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *FileLinkReadModelLinks) SetCreator(v FileLinkReadModelLinksCreator)`

SetCreator sets Creator field to given value.


### GetDelete

`func (o *FileLinkReadModelLinks) GetDelete() FileLinkReadModelLinksDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FileLinkReadModelLinks) GetDeleteOk() (*FileLinkReadModelLinksDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FileLinkReadModelLinks) SetDelete(v FileLinkReadModelLinksDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *FileLinkReadModelLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetStatus

`func (o *FileLinkReadModelLinks) GetStatus() FileLinkReadModelLinksStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileLinkReadModelLinks) GetStatusOk() (*FileLinkReadModelLinksStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileLinkReadModelLinks) SetStatus(v FileLinkReadModelLinksStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileLinkReadModelLinks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOriginOpen

`func (o *FileLinkReadModelLinks) GetOriginOpen() FileLinkReadModelLinksOriginOpen`

GetOriginOpen returns the OriginOpen field if non-nil, zero value otherwise.

### GetOriginOpenOk

`func (o *FileLinkReadModelLinks) GetOriginOpenOk() (*FileLinkReadModelLinksOriginOpen, bool)`

GetOriginOpenOk returns a tuple with the OriginOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOpen

`func (o *FileLinkReadModelLinks) SetOriginOpen(v FileLinkReadModelLinksOriginOpen)`

SetOriginOpen sets OriginOpen field to given value.


### GetStaticOriginOpen

`func (o *FileLinkReadModelLinks) GetStaticOriginOpen() FileLinkReadModelLinksStaticOriginOpen`

GetStaticOriginOpen returns the StaticOriginOpen field if non-nil, zero value otherwise.

### GetStaticOriginOpenOk

`func (o *FileLinkReadModelLinks) GetStaticOriginOpenOk() (*FileLinkReadModelLinksStaticOriginOpen, bool)`

GetStaticOriginOpenOk returns a tuple with the StaticOriginOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOriginOpen

`func (o *FileLinkReadModelLinks) SetStaticOriginOpen(v FileLinkReadModelLinksStaticOriginOpen)`

SetStaticOriginOpen sets StaticOriginOpen field to given value.


### GetOriginOpenLocation

`func (o *FileLinkReadModelLinks) GetOriginOpenLocation() FileLinkReadModelLinksOriginOpenLocation`

GetOriginOpenLocation returns the OriginOpenLocation field if non-nil, zero value otherwise.

### GetOriginOpenLocationOk

`func (o *FileLinkReadModelLinks) GetOriginOpenLocationOk() (*FileLinkReadModelLinksOriginOpenLocation, bool)`

GetOriginOpenLocationOk returns a tuple with the OriginOpenLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOpenLocation

`func (o *FileLinkReadModelLinks) SetOriginOpenLocation(v FileLinkReadModelLinksOriginOpenLocation)`

SetOriginOpenLocation sets OriginOpenLocation field to given value.


### GetStaticOriginOpenLocation

`func (o *FileLinkReadModelLinks) GetStaticOriginOpenLocation() FileLinkReadModelLinksStaticOriginOpenLocation`

GetStaticOriginOpenLocation returns the StaticOriginOpenLocation field if non-nil, zero value otherwise.

### GetStaticOriginOpenLocationOk

`func (o *FileLinkReadModelLinks) GetStaticOriginOpenLocationOk() (*FileLinkReadModelLinksStaticOriginOpenLocation, bool)`

GetStaticOriginOpenLocationOk returns a tuple with the StaticOriginOpenLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOriginOpenLocation

`func (o *FileLinkReadModelLinks) SetStaticOriginOpenLocation(v FileLinkReadModelLinksStaticOriginOpenLocation)`

SetStaticOriginOpenLocation sets StaticOriginOpenLocation field to given value.


### GetStaticOriginDownload

`func (o *FileLinkReadModelLinks) GetStaticOriginDownload() FileLinkReadModelLinksStaticOriginDownload`

GetStaticOriginDownload returns the StaticOriginDownload field if non-nil, zero value otherwise.

### GetStaticOriginDownloadOk

`func (o *FileLinkReadModelLinks) GetStaticOriginDownloadOk() (*FileLinkReadModelLinksStaticOriginDownload, bool)`

GetStaticOriginDownloadOk returns a tuple with the StaticOriginDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOriginDownload

`func (o *FileLinkReadModelLinks) SetStaticOriginDownload(v FileLinkReadModelLinksStaticOriginDownload)`

SetStaticOriginDownload sets StaticOriginDownload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


