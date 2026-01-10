# GuildHomeSettingsResponseResourceChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**Title** | **string** |  | 
**Emoji** | Pointer to [**SettingsEmojiResponse**](SettingsEmojiResponse.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 

## Methods

### NewGuildHomeSettingsResponseResourceChannelsInner

`func NewGuildHomeSettingsResponseResourceChannelsInner(channelId string, title string, description string, ) *GuildHomeSettingsResponseResourceChannelsInner`

NewGuildHomeSettingsResponseResourceChannelsInner instantiates a new GuildHomeSettingsResponseResourceChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildHomeSettingsResponseResourceChannelsInnerWithDefaults

`func NewGuildHomeSettingsResponseResourceChannelsInnerWithDefaults() *GuildHomeSettingsResponseResourceChannelsInner`

NewGuildHomeSettingsResponseResourceChannelsInnerWithDefaults instantiates a new GuildHomeSettingsResponseResourceChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GuildHomeSettingsResponseResourceChannelsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetTitle

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuildHomeSettingsResponseResourceChannelsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEmoji

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetEmoji() SettingsEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetEmojiOk() (*SettingsEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *GuildHomeSettingsResponseResourceChannelsInner) SetEmoji(v SettingsEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *GuildHomeSettingsResponseResourceChannelsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetIcon

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildHomeSettingsResponseResourceChannelsInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildHomeSettingsResponseResourceChannelsInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDescription

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildHomeSettingsResponseResourceChannelsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildHomeSettingsResponseResourceChannelsInner) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


