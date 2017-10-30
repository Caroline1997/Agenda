// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Agenda/ops"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new account",
	Long:  `usage of using this command is to register a new account`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		mail, _ := cmd.Flags().GetString("mail")
		phone, _ := cmd.Flags().GetString("phone")
		ops.Create_acount(username, password, mail, phone)
	},
}

/*var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  `usage of using this command is to login`,
	Run: func(cmd *cobra.Command, args []string) {
		name,_:=cmd.Flags().GetString("name")

		fmt.Println(name)
	},
}*/
func init() {
	RootCmd.AddCommand(registerCmd)
	//RootCmd.AddCommand(loginCmd)
	registerCmd.Flags().StringP("user", "u", "default_name", "Username")
	registerCmd.Flags().StringP("password", "p", "default_password", "Password")
	registerCmd.Flags().StringP("mail", "m", "default_mail", "email_address")
	registerCmd.Flags().StringP("phone", "n", "default_phonenumber", "phone_number")
	//loginCmd.Flags().StringP("name","n","","name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
