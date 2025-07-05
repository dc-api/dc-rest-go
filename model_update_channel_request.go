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
)


// UpdateChannelRequest struct for UpdateChannelRequest
type UpdateChannelRequest struct {
	UpdateDMRequestPartial *UpdateDMRequestPartial
	UpdateGroupDMRequestPartial *UpdateGroupDMRequestPartial
	UpdateGuildChannelRequestPartial *UpdateGuildChannelRequestPartial
	UpdateThreadRequestPartial *UpdateThreadRequestPartial
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpdateChannelRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpdateDMRequestPartial
	err = json.Unmarshal(data, &dst.UpdateDMRequestPartial);
	if err == nil {
		jsonUpdateDMRequestPartial, _ := json.Marshal(dst.UpdateDMRequestPartial)
		if string(jsonUpdateDMRequestPartial) == "{}" { // empty struct
			dst.UpdateDMRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateDMRequestPartial, return on the first match
		}
	} else {
		dst.UpdateDMRequestPartial = nil
	}

	// try to unmarshal JSON data into UpdateGroupDMRequestPartial
	err = json.Unmarshal(data, &dst.UpdateGroupDMRequestPartial);
	if err == nil {
		jsonUpdateGroupDMRequestPartial, _ := json.Marshal(dst.UpdateGroupDMRequestPartial)
		if string(jsonUpdateGroupDMRequestPartial) == "{}" { // empty struct
			dst.UpdateGroupDMRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateGroupDMRequestPartial, return on the first match
		}
	} else {
		dst.UpdateGroupDMRequestPartial = nil
	}

	// try to unmarshal JSON data into UpdateGuildChannelRequestPartial
	err = json.Unmarshal(data, &dst.UpdateGuildChannelRequestPartial);
	if err == nil {
		jsonUpdateGuildChannelRequestPartial, _ := json.Marshal(dst.UpdateGuildChannelRequestPartial)
		if string(jsonUpdateGuildChannelRequestPartial) == "{}" { // empty struct
			dst.UpdateGuildChannelRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateGuildChannelRequestPartial, return on the first match
		}
	} else {
		dst.UpdateGuildChannelRequestPartial = nil
	}

	// try to unmarshal JSON data into UpdateThreadRequestPartial
	err = json.Unmarshal(data, &dst.UpdateThreadRequestPartial);
	if err == nil {
		jsonUpdateThreadRequestPartial, _ := json.Marshal(dst.UpdateThreadRequestPartial)
		if string(jsonUpdateThreadRequestPartial) == "{}" { // empty struct
			dst.UpdateThreadRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateThreadRequestPartial, return on the first match
		}
	} else {
		dst.UpdateThreadRequestPartial = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpdateChannelRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateChannelRequest) MarshalJSON() ([]byte, error) {
	if src.UpdateDMRequestPartial != nil {
		return json.Marshal(&src.UpdateDMRequestPartial)
	}

	if src.UpdateGroupDMRequestPartial != nil {
		return json.Marshal(&src.UpdateGroupDMRequestPartial)
	}

	if src.UpdateGuildChannelRequestPartial != nil {
		return json.Marshal(&src.UpdateGuildChannelRequestPartial)
	}

	if src.UpdateThreadRequestPartial != nil {
		return json.Marshal(&src.UpdateThreadRequestPartial)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUpdateChannelRequest struct {
	value *UpdateChannelRequest
	isSet bool
}

func (v NullableUpdateChannelRequest) Get() *UpdateChannelRequest {
	return v.value
}

func (v *NullableUpdateChannelRequest) Set(val *UpdateChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChannelRequest(val *UpdateChannelRequest) *NullableUpdateChannelRequest {
	return &NullableUpdateChannelRequest{value: val, isSet: true}
}

func (v NullableUpdateChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


