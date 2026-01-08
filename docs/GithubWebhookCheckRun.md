# GithubWebhookCheckRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CheckSuite** | [**GithubCheckSuite**](GithubCheckSuite.md) |  | 
**DetailsUrl** | Pointer to **string** |  | [optional] 
**Output** | Pointer to [**GithubCheckRunOutput**](GithubCheckRunOutput.md) |  | [optional] 
**PullRequests** | Pointer to [**[]GithubCheckPullRequest**](GithubCheckPullRequest.md) |  | [optional] 

## Methods

### NewGithubWebhookCheckRun

`func NewGithubWebhookCheckRun(name string, htmlUrl string, checkSuite GithubCheckSuite, ) *GithubWebhookCheckRun`

NewGithubWebhookCheckRun instantiates a new GithubWebhookCheckRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookCheckRunWithDefaults

`func NewGithubWebhookCheckRunWithDefaults() *GithubWebhookCheckRun`

NewGithubWebhookCheckRunWithDefaults instantiates a new GithubWebhookCheckRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *GithubWebhookCheckRun) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *GithubWebhookCheckRun) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *GithubWebhookCheckRun) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *GithubWebhookCheckRun) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetName

`func (o *GithubWebhookCheckRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubWebhookCheckRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubWebhookCheckRun) SetName(v string)`

SetName sets Name field to given value.


### GetHtmlUrl

`func (o *GithubWebhookCheckRun) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookCheckRun) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookCheckRun) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCheckSuite

`func (o *GithubWebhookCheckRun) GetCheckSuite() GithubCheckSuite`

GetCheckSuite returns the CheckSuite field if non-nil, zero value otherwise.

### GetCheckSuiteOk

`func (o *GithubWebhookCheckRun) GetCheckSuiteOk() (*GithubCheckSuite, bool)`

GetCheckSuiteOk returns a tuple with the CheckSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuite

`func (o *GithubWebhookCheckRun) SetCheckSuite(v GithubCheckSuite)`

SetCheckSuite sets CheckSuite field to given value.


### GetDetailsUrl

`func (o *GithubWebhookCheckRun) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *GithubWebhookCheckRun) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *GithubWebhookCheckRun) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *GithubWebhookCheckRun) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.

### GetOutput

`func (o *GithubWebhookCheckRun) GetOutput() GithubCheckRunOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GithubWebhookCheckRun) GetOutputOk() (*GithubCheckRunOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GithubWebhookCheckRun) SetOutput(v GithubCheckRunOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GithubWebhookCheckRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPullRequests

`func (o *GithubWebhookCheckRun) GetPullRequests() []GithubCheckPullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *GithubWebhookCheckRun) GetPullRequestsOk() (*[]GithubCheckPullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *GithubWebhookCheckRun) SetPullRequests(v []GithubCheckPullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *GithubWebhookCheckRun) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


