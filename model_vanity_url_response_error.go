/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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

// VanityURLResponseError - struct for VanityURLResponseError
type VanityURLResponseError struct {
	VanityURLErrorResponse *VanityURLErrorResponse
}

// VanityURLErrorResponseAsVanityURLResponseError is a convenience function that returns VanityURLErrorResponse wrapped in VanityURLResponseError
func VanityURLErrorResponseAsVanityURLResponseError(v *VanityURLErrorResponse) VanityURLResponseError {
	return VanityURLResponseError{
		VanityURLErrorResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VanityURLResponseError) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into VanityURLErrorResponse
	err = newStrictDecoder(data).Decode(&dst.VanityURLErrorResponse)
	if err == nil {
		jsonVanityURLErrorResponse, _ := json.Marshal(dst.VanityURLErrorResponse)
		if string(jsonVanityURLErrorResponse) == "{}" { // empty struct
			dst.VanityURLErrorResponse = nil
		} else {
			if err = validator.Validate(dst.VanityURLErrorResponse); err != nil {
				dst.VanityURLErrorResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.VanityURLErrorResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.VanityURLErrorResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VanityURLResponseError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VanityURLResponseError)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VanityURLResponseError) MarshalJSON() ([]byte, error) {
	if src.VanityURLErrorResponse != nil {
		return json.Marshal(&src.VanityURLErrorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VanityURLResponseError) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.VanityURLErrorResponse != nil {
		return obj.VanityURLErrorResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj VanityURLResponseError) GetActualInstanceValue() (interface{}) {
	if obj.VanityURLErrorResponse != nil {
		return *obj.VanityURLErrorResponse
	}

	// all schemas are nil
	return nil
}

type NullableVanityURLResponseError struct {
	value *VanityURLResponseError
	isSet bool
}

func (v NullableVanityURLResponseError) Get() *VanityURLResponseError {
	return v.value
}

func (v *NullableVanityURLResponseError) Set(val *VanityURLResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableVanityURLResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableVanityURLResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVanityURLResponseError(val *VanityURLResponseError) *NullableVanityURLResponseError {
	return &NullableVanityURLResponseError{value: val, isSet: true}
}

func (v NullableVanityURLResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVanityURLResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

