# CreateChannelInvite200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Code** | **string** |  | 
**Inviter** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**FriendsCount** | Pointer to **int32** |  | [optional] 
**Channel** | [**InviteChannelResponse**](InviteChannelResponse.md) |  | 
**IsContact** | Pointer to **bool** |  | [optional] 
**Uses** | Pointer to **int32** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**ApproximateMemberCount** | Pointer to **int32** |  | [optional] 
**Guild** | [**InviteGuildResponse**](InviteGuildResponse.md) |  | 
**GuildId** | **string** |  | 
**TargetType** | Pointer to **int32** |  | [optional] 
**TargetUser** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**TargetApplication** | Pointer to [**InviteApplicationResponse**](InviteApplicationResponse.md) |  | [optional] 
**GuildScheduledEvent** | Pointer to [**ScheduledEventResponse**](ScheduledEventResponse.md) |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**ApproximatePresenceCount** | Pointer to **int32** |  | [optional] 
**IsNicknameChangeable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateChannelInvite200Response

`func NewCreateChannelInvite200Response(type_ NullableInt32, code string, channel InviteChannelResponse, guild InviteGuildResponse, guildId string, ) *CreateChannelInvite200Response`

NewCreateChannelInvite200Response instantiates a new CreateChannelInvite200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChannelInvite200ResponseWithDefaults

`func NewCreateChannelInvite200ResponseWithDefaults() *CreateChannelInvite200Response`

NewCreateChannelInvite200ResponseWithDefaults instantiates a new CreateChannelInvite200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateChannelInvite200Response) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateChannelInvite200Response) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateChannelInvite200Response) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateChannelInvite200Response) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateChannelInvite200Response) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *CreateChannelInvite200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateChannelInvite200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateChannelInvite200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetInviter

`func (o *CreateChannelInvite200Response) GetInviter() UserResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *CreateChannelInvite200Response) GetInviterOk() (*UserResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *CreateChannelInvite200Response) SetInviter(v UserResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *CreateChannelInvite200Response) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetMaxAge

`func (o *CreateChannelInvite200Response) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CreateChannelInvite200Response) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CreateChannelInvite200Response) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CreateChannelInvite200Response) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateChannelInvite200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateChannelInvite200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateChannelInvite200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateChannelInvite200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateChannelInvite200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateChannelInvite200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateChannelInvite200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateChannelInvite200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFriendsCount

`func (o *CreateChannelInvite200Response) GetFriendsCount() int32`

GetFriendsCount returns the FriendsCount field if non-nil, zero value otherwise.

### GetFriendsCountOk

`func (o *CreateChannelInvite200Response) GetFriendsCountOk() (*int32, bool)`

GetFriendsCountOk returns a tuple with the FriendsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsCount

`func (o *CreateChannelInvite200Response) SetFriendsCount(v int32)`

SetFriendsCount sets FriendsCount field to given value.

### HasFriendsCount

`func (o *CreateChannelInvite200Response) HasFriendsCount() bool`

HasFriendsCount returns a boolean if a field has been set.

### GetChannel

`func (o *CreateChannelInvite200Response) GetChannel() InviteChannelResponse`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateChannelInvite200Response) GetChannelOk() (*InviteChannelResponse, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateChannelInvite200Response) SetChannel(v InviteChannelResponse)`

SetChannel sets Channel field to given value.


### GetIsContact

`func (o *CreateChannelInvite200Response) GetIsContact() bool`

GetIsContact returns the IsContact field if non-nil, zero value otherwise.

### GetIsContactOk

`func (o *CreateChannelInvite200Response) GetIsContactOk() (*bool, bool)`

GetIsContactOk returns a tuple with the IsContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContact

`func (o *CreateChannelInvite200Response) SetIsContact(v bool)`

SetIsContact sets IsContact field to given value.

### HasIsContact

`func (o *CreateChannelInvite200Response) HasIsContact() bool`

HasIsContact returns a boolean if a field has been set.

### GetUses

`func (o *CreateChannelInvite200Response) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *CreateChannelInvite200Response) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *CreateChannelInvite200Response) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *CreateChannelInvite200Response) HasUses() bool`

HasUses returns a boolean if a field has been set.

### GetMaxUses

`func (o *CreateChannelInvite200Response) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *CreateChannelInvite200Response) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *CreateChannelInvite200Response) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *CreateChannelInvite200Response) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetFlags

