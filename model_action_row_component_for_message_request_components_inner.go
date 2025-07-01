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
	"fmt"
	"gopkg.in/validator.v2"
)

// ActionRowComponentForMessageRequestComponentsInner - struct for ActionRowComponentForMessageRequestComponentsInner
type ActionRowComponentForMessageRequestComponentsInner struct {
	ButtonComponentForMessageRequest *ButtonComponentForMessageRequest
	ChannelSelectComponentForMessageRequest *ChannelSelectComponentForMessageRequest
	MentionableSelectComponentForMessageRequest *MentionableSelectComponentForMessageRequest
	RoleSelectComponentForMessageRequest *RoleSelectComponentForMessageRequest
	StringSelectComponentForMessageRequest *StringSelectComponentForMessageRequest
	UserSelectComponentForMessageRequest *UserSelectComponentForMessageRequest
}

// ButtonComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns ButtonComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func ButtonComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *ButtonComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		ButtonComponentForMessageRequest: v,
	}
}

// ChannelSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns ChannelSelectComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func ChannelSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *ChannelSelectComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		ChannelSelectComponentForMessageRequest: v,
	}
}

// MentionableSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns MentionableSelectComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func MentionableSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *MentionableSelectComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		MentionableSelectComponentForMessageRequest: v,
	}
}

// RoleSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns RoleSelectComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func RoleSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *RoleSelectComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		RoleSelectComponentForMessageRequest: v,
	}
}

// StringSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns StringSelectComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func StringSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *StringSelectComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		StringSelectComponentForMessageRequest: v,
	}
}

// UserSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner is a convenience function that returns UserSelectComponentForMessageRequest wrapped in ActionRowComponentForMessageRequestComponentsInner
func UserSelectComponentForMessageRequestAsActionRowComponentForMessageRequestComponentsInner(v *UserSelectComponentForMessageRequest) ActionRowComponentForMessageRequestComponentsInner {
	return ActionRowComponentForMessageRequestComponentsInner{
		UserSelectComponentForMessageRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ActionRowComponentForMessageRequestComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ButtonComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.ButtonComponentForMessageRequest)
	if err == nil {
		jsonButtonComponentForMessageRequest, _ := json.Marshal(dst.ButtonComponentForMessageRequest)
		if string(jsonButtonComponentForMessageRequest) == "{}" { // empty struct
			dst.ButtonComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.ButtonComponentForMessageRequest); err != nil {
				dst.ButtonComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ButtonComponentForMessageRequest = nil
	}

	// try to unmarshal data into ChannelSelectComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.ChannelSelectComponentForMessageRequest)
	if err == nil {
		jsonChannelSelectComponentForMessageRequest, _ := json.Marshal(dst.ChannelSelectComponentForMessageRequest)
		if string(jsonChannelSelectComponentForMessageRequest) == "{}" { // empty struct
			dst.ChannelSelectComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.ChannelSelectComponentForMessageRequest); err != nil {
				dst.ChannelSelectComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChannelSelectComponentForMessageRequest = nil
	}

	// try to unmarshal data into MentionableSelectComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.MentionableSelectComponentForMessageRequest)
	if err == nil {
		jsonMentionableSelectComponentForMessageRequest, _ := json.Marshal(dst.MentionableSelectComponentForMessageRequest)
		if string(jsonMentionableSelectComponentForMessageRequest) == "{}" { // empty struct
			dst.MentionableSelectComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.MentionableSelectComponentForMessageRequest); err != nil {
				dst.MentionableSelectComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionableSelectComponentForMessageRequest = nil
	}

	// try to unmarshal data into RoleSelectComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.RoleSelectComponentForMessageRequest)
	if err == nil {
		jsonRoleSelectComponentForMessageRequest, _ := json.Marshal(dst.RoleSelectComponentForMessageRequest)
		if string(jsonRoleSelectComponentForMessageRequest) == "{}" { // empty struct
			dst.RoleSelectComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.RoleSelectComponentForMessageRequest); err != nil {
				dst.RoleSelectComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelectComponentForMessageRequest = nil
	}

	// try to unmarshal data into StringSelectComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.StringSelectComponentForMessageRequest)
	if err == nil {
		jsonStringSelectComponentForMessageRequest, _ := json.Marshal(dst.StringSelectComponentForMessageRequest)
		if string(jsonStringSelectComponentForMessageRequest) == "{}" { // empty struct
			dst.StringSelectComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.StringSelectComponentForMessageRequest); err != nil {
				dst.StringSelectComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.StringSelectComponentForMessageRequest = nil
	}

	// try to unmarshal data into UserSelectComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.UserSelectComponentForMessageRequest)
	if err == nil {
		jsonUserSelectComponentForMessageRequest, _ := json.Marshal(dst.UserSelectComponentForMessageRequest)
		if string(jsonUserSelectComponentForMessageRequest) == "{}" { // empty struct
			dst.UserSelectComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.UserSelectComponentForMessageRequest); err != nil {
				dst.UserSelectComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelectComponentForMessageRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ButtonComponentForMessageRequest = nil
		dst.ChannelSelectComponentForMessageRequest = nil
		dst.MentionableSelectComponentForMessageRequest = nil
		dst.RoleSelectComponentForMessageRequest = nil
		dst.StringSelectComponentForMessageRequest = nil
		dst.UserSelectComponentForMessageRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ActionRowComponentForMessageRequestComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ActionRowComponentForMessageRequestComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActionRowComponentForMessageRequestComponentsInner) MarshalJSON() ([]byte, error) {
	if src.ButtonComponentForMessageRequest != nil {
		return json.Marshal(&src.ButtonComponentForMessageRequest)
	}

	if src.ChannelSelectComponentForMessageRequest != nil {
		return json.Marshal(&src.ChannelSelectComponentForMessageRequest)
	}

	if src.MentionableSelectComponentForMessageRequest != nil {
		return json.Marshal(&src.MentionableSelectComponentForMessageRequest)
	}

	if src.RoleSelectComponentForMessageRequest != nil {
		return json.Marshal(&src.RoleSelectComponentForMessageRequest)
	}

	if src.StringSelectComponentForMessageRequest != nil {
		return json.Marshal(&src.StringSelectComponentForMessageRequest)
	}

	if src.UserSelectComponentForMessageRequest != nil {
		return json.Marshal(&src.UserSelectComponentForMessageRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActionRowComponentForMessageRequestComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ButtonComponentForMessageRequest != nil {
		return obj.ButtonComponentForMessageRequest
	}

	if obj.ChannelSelectComponentForMessageRequest != nil {
		return obj.ChannelSelectComponentForMessageRequest
	}

	if obj.MentionableSelectComponentForMessageRequest != nil {
		return obj.MentionableSelectComponentForMessageRequest
	}

	if obj.RoleSelectComponentForMessageRequest != nil {
		return obj.RoleSelectComponentForMessageRequest
	}

	if obj.StringSelectComponentForMessageRequest != nil {
		return obj.StringSelectComponentForMessageRequest
	}

	if obj.UserSelectComponentForMessageRequest != nil {
		return obj.UserSelectComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ActionRowComponentForMessageRequestComponentsInner) GetActualInstanceValue() (interface{}) {
	if obj.ButtonComponentForMessageRequest != nil {
		return *obj.ButtonComponentForMessageRequest
	}

	if obj.ChannelSelectComponentForMessageRequest != nil {
		return *obj.ChannelSelectComponentForMessageRequest
	}

	if obj.MentionableSelectComponentForMessageRequest != nil {
		return *obj.MentionableSelectComponentForMessageRequest
	}

	if obj.RoleSelectComponentForMessageRequest != nil {
		return *obj.RoleSelectComponentForMessageRequest
	}

	if obj.StringSelectComponentForMessageRequest != nil {
		return *obj.StringSelectComponentForMessageRequest
	}

	if obj.UserSelectComponentForMessageRequest != nil {
		return *obj.UserSelectComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

type NullableActionRowComponentForMessageRequestComponentsInner struct {
	value *ActionRowComponentForMessageRequestComponentsInner
	isSet bool
}

func (v NullableActionRowComponentForMessageRequestComponentsInner) Get() *ActionRowComponentForMessageRequestComponentsInner {
	return v.value
}

func (v *NullableActionRowComponentForMessageRequestComponentsInner) Set(val *ActionRowComponentForMessageRequestComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRowComponentForMessageRequestComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRowComponentForMessageRequestComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRowComponentForMessageRequestComponentsInner(val *ActionRowComponentForMessageRequestComponentsInner) *NullableActionRowComponentForMessageRequestComponentsInner {
	return &NullableActionRowComponentForMessageRequestComponentsInner{value: val, isSet: true}
}

func (v NullableActionRowComponentForMessageRequestComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRowComponentForMessageRequestComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


