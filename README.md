# tcare-app
Simple app to make restful api calls. Below are instruction for how to stand up the server and send api requests to it. 2 ways to bring up the server are detailed
1. For a machine with the `Make` command installed
2. For a machine without the `Make` command installed

Both methods assume you have cloned the github repository and navigated to it in your terminal

## With Make installed
1. Run the command `make gomod-tidy` to update dependencies
2. Run command `make standup` to stand up the server

## Without Make installed
1. Run the commands `go mod tidy`, `go mod download`, `go mod vendor`
2. Run the command `cd server` to navigate into the main folder
3. Run the command `go run main.go` to stand up the server

You should see the following message `server is running at localhost:1234` and be able to make api calls to the server
