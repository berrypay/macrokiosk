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

	// ----------- Old Test ---------------
	// var (
	// 	sampleString = "60132899698,26556334019,200,MYR,0.0900|=9.4450,1"
	// )

	// result, err := decodeSingleMTResponse([]byte(sampleString))

	// if result == nil && err != nil {
	// 	t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
	// }

	t.Run("Positive case", func(t *testing.T) {
		const arraySize int = 3
		correctSampleStrings := [arraySize]string{"60132899698,26556334019,200,MYR,0.0900|=9.4450,1", " 60132899698 , 26556334019 , 200, MYR , 0.0900 |= 9.4450,1 ", "60132899698, 26556334019,200,myr,0.0900|=9.4450,1"}

		for i, _ := range correctSampleStrings {
			result, err := decodeSingleMTResponse([]byte(correctSampleStrings[i]))

			if result == nil && err != nil {
				t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
			}
		}
	})
	t.Run("Error case", func(t *testing.T) {
		const arraySize int = 5
		notCorrectSampleStrings := [arraySize]string{"abc456,26556334019,200,MYR,0.0900|=9.4450,1", "60132899698, abc456, 200,MYR,0.0900|=9.4450,1", "60132899698, 26556334019, abc456,MYR,0.0900|=9.4450,1", "60132899698, 26556334019, 200,RANDOM,0.0900|=9.4450,1", "60132899698, 26556334019, 200,MYR,abc456|=9.4450,1"}

		for i, _ := range notCorrectSampleStrings {
			result, err := decodeSingleMTResponse([]byte(notCorrectSampleStrings[i]))

			if result != nil && err == nil {
				t.Errorf("Expected unsuccesfull return, instead of it it was processed, got result: %v", result)
			}
		}
	})
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
