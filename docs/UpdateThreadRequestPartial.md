# UpdateThreadRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Archived** | Pointer to **NullableBool** |  | [optional] 
**Locked** | Pointer to **NullableBool** |  | [optional] 
**Invitable** | Pointer to **NullableBool** |  | [optional] 
**AutoArchiveDuration** | Pointer to **int32** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**UserLimit** | Pointer to **NullableInt32** |  | [optional] 
**RtcRegion** | Pointer to **NullableString** |  | [optional] 
**VideoQualityMode** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateThreadRequestPartial

`func NewUpdateThreadRequestPartial() *UpdateThreadRequestPartial`

NewUpdateThreadRequestPartial instantiates a new UpdateThreadRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThreadRequestPartialWithDefaults

`func NewUpdateThreadRequestPartialWithDefaults() *UpdateThreadRequestPartial`

NewUpdateThreadRequestPartialWithDefaults instantiates a new UpdateThreadRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateThreadRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateThreadRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateThreadRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateThreadRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateThreadRequestPartial) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateThreadRequestPartial) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetArchived

`func (o *UpdateThreadRequestPartial) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateThreadRequestPartial) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateThreadRequestPartial) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateThreadRequestPartial) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *UpdateThreadRequestPartial) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *UpdateThreadRequestPartial) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetLocked

`func (o *UpdateThreadRequestPartial) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateThreadRequestPartial) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateThreadRequestPartial) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateThreadRequestPartial) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *UpdateThreadRequestPartial) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *UpdateThreadRequestPartial) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetInvitable

`func (o *UpdateThreadRequestPartial) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *UpdateThreadRequestPartial) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *UpdateThreadRequestPartial) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *UpdateThreadRequestPartial) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.

### SetInvitableNil

`func (o *UpdateThreadRequestPartial) SetInvitableNil(b bool)`

 SetInvitableNil sets the value for Invitable to be an explicit nil

### UnsetInvitable
`func (o *UpdateThreadRequestPartial) UnsetInvitable()`

UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil
### GetAutoArchiveDuration

`func (o *UpdateThreadRequestPartial) GetAutoArchiveDuration() int32`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *UpdateThreadRequestPartial) GetAutoArchiveDurationOk() (*int32, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *UpdateThreadRequestPartial) SetAutoArchiveDuration(v int32)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *UpdateThreadRequestPartial) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *UpdateThreadRequestPartial) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *UpdateThreadRequestPartial) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *UpdateThreadRequestPartial) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *UpdateThreadRequestPartial) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *UpdateThreadRequestPartial) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *UpdateThreadRequestPartial) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetFlags

`func (o *UpdateThreadRequestPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateThreadRequestPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateThreadRequestPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateThreadRequestPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *UpdateThreadRequestPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *UpdateThreadRequestPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAppliedTags

`func (o *UpdateThreadRequestPartial) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *UpdateThreadRequestPartial) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *UpdateThreadRequestPartial) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *UpdateThreadRequestPartial) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *UpdateThreadRequestPartial) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *UpdateThreadRequestPartial) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil
### GetBitrate

`func (o *UpdateThreadRequestPartial) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *UpdateThreadRequestPartial) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *UpdateThreadRequestPartial) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *UpdateThreadRequestPartial) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *UpdateThreadRequestPartial) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *UpdateThreadRequestPartial) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *UpdateThreadRequestPartial) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UpdateThreadRequestPartial) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UpdateThreadRequestPartial) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UpdateThreadRequestPartial) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *UpdateThreadRequestPartial) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *UpdateThreadRequestPartial) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetRtcRegion

`func (o *UpdateThreadRequestPartial) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *UpdateThreadRequestPartial) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *UpdateThreadRequestPartial) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *UpdateThreadRequestPartial) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *UpdateThreadRequestPartial) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *UpdateThreadRequestPartial) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *UpdateThreadRequestPartial) GetVideoQualityMode() int32`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *UpdateThreadRequestPartial) GetVideoQualityModeOk() (*int32, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *UpdateThreadRequestPartial) SetVideoQualityMode(v int32)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *UpdateThreadRequestPartial) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


