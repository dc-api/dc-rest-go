# UpdateMessageInteractionCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Data** | Pointer to [**NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial**](IncomingWebhookUpdateForInteractionCallbackRequestPartial.md) |  | [optional] 

## Methods

### NewUpdateMessageInteractionCallbackRequest

`func NewUpdateMessageInteractionCallbackRequest(type_ NullableInt32, ) *UpdateMessageInteractionCallbackRequest`

NewUpdateMessageInteractionCallbackRequest instantiates a new UpdateMessageInteractionCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageInteractionCallbackRequestWithDefaults

`func NewUpdateMessageInteractionCallbackRequestWithDefaults() *UpdateMessageInteractionCallbackRequest`

NewUpdateMessageInteractionCallbackRequestWithDefaults instantiates a new UpdateMessageInteractionCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateMessageInteractionCallbackRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMessageInteractionCallbackRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMessageInteractionCallbackRequest) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *UpdateMessageInteractionCallbackRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateMessageInteractionCallbackRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *UpdateMessageInteractionCallbackRequest) GetData() IncomingWebhookUpdateForInteractionCallbackRequestPartial`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateMessageInteractionCallbackRequest) GetDataOk() (*IncomingWebhookUpdateForInteractionCallbackRequestPartial, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateMessageInteractionCallbackRequest) SetData(v IncomingWebhookUpdateForInteractionCallbackRequestPartial)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateMessageInteractionCallbackRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UpdateMessageInteractionCallbackRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UpdateMessageInteractionCallbackRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


