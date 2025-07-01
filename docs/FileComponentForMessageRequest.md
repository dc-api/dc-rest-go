# FileComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Spoiler** | Pointer to **NullableBool** |  | [optional] 
**File** | [**UnfurledMediaRequestWithAttachmentReferenceRequired**](UnfurledMediaRequestWithAttachmentReferenceRequired.md) |  | 

## Methods

### NewFileComponentForMessageRequest

`func NewFileComponentForMessageRequest(type_ int32, file UnfurledMediaRequestWithAttachmentReferenceRequired, ) *FileComponentForMessageRequest`

NewFileComponentForMessageRequest instantiates a new FileComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileComponentForMessageRequestWithDefaults

`func NewFileComponentForMessageRequestWithDefaults() *FileComponentForMessageRequest`

NewFileComponentForMessageRequestWithDefaults instantiates a new FileComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FileComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetSpoiler

`func (o *FileComponentForMessageRequest) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *FileComponentForMessageRequest) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *FileComponentForMessageRequest) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *FileComponentForMessageRequest) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### SetSpoilerNil

`func (o *FileComponentForMessageRequest) SetSpoilerNil(b bool)`

 SetSpoilerNil sets the value for Spoiler to be an explicit nil

### UnsetSpoiler
`func (o *FileComponentForMessageRequest) UnsetSpoiler()`

UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil
### GetFile

`func (o *FileComponentForMessageRequest) GetFile() UnfurledMediaRequestWithAttachmentReferenceRequired`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileComponentForMessageRequest) GetFileOk() (*UnfurledMediaRequestWithAttachmentReferenceRequired, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileComponentForMessageRequest) SetFile(v UnfurledMediaRequestWithAttachmentReferenceRequired)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


