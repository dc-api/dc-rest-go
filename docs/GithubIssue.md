# GithubIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Number** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**User** | [**GithubUser**](GithubUser.md) |  | 
**Title** | **string** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**PullRequest** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGithubIssue

`func NewGithubIssue(id int32, number int32, htmlUrl string, user GithubUser, title string, ) *GithubIssue`

NewGithubIssue instantiates a new GithubIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubIssueWithDefaults

`func NewGithubIssueWithDefaults() *GithubIssue`

NewGithubIssueWithDefaults instantiates a new GithubIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *GithubIssue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GithubIssue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GithubIssue) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetHtmlUrl

`func (o *GithubIssue) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubIssue) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubIssue) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *GithubIssue) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubIssue) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubIssue) SetUser(v GithubUser)`

SetUser sets User field to given value.


### GetTitle

`func (o *GithubIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GithubIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GithubIssue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *GithubIssue) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubIssue) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubIssue) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubIssue) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GithubIssue) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GithubIssue) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetPullRequest

`func (o *GithubIssue) GetPullRequest() interface{}`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *GithubIssue) GetPullRequestOk() (*interface{}, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *GithubIssue) SetPullRequest(v interface{})`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *GithubIssue) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### SetPullRequestNil

`func (o *GithubIssue) SetPullRequestNil(b bool)`

 SetPullRequestNil sets the value for PullRequest to be an explicit nil

### UnsetPullRequest
`func (o *GithubIssue) UnsetPullRequest()`

UnsetPullRequest ensures that no value is present for PullRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


