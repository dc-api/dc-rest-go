# GithubDiscussion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Number** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**AnswerHtmlUrl** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**User** | [**GithubUser**](GithubUser.md) |  | 

## Methods

### NewGithubDiscussion

`func NewGithubDiscussion(title string, number int32, htmlUrl string, user GithubUser, ) *GithubDiscussion`

NewGithubDiscussion instantiates a new GithubDiscussion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubDiscussionWithDefaults

`func NewGithubDiscussionWithDefaults() *GithubDiscussion`

NewGithubDiscussionWithDefaults instantiates a new GithubDiscussion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GithubDiscussion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GithubDiscussion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GithubDiscussion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetNumber

`func (o *GithubDiscussion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GithubDiscussion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GithubDiscussion) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetHtmlUrl

`func (o *GithubDiscussion) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubDiscussion) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubDiscussion) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAnswerHtmlUrl

`func (o *GithubDiscussion) GetAnswerHtmlUrl() string`

GetAnswerHtmlUrl returns the AnswerHtmlUrl field if non-nil, zero value otherwise.

### GetAnswerHtmlUrlOk

`func (o *GithubDiscussion) GetAnswerHtmlUrlOk() (*string, bool)`

GetAnswerHtmlUrlOk returns a tuple with the AnswerHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerHtmlUrl

`func (o *GithubDiscussion) SetAnswerHtmlUrl(v string)`

SetAnswerHtmlUrl sets AnswerHtmlUrl field to given value.

### HasAnswerHtmlUrl

`func (o *GithubDiscussion) HasAnswerHtmlUrl() bool`

HasAnswerHtmlUrl returns a boolean if a field has been set.

### SetAnswerHtmlUrlNil

`func (o *GithubDiscussion) SetAnswerHtmlUrlNil(b bool)`

 SetAnswerHtmlUrlNil sets the value for AnswerHtmlUrl to be an explicit nil

### UnsetAnswerHtmlUrl
`func (o *GithubDiscussion) UnsetAnswerHtmlUrl()`

UnsetAnswerHtmlUrl ensures that no value is present for AnswerHtmlUrl, not even an explicit nil
### GetBody

`func (o *GithubDiscussion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubDiscussion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubDiscussion) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubDiscussion) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GithubDiscussion) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GithubDiscussion) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetUser

`func (o *GithubDiscussion) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubDiscussion) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubDiscussion) SetUser(v GithubUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


