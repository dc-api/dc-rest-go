# ContainerComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**AccentColor** | Pointer to **NullableInt32** |  | [optional] 
**Components** | [**[]ContainerComponentForMessageRequestComponentsInner**](ContainerComponentForMessageRequestComponentsInner.md) |  | 
**Spoiler** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewContainerComponentForMessageRequest

`func NewContainerComponentForMessageRequest(type_ int32, components []ContainerComponentForMessageRequestComponentsInner, ) *ContainerComponentForMessageRequest`

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

`func (o *ContainerComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


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

### SetAccentColorNil

`func (o *ContainerComponentForMessageRequest) SetAccentColorNil(b bool)`

 SetAccentColorNil sets the value for AccentColor to be an explicit nil

### UnsetAccentColor
`func (o *ContainerComponentForMessageRequest) UnsetAccentColor()`

UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
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

### SetSpoilerNil

`func (o *ContainerComponentForMessageRequest) SetSpoilerNil(b bool)`

 SetSpoilerNil sets the value for Spoiler to be an explicit nil

### UnsetSpoiler
`func (o *ContainerComponentForMessageRequest) UnsetSpoiler()`

UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


