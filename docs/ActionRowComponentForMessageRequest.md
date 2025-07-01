# ActionRowComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Components** | [**[]ActionRowComponentForMessageRequestComponentsInner**](ActionRowComponentForMessageRequestComponentsInner.md) |  | 

## Methods

### NewActionRowComponentForMessageRequest

`func NewActionRowComponentForMessageRequest(type_ int32, components []ActionRowComponentForMessageRequestComponentsInner, ) *ActionRowComponentForMessageRequest`

NewActionRowComponentForMessageRequest instantiates a new ActionRowComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowComponentForMessageRequestWithDefaults

`func NewActionRowComponentForMessageRequestWithDefaults() *ActionRowComponentForMessageRequest`

NewActionRowComponentForMessageRequestWithDefaults instantiates a new ActionRowComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRowComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetComponents

`func (o *ActionRowComponentForMessageRequest) GetComponents() []ActionRowComponentForMessageRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ActionRowComponentForMessageRequest) GetComponentsOk() (*[]ActionRowComponentForMessageRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ActionRowComponentForMessageRequest) SetComponents(v []ActionRowComponentForMessageRequestComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


