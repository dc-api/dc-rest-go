# BotAddGuildMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nick** | Pointer to **NullableString** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Mute** | Pointer to **NullableBool** |  | [optional] 
**Deaf** | Pointer to **NullableBool** |  | [optional] 
**AccessToken** | **string** |  | 
**Flags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBotAddGuildMemberRequest

`func NewBotAddGuildMemberRequest(accessToken string, ) *BotAddGuildMemberRequest`

NewBotAddGuildMemberRequest instantiates a new BotAddGuildMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotAddGuildMemberRequestWithDefaults

`func NewBotAddGuildMemberRequestWithDefaults() *BotAddGuildMemberRequest`

NewBotAddGuildMemberRequestWithDefaults instantiates a new BotAddGuildMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNick

`func (o *BotAddGuildMemberRequest) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *BotAddGuildMemberRequest) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *BotAddGuildMemberRequest) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *BotAddGuildMemberRequest) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *BotAddGuildMemberRequest) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *BotAddGuildMemberRequest) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetRoles

`func (o *BotAddGuildMemberRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BotAddGuildMemberRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BotAddGuildMemberRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BotAddGuildMemberRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *BotAddGuildMemberRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *BotAddGuildMemberRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetMute

`func (o *BotAddGuildMemberRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *BotAddGuildMemberRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *BotAddGuildMemberRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *BotAddGuildMemberRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### SetMuteNil

`func (o *BotAddGuildMemberRequest) SetMuteNil(b bool)`

 SetMuteNil sets the value for Mute to be an explicit nil

### UnsetMute
`func (o *BotAddGuildMemberRequest) UnsetMute()`

UnsetMute ensures that no value is present for Mute, not even an explicit nil
### GetDeaf

`func (o *BotAddGuildMemberRequest) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *BotAddGuildMemberRequest) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *BotAddGuildMemberRequest) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.

### HasDeaf

`func (o *BotAddGuildMemberRequest) HasDeaf() bool`

HasDeaf returns a boolean if a field has been set.

### SetDeafNil

`func (o *BotAddGuildMemberRequest) SetDeafNil(b bool)`

 SetDeafNil sets the value for Deaf to be an explicit nil

### UnsetDeaf
`func (o *BotAddGuildMemberRequest) UnsetDeaf()`

UnsetDeaf ensures that no value is present for Deaf, not even an explicit nil
### GetAccessToken

`func (o *BotAddGuildMemberRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BotAddGuildMemberRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BotAddGuildMemberRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetFlags

`func (o *BotAddGuildMemberRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BotAddGuildMemberRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BotAddGuildMemberRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BotAddGuildMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *BotAddGuildMemberRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *BotAddGuildMemberRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


