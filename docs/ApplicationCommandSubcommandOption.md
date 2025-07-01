# ApplicationCommandSubcommandOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Options** | Pointer to [**[]ApplicationCommandSubcommandOptionOptionsInner**](ApplicationCommandSubcommandOptionOptionsInner.md) |  | [optional] 

## Methods

### NewApplicationCommandSubcommandOption

`func NewApplicationCommandSubcommandOption(type_ int32, name string, description string, ) *ApplicationCommandSubcommandOption`

NewApplicationCommandSubcommandOption instantiates a new ApplicationCommandSubcommandOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandSubcommandOptionWithDefaults

`func NewApplicationCommandSubcommandOptionWithDefaults() *ApplicationCommandSubcommandOption`

NewApplicationCommandSubcommandOptionWithDefaults instantiates a new ApplicationCommandSubcommandOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandSubcommandOption) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandSubcommandOption) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandSubcommandOption) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandSubcommandOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandSubcommandOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandSubcommandOption) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandSubcommandOption) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandSubcommandOption) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandSubcommandOption) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandSubcommandOption) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandSubcommandOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandSubcommandOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandSubcommandOption) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOption) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandSubcommandOption) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOption) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOption) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandSubcommandOption) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandSubcommandOption) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandSubcommandOption) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandSubcommandOption) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandSubcommandOption) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandSubcommandOption) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetOptions

`func (o *ApplicationCommandSubcommandOption) GetOptions() []ApplicationCommandSubcommandOptionOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandSubcommandOption) GetOptionsOk() (*[]ApplicationCommandSubcommandOptionOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandSubcommandOption) SetOptions(v []ApplicationCommandSubcommandOptionOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandSubcommandOption) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApplicationCommandSubcommandOption) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandSubcommandOption) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


