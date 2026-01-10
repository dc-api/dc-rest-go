# GuildHomeSettingsResponseNewMemberActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**ActionType** | [**NewMemberActionType**](NewMemberActionType.md) |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Emoji** | Pointer to [**SettingsEmojiResponse**](SettingsEmojiResponse.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewGuildHomeSettingsResponseNewMemberActionsInner

`func NewGuildHomeSettingsResponseNewMemberActionsInner(channelId string, actionType NewMemberActionType, title string, description string, ) *GuildHomeSettingsResponseNewMemberActionsInner`

NewGuildHomeSettingsResponseNewMemberActionsInner instantiates a new GuildHomeSettingsResponseNewMemberActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildHomeSettingsResponseNewMemberActionsInnerWithDefaults

`func NewGuildHomeSettingsResponseNewMemberActionsInnerWithDefaults() *GuildHomeSettingsResponseNewMemberActionsInner`

NewGuildHomeSettingsResponseNewMemberActionsInnerWithDefaults instantiates a new GuildHomeSettingsResponseNewMemberActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetActionType

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetActionType() NewMemberActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetActionTypeOk() (*NewMemberActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetActionType(v NewMemberActionType)`

SetActionType sets ActionType field to given value.


### GetTitle

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmoji

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetEmoji() SettingsEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetEmojiOk() (*SettingsEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetEmoji(v SettingsEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetIcon

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildHomeSettingsResponseNewMemberActionsInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


