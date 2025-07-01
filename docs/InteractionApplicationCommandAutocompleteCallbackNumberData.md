# InteractionApplicationCommandAutocompleteCallbackNumberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to [**[]ApplicationCommandOptionNumberChoice**](ApplicationCommandOptionNumberChoice.md) |  | [optional] 

## Methods

### NewInteractionApplicationCommandAutocompleteCallbackNumberData

`func NewInteractionApplicationCommandAutocompleteCallbackNumberData() *InteractionApplicationCommandAutocompleteCallbackNumberData`

NewInteractionApplicationCommandAutocompleteCallbackNumberData instantiates a new InteractionApplicationCommandAutocompleteCallbackNumberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionApplicationCommandAutocompleteCallbackNumberDataWithDefaults

`func NewInteractionApplicationCommandAutocompleteCallbackNumberDataWithDefaults() *InteractionApplicationCommandAutocompleteCallbackNumberData`

NewInteractionApplicationCommandAutocompleteCallbackNumberDataWithDefaults instantiates a new InteractionApplicationCommandAutocompleteCallbackNumberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) GetChoices() []*ApplicationCommandOptionNumberChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) GetChoicesOk() (*[]*ApplicationCommandOptionNumberChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) SetChoices(v []*ApplicationCommandOptionNumberChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


