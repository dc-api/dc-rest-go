# CreateGuildRequestChannelItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
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
**Id** | Pointer to **string** |  | [optional] 
**AvailableTags** | Pointer to [**[]CreateOrUpdateThreadTagRequest**](CreateOrUpdateThreadTagRequest.md) |  | [optional] 

## Methods

### NewCreateGuildRequestChannelItem

`func NewCreateGuildRequestChannelItem(name string, ) *CreateGuildRequestChannelItem`

NewCreateGuildRequestChannelItem instantiates a new CreateGuildRequestChannelItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildRequestChannelItemWithDefaults

`func NewCreateGuildRequestChannelItemWithDefaults() *CreateGuildRequestChannelItem`

NewCreateGuildRequestChannelItemWithDefaults instantiates a new CreateGuildRequestChannelItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateGuildRequestChannelItem) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGuildRequestChannelItem) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGuildRequestChannelItem) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateGuildRequestChannelItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateGuildRequestChannelItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateGuildRequestChannelItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *CreateGuildRequestChannelItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildRequestChannelItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildRequestChannelItem) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *CreateGuildRequestChannelItem) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateGuildRequestChannelItem) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateGuildRequestChannelItem) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateGuildRequestChannelItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *CreateGuildRequestChannelItem) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CreateGuildRequestChannelItem) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetTopic

`func (o *CreateGuildRequestChannelItem) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CreateGuildRequestChannelItem) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CreateGuildRequestChannelItem) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *CreateGuildRequestChannelItem) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *CreateGuildRequestChannelItem) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *CreateGuildRequestChannelItem) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetBitrate

`func (o *CreateGuildRequestChannelItem) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *CreateGuildRequestChannelItem) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *CreateGuildRequestChannelItem) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *CreateGuildRequestChannelItem) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *CreateGuildRequestChannelItem) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *CreateGuildRequestChannelItem) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *CreateGuildRequestChannelItem) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *CreateGuildRequestChannelItem) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *CreateGuildRequestChannelItem) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *CreateGuildRequestChannelItem) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *CreateGuildRequestChannelItem) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *CreateGuildRequestChannelItem) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetNsfw

