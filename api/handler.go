package api
import (
	"context"
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/PrateekKumar15/go-otp-verify/data"
)

const appTimeout = time.Second * 10
func (app *Config) SendOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()
		app.validateBody(c, &payload) 
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_,err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c,err)
			return
			}
			app.WriteJSON(c,http.StatusAccepted,"OTP Sent Successfully")
		}
}
func (app *Config) VerifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()
		app.validateBody(c, &payload)
		newData := data.VerifyData{
			UserID: payload.UserID,
			Code:   payload.Code,
		}
		fmt.Println(newData.UserID.PhoneNumber)
		resp,err := app.twilioVerifyOTP(newData.UserID.PhoneNumber,newData.Code)
		if err != nil {
			app.errorJSON(c,err)
			return
			}
			if resp == "approved" {
				app.WriteJSON(c,http.StatusAccepted,"User Verified Successfully")
			} else {
				app.WriteJSON(c,http.StatusBadRequest,"Invalid OTP")
			}
		}
}