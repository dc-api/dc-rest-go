# MediaGalleryComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Items** | [**[]MediaGalleryItemRequest**](MediaGalleryItemRequest.md) |  | 

## Methods

### NewMediaGalleryComponentForMessageRequest

`func NewMediaGalleryComponentForMessageRequest(type_ MessageComponentTypes, items []MediaGalleryItemRequest, ) *MediaGalleryComponentForMessageRequest`

NewMediaGalleryComponentForMessageRequest instantiates a new MediaGalleryComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGalleryComponentForMessageRequestWithDefaults

`func NewMediaGalleryComponentForMessageRequestWithDefaults() *MediaGalleryComponentForMessageRequest`

NewMediaGalleryComponentForMessageRequestWithDefaults instantiates a new MediaGalleryComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MediaGalleryComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaGalleryComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaGalleryComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MediaGalleryComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaGalleryComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaGalleryComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MediaGalleryComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *MediaGalleryComponentForMessageRequest) GetItems() []MediaGalleryItemRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MediaGalleryComponentForMessageRequest) GetItemsOk() (*[]MediaGalleryItemRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MediaGalleryComponentForMessageRequest) SetItems(v []MediaGalleryItemRequest)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


