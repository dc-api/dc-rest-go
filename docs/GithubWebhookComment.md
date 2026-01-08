# GithubWebhookComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**User** | [**GithubUser**](GithubUser.md) |  | 
**CommitId** | Pointer to **string** |  | [optional] 
**Body** | **string** |  | 

## Methods

### NewGithubWebhookComment

`func NewGithubWebhookComment(id int32, htmlUrl string, user GithubUser, body string, ) *GithubWebhookComment`

NewGithubWebhookComment instantiates a new GithubWebhookComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookCommentWithDefaults

`func NewGithubWebhookCommentWithDefaults() *GithubWebhookComment`

NewGithubWebhookCommentWithDefaults instantiates a new GithubWebhookComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubWebhookComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubWebhookComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubWebhookComment) SetId(v int32)`

SetId sets Id field to given value.


### GetHtmlUrl

`func (o *GithubWebhookComment) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookComment) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookComment) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *GithubWebhookComment) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubWebhookComment) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubWebhookComment) SetUser(v GithubUser)`

SetUser sets User field to given value.


### GetCommitId

`func (o *GithubWebhookComment) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *GithubWebhookComment) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *GithubWebhookComment) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *GithubWebhookComment) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetBody

`func (o *GithubWebhookComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubWebhookComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubWebhookComment) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


