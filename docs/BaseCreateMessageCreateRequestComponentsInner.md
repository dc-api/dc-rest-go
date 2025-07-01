# BaseCreateMessageCreateRequestComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Components** | [**[]TextDisplayComponentForMessageRequest**](TextDisplayComponentForMessageRequest.md) |  | 
**AccentColor** | Pointer to **int32** |  | [optional] 
**Spoiler** | Pointer to **bool** |  | [optional] 
**File** | [**UnfurledMediaRequestWithAttachmentReferenceRequired**](UnfurledMediaRequestWithAttachmentReferenceRequired.md) |  | 
**Items** | [**[]MediaGalleryItemRequest**](MediaGalleryItemRequest.md) |  | 
**Accessory** | [**SectionComponentForMessageRequestAccessory**](SectionComponentForMessageRequestAccessory.md) |  | 
**Spacing** | Pointer to **int32** |  | [optional] 
**Divider** | Pointer to **bool** |  | [optional] 
**Content** | **string** |  | 

## Methods

### NewBaseCreateMessageCreateRequestComponentsInner

`func NewBaseCreateMessageCreateRequestComponentsInner(type_ int32, components []TextDisplayComponentForMessageRequest, file UnfurledMediaRequestWithAttachmentReferenceRequired, items []MediaGalleryItemRequest, accessory SectionComponentForMessageRequestAccessory, content string, ) *BaseCreateMessageCreateRequestComponentsInner`

NewBaseCreateMessageCreateRequestComponentsInner instantiates a new BaseCreateMessageCreateRequestComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCreateMessageCreateRequestComponentsInnerWithDefaults

`func NewBaseCreateMessageCreateRequestComponentsInnerWithDefaults() *BaseCreateMessageCreateRequestComponentsInner`

NewBaseCreateMessageCreateRequestComponentsInnerWithDefaults instantiates a new BaseCreateMessageCreateRequestComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetComponents

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetComponents() []TextDisplayComponentForMessageRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetComponentsOk() (*[]TextDisplayComponentForMessageRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetComponents(v []TextDisplayComponentForMessageRequest)`

SetComponents sets Components field to given value.


### GetAccentColor

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetAccentColor() int32`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetAccentColorOk() (*int32, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetAccentColor(v int32)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *BaseCreateMessageCreateRequestComponentsInner) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### GetSpoiler

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *BaseCreateMessageCreateRequestComponentsInner) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetFile

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetFile() UnfurledMediaRequestWithAttachmentReferenceRequired`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetFileOk() (*UnfurledMediaRequestWithAttachmentReferenceRequired, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetFile(v UnfurledMediaRequestWithAttachmentReferenceRequired)`

SetFile sets File field to given value.


### GetItems

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetItems() []MediaGalleryItemRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetItemsOk() (*[]MediaGalleryItemRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetItems(v []MediaGalleryItemRequest)`

SetItems sets Items field to given value.


### GetAccessory

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetAccessory() SectionComponentForMessageRequestAccessory`

GetAccessory returns the Accessory field if non-nil, zero value otherwise.

### GetAccessoryOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetAccessoryOk() (*SectionComponentForMessageRequestAccessory, bool)`

GetAccessoryOk returns a tuple with the Accessory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessory

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetAccessory(v SectionComponentForMessageRequestAccessory)`

SetAccessory sets Accessory field to given value.


### GetSpacing

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.

### HasSpacing

`func (o *BaseCreateMessageCreateRequestComponentsInner) HasSpacing() bool`

HasSpacing returns a boolean if a field has been set.

### GetDivider

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetDivider() bool`

GetDivider returns the Divider field if non-nil, zero value otherwise.

### GetDividerOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetDividerOk() (*bool, bool)`

GetDividerOk returns a tuple with the Divider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivider

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetDivider(v bool)`

SetDivider sets Divider field to given value.

### HasDivider

`func (o *BaseCreateMessageCreateRequestComponentsInner) HasDivider() bool`

HasDivider returns a boolean if a field has been set.

### GetContent

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BaseCreateMessageCreateRequestComponentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BaseCreateMessageCreateRequestComponentsInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


