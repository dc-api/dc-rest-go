# CreateInteractionResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Data** | [**IncomingWebhookUpdateForInteractionCallbackRequestPartial**](IncomingWebhookUpdateForInteractionCallbackRequestPartial.md) |  | 

## Methods

### NewCreateInteractionResponseRequest

`func NewCreateInteractionResponseRequest(type_ NullableInt32, data IncomingWebhookUpdateForInteractionCallbackRequestPartial, ) *CreateInteractionResponseRequest`

NewCreateInteractionResponseRequest instantiates a new CreateInteractionResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInteractionResponseRequestWithDefaults

`func NewCreateInteractionResponseRequestWithDefaults() *CreateInteractionResponseRequest`

NewCreateInteractionResponseRequestWithDefaults instantiates a new CreateInteractionResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateInteractionResponseRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateInteractionResponseRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateInteractionResponseRequest) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateInteractionResponseRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateInteractionResponseRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *CreateInteractionResponseRequest) GetData() IncomingWebhookUpdateForInteractionCallbackRequestPartial`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateInteractionResponseRequest) GetDataOk() (*IncomingWebhookUpdateForInteractionCallbackRequestPartial, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateInteractionResponseRequest) SetData(v IncomingWebhookUpdateForInteractionCallbackRequestPartial)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


