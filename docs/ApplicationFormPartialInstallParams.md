# ApplicationFormPartialInstallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to [**[]OAuth2Scopes**](OAuth2Scopes.md) |  | [optional] 
**Permissions** | Pointer to **int32** |  | [optional] 

## Methods

### NewApplicationFormPartialInstallParams

`func NewApplicationFormPartialInstallParams() *ApplicationFormPartialInstallParams`

NewApplicationFormPartialInstallParams instantiates a new ApplicationFormPartialInstallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFormPartialInstallParamsWithDefaults

`func NewApplicationFormPartialInstallParamsWithDefaults() *ApplicationFormPartialInstallParams`

NewApplicationFormPartialInstallParamsWithDefaults instantiates a new ApplicationFormPartialInstallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ApplicationFormPartialInstallParams) GetScopes() []OAuth2Scopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationFormPartialInstallParams) GetScopesOk() (*[]OAuth2Scopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationFormPartialInstallParams) SetScopes(v []OAuth2Scopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationFormPartialInstallParams) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPermissions

`func (o *ApplicationFormPartialInstallParams) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApplicationFormPartialInstallParams) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApplicationFormPartialInstallParams) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApplicationFormPartialInstallParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


