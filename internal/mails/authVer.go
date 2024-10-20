package mails

import (
	"math/rand"
	"strconv"
	"time"
)

type Email interface {
	SendEmail(string, string)
	VerifyEmail(string) bool
	generateCode() string
}

type Mail struct {
	Code string
}

func (m *Mail) TransactionMail(from, receiver string, code int, amount float64) string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Transaction Verification</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			margin: 0;
			padding: 0;
		}
	.email-container {
		background-color: #ffffff;
		max-width: 600px;
		margin: 0 auto;
		padding: 20px;
		border: 1px solid #ddd;
	}
	.header {
		background-color: #4CAF50;
		color: white;
		padding: 10px;
		text-align: center;
	}
	.content {
		padding: 20px;
		color: #333;
	}
	.transaction-details {
		margin-bottom: 20px;
	}
	.transaction-details p {
		line-height: 1.6;
	}
	.verify-btn {
		display: inline-block;
		padding: 10px 20px;
		background-color: #4CAF50;
		color: white;
		text-decoration: none;
		border-radius: 5px;
	}
	.footer {
		text-align: center;
		padding: 10px;
		color: #999;
		font-size: 12px;
	}
	</style>
	</head>
	<body>
	<div class="email-container">
	<div class="header">
	<h1>Transaction Verification</h1>
	</div>
	<div class="content">
	<p>We have received a request to verify a transaction on your account.</p>

	<div class="transaction-details">
	<h2>Transaction Details:</h2>
	<p><strong>Amount:</strong> $` + strconv.Itoa(int(amount)) + `</p>
	<p><strong>Date:</strong>` + time.Now().GoString() + `</p>
	<p><strong>Transaction ID:</strong> TXN-` + strconv.Itoa(rand.Intn(1000000)) + `</p>
	<p><strong>From:</strong> $` + from + `</p>
	<p><strong>Receiver:</strong> $` + receiver + `</p>
	</div>

	<p>If this was you, copy and please paste verification code:</p>
	<p>` + strconv.Itoa(code) + `</p> 
	</div>
	<div class="footer">
	</div>
	</div>
	</body>
	</html>`

}

func (m *Mail) AuthMail(s string) string {
	return "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n" +
		"<head>\n" +
		"    <meta charset=\"UTF-8\">\n" +
		"    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n" +
		"    <title>Email Verification Code</title>\n" +
		"    <style>\n" +
		"        body {\n" +
		"            font-family: Arial, sans-serif;\n" +
		"            color: #333;\n" +
		"            line-height: 1.6;\n" +
		"            margin: 20px;\n" +
		"        }\n" +
		"        .container {\n" +
		"            max-width: 600px;\n" +
		"            margin: auto;\n" +
		"            padding: 20px;\n" +
		"            border: 1px solid #ddd;\n" +
		"            border-radius: 8px;\n" +
		"            background-color: #f9f9f9;\n" +
		"        }\n" +
		"        .code {\n" +
		"            font-size: 18px;\n" +
		"            font-weight: bold;\n" +
		"            color: #007bff;\n" +
		"            margin: 20px 0;\n" +
		"        }\n" +
		"        .note {\n" +
		"            font-size: 14px;\n" +
		"            color: #666;\n" +
		"        }\n" +
		"    </style>\n" +
		"</head>\n" +
		"<body>\n" +
		"    <div class=\"container\">\n" +
		"        <p>Your Email Verification Code is:</p>\n" +
		"        <p class=\"code\">" + s + "</p>\n" +
		"        <p class=\"note\">Do not share this code with anyone.</p>\n" +
		"    </div>\n" +
		"</body>\n" +
		"</html>\n"
}
