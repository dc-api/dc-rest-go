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

// ApplicationCommandSubcommandOptionOptionsInner - struct for ApplicationCommandSubcommandOptionOptionsInner
type ApplicationCommandSubcommandOptionOptionsInner struct {
	ApplicationCommandAttachmentOption *ApplicationCommandAttachmentOption
	ApplicationCommandBooleanOption *ApplicationCommandBooleanOption
	ApplicationCommandChannelOption *ApplicationCommandChannelOption
	ApplicationCommandIntegerOption *ApplicationCommandIntegerOption
	ApplicationCommandMentionableOption *ApplicationCommandMentionableOption
	ApplicationCommandNumberOption *ApplicationCommandNumberOption
	ApplicationCommandRoleOption *ApplicationCommandRoleOption
	ApplicationCommandStringOption *ApplicationCommandStringOption
	ApplicationCommandUserOption *ApplicationCommandUserOption
}

// ApplicationCommandAttachmentOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandAttachmentOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandAttachmentOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandAttachmentOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandAttachmentOption: v,
	}
}

// ApplicationCommandBooleanOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandBooleanOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandBooleanOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandBooleanOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandBooleanOption: v,
	}
}

// ApplicationCommandChannelOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandChannelOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandChannelOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandChannelOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandChannelOption: v,
	}
}

// ApplicationCommandIntegerOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandIntegerOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandIntegerOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandIntegerOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandIntegerOption: v,
	}
}

// ApplicationCommandMentionableOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandMentionableOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandMentionableOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandMentionableOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandMentionableOption: v,
	}
}

// ApplicationCommandNumberOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandNumberOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandNumberOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandNumberOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandNumberOption: v,
	}
}

// ApplicationCommandRoleOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandRoleOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandRoleOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandRoleOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandRoleOption: v,
	}
}

// ApplicationCommandStringOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandStringOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandStringOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandStringOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandStringOption: v,
	}
}

