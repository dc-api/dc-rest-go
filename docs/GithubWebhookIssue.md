# GithubWebhookIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Number** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**User** | [**GithubUser**](GithubUser.md) |  | 
**Title** | **string** |  | 
**Body** | Pointer to **string** |  | [optional] 
**PullRequest** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGithubWebhookIssue

`func NewGithubWebhookIssue(id int32, number int32, htmlUrl string, user GithubUser, title string, ) *GithubWebhookIssue`

NewGithubWebhookIssue instantiates a new GithubWebhookIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookIssueWithDefaults

`func NewGithubWebhookIssueWithDefaults() *GithubWebhookIssue`

NewGithubWebhookIssueWithDefaults instantiates a new GithubWebhookIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubWebhookIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubWebhookIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubWebhookIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *GithubWebhookIssue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GithubWebhookIssue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GithubWebhookIssue) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetHtmlUrl

`func (o *GithubWebhookIssue) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookIssue) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookIssue) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *GithubWebhookIssue) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubWebhookIssue) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubWebhookIssue) SetUser(v GithubUser)`

SetUser sets User field to given value.


### GetTitle

`func (o *GithubWebhookIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GithubWebhookIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GithubWebhookIssue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *GithubWebhookIssue) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubWebhookIssue) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubWebhookIssue) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubWebhookIssue) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPullRequest

`func (o *GithubWebhookIssue) GetPullRequest() interface{}`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *GithubWebhookIssue) GetPullRequestOk() (*interface{}, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *GithubWebhookIssue) SetPullRequest(v interface{})`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *GithubWebhookIssue) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### SetPullRequestNil

`func (o *GithubWebhookIssue) SetPullRequestNil(b bool)`

 SetPullRequestNil sets the value for PullRequest to be an explicit nil

### UnsetPullRequest
`func (o *GithubWebhookIssue) UnsetPullRequest()`

UnsetPullRequest ensures that no value is present for PullRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


