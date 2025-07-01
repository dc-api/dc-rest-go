# CreateGuildRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableInt32** |  | [optional] 
**Color** | Pointer to **NullableInt32** |  | [optional] 
**Hoist** | Pointer to **NullableBool** |  | [optional] 
**Mentionable** | Pointer to **NullableBool** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**UnicodeEmoji** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateGuildRoleRequest

`func NewCreateGuildRoleRequest() *CreateGuildRoleRequest`

NewCreateGuildRoleRequest instantiates a new CreateGuildRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildRoleRequestWithDefaults

`func NewCreateGuildRoleRequestWithDefaults() *CreateGuildRoleRequest`

NewCreateGuildRoleRequestWithDefaults instantiates a new CreateGuildRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGuildRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateGuildRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateGuildRoleRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateGuildRoleRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *CreateGuildRoleRequest) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateGuildRoleRequest) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateGuildRoleRequest) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateGuildRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreateGuildRoleRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreateGuildRoleRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetColor

`func (o *CreateGuildRoleRequest) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateGuildRoleRequest) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateGuildRoleRequest) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateGuildRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateGuildRoleRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateGuildRoleRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetHoist

`func (o *CreateGuildRoleRequest) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *CreateGuildRoleRequest) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *CreateGuildRoleRequest) SetHoist(v bool)`

SetHoist sets Hoist field to given value.

### HasHoist

`func (o *CreateGuildRoleRequest) HasHoist() bool`

HasHoist returns a boolean if a field has been set.

### SetHoistNil

`func (o *CreateGuildRoleRequest) SetHoistNil(b bool)`

 SetHoistNil sets the value for Hoist to be an explicit nil

### UnsetHoist
`func (o *CreateGuildRoleRequest) UnsetHoist()`

UnsetHoist ensures that no value is present for Hoist, not even an explicit nil
### GetMentionable

`func (o *CreateGuildRoleRequest) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *CreateGuildRoleRequest) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *CreateGuildRoleRequest) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.

### HasMentionable

`func (o *CreateGuildRoleRequest) HasMentionable() bool`

HasMentionable returns a boolean if a field has been set.

### SetMentionableNil

`func (o *CreateGuildRoleRequest) SetMentionableNil(b bool)`

 SetMentionableNil sets the value for Mentionable to be an explicit nil

### UnsetMentionable
`func (o *CreateGuildRoleRequest) UnsetMentionable()`

UnsetMentionable ensures that no value is present for Mentionable, not even an explicit nil
### GetIcon

`func (o *CreateGuildRoleRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateGuildRoleRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateGuildRoleRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateGuildRoleRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateGuildRoleRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateGuildRoleRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetUnicodeEmoji

`func (o *CreateGuildRoleRequest) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *CreateGuildRoleRequest) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *CreateGuildRoleRequest) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *CreateGuildRoleRequest) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.

### SetUnicodeEmojiNil

`func (o *CreateGuildRoleRequest) SetUnicodeEmojiNil(b bool)`

 SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil

### UnsetUnicodeEmoji
`func (o *CreateGuildRoleRequest) UnsetUnicodeEmoji()`

UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


