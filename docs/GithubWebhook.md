# GithubWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** |  | [optional] 
**Ref** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to [**NullableGithubComment**](GithubComment.md) |  | [optional] 
**Issue** | Pointer to [**NullableGithubIssue**](GithubIssue.md) |  | [optional] 
**PullRequest** | Pointer to [**NullableGithubIssue**](GithubIssue.md) |  | [optional] 
**Repository** | Pointer to [**NullableGithubRepository**](GithubRepository.md) |  | [optional] 
**Forkee** | Pointer to [**NullableGithubRepository**](GithubRepository.md) |  | [optional] 
**Sender** | [**GithubUser**](GithubUser.md) |  | 
**Member** | Pointer to [**NullableGithubUser**](GithubUser.md) |  | [optional] 
**Release** | Pointer to [**NullableGithubRelease**](GithubRelease.md) |  | [optional] 
**HeadCommit** | Pointer to [**NullableGithubCommit**](GithubCommit.md) |  | [optional] 
**Commits** | Pointer to [**[]GithubCommit**](GithubCommit.md) |  | [optional] 
**Forced** | Pointer to **NullableBool** |  | [optional] 
**Compare** | Pointer to **NullableString** |  | [optional] 
**Review** | Pointer to [**NullableGithubReview**](GithubReview.md) |  | [optional] 
**CheckRun** | Pointer to [**NullableGithubCheckRun**](GithubCheckRun.md) |  | [optional] 
**CheckSuite** | Pointer to [**NullableGithubCheckSuite**](GithubCheckSuite.md) |  | [optional] 
**Discussion** | Pointer to [**NullableGithubDiscussion**](GithubDiscussion.md) |  | [optional] 
**Answer** | Pointer to [**NullableGithubComment**](GithubComment.md) |  | [optional] 

## Methods

### NewGithubWebhook

`func NewGithubWebhook(sender GithubUser, ) *GithubWebhook`

NewGithubWebhook instantiates a new GithubWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookWithDefaults

`func NewGithubWebhookWithDefaults() *GithubWebhook`

NewGithubWebhookWithDefaults instantiates a new GithubWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GithubWebhook) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GithubWebhook) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GithubWebhook) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GithubWebhook) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *GithubWebhook) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *GithubWebhook) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetRef

`func (o *GithubWebhook) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GithubWebhook) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GithubWebhook) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GithubWebhook) HasRef() bool`

HasRef returns a boolean if a field has been set.

### SetRefNil

`func (o *GithubWebhook) SetRefNil(b bool)`

 SetRefNil sets the value for Ref to be an explicit nil

### UnsetRef
`func (o *GithubWebhook) UnsetRef()`

UnsetRef ensures that no value is present for Ref, not even an explicit nil
### GetRefType

`func (o *GithubWebhook) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GithubWebhook) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GithubWebhook) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GithubWebhook) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GithubWebhook) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GithubWebhook) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetComment

`func (o *GithubWebhook) GetComment() GithubComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GithubWebhook) GetCommentOk() (*GithubComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GithubWebhook) SetComment(v GithubComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GithubWebhook) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *GithubWebhook) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *GithubWebhook) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIssue

`func (o *GithubWebhook) GetIssue() GithubIssue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *GithubWebhook) GetIssueOk() (*GithubIssue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *GithubWebhook) SetIssue(v GithubIssue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *GithubWebhook) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### SetIssueNil

`func (o *GithubWebhook) SetIssueNil(b bool)`

 SetIssueNil sets the value for Issue to be an explicit nil

### UnsetIssue
`func (o *GithubWebhook) UnsetIssue()`

UnsetIssue ensures that no value is present for Issue, not even an explicit nil
### GetPullRequest

`func (o *GithubWebhook) GetPullRequest() GithubIssue`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *GithubWebhook) GetPullRequestOk() (*GithubIssue, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *GithubWebhook) SetPullRequest(v GithubIssue)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *GithubWebhook) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### SetPullRequestNil

`func (o *GithubWebhook) SetPullRequestNil(b bool)`

 SetPullRequestNil sets the value for PullRequest to be an explicit nil

### UnsetPullRequest
`func (o *GithubWebhook) UnsetPullRequest()`

UnsetPullRequest ensures that no value is present for PullRequest, not even an explicit nil
### GetRepository

`func (o *GithubWebhook) GetRepository() GithubRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GithubWebhook) GetRepositoryOk() (*GithubRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GithubWebhook) SetRepository(v GithubRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GithubWebhook) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *GithubWebhook) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *GithubWebhook) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetForkee

