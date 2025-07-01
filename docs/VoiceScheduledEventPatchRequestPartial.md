# VoiceScheduledEventPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** |  | [optional] 
**ScheduledEndTime** | Pointer to **NullableTime** |  | [optional] 
**EntityType** | Pointer to **NullableInt32** |  | [optional] 
**PrivacyLevel** | Pointer to **interface{}** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**EntityMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVoiceScheduledEventPatchRequestPartial

`func NewVoiceScheduledEventPatchRequestPartial() *VoiceScheduledEventPatchRequestPartial`

NewVoiceScheduledEventPatchRequestPartial instantiates a new VoiceScheduledEventPatchRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceScheduledEventPatchRequestPartialWithDefaults

`func NewVoiceScheduledEventPatchRequestPartialWithDefaults() *VoiceScheduledEventPatchRequestPartial`

NewVoiceScheduledEventPatchRequestPartialWithDefaults instantiates a new VoiceScheduledEventPatchRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VoiceScheduledEventPatchRequestPartial) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoiceScheduledEventPatchRequestPartial) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoiceScheduledEventPatchRequestPartial) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetName

`func (o *VoiceScheduledEventPatchRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceScheduledEventPatchRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoiceScheduledEventPatchRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VoiceScheduledEventPatchRequestPartial) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoiceScheduledEventPatchRequestPartial) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoiceScheduledEventPatchRequestPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImage

`func (o *VoiceScheduledEventPatchRequestPartial) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VoiceScheduledEventPatchRequestPartial) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VoiceScheduledEventPatchRequestPartial) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *VoiceScheduledEventPatchRequestPartial) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetScheduledEndTime

`func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *VoiceScheduledEventPatchRequestPartial) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetEntityType

`func (o *VoiceScheduledEventPatchRequestPartial) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *VoiceScheduledEventPatchRequestPartial) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *VoiceScheduledEventPatchRequestPartial) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetPrivacyLevel

`func (o *VoiceScheduledEventPatchRequestPartial) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *VoiceScheduledEventPatchRequestPartial) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *VoiceScheduledEventPatchRequestPartial) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### SetPrivacyLevelNil

`func (o *VoiceScheduledEventPatchRequestPartial) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *VoiceScheduledEventPatchRequestPartial) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetChannelId

`func (o *VoiceScheduledEventPatchRequestPartial) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VoiceScheduledEventPatchRequestPartial) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *VoiceScheduledEventPatchRequestPartial) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *VoiceScheduledEventPatchRequestPartial) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *VoiceScheduledEventPatchRequestPartial) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *VoiceScheduledEventPatchRequestPartial) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *VoiceScheduledEventPatchRequestPartial) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


