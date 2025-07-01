# SectionComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Components** | [**[]TextDisplayComponentForMessageRequest**](TextDisplayComponentForMessageRequest.md) |  | 
**Accessory** | [**SectionComponentForMessageRequestAccessory**](SectionComponentForMessageRequestAccessory.md) |  | 

## Methods

### NewSectionComponentForMessageRequest

`func NewSectionComponentForMessageRequest(type_ int32, components []TextDisplayComponentForMessageRequest, accessory SectionComponentForMessageRequestAccessory, ) *SectionComponentForMessageRequest`

NewSectionComponentForMessageRequest instantiates a new SectionComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionComponentForMessageRequestWithDefaults

`func NewSectionComponentForMessageRequestWithDefaults() *SectionComponentForMessageRequest`

NewSectionComponentForMessageRequestWithDefaults instantiates a new SectionComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SectionComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectionComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectionComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetComponents

`func (o *SectionComponentForMessageRequest) GetComponents() []TextDisplayComponentForMessageRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SectionComponentForMessageRequest) GetComponentsOk() (*[]TextDisplayComponentForMessageRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SectionComponentForMessageRequest) SetComponents(v []TextDisplayComponentForMessageRequest)`

SetComponents sets Components field to given value.


### GetAccessory

`func (o *SectionComponentForMessageRequest) GetAccessory() SectionComponentForMessageRequestAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *SectionComponentForMessageRequest) GetAccessoryOk() (*SectionComponentForMessageRequestAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *SectionComponentForMessageRequest) SetAccessory(v SectionComponentForMessageRequestAccessory)`

SetAccessory sets Accessory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


