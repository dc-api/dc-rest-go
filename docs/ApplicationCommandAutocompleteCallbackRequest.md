# ApplicationCommandAutocompleteCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InteractionCallbackTypes**](InteractionCallbackTypes.md) |  | 
**Data** | [**ApplicationCommandAutocompleteCallbackRequestData**](ApplicationCommandAutocompleteCallbackRequestData.md) |  | 

## Methods

### NewApplicationCommandAutocompleteCallbackRequest

`func NewApplicationCommandAutocompleteCallbackRequest(type_ InteractionCallbackTypes, data ApplicationCommandAutocompleteCallbackRequestData, ) *ApplicationCommandAutocompleteCallbackRequest`

NewApplicationCommandAutocompleteCallbackRequest instantiates a new ApplicationCommandAutocompleteCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandAutocompleteCallbackRequestWithDefaults

`func NewApplicationCommandAutocompleteCallbackRequestWithDefaults() *ApplicationCommandAutocompleteCallbackRequest`

NewApplicationCommandAutocompleteCallbackRequestWithDefaults instantiates a new ApplicationCommandAutocompleteCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandAutocompleteCallbackRequest) GetType() InteractionCallbackTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandAutocompleteCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandAutocompleteCallbackRequest) SetType(v InteractionCallbackTypes)`

SetType sets Type field to given value.


### GetData

`func (o *ApplicationCommandAutocompleteCallbackRequest) GetData() ApplicationCommandAutocompleteCallbackRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationCommandAutocompleteCallbackRequest) GetDataOk() (*ApplicationCommandAutocompleteCallbackRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationCommandAutocompleteCallbackRequest) SetData(v ApplicationCommandAutocompleteCallbackRequestData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


