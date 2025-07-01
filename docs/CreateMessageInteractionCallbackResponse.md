# CreateMessageInteractionCallbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Message** | [**MessageResponse**](MessageResponse.md) |  | 

## Methods

### NewCreateMessageInteractionCallbackResponse

`func NewCreateMessageInteractionCallbackResponse(type_ NullableInt32, message MessageResponse, ) *CreateMessageInteractionCallbackResponse`

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

`func (o *CreateMessageInteractionCallbackResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageInteractionCallbackResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageInteractionCallbackResponse) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateMessageInteractionCallbackResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateMessageInteractionCallbackResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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


