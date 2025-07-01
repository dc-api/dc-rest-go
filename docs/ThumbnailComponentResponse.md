# ThumbnailComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**Media** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Spoiler** | **bool** |  | 

## Methods

### NewThumbnailComponentResponse

`func NewThumbnailComponentResponse(type_ int32, id int32, media UnfurledMediaResponse, spoiler bool, ) *ThumbnailComponentResponse`

NewThumbnailComponentResponse instantiates a new ThumbnailComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailComponentResponseWithDefaults

`func NewThumbnailComponentResponseWithDefaults() *ThumbnailComponentResponse`

NewThumbnailComponentResponseWithDefaults instantiates a new ThumbnailComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThumbnailComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThumbnailComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThumbnailComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ThumbnailComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThumbnailComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThumbnailComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetMedia

`func (o *ThumbnailComponentResponse) GetMedia() UnfurledMediaResponse`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ThumbnailComponentResponse) GetMediaOk() (*UnfurledMediaResponse, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ThumbnailComponentResponse) SetMedia(v UnfurledMediaResponse)`

SetMedia sets Media field to given value.


### GetDescription

`func (o *ThumbnailComponentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThumbnailComponentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThumbnailComponentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThumbnailComponentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ThumbnailComponentResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ThumbnailComponentResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSpoiler

`func (o *ThumbnailComponentResponse) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ThumbnailComponentResponse) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ThumbnailComponentResponse) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


