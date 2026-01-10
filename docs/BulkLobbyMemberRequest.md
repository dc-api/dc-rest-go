# BulkLobbyMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flags** | Pointer to [**CreateOrJoinLobbyRequestFlags**](CreateOrJoinLobbyRequestFlags.md) |  | [optional] 
**RemoveMember** | Pointer to **bool** |  | [optional] 

## Methods

### NewBulkLobbyMemberRequest

`func NewBulkLobbyMemberRequest(id string, ) *BulkLobbyMemberRequest`

NewBulkLobbyMemberRequest instantiates a new BulkLobbyMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkLobbyMemberRequestWithDefaults

`func NewBulkLobbyMemberRequestWithDefaults() *BulkLobbyMemberRequest`

NewBulkLobbyMemberRequestWithDefaults instantiates a new BulkLobbyMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkLobbyMemberRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkLobbyMemberRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkLobbyMemberRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *BulkLobbyMemberRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BulkLobbyMemberRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BulkLobbyMemberRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BulkLobbyMemberRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFlags

`func (o *BulkLobbyMemberRequest) GetFlags() CreateOrJoinLobbyRequestFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BulkLobbyMemberRequest) GetFlagsOk() (*CreateOrJoinLobbyRequestFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BulkLobbyMemberRequest) SetFlags(v CreateOrJoinLobbyRequestFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BulkLobbyMemberRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetRemoveMember

`func (o *BulkLobbyMemberRequest) GetRemoveMember() bool`

GetRemoveMember returns the RemoveMember field if non-nil, zero value otherwise.

### GetRemoveMemberOk

`func (o *BulkLobbyMemberRequest) GetRemoveMemberOk() (*bool, bool)`

GetRemoveMemberOk returns a tuple with the RemoveMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveMember

`func (o *BulkLobbyMemberRequest) SetRemoveMember(v bool)`

SetRemoveMember sets RemoveMember field to given value.

### HasRemoveMember

`func (o *BulkLobbyMemberRequest) HasRemoveMember() bool`

HasRemoveMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


