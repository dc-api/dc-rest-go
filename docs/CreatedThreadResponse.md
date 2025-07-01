# CreatedThreadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **NullableTime** |  | [optional] 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**UserLimit** | Pointer to **NullableInt32** |  | [optional] 
**RtcRegion** | Pointer to **NullableString** |  | [optional] 
**VideoQualityMode** | Pointer to **int32** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | **string** |  | 
**ThreadMetadata** | Pointer to [**NullableThreadMetadataResponse**](ThreadMetadataResponse.md) |  | [optional] 
**MessageCount** | **int32** |  | 
**MemberCount** | **int32** |  | 
**TotalMessageSent** | **int32** |  | 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Member** | Pointer to [**NullableThreadMemberResponse**](ThreadMemberResponse.md) |  | [optional] 

## Methods

### NewCreatedThreadResponse

`func NewCreatedThreadResponse(id string, type_ int32, flags int32, guildId string, name string, ownerId string, messageCount int32, memberCount int32, totalMessageSent int32, ) *CreatedThreadResponse`

NewCreatedThreadResponse instantiates a new CreatedThreadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedThreadResponseWithDefaults

`func NewCreatedThreadResponseWithDefaults() *CreatedThreadResponse`

NewCreatedThreadResponseWithDefaults instantiates a new CreatedThreadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedThreadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedThreadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedThreadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CreatedThreadResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatedThreadResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatedThreadResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *CreatedThreadResponse) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *CreatedThreadResponse) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *CreatedThreadResponse) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *CreatedThreadResponse) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *CreatedThreadResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CreatedThreadResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CreatedThreadResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *CreatedThreadResponse) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *CreatedThreadResponse) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *CreatedThreadResponse) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *CreatedThreadResponse) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### SetLastPinTimestampNil

`func (o *CreatedThreadResponse) SetLastPinTimestampNil(b bool)`

 SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil

### UnsetLastPinTimestamp
`func (o *CreatedThreadResponse) UnsetLastPinTimestamp()`

UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
### GetGuildId

`func (o *CreatedThreadResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *CreatedThreadResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *CreatedThreadResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *CreatedThreadResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedThreadResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedThreadResponse) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CreatedThreadResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreatedThreadResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreatedThreadResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreatedThreadResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *CreatedThreadResponse) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreatedThreadResponse) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreatedThreadResponse) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreatedThreadResponse) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *CreatedThreadResponse) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *CreatedThreadResponse) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetBitrate

`func (o *CreatedThreadResponse) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *CreatedThreadResponse) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *CreatedThreadResponse) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *CreatedThreadResponse) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *CreatedThreadResponse) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *CreatedThreadResponse) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *CreatedThreadResponse) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *CreatedThreadResponse) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *CreatedThreadResponse) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *CreatedThreadResponse) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *CreatedThreadResponse) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *CreatedThreadResponse) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetRtcRegion

`func (o *CreatedThreadResponse) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *CreatedThreadResponse) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *CreatedThreadResponse) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *CreatedThreadResponse) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *CreatedThreadResponse) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *CreatedThreadResponse) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *CreatedThreadResponse) GetVideoQualityMode() int32`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *CreatedThreadResponse) GetVideoQualityModeOk() (*int32, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *CreatedThreadResponse) SetVideoQualityMode(v int32)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *CreatedThreadResponse) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetPermissions

`func (o *CreatedThreadResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreatedThreadResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreatedThreadResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreatedThreadResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreatedThreadResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreatedThreadResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetOwnerId

`func (o *CreatedThreadResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreatedThreadResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreatedThreadResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetThreadMetadata

`func (o *CreatedThreadResponse) GetThreadMetadata() ThreadMetadataResponse`

GetThreadMetadata returns the ThreadMetadata field if non-nil, zero value otherwise.

### GetThreadMetadataOk

`func (o *CreatedThreadResponse) GetThreadMetadataOk() (*ThreadMetadataResponse, bool)`

GetThreadMetadataOk returns a tuple with the ThreadMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadMetadata

`func (o *CreatedThreadResponse) SetThreadMetadata(v ThreadMetadataResponse)`

SetThreadMetadata sets ThreadMetadata field to given value.

### HasThreadMetadata

`func (o *CreatedThreadResponse) HasThreadMetadata() bool`

HasThreadMetadata returns a boolean if a field has been set.

### SetThreadMetadataNil

`func (o *CreatedThreadResponse) SetThreadMetadataNil(b bool)`

 SetThreadMetadataNil sets the value for ThreadMetadata to be an explicit nil

### UnsetThreadMetadata
`func (o *CreatedThreadResponse) UnsetThreadMetadata()`

UnsetThreadMetadata ensures that no value is present for ThreadMetadata, not even an explicit nil
### GetMessageCount

`func (o *CreatedThreadResponse) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *CreatedThreadResponse) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *CreatedThreadResponse) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetMemberCount

`func (o *CreatedThreadResponse) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *CreatedThreadResponse) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *CreatedThreadResponse) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetTotalMessageSent

`func (o *CreatedThreadResponse) GetTotalMessageSent() int32`

GetTotalMessageSent returns the TotalMessageSent field if non-nil, zero value otherwise.

### GetTotalMessageSentOk

`func (o *CreatedThreadResponse) GetTotalMessageSentOk() (*int32, bool)`

GetTotalMessageSentOk returns a tuple with the TotalMessageSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageSent

`func (o *CreatedThreadResponse) SetTotalMessageSent(v int32)`

SetTotalMessageSent sets TotalMessageSent field to given value.


### GetAppliedTags

`func (o *CreatedThreadResponse) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *CreatedThreadResponse) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *CreatedThreadResponse) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *CreatedThreadResponse) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *CreatedThreadResponse) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *CreatedThreadResponse) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil
### GetMember

`func (o *CreatedThreadResponse) GetMember() ThreadMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *CreatedThreadResponse) GetMemberOk() (*ThreadMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *CreatedThreadResponse) SetMember(v ThreadMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *CreatedThreadResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *CreatedThreadResponse) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *CreatedThreadResponse) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


