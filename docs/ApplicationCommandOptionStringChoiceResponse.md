# ApplicationCommandOptionStringChoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NameLocalized** | Pointer to **NullableString** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewApplicationCommandOptionStringChoiceResponse

`func NewApplicationCommandOptionStringChoiceResponse(name string, value string, ) *ApplicationCommandOptionStringChoiceResponse`

NewApplicationCommandOptionStringChoiceResponse instantiates a new ApplicationCommandOptionStringChoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandOptionStringChoiceResponseWithDefaults

`func NewApplicationCommandOptionStringChoiceResponseWithDefaults() *ApplicationCommandOptionStringChoiceResponse`

NewApplicationCommandOptionStringChoiceResponseWithDefaults instantiates a new ApplicationCommandOptionStringChoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandOptionStringChoiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandOptionStringChoiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandOptionStringChoiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandOptionStringChoiceResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandOptionStringChoiceResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandOptionStringChoiceResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandOptionStringChoiceResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandOptionStringChoiceResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandOptionStringChoiceResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandOptionStringChoiceResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandOptionStringChoiceResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandOptionStringChoiceResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandOptionStringChoiceResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationCommandOptionStringChoiceResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationCommandOptionStringChoiceResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationCommandOptionStringChoiceResponse) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


