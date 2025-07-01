# ApplicationCommandOptionIntegerChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Value** | **int64** |  | 

## Methods

### NewApplicationCommandOptionIntegerChoice

`func NewApplicationCommandOptionIntegerChoice(name string, value int64, ) *ApplicationCommandOptionIntegerChoice`

NewApplicationCommandOptionIntegerChoice instantiates a new ApplicationCommandOptionIntegerChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandOptionIntegerChoiceWithDefaults

`func NewApplicationCommandOptionIntegerChoiceWithDefaults() *ApplicationCommandOptionIntegerChoice`

NewApplicationCommandOptionIntegerChoiceWithDefaults instantiates a new ApplicationCommandOptionIntegerChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandOptionIntegerChoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandOptionIntegerChoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandOptionIntegerChoice) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoice) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandOptionIntegerChoice) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoice) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoice) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationCommandOptionIntegerChoice) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationCommandOptionIntegerChoice) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationCommandOptionIntegerChoice) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


