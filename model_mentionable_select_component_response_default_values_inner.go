/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// MentionableSelectComponentResponseDefaultValuesInner - struct for MentionableSelectComponentResponseDefaultValuesInner
type MentionableSelectComponentResponseDefaultValuesInner struct {
	RoleSelectDefaultValueResponse *RoleSelectDefaultValueResponse
	UserSelectDefaultValueResponse *UserSelectDefaultValueResponse
}

// RoleSelectDefaultValueResponseAsMentionableSelectComponentResponseDefaultValuesInner is a convenience function that returns RoleSelectDefaultValueResponse wrapped in MentionableSelectComponentResponseDefaultValuesInner
func RoleSelectDefaultValueResponseAsMentionableSelectComponentResponseDefaultValuesInner(v *RoleSelectDefaultValueResponse) MentionableSelectComponentResponseDefaultValuesInner {
	return MentionableSelectComponentResponseDefaultValuesInner{
		RoleSelectDefaultValueResponse: v,
	}
}

// UserSelectDefaultValueResponseAsMentionableSelectComponentResponseDefaultValuesInner is a convenience function that returns UserSelectDefaultValueResponse wrapped in MentionableSelectComponentResponseDefaultValuesInner
func UserSelectDefaultValueResponseAsMentionableSelectComponentResponseDefaultValuesInner(v *UserSelectDefaultValueResponse) MentionableSelectComponentResponseDefaultValuesInner {
	return MentionableSelectComponentResponseDefaultValuesInner{
		UserSelectDefaultValueResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MentionableSelectComponentResponseDefaultValuesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RoleSelectDefaultValueResponse
	err = newStrictDecoder(data).Decode(&dst.RoleSelectDefaultValueResponse)
	if err == nil {
		jsonRoleSelectDefaultValueResponse, _ := json.Marshal(dst.RoleSelectDefaultValueResponse)
		if string(jsonRoleSelectDefaultValueResponse) == "{}" { // empty struct
			dst.RoleSelectDefaultValueResponse = nil
		} else {
			if err = validator.Validate(dst.RoleSelectDefaultValueResponse); err != nil {
				dst.RoleSelectDefaultValueResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelectDefaultValueResponse = nil
	}

	// try to unmarshal data into UserSelectDefaultValueResponse
	err = newStrictDecoder(data).Decode(&dst.UserSelectDefaultValueResponse)
	if err == nil {
		jsonUserSelectDefaultValueResponse, _ := json.Marshal(dst.UserSelectDefaultValueResponse)
		if string(jsonUserSelectDefaultValueResponse) == "{}" { // empty struct
			dst.UserSelectDefaultValueResponse = nil
		} else {
			if err = validator.Validate(dst.UserSelectDefaultValueResponse); err != nil {
				dst.UserSelectDefaultValueResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelectDefaultValueResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RoleSelectDefaultValueResponse = nil
		dst.UserSelectDefaultValueResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MentionableSelectComponentResponseDefaultValuesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MentionableSelectComponentResponseDefaultValuesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MentionableSelectComponentResponseDefaultValuesInner) MarshalJSON() ([]byte, error) {
	if src.RoleSelectDefaultValueResponse != nil {
		return json.Marshal(&src.RoleSelectDefaultValueResponse)
	}

	if src.UserSelectDefaultValueResponse != nil {
		return json.Marshal(&src.UserSelectDefaultValueResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MentionableSelectComponentResponseDefaultValuesInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RoleSelectDefaultValueResponse != nil {
		return obj.RoleSelectDefaultValueResponse
	}

	if obj.UserSelectDefaultValueResponse != nil {
		return obj.UserSelectDefaultValueResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MentionableSelectComponentResponseDefaultValuesInner) GetActualInstanceValue() (interface{}) {
	if obj.RoleSelectDefaultValueResponse != nil {
		return *obj.RoleSelectDefaultValueResponse
	}

	if obj.UserSelectDefaultValueResponse != nil {
		return *obj.UserSelectDefaultValueResponse
	}

	// all schemas are nil
	return nil
}

type NullableMentionableSelectComponentResponseDefaultValuesInner struct {
	value *MentionableSelectComponentResponseDefaultValuesInner
	isSet bool
}

func (v NullableMentionableSelectComponentResponseDefaultValuesInner) Get() *MentionableSelectComponentResponseDefaultValuesInner {
	return v.value
}

func (v *NullableMentionableSelectComponentResponseDefaultValuesInner) Set(val *MentionableSelectComponentResponseDefaultValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionableSelectComponentResponseDefaultValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionableSelectComponentResponseDefaultValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionableSelectComponentResponseDefaultValuesInner(val *MentionableSelectComponentResponseDefaultValuesInner) *NullableMentionableSelectComponentResponseDefaultValuesInner {
	return &NullableMentionableSelectComponentResponseDefaultValuesInner{value: val, isSet: true}
}

func (v NullableMentionableSelectComponentResponseDefaultValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionableSelectComponentResponseDefaultValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


