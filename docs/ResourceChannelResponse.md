# ResourceChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**Title** | **string** |  | 
**Emoji** | Pointer to [**NullableSettingsEmojiResponse**](SettingsEmojiResponse.md) |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 

## Methods

### NewResourceChannelResponse

`func NewResourceChannelResponse(channelId string, title string, description string, ) *ResourceChannelResponse`

NewResourceChannelResponse instantiates a new ResourceChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceChannelResponseWithDefaults

`func NewResourceChannelResponseWithDefaults() *ResourceChannelResponse`

NewResourceChannelResponseWithDefaults instantiates a new ResourceChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *ResourceChannelResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ResourceChannelResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ResourceChannelResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetTitle

`func (o *ResourceChannelResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourceChannelResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourceChannelResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEmoji

`func (o *ResourceChannelResponse) GetEmoji() SettingsEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ResourceChannelResponse) GetEmojiOk() (*SettingsEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ResourceChannelResponse) SetEmoji(v SettingsEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ResourceChannelResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *ResourceChannelResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *ResourceChannelResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetIcon

`func (o *ResourceChannelResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ResourceChannelResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ResourceChannelResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ResourceChannelResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ResourceChannelResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ResourceChannelResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *ResourceChannelResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceChannelResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceChannelResponse) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


