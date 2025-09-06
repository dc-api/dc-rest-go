# GuildInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Code** | **string** |  | 
**Inviter** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**IsContact** | Pointer to **bool** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**Guild** | [**InviteGuildResponse**](InviteGuildResponse.md) |  | 
**GuildId** | **string** |  | 
**Channel** | [**InviteChannelResponse**](InviteChannelResponse.md) |  | 
**TargetType** | Pointer to **NullableInt32** |  | [optional] 
**TargetUser** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**TargetApplication** | Pointer to [**InviteApplicationResponse**](InviteApplicationResponse.md) |  | [optional] 
**GuildScheduledEvent** | Pointer to [**ScheduledEventResponse**](ScheduledEventResponse.md) |  | [optional] 
**Uses** | Pointer to **int32** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**ApproximateMemberCount** | Pointer to **NullableInt32** |  | [optional] 
**ApproximatePresenceCount** | Pointer to **NullableInt32** |  | [optional] 
**IsNicknameChangeable** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuildInviteResponse

`func NewGuildInviteResponse(type_ NullableInt32, code string, guild InviteGuildResponse, guildId string, channel InviteChannelResponse, ) *GuildInviteResponse`

NewGuildInviteResponse instantiates a new GuildInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildInviteResponseWithDefaults

`func NewGuildInviteResponseWithDefaults() *GuildInviteResponse`

NewGuildInviteResponseWithDefaults instantiates a new GuildInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GuildInviteResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildInviteResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildInviteResponse) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GuildInviteResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GuildInviteResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *GuildInviteResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GuildInviteResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GuildInviteResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetInviter

`func (o *GuildInviteResponse) GetInviter() UserResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *GuildInviteResponse) GetInviterOk() (*UserResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *GuildInviteResponse) SetInviter(v UserResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *GuildInviteResponse) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetMaxAge

`func (o *GuildInviteResponse) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *GuildInviteResponse) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *GuildInviteResponse) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *GuildInviteResponse) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GuildInviteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GuildInviteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GuildInviteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GuildInviteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GuildInviteResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GuildInviteResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GuildInviteResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GuildInviteResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GuildInviteResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GuildInviteResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsContact

`func (o *GuildInviteResponse) GetIsContact() bool`

GetIsContact returns the IsContact field if non-nil, zero value otherwise.

### GetIsContactOk

`func (o *GuildInviteResponse) GetIsContactOk() (*bool, bool)`

GetIsContactOk returns a tuple with the IsContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContact

`func (o *GuildInviteResponse) SetIsContact(v bool)`

SetIsContact sets IsContact field to given value.

### HasIsContact

`func (o *GuildInviteResponse) HasIsContact() bool`

HasIsContact returns a boolean if a field has been set.

### GetFlags

`func (o *GuildInviteResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GuildInviteResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GuildInviteResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GuildInviteResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGuild

`func (o *GuildInviteResponse) GetGuild() InviteGuildResponse`

GetGuild returns the Guild field if non-nil, zero value otherwise.

### GetGuildOk

`func (o *GuildInviteResponse) GetGuildOk() (*InviteGuildResponse, bool)`

GetGuildOk returns a tuple with the Guild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuild

`func (o *GuildInviteResponse) SetGuild(v InviteGuildResponse)`

SetGuild sets Guild field to given value.


### GetGuildId

`func (o *GuildInviteResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GuildInviteResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GuildInviteResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetChannel

`func (o *GuildInviteResponse) GetChannel() InviteChannelResponse`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GuildInviteResponse) GetChannelOk() (*InviteChannelResponse, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GuildInviteResponse) SetChannel(v InviteChannelResponse)`

SetChannel sets Channel field to given value.


### GetTargetType

`func (o *GuildInviteResponse) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GuildInviteResponse) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GuildInviteResponse) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *GuildInviteResponse) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *GuildInviteResponse) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *GuildInviteResponse) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargetUser

