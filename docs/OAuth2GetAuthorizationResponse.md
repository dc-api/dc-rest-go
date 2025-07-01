# OAuth2GetAuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | [**ApplicationResponse**](ApplicationResponse.md) |  | 
**Expires** | **time.Time** |  | 
**Scopes** | **[]string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewOAuth2GetAuthorizationResponse

`func NewOAuth2GetAuthorizationResponse(application ApplicationResponse, expires time.Time, scopes []string, ) *OAuth2GetAuthorizationResponse`

NewOAuth2GetAuthorizationResponse instantiates a new OAuth2GetAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2GetAuthorizationResponseWithDefaults

`func NewOAuth2GetAuthorizationResponseWithDefaults() *OAuth2GetAuthorizationResponse`

NewOAuth2GetAuthorizationResponseWithDefaults instantiates a new OAuth2GetAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *OAuth2GetAuthorizationResponse) GetApplication() ApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *OAuth2GetAuthorizationResponse) GetApplicationOk() (*ApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *OAuth2GetAuthorizationResponse) SetApplication(v ApplicationResponse)`

SetApplication sets Application field to given value.


### GetExpires

`func (o *OAuth2GetAuthorizationResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OAuth2GetAuthorizationResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OAuth2GetAuthorizationResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetScopes

`func (o *OAuth2GetAuthorizationResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2GetAuthorizationResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2GetAuthorizationResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetUser

`func (o *OAuth2GetAuthorizationResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OAuth2GetAuthorizationResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OAuth2GetAuthorizationResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *OAuth2GetAuthorizationResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *OAuth2GetAuthorizationResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *OAuth2GetAuthorizationResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


