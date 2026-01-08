# ThumbnailComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Spoiler** | Pointer to **bool** |  | [optional] 
**Media** | [**UnfurledMediaRequest**](UnfurledMediaRequest.md) |  | 

## Methods

### NewThumbnailComponentForMessageRequest

`func NewThumbnailComponentForMessageRequest(type_ MessageComponentTypes, media UnfurledMediaRequest, ) *ThumbnailComponentForMessageRequest`

NewThumbnailComponentForMessageRequest instantiates a new ThumbnailComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailComponentForMessageRequestWithDefaults

`func NewThumbnailComponentForMessageRequestWithDefaults() *ThumbnailComponentForMessageRequest`

NewThumbnailComponentForMessageRequestWithDefaults instantiates a new ThumbnailComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThumbnailComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThumbnailComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThumbnailComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ThumbnailComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThumbnailComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThumbnailComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ThumbnailComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ThumbnailComponentForMessageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThumbnailComponentForMessageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThumbnailComponentForMessageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThumbnailComponentForMessageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpoiler

`func (o *ThumbnailComponentForMessageRequest) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ThumbnailComponentForMessageRequest) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ThumbnailComponentForMessageRequest) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *ThumbnailComponentForMessageRequest) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetMedia

`func (o *ThumbnailComponentForMessageRequest) GetMedia() UnfurledMediaRequest`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ThumbnailComponentForMessageRequest) GetMediaOk() (*UnfurledMediaRequest, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ThumbnailComponentForMessageRequest) SetMedia(v UnfurledMediaRequest)`

SetMedia sets Media field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


