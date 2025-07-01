# LobbyMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**Content** | **string** |  | 
**LobbyId** | **string** |  | 
**ChannelId** | **string** |  | 
**Author** | [**UserResponse**](UserResponse.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flags** | **int32** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 

## Methods

### NewLobbyMessageResponse

`func NewLobbyMessageResponse(id string, type_ int32, content string, lobbyId string, channelId string, author UserResponse, flags int32, ) *LobbyMessageResponse`

NewLobbyMessageResponse instantiates a new LobbyMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyMessageResponseWithDefaults

`func NewLobbyMessageResponseWithDefaults() *LobbyMessageResponse`

NewLobbyMessageResponseWithDefaults instantiates a new LobbyMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LobbyMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LobbyMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LobbyMessageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LobbyMessageResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LobbyMessageResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LobbyMessageResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetContent

`func (o *LobbyMessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *LobbyMessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *LobbyMessageResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetLobbyId

`func (o *LobbyMessageResponse) GetLobbyId() string`

GetLobbyId returns the LobbyId field if non-nil, zero value otherwise.

### GetLobbyIdOk

`func (o *LobbyMessageResponse) GetLobbyIdOk() (*string, bool)`

GetLobbyIdOk returns a tuple with the LobbyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyId

`func (o *LobbyMessageResponse) SetLobbyId(v string)`

SetLobbyId sets LobbyId field to given value.


### GetChannelId

`func (o *LobbyMessageResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *LobbyMessageResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *LobbyMessageResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetAuthor

`func (o *LobbyMessageResponse) GetAuthor() UserResponse`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *LobbyMessageResponse) GetAuthorOk() (*UserResponse, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *LobbyMessageResponse) SetAuthor(v UserResponse)`

SetAuthor sets Author field to given value.


### GetMetadata

`func (o *LobbyMessageResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LobbyMessageResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LobbyMessageResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LobbyMessageResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFlags

`func (o *LobbyMessageResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *LobbyMessageResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *LobbyMessageResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetApplicationId

`func (o *LobbyMessageResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *LobbyMessageResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *LobbyMessageResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *LobbyMessageResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


