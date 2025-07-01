# CreateChannelInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAge** | Pointer to **int32** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**TargetUserId** | Pointer to **string** |  | [optional] 
**TargetApplicationId** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateChannelInviteRequest

`func NewCreateChannelInviteRequest() *CreateChannelInviteRequest`

NewCreateChannelInviteRequest instantiates a new CreateChannelInviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChannelInviteRequestWithDefaults

`func NewCreateChannelInviteRequestWithDefaults() *CreateChannelInviteRequest`

NewCreateChannelInviteRequestWithDefaults instantiates a new CreateChannelInviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAge

`func (o *CreateChannelInviteRequest) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CreateChannelInviteRequest) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CreateChannelInviteRequest) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CreateChannelInviteRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetTemporary

`func (o *CreateChannelInviteRequest) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CreateChannelInviteRequest) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CreateChannelInviteRequest) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CreateChannelInviteRequest) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetMaxUses

`func (o *CreateChannelInviteRequest) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *CreateChannelInviteRequest) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *CreateChannelInviteRequest) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *CreateChannelInviteRequest) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetUnique

`func (o *CreateChannelInviteRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *CreateChannelInviteRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *CreateChannelInviteRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *CreateChannelInviteRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetTargetUserId

`func (o *CreateChannelInviteRequest) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *CreateChannelInviteRequest) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *CreateChannelInviteRequest) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *CreateChannelInviteRequest) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetTargetApplicationId

`func (o *CreateChannelInviteRequest) GetTargetApplicationId() string`

GetTargetApplicationId returns the TargetApplicationId field if non-nil, zero value otherwise.

### GetTargetApplicationIdOk

`func (o *CreateChannelInviteRequest) GetTargetApplicationIdOk() (*string, bool)`

GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationId

`func (o *CreateChannelInviteRequest) SetTargetApplicationId(v string)`

SetTargetApplicationId sets TargetApplicationId field to given value.

### HasTargetApplicationId

`func (o *CreateChannelInviteRequest) HasTargetApplicationId() bool`

HasTargetApplicationId returns a boolean if a field has been set.

### GetTargetType

`func (o *CreateChannelInviteRequest) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CreateChannelInviteRequest) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CreateChannelInviteRequest) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CreateChannelInviteRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *CreateChannelInviteRequest) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *CreateChannelInviteRequest) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


