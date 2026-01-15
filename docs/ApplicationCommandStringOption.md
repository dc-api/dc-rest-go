# ApplicationCommandStringOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ApplicationCommandOptionType**](ApplicationCommandOptionType.md) |  | 
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Autocomplete** | Pointer to **bool** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**Choices** | Pointer to [**[]ApplicationCommandOptionStringChoice**](ApplicationCommandOptionStringChoice.md) |  | [optional] 

## Methods

### NewApplicationCommandStringOption

`func NewApplicationCommandStringOption(type_ ApplicationCommandOptionType, name string, description string, ) *ApplicationCommandStringOption`

NewApplicationCommandStringOption instantiates a new ApplicationCommandStringOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandStringOptionWithDefaults

`func NewApplicationCommandStringOptionWithDefaults() *ApplicationCommandStringOption`

NewApplicationCommandStringOptionWithDefaults instantiates a new ApplicationCommandStringOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandStringOption) GetType() ApplicationCommandOptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandStringOption) GetTypeOk() (*ApplicationCommandOptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandStringOption) SetType(v ApplicationCommandOptionType)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandStringOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandStringOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandStringOption) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandStringOption) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandStringOption) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandStringOption) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandStringOption) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandStringOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandStringOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandStringOption) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationCommandStringOption) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandStringOption) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandStringOption) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandStringOption) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandStringOption) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandStringOption) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandStringOption) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandStringOption) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetAutocomplete

`func (o *ApplicationCommandStringOption) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandStringOption) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandStringOption) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandStringOption) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### GetMinLength

`func (o *ApplicationCommandStringOption) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ApplicationCommandStringOption) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ApplicationCommandStringOption) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ApplicationCommandStringOption) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *ApplicationCommandStringOption) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ApplicationCommandStringOption) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ApplicationCommandStringOption) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ApplicationCommandStringOption) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetChoices

`func (o *ApplicationCommandStringOption) GetChoices() []ApplicationCommandOptionStringChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandStringOption) GetChoicesOk() (*[]ApplicationCommandOptionStringChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandStringOption) SetChoices(v []ApplicationCommandOptionStringChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandStringOption) HasChoices() bool`

HasChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


