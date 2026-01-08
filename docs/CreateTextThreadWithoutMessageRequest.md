# CreateTextThreadWithoutMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AutoArchiveDuration** | Pointer to [**CreateForumThreadRequestAutoArchiveDuration**](CreateForumThreadRequestAutoArchiveDuration.md) |  | [optional] 
**RateLimitPerUser** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**CreateTextThreadWithoutMessageRequestType**](CreateTextThreadWithoutMessageRequestType.md) |  | [optional] 
**Invitable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateTextThreadWithoutMessageRequest

`func NewCreateTextThreadWithoutMessageRequest(name string, ) *CreateTextThreadWithoutMessageRequest`

NewCreateTextThreadWithoutMessageRequest instantiates a new CreateTextThreadWithoutMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTextThreadWithoutMessageRequestWithDefaults

`func NewCreateTextThreadWithoutMessageRequestWithDefaults() *CreateTextThreadWithoutMessageRequest`

NewCreateTextThreadWithoutMessageRequestWithDefaults instantiates a new CreateTextThreadWithoutMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTextThreadWithoutMessageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTextThreadWithoutMessageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTextThreadWithoutMessageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAutoArchiveDuration

`func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *CreateTextThreadWithoutMessageRequest) SetAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *CreateTextThreadWithoutMessageRequest) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *CreateTextThreadWithoutMessageRequest) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreateTextThreadWithoutMessageRequest) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreateTextThreadWithoutMessageRequest) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreateTextThreadWithoutMessageRequest) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### GetType

`func (o *CreateTextThreadWithoutMessageRequest) GetType() CreateTextThreadWithoutMessageRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTextThreadWithoutMessageRequest) GetTypeOk() (*CreateTextThreadWithoutMessageRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTextThreadWithoutMessageRequest) SetType(v CreateTextThreadWithoutMessageRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTextThreadWithoutMessageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInvitable

`func (o *CreateTextThreadWithoutMessageRequest) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *CreateTextThreadWithoutMessageRequest) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *CreateTextThreadWithoutMessageRequest) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *CreateTextThreadWithoutMessageRequest) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


