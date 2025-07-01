# ModalInteractionCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**Data** | [**ModalInteractionCallbackRequestData**](ModalInteractionCallbackRequestData.md) |  | 

## Methods

### NewModalInteractionCallbackRequest

`func NewModalInteractionCallbackRequest(type_ NullableInt32, data ModalInteractionCallbackRequestData, ) *ModalInteractionCallbackRequest`

NewModalInteractionCallbackRequest instantiates a new ModalInteractionCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModalInteractionCallbackRequestWithDefaults

`func NewModalInteractionCallbackRequestWithDefaults() *ModalInteractionCallbackRequest`

NewModalInteractionCallbackRequestWithDefaults instantiates a new ModalInteractionCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModalInteractionCallbackRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModalInteractionCallbackRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModalInteractionCallbackRequest) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ModalInteractionCallbackRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModalInteractionCallbackRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *ModalInteractionCallbackRequest) GetData() ModalInteractionCallbackRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModalInteractionCallbackRequest) GetDataOk() (*ModalInteractionCallbackRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModalInteractionCallbackRequest) SetData(v ModalInteractionCallbackRequestData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


