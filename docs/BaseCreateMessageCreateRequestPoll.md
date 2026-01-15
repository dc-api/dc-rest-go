# BaseCreateMessageCreateRequestPoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | [**PollMedia**](PollMedia.md) |  | 
**Answers** | [**[]PollAnswerCreateRequest**](PollAnswerCreateRequest.md) |  | 
**AllowMultiselect** | Pointer to **bool** |  | [optional] 
**LayoutType** | Pointer to [**PollCreateRequestLayoutType**](PollCreateRequestLayoutType.md) |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 

## Methods

### NewBaseCreateMessageCreateRequestPoll

`func NewBaseCreateMessageCreateRequestPoll(question PollMedia, answers []PollAnswerCreateRequest, ) *BaseCreateMessageCreateRequestPoll`

NewBaseCreateMessageCreateRequestPoll instantiates a new BaseCreateMessageCreateRequestPoll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCreateMessageCreateRequestPollWithDefaults

`func NewBaseCreateMessageCreateRequestPollWithDefaults() *BaseCreateMessageCreateRequestPoll`

NewBaseCreateMessageCreateRequestPollWithDefaults instantiates a new BaseCreateMessageCreateRequestPoll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *BaseCreateMessageCreateRequestPoll) GetQuestion() PollMedia`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *BaseCreateMessageCreateRequestPoll) GetQuestionOk() (*PollMedia, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *BaseCreateMessageCreateRequestPoll) SetQuestion(v PollMedia)`

SetQuestion sets Question field to given value.


### GetAnswers

`func (o *BaseCreateMessageCreateRequestPoll) GetAnswers() []PollAnswerCreateRequest`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *BaseCreateMessageCreateRequestPoll) GetAnswersOk() (*[]PollAnswerCreateRequest, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *BaseCreateMessageCreateRequestPoll) SetAnswers(v []PollAnswerCreateRequest)`

SetAnswers sets Answers field to given value.


### GetAllowMultiselect

`func (o *BaseCreateMessageCreateRequestPoll) GetAllowMultiselect() bool`

GetAllowMultiselect returns the AllowMultiselect field if non-nil, zero value otherwise.

### GetAllowMultiselectOk

`func (o *BaseCreateMessageCreateRequestPoll) GetAllowMultiselectOk() (*bool, bool)`

GetAllowMultiselectOk returns a tuple with the AllowMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiselect

`func (o *BaseCreateMessageCreateRequestPoll) SetAllowMultiselect(v bool)`

SetAllowMultiselect sets AllowMultiselect field to given value.

### HasAllowMultiselect

`func (o *BaseCreateMessageCreateRequestPoll) HasAllowMultiselect() bool`

HasAllowMultiselect returns a boolean if a field has been set.

### GetLayoutType

`func (o *BaseCreateMessageCreateRequestPoll) GetLayoutType() PollCreateRequestLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *BaseCreateMessageCreateRequestPoll) GetLayoutTypeOk() (*PollCreateRequestLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *BaseCreateMessageCreateRequestPoll) SetLayoutType(v PollCreateRequestLayoutType)`

SetLayoutType sets LayoutType field to given value.

### HasLayoutType

`func (o *BaseCreateMessageCreateRequestPoll) HasLayoutType() bool`

HasLayoutType returns a boolean if a field has been set.

### GetDuration

`func (o *BaseCreateMessageCreateRequestPoll) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BaseCreateMessageCreateRequestPoll) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BaseCreateMessageCreateRequestPoll) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BaseCreateMessageCreateRequestPoll) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


