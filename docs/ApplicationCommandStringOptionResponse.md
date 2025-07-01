# ApplicationCommandStringOptionResponse

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
**Choices** | Pointer to [**[]ApplicationCommandOptionStringChoiceResponse**](ApplicationCommandOptionStringChoiceResponse.md) |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewApplicationCommandStringOptionResponse

`func NewApplicationCommandStringOptionResponse(type_ int32, name string, description string, ) *ApplicationCommandStringOptionResponse`

NewApplicationCommandStringOptionResponse instantiates a new ApplicationCommandStringOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandStringOptionResponseWithDefaults

`func NewApplicationCommandStringOptionResponseWithDefaults() *ApplicationCommandStringOptionResponse`

NewApplicationCommandStringOptionResponseWithDefaults instantiates a new ApplicationCommandStringOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandStringOptionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandStringOptionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandStringOptionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandStringOptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandStringOptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandStringOptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandStringOptionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandStringOptionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandStringOptionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandStringOptionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandStringOptionResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandStringOptionResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandStringOptionResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandStringOptionResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandStringOptionResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandStringOptionResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandStringOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandStringOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandStringOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandStringOptionResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandStringOptionResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandStringOptionResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandStringOptionResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### SetDescriptionLocalizedNil

`func (o *ApplicationCommandStringOptionResponse) SetDescriptionLocalizedNil(b bool)`

 SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil

### UnsetDescriptionLocalized
`func (o *ApplicationCommandStringOptionResponse) UnsetDescriptionLocalized()`

UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandStringOptionResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandStringOptionResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandStringOptionResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandStringOptionResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandStringOptionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandStringOptionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandStringOptionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandStringOptionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandStringOptionResponse) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandStringOptionResponse) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetAutocomplete

`func (o *ApplicationCommandStringOptionResponse) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandStringOptionResponse) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandStringOptionResponse) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandStringOptionResponse) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### SetAutocompleteNil

`func (o *ApplicationCommandStringOptionResponse) SetAutocompleteNil(b bool)`

 SetAutocompleteNil sets the value for Autocomplete to be an explicit nil

### UnsetAutocomplete
`func (o *ApplicationCommandStringOptionResponse) UnsetAutocomplete()`

UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
### GetChoices

`func (o *ApplicationCommandStringOptionResponse) GetChoices() []ApplicationCommandOptionStringChoiceResponse`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandStringOptionResponse) GetChoicesOk() (*[]ApplicationCommandOptionStringChoiceResponse, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandStringOptionResponse) SetChoices(v []ApplicationCommandOptionStringChoiceResponse)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandStringOptionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *ApplicationCommandStringOptionResponse) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *ApplicationCommandStringOptionResponse) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetMinLength

`func (o *ApplicationCommandStringOptionResponse) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ApplicationCommandStringOptionResponse) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ApplicationCommandStringOptionResponse) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ApplicationCommandStringOptionResponse) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ApplicationCommandStringOptionResponse) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ApplicationCommandStringOptionResponse) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *ApplicationCommandStringOptionResponse) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ApplicationCommandStringOptionResponse) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ApplicationCommandStringOptionResponse) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ApplicationCommandStringOptionResponse) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ApplicationCommandStringOptionResponse) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ApplicationCommandStringOptionResponse) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


