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

// DefaultKeywordRuleResponseActionsInner - struct for DefaultKeywordRuleResponseActionsInner
type DefaultKeywordRuleResponseActionsInner struct {
	BlockMessageActionResponse *BlockMessageActionResponse
	FlagToChannelActionResponse *FlagToChannelActionResponse
	QuarantineUserActionResponse *QuarantineUserActionResponse
	UserCommunicationDisabledActionResponse *UserCommunicationDisabledActionResponse
}

// BlockMessageActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns BlockMessageActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func BlockMessageActionResponseAsDefaultKeywordRuleResponseActionsInner(v *BlockMessageActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		BlockMessageActionResponse: v,
	}
}

// FlagToChannelActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns FlagToChannelActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func FlagToChannelActionResponseAsDefaultKeywordRuleResponseActionsInner(v *FlagToChannelActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		FlagToChannelActionResponse: v,
	}
}

// QuarantineUserActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns QuarantineUserActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func QuarantineUserActionResponseAsDefaultKeywordRuleResponseActionsInner(v *QuarantineUserActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		QuarantineUserActionResponse: v,
	}
}

// UserCommunicationDisabledActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns UserCommunicationDisabledActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func UserCommunicationDisabledActionResponseAsDefaultKeywordRuleResponseActionsInner(v *UserCommunicationDisabledActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		UserCommunicationDisabledActionResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultKeywordRuleResponseActionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlockMessageActionResponse
	err = newStrictDecoder(data).Decode(&dst.BlockMessageActionResponse)
	if err == nil {
		jsonBlockMessageActionResponse, _ := json.Marshal(dst.BlockMessageActionResponse)
		if string(jsonBlockMessageActionResponse) == "{}" { // empty struct
			dst.BlockMessageActionResponse = nil
		} else {
			if err = validator.Validate(dst.BlockMessageActionResponse); err != nil {
				dst.BlockMessageActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlockMessageActionResponse = nil
	}

	// try to unmarshal data into FlagToChannelActionResponse
	err = newStrictDecoder(data).Decode(&dst.FlagToChannelActionResponse)
	if err == nil {
		jsonFlagToChannelActionResponse, _ := json.Marshal(dst.FlagToChannelActionResponse)
		if string(jsonFlagToChannelActionResponse) == "{}" { // empty struct
			dst.FlagToChannelActionResponse = nil
		} else {
			if err = validator.Validate(dst.FlagToChannelActionResponse); err != nil {
				dst.FlagToChannelActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FlagToChannelActionResponse = nil
	}

	// try to unmarshal data into QuarantineUserActionResponse
	err = newStrictDecoder(data).Decode(&dst.QuarantineUserActionResponse)
	if err == nil {
		jsonQuarantineUserActionResponse, _ := json.Marshal(dst.QuarantineUserActionResponse)
		if string(jsonQuarantineUserActionResponse) == "{}" { // empty struct
			dst.QuarantineUserActionResponse = nil
		} else {
			if err = validator.Validate(dst.QuarantineUserActionResponse); err != nil {
				dst.QuarantineUserActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuarantineUserActionResponse = nil
	}

	// try to unmarshal data into UserCommunicationDisabledActionResponse
	err = newStrictDecoder(data).Decode(&dst.UserCommunicationDisabledActionResponse)
	if err == nil {
		jsonUserCommunicationDisabledActionResponse, _ := json.Marshal(dst.UserCommunicationDisabledActionResponse)
		if string(jsonUserCommunicationDisabledActionResponse) == "{}" { // empty struct
			dst.UserCommunicationDisabledActionResponse = nil
		} else {
			if err = validator.Validate(dst.UserCommunicationDisabledActionResponse); err != nil {
				dst.UserCommunicationDisabledActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserCommunicationDisabledActionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlockMessageActionResponse = nil
		dst.FlagToChannelActionResponse = nil
		dst.QuarantineUserActionResponse = nil
		dst.UserCommunicationDisabledActionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DefaultKeywordRuleResponseActionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DefaultKeywordRuleResponseActionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultKeywordRuleResponseActionsInner) MarshalJSON() ([]byte, error) {
	if src.BlockMessageActionResponse != nil {
		return json.Marshal(&src.BlockMessageActionResponse)
	}

	if src.FlagToChannelActionResponse != nil {
		return json.Marshal(&src.FlagToChannelActionResponse)
	}

	if src.QuarantineUserActionResponse != nil {
		return json.Marshal(&src.QuarantineUserActionResponse)
	}

	if src.UserCommunicationDisabledActionResponse != nil {
		return json.Marshal(&src.UserCommunicationDisabledActionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultKeywordRuleResponseActionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlockMessageActionResponse != nil {
		return obj.BlockMessageActionResponse
	}

	if obj.FlagToChannelActionResponse != nil {
		return obj.FlagToChannelActionResponse
	}

	if obj.QuarantineUserActionResponse != nil {
		return obj.QuarantineUserActionResponse
	}

	if obj.UserCommunicationDisabledActionResponse != nil {
		return obj.UserCommunicationDisabledActionResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj DefaultKeywordRuleResponseActionsInner) GetActualInstanceValue() (interface{}) {
	if obj.BlockMessageActionResponse != nil {
		return *obj.BlockMessageActionResponse
	}

	if obj.FlagToChannelActionResponse != nil {
		return *obj.FlagToChannelActionResponse
	}

	if obj.QuarantineUserActionResponse != nil {
		return *obj.QuarantineUserActionResponse
	}

	if obj.UserCommunicationDisabledActionResponse != nil {
		return *obj.UserCommunicationDisabledActionResponse
	}

	// all schemas are nil
	return nil
}

type NullableDefaultKeywordRuleResponseActionsInner struct {
	value *DefaultKeywordRuleResponseActionsInner
	isSet bool
}

func (v NullableDefaultKeywordRuleResponseActionsInner) Get() *DefaultKeywordRuleResponseActionsInner {
	return v.value
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) Set(val *DefaultKeywordRuleResponseActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordRuleResponseActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordRuleResponseActionsInner(val *DefaultKeywordRuleResponseActionsInner) *NullableDefaultKeywordRuleResponseActionsInner {
	return &NullableDefaultKeywordRuleResponseActionsInner{value: val, isSet: true}
}

func (v NullableDefaultKeywordRuleResponseActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


