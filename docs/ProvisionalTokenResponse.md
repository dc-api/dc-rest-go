# ProvisionalTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | **string** |  | 
**AccessToken** | **string** |  | 
**ExpiresIn** | **int32** |  | 
**Scope** | **string** |  | 
**IdToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**ExpiresAtS** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewProvisionalTokenResponse

`func NewProvisionalTokenResponse(tokenType string, accessToken string, expiresIn int32, scope string, idToken string, ) *ProvisionalTokenResponse`

NewProvisionalTokenResponse instantiates a new ProvisionalTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionalTokenResponseWithDefaults

`func NewProvisionalTokenResponseWithDefaults() *ProvisionalTokenResponse`

NewProvisionalTokenResponseWithDefaults instantiates a new ProvisionalTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *ProvisionalTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ProvisionalTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ProvisionalTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetAccessToken

`func (o *ProvisionalTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ProvisionalTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ProvisionalTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *ProvisionalTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ProvisionalTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ProvisionalTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetScope

`func (o *ProvisionalTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProvisionalTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProvisionalTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.


### GetIdToken

`func (o *ProvisionalTokenResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *ProvisionalTokenResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *ProvisionalTokenResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetRefreshToken

`func (o *ProvisionalTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ProvisionalTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ProvisionalTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ProvisionalTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *ProvisionalTokenResponse) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *ProvisionalTokenResponse) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetScopes

`func (o *ProvisionalTokenResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProvisionalTokenResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProvisionalTokenResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProvisionalTokenResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ProvisionalTokenResponse) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ProvisionalTokenResponse) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetExpiresAtS

`func (o *ProvisionalTokenResponse) GetExpiresAtS() int32`

GetExpiresAtS returns the ExpiresAtS field if non-nil, zero value otherwise.

### GetExpiresAtSOk

`func (o *ProvisionalTokenResponse) GetExpiresAtSOk() (*int32, bool)`

GetExpiresAtSOk returns a tuple with the ExpiresAtS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtS

`func (o *ProvisionalTokenResponse) SetExpiresAtS(v int32)`

SetExpiresAtS sets ExpiresAtS field to given value.

### HasExpiresAtS

`func (o *ProvisionalTokenResponse) HasExpiresAtS() bool`

HasExpiresAtS returns a boolean if a field has been set.

### SetExpiresAtSNil

`func (o *ProvisionalTokenResponse) SetExpiresAtSNil(b bool)`

 SetExpiresAtSNil sets the value for ExpiresAtS to be an explicit nil

### UnsetExpiresAtS
`func (o *ProvisionalTokenResponse) UnsetExpiresAtS()`

UnsetExpiresAtS ensures that no value is present for ExpiresAtS, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


