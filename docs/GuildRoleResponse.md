# GuildRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Permissions** | **string** |  | 
**Position** | **int32** |  | 
**Color** | **int32** |  | 
**Hoist** | **bool** |  | 
**Managed** | **bool** |  | 
**Mentionable** | **bool** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**UnicodeEmoji** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**NullableGuildRoleTagsResponse**](GuildRoleTagsResponse.md) |  | [optional] 

## Methods

### NewGuildRoleResponse

`func NewGuildRoleResponse(id string, name string, permissions string, position int32, color int32, hoist bool, managed bool, mentionable bool, ) *GuildRoleResponse`

NewGuildRoleResponse instantiates a new GuildRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildRoleResponseWithDefaults

`func NewGuildRoleResponseWithDefaults() *GuildRoleResponse`

NewGuildRoleResponseWithDefaults instantiates a new GuildRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildRoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildRoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildRoleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GuildRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildRoleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GuildRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildRoleResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildRoleResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPermissions

`func (o *GuildRoleResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GuildRoleResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GuildRoleResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetPosition

`func (o *GuildRoleResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GuildRoleResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GuildRoleResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetColor

`func (o *GuildRoleResponse) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GuildRoleResponse) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GuildRoleResponse) SetColor(v int32)`

SetColor sets Color field to given value.


### GetHoist

`func (o *GuildRoleResponse) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *GuildRoleResponse) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *GuildRoleResponse) SetHoist(v bool)`

SetHoist sets Hoist field to given value.


### GetManaged

`func (o *GuildRoleResponse) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GuildRoleResponse) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GuildRoleResponse) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetMentionable

`func (o *GuildRoleResponse) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *GuildRoleResponse) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *GuildRoleResponse) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.


### GetIcon

`func (o *GuildRoleResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildRoleResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildRoleResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildRoleResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *GuildRoleResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GuildRoleResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetUnicodeEmoji

`func (o *GuildRoleResponse) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *GuildRoleResponse) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *GuildRoleResponse) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *GuildRoleResponse) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.

### SetUnicodeEmojiNil

`func (o *GuildRoleResponse) SetUnicodeEmojiNil(b bool)`

 SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil

### UnsetUnicodeEmoji
`func (o *GuildRoleResponse) UnsetUnicodeEmoji()`

UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil
### GetTags

`func (o *GuildRoleResponse) GetTags() GuildRoleTagsResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GuildRoleResponse) GetTagsOk() (*GuildRoleTagsResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GuildRoleResponse) SetTags(v GuildRoleTagsResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GuildRoleResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GuildRoleResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GuildRoleResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


