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


// CreateInteractionResponseRequest struct for CreateInteractionResponseRequest
type CreateInteractionResponseRequest struct {
	ApplicationCommandAutocompleteCallbackRequest *ApplicationCommandAutocompleteCallbackRequest
	CreateMessageInteractionCallbackRequest *CreateMessageInteractionCallbackRequest
	LaunchActivityInteractionCallbackRequest *LaunchActivityInteractionCallbackRequest
	ModalInteractionCallbackRequest *ModalInteractionCallbackRequest
	PongInteractionCallbackRequest *PongInteractionCallbackRequest
	UpdateMessageInteractionCallbackRequest *UpdateMessageInteractionCallbackRequest
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateInteractionResponseRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ApplicationCommandAutocompleteCallbackRequest
	err = json.Unmarshal(data, &dst.ApplicationCommandAutocompleteCallbackRequest);
	if err == nil {
		jsonApplicationCommandAutocompleteCallbackRequest, _ := json.Marshal(dst.ApplicationCommandAutocompleteCallbackRequest)
		if string(jsonApplicationCommandAutocompleteCallbackRequest) == "{}" { // empty struct
			dst.ApplicationCommandAutocompleteCallbackRequest = nil
		} else {
			return nil // data stored in dst.ApplicationCommandAutocompleteCallbackRequest, return on the first match
		}
	} else {
		dst.ApplicationCommandAutocompleteCallbackRequest = nil
	}

	// try to unmarshal JSON data into CreateMessageInteractionCallbackRequest
	err = json.Unmarshal(data, &dst.CreateMessageInteractionCallbackRequest);
	if err == nil {
		jsonCreateMessageInteractionCallbackRequest, _ := json.Marshal(dst.CreateMessageInteractionCallbackRequest)
		if string(jsonCreateMessageInteractionCallbackRequest) == "{}" { // empty struct
			dst.CreateMessageInteractionCallbackRequest = nil
		} else {
			return nil // data stored in dst.CreateMessageInteractionCallbackRequest, return on the first match
		}
	} else {
		dst.CreateMessageInteractionCallbackRequest = nil
	}

	// try to unmarshal JSON data into LaunchActivityInteractionCallbackRequest
	err = json.Unmarshal(data, &dst.LaunchActivityInteractionCallbackRequest);
	if err == nil {
		jsonLaunchActivityInteractionCallbackRequest, _ := json.Marshal(dst.LaunchActivityInteractionCallbackRequest)
		if string(jsonLaunchActivityInteractionCallbackRequest) == "{}" { // empty struct
			dst.LaunchActivityInteractionCallbackRequest = nil
		} else {
			return nil // data stored in dst.LaunchActivityInteractionCallbackRequest, return on the first match
		}
	} else {
		dst.LaunchActivityInteractionCallbackRequest = nil
	}

	// try to unmarshal JSON data into ModalInteractionCallbackRequest
	err = json.Unmarshal(data, &dst.ModalInteractionCallbackRequest);
	if err == nil {
		jsonModalInteractionCallbackRequest, _ := json.Marshal(dst.ModalInteractionCallbackRequest)
		if string(jsonModalInteractionCallbackRequest) == "{}" { // empty struct
			dst.ModalInteractionCallbackRequest = nil
		} else {
			return nil // data stored in dst.ModalInteractionCallbackRequest, return on the first match
		}
	} else {
		dst.ModalInteractionCallbackRequest = nil
	}

	// try to unmarshal JSON data into PongInteractionCallbackRequest
	err = json.Unmarshal(data, &dst.PongInteractionCallbackRequest);
	if err == nil {
		jsonPongInteractionCallbackRequest, _ := json.Marshal(dst.PongInteractionCallbackRequest)
		if string(jsonPongInteractionCallbackRequest) == "{}" { // empty struct
			dst.PongInteractionCallbackRequest = nil
		} else {
			return nil // data stored in dst.PongInteractionCallbackRequest, return on the first match
		}
	} else {
		dst.PongInteractionCallbackRequest = nil
	}

	// try to unmarshal JSON data into UpdateMessageInteractionCallbackRequest
	err = json.Unmarshal(data, &dst.UpdateMessageInteractionCallbackRequest);
	if err == nil {
		jsonUpdateMessageInteractionCallbackRequest, _ := json.Marshal(dst.UpdateMessageInteractionCallbackRequest)
		if string(jsonUpdateMessageInteractionCallbackRequest) == "{}" { // empty struct
			dst.UpdateMessageInteractionCallbackRequest = nil
		} else {
			return nil // data stored in dst.UpdateMessageInteractionCallbackRequest, return on the first match
		}
	} else {
		dst.UpdateMessageInteractionCallbackRequest = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateInteractionResponseRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateInteractionResponseRequest) MarshalJSON() ([]byte, error) {
	if src.ApplicationCommandAutocompleteCallbackRequest != nil {
		return json.Marshal(&src.ApplicationCommandAutocompleteCallbackRequest)
	}

	if src.CreateMessageInteractionCallbackRequest != nil {
		return json.Marshal(&src.CreateMessageInteractionCallbackRequest)
	}

	if src.LaunchActivityInteractionCallbackRequest != nil {
		return json.Marshal(&src.LaunchActivityInteractionCallbackRequest)
	}

	if src.ModalInteractionCallbackRequest != nil {
		return json.Marshal(&src.ModalInteractionCallbackRequest)
	}

	if src.PongInteractionCallbackRequest != nil {
		return json.Marshal(&src.PongInteractionCallbackRequest)
	}

	if src.UpdateMessageInteractionCallbackRequest != nil {
		return json.Marshal(&src.UpdateMessageInteractionCallbackRequest)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateInteractionResponseRequest struct {
	value *CreateInteractionResponseRequest
	isSet bool
}

func (v NullableCreateInteractionResponseRequest) Get() *CreateInteractionResponseRequest {
	return v.value
}

func (v *NullableCreateInteractionResponseRequest) Set(val *CreateInteractionResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInteractionResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInteractionResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInteractionResponseRequest(val *CreateInteractionResponseRequest) *NullableCreateInteractionResponseRequest {
	return &NullableCreateInteractionResponseRequest{value: val, isSet: true}
}

func (v NullableCreateInteractionResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInteractionResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


