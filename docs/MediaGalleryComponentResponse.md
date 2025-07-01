# MediaGalleryComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**Items** | [**[]MediaGalleryItemResponse**](MediaGalleryItemResponse.md) |  | 

## Methods

### NewMediaGalleryComponentResponse

`func NewMediaGalleryComponentResponse(type_ int32, id int32, items []MediaGalleryItemResponse, ) *MediaGalleryComponentResponse`

NewMediaGalleryComponentResponse instantiates a new MediaGalleryComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGalleryComponentResponseWithDefaults

`func NewMediaGalleryComponentResponseWithDefaults() *MediaGalleryComponentResponse`

NewMediaGalleryComponentResponseWithDefaults instantiates a new MediaGalleryComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MediaGalleryComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaGalleryComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaGalleryComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *MediaGalleryComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaGalleryComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaGalleryComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetItems

`func (o *MediaGalleryComponentResponse) GetItems() []MediaGalleryItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MediaGalleryComponentResponse) GetItemsOk() (*[]MediaGalleryItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MediaGalleryComponentResponse) SetItems(v []MediaGalleryItemResponse)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


