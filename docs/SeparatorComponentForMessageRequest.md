# SeparatorComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Spacing** | Pointer to [**SeparatorComponentForMessageRequestSpacing**](SeparatorComponentForMessageRequestSpacing.md) |  | [optional] 
**Divider** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeparatorComponentForMessageRequest

`func NewSeparatorComponentForMessageRequest(type_ MessageComponentTypes, ) *SeparatorComponentForMessageRequest`

NewSeparatorComponentForMessageRequest instantiates a new SeparatorComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeparatorComponentForMessageRequestWithDefaults

`func NewSeparatorComponentForMessageRequestWithDefaults() *SeparatorComponentForMessageRequest`

NewSeparatorComponentForMessageRequestWithDefaults instantiates a new SeparatorComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SeparatorComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeparatorComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeparatorComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *SeparatorComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeparatorComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeparatorComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SeparatorComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpacing

`func (o *SeparatorComponentForMessageRequest) GetSpacing() SeparatorComponentForMessageRequestSpacing`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *SeparatorComponentForMessageRequest) GetSpacingOk() (*SeparatorComponentForMessageRequestSpacing, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *SeparatorComponentForMessageRequest) SetSpacing(v SeparatorComponentForMessageRequestSpacing)`

SetSpacing sets Spacing field to given value.

### HasSpacing

`func (o *SeparatorComponentForMessageRequest) HasSpacing() bool`

HasSpacing returns a boolean if a field has been set.

### GetDivider

`func (o *SeparatorComponentForMessageRequest) GetDivider() bool`

GetDivider returns the Divider field if non-nil, zero value otherwise.

### GetDividerOk

`func (o *SeparatorComponentForMessageRequest) GetDividerOk() (*bool, bool)`

GetDividerOk returns a tuple with the Divider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivider

`func (o *SeparatorComponentForMessageRequest) SetDivider(v bool)`

SetDivider sets Divider field to given value.

### HasDivider

`func (o *SeparatorComponentForMessageRequest) HasDivider() bool`

HasDivider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


