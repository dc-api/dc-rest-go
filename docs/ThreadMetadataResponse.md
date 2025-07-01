# ThreadMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** |  | 
**ArchiveTimestamp** | Pointer to **NullableTime** |  | [optional] 
**AutoArchiveDuration** | **int32** |  | 
**Locked** | **bool** |  | 
**CreateTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Invitable** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewThreadMetadataResponse

`func NewThreadMetadataResponse(archived bool, autoArchiveDuration int32, locked bool, ) *ThreadMetadataResponse`

NewThreadMetadataResponse instantiates a new ThreadMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadMetadataResponseWithDefaults

`func NewThreadMetadataResponseWithDefaults() *ThreadMetadataResponse`

NewThreadMetadataResponseWithDefaults instantiates a new ThreadMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *ThreadMetadataResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ThreadMetadataResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ThreadMetadataResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchiveTimestamp

`func (o *ThreadMetadataResponse) GetArchiveTimestamp() time.Time`

GetArchiveTimestamp returns the ArchiveTimestamp field if non-nil, zero value otherwise.

### GetArchiveTimestampOk

`func (o *ThreadMetadataResponse) GetArchiveTimestampOk() (*time.Time, bool)`

GetArchiveTimestampOk returns a tuple with the ArchiveTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTimestamp

`func (o *ThreadMetadataResponse) SetArchiveTimestamp(v time.Time)`

SetArchiveTimestamp sets ArchiveTimestamp field to given value.

### HasArchiveTimestamp

`func (o *ThreadMetadataResponse) HasArchiveTimestamp() bool`

HasArchiveTimestamp returns a boolean if a field has been set.

### SetArchiveTimestampNil

`func (o *ThreadMetadataResponse) SetArchiveTimestampNil(b bool)`

 SetArchiveTimestampNil sets the value for ArchiveTimestamp to be an explicit nil

### UnsetArchiveTimestamp
`func (o *ThreadMetadataResponse) UnsetArchiveTimestamp()`

UnsetArchiveTimestamp ensures that no value is present for ArchiveTimestamp, not even an explicit nil
### GetAutoArchiveDuration

`func (o *ThreadMetadataResponse) GetAutoArchiveDuration() int32`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *ThreadMetadataResponse) GetAutoArchiveDurationOk() (*int32, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *ThreadMetadataResponse) SetAutoArchiveDuration(v int32)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.


### GetLocked

`func (o *ThreadMetadataResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ThreadMetadataResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ThreadMetadataResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetCreateTimestamp

`func (o *ThreadMetadataResponse) GetCreateTimestamp() time.Time`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ThreadMetadataResponse) GetCreateTimestampOk() (*time.Time, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ThreadMetadataResponse) SetCreateTimestamp(v time.Time)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ThreadMetadataResponse) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### SetCreateTimestampNil

`func (o *ThreadMetadataResponse) SetCreateTimestampNil(b bool)`

 SetCreateTimestampNil sets the value for CreateTimestamp to be an explicit nil

### UnsetCreateTimestamp
`func (o *ThreadMetadataResponse) UnsetCreateTimestamp()`

UnsetCreateTimestamp ensures that no value is present for CreateTimestamp, not even an explicit nil
### GetInvitable

`func (o *ThreadMetadataResponse) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *ThreadMetadataResponse) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *ThreadMetadataResponse) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *ThreadMetadataResponse) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.

### SetInvitableNil

`func (o *ThreadMetadataResponse) SetInvitableNil(b bool)`

 SetInvitableNil sets the value for Invitable to be an explicit nil

### UnsetInvitable
`func (o *ThreadMetadataResponse) UnsetInvitable()`

UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


