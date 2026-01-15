/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Discord HTTP API (Preview) service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package dc_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GithubWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubWebhook{}

// GithubWebhook struct for GithubWebhook
type GithubWebhook struct {
	Action *string `json:"action,omitempty"`
	Ref *string `json:"ref,omitempty"`
	RefType *string `json:"ref_type,omitempty"`
	Comment *GithubWebhookComment `json:"comment,omitempty"`
	Issue *GithubWebhookIssue `json:"issue,omitempty"`
	PullRequest *GithubWebhookIssue `json:"pull_request,omitempty"`
	Repository *GithubWebhookRepository `json:"repository,omitempty"`
	Forkee *GithubWebhookRepository `json:"forkee,omitempty"`
	Sender GithubUser `json:"sender"`
	Member *GithubWebhookMember `json:"member,omitempty"`
	Release *GithubWebhookRelease `json:"release,omitempty"`
	HeadCommit *GithubWebhookHeadCommit `json:"head_commit,omitempty"`
	Commits []GithubCommit `json:"commits,omitempty"`
	Forced *bool `json:"forced,omitempty"`
	Compare *string `json:"compare,omitempty"`
	Review *GithubWebhookReview `json:"review,omitempty"`
	CheckRun *GithubWebhookCheckRun `json:"check_run,omitempty"`
	CheckSuite *GithubWebhookCheckSuite `json:"check_suite,omitempty"`
	Discussion *GithubWebhookDiscussion `json:"discussion,omitempty"`
	Answer *GithubWebhookComment `json:"answer,omitempty"`
}

type _GithubWebhook GithubWebhook

// NewGithubWebhook instantiates a new GithubWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubWebhook(sender GithubUser) *GithubWebhook {
	this := GithubWebhook{}
	this.Sender = sender
	return &this
}

// NewGithubWebhookWithDefaults instantiates a new GithubWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubWebhookWithDefaults() *GithubWebhook {
	this := GithubWebhook{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GithubWebhook) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GithubWebhook) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GithubWebhook) SetAction(v string) {
	o.Action = &v
}


// GetRef returns the Ref field value if set, zero value otherwise.
func (o *GithubWebhook) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *GithubWebhook) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *GithubWebhook) SetRef(v string) {
	o.Ref = &v
}


// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *GithubWebhook) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *GithubWebhook) HasRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *GithubWebhook) SetRefType(v string) {
	o.RefType = &v
}


// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GithubWebhook) GetComment() GithubWebhookComment {
	if o == nil || IsNil(o.Comment) {
		var ret GithubWebhookComment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetCommentOk() (*GithubWebhookComment, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GithubWebhook) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given GithubWebhookComment and assigns it to the Comment field.
func (o *GithubWebhook) SetComment(v GithubWebhookComment) {
	o.Comment = &v
}


// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *GithubWebhook) GetIssue() GithubWebhookIssue {
	if o == nil || IsNil(o.Issue) {
		var ret GithubWebhookIssue
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetIssueOk() (*GithubWebhookIssue, bool) {
	if o == nil || IsNil(o.Issue) {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *GithubWebhook) HasIssue() bool {
	if o != nil && !IsNil(o.Issue) {
		return true
	}

	return false
}

// SetIssue gets a reference to the given GithubWebhookIssue and assigns it to the Issue field.
func (o *GithubWebhook) SetIssue(v GithubWebhookIssue) {
	o.Issue = &v
}


// GetPullRequest returns the PullRequest field value if set, zero value otherwise.
func (o *GithubWebhook) GetPullRequest() GithubWebhookIssue {
	if o == nil || IsNil(o.PullRequest) {
		var ret GithubWebhookIssue
		return ret
	}
	return *o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetPullRequestOk() (*GithubWebhookIssue, bool) {
	if o == nil || IsNil(o.PullRequest) {
		return nil, false
	}
	return o.PullRequest, true
}

// HasPullRequest returns a boolean if a field has been set.
func (o *GithubWebhook) HasPullRequest() bool {
	if o != nil && !IsNil(o.PullRequest) {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given GithubWebhookIssue and assigns it to the PullRequest field.
func (o *GithubWebhook) SetPullRequest(v GithubWebhookIssue) {
	o.PullRequest = &v
}


// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GithubWebhook) GetRepository() GithubWebhookRepository {
	if o == nil || IsNil(o.Repository) {
		var ret GithubWebhookRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetRepositoryOk() (*GithubWebhookRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GithubWebhook) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GithubWebhookRepository and assigns it to the Repository field.
func (o *GithubWebhook) SetRepository(v GithubWebhookRepository) {
	o.Repository = &v
}


// GetForkee returns the Forkee field value if set, zero value otherwise.
func (o *GithubWebhook) GetForkee() GithubWebhookRepository {
	if o == nil || IsNil(o.Forkee) {
		var ret GithubWebhookRepository
		return ret
	}
	return *o.Forkee
}

// GetForkeeOk returns a tuple with the Forkee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetForkeeOk() (*GithubWebhookRepository, bool) {
	if o == nil || IsNil(o.Forkee) {
		return nil, false
	}
	return o.Forkee, true
}

// HasForkee returns a boolean if a field has been set.
func (o *GithubWebhook) HasForkee() bool {
	if o != nil && !IsNil(o.Forkee) {
		return true
	}

	return false
}

// SetForkee gets a reference to the given GithubWebhookRepository and assigns it to the Forkee field.
func (o *GithubWebhook) SetForkee(v GithubWebhookRepository) {
	o.Forkee = &v
}


// GetSender returns the Sender field value
func (o *GithubWebhook) GetSender() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetSenderOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *GithubWebhook) SetSender(v GithubUser) {
	o.Sender = v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *GithubWebhook) GetMember() GithubWebhookMember {
	if o == nil || IsNil(o.Member) {
		var ret GithubWebhookMember
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetMemberOk() (*GithubWebhookMember, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *GithubWebhook) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given GithubWebhookMember and assigns it to the Member field.
func (o *GithubWebhook) SetMember(v GithubWebhookMember) {
	o.Member = &v
}


// GetRelease returns the Release field value if set, zero value otherwise.
func (o *GithubWebhook) GetRelease() GithubWebhookRelease {
	if o == nil || IsNil(o.Release) {
		var ret GithubWebhookRelease
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetReleaseOk() (*GithubWebhookRelease, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *GithubWebhook) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given GithubWebhookRelease and assigns it to the Release field.
func (o *GithubWebhook) SetRelease(v GithubWebhookRelease) {
	o.Release = &v
}


// GetHeadCommit returns the HeadCommit field value if set, zero value otherwise.
func (o *GithubWebhook) GetHeadCommit() GithubWebhookHeadCommit {
	if o == nil || IsNil(o.HeadCommit) {
		var ret GithubWebhookHeadCommit
		return ret
	}
	return *o.HeadCommit
}

// GetHeadCommitOk returns a tuple with the HeadCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetHeadCommitOk() (*GithubWebhookHeadCommit, bool) {
	if o == nil || IsNil(o.HeadCommit) {
		return nil, false
	}
	return o.HeadCommit, true
}

// HasHeadCommit returns a boolean if a field has been set.
func (o *GithubWebhook) HasHeadCommit() bool {
	if o != nil && !IsNil(o.HeadCommit) {
		return true
	}

	return false
}

// SetHeadCommit gets a reference to the given GithubWebhookHeadCommit and assigns it to the HeadCommit field.
func (o *GithubWebhook) SetHeadCommit(v GithubWebhookHeadCommit) {
	o.HeadCommit = &v
}


// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *GithubWebhook) GetCommits() []GithubCommit {
	if o == nil || IsNil(o.Commits) {
		var ret []GithubCommit
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetCommitsOk() ([]GithubCommit, bool) {
	if o == nil || IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *GithubWebhook) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []GithubCommit and assigns it to the Commits field.
func (o *GithubWebhook) SetCommits(v []GithubCommit) {
	o.Commits = v
}


// GetForced returns the Forced field value if set, zero value otherwise.
func (o *GithubWebhook) GetForced() bool {
	if o == nil || IsNil(o.Forced) {
		var ret bool
		return ret
	}
	return *o.Forced
}

// GetForcedOk returns a tuple with the Forced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetForcedOk() (*bool, bool) {
	if o == nil || IsNil(o.Forced) {
		return nil, false
	}
	return o.Forced, true
}

// HasForced returns a boolean if a field has been set.
func (o *GithubWebhook) HasForced() bool {
	if o != nil && !IsNil(o.Forced) {
		return true
	}

	return false
}

// SetForced gets a reference to the given bool and assigns it to the Forced field.
func (o *GithubWebhook) SetForced(v bool) {
	o.Forced = &v
}


// GetCompare returns the Compare field value if set, zero value otherwise.
func (o *GithubWebhook) GetCompare() string {
	if o == nil || IsNil(o.Compare) {
		var ret string
		return ret
	}
	return *o.Compare
}

// GetCompareOk returns a tuple with the Compare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetCompareOk() (*string, bool) {
	if o == nil || IsNil(o.Compare) {
		return nil, false
	}
	return o.Compare, true
}

// HasCompare returns a boolean if a field has been set.
func (o *GithubWebhook) HasCompare() bool {
	if o != nil && !IsNil(o.Compare) {
		return true
	}

	return false
}

// SetCompare gets a reference to the given string and assigns it to the Compare field.
func (o *GithubWebhook) SetCompare(v string) {
	o.Compare = &v
}


// GetReview returns the Review field value if set, zero value otherwise.
func (o *GithubWebhook) GetReview() GithubWebhookReview {
	if o == nil || IsNil(o.Review) {
		var ret GithubWebhookReview
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetReviewOk() (*GithubWebhookReview, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *GithubWebhook) HasReview() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given GithubWebhookReview and assigns it to the Review field.
func (o *GithubWebhook) SetReview(v GithubWebhookReview) {
	o.Review = &v
}


// GetCheckRun returns the CheckRun field value if set, zero value otherwise.
func (o *GithubWebhook) GetCheckRun() GithubWebhookCheckRun {
	if o == nil || IsNil(o.CheckRun) {
		var ret GithubWebhookCheckRun
		return ret
	}
	return *o.CheckRun
}

// GetCheckRunOk returns a tuple with the CheckRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetCheckRunOk() (*GithubWebhookCheckRun, bool) {
	if o == nil || IsNil(o.CheckRun) {
		return nil, false
	}
	return o.CheckRun, true
}

// HasCheckRun returns a boolean if a field has been set.
func (o *GithubWebhook) HasCheckRun() bool {
	if o != nil && !IsNil(o.CheckRun) {
		return true
	}

	return false
}

// SetCheckRun gets a reference to the given GithubWebhookCheckRun and assigns it to the CheckRun field.
func (o *GithubWebhook) SetCheckRun(v GithubWebhookCheckRun) {
	o.CheckRun = &v
}


// GetCheckSuite returns the CheckSuite field value if set, zero value otherwise.
func (o *GithubWebhook) GetCheckSuite() GithubWebhookCheckSuite {
	if o == nil || IsNil(o.CheckSuite) {
		var ret GithubWebhookCheckSuite
		return ret
	}
	return *o.CheckSuite
}

// GetCheckSuiteOk returns a tuple with the CheckSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetCheckSuiteOk() (*GithubWebhookCheckSuite, bool) {
	if o == nil || IsNil(o.CheckSuite) {
		return nil, false
	}
	return o.CheckSuite, true
}

// HasCheckSuite returns a boolean if a field has been set.
func (o *GithubWebhook) HasCheckSuite() bool {
	if o != nil && !IsNil(o.CheckSuite) {
		return true
	}

	return false
}

// SetCheckSuite gets a reference to the given GithubWebhookCheckSuite and assigns it to the CheckSuite field.
func (o *GithubWebhook) SetCheckSuite(v GithubWebhookCheckSuite) {
	o.CheckSuite = &v
}


// GetDiscussion returns the Discussion field value if set, zero value otherwise.
func (o *GithubWebhook) GetDiscussion() GithubWebhookDiscussion {
	if o == nil || IsNil(o.Discussion) {
		var ret GithubWebhookDiscussion
		return ret
	}
	return *o.Discussion
}

// GetDiscussionOk returns a tuple with the Discussion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetDiscussionOk() (*GithubWebhookDiscussion, bool) {
	if o == nil || IsNil(o.Discussion) {
		return nil, false
	}
	return o.Discussion, true
}

// HasDiscussion returns a boolean if a field has been set.
func (o *GithubWebhook) HasDiscussion() bool {
	if o != nil && !IsNil(o.Discussion) {
		return true
	}

	return false
}

// SetDiscussion gets a reference to the given GithubWebhookDiscussion and assigns it to the Discussion field.
func (o *GithubWebhook) SetDiscussion(v GithubWebhookDiscussion) {
	o.Discussion = &v
}


// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *GithubWebhook) GetAnswer() GithubWebhookComment {
	if o == nil || IsNil(o.Answer) {
		var ret GithubWebhookComment
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetAnswerOk() (*GithubWebhookComment, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *GithubWebhook) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given GithubWebhookComment and assigns it to the Answer field.
func (o *GithubWebhook) SetAnswer(v GithubWebhookComment) {
	o.Answer = &v
}


func (o GithubWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.RefType) {
		toSerialize["ref_type"] = o.RefType
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Issue) {
		toSerialize["issue"] = o.Issue
	}
	if !IsNil(o.PullRequest) {
		toSerialize["pull_request"] = o.PullRequest
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Forkee) {
		toSerialize["forkee"] = o.Forkee
	}
	toSerialize["sender"] = o.Sender
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.HeadCommit) {
		toSerialize["head_commit"] = o.HeadCommit
	}
	if !IsNil(o.Commits) {
		toSerialize["commits"] = o.Commits
	}
	if !IsNil(o.Forced) {
		toSerialize["forced"] = o.Forced
	}
	if !IsNil(o.Compare) {
		toSerialize["compare"] = o.Compare
	}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !IsNil(o.CheckRun) {
		toSerialize["check_run"] = o.CheckRun
	}
	if !IsNil(o.CheckSuite) {
		toSerialize["check_suite"] = o.CheckSuite
	}
	if !IsNil(o.Discussion) {
		toSerialize["discussion"] = o.Discussion
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	return toSerialize, nil
}

func (o *GithubWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGithubWebhook := _GithubWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubWebhook)

	if err != nil {
		return err
	}

	*o = GithubWebhook(varGithubWebhook)

	return err
}

type NullableGithubWebhook struct {
	value *GithubWebhook
	isSet bool
}

func (v NullableGithubWebhook) Get() *GithubWebhook {
	return v.value
}

func (v *NullableGithubWebhook) Set(val *GithubWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubWebhook(val *GithubWebhook) *NullableGithubWebhook {
	return &NullableGithubWebhook{value: val, isSet: true}
}

func (v NullableGithubWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

