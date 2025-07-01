/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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


// CreateChannelInviteRequest struct for CreateChannelInviteRequest
type CreateChannelInviteRequest struct {
	CreateGroupDMInviteRequest *CreateGroupDMInviteRequest
	CreateGuildInviteRequest *CreateGuildInviteRequest
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateChannelInviteRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CreateGroupDMInviteRequest
	err = json.Unmarshal(data, &dst.CreateGroupDMInviteRequest);
	if err == nil {
		jsonCreateGroupDMInviteRequest, _ := json.Marshal(dst.CreateGroupDMInviteRequest)
		if string(jsonCreateGroupDMInviteRequest) == "{}" { // empty struct
			dst.CreateGroupDMInviteRequest = nil
		} else {
			return nil // data stored in dst.CreateGroupDMInviteRequest, return on the first match
		}
	} else {
		dst.CreateGroupDMInviteRequest = nil
	}

	// try to unmarshal JSON data into CreateGuildInviteRequest
	err = json.Unmarshal(data, &dst.CreateGuildInviteRequest);
	if err == nil {
		jsonCreateGuildInviteRequest, _ := json.Marshal(dst.CreateGuildInviteRequest)
		if string(jsonCreateGuildInviteRequest) == "{}" { // empty struct
			dst.CreateGuildInviteRequest = nil
		} else {
			return nil // data stored in dst.CreateGuildInviteRequest, return on the first match
		}
	} else {
		dst.CreateGuildInviteRequest = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateChannelInviteRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateChannelInviteRequest) MarshalJSON() ([]byte, error) {
	if src.CreateGroupDMInviteRequest != nil {
		return json.Marshal(&src.CreateGroupDMInviteRequest)
	}

	if src.CreateGuildInviteRequest != nil {
		return json.Marshal(&src.CreateGuildInviteRequest)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateChannelInviteRequest struct {
	value *CreateChannelInviteRequest
	isSet bool
}

func (v NullableCreateChannelInviteRequest) Get() *CreateChannelInviteRequest {
	return v.value
}

func (v *NullableCreateChannelInviteRequest) Set(val *CreateChannelInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChannelInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChannelInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChannelInviteRequest(val *CreateChannelInviteRequest) *NullableCreateChannelInviteRequest {
	return &NullableCreateChannelInviteRequest{value: val, isSet: true}
}

func (v NullableCreateChannelInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChannelInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


