# CreateTextThreadWithoutMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AutoArchiveDuration** | Pointer to **int32** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Invitable** | Pointer to **NullableBool** |  | [optional] 

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

`func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDuration() int32`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDurationOk() (*int32, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *CreateTextThreadWithoutMessageRequest) SetAutoArchiveDuration(v int32)`

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

### SetRateLimitPerUserNil

`func (o *CreateTextThreadWithoutMessageRequest) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *CreateTextThreadWithoutMessageRequest) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetType

`func (o *CreateTextThreadWithoutMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTextThreadWithoutMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTextThreadWithoutMessageRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTextThreadWithoutMessageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateTextThreadWithoutMessageRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateTextThreadWithoutMessageRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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

### SetInvitableNil

`func (o *CreateTextThreadWithoutMessageRequest) SetInvitableNil(b bool)`

 SetInvitableNil sets the value for Invitable to be an explicit nil

### UnsetInvitable
`func (o *CreateTextThreadWithoutMessageRequest) UnsetInvitable()`

UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


