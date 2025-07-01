# ThreadSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threads** | [**[]ThreadResponse**](ThreadResponse.md) |  | 
**Members** | [**[]ThreadMemberResponse**](ThreadMemberResponse.md) |  | 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**FirstMessages** | Pointer to [**[]MessageResponse**](MessageResponse.md) |  | [optional] 
**TotalResults** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewThreadSearchResponse

`func NewThreadSearchResponse(threads []ThreadResponse, members []ThreadMemberResponse, ) *ThreadSearchResponse`

NewThreadSearchResponse instantiates a new ThreadSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadSearchResponseWithDefaults

`func NewThreadSearchResponseWithDefaults() *ThreadSearchResponse`

NewThreadSearchResponseWithDefaults instantiates a new ThreadSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreads

`func (o *ThreadSearchResponse) GetThreads() []ThreadResponse`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ThreadSearchResponse) GetThreadsOk() (*[]ThreadResponse, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ThreadSearchResponse) SetThreads(v []ThreadResponse)`

SetThreads sets Threads field to given value.


### GetMembers

`func (o *ThreadSearchResponse) GetMembers() []ThreadMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ThreadSearchResponse) GetMembersOk() (*[]ThreadMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ThreadSearchResponse) SetMembers(v []ThreadMemberResponse)`

SetMembers sets Members field to given value.


### GetHasMore

`func (o *ThreadSearchResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ThreadSearchResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ThreadSearchResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ThreadSearchResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *ThreadSearchResponse) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *ThreadSearchResponse) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetFirstMessages

`func (o *ThreadSearchResponse) GetFirstMessages() []MessageResponse`

GetFirstMessages returns the FirstMessages field if non-nil, zero value otherwise.

### GetFirstMessagesOk

`func (o *ThreadSearchResponse) GetFirstMessagesOk() (*[]MessageResponse, bool)`

GetFirstMessagesOk returns a tuple with the FirstMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessages

`func (o *ThreadSearchResponse) SetFirstMessages(v []MessageResponse)`

SetFirstMessages sets FirstMessages field to given value.

### HasFirstMessages

`func (o *ThreadSearchResponse) HasFirstMessages() bool`

HasFirstMessages returns a boolean if a field has been set.

### SetFirstMessagesNil

`func (o *ThreadSearchResponse) SetFirstMessagesNil(b bool)`

 SetFirstMessagesNil sets the value for FirstMessages to be an explicit nil

### UnsetFirstMessages
`func (o *ThreadSearchResponse) UnsetFirstMessages()`

UnsetFirstMessages ensures that no value is present for FirstMessages, not even an explicit nil
### GetTotalResults

`func (o *ThreadSearchResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ThreadSearchResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ThreadSearchResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ThreadSearchResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### SetTotalResultsNil

`func (o *ThreadSearchResponse) SetTotalResultsNil(b bool)`

 SetTotalResultsNil sets the value for TotalResults to be an explicit nil

### UnsetTotalResults
`func (o *ThreadSearchResponse) UnsetTotalResults()`

UnsetTotalResults ensures that no value is present for TotalResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


