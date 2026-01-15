# ApplicationOAuth2InstallParamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | [**[]OAuth2Scopes**](OAuth2Scopes.md) |  | 
**Permissions** | **string** |  | 

## Methods

### NewApplicationOAuth2InstallParamsResponse

`func NewApplicationOAuth2InstallParamsResponse(scopes []OAuth2Scopes, permissions string, ) *ApplicationOAuth2InstallParamsResponse`

NewApplicationOAuth2InstallParamsResponse instantiates a new ApplicationOAuth2InstallParamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOAuth2InstallParamsResponseWithDefaults

`func NewApplicationOAuth2InstallParamsResponseWithDefaults() *ApplicationOAuth2InstallParamsResponse`

NewApplicationOAuth2InstallParamsResponseWithDefaults instantiates a new ApplicationOAuth2InstallParamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ApplicationOAuth2InstallParamsResponse) GetScopes() []OAuth2Scopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationOAuth2InstallParamsResponse) GetScopesOk() (*[]OAuth2Scopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationOAuth2InstallParamsResponse) SetScopes(v []OAuth2Scopes)`

SetScopes sets Scopes field to given value.


### GetPermissions

`func (o *ApplicationOAuth2InstallParamsResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApplicationOAuth2InstallParamsResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApplicationOAuth2InstallParamsResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


