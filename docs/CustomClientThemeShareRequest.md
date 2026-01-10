# CustomClientThemeShareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colors** | **[]string** |  | 
**GradientAngle** | **int32** |  | 
**BaseMix** | **int32** |  | 
**BaseTheme** | Pointer to [**CustomClientThemeShareRequestBaseTheme**](CustomClientThemeShareRequestBaseTheme.md) |  | [optional] 

## Methods

### NewCustomClientThemeShareRequest

`func NewCustomClientThemeShareRequest(colors []string, gradientAngle int32, baseMix int32, ) *CustomClientThemeShareRequest`

NewCustomClientThemeShareRequest instantiates a new CustomClientThemeShareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomClientThemeShareRequestWithDefaults

`func NewCustomClientThemeShareRequestWithDefaults() *CustomClientThemeShareRequest`

NewCustomClientThemeShareRequestWithDefaults instantiates a new CustomClientThemeShareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColors

`func (o *CustomClientThemeShareRequest) GetColors() []string`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *CustomClientThemeShareRequest) GetColorsOk() (*[]string, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *CustomClientThemeShareRequest) SetColors(v []string)`

SetColors sets Colors field to given value.


### GetGradientAngle

`func (o *CustomClientThemeShareRequest) GetGradientAngle() int32`

GetGradientAngle returns the GradientAngle field if non-nil, zero value otherwise.

### GetGradientAngleOk

`func (o *CustomClientThemeShareRequest) GetGradientAngleOk() (*int32, bool)`

GetGradientAngleOk returns a tuple with the GradientAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradientAngle

`func (o *CustomClientThemeShareRequest) SetGradientAngle(v int32)`

SetGradientAngle sets GradientAngle field to given value.


### GetBaseMix

`func (o *CustomClientThemeShareRequest) GetBaseMix() int32`

GetBaseMix returns the BaseMix field if non-nil, zero value otherwise.

### GetBaseMixOk

`func (o *CustomClientThemeShareRequest) GetBaseMixOk() (*int32, bool)`

GetBaseMixOk returns a tuple with the BaseMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMix

`func (o *CustomClientThemeShareRequest) SetBaseMix(v int32)`

SetBaseMix sets BaseMix field to given value.


### GetBaseTheme

`func (o *CustomClientThemeShareRequest) GetBaseTheme() CustomClientThemeShareRequestBaseTheme`

GetBaseTheme returns the BaseTheme field if non-nil, zero value otherwise.

### GetBaseThemeOk

`func (o *CustomClientThemeShareRequest) GetBaseThemeOk() (*CustomClientThemeShareRequestBaseTheme, bool)`

GetBaseThemeOk returns a tuple with the BaseTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTheme

`func (o *CustomClientThemeShareRequest) SetBaseTheme(v CustomClientThemeShareRequestBaseTheme)`

SetBaseTheme sets BaseTheme field to given value.

### HasBaseTheme

`func (o *CustomClientThemeShareRequest) HasBaseTheme() bool`

HasBaseTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


