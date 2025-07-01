# CreateThreadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AutoArchiveDuration** | Pointer to **int32** |  | [optional] 
**RateLimitPerUser** | Pointer to **int32** |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Message** | [**BaseCreateMessageCreateRequest**](BaseCreateMessageCreateRequest.md) |  | 
**Type** | Pointer to **int32** |  | [optional] 
**Invitable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateThreadRequest

`func NewCreateThreadRequest(name string, message BaseCreateMessageCreateRequest, ) *CreateThreadRequest`

NewCreateThreadRequest instantiates a new CreateThreadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThreadRequestWithDefaults

`func NewCreateThreadRequestWithDefaults() *CreateThreadRequest`

NewCreateThreadRequestWithDefaults instantiates a new CreateThreadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateThreadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateThreadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateThreadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAutoArchiveDuration

`func (o *CreateThreadRequest) GetAutoArchiveDuration() int32`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *CreateThreadRequest) GetAutoArchiveDurationOk() (*int32, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *CreateThreadRequest) SetAutoArchiveDuration(v int32)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *CreateThreadRequest) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *CreateThreadRequest) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreateThreadRequest) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreateThreadRequest) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreateThreadRequest) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### GetAppliedTags

`func (o *CreateThreadRequest) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *CreateThreadRequest) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *CreateThreadRequest) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *CreateThreadRequest) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### GetMessage

`func (o *CreateThreadRequest) GetMessage() BaseCreateMessageCreateRequest`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateThreadRequest) GetMessageOk() (*BaseCreateMessageCreateRequest, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateThreadRequest) SetMessage(v BaseCreateMessageCreateRequest)`

SetMessage sets Message field to given value.


### GetType

`func (o *CreateThreadRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateThreadRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateThreadRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateThreadRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInvitable

`func (o *CreateThreadRequest) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *CreateThreadRequest) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *CreateThreadRequest) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *CreateThreadRequest) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


