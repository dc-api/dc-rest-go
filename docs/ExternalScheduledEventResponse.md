# ExternalScheduledEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**CreatorId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Creator** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **time.Time** |  | [optional] 
**Status** | [**GuildScheduledEventStatuses**](GuildScheduledEventStatuses.md) |  | 
**EntityType** | [**GuildScheduledEventEntityTypes**](GuildScheduledEventEntityTypes.md) |  | 
**EntityId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**UserCount** | Pointer to **int32** |  | [optional] 
**PrivacyLevel** | [**GuildScheduledEventPrivacyLevels**](GuildScheduledEventPrivacyLevels.md) |  | 
**UserRsvp** | Pointer to [**ExternalScheduledEventResponseUserRsvp**](ExternalScheduledEventResponseUserRsvp.md) |  | [optional] 
**EntityMetadata** | [**EntityMetadataExternalResponse**](EntityMetadataExternalResponse.md) |  | 

## Methods

### NewExternalScheduledEventResponse

`func NewExternalScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status GuildScheduledEventStatuses, entityType GuildScheduledEventEntityTypes, privacyLevel GuildScheduledEventPrivacyLevels, entityMetadata EntityMetadataExternalResponse, ) *ExternalScheduledEventResponse`

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

### GetChannelId

`func (o *ExternalScheduledEventResponse) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ExternalScheduledEventResponse) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ExternalScheduledEventResponse) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ExternalScheduledEventResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *ExternalScheduledEventResponse) GetCreatorId() GetEntitlementsSkuIdsParameterOneOfInner`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ExternalScheduledEventResponse) GetCreatorIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ExternalScheduledEventResponse) SetCreatorId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

### GetStatus

`func (o *ExternalScheduledEventResponse) GetStatus() GuildScheduledEventStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalScheduledEventResponse) GetStatusOk() (*GuildScheduledEventStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalScheduledEventResponse) SetStatus(v GuildScheduledEventStatuses)`

SetStatus sets Status field to given value.


### GetEntityType

`func (o *ExternalScheduledEventResponse) GetEntityType() GuildScheduledEventEntityTypes`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ExternalScheduledEventResponse) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ExternalScheduledEventResponse) SetEntityType(v GuildScheduledEventEntityTypes)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *ExternalScheduledEventResponse) GetEntityId() GetEntitlementsSkuIdsParameterOneOfInner`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ExternalScheduledEventResponse) GetEntityIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ExternalScheduledEventResponse) SetEntityId(v GetEntitlementsSkuIdsParameterOneOfInner)`

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

### GetPrivacyLevel

`func (o *ExternalScheduledEventResponse) GetPrivacyLevel() GuildScheduledEventPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *ExternalScheduledEventResponse) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *ExternalScheduledEventResponse) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.


### GetUserRsvp

`func (o *ExternalScheduledEventResponse) GetUserRsvp() ExternalScheduledEventResponseUserRsvp`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *ExternalScheduledEventResponse) GetUserRsvpOk() (*ExternalScheduledEventResponseUserRsvp, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *ExternalScheduledEventResponse) SetUserRsvp(v ExternalScheduledEventResponseUserRsvp)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *ExternalScheduledEventResponse) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

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


