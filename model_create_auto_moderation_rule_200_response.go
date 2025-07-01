/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// CreateAutoModerationRule200Response - struct for CreateAutoModerationRule200Response
type CreateAutoModerationRule200Response struct {
	DefaultKeywordRuleResponse *DefaultKeywordRuleResponse
	KeywordRuleResponse *KeywordRuleResponse
	MLSpamRuleResponse *MLSpamRuleResponse
	MentionSpamRuleResponse *MentionSpamRuleResponse
	SpamLinkRuleResponse *SpamLinkRuleResponse
}

// DefaultKeywordRuleResponseAsCreateAutoModerationRule200Response is a convenience function that returns DefaultKeywordRuleResponse wrapped in CreateAutoModerationRule200Response
func DefaultKeywordRuleResponseAsCreateAutoModerationRule200Response(v *DefaultKeywordRuleResponse) CreateAutoModerationRule200Response {
	return CreateAutoModerationRule200Response{
		DefaultKeywordRuleResponse: v,
	}
}

// KeywordRuleResponseAsCreateAutoModerationRule200Response is a convenience function that returns KeywordRuleResponse wrapped in CreateAutoModerationRule200Response
func KeywordRuleResponseAsCreateAutoModerationRule200Response(v *KeywordRuleResponse) CreateAutoModerationRule200Response {
	return CreateAutoModerationRule200Response{
		KeywordRuleResponse: v,
	}
}

// MLSpamRuleResponseAsCreateAutoModerationRule200Response is a convenience function that returns MLSpamRuleResponse wrapped in CreateAutoModerationRule200Response
func MLSpamRuleResponseAsCreateAutoModerationRule200Response(v *MLSpamRuleResponse) CreateAutoModerationRule200Response {
	return CreateAutoModerationRule200Response{
		MLSpamRuleResponse: v,
	}
}

// MentionSpamRuleResponseAsCreateAutoModerationRule200Response is a convenience function that returns MentionSpamRuleResponse wrapped in CreateAutoModerationRule200Response
func MentionSpamRuleResponseAsCreateAutoModerationRule200Response(v *MentionSpamRuleResponse) CreateAutoModerationRule200Response {
	return CreateAutoModerationRule200Response{
		MentionSpamRuleResponse: v,
	}
}

