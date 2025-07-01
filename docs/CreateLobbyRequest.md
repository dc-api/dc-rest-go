# CreateLobbyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeoutSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Members** | Pointer to [**[]LobbyMemberRequest**](LobbyMemberRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateLobbyRequest

`func NewCreateLobbyRequest() *CreateLobbyRequest`

NewCreateLobbyRequest instantiates a new CreateLobbyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLobbyRequestWithDefaults

`func NewCreateLobbyRequestWithDefaults() *CreateLobbyRequest`

NewCreateLobbyRequestWithDefaults instantiates a new CreateLobbyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeoutSeconds

`func (o *CreateLobbyRequest) GetIdleTimeoutSeconds() int32`

GetIdleTimeoutSeconds returns the IdleTimeoutSeconds field if non-nil, zero value otherwise.

### GetIdleTimeoutSecondsOk

`func (o *CreateLobbyRequest) GetIdleTimeoutSecondsOk() (*int32, bool)`

GetIdleTimeoutSecondsOk returns a tuple with the IdleTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutSeconds

`func (o *CreateLobbyRequest) SetIdleTimeoutSeconds(v int32)`

SetIdleTimeoutSeconds sets IdleTimeoutSeconds field to given value.

### HasIdleTimeoutSeconds

`func (o *CreateLobbyRequest) HasIdleTimeoutSeconds() bool`

HasIdleTimeoutSeconds returns a boolean if a field has been set.

### SetIdleTimeoutSecondsNil

`func (o *CreateLobbyRequest) SetIdleTimeoutSecondsNil(b bool)`

 SetIdleTimeoutSecondsNil sets the value for IdleTimeoutSeconds to be an explicit nil

### UnsetIdleTimeoutSeconds
`func (o *CreateLobbyRequest) UnsetIdleTimeoutSeconds()`

UnsetIdleTimeoutSeconds ensures that no value is present for IdleTimeoutSeconds, not even an explicit nil
### GetMembers

`func (o *CreateLobbyRequest) GetMembers() []LobbyMemberRequest`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateLobbyRequest) GetMembersOk() (*[]LobbyMemberRequest, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateLobbyRequest) SetMembers(v []LobbyMemberRequest)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateLobbyRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *CreateLobbyRequest) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *CreateLobbyRequest) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetMetadata

`func (o *CreateLobbyRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateLobbyRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateLobbyRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateLobbyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


