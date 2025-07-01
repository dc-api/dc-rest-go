# MediaGalleryItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Spoiler** | **bool** |  | 

## Methods

### NewMediaGalleryItemResponse

`func NewMediaGalleryItemResponse(media UnfurledMediaResponse, spoiler bool, ) *MediaGalleryItemResponse`

NewMediaGalleryItemResponse instantiates a new MediaGalleryItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGalleryItemResponseWithDefaults

`func NewMediaGalleryItemResponseWithDefaults() *MediaGalleryItemResponse`

NewMediaGalleryItemResponseWithDefaults instantiates a new MediaGalleryItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *MediaGalleryItemResponse) GetMedia() UnfurledMediaResponse`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MediaGalleryItemResponse) GetMediaOk() (*UnfurledMediaResponse, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MediaGalleryItemResponse) SetMedia(v UnfurledMediaResponse)`

SetMedia sets Media field to given value.


### GetDescription

`func (o *MediaGalleryItemResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MediaGalleryItemResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MediaGalleryItemResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MediaGalleryItemResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MediaGalleryItemResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MediaGalleryItemResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSpoiler

`func (o *MediaGalleryItemResponse) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *MediaGalleryItemResponse) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *MediaGalleryItemResponse) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


