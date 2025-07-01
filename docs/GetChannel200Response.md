# GetChannel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **time.Time** |  | [optional] 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**RateLimitPerUser** | Pointer to **int32** |  | [optional] 
**Bitrate** | Pointer to **int32** |  | [optional] 
**UserLimit** | Pointer to **int32** |  | [optional] 
**RtcRegion** | Pointer to **string** |  | [optional] 
**VideoQualityMode** | Pointer to **int32** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to **int32** |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **int32** |  | [optional] 
**Position** | **int32** |  | 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteResponse**](ChannelPermissionOverwriteResponse.md) |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 
**AvailableTags** | Pointer to [**[]ForumTagResponse**](ForumTagResponse.md) |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**DefaultReactionEmojiResponse**](DefaultReactionEmojiResponse.md) |  | [optional] 
**DefaultSortOrder** | Pointer to **int32** |  | [optional] 
**DefaultForumLayout** | Pointer to **int32** |  | [optional] 
**DefaultTagSetting** | Pointer to **string** |  | [optional] 
**HdStreamingUntil** | Pointer to **time.Time** |  | [optional] 
**HdStreamingBuyerId** | Pointer to **string** |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 
**Icon** | Pointer to **string** |  | [optional] 
**OwnerId** | **string** |  | 
**Managed** | Pointer to **bool** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ThreadMetadata** | Pointer to [**ThreadMetadataResponse**](ThreadMetadataResponse.md) |  | [optional] 
**MessageCount** | **int32** |  | 
**MemberCount** | **int32** |  | 
**TotalMessageSent** | **int32** |  | 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Member** | Pointer to [**ThreadMemberResponse**](ThreadMemberResponse.md) |  | [optional] 

## Methods

### NewGetChannel200Response

`func NewGetChannel200Response(id string, type_ int32, flags int32, guildId string, name string, position int32, recipients []UserResponse, ownerId string, messageCount int32, memberCount int32, totalMessageSent int32, ) *GetChannel200Response`

NewGetChannel200Response instantiates a new GetChannel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannel200ResponseWithDefaults

`func NewGetChannel200ResponseWithDefaults() *GetChannel200Response`

NewGetChannel200ResponseWithDefaults instantiates a new GetChannel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetChannel200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetChannel200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetChannel200Response) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GetChannel200Response) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetChannel200Response) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetChannel200Response) SetType(v int32)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *GetChannel200Response) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *GetChannel200Response) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *GetChannel200Response) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *GetChannel200Response) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *GetChannel200Response) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetChannel200Response) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetChannel200Response) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *GetChannel200Response) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *GetChannel200Response) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *GetChannel200Response) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *GetChannel200Response) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### GetGuildId

`func (o *GetChannel200Response) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GetChannel200Response) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GetChannel200Response) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *GetChannel200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetChannel200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetChannel200Response) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *GetChannel200Response) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetChannel200Response) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetChannel200Response) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetChannel200Response) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *GetChannel200Response) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *GetChannel200Response) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *GetChannel200Response) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *GetChannel200Response) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### GetBitrate

`func (o *GetChannel200Response) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *GetChannel200Response) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *GetChannel200Response) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *GetChannel200Response) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetUserLimit

`func (o *GetChannel200Response) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *GetChannel200Response) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *GetChannel200Response) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *GetChannel200Response) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetRtcRegion

`func (o *GetChannel200Response) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *GetChannel200Response) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *GetChannel200Response) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *GetChannel200Response) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### GetVideoQualityMode

`func (o *GetChannel200Response) GetVideoQualityMode() int32`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *GetChannel200Response) GetVideoQualityModeOk() (*int32, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *GetChannel200Response) SetVideoQualityMode(v int32)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *GetChannel200Response) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetPermissions

`func (o *GetChannel200Response) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetChannel200Response) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetChannel200Response) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetChannel200Response) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTopic

`func (o *GetChannel200Response) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetChannel200Response) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetChannel200Response) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetChannel200Response) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *GetChannel200Response) GetDefaultAutoArchiveDuration() int32`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *GetChannel200Response) GetDefaultAutoArchiveDurationOk() (*int32, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *GetChannel200Response) SetDefaultAutoArchiveDuration(v int32)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *GetChannel200Response) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *GetChannel200Response) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *GetChannel200Response) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *GetChannel200Response) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *GetChannel200Response) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### GetPosition

`func (o *GetChannel200Response) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetChannel200Response) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetChannel200Response) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetPermissionOverwrites

`func (o *GetChannel200Response) GetPermissionOverwrites() []ChannelPermissionOverwriteResponse`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *GetChannel200Response) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteResponse, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *GetChannel200Response) SetPermissionOverwrites(v []ChannelPermissionOverwriteResponse)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *GetChannel200Response) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### GetNsfw

`func (o *GetChannel200Response) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GetChannel200Response) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GetChannel200Response) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *GetChannel200Response) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetAvailableTags

`func (o *GetChannel200Response) GetAvailableTags() []ForumTagResponse`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *GetChannel200Response) GetAvailableTagsOk() (*[]ForumTagResponse, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *GetChannel200Response) SetAvailableTags(v []ForumTagResponse)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *GetChannel200Response) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *GetChannel200Response) GetDefaultReactionEmoji() DefaultReactionEmojiResponse`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *GetChannel200Response) GetDefaultReactionEmojiOk() (*DefaultReactionEmojiResponse, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *GetChannel200Response) SetDefaultReactionEmoji(v DefaultReactionEmojiResponse)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *GetChannel200Response) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultSortOrder

