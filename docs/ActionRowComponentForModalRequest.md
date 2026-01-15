# ActionRowComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Components** | [**[]TextInputComponentForModalRequest**](TextInputComponentForModalRequest.md) |  | 

## Methods

### NewActionRowComponentForModalRequest

`func NewActionRowComponentForModalRequest(type_ MessageComponentTypes, components []TextInputComponentForModalRequest, ) *ActionRowComponentForModalRequest`

NewActionRowComponentForModalRequest instantiates a new ActionRowComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowComponentForModalRequestWithDefaults

`func NewActionRowComponentForModalRequestWithDefaults() *ActionRowComponentForModalRequest`

NewActionRowComponentForModalRequestWithDefaults instantiates a new ActionRowComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRowComponentForModalRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentForModalRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentForModalRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ActionRowComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRowComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRowComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionRowComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponents

`func (o *ActionRowComponentForModalRequest) GetComponents() []TextInputComponentForModalRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ActionRowComponentForModalRequest) GetComponentsOk() (*[]TextInputComponentForModalRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ActionRowComponentForModalRequest) SetComponents(v []TextInputComponentForModalRequest)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


