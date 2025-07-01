# BulkUpdateGuildChannelsRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**LockPermissions** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBulkUpdateGuildChannelsRequestInner

`func NewBulkUpdateGuildChannelsRequestInner() *BulkUpdateGuildChannelsRequestInner`

NewBulkUpdateGuildChannelsRequestInner instantiates a new BulkUpdateGuildChannelsRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateGuildChannelsRequestInnerWithDefaults

`func NewBulkUpdateGuildChannelsRequestInnerWithDefaults() *BulkUpdateGuildChannelsRequestInner`

NewBulkUpdateGuildChannelsRequestInnerWithDefaults instantiates a new BulkUpdateGuildChannelsRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkUpdateGuildChannelsRequestInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateGuildChannelsRequestInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateGuildChannelsRequestInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUpdateGuildChannelsRequestInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *BulkUpdateGuildChannelsRequestInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkUpdateGuildChannelsRequestInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkUpdateGuildChannelsRequestInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkUpdateGuildChannelsRequestInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BulkUpdateGuildChannelsRequestInner) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BulkUpdateGuildChannelsRequestInner) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetParentId

`func (o *BulkUpdateGuildChannelsRequestInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BulkUpdateGuildChannelsRequestInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BulkUpdateGuildChannelsRequestInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BulkUpdateGuildChannelsRequestInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetLockPermissions

`func (o *BulkUpdateGuildChannelsRequestInner) GetLockPermissions() bool`

GetLockPermissions returns the LockPermissions field if non-nil, zero value otherwise.

### GetLockPermissionsOk

`func (o *BulkUpdateGuildChannelsRequestInner) GetLockPermissionsOk() (*bool, bool)`

GetLockPermissionsOk returns a tuple with the LockPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockPermissions

`func (o *BulkUpdateGuildChannelsRequestInner) SetLockPermissions(v bool)`

SetLockPermissions sets LockPermissions field to given value.

### HasLockPermissions

`func (o *BulkUpdateGuildChannelsRequestInner) HasLockPermissions() bool`

HasLockPermissions returns a boolean if a field has been set.

### SetLockPermissionsNil

`func (o *BulkUpdateGuildChannelsRequestInner) SetLockPermissionsNil(b bool)`

 SetLockPermissionsNil sets the value for LockPermissions to be an explicit nil

### UnsetLockPermissions
`func (o *BulkUpdateGuildChannelsRequestInner) UnsetLockPermissions()`

UnsetLockPermissions ensures that no value is present for LockPermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