// SpamLinkRuleResponseAsCreateAutoModerationRule200Response is a convenience function that returns SpamLinkRuleResponse wrapped in CreateAutoModerationRule200Response
func SpamLinkRuleResponseAsCreateAutoModerationRule200Response(v *SpamLinkRuleResponse) CreateAutoModerationRule200Response {
	return CreateAutoModerationRule200Response{
		SpamLinkRuleResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateAutoModerationRule200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DefaultKeywordRuleResponse
	err = newStrictDecoder(data).Decode(&dst.DefaultKeywordRuleResponse)
	if err == nil {
		jsonDefaultKeywordRuleResponse, _ := json.Marshal(dst.DefaultKeywordRuleResponse)
		if string(jsonDefaultKeywordRuleResponse) == "{}" { // empty struct
			dst.DefaultKeywordRuleResponse = nil
		} else {
			if err = validator.Validate(dst.DefaultKeywordRuleResponse); err != nil {
				dst.DefaultKeywordRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.DefaultKeywordRuleResponse = nil
	}

	// try to unmarshal data into KeywordRuleResponse
	err = newStrictDecoder(data).Decode(&dst.KeywordRuleResponse)
	if err == nil {
		jsonKeywordRuleResponse, _ := json.Marshal(dst.KeywordRuleResponse)
		if string(jsonKeywordRuleResponse) == "{}" { // empty struct
			dst.KeywordRuleResponse = nil
		} else {
			if err = validator.Validate(dst.KeywordRuleResponse); err != nil {
				dst.KeywordRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.KeywordRuleResponse = nil
	}

	// try to unmarshal data into MLSpamRuleResponse
	err = newStrictDecoder(data).Decode(&dst.MLSpamRuleResponse)
	if err == nil {
		jsonMLSpamRuleResponse, _ := json.Marshal(dst.MLSpamRuleResponse)
		if string(jsonMLSpamRuleResponse) == "{}" { // empty struct
			dst.MLSpamRuleResponse = nil
		} else {
			if err = validator.Validate(dst.MLSpamRuleResponse); err != nil {
				dst.MLSpamRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MLSpamRuleResponse = nil
	}

	// try to unmarshal data into MentionSpamRuleResponse
	err = newStrictDecoder(data).Decode(&dst.MentionSpamRuleResponse)
	if err == nil {
		jsonMentionSpamRuleResponse, _ := json.Marshal(dst.MentionSpamRuleResponse)
		if string(jsonMentionSpamRuleResponse) == "{}" { // empty struct
			dst.MentionSpamRuleResponse = nil
		} else {
			if err = validator.Validate(dst.MentionSpamRuleResponse); err != nil {
				dst.MentionSpamRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionSpamRuleResponse = nil
	}

	// try to unmarshal data into SpamLinkRuleResponse
	err = newStrictDecoder(data).Decode(&dst.SpamLinkRuleResponse)
	if err == nil {
		jsonSpamLinkRuleResponse, _ := json.Marshal(dst.SpamLinkRuleResponse)
		if string(jsonSpamLinkRuleResponse) == "{}" { // empty struct
			dst.SpamLinkRuleResponse = nil
		} else {
			if err = validator.Validate(dst.SpamLinkRuleResponse); err != nil {
				dst.SpamLinkRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpamLinkRuleResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DefaultKeywordRuleResponse = nil
		dst.KeywordRuleResponse = nil
		dst.MLSpamRuleResponse = nil
		dst.MentionSpamRuleResponse = nil
		dst.SpamLinkRuleResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateAutoModerationRule200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateAutoModerationRule200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateAutoModerationRule200Response) MarshalJSON() ([]byte, error) {
	if src.DefaultKeywordRuleResponse != nil {
		return json.Marshal(&src.DefaultKeywordRuleResponse)
	}

	if src.KeywordRuleResponse != nil {
		return json.Marshal(&src.KeywordRuleResponse)
	}

	if src.MLSpamRuleResponse != nil {
		return json.Marshal(&src.MLSpamRuleResponse)
	}

	if src.MentionSpamRuleResponse != nil {
		return json.Marshal(&src.MentionSpamRuleResponse)
	}

	if src.SpamLinkRuleResponse != nil {
		return json.Marshal(&src.SpamLinkRuleResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateAutoModerationRule200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DefaultKeywordRuleResponse != nil {
		return obj.DefaultKeywordRuleResponse
	}

	if obj.KeywordRuleResponse != nil {
		return obj.KeywordRuleResponse
	}

	if obj.MLSpamRuleResponse != nil {
		return obj.MLSpamRuleResponse
	}

	if obj.MentionSpamRuleResponse != nil {
		return obj.MentionSpamRuleResponse
	}

	if obj.SpamLinkRuleResponse != nil {
		return obj.SpamLinkRuleResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateAutoModerationRule200Response) GetActualInstanceValue() (interface{}) {
	if obj.DefaultKeywordRuleResponse != nil {
		return *obj.DefaultKeywordRuleResponse
	}

	if obj.KeywordRuleResponse != nil {
		return *obj.KeywordRuleResponse
	}

	if obj.MLSpamRuleResponse != nil {
		return *obj.MLSpamRuleResponse
	}

	if obj.MentionSpamRuleResponse != nil {
		return *obj.MentionSpamRuleResponse
	}

	if obj.SpamLinkRuleResponse != nil {
		return *obj.SpamLinkRuleResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateAutoModerationRule200Response struct {
	value *CreateAutoModerationRule200Response
	isSet bool
}

func (v NullableCreateAutoModerationRule200Response) Get() *CreateAutoModerationRule200Response {
	return v.value
}

func (v *NullableCreateAutoModerationRule200Response) Set(val *CreateAutoModerationRule200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutoModerationRule200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutoModerationRule200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutoModerationRule200Response(val *CreateAutoModerationRule200Response) *NullableCreateAutoModerationRule200Response {
	return &NullableCreateAutoModerationRule200Response{value: val, isSet: true}
}

func (v NullableCreateAutoModerationRule200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutoModerationRule200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


