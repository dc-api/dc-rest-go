# ForumTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Moderated** | **bool** |  | 
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewForumTagResponse

`func NewForumTagResponse(id string, name string, moderated bool, ) *ForumTagResponse`

NewForumTagResponse instantiates a new ForumTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTagResponseWithDefaults

`func NewForumTagResponseWithDefaults() *ForumTagResponse`

NewForumTagResponseWithDefaults instantiates a new ForumTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ForumTagResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForumTagResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForumTagResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ForumTagResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumTagResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumTagResponse) SetName(v string)`

SetName sets Name field to given value.


### GetModerated

`func (o *ForumTagResponse) GetModerated() bool`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *ForumTagResponse) GetModeratedOk() (*bool, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *ForumTagResponse) SetModerated(v bool)`

SetModerated sets Moderated field to given value.


### GetEmojiId

`func (o *ForumTagResponse) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *ForumTagResponse) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *ForumTagResponse) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *ForumTagResponse) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *ForumTagResponse) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *ForumTagResponse) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *ForumTagResponse) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *ForumTagResponse) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *ForumTagResponse) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *ForumTagResponse) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