`func (o *GetChannel200Response) GetDefaultSortOrder() int32`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *GetChannel200Response) GetDefaultSortOrderOk() (*int32, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *GetChannel200Response) SetDefaultSortOrder(v int32)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *GetChannel200Response) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *GetChannel200Response) GetDefaultForumLayout() int32`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *GetChannel200Response) GetDefaultForumLayoutOk() (*int32, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *GetChannel200Response) SetDefaultForumLayout(v int32)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *GetChannel200Response) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetDefaultTagSetting

`func (o *GetChannel200Response) GetDefaultTagSetting() string`

GetDefaultTagSetting returns the DefaultTagSetting field if non-nil, zero value otherwise.

### GetDefaultTagSettingOk

`func (o *GetChannel200Response) GetDefaultTagSettingOk() (*string, bool)`

GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTagSetting

`func (o *GetChannel200Response) SetDefaultTagSetting(v string)`

SetDefaultTagSetting sets DefaultTagSetting field to given value.

### HasDefaultTagSetting

`func (o *GetChannel200Response) HasDefaultTagSetting() bool`

HasDefaultTagSetting returns a boolean if a field has been set.

### GetHdStreamingUntil

`func (o *GetChannel200Response) GetHdStreamingUntil() time.Time`

GetHdStreamingUntil returns the HdStreamingUntil field if non-nil, zero value otherwise.

### GetHdStreamingUntilOk

`func (o *GetChannel200Response) GetHdStreamingUntilOk() (*time.Time, bool)`

GetHdStreamingUntilOk returns a tuple with the HdStreamingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdStreamingUntil

`func (o *GetChannel200Response) SetHdStreamingUntil(v time.Time)`

SetHdStreamingUntil sets HdStreamingUntil field to given value.

### HasHdStreamingUntil

`func (o *GetChannel200Response) HasHdStreamingUntil() bool`

HasHdStreamingUntil returns a boolean if a field has been set.

### GetHdStreamingBuyerId

`func (o *GetChannel200Response) GetHdStreamingBuyerId() string`

GetHdStreamingBuyerId returns the HdStreamingBuyerId field if non-nil, zero value otherwise.

### GetHdStreamingBuyerIdOk

`func (o *GetChannel200Response) GetHdStreamingBuyerIdOk() (*string, bool)`

GetHdStreamingBuyerIdOk returns a tuple with the HdStreamingBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdStreamingBuyerId

`func (o *GetChannel200Response) SetHdStreamingBuyerId(v string)`

SetHdStreamingBuyerId sets HdStreamingBuyerId field to given value.

### HasHdStreamingBuyerId

`func (o *GetChannel200Response) HasHdStreamingBuyerId() bool`

HasHdStreamingBuyerId returns a boolean if a field has been set.

### GetRecipients

`func (o *GetChannel200Response) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetChannel200Response) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetChannel200Response) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.


### GetIcon

`func (o *GetChannel200Response) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GetChannel200Response) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GetChannel200Response) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GetChannel200Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetOwnerId

`func (o *GetChannel200Response) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetChannel200Response) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetChannel200Response) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetManaged

`func (o *GetChannel200Response) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GetChannel200Response) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GetChannel200Response) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *GetChannel200Response) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetChannel200Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetChannel200Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetChannel200Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetChannel200Response) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetThreadMetadata

`func (o *GetChannel200Response) GetThreadMetadata() ThreadMetadataResponse`

GetThreadMetadata returns the ThreadMetadata field if non-nil, zero value otherwise.

### GetThreadMetadataOk

`func (o *GetChannel200Response) GetThreadMetadataOk() (*ThreadMetadataResponse, bool)`

GetThreadMetadataOk returns a tuple with the ThreadMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadMetadata

`func (o *GetChannel200Response) SetThreadMetadata(v ThreadMetadataResponse)`

SetThreadMetadata sets ThreadMetadata field to given value.

### HasThreadMetadata

`func (o *GetChannel200Response) HasThreadMetadata() bool`

HasThreadMetadata returns a boolean if a field has been set.

### GetMessageCount

`func (o *GetChannel200Response) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *GetChannel200Response) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *GetChannel200Response) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetMemberCount

`func (o *GetChannel200Response) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GetChannel200Response) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GetChannel200Response) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetTotalMessageSent

`func (o *GetChannel200Response) GetTotalMessageSent() int32`

GetTotalMessageSent returns the TotalMessageSent field if non-nil, zero value otherwise.

### GetTotalMessageSentOk

`func (o *GetChannel200Response) GetTotalMessageSentOk() (*int32, bool)`

GetTotalMessageSentOk returns a tuple with the TotalMessageSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageSent

`func (o *GetChannel200Response) SetTotalMessageSent(v int32)`

SetTotalMessageSent sets TotalMessageSent field to given value.


### GetAppliedTags

`func (o *GetChannel200Response) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *GetChannel200Response) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *GetChannel200Response) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *GetChannel200Response) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### GetMember

`func (o *GetChannel200Response) GetMember() ThreadMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetChannel200Response) GetMemberOk() (*ThreadMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetChannel200Response) SetMember(v ThreadMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetChannel200Response) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


