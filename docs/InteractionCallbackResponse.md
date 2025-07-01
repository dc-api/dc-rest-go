# InteractionCallbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interaction** | [**InteractionResponse**](InteractionResponse.md) |  | 
**Resource** | Pointer to [**NullableInteractionCallbackResponseResource**](InteractionCallbackResponseResource.md) |  | [optional] 

## Methods

### NewInteractionCallbackResponse

`func NewInteractionCallbackResponse(interaction InteractionResponse, ) *InteractionCallbackResponse`

NewInteractionCallbackResponse instantiates a new InteractionCallbackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionCallbackResponseWithDefaults

`func NewInteractionCallbackResponseWithDefaults() *InteractionCallbackResponse`

NewInteractionCallbackResponseWithDefaults instantiates a new InteractionCallbackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteraction

`func (o *InteractionCallbackResponse) GetInteraction() InteractionResponse`

GetInteraction returns the Interaction field if non-nil, zero value otherwise.

### GetInteractionOk

`func (o *InteractionCallbackResponse) GetInteractionOk() (*InteractionResponse, bool)`

GetInteractionOk returns a tuple with the Interaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteraction

`func (o *InteractionCallbackResponse) SetInteraction(v InteractionResponse)`

SetInteraction sets Interaction field to given value.


### GetResource

`func (o *InteractionCallbackResponse) GetResource() InteractionCallbackResponseResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *InteractionCallbackResponse) GetResourceOk() (*InteractionCallbackResponseResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *InteractionCallbackResponse) SetResource(v InteractionCallbackResponseResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *InteractionCallbackResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *InteractionCallbackResponse) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *InteractionCallbackResponse) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