`func (o *GuildInviteResponse) GetTargetUser() UserResponse`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *GuildInviteResponse) GetTargetUserOk() (*UserResponse, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *GuildInviteResponse) SetTargetUser(v UserResponse)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *GuildInviteResponse) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetApplication

`func (o *GuildInviteResponse) GetTargetApplication() InviteApplicationResponse`

GetTargetApplication returns the TargetApplication field if non-nil, zero value otherwise.

### GetTargetApplicationOk

`func (o *GuildInviteResponse) GetTargetApplicationOk() (*InviteApplicationResponse, bool)`

GetTargetApplicationOk returns a tuple with the TargetApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplication

`func (o *GuildInviteResponse) SetTargetApplication(v InviteApplicationResponse)`

SetTargetApplication sets TargetApplication field to given value.

### HasTargetApplication

`func (o *GuildInviteResponse) HasTargetApplication() bool`

HasTargetApplication returns a boolean if a field has been set.

### GetGuildScheduledEvent

`func (o *GuildInviteResponse) GetGuildScheduledEvent() ScheduledEventResponse`

GetGuildScheduledEvent returns the GuildScheduledEvent field if non-nil, zero value otherwise.

### GetGuildScheduledEventOk

`func (o *GuildInviteResponse) GetGuildScheduledEventOk() (*ScheduledEventResponse, bool)`

GetGuildScheduledEventOk returns a tuple with the GuildScheduledEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEvent

`func (o *GuildInviteResponse) SetGuildScheduledEvent(v ScheduledEventResponse)`

SetGuildScheduledEvent sets GuildScheduledEvent field to given value.

### HasGuildScheduledEvent

`func (o *GuildInviteResponse) HasGuildScheduledEvent() bool`

HasGuildScheduledEvent returns a boolean if a field has been set.

### GetUses

`func (o *GuildInviteResponse) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *GuildInviteResponse) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *GuildInviteResponse) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *GuildInviteResponse) HasUses() bool`

HasUses returns a boolean if a field has been set.

### GetMaxUses

`func (o *GuildInviteResponse) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *GuildInviteResponse) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *GuildInviteResponse) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *GuildInviteResponse) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetTemporary

`func (o *GuildInviteResponse) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *GuildInviteResponse) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *GuildInviteResponse) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *GuildInviteResponse) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetApproximateMemberCount

`func (o *GuildInviteResponse) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *GuildInviteResponse) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *GuildInviteResponse) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *GuildInviteResponse) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.

### SetApproximateMemberCountNil

`func (o *GuildInviteResponse) SetApproximateMemberCountNil(b bool)`

 SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil

### UnsetApproximateMemberCount
`func (o *GuildInviteResponse) UnsetApproximateMemberCount()`

UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
### GetApproximatePresenceCount

`func (o *GuildInviteResponse) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *GuildInviteResponse) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *GuildInviteResponse) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.

### HasApproximatePresenceCount

`func (o *GuildInviteResponse) HasApproximatePresenceCount() bool`

HasApproximatePresenceCount returns a boolean if a field has been set.

### SetApproximatePresenceCountNil

`func (o *GuildInviteResponse) SetApproximatePresenceCountNil(b bool)`

 SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil

### UnsetApproximatePresenceCount
`func (o *GuildInviteResponse) UnsetApproximatePresenceCount()`

UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil
### GetIsNicknameChangeable

`func (o *GuildInviteResponse) GetIsNicknameChangeable() bool`

GetIsNicknameChangeable returns the IsNicknameChangeable field if non-nil, zero value otherwise.

### GetIsNicknameChangeableOk

`func (o *GuildInviteResponse) GetIsNicknameChangeableOk() (*bool, bool)`

GetIsNicknameChangeableOk returns a tuple with the IsNicknameChangeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNicknameChangeable

`func (o *GuildInviteResponse) SetIsNicknameChangeable(v bool)`

SetIsNicknameChangeable sets IsNicknameChangeable field to given value.

### HasIsNicknameChangeable

`func (o *GuildInviteResponse) HasIsNicknameChangeable() bool`

HasIsNicknameChangeable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


