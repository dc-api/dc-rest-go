# ScheduledEventResponse

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

## Methods

### NewScheduledEventResponse

`func NewScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status NullableInt32, entityType NullableInt32, privacyLevel interface{}, ) *ScheduledEventResponse`

NewScheduledEventResponse instantiates a new ScheduledEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledEventResponseWithDefaults

`func NewScheduledEventResponseWithDefaults() *ScheduledEventResponse`

NewScheduledEventResponseWithDefaults instantiates a new ScheduledEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *ScheduledEventResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ScheduledEventResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ScheduledEventResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *ScheduledEventResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledEventResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledEventResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScheduledEventResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledEventResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledEventResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledEventResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduledEventResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduledEventResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChannelId

`func (o *ScheduledEventResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ScheduledEventResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ScheduledEventResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ScheduledEventResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *ScheduledEventResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ScheduledEventResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ScheduledEventResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ScheduledEventResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreator

`func (o *ScheduledEventResponse) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ScheduledEventResponse) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ScheduledEventResponse) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ScheduledEventResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *ScheduledEventResponse) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *ScheduledEventResponse) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetImage

`func (o *ScheduledEventResponse) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ScheduledEventResponse) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ScheduledEventResponse) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ScheduledEventResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ScheduledEventResponse) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ScheduledEventResponse) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *ScheduledEventResponse) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ScheduledEventResponse) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *ScheduledEventResponse) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *ScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *ScheduledEventResponse) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *ScheduledEventResponse) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *ScheduledEventResponse) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *ScheduledEventResponse) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetStatus

`func (o *ScheduledEventResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledEventResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledEventResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ScheduledEventResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ScheduledEventResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEntityType

`func (o *ScheduledEventResponse) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ScheduledEventResponse) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ScheduledEventResponse) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *ScheduledEventResponse) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ScheduledEventResponse) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetEntityId

`func (o *ScheduledEventResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ScheduledEventResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ScheduledEventResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ScheduledEventResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetUserCount

`func (o *ScheduledEventResponse) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ScheduledEventResponse) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ScheduledEventResponse) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ScheduledEventResponse) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCountNil

`func (o *ScheduledEventResponse) SetUserCountNil(b bool)`

 SetUserCountNil sets the value for UserCount to be an explicit nil

### UnsetUserCount
`func (o *ScheduledEventResponse) UnsetUserCount()`

UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
### GetPrivacyLevel

`func (o *ScheduledEventResponse) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *ScheduledEventResponse) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *ScheduledEventResponse) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *ScheduledEventResponse) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *ScheduledEventResponse) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetUserRsvp

`func (o *ScheduledEventResponse) GetUserRsvp() ScheduledEventUserResponse`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *ScheduledEventResponse) GetUserRsvpOk() (*ScheduledEventUserResponse, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *ScheduledEventResponse) SetUserRsvp(v ScheduledEventUserResponse)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *ScheduledEventResponse) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

### SetUserRsvpNil

`func (o *ScheduledEventResponse) SetUserRsvpNil(b bool)`

 SetUserRsvpNil sets the value for UserRsvp to be an explicit nil

### UnsetUserRsvp
`func (o *ScheduledEventResponse) UnsetUserRsvp()`

UnsetUserRsvp ensures that no value is present for UserRsvp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