`func (o *CreateGuildRequestChannelItem) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *CreateGuildRequestChannelItem) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *CreateGuildRequestChannelItem) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *CreateGuildRequestChannelItem) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### SetNsfwNil

`func (o *CreateGuildRequestChannelItem) SetNsfwNil(b bool)`

 SetNsfwNil sets the value for Nsfw to be an explicit nil

### UnsetNsfw
`func (o *CreateGuildRequestChannelItem) UnsetNsfw()`

UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
### GetRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreateGuildRequestChannelItem) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *CreateGuildRequestChannelItem) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *CreateGuildRequestChannelItem) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetParentId

`func (o *CreateGuildRequestChannelItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateGuildRequestChannelItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateGuildRequestChannelItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateGuildRequestChannelItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermissionOverwrites

`func (o *CreateGuildRequestChannelItem) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *CreateGuildRequestChannelItem) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteRequest, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *CreateGuildRequestChannelItem) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *CreateGuildRequestChannelItem) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### SetPermissionOverwritesNil

`func (o *CreateGuildRequestChannelItem) SetPermissionOverwritesNil(b bool)`

 SetPermissionOverwritesNil sets the value for PermissionOverwrites to be an explicit nil

### UnsetPermissionOverwrites
`func (o *CreateGuildRequestChannelItem) UnsetPermissionOverwrites()`

UnsetPermissionOverwrites ensures that no value is present for PermissionOverwrites, not even an explicit nil
### GetRtcRegion

`func (o *CreateGuildRequestChannelItem) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *CreateGuildRequestChannelItem) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *CreateGuildRequestChannelItem) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *CreateGuildRequestChannelItem) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *CreateGuildRequestChannelItem) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *CreateGuildRequestChannelItem) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *CreateGuildRequestChannelItem) GetVideoQualityMode() int32`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *CreateGuildRequestChannelItem) GetVideoQualityModeOk() (*int32, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *CreateGuildRequestChannelItem) SetVideoQualityMode(v int32)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *CreateGuildRequestChannelItem) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *CreateGuildRequestChannelItem) GetDefaultAutoArchiveDuration() int32`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *CreateGuildRequestChannelItem) GetDefaultAutoArchiveDurationOk() (*int32, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *CreateGuildRequestChannelItem) SetDefaultAutoArchiveDuration(v int32)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *CreateGuildRequestChannelItem) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *CreateGuildRequestChannelItem) GetDefaultReactionEmoji() UpdateDefaultReactionEmojiRequest`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *CreateGuildRequestChannelItem) GetDefaultReactionEmojiOk() (*UpdateDefaultReactionEmojiRequest, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *CreateGuildRequestChannelItem) SetDefaultReactionEmoji(v UpdateDefaultReactionEmojiRequest)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *CreateGuildRequestChannelItem) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### SetDefaultReactionEmojiNil

`func (o *CreateGuildRequestChannelItem) SetDefaultReactionEmojiNil(b bool)`

 SetDefaultReactionEmojiNil sets the value for DefaultReactionEmoji to be an explicit nil

### UnsetDefaultReactionEmoji
`func (o *CreateGuildRequestChannelItem) UnsetDefaultReactionEmoji()`

UnsetDefaultReactionEmoji ensures that no value is present for DefaultReactionEmoji, not even an explicit nil
### GetDefaultThreadRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *CreateGuildRequestChannelItem) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *CreateGuildRequestChannelItem) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### SetDefaultThreadRateLimitPerUserNil

`func (o *CreateGuildRequestChannelItem) SetDefaultThreadRateLimitPerUserNil(b bool)`

 SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil

### UnsetDefaultThreadRateLimitPerUser
`func (o *CreateGuildRequestChannelItem) UnsetDefaultThreadRateLimitPerUser()`

UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
### GetDefaultSortOrder

`func (o *CreateGuildRequestChannelItem) GetDefaultSortOrder() int32`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *CreateGuildRequestChannelItem) GetDefaultSortOrderOk() (*int32, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *CreateGuildRequestChannelItem) SetDefaultSortOrder(v int32)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *CreateGuildRequestChannelItem) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *CreateGuildRequestChannelItem) GetDefaultForumLayout() int32`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *CreateGuildRequestChannelItem) GetDefaultForumLayoutOk() (*int32, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *CreateGuildRequestChannelItem) SetDefaultForumLayout(v int32)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *CreateGuildRequestChannelItem) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *CreateGuildRequestChannelItem) GetDefaultTagSetting() string`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *CreateGuildRequestChannelItem) GetDefaultTagSettingOk() (*string, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *CreateGuildRequestChannelItem) SetDefaultTagSetting(v string)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *CreateGuildRequestChannelItem) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetId

`func (o *CreateGuildRequestChannelItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGuildRequestChannelItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGuildRequestChannelItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateGuildRequestChannelItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAvailableTags

`func (o *CreateGuildRequestChannelItem) GetAvailableTags() []CreateOrUpdateThreadTagRequest`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *CreateGuildRequestChannelItem) GetAvailableTagsOk() (*[]CreateOrUpdateThreadTagRequest, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *CreateGuildRequestChannelItem) SetAvailableTags(v []CreateOrUpdateThreadTagRequest)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *CreateGuildRequestChannelItem) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### SetAvailableTagsNil

`func (o *CreateGuildRequestChannelItem) SetAvailableTagsNil(b bool)`

 SetAvailableTagsNil sets the value for AvailableTags to be an explicit nil

### UnsetAvailableTags
`func (o *CreateGuildRequestChannelItem) UnsetAvailableTags()`

UnsetAvailableTags ensures that no value is present for AvailableTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


