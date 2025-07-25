/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// MentionableSelectComponentForMessageRequestDefaultValuesInner - struct for MentionableSelectComponentForMessageRequestDefaultValuesInner
type MentionableSelectComponentForMessageRequestDefaultValuesInner struct {
	RoleSelectDefaultValue *RoleSelectDefaultValue
	UserSelectDefaultValue *UserSelectDefaultValue
}

// RoleSelectDefaultValueAsMentionableSelectComponentForMessageRequestDefaultValuesInner is a convenience function that returns RoleSelectDefaultValue wrapped in MentionableSelectComponentForMessageRequestDefaultValuesInner
func RoleSelectDefaultValueAsMentionableSelectComponentForMessageRequestDefaultValuesInner(v *RoleSelectDefaultValue) MentionableSelectComponentForMessageRequestDefaultValuesInner {
	return MentionableSelectComponentForMessageRequestDefaultValuesInner{
		RoleSelectDefaultValue: v,
	}
}

// UserSelectDefaultValueAsMentionableSelectComponentForMessageRequestDefaultValuesInner is a convenience function that returns UserSelectDefaultValue wrapped in MentionableSelectComponentForMessageRequestDefaultValuesInner
func UserSelectDefaultValueAsMentionableSelectComponentForMessageRequestDefaultValuesInner(v *UserSelectDefaultValue) MentionableSelectComponentForMessageRequestDefaultValuesInner {
	return MentionableSelectComponentForMessageRequestDefaultValuesInner{
		UserSelectDefaultValue: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MentionableSelectComponentForMessageRequestDefaultValuesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RoleSelectDefaultValue
	err = newStrictDecoder(data).Decode(&dst.RoleSelectDefaultValue)
	if err == nil {
		jsonRoleSelectDefaultValue, _ := json.Marshal(dst.RoleSelectDefaultValue)
		if string(jsonRoleSelectDefaultValue) == "{}" { // empty struct
			dst.RoleSelectDefaultValue = nil
		} else {
			if err = validator.Validate(dst.RoleSelectDefaultValue); err != nil {
				dst.RoleSelectDefaultValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelectDefaultValue = nil
	}

	// try to unmarshal data into UserSelectDefaultValue
	err = newStrictDecoder(data).Decode(&dst.UserSelectDefaultValue)
	if err == nil {
		jsonUserSelectDefaultValue, _ := json.Marshal(dst.UserSelectDefaultValue)
		if string(jsonUserSelectDefaultValue) == "{}" { // empty struct
			dst.UserSelectDefaultValue = nil
		} else {
			if err = validator.Validate(dst.UserSelectDefaultValue); err != nil {
				dst.UserSelectDefaultValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelectDefaultValue = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RoleSelectDefaultValue = nil
		dst.UserSelectDefaultValue = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MentionableSelectComponentForMessageRequestDefaultValuesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MentionableSelectComponentForMessageRequestDefaultValuesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MentionableSelectComponentForMessageRequestDefaultValuesInner) MarshalJSON() ([]byte, error) {
	if src.RoleSelectDefaultValue != nil {
		return json.Marshal(&src.RoleSelectDefaultValue)
	}

	if src.UserSelectDefaultValue != nil {
		return json.Marshal(&src.UserSelectDefaultValue)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MentionableSelectComponentForMessageRequestDefaultValuesInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RoleSelectDefaultValue != nil {
		return obj.RoleSelectDefaultValue
	}

	if obj.UserSelectDefaultValue != nil {
		return obj.UserSelectDefaultValue
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MentionableSelectComponentForMessageRequestDefaultValuesInner) GetActualInstanceValue() (interface{}) {
	if obj.RoleSelectDefaultValue != nil {
		return *obj.RoleSelectDefaultValue
	}

	if obj.UserSelectDefaultValue != nil {
		return *obj.UserSelectDefaultValue
	}

	// all schemas are nil
	return nil
}

type NullableMentionableSelectComponentForMessageRequestDefaultValuesInner struct {
	value *MentionableSelectComponentForMessageRequestDefaultValuesInner
	isSet bool
}

func (v NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) Get() *MentionableSelectComponentForMessageRequestDefaultValuesInner {
	return v.value
}

func (v *NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) Set(val *MentionableSelectComponentForMessageRequestDefaultValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionableSelectComponentForMessageRequestDefaultValuesInner(val *MentionableSelectComponentForMessageRequestDefaultValuesInner) *NullableMentionableSelectComponentForMessageRequestDefaultValuesInner {
	return &NullableMentionableSelectComponentForMessageRequestDefaultValuesInner{value: val, isSet: true}
}

func (v NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionableSelectComponentForMessageRequestDefaultValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


