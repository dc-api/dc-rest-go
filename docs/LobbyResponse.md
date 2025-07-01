# LobbyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApplicationId** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Members** | Pointer to [**[]LobbyMemberResponse**](LobbyMemberResponse.md) |  | [optional] 
**LinkedChannel** | Pointer to [**NullableGuildChannelResponse**](GuildChannelResponse.md) |  | [optional] 

## Methods

### NewLobbyResponse

`func NewLobbyResponse(id string, applicationId string, ) *LobbyResponse`

NewLobbyResponse instantiates a new LobbyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyResponseWithDefaults

`func NewLobbyResponseWithDefaults() *LobbyResponse`

NewLobbyResponseWithDefaults instantiates a new LobbyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LobbyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LobbyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LobbyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *LobbyResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *LobbyResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *LobbyResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetMetadata

`func (o *LobbyResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LobbyResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LobbyResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LobbyResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMembers

`func (o *LobbyResponse) GetMembers() []LobbyMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *LobbyResponse) GetMembersOk() (*[]LobbyMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *LobbyResponse) SetMembers(v []LobbyMemberResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *LobbyResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *LobbyResponse) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *LobbyResponse) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetLinkedChannel

`func (o *LobbyResponse) GetLinkedChannel() GuildChannelResponse`

GetLinkedChannel returns the LinkedChannel field if non-nil, zero value otherwise.

### GetLinkedChannelOk

`func (o *LobbyResponse) GetLinkedChannelOk() (*GuildChannelResponse, bool)`

GetLinkedChannelOk returns a tuple with the LinkedChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedChannel

`func (o *LobbyResponse) SetLinkedChannel(v GuildChannelResponse)`

SetLinkedChannel sets LinkedChannel field to given value.

### HasLinkedChannel

`func (o *LobbyResponse) HasLinkedChannel() bool`

HasLinkedChannel returns a boolean if a field has been set.

### SetLinkedChannelNil

`func (o *LobbyResponse) SetLinkedChannelNil(b bool)`

 SetLinkedChannelNil sets the value for LinkedChannel to be an explicit nil

### UnsetLinkedChannel
`func (o *LobbyResponse) UnsetLinkedChannel()`

UnsetLinkedChannel ensures that no value is present for LinkedChannel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


