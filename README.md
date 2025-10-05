# go-otp-verify

A small OTP (One-Time Password) verification service written in Go using the Gin web framework and Twilio for SMS delivery.

This project provides a lightweight HTTP API to send and validate OTPs (one-time passwords). It's designed as a starter template you can extend for authentication flows, 2FA, or any scenario that requires phone-based verification.

Module: `github.com/PrateekKumar15/go-otp-verify`
Go version: 1.24.3

## Features

- HTTP API built with `github.com/gin-gonic/gin`.
- Input validation via `github.com/go-playground/validator/v10`.
- Environment configuration via `github.com/joho/godotenv`.
- SMS delivery through Twilio (`github.com/twilio/twilio-go`).

## Prerequisites

- Go 1.24+ installed and available on PATH
- A Twilio account (Account SID, Auth Token) and a Twilio phone number capable of sending SMS
- (Optional) Git for cloning and version control

## Environment variables

Create a `.env` file in the project root with the following values (example):

```
# Twilio configuration
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token_here
TWILIO_PHONE_NUMBER=+1234567890

# Server
PORT=8080

# Optional: change to production if needed
GIN_MODE=debug
```

Notes:
- Keep your Twilio credentials secret. Don't commit `.env` to version control.
- Adjust `PORT` if 8080 is already in use on your machine.

## Setup (Windows PowerShell)

Open PowerShell in the project directory and run:

```powershell
# ensure module dependencies are present
go mod tidy

# build the binary
go build -o .\bin\go-otp-verify.exe ./...

# run directly (alternative)
go run .
```

If you see an error when running `go mod init` like "outside GOPATH, module path must be specified", it's because `go mod init` requires an explicit module path when the project is outside GOPATH. This repository already contains a `go.mod` with module path `github.com/PrateekKumar15/go-otp-verify`, so you normally don't need to run `go mod init`.

To initialize your own module name manually (only if you need to recreate `go.mod`):

```powershell
# example: use a VCS-style path if you plan to publish
go mod init github.com/<your-username>/go-otp-verify
go mod tidy
```

## Running the server

With `.env` in place, start the server:

```powershell
# using go run
go run .

# or run the built binary
.\bin\go-otp-verify.exe
```

The server will listen on the port set by `PORT` (default 8080). If `GIN_MODE=release` the Gin logger will be quieter.

## API (assumptions)

The project implements an OTP send + verify flow. The exact endpoints and payloads may vary in the code; common examples are provided below. If your code uses different endpoint names/payloads, adapt accordingly.

- POST /otp/send
  - Body (JSON): { "phone": "+11234567890" }
  - Response: 200 OK on success (OTP sent)

- POST /otp/verify
  - Body (JSON): { "phone": "+11234567890", "otp": "123456" }
  - Response: 200 OK if OTP valid, 400/401 if invalid or expired

Example curl (PowerShell):

```powershell
# send OTP
curl -Method POST -ContentType "application/json" -Body '{"phone":"+11234567890"}' http://localhost:8080/otp/send

# verify OTP
curl -Method POST -ContentType "application/json" -Body '{"phone":"+11234567890","otp":"123456"}' http://localhost:8080/otp/verify
```

If your project exposes different endpoints or requires additional JSON fields, check `main.go` or the router handlers for exact routes and payloads.

## Tests

There are no tests included by default. Add tests under `*_test.go` files and run:

```powershell
go test ./...
```

## Troubleshooting

- "go: cannot determine module path" — specify a module name with `go mod init <module/path>` or use the existing `go.mod` in this repo.
- Twilio errors — confirm `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, and `TWILIO_PHONE_NUMBER` are correct and that your Twilio account has permissions to send SMS to the destination number.
- Port already in use — change `PORT` in `.env` or set an environment variable when starting the server.

## Next steps / Suggestions

- Add persistence for issued OTPs (Redis or in-memory store with expiry) if not already implemented.
- Add rate limiting per phone number to prevent abuse.
- Add unit tests for the OTP generation and verification logic.
- Add CI workflow for build and test.

## License

This repository doesn't include an explicit license. Add a LICENSE file if you plan to share this publicly.

---

If you'd like, I can:
- Open `main.go` and generate an accurate API reference from the actual routes and payloads.
- Add a sample `.env.example` file and a tiny integration test that simulates send/verify (mocking Twilio).

Tell me which you'd like next and I'll proceed.
