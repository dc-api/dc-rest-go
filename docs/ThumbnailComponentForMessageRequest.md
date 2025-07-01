# ThumbnailComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Spoiler** | Pointer to **NullableBool** |  | [optional] 
**Media** | [**UnfurledMediaRequest**](UnfurledMediaRequest.md) |  | 

## Methods

### NewThumbnailComponentForMessageRequest

`func NewThumbnailComponentForMessageRequest(type_ int32, media UnfurledMediaRequest, ) *ThumbnailComponentForMessageRequest`

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

`func (o *ThumbnailComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThumbnailComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThumbnailComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


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

### SetDescriptionNil

`func (o *ThumbnailComponentForMessageRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ThumbnailComponentForMessageRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetSpoilerNil

`func (o *ThumbnailComponentForMessageRequest) SetSpoilerNil(b bool)`

 SetSpoilerNil sets the value for Spoiler to be an explicit nil

### UnsetSpoiler
`func (o *ThumbnailComponentForMessageRequest) UnsetSpoiler()`

UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil
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


