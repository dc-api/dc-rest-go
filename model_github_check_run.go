/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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

// checks if the GithubCheckRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubCheckRun{}

// GithubCheckRun struct for GithubCheckRun
type GithubCheckRun struct {
	Conclusion *string `json:"conclusion,omitempty"`
	Name string `json:"name"`
	HtmlUrl string `json:"html_url"`
	CheckSuite GithubCheckSuite `json:"check_suite"`
	DetailsUrl *string `json:"details_url,omitempty"`
	Output *GithubCheckRunOutput `json:"output,omitempty"`
	PullRequests []GithubCheckPullRequest `json:"pull_requests,omitempty"`
}

type _GithubCheckRun GithubCheckRun

// NewGithubCheckRun instantiates a new GithubCheckRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubCheckRun(name string, htmlUrl string, checkSuite GithubCheckSuite) *GithubCheckRun {
	this := GithubCheckRun{}
	this.Name = name
	this.HtmlUrl = htmlUrl
	this.CheckSuite = checkSuite
	return &this
}

// NewGithubCheckRunWithDefaults instantiates a new GithubCheckRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubCheckRunWithDefaults() *GithubCheckRun {
	this := GithubCheckRun{}
	return &this
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise.
func (o *GithubCheckRun) GetConclusion() string {
	if o == nil || IsNil(o.Conclusion) {
		var ret string
		return ret
	}
	return *o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetConclusionOk() (*string, bool) {
	if o == nil || IsNil(o.Conclusion) {
		return nil, false
	}
	return o.Conclusion, true
}

// HasConclusion returns a boolean if a field has been set.
func (o *GithubCheckRun) HasConclusion() bool {
	if o != nil && !IsNil(o.Conclusion) {
		return true
	}

	return false
}

// SetConclusion gets a reference to the given string and assigns it to the Conclusion field.
func (o *GithubCheckRun) SetConclusion(v string) {
	o.Conclusion = &v
}


// GetName returns the Name field value
func (o *GithubCheckRun) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GithubCheckRun) SetName(v string) {
	o.Name = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GithubCheckRun) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GithubCheckRun) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetCheckSuite returns the CheckSuite field value
func (o *GithubCheckRun) GetCheckSuite() GithubCheckSuite {
	if o == nil {
		var ret GithubCheckSuite
		return ret
	}

	return o.CheckSuite
}

// GetCheckSuiteOk returns a tuple with the CheckSuite field value
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetCheckSuiteOk() (*GithubCheckSuite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckSuite, true
}

// SetCheckSuite sets field value
func (o *GithubCheckRun) SetCheckSuite(v GithubCheckSuite) {
	o.CheckSuite = v
}

// GetDetailsUrl returns the DetailsUrl field value if set, zero value otherwise.
func (o *GithubCheckRun) GetDetailsUrl() string {
	if o == nil || IsNil(o.DetailsUrl) {
		var ret string
		return ret
	}
	return *o.DetailsUrl
}

// GetDetailsUrlOk returns a tuple with the DetailsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetDetailsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DetailsUrl) {
		return nil, false
	}
	return o.DetailsUrl, true
}

// HasDetailsUrl returns a boolean if a field has been set.
func (o *GithubCheckRun) HasDetailsUrl() bool {
	if o != nil && !IsNil(o.DetailsUrl) {
		return true
	}

	return false
}

// SetDetailsUrl gets a reference to the given string and assigns it to the DetailsUrl field.
func (o *GithubCheckRun) SetDetailsUrl(v string) {
	o.DetailsUrl = &v
}


// GetOutput returns the Output field value if set, zero value otherwise.
func (o *GithubCheckRun) GetOutput() GithubCheckRunOutput {
	if o == nil || IsNil(o.Output) {
		var ret GithubCheckRunOutput
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetOutputOk() (*GithubCheckRunOutput, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *GithubCheckRun) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given GithubCheckRunOutput and assigns it to the Output field.
func (o *GithubCheckRun) SetOutput(v GithubCheckRunOutput) {
	o.Output = &v
}


// GetPullRequests returns the PullRequests field value if set, zero value otherwise.
func (o *GithubCheckRun) GetPullRequests() []GithubCheckPullRequest {
	if o == nil || IsNil(o.PullRequests) {
		var ret []GithubCheckPullRequest
		return ret
	}
	return o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubCheckRun) GetPullRequestsOk() ([]GithubCheckPullRequest, bool) {
	if o == nil || IsNil(o.PullRequests) {
		return nil, false
	}
	return o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *GithubCheckRun) HasPullRequests() bool {
	if o != nil && !IsNil(o.PullRequests) {
		return true
	}

	return false
}

// SetPullRequests gets a reference to the given []GithubCheckPullRequest and assigns it to the PullRequests field.
func (o *GithubCheckRun) SetPullRequests(v []GithubCheckPullRequest) {
	o.PullRequests = v
}


func (o GithubCheckRun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubCheckRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conclusion) {
		toSerialize["conclusion"] = o.Conclusion
	}
	toSerialize["name"] = o.Name
	toSerialize["html_url"] = o.HtmlUrl
	toSerialize["check_suite"] = o.CheckSuite
	if !IsNil(o.DetailsUrl) {
		toSerialize["details_url"] = o.DetailsUrl
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.PullRequests) {
		toSerialize["pull_requests"] = o.PullRequests
	}
	return toSerialize, nil
}

func (o *GithubCheckRun) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"html_url",
		"check_suite",
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

	varGithubCheckRun := _GithubCheckRun{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubCheckRun)

	if err != nil {
		return err
	}

	*o = GithubCheckRun(varGithubCheckRun)

	return err
}

type NullableGithubCheckRun struct {
	value *GithubCheckRun
	isSet bool
}

func (v NullableGithubCheckRun) Get() *GithubCheckRun {
	return v.value
}

func (v *NullableGithubCheckRun) Set(val *GithubCheckRun) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubCheckRun) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubCheckRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubCheckRun(val *GithubCheckRun) *NullableGithubCheckRun {
	return &NullableGithubCheckRun{value: val, isSet: true}
}

func (v NullableGithubCheckRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubCheckRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

