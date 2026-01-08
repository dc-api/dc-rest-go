# TextDisplayComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**Content** | **string** |  | 

## Methods

### NewTextDisplayComponentForModalRequest

`func NewTextDisplayComponentForModalRequest(type_ MessageComponentTypes, content string, ) *TextDisplayComponentForModalRequest`

NewTextDisplayComponentForModalRequest instantiates a new TextDisplayComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextDisplayComponentForModalRequestWithDefaults

`func NewTextDisplayComponentForModalRequestWithDefaults() *TextDisplayComponentForModalRequest`

NewTextDisplayComponentForModalRequestWithDefaults instantiates a new TextDisplayComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextDisplayComponentForModalRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextDisplayComponentForModalRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextDisplayComponentForModalRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *TextDisplayComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextDisplayComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextDisplayComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TextDisplayComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *TextDisplayComponentForModalRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TextDisplayComponentForModalRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TextDisplayComponentForModalRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


