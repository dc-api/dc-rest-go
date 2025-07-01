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
)


// ApplicationCommandAutocompleteCallbackRequestData struct for ApplicationCommandAutocompleteCallbackRequestData
type ApplicationCommandAutocompleteCallbackRequestData struct {
	InteractionApplicationCommandAutocompleteCallbackIntegerData *InteractionApplicationCommandAutocompleteCallbackIntegerData
	InteractionApplicationCommandAutocompleteCallbackNumberData *InteractionApplicationCommandAutocompleteCallbackNumberData
	InteractionApplicationCommandAutocompleteCallbackStringData *InteractionApplicationCommandAutocompleteCallbackStringData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ApplicationCommandAutocompleteCallbackRequestData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackIntegerData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackIntegerData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackIntegerData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackIntegerData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackIntegerData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackIntegerData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackIntegerData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackIntegerData = nil
	}

	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackNumberData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackNumberData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackNumberData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackNumberData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackNumberData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackNumberData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackNumberData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackNumberData = nil
	}

	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackStringData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackStringData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackStringData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackStringData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackStringData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackStringData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackStringData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackStringData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ApplicationCommandAutocompleteCallbackRequestData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplicationCommandAutocompleteCallbackRequestData) MarshalJSON() ([]byte, error) {
	if src.InteractionApplicationCommandAutocompleteCallbackIntegerData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackIntegerData)
	}

	if src.InteractionApplicationCommandAutocompleteCallbackNumberData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackNumberData)
	}

	if src.InteractionApplicationCommandAutocompleteCallbackStringData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackStringData)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableApplicationCommandAutocompleteCallbackRequestData struct {
	value *ApplicationCommandAutocompleteCallbackRequestData
	isSet bool
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) Get() *ApplicationCommandAutocompleteCallbackRequestData {
	return v.value
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) Set(val *ApplicationCommandAutocompleteCallbackRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandAutocompleteCallbackRequestData(val *ApplicationCommandAutocompleteCallbackRequestData) *NullableApplicationCommandAutocompleteCallbackRequestData {
	return &NullableApplicationCommandAutocompleteCallbackRequestData{value: val, isSet: true}
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


