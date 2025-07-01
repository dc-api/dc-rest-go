# OAuth2GetOpenIDConnectUserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**EmailVerified** | Pointer to **NullableBool** |  | [optional] 
**PreferredUsername** | Pointer to **NullableString** |  | [optional] 
**Nickname** | Pointer to **NullableString** |  | [optional] 
**Picture** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuth2GetOpenIDConnectUserInfoResponse

`func NewOAuth2GetOpenIDConnectUserInfoResponse(sub string, ) *OAuth2GetOpenIDConnectUserInfoResponse`

NewOAuth2GetOpenIDConnectUserInfoResponse instantiates a new OAuth2GetOpenIDConnectUserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2GetOpenIDConnectUserInfoResponseWithDefaults

`func NewOAuth2GetOpenIDConnectUserInfoResponseWithDefaults() *OAuth2GetOpenIDConnectUserInfoResponse`

NewOAuth2GetOpenIDConnectUserInfoResponseWithDefaults instantiates a new OAuth2GetOpenIDConnectUserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetSub(v string)`

SetSub sets Sub field to given value.


### GetEmail

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetPreferredUsername

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### SetPreferredUsernameNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPreferredUsernameNil(b bool)`

 SetPreferredUsernameNil sets the value for PreferredUsername to be an explicit nil

### UnsetPreferredUsername
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetPreferredUsername()`

UnsetPreferredUsername ensures that no value is present for PreferredUsername, not even an explicit nil
### GetNickname

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetPicture

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### SetPictureNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetPictureNil(b bool)`

 SetPictureNil sets the value for Picture to be an explicit nil

### UnsetPicture
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetPicture()`

UnsetPicture ensures that no value is present for Picture, not even an explicit nil
### GetLocale

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *OAuth2GetOpenIDConnectUserInfoResponse) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *OAuth2GetOpenIDConnectUserInfoResponse) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


