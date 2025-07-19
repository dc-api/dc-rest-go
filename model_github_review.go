/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
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

// checks if the GithubReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubReview{}

// GithubReview struct for GithubReview
type GithubReview struct {
	User GithubUser `json:"user"`
	Body NullableString `json:"body,omitempty"`
	HtmlUrl string `json:"html_url"`
	State string `json:"state"`
}

type _GithubReview GithubReview

// NewGithubReview instantiates a new GithubReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubReview(user GithubUser, htmlUrl string, state string) *GithubReview {
	this := GithubReview{}
	this.User = user
	this.HtmlUrl = htmlUrl
	this.State = state
	return &this
}

// NewGithubReviewWithDefaults instantiates a new GithubReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubReviewWithDefaults() *GithubReview {
	this := GithubReview{}
	return &this
}

// GetUser returns the User field value
func (o *GithubReview) GetUser() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GithubReview) GetUserOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GithubReview) SetUser(v GithubUser) {
	o.User = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubReview) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubReview) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body.Get()) {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *GithubReview) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *GithubReview) SetBody(v string) {
	o.Body.Set(&v)
}

// SetBodyNil sets the value for Body to be an explicit nil
func (o *GithubReview) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *GithubReview) UnsetBody() {
	o.Body.Unset()
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GithubReview) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GithubReview) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GithubReview) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetState returns the State field value
func (o *GithubReview) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GithubReview) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GithubReview) SetState(v string) {
	o.State = v
}

func (o GithubReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	toSerialize["html_url"] = o.HtmlUrl
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *GithubReview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"html_url",
		"state",
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

	varGithubReview := _GithubReview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubReview)

	if err != nil {
		return err
	}

	*o = GithubReview(varGithubReview)

	return err
}

type NullableGithubReview struct {
	value *GithubReview
	isSet bool
}

func (v NullableGithubReview) Get() *GithubReview {
	return v.value
}

func (v *NullableGithubReview) Set(val *GithubReview) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubReview) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubReview(val *GithubReview) *NullableGithubReview {
	return &NullableGithubReview{value: val, isSet: true}
}

func (v NullableGithubReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


