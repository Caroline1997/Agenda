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
	"log"
	"os"
	//"io"
)

// loginCmd represents the login command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "for user to query",
	Long: `the usage is to query`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Fail to create testquery.log file!")
		}
		querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
		name, _ := cmd.Flags().GetString("name")
		querylog.Printf("[query] 1.get the account's name: "+name+"\n")
		//io.WriteString(file"[query] 1.get the account's name: "+name+"\n")
		password, _ := cmd.Flags().GetString("password")
		querylog.Printf("[query] 1.get the account's password: "+password+"\n")
		//io.WriteString(file,"[query] 2.get the account's password: "+password+"\n")
		flag :=ops.Query_user(name,password)
		if flag==true {
			fmt.Println("Query success!")
			querylog.Printf("[query] 3.get the account's name:Query success! \n")
			//io.WriteString(file,"[query] 3.get the account's name:Query success! \n")
		} else {
		  querylog.Printf("[query] 4.sorry,you should login!\n")
			//io.WriteString(file,"[query] 4.sorry,you should login!")
			fmt.Println("sorry,you should login!")
			file.Close()
	}
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringP("name","n","default_name","user_name")
	queryCmd.Flags().StringP("password", "p", "default_password", "password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
