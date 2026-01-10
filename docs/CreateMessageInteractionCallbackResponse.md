# CreateMessageInteractionCallbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InteractionCallbackTypes**](InteractionCallbackTypes.md) |  | 
**Message** | [**MessageResponse**](MessageResponse.md) |  | 

## Methods

### NewCreateMessageInteractionCallbackResponse

`func NewCreateMessageInteractionCallbackResponse(type_ InteractionCallbackTypes, message MessageResponse, ) *CreateMessageInteractionCallbackResponse`

NewCreateMessageInteractionCallbackResponse instantiates a new CreateMessageInteractionCallbackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageInteractionCallbackResponseWithDefaults

`func NewCreateMessageInteractionCallbackResponseWithDefaults() *CreateMessageInteractionCallbackResponse`

NewCreateMessageInteractionCallbackResponseWithDefaults instantiates a new CreateMessageInteractionCallbackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMessageInteractionCallbackResponse) GetType() InteractionCallbackTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageInteractionCallbackResponse) GetTypeOk() (*InteractionCallbackTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageInteractionCallbackResponse) SetType(v InteractionCallbackTypes)`

SetType sets Type field to given value.


### GetMessage

`func (o *CreateMessageInteractionCallbackResponse) GetMessage() MessageResponse`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateMessageInteractionCallbackResponse) GetMessageOk() (*MessageResponse, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateMessageInteractionCallbackResponse) SetMessage(v MessageResponse)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


