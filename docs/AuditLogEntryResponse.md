# AuditLogEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActionType** | **NullableInt32** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**Changes** | Pointer to [**[]AuditLogObjectChangeResponse**](AuditLogObjectChangeResponse.md) |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuditLogEntryResponse

`func NewAuditLogEntryResponse(id string, actionType NullableInt32, ) *AuditLogEntryResponse`

NewAuditLogEntryResponse instantiates a new AuditLogEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryResponseWithDefaults

`func NewAuditLogEntryResponseWithDefaults() *AuditLogEntryResponse`

NewAuditLogEntryResponseWithDefaults instantiates a new AuditLogEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLogEntryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEntryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEntryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetActionType

`func (o *AuditLogEntryResponse) GetActionType() int32`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditLogEntryResponse) GetActionTypeOk() (*int32, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditLogEntryResponse) SetActionType(v int32)`

SetActionType sets ActionType field to given value.


### SetActionTypeNil

`func (o *AuditLogEntryResponse) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *AuditLogEntryResponse) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetUserId

`func (o *AuditLogEntryResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditLogEntryResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditLogEntryResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditLogEntryResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTargetId

`func (o *AuditLogEntryResponse) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AuditLogEntryResponse) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AuditLogEntryResponse) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AuditLogEntryResponse) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetChanges

`func (o *AuditLogEntryResponse) GetChanges() []AuditLogObjectChangeResponse`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *AuditLogEntryResponse) GetChangesOk() (*[]AuditLogObjectChangeResponse, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *AuditLogEntryResponse) SetChanges(v []AuditLogObjectChangeResponse)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *AuditLogEntryResponse) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *AuditLogEntryResponse) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *AuditLogEntryResponse) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil
### GetOptions

`func (o *AuditLogEntryResponse) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuditLogEntryResponse) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuditLogEntryResponse) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AuditLogEntryResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReason

`func (o *AuditLogEntryResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AuditLogEntryResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AuditLogEntryResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AuditLogEntryResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AuditLogEntryResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AuditLogEntryResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


