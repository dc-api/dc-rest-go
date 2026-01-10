# GuildTemplateChannelResponsePermissionOverwritesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelPermissionOverwrites**](ChannelPermissionOverwrites.md) |  | 
**Allow** | **string** |  | 
**Deny** | **string** |  | 

## Methods

### NewGuildTemplateChannelResponsePermissionOverwritesInner

`func NewGuildTemplateChannelResponsePermissionOverwritesInner(id string, type_ ChannelPermissionOverwrites, allow string, deny string, ) *GuildTemplateChannelResponsePermissionOverwritesInner`

NewGuildTemplateChannelResponsePermissionOverwritesInner instantiates a new GuildTemplateChannelResponsePermissionOverwritesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateChannelResponsePermissionOverwritesInnerWithDefaults

`func NewGuildTemplateChannelResponsePermissionOverwritesInnerWithDefaults() *GuildTemplateChannelResponsePermissionOverwritesInner`

NewGuildTemplateChannelResponsePermissionOverwritesInnerWithDefaults instantiates a new GuildTemplateChannelResponsePermissionOverwritesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetType() ChannelPermissionOverwrites`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetTypeOk() (*ChannelPermissionOverwrites, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) SetType(v ChannelPermissionOverwrites)`

SetType sets Type field to given value.


### GetAllow

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetAllow() string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetAllowOk() (*string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) SetAllow(v string)`

SetAllow sets Allow field to given value.


### GetDeny

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetDeny() string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) GetDenyOk() (*string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *GuildTemplateChannelResponsePermissionOverwritesInner) SetDeny(v string)`

SetDeny sets Deny field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


