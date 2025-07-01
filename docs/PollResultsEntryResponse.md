# PollResultsEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Count** | **int32** |  | 
**MeVoted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPollResultsEntryResponse

`func NewPollResultsEntryResponse(id int32, count int32, ) *PollResultsEntryResponse`

NewPollResultsEntryResponse instantiates a new PollResultsEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollResultsEntryResponseWithDefaults

`func NewPollResultsEntryResponseWithDefaults() *PollResultsEntryResponse`

NewPollResultsEntryResponseWithDefaults instantiates a new PollResultsEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PollResultsEntryResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PollResultsEntryResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PollResultsEntryResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCount

`func (o *PollResultsEntryResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PollResultsEntryResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PollResultsEntryResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetMeVoted

`func (o *PollResultsEntryResponse) GetMeVoted() bool`

GetMeVoted returns the MeVoted field if non-nil, zero value otherwise.

### GetMeVotedOk

`func (o *PollResultsEntryResponse) GetMeVotedOk() (*bool, bool)`

GetMeVotedOk returns a tuple with the MeVoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeVoted

`func (o *PollResultsEntryResponse) SetMeVoted(v bool)`

SetMeVoted sets MeVoted field to given value.

### HasMeVoted

`func (o *PollResultsEntryResponse) HasMeVoted() bool`

HasMeVoted returns a boolean if a field has been set.

### SetMeVotedNil

`func (o *PollResultsEntryResponse) SetMeVotedNil(b bool)`

 SetMeVotedNil sets the value for MeVoted to be an explicit nil

### UnsetMeVoted
`func (o *PollResultsEntryResponse) UnsetMeVoted()`

UnsetMeVoted ensures that no value is present for MeVoted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


