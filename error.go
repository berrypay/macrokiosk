/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /error.go
 * Created Date: Sunday March 12th 2023 00:03:25 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Thursday March 30th 2023 16:11:38 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

import "fmt"

type MkResponseBodyError struct {
	ViolationPart []byte
	Message       string
}

func (m *MkResponseBodyError) Error() string {
	return fmt.Sprintf("Unexpected response body structure found. Violation part: %s, Message: %s", m.ViolationPart, m.Message)
}

type MkError struct {
	Message string
}

func (m *MkError) Error() string {
	return fmt.Sprintf("API Error: %s", m.Message)
}
