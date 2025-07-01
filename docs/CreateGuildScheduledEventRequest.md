# CreateGuildScheduledEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **time.Time** |  | [optional] 
**PrivacyLevel** | **interface{}** |  | 
**EntityType** | **NullableInt32** |  | 
**ChannelId** | Pointer to **string** |  | [optional] 
**EntityMetadata** | **map[string]interface{}** |  | 

## Methods

### NewCreateGuildScheduledEventRequest

`func NewCreateGuildScheduledEventRequest(name string, scheduledStartTime time.Time, privacyLevel interface{}, entityType NullableInt32, entityMetadata map[string]interface{}, ) *CreateGuildScheduledEventRequest`

NewCreateGuildScheduledEventRequest instantiates a new CreateGuildScheduledEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildScheduledEventRequestWithDefaults

`func NewCreateGuildScheduledEventRequestWithDefaults() *CreateGuildScheduledEventRequest`

NewCreateGuildScheduledEventRequestWithDefaults instantiates a new CreateGuildScheduledEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGuildScheduledEventRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildScheduledEventRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildScheduledEventRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGuildScheduledEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGuildScheduledEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGuildScheduledEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGuildScheduledEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *CreateGuildScheduledEventRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateGuildScheduledEventRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateGuildScheduledEventRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateGuildScheduledEventRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *CreateGuildScheduledEventRequest) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *CreateGuildScheduledEventRequest) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *CreateGuildScheduledEventRequest) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *CreateGuildScheduledEventRequest) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *CreateGuildScheduledEventRequest) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *CreateGuildScheduledEventRequest) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *CreateGuildScheduledEventRequest) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *CreateGuildScheduledEventRequest) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *CreateGuildScheduledEventRequest) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *CreateGuildScheduledEventRequest) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *CreateGuildScheduledEventRequest) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *CreateGuildScheduledEventRequest) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetEntityType

`func (o *CreateGuildScheduledEventRequest) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateGuildScheduledEventRequest) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateGuildScheduledEventRequest) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *CreateGuildScheduledEventRequest) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *CreateGuildScheduledEventRequest) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetChannelId

`func (o *CreateGuildScheduledEventRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateGuildScheduledEventRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateGuildScheduledEventRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *CreateGuildScheduledEventRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *CreateGuildScheduledEventRequest) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *CreateGuildScheduledEventRequest) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *CreateGuildScheduledEventRequest) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


