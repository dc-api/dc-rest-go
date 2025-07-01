# StandardStickerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Tags** | **string** |  | 
**Type** | **int32** |  | 
**FormatType** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PackId** | **string** |  | 
**SortValue** | **int32** |  | 

## Methods

### NewStandardStickerResponse

`func NewStandardStickerResponse(id string, name string, tags string, type_ int32, packId string, sortValue int32, ) *StandardStickerResponse`

NewStandardStickerResponse instantiates a new StandardStickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardStickerResponseWithDefaults

`func NewStandardStickerResponseWithDefaults() *StandardStickerResponse`

NewStandardStickerResponseWithDefaults instantiates a new StandardStickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandardStickerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandardStickerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandardStickerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StandardStickerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandardStickerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandardStickerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *StandardStickerResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StandardStickerResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StandardStickerResponse) SetTags(v string)`

SetTags sets Tags field to given value.


### GetType

`func (o *StandardStickerResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StandardStickerResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StandardStickerResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetFormatType

`func (o *StandardStickerResponse) GetFormatType() int32`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *StandardStickerResponse) GetFormatTypeOk() (*int32, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *StandardStickerResponse) SetFormatType(v int32)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *StandardStickerResponse) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.

### SetFormatTypeNil

`func (o *StandardStickerResponse) SetFormatTypeNil(b bool)`

 SetFormatTypeNil sets the value for FormatType to be an explicit nil

### UnsetFormatType
`func (o *StandardStickerResponse) UnsetFormatType()`

UnsetFormatType ensures that no value is present for FormatType, not even an explicit nil
### GetDescription

`func (o *StandardStickerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StandardStickerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StandardStickerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StandardStickerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StandardStickerResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StandardStickerResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPackId

`func (o *StandardStickerResponse) GetPackId() string`

GetPackId returns the PackId field if non-nil, zero value otherwise.

### GetPackIdOk

`func (o *StandardStickerResponse) GetPackIdOk() (*string, bool)`

GetPackIdOk returns a tuple with the PackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackId

`func (o *StandardStickerResponse) SetPackId(v string)`

SetPackId sets PackId field to given value.


### GetSortValue

`func (o *StandardStickerResponse) GetSortValue() int32`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *StandardStickerResponse) GetSortValueOk() (*int32, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *StandardStickerResponse) SetSortValue(v int32)`

SetSortValue sets SortValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


