# PollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | [**PollMediaResponse**](PollMediaResponse.md) |  | 
**Answers** | [**[]PollAnswerResponse**](PollAnswerResponse.md) |  | 
**Expiry** | **time.Time** |  | 
**AllowMultiselect** | **bool** |  | 
**LayoutType** | **int32** |  | 
**Results** | [**PollResultsResponse**](PollResultsResponse.md) |  | 

## Methods

### NewPollResponse

`func NewPollResponse(question PollMediaResponse, answers []PollAnswerResponse, expiry time.Time, allowMultiselect bool, layoutType int32, results PollResultsResponse, ) *PollResponse`

NewPollResponse instantiates a new PollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollResponseWithDefaults

`func NewPollResponseWithDefaults() *PollResponse`

NewPollResponseWithDefaults instantiates a new PollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *PollResponse) GetQuestion() PollMediaResponse`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PollResponse) GetQuestionOk() (*PollMediaResponse, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PollResponse) SetQuestion(v PollMediaResponse)`

SetQuestion sets Question field to given value.


### GetAnswers

`func (o *PollResponse) GetAnswers() []PollAnswerResponse`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *PollResponse) GetAnswersOk() (*[]PollAnswerResponse, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *PollResponse) SetAnswers(v []PollAnswerResponse)`

SetAnswers sets Answers field to given value.


### GetExpiry

`func (o *PollResponse) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PollResponse) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PollResponse) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetAllowMultiselect

`func (o *PollResponse) GetAllowMultiselect() bool`

GetAllowMultiselect returns the AllowMultiselect field if non-nil, zero value otherwise.

### GetAllowMultiselectOk

`func (o *PollResponse) GetAllowMultiselectOk() (*bool, bool)`

GetAllowMultiselectOk returns a tuple with the AllowMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiselect

`func (o *PollResponse) SetAllowMultiselect(v bool)`

SetAllowMultiselect sets AllowMultiselect field to given value.


### GetLayoutType

`func (o *PollResponse) GetLayoutType() int32`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *PollResponse) GetLayoutTypeOk() (*int32, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *PollResponse) SetLayoutType(v int32)`

SetLayoutType sets LayoutType field to given value.


### GetResults

`func (o *PollResponse) GetResults() PollResultsResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PollResponse) GetResultsOk() (*PollResultsResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PollResponse) SetResults(v PollResultsResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


