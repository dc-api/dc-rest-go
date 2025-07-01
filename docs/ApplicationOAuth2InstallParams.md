# ApplicationOAuth2InstallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]string** |  | [optional] 
**Permissions** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewApplicationOAuth2InstallParams

`func NewApplicationOAuth2InstallParams() *ApplicationOAuth2InstallParams`

NewApplicationOAuth2InstallParams instantiates a new ApplicationOAuth2InstallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOAuth2InstallParamsWithDefaults

`func NewApplicationOAuth2InstallParamsWithDefaults() *ApplicationOAuth2InstallParams`

NewApplicationOAuth2InstallParamsWithDefaults instantiates a new ApplicationOAuth2InstallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ApplicationOAuth2InstallParams) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationOAuth2InstallParams) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationOAuth2InstallParams) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationOAuth2InstallParams) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ApplicationOAuth2InstallParams) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ApplicationOAuth2InstallParams) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetPermissions

`func (o *ApplicationOAuth2InstallParams) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApplicationOAuth2InstallParams) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApplicationOAuth2InstallParams) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApplicationOAuth2InstallParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ApplicationOAuth2InstallParams) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApplicationOAuth2InstallParams) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


