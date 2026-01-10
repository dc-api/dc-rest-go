# SetChannelPermissionOverwriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SetChannelPermissionOverwriteRequestType**](SetChannelPermissionOverwriteRequestType.md) |  | [optional] 
**Allow** | Pointer to **int32** |  | [optional] 
**Deny** | Pointer to **int32** |  | [optional] 

## Methods

### NewSetChannelPermissionOverwriteRequest

`func NewSetChannelPermissionOverwriteRequest() *SetChannelPermissionOverwriteRequest`

NewSetChannelPermissionOverwriteRequest instantiates a new SetChannelPermissionOverwriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChannelPermissionOverwriteRequestWithDefaults

`func NewSetChannelPermissionOverwriteRequestWithDefaults() *SetChannelPermissionOverwriteRequest`

NewSetChannelPermissionOverwriteRequestWithDefaults instantiates a new SetChannelPermissionOverwriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetChannelPermissionOverwriteRequest) GetType() SetChannelPermissionOverwriteRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetChannelPermissionOverwriteRequest) GetTypeOk() (*SetChannelPermissionOverwriteRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetChannelPermissionOverwriteRequest) SetType(v SetChannelPermissionOverwriteRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *SetChannelPermissionOverwriteRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAllow

`func (o *SetChannelPermissionOverwriteRequest) GetAllow() int32`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *SetChannelPermissionOverwriteRequest) GetAllowOk() (*int32, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *SetChannelPermissionOverwriteRequest) SetAllow(v int32)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *SetChannelPermissionOverwriteRequest) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetDeny

`func (o *SetChannelPermissionOverwriteRequest) GetDeny() int32`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *SetChannelPermissionOverwriteRequest) GetDenyOk() (*int32, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *SetChannelPermissionOverwriteRequest) SetDeny(v int32)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *SetChannelPermissionOverwriteRequest) HasDeny() bool`

HasDeny returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


