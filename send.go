/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /send.go
 * Created Date: Saturday March 11th 2023 19:06:05 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday March 12th 2023 04:24:02 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	MkAsciiText    string = "0"
	MkUnicodeText  string = "5"
	MkUDHFormatted string = "6"
)

type MTSendResponseInfo struct {
	MSISDN   string `json:"msisdn" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"msisdn"`
	MsgID    string `json:"msgId" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"msgId"`
	Status   string `json:"status" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"status"`
	Currency string `json:"currency" validate:"string" message:"string: The field {field} must be a string" label:"currency"`
	Price    string `json:"priceBalance" validate:"string" message:"string: The field {field} must be a string" label:"price"`
}

type MTSendResponseSummary struct {
	Balance     string `json:"balance" validate:"string" message:"string: The field {field} must be a string" label:"balance"`
	TotalMSISDN string `json:"totalMsisdn" validate:"string" message:"string: The field {field} must be an string" label:"totalMsisdn"`
}

type MTSendResponseSingleData struct {
	Info    MTSendResponseInfo    `json:"info"`
	Summary MTSendResponseSummary `json:"summary"`
}

type MTSendResponseMultiData struct {
	Infos   []MTSendResponseInfo  `json:"infos"`
	Summary MTSendResponseSummary `json:"summary"`
}

type MTSendRequestParams struct {
	To    string `json:"to" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"to"`
	From  string `json:"from" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"from"`
	Text  string `json:"text" validate:"required|string" message:"required: The field {field} is required|string: The field {field} must be a string" label:"text"`
	Title string `json:"title" validate:"string" message:"string: The field {field} must be a string" label:"title"`
}

func SendSingleMT(textType string, to string, from string, message string, title string) (*MTSendResponseSingleData, error) {
	req, err := http.NewRequest("GET", Settings.BaseUrl+Settings.MTSendPath, nil)
	if err != nil {
		return nil, err
	}

	encodedMessage := message
	if textType == MkAsciiText {
		encodedMessage = url.QueryEscape(message)
	} else if textType == MkUnicodeText {
		encodedMessage = convertToUCS2(message)
	}

	params := map[string]string{
		"user":   GetCredential().User,
		"pass":   GetCredential().Pass,
		"type":   textType,
		"to":     to,
		"from":   from,
		"text":   encodedMessage,
		"servid": "MES01",
		"title":  title,
		"detail": "1",
	}

	// Add the parameters to the URL query string
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("API Url: %s\n", req.URL.String())
	}

	// Send the HTTP request and read the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	callResult, err := decodeSingleMTResponse(body)
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Decoded Response: %v\n", callResult)
	}

	return callResult, nil
}

func SendMultiMT(textType string, to []string, from string, message string, title string) (*MTSendResponseMultiData, error) {
	req, err := http.NewRequest("GET", Settings.BaseUrl+Settings.MTSendPath, nil)
	if err != nil {
		return nil, err
	}

	encodedMessage := message
	if textType == MkAsciiText {
		encodedMessage = url.QueryEscape(message)
	} else if textType == MkUnicodeText {
		encodedMessage = convertToUCS2(message)
	}

	params := map[string]string{
		"user":   GetCredential().User,
		"pass":   GetCredential().Pass,
		"type":   textType,
		"to":     strings.Join(to, ","),
		"from":   from,
		"text":   encodedMessage,
		"servid": "MES01",
		"title":  title,
		"detail": "1",
	}

	// Add the parameters to the URL query string
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("API Url: %s\n", req.URL.String())
	}

	// Send the HTTP request and read the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	callResult, err := decodeMultiMTResponse(body)
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Decoded Response: %v\n", callResult)
	}

	return callResult, nil
}

func decodeSingleMTResponse(body []byte) (*MTSendResponseSingleData, error) {
	bodyStr := strings.Split(string(body), "|=")
	if len(bodyStr) != 2 {
		return nil, &MkResponseBodyError{
			ViolationPart: body,
			Message:       "The response body should contains 2 parts separated by '|='.",
		}
	}

	if !(strings.HasSuffix(bodyStr[1], "1")) {
		return nil, &MkResponseBodyError{
			ViolationPart: []byte(bodyStr[1]),
			Message:       "The TotalMSIDN count should be 1.",
		}
	}

	infoPart, err := getSingleInfoResponse(bodyStr[0])
	if err != nil {
		return nil, err
	}

	summaryPart, err := getSummaryResponse(bodyStr[1])
	if err != nil {
		return nil, err
	}

	return &MTSendResponseSingleData{
		Info:    *infoPart,
		Summary: *summaryPart,
	}, nil
}

func decodeMultiMTResponse(body []byte) (*MTSendResponseMultiData, error) {
	bodyStr := strings.Split(string(body), "|=")
	if len(bodyStr) != 2 {
		return nil, &MkResponseBodyError{
			ViolationPart: body,
			Message:       "The response body should contains 2 parts separated by '|='.",
		}
	}

	if !(strings.HasSuffix(bodyStr[1], "1")) {
		return nil, &MkResponseBodyError{
			ViolationPart: []byte(bodyStr[1]),
			Message:       "The TotalMSIDN count should be 1.",
		}
	}

	infoPart, err := getMultiInfoResponse(bodyStr[0])
	if err != nil {
		return nil, err
	}

	summaryPart, err := getSummaryResponse(bodyStr[1])
	if err != nil {
		return nil, err
	}

	return &MTSendResponseMultiData{
		Infos:   infoPart,
		Summary: *summaryPart,
	}, nil
}

func getSingleInfoResponse(infoBody string) (*MTSendResponseInfo, error) {
	ibStr := strings.Split(infoBody, ",")
	if len(ibStr) != 5 {
		return nil, &MkResponseBodyError{
			ViolationPart: []byte(infoBody),
			Message:       "The info body should contains 5 parts separated by ','.",
		}
	}

	return &MTSendResponseInfo{
		MSISDN:   ibStr[0],
		MsgID:    ibStr[1],
		Status:   ibStr[2],
		Currency: ibStr[3],
		Price:    ibStr[4]}, nil
}

func getMultiInfoResponse(infoBody string) ([]MTSendResponseInfo, error) {
	ibStrArray := strings.Split(infoBody, "|")
	if len(ibStrArray) < 2 {
		return nil, &MkResponseBodyError{
			ViolationPart: []byte(infoBody),
			Message:       "The info body should contains multiple values separated by '|'. Less than 2 is found.",
		}
	}

	retData := []MTSendResponseInfo{}

	for _, str := range ibStrArray {
		ibStr := strings.Split(str, ",")
		if len(ibStr) != 5 {
			return nil, &MkResponseBodyError{
				ViolationPart: []byte(infoBody),
				Message:       "The info body should contains 5 parts separated by ','.",
			}
		}

		retData = append(retData, MTSendResponseInfo{
			MSISDN:   ibStr[0],
			MsgID:    ibStr[1],
			Status:   ibStr[2],
			Currency: ibStr[3],
			Price:    ibStr[4],
		})
	}

	return retData, nil
}

func getSummaryResponse(summaryBody string) (*MTSendResponseSummary, error) {
	sumStr := strings.Split(summaryBody, ",")
	if len(sumStr) != 2 {
		return nil, &MkResponseBodyError{
			ViolationPart: []byte(summaryBody),
			Message:       "The summary body should contains 2 parts separated by ','.",
		}
	}

	return &MTSendResponseSummary{
		Balance:     sumStr[0],
		TotalMSISDN: sumStr[1]}, nil
}
