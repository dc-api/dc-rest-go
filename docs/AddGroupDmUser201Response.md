# AddGroupDmUser201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **time.Time** |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 

## Methods

### NewAddGroupDmUser201Response

`func NewAddGroupDmUser201Response(id string, type_ int32, flags int32, recipients []UserResponse, ) *AddGroupDmUser201Response`

NewAddGroupDmUser201Response instantiates a new AddGroupDmUser201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupDmUser201ResponseWithDefaults

`func NewAddGroupDmUser201ResponseWithDefaults() *AddGroupDmUser201Response`

NewAddGroupDmUser201ResponseWithDefaults instantiates a new AddGroupDmUser201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddGroupDmUser201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddGroupDmUser201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddGroupDmUser201Response) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AddGroupDmUser201Response) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddGroupDmUser201Response) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddGroupDmUser201Response) SetType(v int32)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *AddGroupDmUser201Response) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *AddGroupDmUser201Response) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *AddGroupDmUser201Response) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *AddGroupDmUser201Response) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *AddGroupDmUser201Response) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *AddGroupDmUser201Response) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *AddGroupDmUser201Response) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *AddGroupDmUser201Response) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *AddGroupDmUser201Response) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *AddGroupDmUser201Response) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *AddGroupDmUser201Response) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### GetRecipients

`func (o *AddGroupDmUser201Response) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AddGroupDmUser201Response) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AddGroupDmUser201Response) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.


### GetName

`func (o *AddGroupDmUser201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddGroupDmUser201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddGroupDmUser201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddGroupDmUser201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *AddGroupDmUser201Response) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AddGroupDmUser201Response) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AddGroupDmUser201Response) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AddGroupDmUser201Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetOwnerId

`func (o *AddGroupDmUser201Response) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AddGroupDmUser201Response) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AddGroupDmUser201Response) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AddGroupDmUser201Response) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetManaged

`func (o *AddGroupDmUser201Response) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AddGroupDmUser201Response) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AddGroupDmUser201Response) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *AddGroupDmUser201Response) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetApplicationId

`func (o *AddGroupDmUser201Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AddGroupDmUser201Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AddGroupDmUser201Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AddGroupDmUser201Response) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


