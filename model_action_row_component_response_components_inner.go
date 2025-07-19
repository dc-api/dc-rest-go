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
	"fmt"
	"gopkg.in/validator.v2"
)

// ActionRowComponentResponseComponentsInner - struct for ActionRowComponentResponseComponentsInner
type ActionRowComponentResponseComponentsInner struct {
	ButtonComponentResponse *ButtonComponentResponse
	ChannelSelectComponentResponse *ChannelSelectComponentResponse
	MentionableSelectComponentResponse *MentionableSelectComponentResponse
	RoleSelectComponentResponse *RoleSelectComponentResponse
	StringSelectComponentResponse *StringSelectComponentResponse
	TextInputComponentResponse *TextInputComponentResponse
	UserSelectComponentResponse *UserSelectComponentResponse
}

// ButtonComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns ButtonComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func ButtonComponentResponseAsActionRowComponentResponseComponentsInner(v *ButtonComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		ButtonComponentResponse: v,
	}
}

// ChannelSelectComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns ChannelSelectComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func ChannelSelectComponentResponseAsActionRowComponentResponseComponentsInner(v *ChannelSelectComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		ChannelSelectComponentResponse: v,
	}
}

// MentionableSelectComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns MentionableSelectComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func MentionableSelectComponentResponseAsActionRowComponentResponseComponentsInner(v *MentionableSelectComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		MentionableSelectComponentResponse: v,
	}
}

// RoleSelectComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns RoleSelectComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func RoleSelectComponentResponseAsActionRowComponentResponseComponentsInner(v *RoleSelectComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		RoleSelectComponentResponse: v,
	}
}

// StringSelectComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns StringSelectComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func StringSelectComponentResponseAsActionRowComponentResponseComponentsInner(v *StringSelectComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		StringSelectComponentResponse: v,
	}
}

// TextInputComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns TextInputComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func TextInputComponentResponseAsActionRowComponentResponseComponentsInner(v *TextInputComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		TextInputComponentResponse: v,
	}
}

