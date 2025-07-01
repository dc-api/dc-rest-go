# ApplicationCommandSubcommandOptionResponse

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
**Options** | Pointer to [**[]ApplicationCommandSubcommandOptionResponseOptionsInner**](ApplicationCommandSubcommandOptionResponseOptionsInner.md) |  | [optional] 

## Methods

### NewApplicationCommandSubcommandOptionResponse

`func NewApplicationCommandSubcommandOptionResponse(type_ int32, name string, description string, ) *ApplicationCommandSubcommandOptionResponse`

NewApplicationCommandSubcommandOptionResponse instantiates a new ApplicationCommandSubcommandOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandSubcommandOptionResponseWithDefaults

`func NewApplicationCommandSubcommandOptionResponseWithDefaults() *ApplicationCommandSubcommandOptionResponse`

NewApplicationCommandSubcommandOptionResponseWithDefaults instantiates a new ApplicationCommandSubcommandOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandSubcommandOptionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandSubcommandOptionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandSubcommandOptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandSubcommandOptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandSubcommandOptionResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandSubcommandOptionResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandSubcommandOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandSubcommandOptionResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### SetDescriptionLocalizedNil

`func (o *ApplicationCommandSubcommandOptionResponse) SetDescriptionLocalizedNil(b bool)`

 SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil

### UnsetDescriptionLocalized
`func (o *ApplicationCommandSubcommandOptionResponse) UnsetDescriptionLocalized()`

UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandSubcommandOptionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandSubcommandOptionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandSubcommandOptionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandSubcommandOptionResponse) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandSubcommandOptionResponse) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetOptions

`func (o *ApplicationCommandSubcommandOptionResponse) GetOptions() []ApplicationCommandSubcommandOptionResponseOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandSubcommandOptionResponse) GetOptionsOk() (*[]ApplicationCommandSubcommandOptionResponseOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandSubcommandOptionResponse) SetOptions(v []ApplicationCommandSubcommandOptionResponseOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandSubcommandOptionResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApplicationCommandSubcommandOptionResponse) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandSubcommandOptionResponse) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


