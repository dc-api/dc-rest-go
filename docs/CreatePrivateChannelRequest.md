# CreatePrivateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | Pointer to **string** |  | [optional] 
**AccessTokens** | Pointer to **[]string** |  | [optional] 
**Nicks** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreatePrivateChannelRequest

`func NewCreatePrivateChannelRequest() *CreatePrivateChannelRequest`

NewCreatePrivateChannelRequest instantiates a new CreatePrivateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrivateChannelRequestWithDefaults

`func NewCreatePrivateChannelRequestWithDefaults() *CreatePrivateChannelRequest`

NewCreatePrivateChannelRequestWithDefaults instantiates a new CreatePrivateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *CreatePrivateChannelRequest) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *CreatePrivateChannelRequest) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *CreatePrivateChannelRequest) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *CreatePrivateChannelRequest) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetAccessTokens

`func (o *CreatePrivateChannelRequest) GetAccessTokens() []string`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *CreatePrivateChannelRequest) GetAccessTokensOk() (*[]string, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *CreatePrivateChannelRequest) SetAccessTokens(v []string)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *CreatePrivateChannelRequest) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.

### SetAccessTokensNil

`func (o *CreatePrivateChannelRequest) SetAccessTokensNil(b bool)`

 SetAccessTokensNil sets the value for AccessTokens to be an explicit nil

### UnsetAccessTokens
`func (o *CreatePrivateChannelRequest) UnsetAccessTokens()`

UnsetAccessTokens ensures that no value is present for AccessTokens, not even an explicit nil
### GetNicks

`func (o *CreatePrivateChannelRequest) GetNicks() map[string]string`

GetNicks returns the Nicks field if non-nil, zero value otherwise.

### GetNicksOk

`func (o *CreatePrivateChannelRequest) GetNicksOk() (*map[string]string, bool)`

GetNicksOk returns a tuple with the Nicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicks

`func (o *CreatePrivateChannelRequest) SetNicks(v map[string]string)`

SetNicks sets Nicks field to given value.

### HasNicks

`func (o *CreatePrivateChannelRequest) HasNicks() bool`

HasNicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


