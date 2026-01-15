# ContainerComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**AccentColor** | Pointer to **int32** |  | [optional] 
**Components** | [**[]ContainerComponentForMessageRequestComponentsInner**](ContainerComponentForMessageRequestComponentsInner.md) |  | 
**Spoiler** | Pointer to **bool** |  | [optional] 

## Methods

### NewContainerComponentForMessageRequest

`func NewContainerComponentForMessageRequest(type_ MessageComponentTypes, components []ContainerComponentForMessageRequestComponentsInner, ) *ContainerComponentForMessageRequest`

NewContainerComponentForMessageRequest instantiates a new ContainerComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerComponentForMessageRequestWithDefaults

`func NewContainerComponentForMessageRequestWithDefaults() *ContainerComponentForMessageRequest`

NewContainerComponentForMessageRequestWithDefaults instantiates a new ContainerComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ContainerComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccentColor

`func (o *ContainerComponentForMessageRequest) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *ContainerComponentForMessageRequest) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *ContainerComponentForMessageRequest) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *ContainerComponentForMessageRequest) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### GetComponents

`func (o *ContainerComponentForMessageRequest) GetComponents() []ContainerComponentForMessageRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ContainerComponentForMessageRequest) GetComponentsOk() (*[]ContainerComponentForMessageRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ContainerComponentForMessageRequest) SetComponents(v []ContainerComponentForMessageRequestComponentsInner)`

SetComponents sets Components field to given value.


### GetSpoiler

`func (o *ContainerComponentForMessageRequest) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ContainerComponentForMessageRequest) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ContainerComponentForMessageRequest) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *ContainerComponentForMessageRequest) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


