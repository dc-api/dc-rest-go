# PollAnswerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerId** | **int32** |  | 
**PollMedia** | [**PollMediaResponse**](PollMediaResponse.md) |  | 

## Methods

### NewPollAnswerResponse

`func NewPollAnswerResponse(answerId int32, pollMedia PollMediaResponse, ) *PollAnswerResponse`

NewPollAnswerResponse instantiates a new PollAnswerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollAnswerResponseWithDefaults

`func NewPollAnswerResponseWithDefaults() *PollAnswerResponse`

NewPollAnswerResponseWithDefaults instantiates a new PollAnswerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerId

`func (o *PollAnswerResponse) GetAnswerId() int32`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *PollAnswerResponse) GetAnswerIdOk() (*int32, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *PollAnswerResponse) SetAnswerId(v int32)`

SetAnswerId sets AnswerId field to given value.


### GetPollMedia

`func (o *PollAnswerResponse) GetPollMedia() PollMediaResponse`

GetPollMedia returns the PollMedia field if non-nil, zero value otherwise.

### GetPollMediaOk

`func (o *PollAnswerResponse) GetPollMediaOk() (*PollMediaResponse, bool)`

GetPollMediaOk returns a tuple with the PollMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollMedia

`func (o *PollAnswerResponse) SetPollMedia(v PollMediaResponse)`

SetPollMedia sets PollMedia field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


