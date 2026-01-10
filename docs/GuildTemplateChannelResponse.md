# GuildTemplateChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Bitrate** | **int32** |  | 
**UserLimit** | **int32** |  | 
**Nsfw** | **bool** |  | 
**RateLimitPerUser** | **int32** |  | 
**ParentId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to [**CreateForumThreadRequestAutoArchiveDuration**](CreateForumThreadRequestAutoArchiveDuration.md) |  | [optional] 
**PermissionOverwrites** | [**[]GuildTemplateChannelResponsePermissionOverwritesInner**](GuildTemplateChannelResponsePermissionOverwritesInner.md) |  | 
**AvailableTags** | Pointer to [**[]GuildTemplateChannelTags**](GuildTemplateChannelTags.md) |  | [optional] 
**Template** | **string** |  | 
**DefaultReactionEmoji** | Pointer to [**GuildChannelResponseDefaultReactionEmoji**](GuildChannelResponseDefaultReactionEmoji.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **int32** |  | [optional] 
**DefaultSortOrder** | Pointer to [**CreateGuildChannelRequestDefaultSortOrder**](CreateGuildChannelRequestDefaultSortOrder.md) |  | [optional] 
**DefaultForumLayout** | Pointer to [**CreateGuildChannelRequestDefaultForumLayout**](CreateGuildChannelRequestDefaultForumLayout.md) |  | [optional] 
**DefaultTagSetting** | Pointer to [**CreateGuildChannelRequestDefaultTagSetting**](CreateGuildChannelRequestDefaultTagSetting.md) |  | [optional] 
**IconEmoji** | Pointer to [**GuildTemplateChannelResponseIconEmoji**](GuildTemplateChannelResponseIconEmoji.md) |  | [optional] 
**ThemeColor** | Pointer to **int32** |  | [optional] 

## Methods

### NewGuildTemplateChannelResponse

`func NewGuildTemplateChannelResponse(type_ ChannelTypes, bitrate int32, userLimit int32, nsfw bool, rateLimitPerUser int32, permissionOverwrites []GuildTemplateChannelResponsePermissionOverwritesInner, template string, ) *GuildTemplateChannelResponse`

NewGuildTemplateChannelResponse instantiates a new GuildTemplateChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateChannelResponseWithDefaults

`func NewGuildTemplateChannelResponseWithDefaults() *GuildTemplateChannelResponse`

NewGuildTemplateChannelResponseWithDefaults instantiates a new GuildTemplateChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildTemplateChannelResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildTemplateChannelResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildTemplateChannelResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GuildTemplateChannelResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GuildTemplateChannelResponse) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildTemplateChannelResponse) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildTemplateChannelResponse) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetName

`func (o *GuildTemplateChannelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildTemplateChannelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildTemplateChannelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuildTemplateChannelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *GuildTemplateChannelResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GuildTemplateChannelResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GuildTemplateChannelResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GuildTemplateChannelResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTopic

`func (o *GuildTemplateChannelResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GuildTemplateChannelResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GuildTemplateChannelResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GuildTemplateChannelResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetBitrate

`func (o *GuildTemplateChannelResponse) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *GuildTemplateChannelResponse) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *GuildTemplateChannelResponse) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.


### GetUserLimit

`func (o *GuildTemplateChannelResponse) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *GuildTemplateChannelResponse) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *GuildTemplateChannelResponse) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.


### GetNsfw

`func (o *GuildTemplateChannelResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GuildTemplateChannelResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GuildTemplateChannelResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.


### GetRateLimitPerUser

`func (o *GuildTemplateChannelResponse) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *GuildTemplateChannelResponse) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *GuildTemplateChannelResponse) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.


### GetParentId

`func (o *GuildTemplateChannelResponse) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GuildTemplateChannelResponse) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GuildTemplateChannelResponse) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GuildTemplateChannelResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *GuildTemplateChannelResponse) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *GuildTemplateChannelResponse) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetPermissionOverwrites

`func (o *GuildTemplateChannelResponse) GetPermissionOverwrites() []GuildTemplateChannelResponsePermissionOverwritesInner`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *GuildTemplateChannelResponse) GetPermissionOverwritesOk() (*[]GuildTemplateChannelResponsePermissionOverwritesInner, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *GuildTemplateChannelResponse) SetPermissionOverwrites(v []GuildTemplateChannelResponsePermissionOverwritesInner)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.


### GetAvailableTags

`func (o *GuildTemplateChannelResponse) GetAvailableTags() []GuildTemplateChannelTags`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *GuildTemplateChannelResponse) GetAvailableTagsOk() (*[]GuildTemplateChannelTags, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *GuildTemplateChannelResponse) SetAvailableTags(v []GuildTemplateChannelTags)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *GuildTemplateChannelResponse) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GuildTemplateChannelResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GuildTemplateChannelResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GuildTemplateChannelResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetDefaultReactionEmoji

`func (o *GuildTemplateChannelResponse) GetDefaultReactionEmoji() GuildChannelResponseDefaultReactionEmoji`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *GuildTemplateChannelResponse) GetDefaultReactionEmojiOk() (*GuildChannelResponseDefaultReactionEmoji, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *GuildTemplateChannelResponse) SetDefaultReactionEmoji(v GuildChannelResponseDefaultReactionEmoji)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *GuildTemplateChannelResponse) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *GuildTemplateChannelResponse) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *GuildTemplateChannelResponse) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### GetDefaultSortOrder

`func (o *GuildTemplateChannelResponse) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *GuildTemplateChannelResponse) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *GuildTemplateChannelResponse) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *GuildTemplateChannelResponse) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *GuildTemplateChannelResponse) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *GuildTemplateChannelResponse) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *GuildTemplateChannelResponse) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *GuildTemplateChannelResponse) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *GuildTemplateChannelResponse) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *GuildTemplateChannelResponse) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *GuildTemplateChannelResponse) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *GuildTemplateChannelResponse) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetIconEmoji

`func (o *GuildTemplateChannelResponse) GetIconEmoji() GuildTemplateChannelResponseIconEmoji`

GetIconEmoji returns the IconEmoji field if non-nil, zero value otherwise.

### GetIconEmojiOk

`func (o *GuildTemplateChannelResponse) GetIconEmojiOk() (*GuildTemplateChannelResponseIconEmoji, bool)`

GetIconEmojiOk returns a tuple with the IconEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconEmoji

`func (o *GuildTemplateChannelResponse) SetIconEmoji(v GuildTemplateChannelResponseIconEmoji)`

SetIconEmoji sets IconEmoji field to given value.

### HasIconEmoji

`func (o *GuildTemplateChannelResponse) HasIconEmoji() bool`

HasIconEmoji returns a boolean if a field has been set.

### GetThemeColor

`func (o *GuildTemplateChannelResponse) GetThemeColor() int32`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *GuildTemplateChannelResponse) GetThemeColorOk() (*int32, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColor

`func (o *GuildTemplateChannelResponse) SetThemeColor(v int32)`

SetThemeColor sets ThemeColor field to given value.

### HasThemeColor

`func (o *GuildTemplateChannelResponse) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


