# BasicMessageResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**Components** | [**[]TextDisplayComponentResponse**](TextDisplayComponentResponse.md) |  | 
**AccentColor** | Pointer to **int32** |  | [optional] 
**Spoiler** | **bool** |  | 
**File** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Items** | [**[]MediaGalleryItemResponse**](MediaGalleryItemResponse.md) |  | 
**Accessory** | [**SectionComponentResponseAccessory**](SectionComponentResponseAccessory.md) |  | 
**Spacing** | **int32** |  | 
**Divider** | **bool** |  | 
**Content** | **string** |  | 

## Methods

### NewBasicMessageResponseComponentsInner

`func NewBasicMessageResponseComponentsInner(type_ int32, id int32, components []TextDisplayComponentResponse, spoiler bool, file UnfurledMediaResponse, items []MediaGalleryItemResponse, accessory SectionComponentResponseAccessory, spacing int32, divider bool, content string, ) *BasicMessageResponseComponentsInner`

NewBasicMessageResponseComponentsInner instantiates a new BasicMessageResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicMessageResponseComponentsInnerWithDefaults

`func NewBasicMessageResponseComponentsInnerWithDefaults() *BasicMessageResponseComponentsInner`

NewBasicMessageResponseComponentsInnerWithDefaults instantiates a new BasicMessageResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BasicMessageResponseComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicMessageResponseComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicMessageResponseComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *BasicMessageResponseComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicMessageResponseComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicMessageResponseComponentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *BasicMessageResponseComponentsInner) GetComponents() []TextDisplayComponentResponse`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BasicMessageResponseComponentsInner) GetComponentsOk() (*[]TextDisplayComponentResponse, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BasicMessageResponseComponentsInner) SetComponents(v []TextDisplayComponentResponse)`

SetComponents sets Components field to given value.


### GetAccentColor

`func (o *BasicMessageResponseComponentsInner) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *BasicMessageResponseComponentsInner) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *BasicMessageResponseComponentsInner) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *BasicMessageResponseComponentsInner) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### GetSpoiler

`func (o *BasicMessageResponseComponentsInner) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *BasicMessageResponseComponentsInner) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *BasicMessageResponseComponentsInner) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.


### GetFile

`func (o *BasicMessageResponseComponentsInner) GetFile() UnfurledMediaResponse`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *BasicMessageResponseComponentsInner) GetFileOk() (*UnfurledMediaResponse, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *BasicMessageResponseComponentsInner) SetFile(v UnfurledMediaResponse)`

SetFile sets File field to given value.


### GetName

`func (o *BasicMessageResponseComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicMessageResponseComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicMessageResponseComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicMessageResponseComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *BasicMessageResponseComponentsInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BasicMessageResponseComponentsInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BasicMessageResponseComponentsInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BasicMessageResponseComponentsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetItems

`func (o *BasicMessageResponseComponentsInner) GetItems() []MediaGalleryItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BasicMessageResponseComponentsInner) GetItemsOk() (*[]MediaGalleryItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BasicMessageResponseComponentsInner) SetItems(v []MediaGalleryItemResponse)`

SetItems sets Items field to given value.


### GetAccessory

`func (o *BasicMessageResponseComponentsInner) GetAccessory() SectionComponentResponseAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *BasicMessageResponseComponentsInner) GetAccessoryOk() (*SectionComponentResponseAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *BasicMessageResponseComponentsInner) SetAccessory(v SectionComponentResponseAccessory)`

SetAccessory sets Accessory field to given value.


### GetSpacing

`func (o *BasicMessageResponseComponentsInner) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *BasicMessageResponseComponentsInner) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *BasicMessageResponseComponentsInner) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.


### GetDivider

`func (o *BasicMessageResponseComponentsInner) GetDivider() bool`

GetDivider returns the Divider field if non-nil, zero value otherwise.

### GetDividerOk

`func (o *BasicMessageResponseComponentsInner) GetDividerOk() (*bool, bool)`

GetDividerOk returns a tuple with the Divider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivider

`func (o *BasicMessageResponseComponentsInner) SetDivider(v bool)`

SetDivider sets Divider field to given value.


### GetContent

`func (o *BasicMessageResponseComponentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BasicMessageResponseComponentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BasicMessageResponseComponentsInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


