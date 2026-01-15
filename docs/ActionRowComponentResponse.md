# ActionRowComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**Components** | [**[]ActionRowComponentResponseComponentsInner**](ActionRowComponentResponseComponentsInner.md) |  | 

## Methods

### NewActionRowComponentResponse

`func NewActionRowComponentResponse(type_ MessageComponentTypes, id int32, components []ActionRowComponentResponseComponentsInner, ) *ActionRowComponentResponse`

NewActionRowComponentResponse instantiates a new ActionRowComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowComponentResponseWithDefaults

`func NewActionRowComponentResponseWithDefaults() *ActionRowComponentResponse`

NewActionRowComponentResponseWithDefaults instantiates a new ActionRowComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRowComponentResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ActionRowComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRowComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRowComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *ActionRowComponentResponse) GetComponents() []ActionRowComponentResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ActionRowComponentResponse) GetComponentsOk() (*[]ActionRowComponentResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ActionRowComponentResponse) SetComponents(v []ActionRowComponentResponseComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


