# CreateMessageInteractionCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InteractionCallbackTypes**](InteractionCallbackTypes.md) |  | 
**Data** | Pointer to [**CreateMessageInteractionCallbackRequestData**](CreateMessageInteractionCallbackRequestData.md) |  | [optional] 

## Methods

### NewCreateMessageInteractionCallbackRequest

`func NewCreateMessageInteractionCallbackRequest(type_ InteractionCallbackTypes, ) *CreateMessageInteractionCallbackRequest`

NewCreateMessageInteractionCallbackRequest instantiates a new CreateMessageInteractionCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageInteractionCallbackRequestWithDefaults

`func NewCreateMessageInteractionCallbackRequestWithDefaults() *CreateMessageInteractionCallbackRequest`

NewCreateMessageInteractionCallbackRequestWithDefaults instantiates a new CreateMessageInteractionCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMessageInteractionCallbackRequest) GetType() InteractionCallbackTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageInteractionCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageInteractionCallbackRequest) SetType(v InteractionCallbackTypes)`

SetType sets Type field to given value.


### GetData

`func (o *CreateMessageInteractionCallbackRequest) GetData() CreateMessageInteractionCallbackRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateMessageInteractionCallbackRequest) GetDataOk() (*CreateMessageInteractionCallbackRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateMessageInteractionCallbackRequest) SetData(v CreateMessageInteractionCallbackRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateMessageInteractionCallbackRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


