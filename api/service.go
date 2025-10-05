package api

import (
	"errors"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})


func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	params.SetLocale("en")
	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	if resp.Status == nil {
		return "", errors.New("failed to send OTP")
	}
	return *resp.Sid, nil
}
func (app *Config) twilioVerifyOTP(phoneNumber, code string)(string, error)  {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	if *resp.Status != "approved" {
		return "", errors.New("failed to verify OTP")
	}
	return  *resp.Status, nil
}