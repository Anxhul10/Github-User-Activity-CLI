package main

import (
 "fmt"
 "github.com/spf13/cobra"
 "strings"
)

func runner(cmd *cobra.Command, args []string) {
  var text string
  if len(args) > 0 {
    text = strings.Join(args, " ")
	} else {
    fmt.Println("please pass username!!");
    return
  }
  fmt.Println(text)
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

func main() {
  Execute()
}