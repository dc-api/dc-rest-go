# ApplicationCommandNumberOptionResponse

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
**Autocomplete** | Pointer to **bool** |  | [optional] 
**Choices** | Pointer to [**[]ApplicationCommandOptionNumberChoiceResponse**](ApplicationCommandOptionNumberChoiceResponse.md) |  | [optional] 
**MinValue** | Pointer to **float64** |  | [optional] 
**MaxValue** | Pointer to **float64** |  | [optional] 

## Methods

### NewApplicationCommandNumberOptionResponse

`func NewApplicationCommandNumberOptionResponse(type_ int32, name string, description string, ) *ApplicationCommandNumberOptionResponse`

NewApplicationCommandNumberOptionResponse instantiates a new ApplicationCommandNumberOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandNumberOptionResponseWithDefaults

`func NewApplicationCommandNumberOptionResponseWithDefaults() *ApplicationCommandNumberOptionResponse`

NewApplicationCommandNumberOptionResponseWithDefaults instantiates a new ApplicationCommandNumberOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandNumberOptionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandNumberOptionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandNumberOptionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandNumberOptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandNumberOptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandNumberOptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandNumberOptionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandNumberOptionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandNumberOptionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandNumberOptionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### GetNameLocalizations

`func (o *ApplicationCommandNumberOptionResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandNumberOptionResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandNumberOptionResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandNumberOptionResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandNumberOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandNumberOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandNumberOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandNumberOptionResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandNumberOptionResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandNumberOptionResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandNumberOptionResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### GetDescriptionLocalizations

`func (o *ApplicationCommandNumberOptionResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandNumberOptionResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandNumberOptionResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandNumberOptionResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandNumberOptionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandNumberOptionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandNumberOptionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandNumberOptionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetAutocomplete

`func (o *ApplicationCommandNumberOptionResponse) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandNumberOptionResponse) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandNumberOptionResponse) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandNumberOptionResponse) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### GetChoices

`func (o *ApplicationCommandNumberOptionResponse) GetChoices() []ApplicationCommandOptionNumberChoiceResponse`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandNumberOptionResponse) GetChoicesOk() (*[]ApplicationCommandOptionNumberChoiceResponse, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandNumberOptionResponse) SetChoices(v []ApplicationCommandOptionNumberChoiceResponse)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandNumberOptionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetMinValue

`func (o *ApplicationCommandNumberOptionResponse) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ApplicationCommandNumberOptionResponse) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ApplicationCommandNumberOptionResponse) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ApplicationCommandNumberOptionResponse) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ApplicationCommandNumberOptionResponse) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ApplicationCommandNumberOptionResponse) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ApplicationCommandNumberOptionResponse) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ApplicationCommandNumberOptionResponse) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


