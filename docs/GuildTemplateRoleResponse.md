# GuildTemplateRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Permissions** | **string** |  | 
**Color** | **int32** |  | 
**Colors** | Pointer to [**GuildTemplateRoleResponseColors**](GuildTemplateRoleResponseColors.md) |  | [optional] 
**Hoist** | **bool** |  | 
**Mentionable** | **bool** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**UnicodeEmoji** | Pointer to **string** |  | [optional] 

## Methods

### NewGuildTemplateRoleResponse

`func NewGuildTemplateRoleResponse(id int32, name string, permissions string, color int32, hoist bool, mentionable bool, ) *GuildTemplateRoleResponse`

NewGuildTemplateRoleResponse instantiates a new GuildTemplateRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateRoleResponseWithDefaults

`func NewGuildTemplateRoleResponseWithDefaults() *GuildTemplateRoleResponse`

NewGuildTemplateRoleResponseWithDefaults instantiates a new GuildTemplateRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildTemplateRoleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildTemplateRoleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildTemplateRoleResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GuildTemplateRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildTemplateRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildTemplateRoleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *GuildTemplateRoleResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GuildTemplateRoleResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GuildTemplateRoleResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetColor

`func (o *GuildTemplateRoleResponse) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GuildTemplateRoleResponse) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GuildTemplateRoleResponse) SetColor(v int32)`

SetColor sets Color field to given value.


### GetColors

`func (o *GuildTemplateRoleResponse) GetColors() GuildTemplateRoleResponseColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *GuildTemplateRoleResponse) GetColorsOk() (*GuildTemplateRoleResponseColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *GuildTemplateRoleResponse) SetColors(v GuildTemplateRoleResponseColors)`

SetColors sets Colors field to given value.

### HasColors

`func (o *GuildTemplateRoleResponse) HasColors() bool`

HasColors returns a boolean if a field has been set.

### GetHoist

`func (o *GuildTemplateRoleResponse) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *GuildTemplateRoleResponse) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *GuildTemplateRoleResponse) SetHoist(v bool)`

SetHoist sets Hoist field to given value.


### GetMentionable

`func (o *GuildTemplateRoleResponse) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *GuildTemplateRoleResponse) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *GuildTemplateRoleResponse) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.


### GetIcon

`func (o *GuildTemplateRoleResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildTemplateRoleResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildTemplateRoleResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildTemplateRoleResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUnicodeEmoji

`func (o *GuildTemplateRoleResponse) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *GuildTemplateRoleResponse) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *GuildTemplateRoleResponse) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *GuildTemplateRoleResponse) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


