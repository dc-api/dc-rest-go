# AddGuildMemberRequest

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

### NewAddGuildMemberRequest

`func NewAddGuildMemberRequest(accessToken string, ) *AddGuildMemberRequest`

NewAddGuildMemberRequest instantiates a new AddGuildMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGuildMemberRequestWithDefaults

`func NewAddGuildMemberRequestWithDefaults() *AddGuildMemberRequest`

NewAddGuildMemberRequestWithDefaults instantiates a new AddGuildMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNick

`func (o *AddGuildMemberRequest) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *AddGuildMemberRequest) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *AddGuildMemberRequest) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *AddGuildMemberRequest) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *AddGuildMemberRequest) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *AddGuildMemberRequest) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetRoles

`func (o *AddGuildMemberRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddGuildMemberRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddGuildMemberRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddGuildMemberRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AddGuildMemberRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AddGuildMemberRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetMute

`func (o *AddGuildMemberRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *AddGuildMemberRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *AddGuildMemberRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *AddGuildMemberRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### SetMuteNil

`func (o *AddGuildMemberRequest) SetMuteNil(b bool)`

 SetMuteNil sets the value for Mute to be an explicit nil

### UnsetMute
`func (o *AddGuildMemberRequest) UnsetMute()`

UnsetMute ensures that no value is present for Mute, not even an explicit nil
### GetDeaf

`func (o *AddGuildMemberRequest) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *AddGuildMemberRequest) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *AddGuildMemberRequest) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.

### HasDeaf

`func (o *AddGuildMemberRequest) HasDeaf() bool`

HasDeaf returns a boolean if a field has been set.

### SetDeafNil

`func (o *AddGuildMemberRequest) SetDeafNil(b bool)`

 SetDeafNil sets the value for Deaf to be an explicit nil

### UnsetDeaf
`func (o *AddGuildMemberRequest) UnsetDeaf()`

UnsetDeaf ensures that no value is present for Deaf, not even an explicit nil
### GetAccessToken

`func (o *AddGuildMemberRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AddGuildMemberRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AddGuildMemberRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetFlags

`func (o *AddGuildMemberRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *AddGuildMemberRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *AddGuildMemberRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *AddGuildMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *AddGuildMemberRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *AddGuildMemberRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


