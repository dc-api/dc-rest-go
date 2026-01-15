# StageScheduledEventCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **time.Time** |  | [optional] 
**PrivacyLevel** | [**GuildScheduledEventPrivacyLevels**](GuildScheduledEventPrivacyLevels.md) |  | 
**EntityType** | [**GuildScheduledEventEntityTypes**](GuildScheduledEventEntityTypes.md) |  | 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**EntityMetadata** | Pointer to [**StageScheduledEventCreateRequestEntityMetadata**](StageScheduledEventCreateRequestEntityMetadata.md) |  | [optional] 

## Methods

### NewStageScheduledEventCreateRequest

`func NewStageScheduledEventCreateRequest(name string, scheduledStartTime time.Time, privacyLevel GuildScheduledEventPrivacyLevels, entityType GuildScheduledEventEntityTypes, ) *StageScheduledEventCreateRequest`

NewStageScheduledEventCreateRequest instantiates a new StageScheduledEventCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageScheduledEventCreateRequestWithDefaults

`func NewStageScheduledEventCreateRequestWithDefaults() *StageScheduledEventCreateRequest`

NewStageScheduledEventCreateRequestWithDefaults instantiates a new StageScheduledEventCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StageScheduledEventCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageScheduledEventCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageScheduledEventCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StageScheduledEventCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StageScheduledEventCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StageScheduledEventCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StageScheduledEventCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *StageScheduledEventCreateRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *StageScheduledEventCreateRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *StageScheduledEventCreateRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *StageScheduledEventCreateRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *StageScheduledEventCreateRequest) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *StageScheduledEventCreateRequest) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *StageScheduledEventCreateRequest) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *StageScheduledEventCreateRequest) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *StageScheduledEventCreateRequest) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *StageScheduledEventCreateRequest) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *StageScheduledEventCreateRequest) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *StageScheduledEventCreateRequest) GetPrivacyLevel() GuildScheduledEventPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *StageScheduledEventCreateRequest) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *StageScheduledEventCreateRequest) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.


### GetEntityType

`func (o *StageScheduledEventCreateRequest) GetEntityType() GuildScheduledEventEntityTypes`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StageScheduledEventCreateRequest) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StageScheduledEventCreateRequest) SetEntityType(v GuildScheduledEventEntityTypes)`

SetEntityType sets EntityType field to given value.


### GetChannelId

`func (o *StageScheduledEventCreateRequest) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *StageScheduledEventCreateRequest) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *StageScheduledEventCreateRequest) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *StageScheduledEventCreateRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *StageScheduledEventCreateRequest) GetEntityMetadata() StageScheduledEventCreateRequestEntityMetadata`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *StageScheduledEventCreateRequest) GetEntityMetadataOk() (*StageScheduledEventCreateRequestEntityMetadata, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *StageScheduledEventCreateRequest) SetEntityMetadata(v StageScheduledEventCreateRequestEntityMetadata)`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *StageScheduledEventCreateRequest) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


