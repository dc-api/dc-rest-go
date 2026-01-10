/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// LabelComponentForModalRequestComponent - struct for LabelComponentForModalRequestComponent
type LabelComponentForModalRequestComponent struct {
	ChannelSelectComponentForModalRequest *ChannelSelectComponentForModalRequest
	FileUploadComponentForModalRequest *FileUploadComponentForModalRequest
	MentionableSelectComponentForModalRequest *MentionableSelectComponentForModalRequest
	RoleSelectComponentForModalRequest *RoleSelectComponentForModalRequest
	StringSelectComponentForModalRequest *StringSelectComponentForModalRequest
	TextInputComponentForModalRequest *TextInputComponentForModalRequest
	UserSelectComponentForModalRequest *UserSelectComponentForModalRequest
}

// ChannelSelectComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns ChannelSelectComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func ChannelSelectComponentForModalRequestAsLabelComponentForModalRequestComponent(v *ChannelSelectComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		ChannelSelectComponentForModalRequest: v,
	}
}

// FileUploadComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns FileUploadComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func FileUploadComponentForModalRequestAsLabelComponentForModalRequestComponent(v *FileUploadComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		FileUploadComponentForModalRequest: v,
	}
}

// MentionableSelectComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns MentionableSelectComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func MentionableSelectComponentForModalRequestAsLabelComponentForModalRequestComponent(v *MentionableSelectComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		MentionableSelectComponentForModalRequest: v,
	}
}

// RoleSelectComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns RoleSelectComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func RoleSelectComponentForModalRequestAsLabelComponentForModalRequestComponent(v *RoleSelectComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		RoleSelectComponentForModalRequest: v,
	}
}

// StringSelectComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns StringSelectComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func StringSelectComponentForModalRequestAsLabelComponentForModalRequestComponent(v *StringSelectComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		StringSelectComponentForModalRequest: v,
	}
}

// TextInputComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns TextInputComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func TextInputComponentForModalRequestAsLabelComponentForModalRequestComponent(v *TextInputComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		TextInputComponentForModalRequest: v,
	}
}

