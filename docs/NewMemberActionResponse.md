# NewMemberActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**ActionType** | **NullableInt32** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Emoji** | Pointer to [**NullableSettingsEmojiResponse**](SettingsEmojiResponse.md) |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewMemberActionResponse

`func NewNewMemberActionResponse(channelId string, actionType NullableInt32, title string, description string, ) *NewMemberActionResponse`

NewNewMemberActionResponse instantiates a new NewMemberActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMemberActionResponseWithDefaults

`func NewNewMemberActionResponseWithDefaults() *NewMemberActionResponse`

NewNewMemberActionResponseWithDefaults instantiates a new NewMemberActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *NewMemberActionResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *NewMemberActionResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *NewMemberActionResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetActionType

`func (o *NewMemberActionResponse) GetActionType() int32`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *NewMemberActionResponse) GetActionTypeOk() (*int32, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *NewMemberActionResponse) SetActionType(v int32)`

SetActionType sets ActionType field to given value.


### SetActionTypeNil

`func (o *NewMemberActionResponse) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *NewMemberActionResponse) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetTitle

`func (o *NewMemberActionResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewMemberActionResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewMemberActionResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewMemberActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewMemberActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewMemberActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmoji

`func (o *NewMemberActionResponse) GetEmoji() SettingsEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *NewMemberActionResponse) GetEmojiOk() (*SettingsEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *NewMemberActionResponse) SetEmoji(v SettingsEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *NewMemberActionResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *NewMemberActionResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *NewMemberActionResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetIcon

`func (o *NewMemberActionResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NewMemberActionResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NewMemberActionResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *NewMemberActionResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *NewMemberActionResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *NewMemberActionResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


