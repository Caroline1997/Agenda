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
	"github.com/spf13/cobra"
	"io"
	"bufio"
	"fmt"
	//"io/ioutil"
	"os"
	"strings"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new account",
	Long:  `usage of using this command is to register a new account`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		create_acount(username)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func create_acount(usrname string) {	
	//read from document 'account'
	fin,err := os.Open("data/account")
	if err != nil{
		panic(err)
	}
	buff := bufio.NewReader(fin)
	//check if the name has been registered
	flag := true
	for true {
		//read the first line of id
		info, e := buff.ReadString('\n')
		if e != nil{
			//end of the account		
			break
		}
		info = strings.Replace(info,"\n","", -1) //delete \n from the string
		//if the id alredy exist		
		if info == usrname{
			fmt.Println("Sorry, the id has been registered, please try another id :)")
			flag = false
			break		
		}
		//skip two lines of email and mobilephone number
		info, e1 := buff.ReadString('\n')
		if e1 != nil{		
			break
		}
		info, e2 := buff.ReadString('\n')
		if e2 != nil{		
			break
		}
	}
	if flag{
		fout, err := os.OpenFile("data/account", os.O_WRONLY|os.O_APPEND, 0666)
		if err!= nil {
			panic(err)
			fmt.Println("Sorry, the opration 'register' failed!\n")
		}else{
			var email, number string
			fmt.Println("Please write down your email-address and your mobile number:")
			fmt.Scanln(&email, &number)
			io.WriteString(fout, usrname+"\n")			
			io.WriteString(fout, email+"\n")
			io.WriteString(fout, number + "\n")
			fmt.Println("register success!thanks for your cooperation!")
		}
	}
}
