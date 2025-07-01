# CreateOrJoinLobbyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeoutSeconds** | Pointer to **NullableInt32** |  | [optional] 
**LobbyMetadata** | Pointer to **map[string]string** |  | [optional] 
**MemberMetadata** | Pointer to **map[string]string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewCreateOrJoinLobbyRequest

`func NewCreateOrJoinLobbyRequest(secret string, ) *CreateOrJoinLobbyRequest`

NewCreateOrJoinLobbyRequest instantiates a new CreateOrJoinLobbyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrJoinLobbyRequestWithDefaults

`func NewCreateOrJoinLobbyRequestWithDefaults() *CreateOrJoinLobbyRequest`

NewCreateOrJoinLobbyRequestWithDefaults instantiates a new CreateOrJoinLobbyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeoutSeconds

`func (o *CreateOrJoinLobbyRequest) GetIdleTimeoutSeconds() int32`

GetIdleTimeoutSeconds returns the IdleTimeoutSeconds field if non-nil, zero value otherwise.

### GetIdleTimeoutSecondsOk

`func (o *CreateOrJoinLobbyRequest) GetIdleTimeoutSecondsOk() (*int32, bool)`

GetIdleTimeoutSecondsOk returns a tuple with the IdleTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutSeconds

`func (o *CreateOrJoinLobbyRequest) SetIdleTimeoutSeconds(v int32)`

SetIdleTimeoutSeconds sets IdleTimeoutSeconds field to given value.

### HasIdleTimeoutSeconds

`func (o *CreateOrJoinLobbyRequest) HasIdleTimeoutSeconds() bool`

HasIdleTimeoutSeconds returns a boolean if a field has been set.

### SetIdleTimeoutSecondsNil

`func (o *CreateOrJoinLobbyRequest) SetIdleTimeoutSecondsNil(b bool)`

 SetIdleTimeoutSecondsNil sets the value for IdleTimeoutSeconds to be an explicit nil

### UnsetIdleTimeoutSeconds
`func (o *CreateOrJoinLobbyRequest) UnsetIdleTimeoutSeconds()`

UnsetIdleTimeoutSeconds ensures that no value is present for IdleTimeoutSeconds, not even an explicit nil
### GetLobbyMetadata

`func (o *CreateOrJoinLobbyRequest) GetLobbyMetadata() map[string]string`

GetLobbyMetadata returns the LobbyMetadata field if non-nil, zero value otherwise.

### GetLobbyMetadataOk

`func (o *CreateOrJoinLobbyRequest) GetLobbyMetadataOk() (*map[string]string, bool)`

GetLobbyMetadataOk returns a tuple with the LobbyMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyMetadata

`func (o *CreateOrJoinLobbyRequest) SetLobbyMetadata(v map[string]string)`

SetLobbyMetadata sets LobbyMetadata field to given value.

### HasLobbyMetadata

`func (o *CreateOrJoinLobbyRequest) HasLobbyMetadata() bool`

HasLobbyMetadata returns a boolean if a field has been set.

### GetMemberMetadata

`func (o *CreateOrJoinLobbyRequest) GetMemberMetadata() map[string]string`

GetMemberMetadata returns the MemberMetadata field if non-nil, zero value otherwise.

### GetMemberMetadataOk

`func (o *CreateOrJoinLobbyRequest) GetMemberMetadataOk() (*map[string]string, bool)`

GetMemberMetadataOk returns a tuple with the MemberMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMetadata

`func (o *CreateOrJoinLobbyRequest) SetMemberMetadata(v map[string]string)`

SetMemberMetadata sets MemberMetadata field to given value.

### HasMemberMetadata

`func (o *CreateOrJoinLobbyRequest) HasMemberMetadata() bool`

HasMemberMetadata returns a boolean if a field has been set.

### GetSecret

`func (o *CreateOrJoinLobbyRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateOrJoinLobbyRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateOrJoinLobbyRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


