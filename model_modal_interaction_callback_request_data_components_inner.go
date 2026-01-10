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

// ModalInteractionCallbackRequestDataComponentsInner - struct for ModalInteractionCallbackRequestDataComponentsInner
type ModalInteractionCallbackRequestDataComponentsInner struct {
	ActionRowComponentForModalRequest *ActionRowComponentForModalRequest
	LabelComponentForModalRequest *LabelComponentForModalRequest
	TextDisplayComponentForModalRequest *TextDisplayComponentForModalRequest
}

// ActionRowComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner is a convenience function that returns ActionRowComponentForModalRequest wrapped in ModalInteractionCallbackRequestDataComponentsInner
func ActionRowComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner(v *ActionRowComponentForModalRequest) ModalInteractionCallbackRequestDataComponentsInner {
	return ModalInteractionCallbackRequestDataComponentsInner{
		ActionRowComponentForModalRequest: v,
	}
}

// LabelComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner is a convenience function that returns LabelComponentForModalRequest wrapped in ModalInteractionCallbackRequestDataComponentsInner
func LabelComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner(v *LabelComponentForModalRequest) ModalInteractionCallbackRequestDataComponentsInner {
	return ModalInteractionCallbackRequestDataComponentsInner{
		LabelComponentForModalRequest: v,
	}
}

// TextDisplayComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner is a convenience function that returns TextDisplayComponentForModalRequest wrapped in ModalInteractionCallbackRequestDataComponentsInner
func TextDisplayComponentForModalRequestAsModalInteractionCallbackRequestDataComponentsInner(v *TextDisplayComponentForModalRequest) ModalInteractionCallbackRequestDataComponentsInner {
	return ModalInteractionCallbackRequestDataComponentsInner{
		TextDisplayComponentForModalRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ModalInteractionCallbackRequestDataComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionRowComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.ActionRowComponentForModalRequest)
	if err == nil {
		jsonActionRowComponentForModalRequest, _ := json.Marshal(dst.ActionRowComponentForModalRequest)
		if string(jsonActionRowComponentForModalRequest) == "{}" { // empty struct
			dst.ActionRowComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.ActionRowComponentForModalRequest); err != nil {
				dst.ActionRowComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActionRowComponentForModalRequest = nil
	}

	// try to unmarshal data into LabelComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.LabelComponentForModalRequest)
	if err == nil {
		jsonLabelComponentForModalRequest, _ := json.Marshal(dst.LabelComponentForModalRequest)
		if string(jsonLabelComponentForModalRequest) == "{}" { // empty struct
			dst.LabelComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.LabelComponentForModalRequest); err != nil {
				dst.LabelComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.LabelComponentForModalRequest = nil
	}

	// try to unmarshal data into TextDisplayComponentForModalRequest
	err = newStrictDecoder(data).Decode(&dst.TextDisplayComponentForModalRequest)
	if err == nil {
		jsonTextDisplayComponentForModalRequest, _ := json.Marshal(dst.TextDisplayComponentForModalRequest)
		if string(jsonTextDisplayComponentForModalRequest) == "{}" { // empty struct
			dst.TextDisplayComponentForModalRequest = nil
		} else {
			if err = validator.Validate(dst.TextDisplayComponentForModalRequest); err != nil {
				dst.TextDisplayComponentForModalRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.TextDisplayComponentForModalRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionRowComponentForModalRequest = nil
		dst.LabelComponentForModalRequest = nil
		dst.TextDisplayComponentForModalRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ModalInteractionCallbackRequestDataComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ModalInteractionCallbackRequestDataComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ModalInteractionCallbackRequestDataComponentsInner) MarshalJSON() ([]byte, error) {
	if src.ActionRowComponentForModalRequest != nil {
		return json.Marshal(&src.ActionRowComponentForModalRequest)
	}

	if src.LabelComponentForModalRequest != nil {
		return json.Marshal(&src.LabelComponentForModalRequest)
	}

	if src.TextDisplayComponentForModalRequest != nil {
		return json.Marshal(&src.TextDisplayComponentForModalRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ModalInteractionCallbackRequestDataComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActionRowComponentForModalRequest != nil {
		return obj.ActionRowComponentForModalRequest
	}

	if obj.LabelComponentForModalRequest != nil {
		return obj.LabelComponentForModalRequest
	}

	if obj.TextDisplayComponentForModalRequest != nil {
		return obj.TextDisplayComponentForModalRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ModalInteractionCallbackRequestDataComponentsInner) GetActualInstanceValue() (interface{}) {
	if obj.ActionRowComponentForModalRequest != nil {
		return *obj.ActionRowComponentForModalRequest
	}

	if obj.LabelComponentForModalRequest != nil {
		return *obj.LabelComponentForModalRequest
	}

	if obj.TextDisplayComponentForModalRequest != nil {
		return *obj.TextDisplayComponentForModalRequest
	}

	// all schemas are nil
	return nil
}

type NullableModalInteractionCallbackRequestDataComponentsInner struct {
	value *ModalInteractionCallbackRequestDataComponentsInner
	isSet bool
}

func (v NullableModalInteractionCallbackRequestDataComponentsInner) Get() *ModalInteractionCallbackRequestDataComponentsInner {
	return v.value
}

func (v *NullableModalInteractionCallbackRequestDataComponentsInner) Set(val *ModalInteractionCallbackRequestDataComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModalInteractionCallbackRequestDataComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModalInteractionCallbackRequestDataComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModalInteractionCallbackRequestDataComponentsInner(val *ModalInteractionCallbackRequestDataComponentsInner) *NullableModalInteractionCallbackRequestDataComponentsInner {
	return &NullableModalInteractionCallbackRequestDataComponentsInner{value: val, isSet: true}
}

func (v NullableModalInteractionCallbackRequestDataComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModalInteractionCallbackRequestDataComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