// UserSelectComponentForModalRequestAsLabelComponentForModalRequestComponent is a convenience function that returns UserSelectComponentForModalRequest wrapped in LabelComponentForModalRequestComponent
func UserSelectComponentForModalRequestAsLabelComponentForModalRequestComponent(v *UserSelectComponentForModalRequest) LabelComponentForModalRequestComponent {
	return LabelComponentForModalRequestComponent{
		UserSelectComponentForModalRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LabelComponentForModalRequestComponent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChannelSelectComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.ChannelSelectComponentForModalRequest)
	if err == nil {
		jsonChannelSelectComponentForModalRequest, _ := json.Marshal(dst.ChannelSelectComponentForModalRequest)
		if string(jsonChannelSelectComponentForModalRequest) == "{}" { // empty struct
			dst.ChannelSelectComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.ChannelSelectComponentForModalRequest); err != nil {
				dst.ChannelSelectComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChannelSelectComponentForModalRequest = nil
	}

	// try to unmarshal data into FileUploadComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.FileUploadComponentForModalRequest)
	if err == nil {
		jsonFileUploadComponentForModalRequest, _ := json.Marshal(dst.FileUploadComponentForModalRequest)
		if string(jsonFileUploadComponentForModalRequest) == "{}" { // empty struct
			dst.FileUploadComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.FileUploadComponentForModalRequest); err != nil {
				dst.FileUploadComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.FileUploadComponentForModalRequest = nil
	}

	// try to unmarshal data into MentionableSelectComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.MentionableSelectComponentForModalRequest)
	if err == nil {
		jsonMentionableSelectComponentForModalRequest, _ := json.Marshal(dst.MentionableSelectComponentForModalRequest)
		if string(jsonMentionableSelectComponentForModalRequest) == "{}" { // empty struct
			dst.MentionableSelectComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.MentionableSelectComponentForModalRequest); err != nil {
				dst.MentionableSelectComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionableSelectComponentForModalRequest = nil
	}

	// try to unmarshal data into RoleSelectComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.RoleSelectComponentForModalRequest)
	if err == nil {
		jsonRoleSelectComponentForModalRequest, _ := json.Marshal(dst.RoleSelectComponentForModalRequest)
		if string(jsonRoleSelectComponentForModalRequest) == "{}" { // empty struct
			dst.RoleSelectComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.RoleSelectComponentForModalRequest); err != nil {
				dst.RoleSelectComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelectComponentForModalRequest = nil
	}

	// try to unmarshal data into StringSelectComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.StringSelectComponentForModalRequest)
	if err == nil {
		jsonStringSelectComponentForModalRequest, _ := json.Marshal(dst.StringSelectComponentForModalRequest)
		if string(jsonStringSelectComponentForModalRequest) == "{}" { // empty struct
			dst.StringSelectComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.StringSelectComponentForModalRequest); err != nil {
				dst.StringSelectComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.StringSelectComponentForModalRequest = nil
	}

	// try to unmarshal data into TextInputComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.TextInputComponentForModalRequest)
	if err == nil {
		jsonTextInputComponentForModalRequest, _ := json.Marshal(dst.TextInputComponentForModalRequest)
		if string(jsonTextInputComponentForModalRequest) == "{}" { // empty struct
			dst.TextInputComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.TextInputComponentForModalRequest); err != nil {
				dst.TextInputComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.TextInputComponentForModalRequest = nil
	}

	// try to unmarshal data into UserSelectComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.UserSelectComponentForModalRequest)
	if err == nil {
		jsonUserSelectComponentForModalRequest, _ := json.Marshal(dst.UserSelectComponentForModalRequest)
		if string(jsonUserSelectComponentForModalRequest) == "{}" { // empty struct
			dst.UserSelectComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.UserSelectComponentForModalRequest); err != nil {
				dst.UserSelectComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelectComponentForModalRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChannelSelectComponentForModalRequest = nil
		dst.FileUploadComponentForModalRequest = nil
		dst.MentionableSelectComponentForModalRequest = nil
		dst.RoleSelectComponentForModalRequest = nil
		dst.StringSelectComponentForModalRequest = nil
		dst.TextInputComponentForModalRequest = nil
		dst.UserSelectComponentForModalRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LabelComponentForModalRequestComponent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LabelComponentForModalRequestComponent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LabelComponentForModalRequestComponent) MarshalJSON() ([]byte, error) {
	if src.ChannelSelectComponentForModalRequest != nil {
		return json.Marshal(&src.ChannelSelectComponentForModalRequest)
	}

	if src.FileUploadComponentForModalRequest != nil {
		return json.Marshal(&src.FileUploadComponentForModalRequest)
	}

	if src.MentionableSelectComponentForModalRequest != nil {
		return json.Marshal(&src.MentionableSelectComponentForModalRequest)
	}

	if src.RoleSelectComponentForModalRequest != nil {
		return json.Marshal(&src.RoleSelectComponentForModalRequest)
	}

	if src.StringSelectComponentForModalRequest != nil {
		return json.Marshal(&src.StringSelectComponentForModalRequest)
	}

	if src.TextInputComponentForModalRequest != nil {
		return json.Marshal(&src.TextInputComponentForModalRequest)
	}

	if src.UserSelectComponentForModalRequest != nil {
		return json.Marshal(&src.UserSelectComponentForModalRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LabelComponentForModalRequestComponent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChannelSelectComponentForModalRequest != nil {
		return obj.ChannelSelectComponentForModalRequest
	}

	if obj.FileUploadComponentForModalRequest != nil {
		return obj.FileUploadComponentForModalRequest
	}

	if obj.MentionableSelectComponentForModalRequest != nil {
		return obj.MentionableSelectComponentForModalRequest
	}

	if obj.RoleSelectComponentForModalRequest != nil {
		return obj.RoleSelectComponentForModalRequest
	}

	if obj.StringSelectComponentForModalRequest != nil {
		return obj.StringSelectComponentForModalRequest
	}

	if obj.TextInputComponentForModalRequest != nil {
		return obj.TextInputComponentForModalRequest
	}

	if obj.UserSelectComponentForModalRequest != nil {
		return obj.UserSelectComponentForModalRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj LabelComponentForModalRequestComponent) GetActualInstanceValue() (interface{}) {
	if obj.ChannelSelectComponentForModalRequest != nil {
		return *obj.ChannelSelectComponentForModalRequest
	}

	if obj.FileUploadComponentForModalRequest != nil {
		return *obj.FileUploadComponentForModalRequest
	}

	if obj.MentionableSelectComponentForModalRequest != nil {
		return *obj.MentionableSelectComponentForModalRequest
	}

	if obj.RoleSelectComponentForModalRequest != nil {
		return *obj.RoleSelectComponentForModalRequest
	}

	if obj.StringSelectComponentForModalRequest != nil {
		return *obj.StringSelectComponentForModalRequest
	}

	if obj.TextInputComponentForModalRequest != nil {
		return *obj.TextInputComponentForModalRequest
	}

	if obj.UserSelectComponentForModalRequest != nil {
		return *obj.UserSelectComponentForModalRequest
	}

	// all schemas are nil
	return nil
}

type NullableLabelComponentForModalRequestComponent struct {
	value *LabelComponentForModalRequestComponent
	isSet bool
}

func (v NullableLabelComponentForModalRequestComponent) Get() *LabelComponentForModalRequestComponent {
	return v.value
}

func (v *NullableLabelComponentForModalRequestComponent) Set(val *LabelComponentForModalRequestComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelComponentForModalRequestComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelComponentForModalRequestComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelComponentForModalRequestComponent(val *LabelComponentForModalRequestComponent) *NullableLabelComponentForModalRequestComponent {
	return &NullableLabelComponentForModalRequestComponent{value: val, isSet: true}
}

func (v NullableLabelComponentForModalRequestComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelComponentForModalRequestComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

