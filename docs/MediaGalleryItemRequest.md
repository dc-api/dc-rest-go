# MediaGalleryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Spoiler** | Pointer to **NullableBool** |  | [optional] 
**Media** | [**UnfurledMediaRequest**](UnfurledMediaRequest.md) |  | 

## Methods

### NewMediaGalleryItemRequest

`func NewMediaGalleryItemRequest(media UnfurledMediaRequest, ) *MediaGalleryItemRequest`

NewMediaGalleryItemRequest instantiates a new MediaGalleryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGalleryItemRequestWithDefaults

`func NewMediaGalleryItemRequestWithDefaults() *MediaGalleryItemRequest`

NewMediaGalleryItemRequestWithDefaults instantiates a new MediaGalleryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MediaGalleryItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MediaGalleryItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MediaGalleryItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MediaGalleryItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MediaGalleryItemRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MediaGalleryItemRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSpoiler

`func (o *MediaGalleryItemRequest) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *MediaGalleryItemRequest) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *MediaGalleryItemRequest) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *MediaGalleryItemRequest) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### SetSpoilerNil

`func (o *MediaGalleryItemRequest) SetSpoilerNil(b bool)`

 SetSpoilerNil sets the value for Spoiler to be an explicit nil

### UnsetSpoiler
`func (o *MediaGalleryItemRequest) UnsetSpoiler()`

UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil
### GetMedia

`func (o *MediaGalleryItemRequest) GetMedia() UnfurledMediaRequest`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MediaGalleryItemRequest) GetMediaOk() (*UnfurledMediaRequest, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MediaGalleryItemRequest) SetMedia(v UnfurledMediaRequest)`

SetMedia sets Media field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


