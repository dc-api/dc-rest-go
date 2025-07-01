# ContainerComponentForMessageRequestComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Components** | [**[]TextDisplayComponentForMessageRequest**](TextDisplayComponentForMessageRequest.md) |  | 
**Spoiler** | Pointer to **bool** |  | [optional] 
**File** | [**UnfurledMediaRequestWithAttachmentReferenceRequired**](UnfurledMediaRequestWithAttachmentReferenceRequired.md) |  | 
**Items** | [**[]MediaGalleryItemRequest**](MediaGalleryItemRequest.md) |  | 
**Accessory** | [**SectionComponentForMessageRequestAccessory**](SectionComponentForMessageRequestAccessory.md) |  | 
**Spacing** | Pointer to **int32** |  | [optional] 
**Divider** | Pointer to **bool** |  | [optional] 
**Content** | **string** |  | 

## Methods

### NewContainerComponentForMessageRequestComponentsInner

`func NewContainerComponentForMessageRequestComponentsInner(type_ int32, components []TextDisplayComponentForMessageRequest, file UnfurledMediaRequestWithAttachmentReferenceRequired, items []MediaGalleryItemRequest, accessory SectionComponentForMessageRequestAccessory, content string, ) *ContainerComponentForMessageRequestComponentsInner`

NewContainerComponentForMessageRequestComponentsInner instantiates a new ContainerComponentForMessageRequestComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerComponentForMessageRequestComponentsInnerWithDefaults

`func NewContainerComponentForMessageRequestComponentsInnerWithDefaults() *ContainerComponentForMessageRequestComponentsInner`

NewContainerComponentForMessageRequestComponentsInnerWithDefaults instantiates a new ContainerComponentForMessageRequestComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerComponentForMessageRequestComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerComponentForMessageRequestComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetComponents

`func (o *ContainerComponentForMessageRequestComponentsInner) GetComponents() []TextDisplayComponentForMessageRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetComponentsOk() (*[]TextDisplayComponentForMessageRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ContainerComponentForMessageRequestComponentsInner) SetComponents(v []TextDisplayComponentForMessageRequest)`

SetComponents sets Components field to given value.


### GetSpoiler

`func (o *ContainerComponentForMessageRequestComponentsInner) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *ContainerComponentForMessageRequestComponentsInner) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *ContainerComponentForMessageRequestComponentsInner) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetFile

`func (o *ContainerComponentForMessageRequestComponentsInner) GetFile() UnfurledMediaRequestWithAttachmentReferenceRequired`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetFileOk() (*UnfurledMediaRequestWithAttachmentReferenceRequired, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ContainerComponentForMessageRequestComponentsInner) SetFile(v UnfurledMediaRequestWithAttachmentReferenceRequired)`

SetFile sets File field to given value.


### GetItems

`func (o *ContainerComponentForMessageRequestComponentsInner) GetItems() []MediaGalleryItemRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetItemsOk() (*[]MediaGalleryItemRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContainerComponentForMessageRequestComponentsInner) SetItems(v []MediaGalleryItemRequest)`

SetItems sets Items field to given value.


### GetAccessory

`func (o *ContainerComponentForMessageRequestComponentsInner) GetAccessory() SectionComponentForMessageRequestAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetAccessoryOk() (*SectionComponentForMessageRequestAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *ContainerComponentForMessageRequestComponentsInner) SetAccessory(v SectionComponentForMessageRequestAccessory)`

SetAccessory sets Accessory field to given value.


### GetSpacing

`func (o *ContainerComponentForMessageRequestComponentsInner) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *ContainerComponentForMessageRequestComponentsInner) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.

### HasSpacing

`func (o *ContainerComponentForMessageRequestComponentsInner) HasSpacing() bool`

HasSpacing returns a boolean if a field has been set.

### GetDivider

`func (o *ContainerComponentForMessageRequestComponentsInner) GetDivider() bool`

GetDivider returns the Divider field if non-nil, zero value otherwise.

### GetDividerOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetDividerOk() (*bool, bool)`

GetDividerOk returns a tuple with the Divider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivider

`func (o *ContainerComponentForMessageRequestComponentsInner) SetDivider(v bool)`

SetDivider sets Divider field to given value.

### HasDivider

`func (o *ContainerComponentForMessageRequestComponentsInner) HasDivider() bool`

HasDivider returns a boolean if a field has been set.

### GetContent

`func (o *ContainerComponentForMessageRequestComponentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContainerComponentForMessageRequestComponentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContainerComponentForMessageRequestComponentsInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