`func (o *GithubWebhook) GetForkee() GithubRepository`

GetForkee returns the Forkee field if non-nil, zero value otherwise.

### GetForkeeOk

`func (o *GithubWebhook) GetForkeeOk() (*GithubRepository, bool)`

GetForkeeOk returns a tuple with the Forkee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkee

`func (o *GithubWebhook) SetForkee(v GithubRepository)`

SetForkee sets Forkee field to given value.

### HasForkee

`func (o *GithubWebhook) HasForkee() bool`

HasForkee returns a boolean if a field has been set.

### SetForkeeNil

`func (o *GithubWebhook) SetForkeeNil(b bool)`

 SetForkeeNil sets the value for Forkee to be an explicit nil

### UnsetForkee
`func (o *GithubWebhook) UnsetForkee()`

UnsetForkee ensures that no value is present for Forkee, not even an explicit nil
### GetSender

`func (o *GithubWebhook) GetSender() GithubUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GithubWebhook) GetSenderOk() (*GithubUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GithubWebhook) SetSender(v GithubUser)`

SetSender sets Sender field to given value.


### GetMember

`func (o *GithubWebhook) GetMember() GithubUser`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GithubWebhook) GetMemberOk() (*GithubUser, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GithubWebhook) SetMember(v GithubUser)`

SetMember sets Member field to given value.

### HasMember

`func (o *GithubWebhook) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *GithubWebhook) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *GithubWebhook) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetRelease

`func (o *GithubWebhook) GetRelease() GithubRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *GithubWebhook) GetReleaseOk() (*GithubRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *GithubWebhook) SetRelease(v GithubRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *GithubWebhook) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *GithubWebhook) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *GithubWebhook) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetHeadCommit

`func (o *GithubWebhook) GetHeadCommit() GithubCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *GithubWebhook) GetHeadCommitOk() (*GithubCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *GithubWebhook) SetHeadCommit(v GithubCommit)`

SetHeadCommit sets HeadCommit field to given value.

### HasHeadCommit

`func (o *GithubWebhook) HasHeadCommit() bool`

HasHeadCommit returns a boolean if a field has been set.

### SetHeadCommitNil

`func (o *GithubWebhook) SetHeadCommitNil(b bool)`

 SetHeadCommitNil sets the value for HeadCommit to be an explicit nil

### UnsetHeadCommit
`func (o *GithubWebhook) UnsetHeadCommit()`

UnsetHeadCommit ensures that no value is present for HeadCommit, not even an explicit nil
### GetCommits

`func (o *GithubWebhook) GetCommits() []GithubCommit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GithubWebhook) GetCommitsOk() (*[]GithubCommit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GithubWebhook) SetCommits(v []GithubCommit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GithubWebhook) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### SetCommitsNil

`func (o *GithubWebhook) SetCommitsNil(b bool)`

 SetCommitsNil sets the value for Commits to be an explicit nil

### UnsetCommits
`func (o *GithubWebhook) UnsetCommits()`

UnsetCommits ensures that no value is present for Commits, not even an explicit nil
### GetForced

`func (o *GithubWebhook) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *GithubWebhook) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *GithubWebhook) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *GithubWebhook) HasForced() bool`

HasForced returns a boolean if a field has been set.

### SetForcedNil

`func (o *GithubWebhook) SetForcedNil(b bool)`

 SetForcedNil sets the value for Forced to be an explicit nil

### UnsetForced
`func (o *GithubWebhook) UnsetForced()`

UnsetForced ensures that no value is present for Forced, not even an explicit nil
### GetCompare

`func (o *GithubWebhook) GetCompare() string`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *GithubWebhook) GetCompareOk() (*string, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *GithubWebhook) SetCompare(v string)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *GithubWebhook) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### SetCompareNil

`func (o *GithubWebhook) SetCompareNil(b bool)`

 SetCompareNil sets the value for Compare to be an explicit nil

### UnsetCompare
`func (o *GithubWebhook) UnsetCompare()`

