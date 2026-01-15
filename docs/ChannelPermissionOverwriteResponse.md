# ChannelPermissionOverwriteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelPermissionOverwrites**](ChannelPermissionOverwrites.md) |  | 
**Allow** | **string** |  | 
**Deny** | **string** |  | 

## Methods

### NewChannelPermissionOverwriteResponse

`func NewChannelPermissionOverwriteResponse(id string, type_ ChannelPermissionOverwrites, allow string, deny string, ) *ChannelPermissionOverwriteResponse`

NewChannelPermissionOverwriteResponse instantiates a new ChannelPermissionOverwriteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPermissionOverwriteResponseWithDefaults

`func NewChannelPermissionOverwriteResponseWithDefaults() *ChannelPermissionOverwriteResponse`

NewChannelPermissionOverwriteResponseWithDefaults instantiates a new ChannelPermissionOverwriteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelPermissionOverwriteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelPermissionOverwriteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelPermissionOverwriteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ChannelPermissionOverwriteResponse) GetType() ChannelPermissionOverwrites`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelPermissionOverwriteResponse) GetTypeOk() (*ChannelPermissionOverwrites, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelPermissionOverwriteResponse) SetType(v ChannelPermissionOverwrites)`

SetType sets Type field to given value.


### GetAllow

`func (o *ChannelPermissionOverwriteResponse) GetAllow() string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *ChannelPermissionOverwriteResponse) GetAllowOk() (*string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *ChannelPermissionOverwriteResponse) SetAllow(v string)`

SetAllow sets Allow field to given value.


### GetDeny

`func (o *ChannelPermissionOverwriteResponse) GetDeny() string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *ChannelPermissionOverwriteResponse) GetDenyOk() (*string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *ChannelPermissionOverwriteResponse) SetDeny(v string)`

SetDeny sets Deny field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


