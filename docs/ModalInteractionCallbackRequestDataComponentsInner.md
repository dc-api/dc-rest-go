# ModalInteractionCallbackRequestDataComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Components** | [**[]TextInputComponentForModalRequest**](TextInputComponentForModalRequest.md) |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Component** | [**LabelComponentForModalRequestComponent**](LabelComponentForModalRequestComponent.md) |  | 
**Content** | **string** |  | 

## Methods

### NewModalInteractionCallbackRequestDataComponentsInner

`func NewModalInteractionCallbackRequestDataComponentsInner(type_ int32, components []TextInputComponentForModalRequest, label string, component LabelComponentForModalRequestComponent, content string, ) *ModalInteractionCallbackRequestDataComponentsInner`

NewModalInteractionCallbackRequestDataComponentsInner instantiates a new ModalInteractionCallbackRequestDataComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModalInteractionCallbackRequestDataComponentsInnerWithDefaults

`func NewModalInteractionCallbackRequestDataComponentsInnerWithDefaults() *ModalInteractionCallbackRequestDataComponentsInner`

NewModalInteractionCallbackRequestDataComponentsInnerWithDefaults instantiates a new ModalInteractionCallbackRequestDataComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModalInteractionCallbackRequestDataComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponents

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetComponents() []TextInputComponentForModalRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetComponentsOk() (*[]TextInputComponentForModalRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetComponents(v []TextInputComponentForModalRequest)`

SetComponents sets Components field to given value.


### GetLabel

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModalInteractionCallbackRequestDataComponentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponent

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetComponent() LabelComponentForModalRequestComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetComponentOk() (*LabelComponentForModalRequestComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetComponent(v LabelComponentForModalRequestComponent)`

SetComponent sets Component field to given value.


### GetContent

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModalInteractionCallbackRequestDataComponentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModalInteractionCallbackRequestDataComponentsInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


