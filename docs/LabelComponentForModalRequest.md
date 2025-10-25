# LabelComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Label** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Component** | [**LabelComponentForModalRequestComponent**](LabelComponentForModalRequestComponent.md) |  | 

## Methods

### NewLabelComponentForModalRequest

`func NewLabelComponentForModalRequest(type_ int32, label string, component LabelComponentForModalRequestComponent, ) *LabelComponentForModalRequest`

NewLabelComponentForModalRequest instantiates a new LabelComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelComponentForModalRequestWithDefaults

`func NewLabelComponentForModalRequestWithDefaults() *LabelComponentForModalRequest`

NewLabelComponentForModalRequestWithDefaults instantiates a new LabelComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LabelComponentForModalRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelComponentForModalRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelComponentForModalRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *LabelComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LabelComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LabelComponentForModalRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LabelComponentForModalRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *LabelComponentForModalRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabelComponentForModalRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabelComponentForModalRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *LabelComponentForModalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelComponentForModalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelComponentForModalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LabelComponentForModalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LabelComponentForModalRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LabelComponentForModalRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetComponent

`func (o *LabelComponentForModalRequest) GetComponent() LabelComponentForModalRequestComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *LabelComponentForModalRequest) GetComponentOk() (*LabelComponentForModalRequestComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *LabelComponentForModalRequest) SetComponent(v LabelComponentForModalRequestComponent)`

SetComponent sets Component field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


