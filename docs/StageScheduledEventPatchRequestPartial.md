# StageScheduledEventPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ExternalScheduledEventPatchRequestPartialStatus**](ExternalScheduledEventPatchRequestPartialStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** |  | [optional] 
**ScheduledEndTime** | Pointer to **time.Time** |  | [optional] 
**EntityType** | Pointer to [**StageScheduledEventPatchRequestPartialEntityType**](StageScheduledEventPatchRequestPartialEntityType.md) |  | [optional] 
**PrivacyLevel** | Pointer to [**GuildScheduledEventPrivacyLevels**](GuildScheduledEventPrivacyLevels.md) |  | [optional] 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**EntityMetadata** | Pointer to [**StageScheduledEventCreateRequestEntityMetadata**](StageScheduledEventCreateRequestEntityMetadata.md) |  | [optional] 

## Methods

### NewStageScheduledEventPatchRequestPartial

`func NewStageScheduledEventPatchRequestPartial() *StageScheduledEventPatchRequestPartial`

NewStageScheduledEventPatchRequestPartial instantiates a new StageScheduledEventPatchRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageScheduledEventPatchRequestPartialWithDefaults

`func NewStageScheduledEventPatchRequestPartialWithDefaults() *StageScheduledEventPatchRequestPartial`

NewStageScheduledEventPatchRequestPartialWithDefaults instantiates a new StageScheduledEventPatchRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StageScheduledEventPatchRequestPartial) GetStatus() ExternalScheduledEventPatchRequestPartialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageScheduledEventPatchRequestPartial) GetStatusOk() (*ExternalScheduledEventPatchRequestPartialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageScheduledEventPatchRequestPartial) SetStatus(v ExternalScheduledEventPatchRequestPartialStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StageScheduledEventPatchRequestPartial) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *StageScheduledEventPatchRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageScheduledEventPatchRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageScheduledEventPatchRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageScheduledEventPatchRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StageScheduledEventPatchRequestPartial) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StageScheduledEventPatchRequestPartial) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StageScheduledEventPatchRequestPartial) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StageScheduledEventPatchRequestPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *StageScheduledEventPatchRequestPartial) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *StageScheduledEventPatchRequestPartial) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *StageScheduledEventPatchRequestPartial) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *StageScheduledEventPatchRequestPartial) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *StageScheduledEventPatchRequestPartial) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *StageScheduledEventPatchRequestPartial) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *StageScheduledEventPatchRequestPartial) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *StageScheduledEventPatchRequestPartial) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetScheduledEndTime

`func (o *StageScheduledEventPatchRequestPartial) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *StageScheduledEventPatchRequestPartial) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *StageScheduledEventPatchRequestPartial) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *StageScheduledEventPatchRequestPartial) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### GetEntityType

`func (o *StageScheduledEventPatchRequestPartial) GetEntityType() StageScheduledEventPatchRequestPartialEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StageScheduledEventPatchRequestPartial) GetEntityTypeOk() (*StageScheduledEventPatchRequestPartialEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StageScheduledEventPatchRequestPartial) SetEntityType(v StageScheduledEventPatchRequestPartialEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *StageScheduledEventPatchRequestPartial) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *StageScheduledEventPatchRequestPartial) GetPrivacyLevel() GuildScheduledEventPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *StageScheduledEventPatchRequestPartial) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *StageScheduledEventPatchRequestPartial) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *StageScheduledEventPatchRequestPartial) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetChannelId

`func (o *StageScheduledEventPatchRequestPartial) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *StageScheduledEventPatchRequestPartial) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *StageScheduledEventPatchRequestPartial) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *StageScheduledEventPatchRequestPartial) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *StageScheduledEventPatchRequestPartial) GetEntityMetadata() StageScheduledEventCreateRequestEntityMetadata`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *StageScheduledEventPatchRequestPartial) GetEntityMetadataOk() (*StageScheduledEventCreateRequestEntityMetadata, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *StageScheduledEventPatchRequestPartial) SetEntityMetadata(v StageScheduledEventCreateRequestEntityMetadata)`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *StageScheduledEventPatchRequestPartial) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


