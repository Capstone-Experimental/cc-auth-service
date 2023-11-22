package helper

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

const (
	otpChars  = "0123456789"
	otpLength = 6
)

func generateOTP(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	otp := make([]byte, length)
	for i := range otp {
		otp[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(otp)
}

func SendOTP(email string) (string, error) {
	from := "asslinkmiddle@gmail.com"
	password := "Nx9K2h6cAE4t"
	to := []string{email}
	smtpHost := "smtp.zoho.com"
	smtpPort := "587"
	strOtp := generateOTP(6)

	message := []byte("To: " + email + "\r\n" +
		"Subject: OTP for your application\r\n" +
		"\r\n" +
		"Your OTP is: " + strOtp + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return "", err
	}
	fmt.Println("otp : ", strOtp)
	return strOtp, nil
}
