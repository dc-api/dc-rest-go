# CreateGuildRequestRoleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableInt32** |  | [optional] 
**Color** | Pointer to **NullableInt32** |  | [optional] 
**Hoist** | Pointer to **NullableBool** |  | [optional] 
**Mentionable** | Pointer to **NullableBool** |  | [optional] 
**UnicodeEmoji** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateGuildRequestRoleItem

`func NewCreateGuildRequestRoleItem(id int32, ) *CreateGuildRequestRoleItem`

NewCreateGuildRequestRoleItem instantiates a new CreateGuildRequestRoleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildRequestRoleItemWithDefaults

`func NewCreateGuildRequestRoleItemWithDefaults() *CreateGuildRequestRoleItem`

NewCreateGuildRequestRoleItemWithDefaults instantiates a new CreateGuildRequestRoleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateGuildRequestRoleItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGuildRequestRoleItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGuildRequestRoleItem) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CreateGuildRequestRoleItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildRequestRoleItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildRequestRoleItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateGuildRequestRoleItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateGuildRequestRoleItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateGuildRequestRoleItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *CreateGuildRequestRoleItem) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateGuildRequestRoleItem) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateGuildRequestRoleItem) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateGuildRequestRoleItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreateGuildRequestRoleItem) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreateGuildRequestRoleItem) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetColor

`func (o *CreateGuildRequestRoleItem) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateGuildRequestRoleItem) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateGuildRequestRoleItem) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateGuildRequestRoleItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateGuildRequestRoleItem) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateGuildRequestRoleItem) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetHoist

`func (o *CreateGuildRequestRoleItem) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *CreateGuildRequestRoleItem) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *CreateGuildRequestRoleItem) SetHoist(v bool)`

SetHoist sets Hoist field to given value.

### HasHoist

`func (o *CreateGuildRequestRoleItem) HasHoist() bool`

HasHoist returns a boolean if a field has been set.

### SetHoistNil

`func (o *CreateGuildRequestRoleItem) SetHoistNil(b bool)`

 SetHoistNil sets the value for Hoist to be an explicit nil

### UnsetHoist
`func (o *CreateGuildRequestRoleItem) UnsetHoist()`

UnsetHoist ensures that no value is present for Hoist, not even an explicit nil
### GetMentionable

`func (o *CreateGuildRequestRoleItem) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *CreateGuildRequestRoleItem) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *CreateGuildRequestRoleItem) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.

### HasMentionable

`func (o *CreateGuildRequestRoleItem) HasMentionable() bool`

HasMentionable returns a boolean if a field has been set.

### SetMentionableNil

`func (o *CreateGuildRequestRoleItem) SetMentionableNil(b bool)`

 SetMentionableNil sets the value for Mentionable to be an explicit nil

### UnsetMentionable
`func (o *CreateGuildRequestRoleItem) UnsetMentionable()`

UnsetMentionable ensures that no value is present for Mentionable, not even an explicit nil
### GetUnicodeEmoji

`func (o *CreateGuildRequestRoleItem) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *CreateGuildRequestRoleItem) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *CreateGuildRequestRoleItem) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *CreateGuildRequestRoleItem) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.

### SetUnicodeEmojiNil

`func (o *CreateGuildRequestRoleItem) SetUnicodeEmojiNil(b bool)`

 SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil

### UnsetUnicodeEmoji
`func (o *CreateGuildRequestRoleItem) UnsetUnicodeEmoji()`

UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