// ApplicationCommandUserOptionAsApplicationCommandSubcommandOptionOptionsInner is a convenience function that returns ApplicationCommandUserOption wrapped in ApplicationCommandSubcommandOptionOptionsInner
func ApplicationCommandUserOptionAsApplicationCommandSubcommandOptionOptionsInner(v *ApplicationCommandUserOption) ApplicationCommandSubcommandOptionOptionsInner {
	return ApplicationCommandSubcommandOptionOptionsInner{
		ApplicationCommandUserOption: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApplicationCommandSubcommandOptionOptionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplicationCommandAttachmentOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandAttachmentOption)
	if err == nil {
		jsonApplicationCommandAttachmentOption, _ := json.Marshal(dst.ApplicationCommandAttachmentOption)
		if string(jsonApplicationCommandAttachmentOption) == "{}" { // empty struct
			dst.ApplicationCommandAttachmentOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandAttachmentOption); err != nil {
				dst.ApplicationCommandAttachmentOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandAttachmentOption = nil
	}

	// try to unmarshal data into ApplicationCommandBooleanOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandBooleanOption)
	if err == nil {
		jsonApplicationCommandBooleanOption, _ := json.Marshal(dst.ApplicationCommandBooleanOption)
		if string(jsonApplicationCommandBooleanOption) == "{}" { // empty struct
			dst.ApplicationCommandBooleanOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandBooleanOption); err != nil {
				dst.ApplicationCommandBooleanOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandBooleanOption = nil
	}

	// try to unmarshal data into ApplicationCommandChannelOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandChannelOption)
	if err == nil {
		jsonApplicationCommandChannelOption, _ := json.Marshal(dst.ApplicationCommandChannelOption)
		if string(jsonApplicationCommandChannelOption) == "{}" { // empty struct
			dst.ApplicationCommandChannelOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandChannelOption); err != nil {
				dst.ApplicationCommandChannelOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandChannelOption = nil
	}

	// try to unmarshal data into ApplicationCommandIntegerOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandIntegerOption)
	if err == nil {
		jsonApplicationCommandIntegerOption, _ := json.Marshal(dst.ApplicationCommandIntegerOption)
		if string(jsonApplicationCommandIntegerOption) == "{}" { // empty struct
			dst.ApplicationCommandIntegerOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandIntegerOption); err != nil {
				dst.ApplicationCommandIntegerOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandIntegerOption = nil
	}

	// try to unmarshal data into ApplicationCommandMentionableOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandMentionableOption)
	if err == nil {
		jsonApplicationCommandMentionableOption, _ := json.Marshal(dst.ApplicationCommandMentionableOption)
		if string(jsonApplicationCommandMentionableOption) == "{}" { // empty struct
			dst.ApplicationCommandMentionableOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandMentionableOption); err != nil {
				dst.ApplicationCommandMentionableOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandMentionableOption = nil
	}

	// try to unmarshal data into ApplicationCommandNumberOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandNumberOption)
	if err == nil {
		jsonApplicationCommandNumberOption, _ := json.Marshal(dst.ApplicationCommandNumberOption)
		if string(jsonApplicationCommandNumberOption) == "{}" { // empty struct
			dst.ApplicationCommandNumberOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandNumberOption); err != nil {
				dst.ApplicationCommandNumberOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandNumberOption = nil
	}

	// try to unmarshal data into ApplicationCommandRoleOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandRoleOption)
	if err == nil {
		jsonApplicationCommandRoleOption, _ := json.Marshal(dst.ApplicationCommandRoleOption)
		if string(jsonApplicationCommandRoleOption) == "{}" { // empty struct
			dst.ApplicationCommandRoleOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandRoleOption); err != nil {
				dst.ApplicationCommandRoleOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandRoleOption = nil
	}

	// try to unmarshal data into ApplicationCommandStringOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandStringOption)
	if err == nil {
		jsonApplicationCommandStringOption, _ := json.Marshal(dst.ApplicationCommandStringOption)
		if string(jsonApplicationCommandStringOption) == "{}" { // empty struct
			dst.ApplicationCommandStringOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandStringOption); err != nil {
				dst.ApplicationCommandStringOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandStringOption = nil
	}

	// try to unmarshal data into ApplicationCommandUserOption
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandUserOption)
	if err == nil {
		jsonApplicationCommandUserOption, _ := json.Marshal(dst.ApplicationCommandUserOption)
		if string(jsonApplicationCommandUserOption) == "{}" { // empty struct
			dst.ApplicationCommandUserOption = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandUserOption); err != nil {
				dst.ApplicationCommandUserOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandUserOption = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationCommandAttachmentOption = nil
		dst.ApplicationCommandBooleanOption = nil
		dst.ApplicationCommandChannelOption = nil
		dst.ApplicationCommandIntegerOption = nil
		dst.ApplicationCommandMentionableOption = nil
		dst.ApplicationCommandNumberOption = nil
		dst.ApplicationCommandRoleOption = nil
		dst.ApplicationCommandStringOption = nil
		dst.ApplicationCommandUserOption = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApplicationCommandSubcommandOptionOptionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApplicationCommandSubcommandOptionOptionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplicationCommandSubcommandOptionOptionsInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationCommandAttachmentOption != nil {
		return json.Marshal(&src.ApplicationCommandAttachmentOption)
	}

	if src.ApplicationCommandBooleanOption != nil {
		return json.Marshal(&src.ApplicationCommandBooleanOption)
	}

	if src.ApplicationCommandChannelOption != nil {
		return json.Marshal(&src.ApplicationCommandChannelOption)
	}

	if src.ApplicationCommandIntegerOption != nil {
		return json.Marshal(&src.ApplicationCommandIntegerOption)
	}

	if src.ApplicationCommandMentionableOption != nil {
		return json.Marshal(&src.ApplicationCommandMentionableOption)
	}

	if src.ApplicationCommandNumberOption != nil {
		return json.Marshal(&src.ApplicationCommandNumberOption)
	}

	if src.ApplicationCommandRoleOption != nil {
		return json.Marshal(&src.ApplicationCommandRoleOption)
	}

	if src.ApplicationCommandStringOption != nil {
		return json.Marshal(&src.ApplicationCommandStringOption)
	}

	if src.ApplicationCommandUserOption != nil {
		return json.Marshal(&src.ApplicationCommandUserOption)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApplicationCommandSubcommandOptionOptionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationCommandAttachmentOption != nil {
		return obj.ApplicationCommandAttachmentOption
	}

	if obj.ApplicationCommandBooleanOption != nil {
		return obj.ApplicationCommandBooleanOption
	}

	if obj.ApplicationCommandChannelOption != nil {
		return obj.ApplicationCommandChannelOption
	}

	if obj.ApplicationCommandIntegerOption != nil {
		return obj.ApplicationCommandIntegerOption
	}

	if obj.ApplicationCommandMentionableOption != nil {
		return obj.ApplicationCommandMentionableOption
	}

	if obj.ApplicationCommandNumberOption != nil {
		return obj.ApplicationCommandNumberOption
	}

	if obj.ApplicationCommandRoleOption != nil {
		return obj.ApplicationCommandRoleOption
	}

	if obj.ApplicationCommandStringOption != nil {
		return obj.ApplicationCommandStringOption
	}

	if obj.ApplicationCommandUserOption != nil {
		return obj.ApplicationCommandUserOption
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApplicationCommandSubcommandOptionOptionsInner) GetActualInstanceValue() (interface{}) {
	if obj.ApplicationCommandAttachmentOption != nil {
		return *obj.ApplicationCommandAttachmentOption
	}

	if obj.ApplicationCommandBooleanOption != nil {
		return *obj.ApplicationCommandBooleanOption
	}

	if obj.ApplicationCommandChannelOption != nil {
		return *obj.ApplicationCommandChannelOption
	}

	if obj.ApplicationCommandIntegerOption != nil {
		return *obj.ApplicationCommandIntegerOption
	}

	if obj.ApplicationCommandMentionableOption != nil {
		return *obj.ApplicationCommandMentionableOption
	}

	if obj.ApplicationCommandNumberOption != nil {
		return *obj.ApplicationCommandNumberOption
	}

	if obj.ApplicationCommandRoleOption != nil {
		return *obj.ApplicationCommandRoleOption
	}

	if obj.ApplicationCommandStringOption != nil {
		return *obj.ApplicationCommandStringOption
	}

	if obj.ApplicationCommandUserOption != nil {
		return *obj.ApplicationCommandUserOption
	}

	// all schemas are nil
	return nil
}

type NullableApplicationCommandSubcommandOptionOptionsInner struct {
	value *ApplicationCommandSubcommandOptionOptionsInner
	isSet bool
}

func (v NullableApplicationCommandSubcommandOptionOptionsInner) Get() *ApplicationCommandSubcommandOptionOptionsInner {
	return v.value
}

func (v *NullableApplicationCommandSubcommandOptionOptionsInner) Set(val *ApplicationCommandSubcommandOptionOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandSubcommandOptionOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandSubcommandOptionOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandSubcommandOptionOptionsInner(val *ApplicationCommandSubcommandOptionOptionsInner) *NullableApplicationCommandSubcommandOptionOptionsInner {
	return &NullableApplicationCommandSubcommandOptionOptionsInner{value: val, isSet: true}
}

func (v NullableApplicationCommandSubcommandOptionOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandSubcommandOptionOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


