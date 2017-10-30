package ops

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//check if the username has been registered
func check_user(usrname string) bool {
	//read from document 'account'
	fin, err := os.Open("data/account")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(fin)
	flag := true
	for true {
		//read the first line of id
		info, e := buff.ReadString('\n')
		if e != nil {
			//end of the account
			break
		}
		info = strings.Replace(info, "\n", "", -1) //delete \n from the string
		//if the id alredy exist
		if info == usrname {
			fmt.Println("Sorry, the id has been registered, please try another id :)")
			flag = false
			break
		}
		//skip 3 lines of email and mobilephone number
		info, e1 := buff.ReadString('\n')
		if e1 != nil {
			break
		}
		info, e2 := buff.ReadString('\n')
		if e2 != nil {
			break
		}
		info, e3 := buff.ReadString('\n')
		if e3 != nil {
			break
		}
	}
	return flag
}

func check_user_password(name string, password string) bool {
	//read from document 'account'
	fin, err := os.Open("data/account")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(fin)
	flag := false
	for true {
		//read the first line of id
		info, e := buff.ReadString('\n')
		if e != nil { //end of the account
			break
		}
		info = strings.Replace(info, "\n", "", -1) //delete \n from the string
		//if the id alredy exist
		if info == name {
			info, e := buff.ReadString('\n')
			if e != nil {
				break
			}
			info = strings.Replace(info, "\n", "", -1) //delete \n from the string
			if info == password {
				return true
			}
			return false
		}
		//skip 3 lines of email and mobilephone number
		info, e1 := buff.ReadString('\n')
		if e1 != nil {
			break
		}
		info, e2 := buff.ReadString('\n')
		if e2 != nil {
			break
		}
		info, e3 := buff.ReadString('\n')
		if e3 != nil {
			break
		}
	}
	return flag
}
func Create_acount(usrname string, password string, mail string, phone string) {
	//flag = true if usrname has not been used
	flag := check_user(usrname)
	if flag {
		fout, err := os.OpenFile("data/account", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
			fmt.Println("Sorry, can't open the account file, the opration 'register' failed!\n")
		} else {
			io.WriteString(fout, usrname+"\n")
			io.WriteString(fout, password+"\n")
			io.WriteString(fout, mail+"\n")
			io.WriteString(fout, phone+"\n")
			fmt.Println("register success!thanks for your cooperation!")
		}
	}
}

func Login(name string, password string) {
	flag := able_to_login()
	if flag {
		check := check_user_password(name, password)
		if check {
			login_writefile(name, password)
		} else {
			fmt.Printf("your name or password is not correct, please try again\n")
		}
	} else {
		fmt.Printf("you have to log out before log in\n")
	}
}

func able_to_login() bool {
	//read from document 'user_login'
	fin, err := os.Open("data/user_login")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(fin)
	info, e := buff.ReadString('\n')
	info = strings.Replace(info, "\n", "", -1) //delete \n from the string
	if e != nil {
		//no one log in now
		return true
	}
	return false
}
func login_writefile(name, password string) {
	fout, err := os.OpenFile("data/user_login", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		fmt.Println("Sorry, can't open the user_login file, the opration 'login' failed!\n")
	} else {
		io.WriteString(fout, name+"\n")
		io.WriteString(fout, password+"\n")
		fmt.Println("log in success!")
	}
}

func Logout(name, password string) {
	logout_readfile(name, password)
}
func logout_readfile(name, password string) {
	//read from document 'user_login'
	fin, err := os.Open("data/user_login")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(fin)
	info, e := buff.ReadString('\n')
	info = strings.Replace(info, "\n", "", -1) //delete \n from the string
	if e != nil {
		fmt.Println("Sorry, there is no need to log out because no one log in")
	} else {
		pass, e := buff.ReadString('\n')
		if e != nil {
			panic(e)
		}
		pass = strings.Replace(pass, "\n", "", -1) //delete \n from the string
		if info == name && password == pass {
			logout_clearfile()
		} else {
			fmt.Println("Sorry, the name or password is not correct, please try again!")
		}
	}
}
func logout_clearfile() {
	fout, err := os.OpenFile("data/user_login", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
		fmt.Println("Sorry, can't open the user_login file, the opration 'logout' failed!\n")
	} else {
		io.WriteString(fout, "")
		fmt.Println("log out success!")
	}
}
