# ApplicationCommandIntegerOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Name** | **string** |  | 
**NameLocalized** | Pointer to **NullableString** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalized** | Pointer to **NullableString** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Autocomplete** | Pointer to **NullableBool** |  | [optional] 
**Choices** | Pointer to [**[]ApplicationCommandOptionIntegerChoiceResponse**](ApplicationCommandOptionIntegerChoiceResponse.md) |  | [optional] 
**MinValue** | Pointer to **int64** |  | [optional] 
**MaxValue** | Pointer to **int64** |  | [optional] 

## Methods

### NewApplicationCommandIntegerOptionResponse

`func NewApplicationCommandIntegerOptionResponse(type_ int32, name string, description string, ) *ApplicationCommandIntegerOptionResponse`

NewApplicationCommandIntegerOptionResponse instantiates a new ApplicationCommandIntegerOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandIntegerOptionResponseWithDefaults

`func NewApplicationCommandIntegerOptionResponseWithDefaults() *ApplicationCommandIntegerOptionResponse`

NewApplicationCommandIntegerOptionResponseWithDefaults instantiates a new ApplicationCommandIntegerOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandIntegerOptionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandIntegerOptionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandIntegerOptionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandIntegerOptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandIntegerOptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandIntegerOptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandIntegerOptionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandIntegerOptionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandIntegerOptionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandIntegerOptionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandIntegerOptionResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandIntegerOptionResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandIntegerOptionResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandIntegerOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandIntegerOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandIntegerOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandIntegerOptionResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandIntegerOptionResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandIntegerOptionResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandIntegerOptionResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### SetDescriptionLocalizedNil

`func (o *ApplicationCommandIntegerOptionResponse) SetDescriptionLocalizedNil(b bool)`

 SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil

### UnsetDescriptionLocalized
`func (o *ApplicationCommandIntegerOptionResponse) UnsetDescriptionLocalized()`

UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandIntegerOptionResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandIntegerOptionResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandIntegerOptionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandIntegerOptionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandIntegerOptionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandIntegerOptionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandIntegerOptionResponse) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandIntegerOptionResponse) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetAutocomplete

`func (o *ApplicationCommandIntegerOptionResponse) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandIntegerOptionResponse) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandIntegerOptionResponse) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandIntegerOptionResponse) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### SetAutocompleteNil

`func (o *ApplicationCommandIntegerOptionResponse) SetAutocompleteNil(b bool)`

 SetAutocompleteNil sets the value for Autocomplete to be an explicit nil

### UnsetAutocomplete
`func (o *ApplicationCommandIntegerOptionResponse) UnsetAutocomplete()`

UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
### GetChoices

`func (o *ApplicationCommandIntegerOptionResponse) GetChoices() []ApplicationCommandOptionIntegerChoiceResponse`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandIntegerOptionResponse) GetChoicesOk() (*[]ApplicationCommandOptionIntegerChoiceResponse, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandIntegerOptionResponse) SetChoices(v []ApplicationCommandOptionIntegerChoiceResponse)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandIntegerOptionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *ApplicationCommandIntegerOptionResponse) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *ApplicationCommandIntegerOptionResponse) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetMinValue

`func (o *ApplicationCommandIntegerOptionResponse) GetMinValue() int64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ApplicationCommandIntegerOptionResponse) GetMinValueOk() (*int64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ApplicationCommandIntegerOptionResponse) SetMinValue(v int64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ApplicationCommandIntegerOptionResponse) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ApplicationCommandIntegerOptionResponse) GetMaxValue() int64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ApplicationCommandIntegerOptionResponse) GetMaxValueOk() (*int64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ApplicationCommandIntegerOptionResponse) SetMaxValue(v int64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ApplicationCommandIntegerOptionResponse) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


