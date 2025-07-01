# PollCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | [**PollMedia**](PollMedia.md) |  | 
**Answers** | [**[]PollAnswerCreateRequest**](PollAnswerCreateRequest.md) |  | 
**AllowMultiselect** | Pointer to **NullableBool** |  | [optional] 
**LayoutType** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPollCreateRequest

`func NewPollCreateRequest(question PollMedia, answers []PollAnswerCreateRequest, ) *PollCreateRequest`

NewPollCreateRequest instantiates a new PollCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollCreateRequestWithDefaults

`func NewPollCreateRequestWithDefaults() *PollCreateRequest`

NewPollCreateRequestWithDefaults instantiates a new PollCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *PollCreateRequest) GetQuestion() PollMedia`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PollCreateRequest) GetQuestionOk() (*PollMedia, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PollCreateRequest) SetQuestion(v PollMedia)`

SetQuestion sets Question field to given value.


### GetAnswers

`func (o *PollCreateRequest) GetAnswers() []PollAnswerCreateRequest`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *PollCreateRequest) GetAnswersOk() (*[]PollAnswerCreateRequest, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *PollCreateRequest) SetAnswers(v []PollAnswerCreateRequest)`

SetAnswers sets Answers field to given value.


### GetAllowMultiselect

`func (o *PollCreateRequest) GetAllowMultiselect() bool`

GetAllowMultiselect returns the AllowMultiselect field if non-nil, zero value otherwise.

### GetAllowMultiselectOk

`func (o *PollCreateRequest) GetAllowMultiselectOk() (*bool, bool)`

GetAllowMultiselectOk returns a tuple with the AllowMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiselect

`func (o *PollCreateRequest) SetAllowMultiselect(v bool)`

SetAllowMultiselect sets AllowMultiselect field to given value.

### HasAllowMultiselect

`func (o *PollCreateRequest) HasAllowMultiselect() bool`

HasAllowMultiselect returns a boolean if a field has been set.

### SetAllowMultiselectNil

`func (o *PollCreateRequest) SetAllowMultiselectNil(b bool)`

 SetAllowMultiselectNil sets the value for AllowMultiselect to be an explicit nil

### UnsetAllowMultiselect
`func (o *PollCreateRequest) UnsetAllowMultiselect()`

UnsetAllowMultiselect ensures that no value is present for AllowMultiselect, not even an explicit nil
### GetLayoutType

`func (o *PollCreateRequest) GetLayoutType() int32`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *PollCreateRequest) GetLayoutTypeOk() (*int32, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *PollCreateRequest) SetLayoutType(v int32)`

SetLayoutType sets LayoutType field to given value.

### HasLayoutType

`func (o *PollCreateRequest) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### GetDuration

`func (o *PollCreateRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PollCreateRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PollCreateRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PollCreateRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *PollCreateRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *PollCreateRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


