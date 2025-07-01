# ListGuildScheduledEvents200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **time.Time** |  | [optional] 
**Status** | **int32** |  | 
**EntityType** | **NullableInt32** |  | 
**EntityId** | Pointer to **string** |  | [optional] 
**UserCount** | Pointer to **int32** |  | [optional] 
**PrivacyLevel** | **interface{}** |  | 
**UserRsvp** | Pointer to [**ScheduledEventUserResponse**](ScheduledEventUserResponse.md) |  | [optional] 
**EntityMetadata** | **map[string]interface{}** |  | 

## Methods

### NewListGuildScheduledEvents200ResponseInner

`func NewListGuildScheduledEvents200ResponseInner(id string, guildId string, name string, scheduledStartTime time.Time, status int32, entityType NullableInt32, privacyLevel interface{}, entityMetadata map[string]interface{}, ) *ListGuildScheduledEvents200ResponseInner`

NewListGuildScheduledEvents200ResponseInner instantiates a new ListGuildScheduledEvents200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuildScheduledEvents200ResponseInnerWithDefaults

`func NewListGuildScheduledEvents200ResponseInnerWithDefaults() *ListGuildScheduledEvents200ResponseInner`

NewListGuildScheduledEvents200ResponseInnerWithDefaults instantiates a new ListGuildScheduledEvents200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuildScheduledEvents200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuildScheduledEvents200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *ListGuildScheduledEvents200ResponseInner) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ListGuildScheduledEvents200ResponseInner) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *ListGuildScheduledEvents200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuildScheduledEvents200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListGuildScheduledEvents200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListGuildScheduledEvents200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListGuildScheduledEvents200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChannelId

`func (o *ListGuildScheduledEvents200ResponseInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ListGuildScheduledEvents200ResponseInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ListGuildScheduledEvents200ResponseInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *ListGuildScheduledEvents200ResponseInner) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ListGuildScheduledEvents200ResponseInner) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ListGuildScheduledEvents200ResponseInner) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreator

`func (o *ListGuildScheduledEvents200ResponseInner) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ListGuildScheduledEvents200ResponseInner) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ListGuildScheduledEvents200ResponseInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetImage

`func (o *ListGuildScheduledEvents200ResponseInner) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ListGuildScheduledEvents200ResponseInner) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ListGuildScheduledEvents200ResponseInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *ListGuildScheduledEvents200ResponseInner) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ListGuildScheduledEvents200ResponseInner) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *ListGuildScheduledEvents200ResponseInner) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *ListGuildScheduledEvents200ResponseInner) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *ListGuildScheduledEvents200ResponseInner) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### GetStatus

`func (o *ListGuildScheduledEvents200ResponseInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListGuildScheduledEvents200ResponseInner) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetEntityType

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ListGuildScheduledEvents200ResponseInner) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### SetEntityTypeNil

`func (o *ListGuildScheduledEvents200ResponseInner) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ListGuildScheduledEvents200ResponseInner) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetEntityId

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ListGuildScheduledEvents200ResponseInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ListGuildScheduledEvents200ResponseInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetUserCount

`func (o *ListGuildScheduledEvents200ResponseInner) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ListGuildScheduledEvents200ResponseInner) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ListGuildScheduledEvents200ResponseInner) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *ListGuildScheduledEvents200ResponseInner) GetPrivacyLevel() interface{}`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetPrivacyLevelOk() (*interface{}, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *ListGuildScheduledEvents200ResponseInner) SetPrivacyLevel(v interface{})`

SetPrivacyLevel sets PrivacyLevel field to given value.


### SetPrivacyLevelNil

`func (o *ListGuildScheduledEvents200ResponseInner) SetPrivacyLevelNil(b bool)`

 SetPrivacyLevelNil sets the value for PrivacyLevel to be an explicit nil

### UnsetPrivacyLevel
`func (o *ListGuildScheduledEvents200ResponseInner) UnsetPrivacyLevel()`

UnsetPrivacyLevel ensures that no value is present for PrivacyLevel, not even an explicit nil
### GetUserRsvp

`func (o *ListGuildScheduledEvents200ResponseInner) GetUserRsvp() ScheduledEventUserResponse`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetUserRsvpOk() (*ScheduledEventUserResponse, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *ListGuildScheduledEvents200ResponseInner) SetUserRsvp(v ScheduledEventUserResponse)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *ListGuildScheduledEvents200ResponseInner) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *ListGuildScheduledEvents200ResponseInner) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *ListGuildScheduledEvents200ResponseInner) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


