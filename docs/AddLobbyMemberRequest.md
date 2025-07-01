# AddLobbyMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAddLobbyMemberRequest

`func NewAddLobbyMemberRequest() *AddLobbyMemberRequest`

NewAddLobbyMemberRequest instantiates a new AddLobbyMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLobbyMemberRequestWithDefaults

`func NewAddLobbyMemberRequestWithDefaults() *AddLobbyMemberRequest`

NewAddLobbyMemberRequestWithDefaults instantiates a new AddLobbyMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AddLobbyMemberRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddLobbyMemberRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddLobbyMemberRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddLobbyMemberRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFlags

`func (o *AddLobbyMemberRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *AddLobbyMemberRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *AddLobbyMemberRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *AddLobbyMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *AddLobbyMemberRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *AddLobbyMemberRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


