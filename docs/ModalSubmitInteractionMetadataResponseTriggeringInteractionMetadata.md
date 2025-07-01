# ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata

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

## Methods

### NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata

`func NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata(id string, type_ int32, authorizingIntegrationOwners map[string]string, interactedMessageId string, ) *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata`

NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata instantiates a new ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadataWithDefaults

`func NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadataWithDefaults() *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata`

NewModalSubmitInteractionMetadataResponseTriggeringInteractionMetadataWithDefaults instantiates a new ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetType(v int32)`

SetType sets Type field to given value.


### GetUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthorizingIntegrationOwners

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetAuthorizingIntegrationOwners() map[string]string`

GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field if non-nil, zero value otherwise.

### GetAuthorizingIntegrationOwnersOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetAuthorizingIntegrationOwnersOk() (*map[string]string, bool)`

GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingIntegrationOwners

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetAuthorizingIntegrationOwners(v map[string]string)`

SetAuthorizingIntegrationOwners sets AuthorizingIntegrationOwners field to given value.


### GetOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetOriginalResponseMessageId() string`

GetOriginalResponseMessageId returns the OriginalResponseMessageId field if non-nil, zero value otherwise.

### GetOriginalResponseMessageIdOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetOriginalResponseMessageIdOk() (*string, bool)`

GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetOriginalResponseMessageId(v string)`

SetOriginalResponseMessageId sets OriginalResponseMessageId field to given value.

### HasOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) HasOriginalResponseMessageId() bool`

HasOriginalResponseMessageId returns a boolean if a field has been set.

### GetTargetUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetTargetUser() UserResponse`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetTargetUserOk() (*UserResponse, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetTargetUser(v UserResponse)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetTargetMessageId() string`

GetTargetMessageId returns the TargetMessageId field if non-nil, zero value otherwise.

### GetTargetMessageIdOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetTargetMessageIdOk() (*string, bool)`

GetTargetMessageIdOk returns a tuple with the TargetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetTargetMessageId(v string)`

SetTargetMessageId sets TargetMessageId field to given value.

### HasTargetMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) HasTargetMessageId() bool`

HasTargetMessageId returns a boolean if a field has been set.

### GetInteractedMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetInteractedMessageId() string`

GetInteractedMessageId returns the InteractedMessageId field if non-nil, zero value otherwise.

### GetInteractedMessageIdOk

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) GetInteractedMessageIdOk() (*string, bool)`

GetInteractedMessageIdOk returns a tuple with the InteractedMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedMessageId

`func (o *ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata) SetInteractedMessageId(v string)`

SetInteractedMessageId sets InteractedMessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


