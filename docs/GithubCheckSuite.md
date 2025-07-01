# GithubCheckSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **NullableString** |  | [optional] 
**HeadBranch** | Pointer to **NullableString** |  | [optional] 
**HeadSha** | **string** |  | 
**PullRequests** | Pointer to [**[]GithubCheckPullRequest**](GithubCheckPullRequest.md) |  | [optional] 
**App** | [**GithubCheckApp**](GithubCheckApp.md) |  | 

## Methods

### NewGithubCheckSuite

`func NewGithubCheckSuite(headSha string, app GithubCheckApp, ) *GithubCheckSuite`

NewGithubCheckSuite instantiates a new GithubCheckSuite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubCheckSuiteWithDefaults

`func NewGithubCheckSuiteWithDefaults() *GithubCheckSuite`

NewGithubCheckSuiteWithDefaults instantiates a new GithubCheckSuite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *GithubCheckSuite) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *GithubCheckSuite) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *GithubCheckSuite) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *GithubCheckSuite) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### SetConclusionNil

`func (o *GithubCheckSuite) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *GithubCheckSuite) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetHeadBranch

`func (o *GithubCheckSuite) GetHeadBranch() string`

GetHeadBranch returns the HeadBranch field if non-nil, zero value otherwise.

### GetHeadBranchOk

`func (o *GithubCheckSuite) GetHeadBranchOk() (*string, bool)`

GetHeadBranchOk returns a tuple with the HeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadBranch

`func (o *GithubCheckSuite) SetHeadBranch(v string)`

SetHeadBranch sets HeadBranch field to given value.

### HasHeadBranch

`func (o *GithubCheckSuite) HasHeadBranch() bool`

HasHeadBranch returns a boolean if a field has been set.

### SetHeadBranchNil

`func (o *GithubCheckSuite) SetHeadBranchNil(b bool)`

 SetHeadBranchNil sets the value for HeadBranch to be an explicit nil

### UnsetHeadBranch
`func (o *GithubCheckSuite) UnsetHeadBranch()`

UnsetHeadBranch ensures that no value is present for HeadBranch, not even an explicit nil
### GetHeadSha

`func (o *GithubCheckSuite) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *GithubCheckSuite) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *GithubCheckSuite) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetPullRequests

`func (o *GithubCheckSuite) GetPullRequests() []GithubCheckPullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *GithubCheckSuite) GetPullRequestsOk() (*[]GithubCheckPullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *GithubCheckSuite) SetPullRequests(v []GithubCheckPullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *GithubCheckSuite) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### SetPullRequestsNil

`func (o *GithubCheckSuite) SetPullRequestsNil(b bool)`

 SetPullRequestsNil sets the value for PullRequests to be an explicit nil

### UnsetPullRequests
`func (o *GithubCheckSuite) UnsetPullRequests()`

UnsetPullRequests ensures that no value is present for PullRequests, not even an explicit nil
### GetApp

`func (o *GithubCheckSuite) GetApp() GithubCheckApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GithubCheckSuite) GetAppOk() (*GithubCheckApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GithubCheckSuite) SetApp(v GithubCheckApp)`

SetApp sets App field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


