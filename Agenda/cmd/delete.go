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
	"fmt"
	"Agenda/ops"
	"github.com/spf13/cobra"
	"os"
	"log"
)

// loginCmd represents the login command
var deleteCmd = &cobra.Command{
	Use:   "delete user",
	Short: "for user to delete the account",
	Long: `the usage is to delete his own acccount`,
	Run: func(cmd *cobra.Command, args []string) {
		file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Fail to create testquery.log file!")
		}
		querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
		//fmt.Println("login called")
		name, _ := cmd.Flags().GetString("name")
		querylog.Printf("[delete] 1.get the account's name: "+name+"\n")
		password, _ := cmd.Flags().GetString("password")
		querylog.Printf("[delete] 2.get the account's password: "+password+"\n")
		flag := ops.Delete_user(name, password)
		//var del ops.Delete
		//del = ops.GetDelete()
		if flag==false {
			fmt.Println("Delete failed!")
			del := "[delete] 14.Delete user failed!"
			querylog.Println(del)
		}

	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name","n","default_name","user_name")
	deleteCmd.Flags().StringP("password", "p", "default_password", "password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
