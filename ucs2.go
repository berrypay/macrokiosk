/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /ucs2.go
 * Created Date: Saturday March 11th 2023 21:10:21 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Saturday March 11th 2023 21:19:19 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

import (
	"fmt"
)

func convertToUCS2(s string) string {
	var ucs2String string

	for _, c := range s {
		ucs2String += fmt.Sprintf("%04X", c)
	}

	return ucs2String
}
