/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /send_test.go
 * Created Date: Sunday March 12th 2023 02:00:24 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday March 12th 2023 02:13:38 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

import "testing"

func TestDecodeSingleMTResponse(t *testing.T) {
	var (
		sampleString = "60132899698,26556334019,200,MYR,0.0900|=9.4450,1"
	)

	result, err := decodeSingleMTResponse([]byte(sampleString))

	if result == nil && err != nil {
		t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
	}
}

func TestDecodeMultiMTResponse(t *testing.T) {
	var (
		sampleString = "60132899698,38542873271,200,MYR,0.0450|601116904197,38542873272,200,MYR,0.0900|=9.1300,2"
	)

	result, err := decodeMultiMTResponse([]byte(sampleString))

	if result == nil && err != nil {
		t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
	}
}
