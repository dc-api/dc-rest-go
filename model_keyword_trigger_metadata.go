/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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
)

// checks if the KeywordTriggerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordTriggerMetadata{}

// KeywordTriggerMetadata struct for KeywordTriggerMetadata
type KeywordTriggerMetadata struct {
	KeywordFilter []string `json:"keyword_filter,omitempty"`
	RegexPatterns []string `json:"regex_patterns,omitempty"`
	AllowList []string `json:"allow_list,omitempty"`
}

// NewKeywordTriggerMetadata instantiates a new KeywordTriggerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordTriggerMetadata() *KeywordTriggerMetadata {
	this := KeywordTriggerMetadata{}
	return &this
}

// NewKeywordTriggerMetadataWithDefaults instantiates a new KeywordTriggerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordTriggerMetadataWithDefaults() *KeywordTriggerMetadata {
	this := KeywordTriggerMetadata{}
	return &this
}

// GetKeywordFilter returns the KeywordFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordTriggerMetadata) GetKeywordFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.KeywordFilter
}

// GetKeywordFilterOk returns a tuple with the KeywordFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordTriggerMetadata) GetKeywordFilterOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeywordFilter, true
}

// HasKeywordFilter returns a boolean if a field has been set.
func (o *KeywordTriggerMetadata) HasKeywordFilter() bool {
	if o != nil && !IsNil(o.KeywordFilter) {
		return true
	}

	return false
}

// SetKeywordFilter gets a reference to the given []string and assigns it to the KeywordFilter field.
func (o *KeywordTriggerMetadata) SetKeywordFilter(v []string) {
	o.KeywordFilter = v
}


// GetRegexPatterns returns the RegexPatterns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordTriggerMetadata) GetRegexPatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RegexPatterns
}

// GetRegexPatternsOk returns a tuple with the RegexPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordTriggerMetadata) GetRegexPatternsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegexPatterns, true
}

// HasRegexPatterns returns a boolean if a field has been set.
func (o *KeywordTriggerMetadata) HasRegexPatterns() bool {
	if o != nil && !IsNil(o.RegexPatterns) {
		return true
	}

	return false
}

// SetRegexPatterns gets a reference to the given []string and assigns it to the RegexPatterns field.
func (o *KeywordTriggerMetadata) SetRegexPatterns(v []string) {
	o.RegexPatterns = v
}


// GetAllowList returns the AllowList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordTriggerMetadata) GetAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordTriggerMetadata) GetAllowListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *KeywordTriggerMetadata) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *KeywordTriggerMetadata) SetAllowList(v []string) {
	o.AllowList = v
}


func (o KeywordTriggerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeywordTriggerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.KeywordFilter != nil {
		toSerialize["keyword_filter"] = o.KeywordFilter
	}
	if o.RegexPatterns != nil {
		toSerialize["regex_patterns"] = o.RegexPatterns
	}
	if o.AllowList != nil {
		toSerialize["allow_list"] = o.AllowList
	}
	return toSerialize, nil
}

type NullableKeywordTriggerMetadata struct {
	value *KeywordTriggerMetadata
	isSet bool
}

func (v NullableKeywordTriggerMetadata) Get() *KeywordTriggerMetadata {
	return v.value
}

func (v *NullableKeywordTriggerMetadata) Set(val *KeywordTriggerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordTriggerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordTriggerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordTriggerMetadata(val *KeywordTriggerMetadata) *NullableKeywordTriggerMetadata {
	return &NullableKeywordTriggerMetadata{value: val, isSet: true}
}

func (v NullableKeywordTriggerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeywordTriggerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