// UserSelectComponentResponseAsActionRowComponentResponseComponentsInner is a convenience function that returns UserSelectComponentResponse wrapped in ActionRowComponentResponseComponentsInner
func UserSelectComponentResponseAsActionRowComponentResponseComponentsInner(v *UserSelectComponentResponse) ActionRowComponentResponseComponentsInner {
	return ActionRowComponentResponseComponentsInner{
		UserSelectComponentResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ActionRowComponentResponseComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ButtonComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ButtonComponentResponse)
	if err == nil {
		jsonButtonComponentResponse, _ := json.Marshal(dst.ButtonComponentResponse)
		if string(jsonButtonComponentResponse) == "{}" { // empty struct
			dst.ButtonComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ButtonComponentResponse); err != nil {
				dst.ButtonComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ButtonComponentResponse = nil
	}

	// try to unmarshal data into ChannelSelectComponentResponse
	err = newStrictDecoder(data).Decode(&dst.ChannelSelectComponentResponse)
	if err == nil {
		jsonChannelSelectComponentResponse, _ := json.Marshal(dst.ChannelSelectComponentResponse)
		if string(jsonChannelSelectComponentResponse) == "{}" { // empty struct
			dst.ChannelSelectComponentResponse = nil
		} else {
			if err = validator.Validate(dst.ChannelSelectComponentResponse); err != nil {
				dst.ChannelSelectComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChannelSelectComponentResponse = nil
	}

	// try to unmarshal data into MentionableSelectComponentResponse
	err = newStrictDecoder(data).Decode(&dst.MentionableSelectComponentResponse)
	if err == nil {
		jsonMentionableSelectComponentResponse, _ := json.Marshal(dst.MentionableSelectComponentResponse)
		if string(jsonMentionableSelectComponentResponse) == "{}" { // empty struct
			dst.MentionableSelectComponentResponse = nil
		} else {
			if err = validator.Validate(dst.MentionableSelectComponentResponse); err != nil {
				dst.MentionableSelectComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionableSelectComponentResponse = nil
	}

	// try to unmarshal data into RoleSelectComponentResponse
	err = newStrictDecoder(data).Decode(&dst.RoleSelectComponentResponse)
	if err == nil {
		jsonRoleSelectComponentResponse, _ := json.Marshal(dst.RoleSelectComponentResponse)
		if string(jsonRoleSelectComponentResponse) == "{}" { // empty struct
			dst.RoleSelectComponentResponse = nil
		} else {
			if err = validator.Validate(dst.RoleSelectComponentResponse); err != nil {
				dst.RoleSelectComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelectComponentResponse = nil
	}

	// try to unmarshal data into StringSelectComponentResponse
	err = newStrictDecoder(data).Decode(&dst.StringSelectComponentResponse)
	if err == nil {
		jsonStringSelectComponentResponse, _ := json.Marshal(dst.StringSelectComponentResponse)
		if string(jsonStringSelectComponentResponse) == "{}" { // empty struct
			dst.StringSelectComponentResponse = nil
		} else {
			if err = validator.Validate(dst.StringSelectComponentResponse); err != nil {
				dst.StringSelectComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StringSelectComponentResponse = nil
	}

	// try to unmarshal data into TextInputComponentResponse
	err = newStrictDecoder(data).Decode(&dst.TextInputComponentResponse)
	if err == nil {
		jsonTextInputComponentResponse, _ := json.Marshal(dst.TextInputComponentResponse)
		if string(jsonTextInputComponentResponse) == "{}" { // empty struct
			dst.TextInputComponentResponse = nil
		} else {
			if err = validator.Validate(dst.TextInputComponentResponse); err != nil {
				dst.TextInputComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.TextInputComponentResponse = nil
	}

	// try to unmarshal data into UserSelectComponentResponse
	err = newStrictDecoder(data).Decode(&dst.UserSelectComponentResponse)
	if err == nil {
		jsonUserSelectComponentResponse, _ := json.Marshal(dst.UserSelectComponentResponse)
		if string(jsonUserSelectComponentResponse) == "{}" { // empty struct
			dst.UserSelectComponentResponse = nil
		} else {
			if err = validator.Validate(dst.UserSelectComponentResponse); err != nil {
				dst.UserSelectComponentResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelectComponentResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ButtonComponentResponse = nil
		dst.ChannelSelectComponentResponse = nil
		dst.MentionableSelectComponentResponse = nil
		dst.RoleSelectComponentResponse = nil
		dst.StringSelectComponentResponse = nil
		dst.TextInputComponentResponse = nil
		dst.UserSelectComponentResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ActionRowComponentResponseComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ActionRowComponentResponseComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActionRowComponentResponseComponentsInner) MarshalJSON() ([]byte, error) {
	if src.ButtonComponentResponse != nil {
		return json.Marshal(&src.ButtonComponentResponse)
	}

	if src.ChannelSelectComponentResponse != nil {
		return json.Marshal(&src.ChannelSelectComponentResponse)
	}

	if src.MentionableSelectComponentResponse != nil {
		return json.Marshal(&src.MentionableSelectComponentResponse)
	}

	if src.RoleSelectComponentResponse != nil {
		return json.Marshal(&src.RoleSelectComponentResponse)
	}

	if src.StringSelectComponentResponse != nil {
		return json.Marshal(&src.StringSelectComponentResponse)
	}

	if src.TextInputComponentResponse != nil {
		return json.Marshal(&src.TextInputComponentResponse)
	}

	if src.UserSelectComponentResponse != nil {
		return json.Marshal(&src.UserSelectComponentResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActionRowComponentResponseComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ButtonComponentResponse != nil {
		return obj.ButtonComponentResponse
	}

	if obj.ChannelSelectComponentResponse != nil {
		return obj.ChannelSelectComponentResponse
	}

	if obj.MentionableSelectComponentResponse != nil {
		return obj.MentionableSelectComponentResponse
	}

	if obj.RoleSelectComponentResponse != nil {
		return obj.RoleSelectComponentResponse
	}

	if obj.StringSelectComponentResponse != nil {
		return obj.StringSelectComponentResponse
	}

	if obj.TextInputComponentResponse != nil {
		return obj.TextInputComponentResponse
	}

	if obj.UserSelectComponentResponse != nil {
		return obj.UserSelectComponentResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ActionRowComponentResponseComponentsInner) GetActualInstanceValue() (interface{}) {
	if obj.ButtonComponentResponse != nil {
		return *obj.ButtonComponentResponse
	}

	if obj.ChannelSelectComponentResponse != nil {
		return *obj.ChannelSelectComponentResponse
	}

	if obj.MentionableSelectComponentResponse != nil {
		return *obj.MentionableSelectComponentResponse
	}

	if obj.RoleSelectComponentResponse != nil {
		return *obj.RoleSelectComponentResponse
	}

	if obj.StringSelectComponentResponse != nil {
		return *obj.StringSelectComponentResponse
	}

	if obj.TextInputComponentResponse != nil {
		return *obj.TextInputComponentResponse
	}

	if obj.UserSelectComponentResponse != nil {
		return *obj.UserSelectComponentResponse
	}

	// all schemas are nil
	return nil
}

type NullableActionRowComponentResponseComponentsInner struct {
	value *ActionRowComponentResponseComponentsInner
	isSet bool
}

func (v NullableActionRowComponentResponseComponentsInner) Get() *ActionRowComponentResponseComponentsInner {
	return v.value
}

func (v *NullableActionRowComponentResponseComponentsInner) Set(val *ActionRowComponentResponseComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRowComponentResponseComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRowComponentResponseComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRowComponentResponseComponentsInner(val *ActionRowComponentResponseComponentsInner) *NullableActionRowComponentResponseComponentsInner {
	return &NullableActionRowComponentResponseComponentsInner{value: val, isSet: true}
}

func (v NullableActionRowComponentResponseComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRowComponentResponseComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


