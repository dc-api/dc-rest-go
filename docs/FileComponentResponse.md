# FileComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**File** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Spoiler** | **bool** |  | 

## Methods

### NewFileComponentResponse

`func NewFileComponentResponse(type_ int32, id int32, file UnfurledMediaResponse, spoiler bool, ) *FileComponentResponse`

NewFileComponentResponse instantiates a new FileComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileComponentResponseWithDefaults

`func NewFileComponentResponseWithDefaults() *FileComponentResponse`

NewFileComponentResponseWithDefaults instantiates a new FileComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FileComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *FileComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetFile

`func (o *FileComponentResponse) GetFile() UnfurledMediaResponse`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileComponentResponse) GetFileOk() (*UnfurledMediaResponse, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileComponentResponse) SetFile(v UnfurledMediaResponse)`

SetFile sets File field to given value.


### GetName

`func (o *FileComponentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileComponentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileComponentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileComponentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileComponentResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileComponentResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *FileComponentResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileComponentResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileComponentResponse) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileComponentResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *FileComponentResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *FileComponentResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSpoiler

`func (o *FileComponentResponse) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *FileComponentResponse) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *FileComponentResponse) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


