# EmbeddedActivityInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**InstanceId** | **string** |  | 
**LaunchId** | **string** |  | 
**Location** | Pointer to [**NullableEmbeddedActivityInstanceLocation**](EmbeddedActivityInstanceLocation.md) |  | [optional] 
**Users** | **[]string** |  | 

## Methods

### NewEmbeddedActivityInstance

`func NewEmbeddedActivityInstance(applicationId string, instanceId string, launchId string, users []string, ) *EmbeddedActivityInstance`

NewEmbeddedActivityInstance instantiates a new EmbeddedActivityInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedActivityInstanceWithDefaults

`func NewEmbeddedActivityInstanceWithDefaults() *EmbeddedActivityInstance`

NewEmbeddedActivityInstanceWithDefaults instantiates a new EmbeddedActivityInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *EmbeddedActivityInstance) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *EmbeddedActivityInstance) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *EmbeddedActivityInstance) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetInstanceId

`func (o *EmbeddedActivityInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *EmbeddedActivityInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *EmbeddedActivityInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetLaunchId

`func (o *EmbeddedActivityInstance) GetLaunchId() string`

GetLaunchId returns the LaunchId field if non-nil, zero value otherwise.

### GetLaunchIdOk

`func (o *EmbeddedActivityInstance) GetLaunchIdOk() (*string, bool)`

GetLaunchIdOk returns a tuple with the LaunchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchId

`func (o *EmbeddedActivityInstance) SetLaunchId(v string)`

SetLaunchId sets LaunchId field to given value.


### GetLocation

`func (o *EmbeddedActivityInstance) GetLocation() EmbeddedActivityInstanceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmbeddedActivityInstance) GetLocationOk() (*EmbeddedActivityInstanceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmbeddedActivityInstance) SetLocation(v EmbeddedActivityInstanceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmbeddedActivityInstance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *EmbeddedActivityInstance) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *EmbeddedActivityInstance) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetUsers

`func (o *EmbeddedActivityInstance) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EmbeddedActivityInstance) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EmbeddedActivityInstance) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


