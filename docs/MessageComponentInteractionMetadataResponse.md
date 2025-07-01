# MessageComponentInteractionMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**AuthorizingIntegrationOwners** | **map[string]string** |  | 
**OriginalResponseMessageId** | Pointer to **string** |  | [optional] 
**InteractedMessageId** | **string** |  | 

## Methods

### NewMessageComponentInteractionMetadataResponse

`func NewMessageComponentInteractionMetadataResponse(id string, type_ int32, authorizingIntegrationOwners map[string]string, interactedMessageId string, ) *MessageComponentInteractionMetadataResponse`

NewMessageComponentInteractionMetadataResponse instantiates a new MessageComponentInteractionMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentInteractionMetadataResponseWithDefaults

`func NewMessageComponentInteractionMetadataResponseWithDefaults() *MessageComponentInteractionMetadataResponse`

NewMessageComponentInteractionMetadataResponseWithDefaults instantiates a new MessageComponentInteractionMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageComponentInteractionMetadataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentInteractionMetadataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentInteractionMetadataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MessageComponentInteractionMetadataResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentInteractionMetadataResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentInteractionMetadataResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetUser

`func (o *MessageComponentInteractionMetadataResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MessageComponentInteractionMetadataResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MessageComponentInteractionMetadataResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *MessageComponentInteractionMetadataResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MessageComponentInteractionMetadataResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MessageComponentInteractionMetadataResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAuthorizingIntegrationOwners

`func (o *MessageComponentInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string`

GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field if non-nil, zero value otherwise.

### GetAuthorizingIntegrationOwnersOk

`func (o *MessageComponentInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (*map[string]string, bool)`

GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingIntegrationOwners

`func (o *MessageComponentInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string)`

SetAuthorizingIntegrationOwners sets AuthorizingIntegrationOwners field to given value.


### GetOriginalResponseMessageId

`func (o *MessageComponentInteractionMetadataResponse) GetOriginalResponseMessageId() string`

GetOriginalResponseMessageId returns the OriginalResponseMessageId field if non-nil, zero value otherwise.

### GetOriginalResponseMessageIdOk

`func (o *MessageComponentInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool)`

GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessageId

`func (o *MessageComponentInteractionMetadataResponse) SetOriginalResponseMessageId(v string)`

SetOriginalResponseMessageId sets OriginalResponseMessageId field to given value.

### HasOriginalResponseMessageId

`func (o *MessageComponentInteractionMetadataResponse) HasOriginalResponseMessageId() bool`

HasOriginalResponseMessageId returns a boolean if a field has been set.

### GetInteractedMessageId

`func (o *MessageComponentInteractionMetadataResponse) GetInteractedMessageId() string`

GetInteractedMessageId returns the InteractedMessageId field if non-nil, zero value otherwise.

### GetInteractedMessageIdOk

`func (o *MessageComponentInteractionMetadataResponse) GetInteractedMessageIdOk() (*string, bool)`

GetInteractedMessageIdOk returns a tuple with the InteractedMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedMessageId

`func (o *MessageComponentInteractionMetadataResponse) SetInteractedMessageId(v string)`

SetInteractedMessageId sets InteractedMessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