UnsetCompare ensures that no value is present for Compare, not even an explicit nil
### GetReview

`func (o *GithubWebhook) GetReview() GithubReview`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *GithubWebhook) GetReviewOk() (*GithubReview, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *GithubWebhook) SetReview(v GithubReview)`

SetReview sets Review field to given value.

### HasReview

`func (o *GithubWebhook) HasReview() bool`

HasReview returns a boolean if a field has been set.

### SetReviewNil

`func (o *GithubWebhook) SetReviewNil(b bool)`

 SetReviewNil sets the value for Review to be an explicit nil

### UnsetReview
`func (o *GithubWebhook) UnsetReview()`

UnsetReview ensures that no value is present for Review, not even an explicit nil
### GetCheckRun

`func (o *GithubWebhook) GetCheckRun() GithubCheckRun`

GetCheckRun returns the CheckRun field if non-nil, zero value otherwise.

### GetCheckRunOk

`func (o *GithubWebhook) GetCheckRunOk() (*GithubCheckRun, bool)`

GetCheckRunOk returns a tuple with the CheckRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRun

`func (o *GithubWebhook) SetCheckRun(v GithubCheckRun)`

SetCheckRun sets CheckRun field to given value.

### HasCheckRun

`func (o *GithubWebhook) HasCheckRun() bool`

HasCheckRun returns a boolean if a field has been set.

### SetCheckRunNil

`func (o *GithubWebhook) SetCheckRunNil(b bool)`

 SetCheckRunNil sets the value for CheckRun to be an explicit nil

### UnsetCheckRun
`func (o *GithubWebhook) UnsetCheckRun()`

UnsetCheckRun ensures that no value is present for CheckRun, not even an explicit nil
### GetCheckSuite

`func (o *GithubWebhook) GetCheckSuite() GithubCheckSuite`

GetCheckSuite returns the CheckSuite field if non-nil, zero value otherwise.

### GetCheckSuiteOk

`func (o *GithubWebhook) GetCheckSuiteOk() (*GithubCheckSuite, bool)`

GetCheckSuiteOk returns a tuple with the CheckSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuite

`func (o *GithubWebhook) SetCheckSuite(v GithubCheckSuite)`

SetCheckSuite sets CheckSuite field to given value.

### HasCheckSuite

`func (o *GithubWebhook) HasCheckSuite() bool`

HasCheckSuite returns a boolean if a field has been set.

### SetCheckSuiteNil

`func (o *GithubWebhook) SetCheckSuiteNil(b bool)`

 SetCheckSuiteNil sets the value for CheckSuite to be an explicit nil

### UnsetCheckSuite
`func (o *GithubWebhook) UnsetCheckSuite()`

UnsetCheckSuite ensures that no value is present for CheckSuite, not even an explicit nil
### GetDiscussion

`func (o *GithubWebhook) GetDiscussion() GithubDiscussion`

GetDiscussion returns the Discussion field if non-nil, zero value otherwise.

### GetDiscussionOk

`func (o *GithubWebhook) GetDiscussionOk() (*GithubDiscussion, bool)`

GetDiscussionOk returns a tuple with the Discussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussion

`func (o *GithubWebhook) SetDiscussion(v GithubDiscussion)`

SetDiscussion sets Discussion field to given value.

### HasDiscussion

`func (o *GithubWebhook) HasDiscussion() bool`

HasDiscussion returns a boolean if a field has been set.

### SetDiscussionNil

`func (o *GithubWebhook) SetDiscussionNil(b bool)`

 SetDiscussionNil sets the value for Discussion to be an explicit nil

### UnsetDiscussion
`func (o *GithubWebhook) UnsetDiscussion()`

UnsetDiscussion ensures that no value is present for Discussion, not even an explicit nil
### GetAnswer

`func (o *GithubWebhook) GetAnswer() GithubComment`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GithubWebhook) GetAnswerOk() (*GithubComment, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GithubWebhook) SetAnswer(v GithubComment)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GithubWebhook) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### SetAnswerNil

`func (o *GithubWebhook) SetAnswerNil(b bool)`

 SetAnswerNil sets the value for Answer to be an explicit nil

### UnsetAnswer
`func (o *GithubWebhook) UnsetAnswer()`

UnsetAnswer ensures that no value is present for Answer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


