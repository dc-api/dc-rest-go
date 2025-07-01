# GuildBanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserResponse**](UserResponse.md) |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGuildBanResponse

`func NewGuildBanResponse(user UserResponse, ) *GuildBanResponse`

NewGuildBanResponse instantiates a new GuildBanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildBanResponseWithDefaults

`func NewGuildBanResponseWithDefaults() *GuildBanResponse`

NewGuildBanResponseWithDefaults instantiates a new GuildBanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GuildBanResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GuildBanResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GuildBanResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetReason

`func (o *GuildBanResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GuildBanResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GuildBanResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GuildBanResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GuildBanResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GuildBanResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


