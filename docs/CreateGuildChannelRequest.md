# CreateGuildChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CreateGuildChannelRequestType**](CreateGuildChannelRequestType.md) |  | [optional] 
**Name** | **string** |  | 
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
**AvailableTags** | Pointer to [**[]CreateGuildChannelRequestAvailableTagsInner**](CreateGuildChannelRequestAvailableTagsInner.md) |  | [optional] 

## Methods

### NewCreateGuildChannelRequest

`func NewCreateGuildChannelRequest(name string, ) *CreateGuildChannelRequest`

NewCreateGuildChannelRequest instantiates a new CreateGuildChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildChannelRequestWithDefaults

`func NewCreateGuildChannelRequestWithDefaults() *CreateGuildChannelRequest`

NewCreateGuildChannelRequestWithDefaults instantiates a new CreateGuildChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateGuildChannelRequest) GetType() CreateGuildChannelRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGuildChannelRequest) GetTypeOk() (*CreateGuildChannelRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGuildChannelRequest) SetType(v CreateGuildChannelRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateGuildChannelRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CreateGuildChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *CreateGuildChannelRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateGuildChannelRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateGuildChannelRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateGuildChannelRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTopic

`func (o *CreateGuildChannelRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CreateGuildChannelRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CreateGuildChannelRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *CreateGuildChannelRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetBitrate

`func (o *CreateGuildChannelRequest) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *CreateGuildChannelRequest) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *CreateGuildChannelRequest) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *CreateGuildChannelRequest) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetUserLimit

`func (o *CreateGuildChannelRequest) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *CreateGuildChannelRequest) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *CreateGuildChannelRequest) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *CreateGuildChannelRequest) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetNsfw

`func (o *CreateGuildChannelRequest) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *CreateGuildChannelRequest) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *CreateGuildChannelRequest) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *CreateGuildChannelRequest) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *CreateGuildChannelRequest) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreateGuildChannelRequest) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreateGuildChannelRequest) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreateGuildChannelRequest) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### GetParentId

`func (o *CreateGuildChannelRequest) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateGuildChannelRequest) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateGuildChannelRequest) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateGuildChannelRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermissionOverwrites

`func (o *CreateGuildChannelRequest) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *CreateGuildChannelRequest) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteRequest, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *CreateGuildChannelRequest) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *CreateGuildChannelRequest) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### GetRtcRegion

`func (o *CreateGuildChannelRequest) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *CreateGuildChannelRequest) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *CreateGuildChannelRequest) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *CreateGuildChannelRequest) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### GetVideoQualityMode

`func (o *CreateGuildChannelRequest) GetVideoQualityMode() CreateGuildChannelRequestVideoQualityMode`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *CreateGuildChannelRequest) GetVideoQualityModeOk() (*CreateGuildChannelRequestVideoQualityMode, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *CreateGuildChannelRequest) SetVideoQualityMode(v CreateGuildChannelRequestVideoQualityMode)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *CreateGuildChannelRequest) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *CreateGuildChannelRequest) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *CreateGuildChannelRequest) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *CreateGuildChannelRequest) GetDefaultReactionEmoji() CreateGuildChannelRequestDefaultReactionEmoji`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *CreateGuildChannelRequest) GetDefaultReactionEmojiOk() (*CreateGuildChannelRequestDefaultReactionEmoji, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *CreateGuildChannelRequest) SetDefaultReactionEmoji(v CreateGuildChannelRequestDefaultReactionEmoji)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *CreateGuildChannelRequest) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *CreateGuildChannelRequest) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *CreateGuildChannelRequest) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### GetDefaultSortOrder

`func (o *CreateGuildChannelRequest) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *CreateGuildChannelRequest) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *CreateGuildChannelRequest) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *CreateGuildChannelRequest) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *CreateGuildChannelRequest) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *CreateGuildChannelRequest) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *CreateGuildChannelRequest) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *CreateGuildChannelRequest) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *CreateGuildChannelRequest) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *CreateGuildChannelRequest) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *CreateGuildChannelRequest) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *CreateGuildChannelRequest) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetAvailableTags

`func (o *CreateGuildChannelRequest) GetAvailableTags() []CreateGuildChannelRequestAvailableTagsInner`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *CreateGuildChannelRequest) GetAvailableTagsOk() (*[]CreateGuildChannelRequestAvailableTagsInner, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *CreateGuildChannelRequest) SetAvailableTags(v []CreateGuildChannelRequestAvailableTagsInner)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *CreateGuildChannelRequest) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


