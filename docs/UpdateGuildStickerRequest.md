# UpdateGuildStickerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateGuildStickerRequest

`func NewUpdateGuildStickerRequest() *UpdateGuildStickerRequest`

NewUpdateGuildStickerRequest instantiates a new UpdateGuildStickerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildStickerRequestWithDefaults

`func NewUpdateGuildStickerRequestWithDefaults() *UpdateGuildStickerRequest`

NewUpdateGuildStickerRequestWithDefaults instantiates a new UpdateGuildStickerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGuildStickerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGuildStickerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGuildStickerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGuildStickerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *UpdateGuildStickerRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateGuildStickerRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateGuildStickerRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateGuildStickerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGuildStickerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGuildStickerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGuildStickerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGuildStickerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateGuildStickerRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateGuildStickerRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


