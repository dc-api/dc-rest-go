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

// BasicMessageResponseInteractionMetadata - struct for BasicMessageResponseInteractionMetadata
type BasicMessageResponseInteractionMetadata struct {
	ApplicationCommandInteractionMetadataResponse *ApplicationCommandInteractionMetadataResponse
	MessageComponentInteractionMetadataResponse *MessageComponentInteractionMetadataResponse
	ModalSubmitInteractionMetadataResponse *ModalSubmitInteractionMetadataResponse
}

// ApplicationCommandInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata is a convenience function that returns ApplicationCommandInteractionMetadataResponse wrapped in BasicMessageResponseInteractionMetadata
func ApplicationCommandInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata(v *ApplicationCommandInteractionMetadataResponse) BasicMessageResponseInteractionMetadata {
	return BasicMessageResponseInteractionMetadata{
		ApplicationCommandInteractionMetadataResponse: v,
	}
}

// MessageComponentInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata is a convenience function that returns MessageComponentInteractionMetadataResponse wrapped in BasicMessageResponseInteractionMetadata
func MessageComponentInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata(v *MessageComponentInteractionMetadataResponse) BasicMessageResponseInteractionMetadata {
	return BasicMessageResponseInteractionMetadata{
		MessageComponentInteractionMetadataResponse: v,
	}
}

// ModalSubmitInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata is a convenience function that returns ModalSubmitInteractionMetadataResponse wrapped in BasicMessageResponseInteractionMetadata
func ModalSubmitInteractionMetadataResponseAsBasicMessageResponseInteractionMetadata(v *ModalSubmitInteractionMetadataResponse) BasicMessageResponseInteractionMetadata {
	return BasicMessageResponseInteractionMetadata{
		ModalSubmitInteractionMetadataResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BasicMessageResponseInteractionMetadata) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ApplicationCommandInteractionMetadataResponse
	err = newStrictDecoder(data).Decode(&dst.ApplicationCommandInteractionMetadataResponse)
	if err == nil {
		jsonApplicationCommandInteractionMetadataResponse, _ := json.Marshal(dst.ApplicationCommandInteractionMetadataResponse)
		if string(jsonApplicationCommandInteractionMetadataResponse) == "{}" { // empty struct
			dst.ApplicationCommandInteractionMetadataResponse = nil
		} else {
			if err = validator.Validate(dst.ApplicationCommandInteractionMetadataResponse); err != nil {
				dst.ApplicationCommandInteractionMetadataResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationCommandInteractionMetadataResponse = nil
	}

	// try to unmarshal data into MessageComponentInteractionMetadataResponse
	err = newStrictDecoder(data).Decode(&dst.MessageComponentInteractionMetadataResponse)
	if err == nil {
		jsonMessageComponentInteractionMetadataResponse, _ := json.Marshal(dst.MessageComponentInteractionMetadataResponse)
		if string(jsonMessageComponentInteractionMetadataResponse) == "{}" { // empty struct
			dst.MessageComponentInteractionMetadataResponse = nil
		} else {
			if err = validator.Validate(dst.MessageComponentInteractionMetadataResponse); err != nil {
				dst.MessageComponentInteractionMetadataResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageComponentInteractionMetadataResponse = nil
	}

	// try to unmarshal data into ModalSubmitInteractionMetadataResponse
	err = newStrictDecoder(data).Decode(&dst.ModalSubmitInteractionMetadataResponse)
	if err == nil {
		jsonModalSubmitInteractionMetadataResponse, _ := json.Marshal(dst.ModalSubmitInteractionMetadataResponse)
		if string(jsonModalSubmitInteractionMetadataResponse) == "{}" { // empty struct
			dst.ModalSubmitInteractionMetadataResponse = nil
		} else {
			if err = validator.Validate(dst.ModalSubmitInteractionMetadataResponse); err != nil {
				dst.ModalSubmitInteractionMetadataResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ModalSubmitInteractionMetadataResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationCommandInteractionMetadataResponse = nil
		dst.MessageComponentInteractionMetadataResponse = nil
		dst.ModalSubmitInteractionMetadataResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BasicMessageResponseInteractionMetadata)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BasicMessageResponseInteractionMetadata)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BasicMessageResponseInteractionMetadata) MarshalJSON() ([]byte, error) {
	if src.ApplicationCommandInteractionMetadataResponse != nil {
		return json.Marshal(&src.ApplicationCommandInteractionMetadataResponse)
	}

	if src.MessageComponentInteractionMetadataResponse != nil {
		return json.Marshal(&src.MessageComponentInteractionMetadataResponse)
	}

	if src.ModalSubmitInteractionMetadataResponse != nil {
		return json.Marshal(&src.ModalSubmitInteractionMetadataResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BasicMessageResponseInteractionMetadata) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationCommandInteractionMetadataResponse != nil {
		return obj.ApplicationCommandInteractionMetadataResponse
	}

	if obj.MessageComponentInteractionMetadataResponse != nil {
		return obj.MessageComponentInteractionMetadataResponse
	}

	if obj.ModalSubmitInteractionMetadataResponse != nil {
		return obj.ModalSubmitInteractionMetadataResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BasicMessageResponseInteractionMetadata) GetActualInstanceValue() (interface{}) {
	if obj.ApplicationCommandInteractionMetadataResponse != nil {
		return *obj.ApplicationCommandInteractionMetadataResponse
	}

	if obj.MessageComponentInteractionMetadataResponse != nil {
		return *obj.MessageComponentInteractionMetadataResponse
	}

	if obj.ModalSubmitInteractionMetadataResponse != nil {
		return *obj.ModalSubmitInteractionMetadataResponse
	}

	// all schemas are nil
	return nil
}

type NullableBasicMessageResponseInteractionMetadata struct {
	value *BasicMessageResponseInteractionMetadata
	isSet bool
}

func (v NullableBasicMessageResponseInteractionMetadata) Get() *BasicMessageResponseInteractionMetadata {
	return v.value
}

func (v *NullableBasicMessageResponseInteractionMetadata) Set(val *BasicMessageResponseInteractionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMessageResponseInteractionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMessageResponseInteractionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMessageResponseInteractionMetadata(val *BasicMessageResponseInteractionMetadata) *NullableBasicMessageResponseInteractionMetadata {
	return &NullableBasicMessageResponseInteractionMetadata{value: val, isSet: true}
}

func (v NullableBasicMessageResponseInteractionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMessageResponseInteractionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


