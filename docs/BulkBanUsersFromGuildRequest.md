# BulkBanUsersFromGuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** |  | 
**DeleteMessageSeconds** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBulkBanUsersFromGuildRequest

`func NewBulkBanUsersFromGuildRequest(userIds []string, ) *BulkBanUsersFromGuildRequest`

NewBulkBanUsersFromGuildRequest instantiates a new BulkBanUsersFromGuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkBanUsersFromGuildRequestWithDefaults

`func NewBulkBanUsersFromGuildRequestWithDefaults() *BulkBanUsersFromGuildRequest`

NewBulkBanUsersFromGuildRequestWithDefaults instantiates a new BulkBanUsersFromGuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *BulkBanUsersFromGuildRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkBanUsersFromGuildRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkBanUsersFromGuildRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetDeleteMessageSeconds

`func (o *BulkBanUsersFromGuildRequest) GetDeleteMessageSeconds() int32`

GetDeleteMessageSeconds returns the DeleteMessageSeconds field if non-nil, zero value otherwise.

### GetDeleteMessageSecondsOk

`func (o *BulkBanUsersFromGuildRequest) GetDeleteMessageSecondsOk() (*int32, bool)`

GetDeleteMessageSecondsOk returns a tuple with the DeleteMessageSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessageSeconds

`func (o *BulkBanUsersFromGuildRequest) SetDeleteMessageSeconds(v int32)`

SetDeleteMessageSeconds sets DeleteMessageSeconds field to given value.

### HasDeleteMessageSeconds

`func (o *BulkBanUsersFromGuildRequest) HasDeleteMessageSeconds() bool`

HasDeleteMessageSeconds returns a boolean if a field has been set.

### SetDeleteMessageSecondsNil

`func (o *BulkBanUsersFromGuildRequest) SetDeleteMessageSecondsNil(b bool)`

 SetDeleteMessageSecondsNil sets the value for DeleteMessageSeconds to be an explicit nil

### UnsetDeleteMessageSeconds
`func (o *BulkBanUsersFromGuildRequest) UnsetDeleteMessageSeconds()`

UnsetDeleteMessageSeconds ensures that no value is present for DeleteMessageSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


