# VoiceScheduledEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | **NullableInt32** |  | 
**EntityType** | **NullableInt32** |  | 
**EntityId** | Pointer to **string** |  | [optional] 
**UserCount** | Pointer to **NullableInt32** |  | [optional] 
**PrivacyLevel** | **interface{}** |  | 
**UserRsvp** | Pointer to [**NullableScheduledEventUserResponse**](ScheduledEventUserResponse.md) |  | [optional] 
**EntityMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVoiceScheduledEventResponse

`func NewVoiceScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status NullableInt32, entityType NullableInt32, privacyLevel interface{}, ) *VoiceScheduledEventResponse`

NewVoiceScheduledEventResponse instantiates a new VoiceScheduledEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceScheduledEventResponseWithDefaults

`func NewVoiceScheduledEventResponseWithDefaults() *VoiceScheduledEventResponse`

NewVoiceScheduledEventResponseWithDefaults instantiates a new VoiceScheduledEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoiceScheduledEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoiceScheduledEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoiceScheduledEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *VoiceScheduledEventResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *VoiceScheduledEventResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *VoiceScheduledEventResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *VoiceScheduledEventResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceScheduledEventResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceScheduledEventResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VoiceScheduledEventResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoiceScheduledEventResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoiceScheduledEventResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoiceScheduledEventResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VoiceScheduledEventResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VoiceScheduledEventResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChannelId

`func (o *VoiceScheduledEventResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VoiceScheduledEventResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VoiceScheduledEventResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *VoiceScheduledEventResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *VoiceScheduledEventResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *VoiceScheduledEventResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *VoiceScheduledEventResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *VoiceScheduledEventResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreator

`func (o *VoiceScheduledEventResponse) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *VoiceScheduledEventResponse) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *VoiceScheduledEventResponse) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *VoiceScheduledEventResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *VoiceScheduledEventResponse) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *VoiceScheduledEventResponse) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetImage

`func (o *VoiceScheduledEventResponse) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VoiceScheduledEventResponse) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VoiceScheduledEventResponse) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VoiceScheduledEventResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *VoiceScheduledEventResponse) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *VoiceScheduledEventResponse) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *VoiceScheduledEventResponse) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *VoiceScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *VoiceScheduledEventResponse) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *VoiceScheduledEventResponse) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *VoiceScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *VoiceScheduledEventResponse) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *VoiceScheduledEventResponse) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *VoiceScheduledEventResponse) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *VoiceScheduledEventResponse) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetStatus

`func (o *VoiceScheduledEventResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoiceScheduledEventResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoiceScheduledEventResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *VoiceScheduledEventResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VoiceScheduledEventResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEntityType

`func (o *VoiceScheduledEventResponse) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *VoiceScheduledEventResponse) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *VoiceScheduledEventResponse) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *VoiceScheduledEventResponse) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *VoiceScheduledEventResponse) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetEntityId

`func (o *VoiceScheduledEventResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *VoiceScheduledEventResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *VoiceScheduledEventResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *VoiceScheduledEventResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetUserCount

`func (o *VoiceScheduledEventResponse) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *VoiceScheduledEventResponse) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *VoiceScheduledEventResponse) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *VoiceScheduledEventResponse) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCountNil

`func (o *VoiceScheduledEventResponse) SetUserCountNil(b bool)`

 SetUserCountNil sets the value for UserCount to be an explicit nil

### UnsetUserCount
`func (o *VoiceScheduledEventResponse) UnsetUserCount()`

UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
### GetPrivacyLevel

`func (o *VoiceScheduledEventResponse) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *VoiceScheduledEventResponse) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *VoiceScheduledEventResponse) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *VoiceScheduledEventResponse) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *VoiceScheduledEventResponse) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetUserRsvp

`func (o *VoiceScheduledEventResponse) GetUserRsvp() ScheduledEventUserResponse`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *VoiceScheduledEventResponse) GetUserRsvpOk() (*ScheduledEventUserResponse, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *VoiceScheduledEventResponse) SetUserRsvp(v ScheduledEventUserResponse)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *VoiceScheduledEventResponse) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

### SetUserRsvpNil

`func (o *VoiceScheduledEventResponse) SetUserRsvpNil(b bool)`

 SetUserRsvpNil sets the value for UserRsvp to be an explicit nil

### UnsetUserRsvp
`func (o *VoiceScheduledEventResponse) UnsetUserRsvp()`

UnsetUserRsvp ensures that no value is present for UserRsvp, not even an explicit nil
### GetEntityMetadata

`func (o *VoiceScheduledEventResponse) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *VoiceScheduledEventResponse) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *VoiceScheduledEventResponse) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *VoiceScheduledEventResponse) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


