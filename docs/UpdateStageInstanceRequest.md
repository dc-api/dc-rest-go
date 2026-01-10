# UpdateStageInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**PrivacyLevel** | Pointer to [**StageInstancesPrivacyLevels**](StageInstancesPrivacyLevels.md) |  | [optional] 

## Methods

### NewUpdateStageInstanceRequest

`func NewUpdateStageInstanceRequest() *UpdateStageInstanceRequest`

NewUpdateStageInstanceRequest instantiates a new UpdateStageInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStageInstanceRequestWithDefaults

`func NewUpdateStageInstanceRequestWithDefaults() *UpdateStageInstanceRequest`

NewUpdateStageInstanceRequestWithDefaults instantiates a new UpdateStageInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *UpdateStageInstanceRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *UpdateStageInstanceRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *UpdateStageInstanceRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *UpdateStageInstanceRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *UpdateStageInstanceRequest) GetPrivacyLevel() StageInstancesPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *UpdateStageInstanceRequest) GetPrivacyLevelOk() (*StageInstancesPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *UpdateStageInstanceRequest) SetPrivacyLevel(v StageInstancesPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *UpdateStageInstanceRequest) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


