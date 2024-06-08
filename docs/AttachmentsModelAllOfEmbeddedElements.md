# AttachmentsModelAllOfEmbeddedElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Attachment&#39;s id | [optional] 
**Title** | Pointer to **string** | The name of the file | [optional] 
**FileName** | **string** | The name of the uploaded file | 
**FileSize** | Pointer to **int64** | The size of the uploaded file in Bytes | [optional] 
**Description** | [**AttachmentModelDescription**](AttachmentModelDescription.md) |  | 
**ContentType** | **string** | The files MIME-Type as determined by the server | 
**Digest** | [**AttachmentModelDigest**](AttachmentModelDigest.md) |  | 
**CreatedAt** | **time.Time** | Time of creation | 
**Links** | Pointer to [**AttachmentModelLinks**](AttachmentModelLinks.md) |  | [optional] 

## Methods

### NewAttachmentsModelAllOfEmbeddedElements

`func NewAttachmentsModelAllOfEmbeddedElements(fileName string, description AttachmentModelDescription, contentType string, digest AttachmentModelDigest, createdAt time.Time, ) *AttachmentsModelAllOfEmbeddedElements`

NewAttachmentsModelAllOfEmbeddedElements instantiates a new AttachmentsModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsModelAllOfEmbeddedElementsWithDefaults

`func NewAttachmentsModelAllOfEmbeddedElementsWithDefaults() *AttachmentsModelAllOfEmbeddedElements`

NewAttachmentsModelAllOfEmbeddedElementsWithDefaults instantiates a new AttachmentsModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachmentsModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentsModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentsModelAllOfEmbeddedElements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *AttachmentsModelAllOfEmbeddedElements) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentsModelAllOfEmbeddedElements) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentsModelAllOfEmbeddedElements) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFileName

`func (o *AttachmentsModelAllOfEmbeddedElements) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AttachmentsModelAllOfEmbeddedElements) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *AttachmentsModelAllOfEmbeddedElements) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AttachmentsModelAllOfEmbeddedElements) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AttachmentsModelAllOfEmbeddedElements) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetDescription

`func (o *AttachmentsModelAllOfEmbeddedElements) GetDescription() AttachmentModelDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetDescriptionOk() (*AttachmentModelDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentsModelAllOfEmbeddedElements) SetDescription(v AttachmentModelDescription)`

SetDescription sets Description field to given value.


### GetContentType

`func (o *AttachmentsModelAllOfEmbeddedElements) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AttachmentsModelAllOfEmbeddedElements) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDigest

`func (o *AttachmentsModelAllOfEmbeddedElements) GetDigest() AttachmentModelDigest`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetDigestOk() (*AttachmentModelDigest, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AttachmentsModelAllOfEmbeddedElements) SetDigest(v AttachmentModelDigest)`

SetDigest sets Digest field to given value.


### GetCreatedAt

`func (o *AttachmentsModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentsModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLinks

`func (o *AttachmentsModelAllOfEmbeddedElements) GetLinks() AttachmentModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachmentsModelAllOfEmbeddedElements) GetLinksOk() (*AttachmentModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachmentsModelAllOfEmbeddedElements) SetLinks(v AttachmentModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachmentsModelAllOfEmbeddedElements) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


