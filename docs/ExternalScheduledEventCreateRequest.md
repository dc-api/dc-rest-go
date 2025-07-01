# ExternalScheduledEventCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **NullableTime** |  | [optional] 
**PrivacyLevel** | **interface{}** |  | 
**EntityType** | **NullableInt32** |  | 
**ChannelId** | Pointer to **string** |  | [optional] 
**EntityMetadata** | [**EntityMetadataExternal**](EntityMetadataExternal.md) |  | 

## Methods

### NewExternalScheduledEventCreateRequest

`func NewExternalScheduledEventCreateRequest(name string, scheduledStartTime time.Time, privacyLevel interface{}, entityType NullableInt32, entityMetadata EntityMetadataExternal, ) *ExternalScheduledEventCreateRequest`

NewExternalScheduledEventCreateRequest instantiates a new ExternalScheduledEventCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalScheduledEventCreateRequestWithDefaults

`func NewExternalScheduledEventCreateRequestWithDefaults() *ExternalScheduledEventCreateRequest`

NewExternalScheduledEventCreateRequestWithDefaults instantiates a new ExternalScheduledEventCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalScheduledEventCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalScheduledEventCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalScheduledEventCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExternalScheduledEventCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalScheduledEventCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalScheduledEventCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalScheduledEventCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalScheduledEventCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalScheduledEventCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImage

`func (o *ExternalScheduledEventCreateRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ExternalScheduledEventCreateRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ExternalScheduledEventCreateRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ExternalScheduledEventCreateRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ExternalScheduledEventCreateRequest) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ExternalScheduledEventCreateRequest) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *ExternalScheduledEventCreateRequest) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ExternalScheduledEventCreateRequest) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ExternalScheduledEventCreateRequest) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *ExternalScheduledEventCreateRequest) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *ExternalScheduledEventCreateRequest) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *ExternalScheduledEventCreateRequest) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *ExternalScheduledEventCreateRequest) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *ExternalScheduledEventCreateRequest) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *ExternalScheduledEventCreateRequest) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetPrivacyLevel

`func (o *ExternalScheduledEventCreateRequest) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *ExternalScheduledEventCreateRequest) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *ExternalScheduledEventCreateRequest) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *ExternalScheduledEventCreateRequest) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *ExternalScheduledEventCreateRequest) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetEntityType

`func (o *ExternalScheduledEventCreateRequest) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ExternalScheduledEventCreateRequest) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ExternalScheduledEventCreateRequest) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *ExternalScheduledEventCreateRequest) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ExternalScheduledEventCreateRequest) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetChannelId

`func (o *ExternalScheduledEventCreateRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ExternalScheduledEventCreateRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ExternalScheduledEventCreateRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ExternalScheduledEventCreateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *ExternalScheduledEventCreateRequest) GetEntityMetadata() EntityMetadataExternal`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *ExternalScheduledEventCreateRequest) GetEntityMetadataOk() (*EntityMetadataExternal, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *ExternalScheduledEventCreateRequest) SetEntityMetadata(v EntityMetadataExternal)`

SetEntityMetadata sets EntityMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


