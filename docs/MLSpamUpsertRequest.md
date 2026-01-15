# MLSpamUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EventType** | [**AutomodEventType**](AutomodEventType.md) |  | 
**Actions** | Pointer to [**[]DefaultKeywordListUpsertRequestActionsInner**](DefaultKeywordListUpsertRequestActionsInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerType** | [**AutomodTriggerType**](AutomodTriggerType.md) |  | 
**TriggerMetadata** | Pointer to [**MLSpamUpsertRequestTriggerMetadata**](MLSpamUpsertRequestTriggerMetadata.md) |  | [optional] 

## Methods

### NewMLSpamUpsertRequest

`func NewMLSpamUpsertRequest(name string, eventType AutomodEventType, triggerType AutomodTriggerType, ) *MLSpamUpsertRequest`

NewMLSpamUpsertRequest instantiates a new MLSpamUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLSpamUpsertRequestWithDefaults

`func NewMLSpamUpsertRequestWithDefaults() *MLSpamUpsertRequest`

NewMLSpamUpsertRequestWithDefaults instantiates a new MLSpamUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MLSpamUpsertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MLSpamUpsertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MLSpamUpsertRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *MLSpamUpsertRequest) GetEventType() AutomodEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MLSpamUpsertRequest) GetEventTypeOk() (*AutomodEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MLSpamUpsertRequest) SetEventType(v AutomodEventType)`

SetEventType sets EventType field to given value.


### GetActions

`func (o *MLSpamUpsertRequest) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MLSpamUpsertRequest) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MLSpamUpsertRequest) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MLSpamUpsertRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEnabled

`func (o *MLSpamUpsertRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MLSpamUpsertRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MLSpamUpsertRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MLSpamUpsertRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExemptRoles

`func (o *MLSpamUpsertRequest) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *MLSpamUpsertRequest) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *MLSpamUpsertRequest) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *MLSpamUpsertRequest) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### GetExemptChannels

`func (o *MLSpamUpsertRequest) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *MLSpamUpsertRequest) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *MLSpamUpsertRequest) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *MLSpamUpsertRequest) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### GetTriggerType

`func (o *MLSpamUpsertRequest) GetTriggerType() AutomodTriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MLSpamUpsertRequest) GetTriggerTypeOk() (*AutomodTriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MLSpamUpsertRequest) SetTriggerType(v AutomodTriggerType)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerMetadata

`func (o *MLSpamUpsertRequest) GetTriggerMetadata() MLSpamUpsertRequestTriggerMetadata`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *MLSpamUpsertRequest) GetTriggerMetadataOk() (*MLSpamUpsertRequestTriggerMetadata, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *MLSpamUpsertRequest) SetTriggerMetadata(v MLSpamUpsertRequestTriggerMetadata)`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *MLSpamUpsertRequest) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


