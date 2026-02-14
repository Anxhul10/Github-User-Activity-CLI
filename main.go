package main

import (
 "fmt"
 "io"
 "github.com/spf13/cobra"
 "strings"
 "net/http"
  "encoding/json"
  "strconv"
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

  commit := 0
  repo := ""
  created_repo := ""
  pr_repo := ""
  starred := ""
  fork_name := ""
  _ = fork_name
  _ = created_repo
  _ = pr_repo
  _ = repo
  _ = starred
  visited_cr := false
  visited_starred := false
  visited_pr := false
  visited_fork := false

  limiter := 0
  for _, event := range events {
    if event.Type == "PushEvent" {
      if repo == "" {
        repo = event.Repo.Name
        commit += 1
      } else {
        if repo == event.Repo.Name {
          commit += 1
        }
      }
      if limiter >= 10 {
        break
      }
      limiter += 1
    }
    if event.Payload.Action == "started" && visited_starred == false{
      starred = event.Repo.Name
      visited_starred = true
    }
    if event.Type == "CreateEvent" && visited_cr == false{
      created_repo = event.Repo.Name
      visited_cr = true
    }
    if event.Type == "PullRequestEvent" && visited_pr == false {
      pr_repo = event.Repo.Name
      visited_pr = true
    }
    if event.Type == "ForkEvent" && visited_fork == false {
      fork_name = event.Repo.Name
      visited_fork = true
    }
  }
  
  // logs
  if repo != "" {
    fmt.Println(username + " pushed " + strconv.Itoa(commit) + " commits on "+ repo)
  }
  if created_repo != "" {
    fmt.Println(username + " created " + created_repo + " Repository")
  }
  if pr_repo != "" {
    fmt.Println(username + " openend pull request on " + pr_repo)
  }
  if starred != "" {
    fmt.Println(username +" starred " + starred +" Repository")  
  }
  if fork_name != "" {
    fmt.Println(username +" forked " + fork_name + "Repository")
  }
}
func main() {
  Execute()
}