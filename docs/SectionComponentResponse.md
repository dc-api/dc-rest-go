# SectionComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**Components** | [**[]TextDisplayComponentResponse**](TextDisplayComponentResponse.md) |  | 
**Accessory** | [**SectionComponentResponseAccessory**](SectionComponentResponseAccessory.md) |  | 

## Methods

### NewSectionComponentResponse

`func NewSectionComponentResponse(type_ int32, id int32, components []TextDisplayComponentResponse, accessory SectionComponentResponseAccessory, ) *SectionComponentResponse`

NewSectionComponentResponse instantiates a new SectionComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionComponentResponseWithDefaults

`func NewSectionComponentResponseWithDefaults() *SectionComponentResponse`

NewSectionComponentResponseWithDefaults instantiates a new SectionComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SectionComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectionComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectionComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *SectionComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *SectionComponentResponse) GetComponents() []TextDisplayComponentResponse`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SectionComponentResponse) GetComponentsOk() (*[]TextDisplayComponentResponse, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SectionComponentResponse) SetComponents(v []TextDisplayComponentResponse)`

SetComponents sets Components field to given value.


### GetAccessory

`func (o *SectionComponentResponse) GetAccessory() SectionComponentResponseAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *SectionComponentResponse) GetAccessoryOk() (*SectionComponentResponseAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *SectionComponentResponse) SetAccessory(v SectionComponentResponseAccessory)`

SetAccessory sets Accessory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


