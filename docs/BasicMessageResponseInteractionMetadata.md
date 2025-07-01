# BasicMessageResponseInteractionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**AuthorizingIntegrationOwners** | **map[string]string** |  | 
**OriginalResponseMessageId** | Pointer to **string** |  | [optional] 
**TargetUser** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**TargetMessageId** | Pointer to **string** |  | [optional] 
**InteractedMessageId** | **string** |  | 
**TriggeringInteractionMetadata** | [**ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata**](ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata.md) |  | 

## Methods

### NewBasicMessageResponseInteractionMetadata

`func NewBasicMessageResponseInteractionMetadata(id string, type_ int32, authorizingIntegrationOwners map[string]string, interactedMessageId string, triggeringInteractionMetadata ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata, ) *BasicMessageResponseInteractionMetadata`

NewBasicMessageResponseInteractionMetadata instantiates a new BasicMessageResponseInteractionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicMessageResponseInteractionMetadataWithDefaults

`func NewBasicMessageResponseInteractionMetadataWithDefaults() *BasicMessageResponseInteractionMetadata`

NewBasicMessageResponseInteractionMetadataWithDefaults instantiates a new BasicMessageResponseInteractionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicMessageResponseInteractionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicMessageResponseInteractionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicMessageResponseInteractionMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BasicMessageResponseInteractionMetadata) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicMessageResponseInteractionMetadata) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicMessageResponseInteractionMetadata) SetType(v int32)`

SetType sets Type field to given value.


### GetUser

`func (o *BasicMessageResponseInteractionMetadata) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BasicMessageResponseInteractionMetadata) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BasicMessageResponseInteractionMetadata) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *BasicMessageResponseInteractionMetadata) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthorizingIntegrationOwners

`func (o *BasicMessageResponseInteractionMetadata) GetAuthorizingIntegrationOwners() map[string]string`

GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field if non-nil, zero value otherwise.

### GetAuthorizingIntegrationOwnersOk

`func (o *BasicMessageResponseInteractionMetadata) GetAuthorizingIntegrationOwnersOk() (*map[string]string, bool)`

GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingIntegrationOwners

`func (o *BasicMessageResponseInteractionMetadata) SetAuthorizingIntegrationOwners(v map[string]string)`

SetAuthorizingIntegrationOwners sets AuthorizingIntegrationOwners field to given value.


### GetOriginalResponseMessageId

`func (o *BasicMessageResponseInteractionMetadata) GetOriginalResponseMessageId() string`

GetOriginalResponseMessageId returns the OriginalResponseMessageId field if non-nil, zero value otherwise.

### GetOriginalResponseMessageIdOk

`func (o *BasicMessageResponseInteractionMetadata) GetOriginalResponseMessageIdOk() (*string, bool)`

GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessageId

`func (o *BasicMessageResponseInteractionMetadata) SetOriginalResponseMessageId(v string)`

SetOriginalResponseMessageId sets OriginalResponseMessageId field to given value.

### HasOriginalResponseMessageId

`func (o *BasicMessageResponseInteractionMetadata) HasOriginalResponseMessageId() bool`

HasOriginalResponseMessageId returns a boolean if a field has been set.

### GetTargetUser

`func (o *BasicMessageResponseInteractionMetadata) GetTargetUser() UserResponse`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *BasicMessageResponseInteractionMetadata) GetTargetUserOk() (*UserResponse, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *BasicMessageResponseInteractionMetadata) SetTargetUser(v UserResponse)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *BasicMessageResponseInteractionMetadata) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetMessageId

`func (o *BasicMessageResponseInteractionMetadata) GetTargetMessageId() string`

GetTargetMessageId returns the TargetMessageId field if non-nil, zero value otherwise.

### GetTargetMessageIdOk

`func (o *BasicMessageResponseInteractionMetadata) GetTargetMessageIdOk() (*string, bool)`

GetTargetMessageIdOk returns a tuple with the TargetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMessageId

`func (o *BasicMessageResponseInteractionMetadata) SetTargetMessageId(v string)`

SetTargetMessageId sets TargetMessageId field to given value.

### HasTargetMessageId

`func (o *BasicMessageResponseInteractionMetadata) HasTargetMessageId() bool`

HasTargetMessageId returns a boolean if a field has been set.

### GetInteractedMessageId

`func (o *BasicMessageResponseInteractionMetadata) GetInteractedMessageId() string`

GetInteractedMessageId returns the InteractedMessageId field if non-nil, zero value otherwise.

### GetInteractedMessageIdOk

`func (o *BasicMessageResponseInteractionMetadata) GetInteractedMessageIdOk() (*string, bool)`

GetInteractedMessageIdOk returns a tuple with the InteractedMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedMessageId

`func (o *BasicMessageResponseInteractionMetadata) SetInteractedMessageId(v string)`

SetInteractedMessageId sets InteractedMessageId field to given value.


### GetTriggeringInteractionMetadata

`func (o *BasicMessageResponseInteractionMetadata) GetTriggeringInteractionMetadata() ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata`

GetTriggeringInteractionMetadata returns the TriggeringInteractionMetadata field if non-nil, zero value otherwise.

### GetTriggeringInteractionMetadataOk

`func (o *BasicMessageResponseInteractionMetadata) GetTriggeringInteractionMetadataOk() (*ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata, bool)`

GetTriggeringInteractionMetadataOk returns a tuple with the TriggeringInteractionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringInteractionMetadata

`func (o *BasicMessageResponseInteractionMetadata) SetTriggeringInteractionMetadata(v ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata)`

SetTriggeringInteractionMetadata sets TriggeringInteractionMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


