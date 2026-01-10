# UpdateGuildChannelRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CreateGuildChannelRequestType**](CreateGuildChannelRequestType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Bitrate** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 
**RateLimitPerUser** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteRequest**](ChannelPermissionOverwriteRequest.md) |  | [optional] 
**RtcRegion** | Pointer to **string** |  | [optional] 
**VideoQualityMode** | Pointer to [**CreateGuildChannelRequestVideoQualityMode**](CreateGuildChannelRequestVideoQualityMode.md) |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to [**CreateForumThreadRequestAutoArchiveDuration**](CreateForumThreadRequestAutoArchiveDuration.md) |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**CreateGuildChannelRequestDefaultReactionEmoji**](CreateGuildChannelRequestDefaultReactionEmoji.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **int32** |  | [optional] 
**DefaultSortOrder** | Pointer to [**CreateGuildChannelRequestDefaultSortOrder**](CreateGuildChannelRequestDefaultSortOrder.md) |  | [optional] 
**DefaultForumLayout** | Pointer to [**CreateGuildChannelRequestDefaultForumLayout**](CreateGuildChannelRequestDefaultForumLayout.md) |  | [optional] 
**DefaultTagSetting** | Pointer to [**CreateGuildChannelRequestDefaultTagSetting**](CreateGuildChannelRequestDefaultTagSetting.md) |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
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

`func (o *UpdateGuildChannelRequestPartial) GetType() CreateGuildChannelRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateGuildChannelRequestPartial) GetTypeOk() (*CreateGuildChannelRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateGuildChannelRequestPartial) SetType(v CreateGuildChannelRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateGuildChannelRequestPartial) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetParentId

`func (o *UpdateGuildChannelRequestPartial) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateGuildChannelRequestPartial) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateGuildChannelRequestPartial) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

### GetVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) GetVideoQualityMode() CreateGuildChannelRequestVideoQualityMode`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *UpdateGuildChannelRequestPartial) GetVideoQualityModeOk() (*CreateGuildChannelRequestVideoQualityMode, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) SetVideoQualityMode(v CreateGuildChannelRequestVideoQualityMode)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *UpdateGuildChannelRequestPartial) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *UpdateGuildChannelRequestPartial) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmoji() CreateGuildChannelRequestDefaultReactionEmoji`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmojiOk() (*CreateGuildChannelRequestDefaultReactionEmoji, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) SetDefaultReactionEmoji(v CreateGuildChannelRequestDefaultReactionEmoji)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *UpdateGuildChannelRequestPartial) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

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

### GetDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *UpdateGuildChannelRequestPartial) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *UpdateGuildChannelRequestPartial) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *UpdateGuildChannelRequestPartial) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


