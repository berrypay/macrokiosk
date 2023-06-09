/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /setting.go
 * Created Date: Saturday March 11th 2023 21:04:51 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Monday March 13th 2023 08:52:42 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package macrokiosk

const (
	MkAsciiText    string = "0"
	MkUnicodeText  string = "5"
	MkUDHFormatted string = "6"
)

type MacroKioskCredential struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type MacroKioskSettings struct {
	BaseUrl                string                `json:"baseUrl"`
	Credential             *MacroKioskCredential `json:"credential"`
	DefaultFrom            string                `json:"defaultFrom"`
	DefaultMessageEncoding string                `json:"defaultMessageEncoding"`
	MTSendPath             string                `json:"mtSendPath"`
}

var Settings *MacroKioskSettings

func init() {
	Settings = &MacroKioskSettings{
		BaseUrl: "https://www.etracker.cc",
		Credential: &MacroKioskCredential{
			User: "TEST000",
			Pass: "",
		},
		DefaultFrom:            "Private Sender",
		DefaultMessageEncoding: MkAsciiText,
		MTSendPath:             "/bulksms/mesapi.aspx?",
	}
}

func SetCredential(credential *MacroKioskCredential) {
	Settings.Credential = credential
}

func GetCredential() *MacroKioskCredential {
	return Settings.Credential
}

func SetDefaultFrom(from string) {
	Settings.DefaultFrom = from
}

func GetDefaultFrom() string {
	return Settings.DefaultFrom
}

func SetBaseUrl(baseUrl string) {
	Settings.BaseUrl = baseUrl
}

func GetBaseUrl() string {
	return Settings.BaseUrl
}

func SetMTSendPath(path string) {
	Settings.MTSendPath = path
}

func GetMTSendPath() string {
	return Settings.MTSendPath
}

func GetDefaultMessageEncoding() string {
	return Settings.DefaultMessageEncoding
}

func SetDefaultMessageEncoding(messageEncoding string) {
	Settings.DefaultMessageEncoding = messageEncoding
}
