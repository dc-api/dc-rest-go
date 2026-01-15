# SlackWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]WebhookSlackEmbed**](WebhookSlackEmbed.md) |  | [optional] 

## Methods

### NewSlackWebhook

`func NewSlackWebhook() *SlackWebhook`

NewSlackWebhook instantiates a new SlackWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackWebhookWithDefaults

`func NewSlackWebhookWithDefaults() *SlackWebhook`

NewSlackWebhookWithDefaults instantiates a new SlackWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SlackWebhook) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SlackWebhook) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SlackWebhook) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SlackWebhook) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUsername

`func (o *SlackWebhook) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SlackWebhook) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SlackWebhook) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SlackWebhook) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIconUrl

`func (o *SlackWebhook) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *SlackWebhook) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *SlackWebhook) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *SlackWebhook) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAttachments

`func (o *SlackWebhook) GetAttachments() []WebhookSlackEmbed`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SlackWebhook) GetAttachmentsOk() (*[]WebhookSlackEmbed, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SlackWebhook) SetAttachments(v []WebhookSlackEmbed)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SlackWebhook) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


