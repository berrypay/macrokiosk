/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /error.go
 * Created Date: Sunday March 12th 2023 00:03:25 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday March 12th 2023 01:20:16 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

type MkResponseBodyError struct {
	ViolationPart []byte
	Message       string
}

func (m *MkResponseBodyError) Error() string {
	return "Unexpected response body structure found"
}
