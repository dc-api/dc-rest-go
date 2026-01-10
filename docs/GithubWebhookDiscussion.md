# GithubWebhookDiscussion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Number** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**AnswerHtmlUrl** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**User** | [**GithubUser**](GithubUser.md) |  | 

## Methods

### NewGithubWebhookDiscussion

`func NewGithubWebhookDiscussion(title string, number int32, htmlUrl string, user GithubUser, ) *GithubWebhookDiscussion`

NewGithubWebhookDiscussion instantiates a new GithubWebhookDiscussion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookDiscussionWithDefaults

`func NewGithubWebhookDiscussionWithDefaults() *GithubWebhookDiscussion`

NewGithubWebhookDiscussionWithDefaults instantiates a new GithubWebhookDiscussion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GithubWebhookDiscussion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GithubWebhookDiscussion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GithubWebhookDiscussion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetNumber

`func (o *GithubWebhookDiscussion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GithubWebhookDiscussion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GithubWebhookDiscussion) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetHtmlUrl

`func (o *GithubWebhookDiscussion) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookDiscussion) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookDiscussion) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAnswerHtmlUrl

`func (o *GithubWebhookDiscussion) GetAnswerHtmlUrl() string`

GetAnswerHtmlUrl returns the AnswerHtmlUrl field if non-nil, zero value otherwise.

### GetAnswerHtmlUrlOk

`func (o *GithubWebhookDiscussion) GetAnswerHtmlUrlOk() (*string, bool)`

GetAnswerHtmlUrlOk returns a tuple with the AnswerHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerHtmlUrl

`func (o *GithubWebhookDiscussion) SetAnswerHtmlUrl(v string)`

SetAnswerHtmlUrl sets AnswerHtmlUrl field to given value.

### HasAnswerHtmlUrl

`func (o *GithubWebhookDiscussion) HasAnswerHtmlUrl() bool`

HasAnswerHtmlUrl returns a boolean if a field has been set.

### GetBody

`func (o *GithubWebhookDiscussion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubWebhookDiscussion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubWebhookDiscussion) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubWebhookDiscussion) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetUser

`func (o *GithubWebhookDiscussion) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubWebhookDiscussion) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubWebhookDiscussion) SetUser(v GithubUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


