# BanUserFromGuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteMessageSeconds** | Pointer to **int32** |  | [optional] 
**DeleteMessageDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewBanUserFromGuildRequest

`func NewBanUserFromGuildRequest() *BanUserFromGuildRequest`

NewBanUserFromGuildRequest instantiates a new BanUserFromGuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanUserFromGuildRequestWithDefaults

`func NewBanUserFromGuildRequestWithDefaults() *BanUserFromGuildRequest`

NewBanUserFromGuildRequestWithDefaults instantiates a new BanUserFromGuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteMessageSeconds

`func (o *BanUserFromGuildRequest) GetDeleteMessageSeconds() int32`

GetDeleteMessageSeconds returns the DeleteMessageSeconds field if non-nil, zero value otherwise.

### GetDeleteMessageSecondsOk

`func (o *BanUserFromGuildRequest) GetDeleteMessageSecondsOk() (*int32, bool)`

GetDeleteMessageSecondsOk returns a tuple with the DeleteMessageSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessageSeconds

`func (o *BanUserFromGuildRequest) SetDeleteMessageSeconds(v int32)`

SetDeleteMessageSeconds sets DeleteMessageSeconds field to given value.

### HasDeleteMessageSeconds

`func (o *BanUserFromGuildRequest) HasDeleteMessageSeconds() bool`

HasDeleteMessageSeconds returns a boolean if a field has been set.

### GetDeleteMessageDays

`func (o *BanUserFromGuildRequest) GetDeleteMessageDays() int32`

GetDeleteMessageDays returns the DeleteMessageDays field if non-nil, zero value otherwise.

### GetDeleteMessageDaysOk

`func (o *BanUserFromGuildRequest) GetDeleteMessageDaysOk() (*int32, bool)`

GetDeleteMessageDaysOk returns a tuple with the DeleteMessageDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessageDays

`func (o *BanUserFromGuildRequest) SetDeleteMessageDays(v int32)`

SetDeleteMessageDays sets DeleteMessageDays field to given value.

### HasDeleteMessageDays

`func (o *BanUserFromGuildRequest) HasDeleteMessageDays() bool`

HasDeleteMessageDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


