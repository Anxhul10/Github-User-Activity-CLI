package main

import (
 "fmt"
 "io"
 "github.com/spf13/cobra"
 "strings"
 "net/http"
  "encoding/json"
)

type Event struct {
    ID   string `json:"id"`
    Type string `json:"type"`
    Repo struct {
        Name string `json:"name"`
    } `json:"repo"`
    Payload struct {
      Action string `json: "action"`
    } `json: payload`
    CreatedAt string `json:"created_at"`
}

var rootCmd = &cobra.Command{
  Use: "username",
  Short: "gets the username from the CLI",
  Args: cobra.ArbitraryArgs,
  Run:  runner,
}

func Execute() {
 cobra.CheckErr(rootCmd.Execute())
}

func runner(cmd *cobra.Command, args []string) {
  var username string
  if len(args) > 0 {
    username = strings.Join(args, " ")
	} else {
    fmt.Println("please pass username!!");
    return
  }
  logResponse(username)
}

func logResponse(username string) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://api.github.com/users/"+username+"/events/public", nil)

  if err != nil {
    fmt.Println("Request failed: %s", err)
    return
  }
  req.Header.Add("X-GitHub-Api-Version","2022-11-28")

  resp, err := client.Do(req)

  defer resp.Body.Close()

  var events []Event 
  resp_body, err := io.ReadAll(resp.Body)

  json.Unmarshal(resp_body, &events)

  for _, event := range events {
    fmt.Print(event)
  }
}
func main() {
  Execute()
}