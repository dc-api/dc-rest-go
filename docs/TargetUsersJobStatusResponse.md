# TargetUsersJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TargetUsersJobStatusTypes**](TargetUsersJobStatusTypes.md) | The status of the job processing the target users. | 
**TotalUsers** | **int64** |  | 
**ProcessedUsers** | **int64** |  | 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the job was created. | [optional] 
**CompletedAt** | Pointer to **time.Time** | The timestamp when the job was successfully completed. | [optional] 
**ErrorMessage** | Pointer to **string** | The error message if the job failed. | [optional] 

## Methods

### NewTargetUsersJobStatusResponse

`func NewTargetUsersJobStatusResponse(status TargetUsersJobStatusTypes, totalUsers int64, processedUsers int64, ) *TargetUsersJobStatusResponse`

NewTargetUsersJobStatusResponse instantiates a new TargetUsersJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUsersJobStatusResponseWithDefaults

`func NewTargetUsersJobStatusResponseWithDefaults() *TargetUsersJobStatusResponse`

NewTargetUsersJobStatusResponseWithDefaults instantiates a new TargetUsersJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TargetUsersJobStatusResponse) GetStatus() TargetUsersJobStatusTypes`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetUsersJobStatusResponse) GetStatusOk() (*TargetUsersJobStatusTypes, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetUsersJobStatusResponse) SetStatus(v TargetUsersJobStatusTypes)`

SetStatus sets Status field to given value.


### GetTotalUsers

`func (o *TargetUsersJobStatusResponse) GetTotalUsers() int64`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *TargetUsersJobStatusResponse) GetTotalUsersOk() (*int64, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *TargetUsersJobStatusResponse) SetTotalUsers(v int64)`

SetTotalUsers sets TotalUsers field to given value.


### GetProcessedUsers

`func (o *TargetUsersJobStatusResponse) GetProcessedUsers() int64`

GetProcessedUsers returns the ProcessedUsers field if non-nil, zero value otherwise.

### GetProcessedUsersOk

`func (o *TargetUsersJobStatusResponse) GetProcessedUsersOk() (*int64, bool)`

GetProcessedUsersOk returns a tuple with the ProcessedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedUsers

`func (o *TargetUsersJobStatusResponse) SetProcessedUsers(v int64)`

SetProcessedUsers sets ProcessedUsers field to given value.


### GetCreatedAt

`func (o *TargetUsersJobStatusResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TargetUsersJobStatusResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TargetUsersJobStatusResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TargetUsersJobStatusResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TargetUsersJobStatusResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TargetUsersJobStatusResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TargetUsersJobStatusResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TargetUsersJobStatusResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TargetUsersJobStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TargetUsersJobStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TargetUsersJobStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TargetUsersJobStatusResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


