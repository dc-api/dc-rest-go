# LobbyMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flags** | **int32** |  | 

## Methods

### NewLobbyMemberResponse

`func NewLobbyMemberResponse(id string, flags int32, ) *LobbyMemberResponse`

NewLobbyMemberResponse instantiates a new LobbyMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLobbyMemberResponseWithDefaults

`func NewLobbyMemberResponseWithDefaults() *LobbyMemberResponse`

NewLobbyMemberResponseWithDefaults instantiates a new LobbyMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LobbyMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LobbyMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LobbyMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *LobbyMemberResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LobbyMemberResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LobbyMemberResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LobbyMemberResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFlags

`func (o *LobbyMemberResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *LobbyMemberResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *LobbyMemberResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


