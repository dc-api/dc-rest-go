# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **int32** |  | [optional] 
**Color** | Pointer to **int32** |  | [optional] 
**Colors** | Pointer to [**CreateRoleRequestColors**](CreateRoleRequestColors.md) |  | [optional] 
**Hoist** | Pointer to **bool** |  | [optional] 
**Mentionable** | Pointer to **bool** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**UnicodeEmoji** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest() *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateRoleRequest) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleRequest) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleRequest) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetColor

`func (o *CreateRoleRequest) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateRoleRequest) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateRoleRequest) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetColors

`func (o *CreateRoleRequest) GetColors() CreateRoleRequestColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *CreateRoleRequest) GetColorsOk() (*CreateRoleRequestColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *CreateRoleRequest) SetColors(v CreateRoleRequestColors)`

SetColors sets Colors field to given value.

### HasColors

`func (o *CreateRoleRequest) HasColors() bool`

HasColors returns a boolean if a field has been set.

### GetHoist

`func (o *CreateRoleRequest) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *CreateRoleRequest) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *CreateRoleRequest) SetHoist(v bool)`

SetHoist sets Hoist field to given value.

### HasHoist

`func (o *CreateRoleRequest) HasHoist() bool`

HasHoist returns a boolean if a field has been set.

### GetMentionable

`func (o *CreateRoleRequest) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *CreateRoleRequest) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *CreateRoleRequest) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.

### HasMentionable

`func (o *CreateRoleRequest) HasMentionable() bool`

HasMentionable returns a boolean if a field has been set.

### GetIcon

`func (o *CreateRoleRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateRoleRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateRoleRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateRoleRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUnicodeEmoji

`func (o *CreateRoleRequest) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *CreateRoleRequest) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *CreateRoleRequest) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *CreateRoleRequest) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


