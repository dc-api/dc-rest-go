# ModalInteractionCallbackRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomId** | **string** |  | 
**Title** | **string** |  | 
**Components** | [**[]ActionRowComponentForModalRequest**](ActionRowComponentForModalRequest.md) |  | 

## Methods

### NewModalInteractionCallbackRequestData

`func NewModalInteractionCallbackRequestData(customId string, title string, components []ActionRowComponentForModalRequest, ) *ModalInteractionCallbackRequestData`

NewModalInteractionCallbackRequestData instantiates a new ModalInteractionCallbackRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModalInteractionCallbackRequestDataWithDefaults

`func NewModalInteractionCallbackRequestDataWithDefaults() *ModalInteractionCallbackRequestData`

NewModalInteractionCallbackRequestDataWithDefaults instantiates a new ModalInteractionCallbackRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomId

`func (o *ModalInteractionCallbackRequestData) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ModalInteractionCallbackRequestData) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ModalInteractionCallbackRequestData) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetTitle

`func (o *ModalInteractionCallbackRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModalInteractionCallbackRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModalInteractionCallbackRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetComponents

`func (o *ModalInteractionCallbackRequestData) GetComponents() []ActionRowComponentForModalRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ModalInteractionCallbackRequestData) GetComponentsOk() (*[]ActionRowComponentForModalRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ModalInteractionCallbackRequestData) SetComponents(v []ActionRowComponentForModalRequest)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


