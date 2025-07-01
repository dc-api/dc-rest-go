# GithubCheckRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CheckSuite** | [**GithubCheckSuite**](GithubCheckSuite.md) |  | 
**DetailsUrl** | Pointer to **NullableString** |  | [optional] 
**Output** | Pointer to [**NullableGithubCheckRunOutput**](GithubCheckRunOutput.md) |  | [optional] 
**PullRequests** | Pointer to [**[]GithubCheckPullRequest**](GithubCheckPullRequest.md) |  | [optional] 

## Methods

### NewGithubCheckRun

`func NewGithubCheckRun(name string, htmlUrl string, checkSuite GithubCheckSuite, ) *GithubCheckRun`

NewGithubCheckRun instantiates a new GithubCheckRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubCheckRunWithDefaults

`func NewGithubCheckRunWithDefaults() *GithubCheckRun`

NewGithubCheckRunWithDefaults instantiates a new GithubCheckRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *GithubCheckRun) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *GithubCheckRun) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *GithubCheckRun) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *GithubCheckRun) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### SetConclusionNil

`func (o *GithubCheckRun) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *GithubCheckRun) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetName

`func (o *GithubCheckRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubCheckRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubCheckRun) SetName(v string)`

SetName sets Name field to given value.


### GetHtmlUrl

`func (o *GithubCheckRun) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubCheckRun) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubCheckRun) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCheckSuite

`func (o *GithubCheckRun) GetCheckSuite() GithubCheckSuite`

GetCheckSuite returns the CheckSuite field if non-nil, zero value otherwise.

### GetCheckSuiteOk

`func (o *GithubCheckRun) GetCheckSuiteOk() (*GithubCheckSuite, bool)`

GetCheckSuiteOk returns a tuple with the CheckSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuite

`func (o *GithubCheckRun) SetCheckSuite(v GithubCheckSuite)`

SetCheckSuite sets CheckSuite field to given value.


### GetDetailsUrl

`func (o *GithubCheckRun) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *GithubCheckRun) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *GithubCheckRun) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *GithubCheckRun) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.

### SetDetailsUrlNil

`func (o *GithubCheckRun) SetDetailsUrlNil(b bool)`

 SetDetailsUrlNil sets the value for DetailsUrl to be an explicit nil

### UnsetDetailsUrl
`func (o *GithubCheckRun) UnsetDetailsUrl()`

UnsetDetailsUrl ensures that no value is present for DetailsUrl, not even an explicit nil
### GetOutput

`func (o *GithubCheckRun) GetOutput() GithubCheckRunOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GithubCheckRun) GetOutputOk() (*GithubCheckRunOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GithubCheckRun) SetOutput(v GithubCheckRunOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GithubCheckRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *GithubCheckRun) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *GithubCheckRun) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPullRequests

`func (o *GithubCheckRun) GetPullRequests() []GithubCheckPullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *GithubCheckRun) GetPullRequestsOk() (*[]GithubCheckPullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *GithubCheckRun) SetPullRequests(v []GithubCheckPullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *GithubCheckRun) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### SetPullRequestsNil

`func (o *GithubCheckRun) SetPullRequestsNil(b bool)`

 SetPullRequestsNil sets the value for PullRequests to be an explicit nil

### UnsetPullRequests
`func (o *GithubCheckRun) UnsetPullRequests()`

UnsetPullRequests ensures that no value is present for PullRequests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


