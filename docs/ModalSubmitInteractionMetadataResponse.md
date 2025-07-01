# ModalSubmitInteractionMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**AuthorizingIntegrationOwners** | **map[string]string** |  | 
**OriginalResponseMessageId** | Pointer to **string** |  | [optional] 
**TriggeringInteractionMetadata** | [**ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata**](ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata.md) |  | 

## Methods

### NewModalSubmitInteractionMetadataResponse

`func NewModalSubmitInteractionMetadataResponse(id string, type_ int32, authorizingIntegrationOwners map[string]string, triggeringInteractionMetadata ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata, ) *ModalSubmitInteractionMetadataResponse`

NewModalSubmitInteractionMetadataResponse instantiates a new ModalSubmitInteractionMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModalSubmitInteractionMetadataResponseWithDefaults

`func NewModalSubmitInteractionMetadataResponseWithDefaults() *ModalSubmitInteractionMetadataResponse`

NewModalSubmitInteractionMetadataResponseWithDefaults instantiates a new ModalSubmitInteractionMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModalSubmitInteractionMetadataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModalSubmitInteractionMetadataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModalSubmitInteractionMetadataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ModalSubmitInteractionMetadataResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModalSubmitInteractionMetadataResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModalSubmitInteractionMetadataResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetUser

`func (o *ModalSubmitInteractionMetadataResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModalSubmitInteractionMetadataResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModalSubmitInteractionMetadataResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ModalSubmitInteractionMetadataResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ModalSubmitInteractionMetadataResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ModalSubmitInteractionMetadataResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAuthorizingIntegrationOwners

`func (o *ModalSubmitInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string`

GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field if non-nil, zero value otherwise.

### GetAuthorizingIntegrationOwnersOk

`func (o *ModalSubmitInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (*map[string]string, bool)`

GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingIntegrationOwners

`func (o *ModalSubmitInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string)`

SetAuthorizingIntegrationOwners sets AuthorizingIntegrationOwners field to given value.


### GetOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponse) GetOriginalResponseMessageId() string`

GetOriginalResponseMessageId returns the OriginalResponseMessageId field if non-nil, zero value otherwise.

### GetOriginalResponseMessageIdOk

`func (o *ModalSubmitInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool)`

GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponse) SetOriginalResponseMessageId(v string)`

SetOriginalResponseMessageId sets OriginalResponseMessageId field to given value.

### HasOriginalResponseMessageId

`func (o *ModalSubmitInteractionMetadataResponse) HasOriginalResponseMessageId() bool`

HasOriginalResponseMessageId returns a boolean if a field has been set.

### GetTriggeringInteractionMetadata

`func (o *ModalSubmitInteractionMetadataResponse) GetTriggeringInteractionMetadata() ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata`

GetTriggeringInteractionMetadata returns the TriggeringInteractionMetadata field if non-nil, zero value otherwise.

### GetTriggeringInteractionMetadataOk

`func (o *ModalSubmitInteractionMetadataResponse) GetTriggeringInteractionMetadataOk() (*ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata, bool)`

GetTriggeringInteractionMetadataOk returns a tuple with the TriggeringInteractionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringInteractionMetadata

`func (o *ModalSubmitInteractionMetadataResponse) SetTriggeringInteractionMetadata(v ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata)`

SetTriggeringInteractionMetadata sets TriggeringInteractionMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


