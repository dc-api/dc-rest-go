# RatelimitedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Discord internal error code. See error code reference | 
**Message** | **string** | Human-readable error message | 
**RetryAfter** | **float32** | The number of seconds to wait before retrying your request | 
**Global** | **bool** | Whether you are being ratelimited by the global ratelimit or a per-endpoint ratelimit | 

## Methods

### NewRatelimitedResponse

`func NewRatelimitedResponse(code int32, message string, retryAfter float32, global bool, ) *RatelimitedResponse`

NewRatelimitedResponse instantiates a new RatelimitedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatelimitedResponseWithDefaults

`func NewRatelimitedResponseWithDefaults() *RatelimitedResponse`

NewRatelimitedResponseWithDefaults instantiates a new RatelimitedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RatelimitedResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RatelimitedResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RatelimitedResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RatelimitedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RatelimitedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RatelimitedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRetryAfter

`func (o *RatelimitedResponse) GetRetryAfter() float32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *RatelimitedResponse) GetRetryAfterOk() (*float32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *RatelimitedResponse) SetRetryAfter(v float32)`

SetRetryAfter sets RetryAfter field to given value.


### GetGlobal

`func (o *RatelimitedResponse) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *RatelimitedResponse) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *RatelimitedResponse) SetGlobal(v bool)`

SetGlobal sets Global field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


