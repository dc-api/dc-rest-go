# ChannelPermissionOverwriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Allow** | Pointer to **NullableInt32** |  | [optional] 
**Deny** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewChannelPermissionOverwriteRequest

`func NewChannelPermissionOverwriteRequest(id string, ) *ChannelPermissionOverwriteRequest`

NewChannelPermissionOverwriteRequest instantiates a new ChannelPermissionOverwriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPermissionOverwriteRequestWithDefaults

`func NewChannelPermissionOverwriteRequestWithDefaults() *ChannelPermissionOverwriteRequest`

NewChannelPermissionOverwriteRequestWithDefaults instantiates a new ChannelPermissionOverwriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelPermissionOverwriteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelPermissionOverwriteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelPermissionOverwriteRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ChannelPermissionOverwriteRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelPermissionOverwriteRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelPermissionOverwriteRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelPermissionOverwriteRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ChannelPermissionOverwriteRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChannelPermissionOverwriteRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAllow

`func (o *ChannelPermissionOverwriteRequest) GetAllow() int32`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *ChannelPermissionOverwriteRequest) GetAllowOk() (*int32, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *ChannelPermissionOverwriteRequest) SetAllow(v int32)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *ChannelPermissionOverwriteRequest) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### SetAllowNil

`func (o *ChannelPermissionOverwriteRequest) SetAllowNil(b bool)`

 SetAllowNil sets the value for Allow to be an explicit nil

### UnsetAllow
`func (o *ChannelPermissionOverwriteRequest) UnsetAllow()`

UnsetAllow ensures that no value is present for Allow, not even an explicit nil
### GetDeny

`func (o *ChannelPermissionOverwriteRequest) GetDeny() int32`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *ChannelPermissionOverwriteRequest) GetDenyOk() (*int32, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *ChannelPermissionOverwriteRequest) SetDeny(v int32)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *ChannelPermissionOverwriteRequest) HasDeny() bool`

HasDeny returns a boolean if a field has been set.

### SetDenyNil

`func (o *ChannelPermissionOverwriteRequest) SetDenyNil(b bool)`

 SetDenyNil sets the value for Deny to be an explicit nil

### UnsetDeny
`func (o *ChannelPermissionOverwriteRequest) UnsetDeny()`

UnsetDeny ensures that no value is present for Deny, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


