# UpdateWebhookByTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateWebhookByTokenRequest

`func NewUpdateWebhookByTokenRequest() *UpdateWebhookByTokenRequest`

NewUpdateWebhookByTokenRequest instantiates a new UpdateWebhookByTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookByTokenRequestWithDefaults

`func NewUpdateWebhookByTokenRequestWithDefaults() *UpdateWebhookByTokenRequest`

NewUpdateWebhookByTokenRequestWithDefaults instantiates a new UpdateWebhookByTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhookByTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookByTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookByTokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookByTokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *UpdateWebhookByTokenRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UpdateWebhookByTokenRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UpdateWebhookByTokenRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UpdateWebhookByTokenRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UpdateWebhookByTokenRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UpdateWebhookByTokenRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


