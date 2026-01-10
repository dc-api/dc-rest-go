# FriendInviteResponseChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to [**[]InviteChannelRecipientResponse**](InviteChannelRecipientResponse.md) |  | [optional] 

## Methods

### NewFriendInviteResponseChannel

`func NewFriendInviteResponseChannel(id string, type_ ChannelTypes, ) *FriendInviteResponseChannel`

NewFriendInviteResponseChannel instantiates a new FriendInviteResponseChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFriendInviteResponseChannelWithDefaults

`func NewFriendInviteResponseChannelWithDefaults() *FriendInviteResponseChannel`

NewFriendInviteResponseChannelWithDefaults instantiates a new FriendInviteResponseChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FriendInviteResponseChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FriendInviteResponseChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FriendInviteResponseChannel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *FriendInviteResponseChannel) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FriendInviteResponseChannel) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FriendInviteResponseChannel) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetName

`func (o *FriendInviteResponseChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FriendInviteResponseChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FriendInviteResponseChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FriendInviteResponseChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *FriendInviteResponseChannel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FriendInviteResponseChannel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FriendInviteResponseChannel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FriendInviteResponseChannel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetRecipients

`func (o *FriendInviteResponseChannel) GetRecipients() []InviteChannelRecipientResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *FriendInviteResponseChannel) GetRecipientsOk() (*[]InviteChannelRecipientResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *FriendInviteResponseChannel) SetRecipients(v []InviteChannelRecipientResponse)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *FriendInviteResponseChannel) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


