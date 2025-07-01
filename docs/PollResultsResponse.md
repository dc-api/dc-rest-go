# PollResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerCounts** | Pointer to [**[]PollResultsEntryResponse**](PollResultsEntryResponse.md) |  | [optional] 
**IsFinalized** | **bool** |  | 

## Methods

### NewPollResultsResponse

`func NewPollResultsResponse(isFinalized bool, ) *PollResultsResponse`

NewPollResultsResponse instantiates a new PollResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollResultsResponseWithDefaults

`func NewPollResultsResponseWithDefaults() *PollResultsResponse`

NewPollResultsResponseWithDefaults instantiates a new PollResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerCounts

`func (o *PollResultsResponse) GetAnswerCounts() []PollResultsEntryResponse`

GetAnswerCounts returns the AnswerCounts field if non-nil, zero value otherwise.

### GetAnswerCountsOk

`func (o *PollResultsResponse) GetAnswerCountsOk() (*[]PollResultsEntryResponse, bool)`

GetAnswerCountsOk returns a tuple with the AnswerCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerCounts

`func (o *PollResultsResponse) SetAnswerCounts(v []PollResultsEntryResponse)`

SetAnswerCounts sets AnswerCounts field to given value.

### HasAnswerCounts

`func (o *PollResultsResponse) HasAnswerCounts() bool`

HasAnswerCounts returns a boolean if a field has been set.

### SetAnswerCountsNil

`func (o *PollResultsResponse) SetAnswerCountsNil(b bool)`

 SetAnswerCountsNil sets the value for AnswerCounts to be an explicit nil

### UnsetAnswerCounts
`func (o *PollResultsResponse) UnsetAnswerCounts()`

UnsetAnswerCounts ensures that no value is present for AnswerCounts, not even an explicit nil
### GetIsFinalized

`func (o *PollResultsResponse) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *PollResultsResponse) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *PollResultsResponse) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


