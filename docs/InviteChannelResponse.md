# InviteChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to [**[]InviteChannelRecipientResponse**](InviteChannelRecipientResponse.md) |  | [optional] 

## Methods

### NewInviteChannelResponse

`func NewInviteChannelResponse(id string, type_ int32, ) *InviteChannelResponse`

NewInviteChannelResponse instantiates a new InviteChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteChannelResponseWithDefaults

`func NewInviteChannelResponseWithDefaults() *InviteChannelResponse`

NewInviteChannelResponseWithDefaults instantiates a new InviteChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InviteChannelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InviteChannelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InviteChannelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InviteChannelResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InviteChannelResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InviteChannelResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *InviteChannelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InviteChannelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InviteChannelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InviteChannelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InviteChannelResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InviteChannelResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *InviteChannelResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *InviteChannelResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *InviteChannelResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *InviteChannelResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *InviteChannelResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *InviteChannelResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetRecipients

`func (o *InviteChannelResponse) GetRecipients() []InviteChannelRecipientResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InviteChannelResponse) GetRecipientsOk() (*[]InviteChannelRecipientResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InviteChannelResponse) SetRecipients(v []InviteChannelRecipientResponse)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InviteChannelResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *InviteChannelResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *InviteChannelResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


