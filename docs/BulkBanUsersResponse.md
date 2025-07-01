# BulkBanUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannedUsers** | **[]string** |  | 
**FailedUsers** | **[]string** |  | 

## Methods

### NewBulkBanUsersResponse

`func NewBulkBanUsersResponse(bannedUsers []string, failedUsers []string, ) *BulkBanUsersResponse`

NewBulkBanUsersResponse instantiates a new BulkBanUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkBanUsersResponseWithDefaults

`func NewBulkBanUsersResponseWithDefaults() *BulkBanUsersResponse`

NewBulkBanUsersResponseWithDefaults instantiates a new BulkBanUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannedUsers

`func (o *BulkBanUsersResponse) GetBannedUsers() []string`

GetBannedUsers returns the BannedUsers field if non-nil, zero value otherwise.

### GetBannedUsersOk

`func (o *BulkBanUsersResponse) GetBannedUsersOk() (*[]string, bool)`

GetBannedUsersOk returns a tuple with the BannedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUsers

`func (o *BulkBanUsersResponse) SetBannedUsers(v []string)`

SetBannedUsers sets BannedUsers field to given value.


### GetFailedUsers

`func (o *BulkBanUsersResponse) GetFailedUsers() []string`

GetFailedUsers returns the FailedUsers field if non-nil, zero value otherwise.

### GetFailedUsersOk

`func (o *BulkBanUsersResponse) GetFailedUsersOk() (*[]string, bool)`

GetFailedUsersOk returns a tuple with the FailedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUsers

`func (o *BulkBanUsersResponse) SetFailedUsers(v []string)`

SetFailedUsers sets FailedUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


