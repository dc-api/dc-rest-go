# GuildChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**LastMessageId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **time.Time** |  | [optional] 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**ParentId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**RateLimitPerUser** | Pointer to **int32** |  | [optional] 
**Bitrate** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**RtcRegion** | Pointer to **string** |  | [optional] 
**VideoQualityMode** | Pointer to [**VideoQualityModes**](VideoQualityModes.md) |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to [**ThreadAutoArchiveDuration**](ThreadAutoArchiveDuration.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **int32** |  | [optional] 
**Position** | **int32** |  | 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteResponse**](ChannelPermissionOverwriteResponse.md) |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 
**AvailableTags** | Pointer to [**[]ForumTagResponse**](ForumTagResponse.md) |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**GuildChannelResponseDefaultReactionEmoji**](GuildChannelResponseDefaultReactionEmoji.md) |  | [optional] 
**DefaultSortOrder** | Pointer to [**CreateGuildChannelRequestDefaultSortOrder**](CreateGuildChannelRequestDefaultSortOrder.md) |  | [optional] 
**DefaultForumLayout** | Pointer to [**CreateGuildChannelRequestDefaultForumLayout**](CreateGuildChannelRequestDefaultForumLayout.md) |  | [optional] 
**DefaultTagSetting** | Pointer to [**CreateGuildChannelRequestDefaultTagSetting**](CreateGuildChannelRequestDefaultTagSetting.md) |  | [optional] 
**HdStreamingUntil** | Pointer to **time.Time** |  | [optional] 
**HdStreamingBuyerId** | Pointer to **string** |  | [optional] 

## Methods

### NewGuildChannelResponse

`func NewGuildChannelResponse(id string, type_ ChannelTypes, flags int32, guildId string, name string, position int32, ) *GuildChannelResponse`

NewGuildChannelResponse instantiates a new GuildChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildChannelResponseWithDefaults

`func NewGuildChannelResponseWithDefaults() *GuildChannelResponse`

NewGuildChannelResponseWithDefaults instantiates a new GuildChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildChannelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildChannelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildChannelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GuildChannelResponse) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildChannelResponse) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildChannelResponse) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *GuildChannelResponse) GetLastMessageId() GetEntitlementsSkuIdsParameterOneOfInner`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *GuildChannelResponse) GetLastMessageIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *GuildChannelResponse) SetLastMessageId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *GuildChannelResponse) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *GuildChannelResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GuildChannelResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GuildChannelResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *GuildChannelResponse) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *GuildChannelResponse) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *GuildChannelResponse) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *GuildChannelResponse) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### GetGuildId

`func (o *GuildChannelResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GuildChannelResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GuildChannelResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *GuildChannelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildChannelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildChannelResponse) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *GuildChannelResponse) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GuildChannelResponse) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GuildChannelResponse) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GuildChannelResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *GuildChannelResponse) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *GuildChannelResponse) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *GuildChannelResponse) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *GuildChannelResponse) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### GetBitrate

`func (o *GuildChannelResponse) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *GuildChannelResponse) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *GuildChannelResponse) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *GuildChannelResponse) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetUserLimit

`func (o *GuildChannelResponse) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *GuildChannelResponse) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *GuildChannelResponse) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *GuildChannelResponse) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetRtcRegion

`func (o *GuildChannelResponse) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *GuildChannelResponse) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *GuildChannelResponse) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *GuildChannelResponse) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### GetVideoQualityMode

`func (o *GuildChannelResponse) GetVideoQualityMode() VideoQualityModes`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *GuildChannelResponse) GetVideoQualityModeOk() (*VideoQualityModes, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *GuildChannelResponse) SetVideoQualityMode(v VideoQualityModes)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *GuildChannelResponse) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetPermissions

`func (o *GuildChannelResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GuildChannelResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GuildChannelResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GuildChannelResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTopic

`func (o *GuildChannelResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GuildChannelResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GuildChannelResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GuildChannelResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *GuildChannelResponse) GetDefaultAutoArchiveDuration() ThreadAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *GuildChannelResponse) GetDefaultAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *GuildChannelResponse) SetDefaultAutoArchiveDuration(v ThreadAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *GuildChannelResponse) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *GuildChannelResponse) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *GuildChannelResponse) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *GuildChannelResponse) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *GuildChannelResponse) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### GetPosition

`func (o *GuildChannelResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GuildChannelResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GuildChannelResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetPermissionOverwrites

`func (o *GuildChannelResponse) GetPermissionOverwrites() []ChannelPermissionOverwriteResponse`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *GuildChannelResponse) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteResponse, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *GuildChannelResponse) SetPermissionOverwrites(v []ChannelPermissionOverwriteResponse)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *GuildChannelResponse) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### GetNsfw

`func (o *GuildChannelResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GuildChannelResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GuildChannelResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *GuildChannelResponse) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetAvailableTags

`func (o *GuildChannelResponse) GetAvailableTags() []ForumTagResponse`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *GuildChannelResponse) GetAvailableTagsOk() (*[]ForumTagResponse, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *GuildChannelResponse) SetAvailableTags(v []ForumTagResponse)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *GuildChannelResponse) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *GuildChannelResponse) GetDefaultReactionEmoji() GuildChannelResponseDefaultReactionEmoji`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *GuildChannelResponse) GetDefaultReactionEmojiOk() (*GuildChannelResponseDefaultReactionEmoji, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *GuildChannelResponse) SetDefaultReactionEmoji(v GuildChannelResponseDefaultReactionEmoji)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *GuildChannelResponse) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultSortOrder

`func (o *GuildChannelResponse) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *GuildChannelResponse) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *GuildChannelResponse) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *GuildChannelResponse) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *GuildChannelResponse) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *GuildChannelResponse) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *GuildChannelResponse) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *GuildChannelResponse) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *GuildChannelResponse) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *GuildChannelResponse) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *GuildChannelResponse) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *GuildChannelResponse) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetHdStreamingUntil

`func (o *GuildChannelResponse) GetHdStreamingUntil() time.Time`

GetHdStreamingUntil returns the HdStreamingUntil field if non-nil, zero value otherwise.

### GetHdStreamingUntilOk

`func (o *GuildChannelResponse) GetHdStreamingUntilOk() (*time.Time, bool)`

GetHdStreamingUntilOk returns a tuple with the HdStreamingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdStreamingUntil

`func (o *GuildChannelResponse) SetHdStreamingUntil(v time.Time)`

SetHdStreamingUntil sets HdStreamingUntil field to given value.

### HasHdStreamingUntil

`func (o *GuildChannelResponse) HasHdStreamingUntil() bool`

HasHdStreamingUntil returns a boolean if a field has been set.

### GetHdStreamingBuyerId

`func (o *GuildChannelResponse) GetHdStreamingBuyerId() string`

GetHdStreamingBuyerId returns the HdStreamingBuyerId field if non-nil, zero value otherwise.

### GetHdStreamingBuyerIdOk

`func (o *GuildChannelResponse) GetHdStreamingBuyerIdOk() (*string, bool)`

GetHdStreamingBuyerIdOk returns a tuple with the HdStreamingBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdStreamingBuyerId

`func (o *GuildChannelResponse) SetHdStreamingBuyerId(v string)`

SetHdStreamingBuyerId sets HdStreamingBuyerId field to given value.

### HasHdStreamingBuyerId

`func (o *GuildChannelResponse) HasHdStreamingBuyerId() bool`

HasHdStreamingBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


