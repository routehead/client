package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/routehead/client/pkg/confs"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	apiURL   string
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	versionCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	versionCmd.Flags().StringVarP(&apiURL, "apiURL", "l", "https://api.routehead.com", "API URL")
}

var versionCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the routehead server",
	Run: func(cmd *cobra.Command, args []string) {
		body := map[string]string{"username": username, "password": password}
		b, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		response, err := http.Post(fmt.Sprintf("%s/login", apiURL), "application/json", bytes.NewBuffer(b))
		if err != nil {
			cmd.Println("Could not login")
		}
		defer response.Body.Close()

		var j map[string]string
		err = json.NewDecoder(response.Body).Decode(&j)
		if err != nil {
			panic(err)
		}

		if response.StatusCode == 200 {
			err := confs.CreateUserConfigFile(confs.Config{
				Username: username,
				Token:    j["token"],
			})
			if err != nil {
				fmt.Printf("Token file exists.\nLogout first\n")
				return
			}
			fmt.Printf("Logged in user %s", username)
		} else if response.StatusCode == 401 || response.StatusCode == 403 {
			fmt.Printf("Error in logging in: %s\n", j["reason"])
		}
	},
}
