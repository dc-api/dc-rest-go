# BulkBanUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** |  | 
**DeleteMessageSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewBulkBanUsersRequest

`func NewBulkBanUsersRequest(userIds []string, ) *BulkBanUsersRequest`

NewBulkBanUsersRequest instantiates a new BulkBanUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkBanUsersRequestWithDefaults

`func NewBulkBanUsersRequestWithDefaults() *BulkBanUsersRequest`

NewBulkBanUsersRequestWithDefaults instantiates a new BulkBanUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *BulkBanUsersRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkBanUsersRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkBanUsersRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetDeleteMessageSeconds

`func (o *BulkBanUsersRequest) GetDeleteMessageSeconds() int32`

GetDeleteMessageSeconds returns the DeleteMessageSeconds field if non-nil, zero value otherwise.

### GetDeleteMessageSecondsOk

`func (o *BulkBanUsersRequest) GetDeleteMessageSecondsOk() (*int32, bool)`

GetDeleteMessageSecondsOk returns a tuple with the DeleteMessageSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMessageSeconds

`func (o *BulkBanUsersRequest) SetDeleteMessageSeconds(v int32)`

SetDeleteMessageSeconds sets DeleteMessageSeconds field to given value.

### HasDeleteMessageSeconds

`func (o *BulkBanUsersRequest) HasDeleteMessageSeconds() bool`

HasDeleteMessageSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


