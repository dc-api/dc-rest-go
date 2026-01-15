# MessageAllowedMentionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parse** | Pointer to [**[]MessageAllowedMentionsRequestParseInner**](MessageAllowedMentionsRequestParseInner.md) |  | [optional] 
**Users** | Pointer to [**[]GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Roles** | Pointer to [**[]GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**RepliedUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewMessageAllowedMentionsRequest

`func NewMessageAllowedMentionsRequest() *MessageAllowedMentionsRequest`

NewMessageAllowedMentionsRequest instantiates a new MessageAllowedMentionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAllowedMentionsRequestWithDefaults

`func NewMessageAllowedMentionsRequestWithDefaults() *MessageAllowedMentionsRequest`

NewMessageAllowedMentionsRequestWithDefaults instantiates a new MessageAllowedMentionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParse

`func (o *MessageAllowedMentionsRequest) GetParse() []MessageAllowedMentionsRequestParseInner`

GetParse returns the Parse field if non-nil, zero value otherwise.

### GetParseOk

`func (o *MessageAllowedMentionsRequest) GetParseOk() (*[]MessageAllowedMentionsRequestParseInner, bool)`

GetParseOk returns a tuple with the Parse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParse

`func (o *MessageAllowedMentionsRequest) SetParse(v []MessageAllowedMentionsRequestParseInner)`

SetParse sets Parse field to given value.

### HasParse

`func (o *MessageAllowedMentionsRequest) HasParse() bool`

HasParse returns a boolean if a field has been set.

### GetUsers

`func (o *MessageAllowedMentionsRequest) GetUsers() []GetEntitlementsSkuIdsParameterOneOfInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MessageAllowedMentionsRequest) GetUsersOk() (*[]GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MessageAllowedMentionsRequest) SetUsers(v []GetEntitlementsSkuIdsParameterOneOfInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MessageAllowedMentionsRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetRoles

`func (o *MessageAllowedMentionsRequest) GetRoles() []GetEntitlementsSkuIdsParameterOneOfInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MessageAllowedMentionsRequest) GetRolesOk() (*[]GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MessageAllowedMentionsRequest) SetRoles(v []GetEntitlementsSkuIdsParameterOneOfInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MessageAllowedMentionsRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRepliedUser

`func (o *MessageAllowedMentionsRequest) GetRepliedUser() bool`

GetRepliedUser returns the RepliedUser field if non-nil, zero value otherwise.

### GetRepliedUserOk

`func (o *MessageAllowedMentionsRequest) GetRepliedUserOk() (*bool, bool)`

GetRepliedUserOk returns a tuple with the RepliedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedUser

`func (o *MessageAllowedMentionsRequest) SetRepliedUser(v bool)`

SetRepliedUser sets RepliedUser field to given value.

### HasRepliedUser

`func (o *MessageAllowedMentionsRequest) HasRepliedUser() bool`

HasRepliedUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


