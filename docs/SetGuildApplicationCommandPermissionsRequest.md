# SetGuildApplicationCommandPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**[]ApplicationCommandPermission**](ApplicationCommandPermission.md) |  | [optional] 

## Methods

### NewSetGuildApplicationCommandPermissionsRequest

`func NewSetGuildApplicationCommandPermissionsRequest() *SetGuildApplicationCommandPermissionsRequest`

NewSetGuildApplicationCommandPermissionsRequest instantiates a new SetGuildApplicationCommandPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGuildApplicationCommandPermissionsRequestWithDefaults

`func NewSetGuildApplicationCommandPermissionsRequestWithDefaults() *SetGuildApplicationCommandPermissionsRequest`

NewSetGuildApplicationCommandPermissionsRequestWithDefaults instantiates a new SetGuildApplicationCommandPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *SetGuildApplicationCommandPermissionsRequest) GetPermissions() []ApplicationCommandPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SetGuildApplicationCommandPermissionsRequest) GetPermissionsOk() (*[]ApplicationCommandPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SetGuildApplicationCommandPermissionsRequest) SetPermissions(v []ApplicationCommandPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SetGuildApplicationCommandPermissionsRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SetGuildApplicationCommandPermissionsRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SetGuildApplicationCommandPermissionsRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


