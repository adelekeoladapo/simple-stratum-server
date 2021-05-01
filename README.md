# Simple Stratum Server

## Requirements
- Golang. <br />
- Docker. <br />
- Docker Compose. <br />

## Instructions
- Clone the project to your device
- From the terminal, navigate to the root directory of the project
- Run <b>go run server.go</b> to start the server on port 3000
- Run <b>go run client.go</b> to run the client
- It can also be tested by using <b>telnet</b> <br />
<b><i>telnet 127.0.0.1 3000</i></b><br />
<b>Sample Requests</b> <br />
{"jsonrpc":"2.0","method":"Mining.Authorize","params": [{"Username":"dapoadeleke","Password":"password123"}],"id":1} <br />
{"jsonrpc":"2.0","method":"Mining.Subscribe","params": [],"id":1}

