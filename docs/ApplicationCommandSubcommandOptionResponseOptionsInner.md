# ApplicationCommandSubcommandOptionResponseOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Name** | **string** |  | 
**NameLocalized** | Pointer to **string** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalized** | Pointer to **string** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**ChannelTypes** | Pointer to **[]int32** |  | [optional] 
**Autocomplete** | Pointer to **bool** |  | [optional] 
**Choices** | Pointer to [**[]ApplicationCommandOptionStringChoiceResponse**](ApplicationCommandOptionStringChoiceResponse.md) |  | [optional] 
**MinValue** | Pointer to **float64** |  | [optional] 
**MaxValue** | Pointer to **float64** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewApplicationCommandSubcommandOptionResponseOptionsInner

`func NewApplicationCommandSubcommandOptionResponseOptionsInner(type_ int32, name string, description string, ) *ApplicationCommandSubcommandOptionResponseOptionsInner`

NewApplicationCommandSubcommandOptionResponseOptionsInner instantiates a new ApplicationCommandSubcommandOptionResponseOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandSubcommandOptionResponseOptionsInnerWithDefaults

`func NewApplicationCommandSubcommandOptionResponseOptionsInnerWithDefaults() *ApplicationCommandSubcommandOptionResponseOptionsInner`

NewApplicationCommandSubcommandOptionResponseOptionsInnerWithDefaults instantiates a new ApplicationCommandSubcommandOptionResponseOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### GetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### GetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetChannelTypes

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetChannelTypes() []int32`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetChannelTypesOk() (*[]int32, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetChannelTypes(v []int32)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### GetAutocomplete

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### GetChoices

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetChoices() []ApplicationCommandOptionStringChoiceResponse`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetChoicesOk() (*[]ApplicationCommandOptionStringChoiceResponse, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetChoices(v []ApplicationCommandOptionStringChoiceResponse)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetMinValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ApplicationCommandSubcommandOptionResponseOptionsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


