# UpdateGuildMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nick** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Mute** | Pointer to **bool** |  | [optional] 
**Deaf** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**CommunicationDisabledUntil** | Pointer to **time.Time** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateGuildMemberRequest

`func NewUpdateGuildMemberRequest() *UpdateGuildMemberRequest`

NewUpdateGuildMemberRequest instantiates a new UpdateGuildMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildMemberRequestWithDefaults

`func NewUpdateGuildMemberRequestWithDefaults() *UpdateGuildMemberRequest`

NewUpdateGuildMemberRequestWithDefaults instantiates a new UpdateGuildMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNick

`func (o *UpdateGuildMemberRequest) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *UpdateGuildMemberRequest) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *UpdateGuildMemberRequest) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *UpdateGuildMemberRequest) HasNick() bool`

HasNick returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateGuildMemberRequest) GetRoles() []GetEntitlementsSkuIdsParameterOneOfInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateGuildMemberRequest) GetRolesOk() (*[]GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateGuildMemberRequest) SetRoles(v []GetEntitlementsSkuIdsParameterOneOfInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateGuildMemberRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetMute

`func (o *UpdateGuildMemberRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *UpdateGuildMemberRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *UpdateGuildMemberRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *UpdateGuildMemberRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetDeaf

`func (o *UpdateGuildMemberRequest) GetDeaf() bool`

GetDeaf returns the Deaf field if non-nil, zero value otherwise.

### GetDeafOk

`func (o *UpdateGuildMemberRequest) GetDeafOk() (*bool, bool)`

GetDeafOk returns a tuple with the Deaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaf

`func (o *UpdateGuildMemberRequest) SetDeaf(v bool)`

SetDeaf sets Deaf field to given value.

### HasDeaf

`func (o *UpdateGuildMemberRequest) HasDeaf() bool`

HasDeaf returns a boolean if a field has been set.

### GetChannelId

`func (o *UpdateGuildMemberRequest) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateGuildMemberRequest) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateGuildMemberRequest) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateGuildMemberRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) GetCommunicationDisabledUntil() time.Time`

GetCommunicationDisabledUntil returns the CommunicationDisabledUntil field if non-nil, zero value otherwise.

### GetCommunicationDisabledUntilOk

`func (o *UpdateGuildMemberRequest) GetCommunicationDisabledUntilOk() (*time.Time, bool)`

GetCommunicationDisabledUntilOk returns a tuple with the CommunicationDisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) SetCommunicationDisabledUntil(v time.Time)`

SetCommunicationDisabledUntil sets CommunicationDisabledUntil field to given value.

### HasCommunicationDisabledUntil

`func (o *UpdateGuildMemberRequest) HasCommunicationDisabledUntil() bool`

HasCommunicationDisabledUntil returns a boolean if a field has been set.

### GetFlags

`func (o *UpdateGuildMemberRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateGuildMemberRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateGuildMemberRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateGuildMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


