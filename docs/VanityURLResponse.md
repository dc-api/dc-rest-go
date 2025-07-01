# VanityURLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Uses** | **int32** |  | 
**Error** | Pointer to [**NullableVanityURLErrorResponse**](VanityURLErrorResponse.md) |  | [optional] 

## Methods

### NewVanityURLResponse

`func NewVanityURLResponse(uses int32, ) *VanityURLResponse`

NewVanityURLResponse instantiates a new VanityURLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVanityURLResponseWithDefaults

`func NewVanityURLResponseWithDefaults() *VanityURLResponse`

NewVanityURLResponseWithDefaults instantiates a new VanityURLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VanityURLResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VanityURLResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VanityURLResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VanityURLResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VanityURLResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VanityURLResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetUses

`func (o *VanityURLResponse) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *VanityURLResponse) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *VanityURLResponse) SetUses(v int32)`

SetUses sets Uses field to given value.


### GetError

`func (o *VanityURLResponse) GetError() VanityURLErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VanityURLResponse) GetErrorOk() (*VanityURLErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VanityURLResponse) SetError(v VanityURLErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *VanityURLResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *VanityURLResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *VanityURLResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


