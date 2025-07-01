# ExternalScheduledEventResponse

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
**EntityMetadata** | [**EntityMetadataExternalResponse**](EntityMetadataExternalResponse.md) |  | 

## Methods

### NewExternalScheduledEventResponse

`func NewExternalScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status NullableInt32, entityType NullableInt32, privacyLevel interface{}, entityMetadata EntityMetadataExternalResponse, ) *ExternalScheduledEventResponse`

NewExternalScheduledEventResponse instantiates a new ExternalScheduledEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalScheduledEventResponseWithDefaults

`func NewExternalScheduledEventResponseWithDefaults() *ExternalScheduledEventResponse`

NewExternalScheduledEventResponseWithDefaults instantiates a new ExternalScheduledEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalScheduledEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalScheduledEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalScheduledEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *ExternalScheduledEventResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ExternalScheduledEventResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ExternalScheduledEventResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *ExternalScheduledEventResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalScheduledEventResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalScheduledEventResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExternalScheduledEventResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalScheduledEventResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalScheduledEventResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalScheduledEventResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalScheduledEventResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalScheduledEventResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChannelId

`func (o *ExternalScheduledEventResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ExternalScheduledEventResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ExternalScheduledEventResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ExternalScheduledEventResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *ExternalScheduledEventResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ExternalScheduledEventResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ExternalScheduledEventResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ExternalScheduledEventResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreator

`func (o *ExternalScheduledEventResponse) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ExternalScheduledEventResponse) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ExternalScheduledEventResponse) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ExternalScheduledEventResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *ExternalScheduledEventResponse) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *ExternalScheduledEventResponse) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetImage

`func (o *ExternalScheduledEventResponse) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ExternalScheduledEventResponse) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ExternalScheduledEventResponse) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ExternalScheduledEventResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ExternalScheduledEventResponse) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ExternalScheduledEventResponse) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *ExternalScheduledEventResponse) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ExternalScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ExternalScheduledEventResponse) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *ExternalScheduledEventResponse) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *ExternalScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *ExternalScheduledEventResponse) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *ExternalScheduledEventResponse) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *ExternalScheduledEventResponse) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *ExternalScheduledEventResponse) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetStatus

`func (o *ExternalScheduledEventResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalScheduledEventResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalScheduledEventResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ExternalScheduledEventResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExternalScheduledEventResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEntityType

`func (o *ExternalScheduledEventResponse) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ExternalScheduledEventResponse) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ExternalScheduledEventResponse) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *ExternalScheduledEventResponse) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ExternalScheduledEventResponse) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetEntityId

`func (o *ExternalScheduledEventResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ExternalScheduledEventResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ExternalScheduledEventResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ExternalScheduledEventResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetUserCount

`func (o *ExternalScheduledEventResponse) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ExternalScheduledEventResponse) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ExternalScheduledEventResponse) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ExternalScheduledEventResponse) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCountNil

`func (o *ExternalScheduledEventResponse) SetUserCountNil(b bool)`

 SetUserCountNil sets the value for UserCount to be an explicit nil

### UnsetUserCount
`func (o *ExternalScheduledEventResponse) UnsetUserCount()`

UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
### GetPrivacyLevel

`func (o *ExternalScheduledEventResponse) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *ExternalScheduledEventResponse) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *ExternalScheduledEventResponse) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *ExternalScheduledEventResponse) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *ExternalScheduledEventResponse) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetUserRsvp

`func (o *ExternalScheduledEventResponse) GetUserRsvp() ScheduledEventUserResponse`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *ExternalScheduledEventResponse) GetUserRsvpOk() (*ScheduledEventUserResponse, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *ExternalScheduledEventResponse) SetUserRsvp(v ScheduledEventUserResponse)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *ExternalScheduledEventResponse) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

### SetUserRsvpNil

`func (o *ExternalScheduledEventResponse) SetUserRsvpNil(b bool)`

 SetUserRsvpNil sets the value for UserRsvp to be an explicit nil

### UnsetUserRsvp
`func (o *ExternalScheduledEventResponse) UnsetUserRsvp()`

UnsetUserRsvp ensures that no value is present for UserRsvp, not even an explicit nil
### GetEntityMetadata

`func (o *ExternalScheduledEventResponse) GetEntityMetadata() EntityMetadataExternalResponse`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *ExternalScheduledEventResponse) GetEntityMetadataOk() (*EntityMetadataExternalResponse, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *ExternalScheduledEventResponse) SetEntityMetadata(v EntityMetadataExternalResponse)`

SetEntityMetadata sets EntityMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


