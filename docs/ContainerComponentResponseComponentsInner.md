# ContainerComponentResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**Components** | [**[]TextDisplayComponentResponse**](TextDisplayComponentResponse.md) |  | 
**File** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Spoiler** | **bool** |  | 
**Items** | [**[]MediaGalleryItemResponse**](MediaGalleryItemResponse.md) |  | 
**Accessory** | [**SectionComponentResponseAccessory**](SectionComponentResponseAccessory.md) |  | 
**Spacing** | **int32** |  | 
**Divider** | **bool** |  | 
**Content** | **string** |  | 

## Methods

### NewContainerComponentResponseComponentsInner

`func NewContainerComponentResponseComponentsInner(type_ int32, id int32, components []TextDisplayComponentResponse, file UnfurledMediaResponse, spoiler bool, items []MediaGalleryItemResponse, accessory SectionComponentResponseAccessory, spacing int32, divider bool, content string, ) *ContainerComponentResponseComponentsInner`

NewContainerComponentResponseComponentsInner instantiates a new ContainerComponentResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerComponentResponseComponentsInnerWithDefaults

`func NewContainerComponentResponseComponentsInnerWithDefaults() *ContainerComponentResponseComponentsInner`

NewContainerComponentResponseComponentsInnerWithDefaults instantiates a new ContainerComponentResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerComponentResponseComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerComponentResponseComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerComponentResponseComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ContainerComponentResponseComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerComponentResponseComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerComponentResponseComponentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *ContainerComponentResponseComponentsInner) GetComponents() []TextDisplayComponentResponse`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ContainerComponentResponseComponentsInner) GetComponentsOk() (*[]TextDisplayComponentResponse, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ContainerComponentResponseComponentsInner) SetComponents(v []TextDisplayComponentResponse)`

SetComponents sets Components field to given value.


### GetFile

`func (o *ContainerComponentResponseComponentsInner) GetFile() UnfurledMediaResponse`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ContainerComponentResponseComponentsInner) GetFileOk() (*UnfurledMediaResponse, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ContainerComponentResponseComponentsInner) SetFile(v UnfurledMediaResponse)`

SetFile sets File field to given value.


### GetName

`func (o *ContainerComponentResponseComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerComponentResponseComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerComponentResponseComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerComponentResponseComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ContainerComponentResponseComponentsInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContainerComponentResponseComponentsInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContainerComponentResponseComponentsInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContainerComponentResponseComponentsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSpoiler

`func (o *ContainerComponentResponseComponentsInner) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ContainerComponentResponseComponentsInner) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ContainerComponentResponseComponentsInner) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.


### GetItems

`func (o *ContainerComponentResponseComponentsInner) GetItems() []MediaGalleryItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContainerComponentResponseComponentsInner) GetItemsOk() (*[]MediaGalleryItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContainerComponentResponseComponentsInner) SetItems(v []MediaGalleryItemResponse)`

SetItems sets Items field to given value.


### GetAccessory

`func (o *ContainerComponentResponseComponentsInner) GetAccessory() SectionComponentResponseAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *ContainerComponentResponseComponentsInner) GetAccessoryOk() (*SectionComponentResponseAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *ContainerComponentResponseComponentsInner) SetAccessory(v SectionComponentResponseAccessory)`

SetAccessory sets Accessory field to given value.


### GetSpacing

`func (o *ContainerComponentResponseComponentsInner) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *ContainerComponentResponseComponentsInner) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *ContainerComponentResponseComponentsInner) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.


### GetDivider

`func (o *ContainerComponentResponseComponentsInner) GetDivider() bool`

GetDivider returns the Divider field if non-nil, zero value otherwise.

### GetDividerOk

`func (o *ContainerComponentResponseComponentsInner) GetDividerOk() (*bool, bool)`

GetDividerOk returns a tuple with the Divider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivider

`func (o *ContainerComponentResponseComponentsInner) SetDivider(v bool)`

SetDivider sets Divider field to given value.


### GetContent

`func (o *ContainerComponentResponseComponentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContainerComponentResponseComponentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContainerComponentResponseComponentsInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


