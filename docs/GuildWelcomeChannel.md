# GuildWelcomeChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**Description** | **string** |  | 
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGuildWelcomeChannel

`func NewGuildWelcomeChannel(channelId string, description string, ) *GuildWelcomeChannel`

NewGuildWelcomeChannel instantiates a new GuildWelcomeChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildWelcomeChannelWithDefaults

`func NewGuildWelcomeChannelWithDefaults() *GuildWelcomeChannel`

NewGuildWelcomeChannelWithDefaults instantiates a new GuildWelcomeChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *GuildWelcomeChannel) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GuildWelcomeChannel) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GuildWelcomeChannel) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetDescription

`func (o *GuildWelcomeChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildWelcomeChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildWelcomeChannel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmojiId

`func (o *GuildWelcomeChannel) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *GuildWelcomeChannel) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *GuildWelcomeChannel) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *GuildWelcomeChannel) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *GuildWelcomeChannel) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *GuildWelcomeChannel) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *GuildWelcomeChannel) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *GuildWelcomeChannel) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *GuildWelcomeChannel) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *GuildWelcomeChannel) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


