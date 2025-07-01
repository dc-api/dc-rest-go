# PinnedMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PinnedMessageResponse**](PinnedMessageResponse.md) |  | [optional] 
**HasMore** | **bool** |  | 

## Methods

### NewPinnedMessagesResponse

`func NewPinnedMessagesResponse(hasMore bool, ) *PinnedMessagesResponse`

NewPinnedMessagesResponse instantiates a new PinnedMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedMessagesResponseWithDefaults

`func NewPinnedMessagesResponseWithDefaults() *PinnedMessagesResponse`

NewPinnedMessagesResponseWithDefaults instantiates a new PinnedMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PinnedMessagesResponse) GetItems() []PinnedMessageResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PinnedMessagesResponse) GetItemsOk() (*[]PinnedMessageResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PinnedMessagesResponse) SetItems(v []PinnedMessageResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *PinnedMessagesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *PinnedMessagesResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *PinnedMessagesResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetHasMore

`func (o *PinnedMessagesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PinnedMessagesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PinnedMessagesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


