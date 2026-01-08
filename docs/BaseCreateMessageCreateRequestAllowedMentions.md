# BaseCreateMessageCreateRequestAllowedMentions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parse** | Pointer to [**[]MessageAllowedMentionsRequestParseInner**](MessageAllowedMentionsRequestParseInner.md) |  | [optional] 
**Users** | Pointer to [**[]GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Roles** | Pointer to [**[]GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**RepliedUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewBaseCreateMessageCreateRequestAllowedMentions

`func NewBaseCreateMessageCreateRequestAllowedMentions() *BaseCreateMessageCreateRequestAllowedMentions`

NewBaseCreateMessageCreateRequestAllowedMentions instantiates a new BaseCreateMessageCreateRequestAllowedMentions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCreateMessageCreateRequestAllowedMentionsWithDefaults

`func NewBaseCreateMessageCreateRequestAllowedMentionsWithDefaults() *BaseCreateMessageCreateRequestAllowedMentions`

NewBaseCreateMessageCreateRequestAllowedMentionsWithDefaults instantiates a new BaseCreateMessageCreateRequestAllowedMentions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParse

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetParse() []MessageAllowedMentionsRequestParseInner`

GetParse returns the Parse field if non-nil, zero value otherwise.

### GetParseOk

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetParseOk() (*[]MessageAllowedMentionsRequestParseInner, bool)`

GetParseOk returns a tuple with the Parse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParse

`func (o *BaseCreateMessageCreateRequestAllowedMentions) SetParse(v []MessageAllowedMentionsRequestParseInner)`

SetParse sets Parse field to given value.

### HasParse

`func (o *BaseCreateMessageCreateRequestAllowedMentions) HasParse() bool`

HasParse returns a boolean if a field has been set.

### GetUsers

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetUsers() []GetEntitlementsSkuIdsParameterOneOfInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetUsersOk() (*[]GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BaseCreateMessageCreateRequestAllowedMentions) SetUsers(v []GetEntitlementsSkuIdsParameterOneOfInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BaseCreateMessageCreateRequestAllowedMentions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetRoles

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetRoles() []GetEntitlementsSkuIdsParameterOneOfInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetRolesOk() (*[]GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BaseCreateMessageCreateRequestAllowedMentions) SetRoles(v []GetEntitlementsSkuIdsParameterOneOfInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BaseCreateMessageCreateRequestAllowedMentions) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRepliedUser

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetRepliedUser() bool`

GetRepliedUser returns the RepliedUser field if non-nil, zero value otherwise.

### GetRepliedUserOk

`func (o *BaseCreateMessageCreateRequestAllowedMentions) GetRepliedUserOk() (*bool, bool)`

GetRepliedUserOk returns a tuple with the RepliedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedUser

`func (o *BaseCreateMessageCreateRequestAllowedMentions) SetRepliedUser(v bool)`

SetRepliedUser sets RepliedUser field to given value.

### HasRepliedUser

`func (o *BaseCreateMessageCreateRequestAllowedMentions) HasRepliedUser() bool`

HasRepliedUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


