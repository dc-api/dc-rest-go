/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the GithubCheckPullRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubCheckPullRequest{}

// GithubCheckPullRequest struct for GithubCheckPullRequest
type GithubCheckPullRequest struct {
	Number int32 `json:"number"`
}

type _GithubCheckPullRequest GithubCheckPullRequest

// NewGithubCheckPullRequest instantiates a new GithubCheckPullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubCheckPullRequest(number int32) *GithubCheckPullRequest {
	this := GithubCheckPullRequest{}
	this.Number = number
	return &this
}

// NewGithubCheckPullRequestWithDefaults instantiates a new GithubCheckPullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubCheckPullRequestWithDefaults() *GithubCheckPullRequest {
	this := GithubCheckPullRequest{}
	return &this
}

// GetNumber returns the Number field value
func (o *GithubCheckPullRequest) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *GithubCheckPullRequest) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *GithubCheckPullRequest) SetNumber(v int32) {
	o.Number = v
}

func (o GithubCheckPullRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubCheckPullRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	return toSerialize, nil
}

func (o *GithubCheckPullRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
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

	varGithubCheckPullRequest := _GithubCheckPullRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubCheckPullRequest)

	if err != nil {
		return err
	}

	*o = GithubCheckPullRequest(varGithubCheckPullRequest)

	return err
}

type NullableGithubCheckPullRequest struct {
	value *GithubCheckPullRequest
	isSet bool
}

func (v NullableGithubCheckPullRequest) Get() *GithubCheckPullRequest {
	return v.value
}

func (v *NullableGithubCheckPullRequest) Set(val *GithubCheckPullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubCheckPullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubCheckPullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubCheckPullRequest(val *GithubCheckPullRequest) *NullableGithubCheckPullRequest {
	return &NullableGithubCheckPullRequest{value: val, isSet: true}
}

func (v NullableGithubCheckPullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubCheckPullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