`func (o *CreateChannelInvite200Response) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CreateChannelInvite200Response) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CreateChannelInvite200Response) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *CreateChannelInvite200Response) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetApproximateMemberCount

`func (o *CreateChannelInvite200Response) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *CreateChannelInvite200Response) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *CreateChannelInvite200Response) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *CreateChannelInvite200Response) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.

### GetGuild

`func (o *CreateChannelInvite200Response) GetGuild() InviteGuildResponse`

GetGuild returns the Guild field if non-nil, zero value otherwise.

### GetGuildOk

`func (o *CreateChannelInvite200Response) GetGuildOk() (*InviteGuildResponse, bool)`

GetGuildOk returns a tuple with the Guild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuild

`func (o *CreateChannelInvite200Response) SetGuild(v InviteGuildResponse)`

SetGuild sets Guild field to given value.


### GetGuildId

`func (o *CreateChannelInvite200Response) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *CreateChannelInvite200Response) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *CreateChannelInvite200Response) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetTargetType

`func (o *CreateChannelInvite200Response) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CreateChannelInvite200Response) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CreateChannelInvite200Response) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CreateChannelInvite200Response) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetUser

`func (o *CreateChannelInvite200Response) GetTargetUser() UserResponse`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *CreateChannelInvite200Response) GetTargetUserOk() (*UserResponse, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *CreateChannelInvite200Response) SetTargetUser(v UserResponse)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *CreateChannelInvite200Response) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetApplication

`func (o *CreateChannelInvite200Response) GetTargetApplication() InviteApplicationResponse`

GetTargetApplication returns the TargetApplication field if non-nil, zero value otherwise.

### GetTargetApplicationOk

`func (o *CreateChannelInvite200Response) GetTargetApplicationOk() (*InviteApplicationResponse, bool)`

GetTargetApplicationOk returns a tuple with the TargetApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplication

`func (o *CreateChannelInvite200Response) SetTargetApplication(v InviteApplicationResponse)`

SetTargetApplication sets TargetApplication field to given value.

### HasTargetApplication

`func (o *CreateChannelInvite200Response) HasTargetApplication() bool`

HasTargetApplication returns a boolean if a field has been set.

### GetGuildScheduledEvent

`func (o *CreateChannelInvite200Response) GetGuildScheduledEvent() ScheduledEventResponse`

GetGuildScheduledEvent returns the GuildScheduledEvent field if non-nil, zero value otherwise.

### GetGuildScheduledEventOk

`func (o *CreateChannelInvite200Response) GetGuildScheduledEventOk() (*ScheduledEventResponse, bool)`

GetGuildScheduledEventOk returns a tuple with the GuildScheduledEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEvent

`func (o *CreateChannelInvite200Response) SetGuildScheduledEvent(v ScheduledEventResponse)`

SetGuildScheduledEvent sets GuildScheduledEvent field to given value.

### HasGuildScheduledEvent

`func (o *CreateChannelInvite200Response) HasGuildScheduledEvent() bool`

HasGuildScheduledEvent returns a boolean if a field has been set.

### GetTemporary

`func (o *CreateChannelInvite200Response) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CreateChannelInvite200Response) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CreateChannelInvite200Response) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CreateChannelInvite200Response) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetApproximatePresenceCount

`func (o *CreateChannelInvite200Response) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *CreateChannelInvite200Response) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *CreateChannelInvite200Response) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.

### HasApproximatePresenceCount

`func (o *CreateChannelInvite200Response) HasApproximatePresenceCount() bool`

HasApproximatePresenceCount returns a boolean if a field has been set.

### GetIsNicknameChangeable

`func (o *CreateChannelInvite200Response) GetIsNicknameChangeable() bool`

GetIsNicknameChangeable returns the IsNicknameChangeable field if non-nil, zero value otherwise.

### GetIsNicknameChangeableOk

`func (o *CreateChannelInvite200Response) GetIsNicknameChangeableOk() (*bool, bool)`

GetIsNicknameChangeableOk returns a tuple with the IsNicknameChangeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNicknameChangeable

`func (o *CreateChannelInvite200Response) SetIsNicknameChangeable(v bool)`

SetIsNicknameChangeable sets IsNicknameChangeable field to given value.

### HasIsNicknameChangeable

`func (o *CreateChannelInvite200Response) HasIsNicknameChangeable() bool`

HasIsNicknameChangeable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


