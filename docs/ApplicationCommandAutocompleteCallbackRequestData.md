# ApplicationCommandAutocompleteCallbackRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to [**[]ApplicationCommandOptionStringChoice**](ApplicationCommandOptionStringChoice.md) |  | [optional] 

## Methods

### NewApplicationCommandAutocompleteCallbackRequestData

`func NewApplicationCommandAutocompleteCallbackRequestData() *ApplicationCommandAutocompleteCallbackRequestData`

NewApplicationCommandAutocompleteCallbackRequestData instantiates a new ApplicationCommandAutocompleteCallbackRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandAutocompleteCallbackRequestDataWithDefaults

`func NewApplicationCommandAutocompleteCallbackRequestDataWithDefaults() *ApplicationCommandAutocompleteCallbackRequestData`

NewApplicationCommandAutocompleteCallbackRequestDataWithDefaults instantiates a new ApplicationCommandAutocompleteCallbackRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *ApplicationCommandAutocompleteCallbackRequestData) GetChoices() []ApplicationCommandOptionStringChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandAutocompleteCallbackRequestData) GetChoicesOk() (*[]ApplicationCommandOptionStringChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandAutocompleteCallbackRequestData) SetChoices(v []ApplicationCommandOptionStringChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandAutocompleteCallbackRequestData) HasChoices() bool`

HasChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


