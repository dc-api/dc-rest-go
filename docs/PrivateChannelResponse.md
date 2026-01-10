# PrivateChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**LastMessageId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **time.Time** |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 

## Methods

### NewPrivateChannelResponse

`func NewPrivateChannelResponse(id string, type_ ChannelTypes, flags int32, recipients []UserResponse, ) *PrivateChannelResponse`

NewPrivateChannelResponse instantiates a new PrivateChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateChannelResponseWithDefaults

`func NewPrivateChannelResponseWithDefaults() *PrivateChannelResponse`

NewPrivateChannelResponseWithDefaults instantiates a new PrivateChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateChannelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateChannelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateChannelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PrivateChannelResponse) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateChannelResponse) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateChannelResponse) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *PrivateChannelResponse) GetLastMessageId() GetEntitlementsSkuIdsParameterOneOfInner`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *PrivateChannelResponse) GetLastMessageIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *PrivateChannelResponse) SetLastMessageId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *PrivateChannelResponse) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *PrivateChannelResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PrivateChannelResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PrivateChannelResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *PrivateChannelResponse) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *PrivateChannelResponse) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *PrivateChannelResponse) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *PrivateChannelResponse) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### GetRecipients

`func (o *PrivateChannelResponse) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PrivateChannelResponse) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PrivateChannelResponse) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


