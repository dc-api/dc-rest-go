# WebhookSourceGuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewWebhookSourceGuildResponse

`func NewWebhookSourceGuildResponse(id string, name string, ) *WebhookSourceGuildResponse`

NewWebhookSourceGuildResponse instantiates a new WebhookSourceGuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSourceGuildResponseWithDefaults

`func NewWebhookSourceGuildResponseWithDefaults() *WebhookSourceGuildResponse`

NewWebhookSourceGuildResponseWithDefaults instantiates a new WebhookSourceGuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookSourceGuildResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSourceGuildResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSourceGuildResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *WebhookSourceGuildResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *WebhookSourceGuildResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *WebhookSourceGuildResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *WebhookSourceGuildResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *WebhookSourceGuildResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *WebhookSourceGuildResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetName

`func (o *WebhookSourceGuildResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookSourceGuildResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookSourceGuildResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


