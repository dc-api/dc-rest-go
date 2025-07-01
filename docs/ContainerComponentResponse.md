# ContainerComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**AccentColor** | Pointer to **NullableInt32** |  | [optional] 
**Components** | [**[]ContainerComponentResponseComponentsInner**](ContainerComponentResponseComponentsInner.md) |  | 
**Spoiler** | **bool** |  | 

## Methods

### NewContainerComponentResponse

`func NewContainerComponentResponse(type_ int32, id int32, components []ContainerComponentResponseComponentsInner, spoiler bool, ) *ContainerComponentResponse`

NewContainerComponentResponse instantiates a new ContainerComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerComponentResponseWithDefaults

`func NewContainerComponentResponseWithDefaults() *ContainerComponentResponse`

NewContainerComponentResponseWithDefaults instantiates a new ContainerComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ContainerComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetAccentColor

`func (o *ContainerComponentResponse) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *ContainerComponentResponse) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *ContainerComponentResponse) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *ContainerComponentResponse) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### SetAccentColorNil

`func (o *ContainerComponentResponse) SetAccentColorNil(b bool)`

 SetAccentColorNil sets the value for AccentColor to be an explicit nil

### UnsetAccentColor
`func (o *ContainerComponentResponse) UnsetAccentColor()`

UnsetAccentColor ensures that no value is present for AccentColor, not even an explicit nil
### GetComponents

`func (o *ContainerComponentResponse) GetComponents() []ContainerComponentResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ContainerComponentResponse) GetComponentsOk() (*[]ContainerComponentResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ContainerComponentResponse) SetComponents(v []ContainerComponentResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetSpoiler

`func (o *ContainerComponentResponse) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ContainerComponentResponse) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ContainerComponentResponse) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


