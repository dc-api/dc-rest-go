# GithubWebhookCheckSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **string** |  | [optional] 
**HeadBranch** | Pointer to **string** |  | [optional] 
**HeadSha** | **string** |  | 
**PullRequests** | Pointer to [**[]GithubCheckPullRequest**](GithubCheckPullRequest.md) |  | [optional] 
**App** | [**GithubCheckApp**](GithubCheckApp.md) |  | 

## Methods

### NewGithubWebhookCheckSuite

`func NewGithubWebhookCheckSuite(headSha string, app GithubCheckApp, ) *GithubWebhookCheckSuite`

NewGithubWebhookCheckSuite instantiates a new GithubWebhookCheckSuite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookCheckSuiteWithDefaults

`func NewGithubWebhookCheckSuiteWithDefaults() *GithubWebhookCheckSuite`

NewGithubWebhookCheckSuiteWithDefaults instantiates a new GithubWebhookCheckSuite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *GithubWebhookCheckSuite) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *GithubWebhookCheckSuite) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *GithubWebhookCheckSuite) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *GithubWebhookCheckSuite) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetHeadBranch

`func (o *GithubWebhookCheckSuite) GetHeadBranch() string`

GetHeadBranch returns the HeadBranch field if non-nil, zero value otherwise.

### GetHeadBranchOk

`func (o *GithubWebhookCheckSuite) GetHeadBranchOk() (*string, bool)`

GetHeadBranchOk returns a tuple with the HeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadBranch

`func (o *GithubWebhookCheckSuite) SetHeadBranch(v string)`

SetHeadBranch sets HeadBranch field to given value.

### HasHeadBranch

`func (o *GithubWebhookCheckSuite) HasHeadBranch() bool`

HasHeadBranch returns a boolean if a field has been set.

### GetHeadSha

`func (o *GithubWebhookCheckSuite) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *GithubWebhookCheckSuite) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *GithubWebhookCheckSuite) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetPullRequests

`func (o *GithubWebhookCheckSuite) GetPullRequests() []GithubCheckPullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *GithubWebhookCheckSuite) GetPullRequestsOk() (*[]GithubCheckPullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *GithubWebhookCheckSuite) SetPullRequests(v []GithubCheckPullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *GithubWebhookCheckSuite) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetApp

`func (o *GithubWebhookCheckSuite) GetApp() GithubCheckApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GithubWebhookCheckSuite) GetAppOk() (*GithubCheckApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GithubWebhookCheckSuite) SetApp(v GithubCheckApp)`

SetApp sets App field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


