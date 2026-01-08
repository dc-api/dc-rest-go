# BotAccountPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Avatar** | Pointer to **string** |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 

## Methods

### NewBotAccountPatchRequest

`func NewBotAccountPatchRequest(username string, ) *BotAccountPatchRequest`

NewBotAccountPatchRequest instantiates a new BotAccountPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotAccountPatchRequestWithDefaults

`func NewBotAccountPatchRequestWithDefaults() *BotAccountPatchRequest`

NewBotAccountPatchRequestWithDefaults instantiates a new BotAccountPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BotAccountPatchRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BotAccountPatchRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BotAccountPatchRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAvatar

`func (o *BotAccountPatchRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *BotAccountPatchRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *BotAccountPatchRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *BotAccountPatchRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBanner

`func (o *BotAccountPatchRequest) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *BotAccountPatchRequest) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *BotAccountPatchRequest) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *BotAccountPatchRequest) HasBanner() bool`

HasBanner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


