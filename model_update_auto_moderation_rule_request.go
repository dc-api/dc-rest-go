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
	"fmt"
)


// UpdateAutoModerationRuleRequest struct for UpdateAutoModerationRuleRequest
type UpdateAutoModerationRuleRequest struct {
	DefaultKeywordListUpsertRequestPartial *DefaultKeywordListUpsertRequestPartial
	KeywordUpsertRequestPartial *KeywordUpsertRequestPartial
	MLSpamUpsertRequestPartial *MLSpamUpsertRequestPartial
	MentionSpamUpsertRequestPartial *MentionSpamUpsertRequestPartial
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpdateAutoModerationRuleRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DefaultKeywordListUpsertRequestPartial
	err = json.Unmarshal(data, &dst.DefaultKeywordListUpsertRequestPartial);
	if err == nil {
		jsonDefaultKeywordListUpsertRequestPartial, _ := json.Marshal(dst.DefaultKeywordListUpsertRequestPartial)
		if string(jsonDefaultKeywordListUpsertRequestPartial) == "{}" { // empty struct
			dst.DefaultKeywordListUpsertRequestPartial = nil
		} else {
			return nil // data stored in dst.DefaultKeywordListUpsertRequestPartial, return on the first match
		}
	} else {
		dst.DefaultKeywordListUpsertRequestPartial = nil
	}

	// try to unmarshal JSON data into KeywordUpsertRequestPartial
	err = json.Unmarshal(data, &dst.KeywordUpsertRequestPartial);
	if err == nil {
		jsonKeywordUpsertRequestPartial, _ := json.Marshal(dst.KeywordUpsertRequestPartial)
		if string(jsonKeywordUpsertRequestPartial) == "{}" { // empty struct
			dst.KeywordUpsertRequestPartial = nil
		} else {
			return nil // data stored in dst.KeywordUpsertRequestPartial, return on the first match
		}
	} else {
		dst.KeywordUpsertRequestPartial = nil
	}

	// try to unmarshal JSON data into MLSpamUpsertRequestPartial
	err = json.Unmarshal(data, &dst.MLSpamUpsertRequestPartial);
	if err == nil {
		jsonMLSpamUpsertRequestPartial, _ := json.Marshal(dst.MLSpamUpsertRequestPartial)
		if string(jsonMLSpamUpsertRequestPartial) == "{}" { // empty struct
			dst.MLSpamUpsertRequestPartial = nil
		} else {
			return nil // data stored in dst.MLSpamUpsertRequestPartial, return on the first match
		}
	} else {
		dst.MLSpamUpsertRequestPartial = nil
	}

	// try to unmarshal JSON data into MentionSpamUpsertRequestPartial
	err = json.Unmarshal(data, &dst.MentionSpamUpsertRequestPartial);
	if err == nil {
		jsonMentionSpamUpsertRequestPartial, _ := json.Marshal(dst.MentionSpamUpsertRequestPartial)
		if string(jsonMentionSpamUpsertRequestPartial) == "{}" { // empty struct
			dst.MentionSpamUpsertRequestPartial = nil
		} else {
			return nil // data stored in dst.MentionSpamUpsertRequestPartial, return on the first match
		}
	} else {
		dst.MentionSpamUpsertRequestPartial = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpdateAutoModerationRuleRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateAutoModerationRuleRequest) MarshalJSON() ([]byte, error) {
	if src.DefaultKeywordListUpsertRequestPartial != nil {
		return json.Marshal(&src.DefaultKeywordListUpsertRequestPartial)
	}

	if src.KeywordUpsertRequestPartial != nil {
		return json.Marshal(&src.KeywordUpsertRequestPartial)
	}

	if src.MLSpamUpsertRequestPartial != nil {
		return json.Marshal(&src.MLSpamUpsertRequestPartial)
	}

	if src.MentionSpamUpsertRequestPartial != nil {
		return json.Marshal(&src.MentionSpamUpsertRequestPartial)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUpdateAutoModerationRuleRequest struct {
	value *UpdateAutoModerationRuleRequest
	isSet bool
}

func (v NullableUpdateAutoModerationRuleRequest) Get() *UpdateAutoModerationRuleRequest {
	return v.value
}

func (v *NullableUpdateAutoModerationRuleRequest) Set(val *UpdateAutoModerationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAutoModerationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAutoModerationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAutoModerationRuleRequest(val *UpdateAutoModerationRuleRequest) *NullableUpdateAutoModerationRuleRequest {
	return &NullableUpdateAutoModerationRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateAutoModerationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAutoModerationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


