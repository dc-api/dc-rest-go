# SectionComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Components** | [**[]TextDisplayComponentForMessageRequest**](TextDisplayComponentForMessageRequest.md) |  | 
**Accessory** | [**SectionComponentForMessageRequestAccessory**](SectionComponentForMessageRequestAccessory.md) |  | 

## Methods

### NewSectionComponentForMessageRequest

`func NewSectionComponentForMessageRequest(type_ MessageComponentTypes, components []TextDisplayComponentForMessageRequest, accessory SectionComponentForMessageRequestAccessory, ) *SectionComponentForMessageRequest`

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

`func (o *SectionComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectionComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectionComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *SectionComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SectionComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

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


