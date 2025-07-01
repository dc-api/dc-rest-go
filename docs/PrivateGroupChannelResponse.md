# PrivateGroupChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **NullableBool** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivateGroupChannelResponse

`func NewPrivateGroupChannelResponse(id string, type_ int32, flags int32, recipients []UserResponse, ) *PrivateGroupChannelResponse`

NewPrivateGroupChannelResponse instantiates a new PrivateGroupChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateGroupChannelResponseWithDefaults

`func NewPrivateGroupChannelResponseWithDefaults() *PrivateGroupChannelResponse`

NewPrivateGroupChannelResponseWithDefaults instantiates a new PrivateGroupChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateGroupChannelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateGroupChannelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateGroupChannelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PrivateGroupChannelResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateGroupChannelResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateGroupChannelResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *PrivateGroupChannelResponse) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *PrivateGroupChannelResponse) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *PrivateGroupChannelResponse) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *PrivateGroupChannelResponse) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *PrivateGroupChannelResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PrivateGroupChannelResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PrivateGroupChannelResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *PrivateGroupChannelResponse) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *PrivateGroupChannelResponse) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *PrivateGroupChannelResponse) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *PrivateGroupChannelResponse) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### SetLastPinTimestampNil

`func (o *PrivateGroupChannelResponse) SetLastPinTimestampNil(b bool)`

 SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil

### UnsetLastPinTimestamp
`func (o *PrivateGroupChannelResponse) UnsetLastPinTimestamp()`

UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
### GetRecipients

`func (o *PrivateGroupChannelResponse) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PrivateGroupChannelResponse) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PrivateGroupChannelResponse) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.


### GetName

`func (o *PrivateGroupChannelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateGroupChannelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateGroupChannelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateGroupChannelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PrivateGroupChannelResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PrivateGroupChannelResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *PrivateGroupChannelResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PrivateGroupChannelResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PrivateGroupChannelResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PrivateGroupChannelResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PrivateGroupChannelResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PrivateGroupChannelResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetOwnerId

`func (o *PrivateGroupChannelResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PrivateGroupChannelResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PrivateGroupChannelResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PrivateGroupChannelResponse) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetManaged

`func (o *PrivateGroupChannelResponse) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PrivateGroupChannelResponse) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PrivateGroupChannelResponse) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PrivateGroupChannelResponse) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PrivateGroupChannelResponse) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PrivateGroupChannelResponse) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetApplicationId

`func (o *PrivateGroupChannelResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PrivateGroupChannelResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PrivateGroupChannelResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PrivateGroupChannelResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


