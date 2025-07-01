# ThreadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threads** | [**[]ThreadResponse**](ThreadResponse.md) |  | 
**Members** | [**[]ThreadMemberResponse**](ThreadMemberResponse.md) |  | 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**FirstMessages** | Pointer to [**[]MessageResponse**](MessageResponse.md) |  | [optional] 

## Methods

### NewThreadsResponse

`func NewThreadsResponse(threads []ThreadResponse, members []ThreadMemberResponse, ) *ThreadsResponse`

NewThreadsResponse instantiates a new ThreadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadsResponseWithDefaults

`func NewThreadsResponseWithDefaults() *ThreadsResponse`

NewThreadsResponseWithDefaults instantiates a new ThreadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreads

`func (o *ThreadsResponse) GetThreads() []ThreadResponse`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ThreadsResponse) GetThreadsOk() (*[]ThreadResponse, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ThreadsResponse) SetThreads(v []ThreadResponse)`

SetThreads sets Threads field to given value.


### GetMembers

`func (o *ThreadsResponse) GetMembers() []ThreadMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ThreadsResponse) GetMembersOk() (*[]ThreadMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ThreadsResponse) SetMembers(v []ThreadMemberResponse)`

SetMembers sets Members field to given value.


### GetHasMore

`func (o *ThreadsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ThreadsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ThreadsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ThreadsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *ThreadsResponse) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *ThreadsResponse) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetFirstMessages

`func (o *ThreadsResponse) GetFirstMessages() []MessageResponse`

GetFirstMessages returns the FirstMessages field if non-nil, zero value otherwise.

### GetFirstMessagesOk

`func (o *ThreadsResponse) GetFirstMessagesOk() (*[]MessageResponse, bool)`

GetFirstMessagesOk returns a tuple with the FirstMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessages

`func (o *ThreadsResponse) SetFirstMessages(v []MessageResponse)`

SetFirstMessages sets FirstMessages field to given value.

### HasFirstMessages

`func (o *ThreadsResponse) HasFirstMessages() bool`

HasFirstMessages returns a boolean if a field has been set.

### SetFirstMessagesNil

`func (o *ThreadsResponse) SetFirstMessagesNil(b bool)`

 SetFirstMessagesNil sets the value for FirstMessages to be an explicit nil

### UnsetFirstMessages
`func (o *ThreadsResponse) UnsetFirstMessages()`

UnsetFirstMessages ensures that no value is present for FirstMessages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


