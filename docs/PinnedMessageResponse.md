# PinnedMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PinnedAt** | **time.Time** |  | 
**Message** | [**MessageResponse**](MessageResponse.md) |  | 

## Methods

### NewPinnedMessageResponse

`func NewPinnedMessageResponse(pinnedAt time.Time, message MessageResponse, ) *PinnedMessageResponse`

NewPinnedMessageResponse instantiates a new PinnedMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedMessageResponseWithDefaults

`func NewPinnedMessageResponseWithDefaults() *PinnedMessageResponse`

NewPinnedMessageResponseWithDefaults instantiates a new PinnedMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPinnedAt

`func (o *PinnedMessageResponse) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *PinnedMessageResponse) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *PinnedMessageResponse) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.


### GetMessage

`func (o *PinnedMessageResponse) GetMessage() MessageResponse`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PinnedMessageResponse) GetMessageOk() (*MessageResponse, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PinnedMessageResponse) SetMessage(v MessageResponse)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


