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
	"gopkg.in/validator.v2"
)

// CreateAutoModerationRuleRequest - struct for CreateAutoModerationRuleRequest
type CreateAutoModerationRuleRequest struct {
	DefaultKeywordListUpsertRequest *DefaultKeywordListUpsertRequest
	KeywordUpsertRequest *KeywordUpsertRequest
	MLSpamUpsertRequest *MLSpamUpsertRequest
	MentionSpamUpsertRequest *MentionSpamUpsertRequest
}

// DefaultKeywordListUpsertRequestAsCreateAutoModerationRuleRequest is a convenience function that returns DefaultKeywordListUpsertRequest wrapped in CreateAutoModerationRuleRequest
func DefaultKeywordListUpsertRequestAsCreateAutoModerationRuleRequest(v *DefaultKeywordListUpsertRequest) CreateAutoModerationRuleRequest {
	return CreateAutoModerationRuleRequest{
		DefaultKeywordListUpsertRequest: v,
	}
}

// KeywordUpsertRequestAsCreateAutoModerationRuleRequest is a convenience function that returns KeywordUpsertRequest wrapped in CreateAutoModerationRuleRequest
func KeywordUpsertRequestAsCreateAutoModerationRuleRequest(v *KeywordUpsertRequest) CreateAutoModerationRuleRequest {
	return CreateAutoModerationRuleRequest{
		KeywordUpsertRequest: v,
	}
}

// MLSpamUpsertRequestAsCreateAutoModerationRuleRequest is a convenience function that returns MLSpamUpsertRequest wrapped in CreateAutoModerationRuleRequest
func MLSpamUpsertRequestAsCreateAutoModerationRuleRequest(v *MLSpamUpsertRequest) CreateAutoModerationRuleRequest {
	return CreateAutoModerationRuleRequest{
		MLSpamUpsertRequest: v,
	}
}

// MentionSpamUpsertRequestAsCreateAutoModerationRuleRequest is a convenience function that returns MentionSpamUpsertRequest wrapped in CreateAutoModerationRuleRequest
func MentionSpamUpsertRequestAsCreateAutoModerationRuleRequest(v *MentionSpamUpsertRequest) CreateAutoModerationRuleRequest {
	return CreateAutoModerationRuleRequest{
		MentionSpamUpsertRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateAutoModerationRuleRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DefaultKeywordListUpsertRequest
	err = newStrictDecoder(data).Decode(&dst.DefaultKeywordListUpsertRequest)
	if err == nil {
		jsonDefaultKeywordListUpsertRequest, _ := json.Marshal(dst.DefaultKeywordListUpsertRequest)
		if string(jsonDefaultKeywordListUpsertRequest) == "{}" { // empty struct
			dst.DefaultKeywordListUpsertRequest = nil
		} else {
			if err = validator.Validate(dst.DefaultKeywordListUpsertRequest); err != nil {
				dst.DefaultKeywordListUpsertRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.DefaultKeywordListUpsertRequest = nil
	}

	// try to unmarshal data into KeywordUpsertRequest
	err = newStrictDecoder(data).Decode(&dst.KeywordUpsertRequest)
	if err == nil {
		jsonKeywordUpsertRequest, _ := json.Marshal(dst.KeywordUpsertRequest)
		if string(jsonKeywordUpsertRequest) == "{}" { // empty struct
			dst.KeywordUpsertRequest = nil
		} else {
			if err = validator.Validate(dst.KeywordUpsertRequest); err != nil {
				dst.KeywordUpsertRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.KeywordUpsertRequest = nil
	}

	// try to unmarshal data into MLSpamUpsertRequest
	err = newStrictDecoder(data).Decode(&dst.MLSpamUpsertRequest)
	if err == nil {
		jsonMLSpamUpsertRequest, _ := json.Marshal(dst.MLSpamUpsertRequest)
		if string(jsonMLSpamUpsertRequest) == "{}" { // empty struct
			dst.MLSpamUpsertRequest = nil
		} else {
			if err = validator.Validate(dst.MLSpamUpsertRequest); err != nil {
				dst.MLSpamUpsertRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.MLSpamUpsertRequest = nil
	}

	// try to unmarshal data into MentionSpamUpsertRequest
	err = newStrictDecoder(data).Decode(&dst.MentionSpamUpsertRequest)
	if err == nil {
		jsonMentionSpamUpsertRequest, _ := json.Marshal(dst.MentionSpamUpsertRequest)
		if string(jsonMentionSpamUpsertRequest) == "{}" { // empty struct
			dst.MentionSpamUpsertRequest = nil
		} else {
			if err = validator.Validate(dst.MentionSpamUpsertRequest); err != nil {
				dst.MentionSpamUpsertRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionSpamUpsertRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DefaultKeywordListUpsertRequest = nil
		dst.KeywordUpsertRequest = nil
		dst.MLSpamUpsertRequest = nil
		dst.MentionSpamUpsertRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateAutoModerationRuleRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateAutoModerationRuleRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateAutoModerationRuleRequest) MarshalJSON() ([]byte, error) {
	if src.DefaultKeywordListUpsertRequest != nil {
		return json.Marshal(&src.DefaultKeywordListUpsertRequest)
	}

	if src.KeywordUpsertRequest != nil {
		return json.Marshal(&src.KeywordUpsertRequest)
	}

	if src.MLSpamUpsertRequest != nil {
		return json.Marshal(&src.MLSpamUpsertRequest)
	}

	if src.MentionSpamUpsertRequest != nil {
		return json.Marshal(&src.MentionSpamUpsertRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateAutoModerationRuleRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DefaultKeywordListUpsertRequest != nil {
		return obj.DefaultKeywordListUpsertRequest
	}

	if obj.KeywordUpsertRequest != nil {
		return obj.KeywordUpsertRequest
	}

	if obj.MLSpamUpsertRequest != nil {
		return obj.MLSpamUpsertRequest
	}

	if obj.MentionSpamUpsertRequest != nil {
		return obj.MentionSpamUpsertRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateAutoModerationRuleRequest) GetActualInstanceValue() (interface{}) {
	if obj.DefaultKeywordListUpsertRequest != nil {
		return *obj.DefaultKeywordListUpsertRequest
	}

	if obj.KeywordUpsertRequest != nil {
		return *obj.KeywordUpsertRequest
	}

	if obj.MLSpamUpsertRequest != nil {
		return *obj.MLSpamUpsertRequest
	}

	if obj.MentionSpamUpsertRequest != nil {
		return *obj.MentionSpamUpsertRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateAutoModerationRuleRequest struct {
	value *CreateAutoModerationRuleRequest
	isSet bool
}

func (v NullableCreateAutoModerationRuleRequest) Get() *CreateAutoModerationRuleRequest {
	return v.value
}

func (v *NullableCreateAutoModerationRuleRequest) Set(val *CreateAutoModerationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutoModerationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutoModerationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutoModerationRuleRequest(val *CreateAutoModerationRuleRequest) *NullableCreateAutoModerationRuleRequest {
	return &NullableCreateAutoModerationRuleRequest{value: val, isSet: true}
}

func (v NullableCreateAutoModerationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutoModerationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


