# UpdateGuildChannelRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Topic** | Pointer to **NullableString** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**UserLimit** | Pointer to **NullableInt32** |  | [optional] 
**Nsfw** | Pointer to **NullableBool** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteRequest**](ChannelPermissionOverwriteRequest.md) |  | [optional] 
**RtcRegion** | Pointer to **NullableString** |  | [optional] 
**VideoQualityMode** | Pointer to **int32** |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to **int32** |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**NullableUpdateDefaultReactionEmojiRequest**](UpdateDefaultReactionEmojiRequest.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSortOrder** | Pointer to **int32** |  | [optional] 
**DefaultForumLayout** | Pointer to **int32** |  | [optional] 
**DefaultTagSetting** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**AvailableTags** | Pointer to [**[]UpdateThreadTagRequest**](UpdateThreadTagRequest.md) |  | [optional] 

## Methods

### NewUpdateGuildChannelRequestPartial

`func NewUpdateGuildChannelRequestPartial() *UpdateGuildChannelRequestPartial`

NewUpdateGuildChannelRequestPartial instantiates a new UpdateGuildChannelRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildChannelRequestPartialWithDefaults

`func NewUpdateGuildChannelRequestPartialWithDefaults() *UpdateGuildChannelRequestPartial`

NewUpdateGuildChannelRequestPartialWithDefaults instantiates a new UpdateGuildChannelRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateGuildChannelRequestPartial) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateGuildChannelRequestPartial) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateGuildChannelRequestPartial) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateGuildChannelRequestPartial) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdateGuildChannelRequestPartial) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateGuildChannelRequestPartial) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *UpdateGuildChannelRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGuildChannelRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGuildChannelRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGuildChannelRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateGuildChannelRequestPartial) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateGuildChannelRequestPartial) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateGuildChannelRequestPartial) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateGuildChannelRequestPartial) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *UpdateGuildChannelRequestPartial) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *UpdateGuildChannelRequestPartial) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetTopic

`func (o *UpdateGuildChannelRequestPartial) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *UpdateGuildChannelRequestPartial) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *UpdateGuildChannelRequestPartial) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *UpdateGuildChannelRequestPartial) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *UpdateGuildChannelRequestPartial) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *UpdateGuildChannelRequestPartial) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetBitrate

`func (o *UpdateGuildChannelRequestPartial) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *UpdateGuildChannelRequestPartial) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *UpdateGuildChannelRequestPartial) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *UpdateGuildChannelRequestPartial) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *UpdateGuildChannelRequestPartial) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *UpdateGuildChannelRequestPartial) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *UpdateGuildChannelRequestPartial) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UpdateGuildChannelRequestPartial) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UpdateGuildChannelRequestPartial) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UpdateGuildChannelRequestPartial) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *UpdateGuildChannelRequestPartial) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *UpdateGuildChannelRequestPartial) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetNsfw

`func (o *UpdateGuildChannelRequestPartial) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *UpdateGuildChannelRequestPartial) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *UpdateGuildChannelRequestPartial) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *UpdateGuildChannelRequestPartial) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### SetNsfwNil

`func (o *UpdateGuildChannelRequestPartial) SetNsfwNil(b bool)`

 SetNsfwNil sets the value for Nsfw to be an explicit nil

### UnsetNsfw
`func (o *UpdateGuildChannelRequestPartial) UnsetNsfw()`

UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
### GetRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *UpdateGuildChannelRequestPartial) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *UpdateGuildChannelRequestPartial) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *UpdateGuildChannelRequestPartial) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetParentId

`func (o *UpdateGuildChannelRequestPartial) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateGuildChannelRequestPartial) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateGuildChannelRequestPartial) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateGuildChannelRequestPartial) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermissionOverwrites

`func (o *UpdateGuildChannelRequestPartial) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *UpdateGuildChannelRequestPartial) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteRequest, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *UpdateGuildChannelRequestPartial) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *UpdateGuildChannelRequestPartial) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### SetPermissionOverwritesNil

`func (o *UpdateGuildChannelRequestPartial) SetPermissionOverwritesNil(b bool)`

 SetPermissionOverwritesNil sets the value for PermissionOverwrites to be an explicit nil

### UnsetPermissionOverwrites
`func (o *UpdateGuildChannelRequestPartial) UnsetPermissionOverwrites()`

UnsetPermissionOverwrites ensures that no value is present for PermissionOverwrites, not even an explicit nil
### GetRtcRegion

`func (o *UpdateGuildChannelRequestPartial) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *UpdateGuildChannelRequestPartial) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *UpdateGuildChannelRequestPartial) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *UpdateGuildChannelRequestPartial) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *UpdateGuildChannelRequestPartial) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *UpdateGuildChannelRequestPartial) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) GetVideoQualityMode() int32`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *UpdateGuildChannelRequestPartial) GetVideoQualityModeOk() (*int32, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) SetVideoQualityMode(v int32)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDuration() int32`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDurationOk() (*int32, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) SetDefaultAutoArchiveDuration(v int32)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmoji() UpdateDefaultReactionEmojiRequest`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmojiOk() (*UpdateDefaultReactionEmojiRequest, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) SetDefaultReactionEmoji(v UpdateDefaultReactionEmojiRequest)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### SetDefaultReactionEmojiNil

`func (o *UpdateGuildChannelRequestPartial) SetDefaultReactionEmojiNil(b bool)`

 SetDefaultReactionEmojiNil sets the value for DefaultReactionEmoji to be an explicit nil

### UnsetDefaultReactionEmoji
`func (o *UpdateGuildChannelRequestPartial) UnsetDefaultReactionEmoji()`

UnsetDefaultReactionEmoji ensures that no value is present for DefaultReactionEmoji, not even an explicit nil
### GetDefaultThreadRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *UpdateGuildChannelRequestPartial) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### SetDefaultThreadRateLimitPerUserNil

`func (o *UpdateGuildChannelRequestPartial) SetDefaultThreadRateLimitPerUserNil(b bool)`

 SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil

### UnsetDefaultThreadRateLimitPerUser
`func (o *UpdateGuildChannelRequestPartial) UnsetDefaultThreadRateLimitPerUser()`

UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
### GetDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrder() int32`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrderOk() (*int32, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) SetDefaultSortOrder(v int32)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayout() int32`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayoutOk() (*int32, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) SetDefaultForumLayout(v int32)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSetting() string`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSettingOk() (*string, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *UpdateGuildChannelRequestPartial) SetDefaultTagSetting(v string)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *UpdateGuildChannelRequestPartial) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetFlags

`func (o *UpdateGuildChannelRequestPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateGuildChannelRequestPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateGuildChannelRequestPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateGuildChannelRequestPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *UpdateGuildChannelRequestPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *UpdateGuildChannelRequestPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAvailableTags

`func (o *UpdateGuildChannelRequestPartial) GetAvailableTags() []UpdateThreadTagRequest`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *UpdateGuildChannelRequestPartial) GetAvailableTagsOk() (*[]UpdateThreadTagRequest, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *UpdateGuildChannelRequestPartial) SetAvailableTags(v []UpdateThreadTagRequest)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *UpdateGuildChannelRequestPartial) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### SetAvailableTagsNil

`func (o *UpdateGuildChannelRequestPartial) SetAvailableTagsNil(b bool)`

 SetAvailableTagsNil sets the value for AvailableTags to be an explicit nil

### UnsetAvailableTags
`func (o *UpdateGuildChannelRequestPartial) UnsetAvailableTags()`

UnsetAvailableTags ensures that no value is present for AvailableTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


