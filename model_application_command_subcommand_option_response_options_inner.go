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

// ApplicationCommandSubcommandOptionResponseOptionsInner - struct for ApplicationCommandSubcommandOptionResponseOptionsInner
type ApplicationCommandSubcommandOptionResponseOptionsInner struct {
	ApplicationCommandAttachmentOptionResponse *ApplicationCommandAttachmentOptionResponse
	ApplicationCommandBooleanOptionResponse *ApplicationCommandBooleanOptionResponse
	ApplicationCommandChannelOptionResponse *ApplicationCommandChannelOptionResponse
	ApplicationCommandIntegerOptionResponse *ApplicationCommandIntegerOptionResponse
	ApplicationCommandMentionableOptionResponse *ApplicationCommandMentionableOptionResponse
	ApplicationCommandNumberOptionResponse *ApplicationCommandNumberOptionResponse
	ApplicationCommandRoleOptionResponse *ApplicationCommandRoleOptionResponse
	ApplicationCommandStringOptionResponse *ApplicationCommandStringOptionResponse
	ApplicationCommandUserOptionResponse *ApplicationCommandUserOptionResponse
}

// ApplicationCommandAttachmentOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandAttachmentOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandAttachmentOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandAttachmentOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandAttachmentOptionResponse: v,
	}
}

// ApplicationCommandBooleanOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandBooleanOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandBooleanOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandBooleanOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandBooleanOptionResponse: v,
	}
}

// ApplicationCommandChannelOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandChannelOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandChannelOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandChannelOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandChannelOptionResponse: v,
	}
}

// ApplicationCommandIntegerOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandIntegerOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandIntegerOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandIntegerOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandIntegerOptionResponse: v,
	}
}

// ApplicationCommandMentionableOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandMentionableOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandMentionableOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandMentionableOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandMentionableOptionResponse: v,
	}
}

// ApplicationCommandNumberOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandNumberOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandNumberOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandNumberOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandNumberOptionResponse: v,
	}
}

// ApplicationCommandRoleOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandRoleOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandRoleOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandRoleOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandRoleOptionResponse: v,
	}
}

// ApplicationCommandStringOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandStringOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandStringOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandStringOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandStringOptionResponse: v,
	}
}

// ApplicationCommandUserOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner is a convenience function that returns ApplicationCommandUserOptionResponse wrapped in ApplicationCommandSubcommandOptionResponseOptionsInner
func ApplicationCommandUserOptionResponseAsApplicationCommandSubcommandOptionResponseOptionsInner(v *ApplicationCommandUserOptionResponse) ApplicationCommandSubcommandOptionResponseOptionsInner {
	return ApplicationCommandSubcommandOptionResponseOptionsInner{
		ApplicationCommandUserOptionResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApplicationCommandSubcommandOptionResponseOptionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplicationCommandAttachmentOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandAttachmentOptionResponse)
	if err == nil {
		jsonApplicationCommandAttachmentOptionResponse, _ := json.Marshal(dst.ApplicationCommandAttachmentOptionResponse)
		if string(jsonApplicationCommandAttachmentOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandAttachmentOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandAttachmentOptionResponse); err != nil {
				dst.ApplicationCommandAttachmentOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandAttachmentOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandBooleanOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandBooleanOptionResponse)
	if err == nil {
		jsonApplicationCommandBooleanOptionResponse, _ := json.Marshal(dst.ApplicationCommandBooleanOptionResponse)
		if string(jsonApplicationCommandBooleanOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandBooleanOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandBooleanOptionResponse); err != nil {
				dst.ApplicationCommandBooleanOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandBooleanOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandChannelOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandChannelOptionResponse)
	if err == nil {
		jsonApplicationCommandChannelOptionResponse, _ := json.Marshal(dst.ApplicationCommandChannelOptionResponse)
		if string(jsonApplicationCommandChannelOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandChannelOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandChannelOptionResponse); err != nil {
				dst.ApplicationCommandChannelOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandChannelOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandIntegerOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandIntegerOptionResponse)
	if err == nil {
		jsonApplicationCommandIntegerOptionResponse, _ := json.Marshal(dst.ApplicationCommandIntegerOptionResponse)
		if string(jsonApplicationCommandIntegerOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandIntegerOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandIntegerOptionResponse); err != nil {
				dst.ApplicationCommandIntegerOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandIntegerOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandMentionableOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandMentionableOptionResponse)
	if err == nil {
		jsonApplicationCommandMentionableOptionResponse, _ := json.Marshal(dst.ApplicationCommandMentionableOptionResponse)
		if string(jsonApplicationCommandMentionableOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandMentionableOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandMentionableOptionResponse); err != nil {
				dst.ApplicationCommandMentionableOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandMentionableOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandNumberOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandNumberOptionResponse)
	if err == nil {
		jsonApplicationCommandNumberOptionResponse, _ := json.Marshal(dst.ApplicationCommandNumberOptionResponse)
		if string(jsonApplicationCommandNumberOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandNumberOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandNumberOptionResponse); err != nil {
				dst.ApplicationCommandNumberOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandNumberOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandRoleOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandRoleOptionResponse)
	if err == nil {
		jsonApplicationCommandRoleOptionResponse, _ := json.Marshal(dst.ApplicationCommandRoleOptionResponse)
		if string(jsonApplicationCommandRoleOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandRoleOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandRoleOptionResponse); err != nil {
				dst.ApplicationCommandRoleOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandRoleOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandStringOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandStringOptionResponse)
	if err == nil {
		jsonApplicationCommandStringOptionResponse, _ := json.Marshal(dst.ApplicationCommandStringOptionResponse)
		if string(jsonApplicationCommandStringOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandStringOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandStringOptionResponse); err != nil {
				dst.ApplicationCommandStringOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandStringOptionResponse = nil
	}

	// try to unmarshal data into ApplicationCommandUserOptionResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandUserOptionResponse)
	if err == nil {
		jsonApplicationCommandUserOptionResponse, _ := json.Marshal(dst.ApplicationCommandUserOptionResponse)
		if string(jsonApplicationCommandUserOptionResponse) == "{}" { // empty struct
			dst.ApplicationCommandUserOptionResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandUserOptionResponse); err != nil {
				dst.ApplicationCommandUserOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandUserOptionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationCommandAttachmentOptionResponse = nil
		dst.ApplicationCommandBooleanOptionResponse = nil
		dst.ApplicationCommandChannelOptionResponse = nil
		dst.ApplicationCommandIntegerOptionResponse = nil
		dst.ApplicationCommandMentionableOptionResponse = nil
		dst.ApplicationCommandNumberOptionResponse = nil
		dst.ApplicationCommandRoleOptionResponse = nil
		dst.ApplicationCommandStringOptionResponse = nil
		dst.ApplicationCommandUserOptionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApplicationCommandSubcommandOptionResponseOptionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApplicationCommandSubcommandOptionResponseOptionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplicationCommandSubcommandOptionResponseOptionsInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationCommandAttachmentOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandAttachmentOptionResponse)
	}

	if src.ApplicationCommandBooleanOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandBooleanOptionResponse)
	}

	if src.ApplicationCommandChannelOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandChannelOptionResponse)
	}

	if src.ApplicationCommandIntegerOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandIntegerOptionResponse)
	}

	if src.ApplicationCommandMentionableOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandMentionableOptionResponse)
	}

	if src.ApplicationCommandNumberOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandNumberOptionResponse)
	}

	if src.ApplicationCommandRoleOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandRoleOptionResponse)
	}

	if src.ApplicationCommandStringOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandStringOptionResponse)
	}

	if src.ApplicationCommandUserOptionResponse != nil {
		return json.Marshal(&src.ApplicationCommandUserOptionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApplicationCommandSubcommandOptionResponseOptionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationCommandAttachmentOptionResponse != nil {
		return obj.ApplicationCommandAttachmentOptionResponse
	}

	if obj.ApplicationCommandBooleanOptionResponse != nil {
		return obj.ApplicationCommandBooleanOptionResponse
	}

	if obj.ApplicationCommandChannelOptionResponse != nil {
		return obj.ApplicationCommandChannelOptionResponse
	}

	if obj.ApplicationCommandIntegerOptionResponse != nil {
		return obj.ApplicationCommandIntegerOptionResponse
	}

	if obj.ApplicationCommandMentionableOptionResponse != nil {
		return obj.ApplicationCommandMentionableOptionResponse
	}

	if obj.ApplicationCommandNumberOptionResponse != nil {
		return obj.ApplicationCommandNumberOptionResponse
	}

	if obj.ApplicationCommandRoleOptionResponse != nil {
		return obj.ApplicationCommandRoleOptionResponse
	}

	if obj.ApplicationCommandStringOptionResponse != nil {
		return obj.ApplicationCommandStringOptionResponse
	}

	if obj.ApplicationCommandUserOptionResponse != nil {
		return obj.ApplicationCommandUserOptionResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApplicationCommandSubcommandOptionResponseOptionsInner) GetActualInstanceValue() (interface{}) {
	if obj.ApplicationCommandAttachmentOptionResponse != nil {
		return *obj.ApplicationCommandAttachmentOptionResponse
	}

	if obj.ApplicationCommandBooleanOptionResponse != nil {
		return *obj.ApplicationCommandBooleanOptionResponse
	}

	if obj.ApplicationCommandChannelOptionResponse != nil {
		return *obj.ApplicationCommandChannelOptionResponse
	}

	if obj.ApplicationCommandIntegerOptionResponse != nil {
		return *obj.ApplicationCommandIntegerOptionResponse
	}

	if obj.ApplicationCommandMentionableOptionResponse != nil {
		return *obj.ApplicationCommandMentionableOptionResponse
	}

	if obj.ApplicationCommandNumberOptionResponse != nil {
		return *obj.ApplicationCommandNumberOptionResponse
	}

	if obj.ApplicationCommandRoleOptionResponse != nil {
		return *obj.ApplicationCommandRoleOptionResponse
	}

	if obj.ApplicationCommandStringOptionResponse != nil {
		return *obj.ApplicationCommandStringOptionResponse
	}

	if obj.ApplicationCommandUserOptionResponse != nil {
		return *obj.ApplicationCommandUserOptionResponse
	}

	// all schemas are nil
	return nil
}

type NullableApplicationCommandSubcommandOptionResponseOptionsInner struct {
	value *ApplicationCommandSubcommandOptionResponseOptionsInner
	isSet bool
}

func (v NullableApplicationCommandSubcommandOptionResponseOptionsInner) Get() *ApplicationCommandSubcommandOptionResponseOptionsInner {
	return v.value
}

func (v *NullableApplicationCommandSubcommandOptionResponseOptionsInner) Set(val *ApplicationCommandSubcommandOptionResponseOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandSubcommandOptionResponseOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandSubcommandOptionResponseOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandSubcommandOptionResponseOptionsInner(val *ApplicationCommandSubcommandOptionResponseOptionsInner) *NullableApplicationCommandSubcommandOptionResponseOptionsInner {
	return &NullableApplicationCommandSubcommandOptionResponseOptionsInner{value: val, isSet: true}
}

func (v NullableApplicationCommandSubcommandOptionResponseOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandSubcommandOptionResponseOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


