package ops

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"Agenda/model"
	"log"
)

type Reg string
type Logtoin string
type Logtoout string
type Delete string
var reg Reg
var logtoin Logtoin
var logtoout Logtoout
var del Delete

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
func Query_user(usrname string,password string) bool{
	if able_to_login()==true {
		return false
	}
	var info_user string
	fin, err := os.Open("data/account")
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(fin)
	for true {
	  info_user = ""
		//read the first line of id
		info, e := buff.ReadString('\n')
		if e != nil {
			//end of the account
			break
		}
		info = strings.Replace(info, "\n", "", -1) //delete \n from the string
		info_user = "username: "+info + " "
		//skip 3 lines of email and mobilephone number
		info, e1 := buff.ReadString('\n')
		if e1 != nil {
			break
		}
		info = strings.Replace(info, "\n", "", -1)
		info, e2 := buff.ReadString('\n')
		if e2 != nil {
			break
		}
		info = strings.Replace(info, "\n", "", -1)
		info_user = info_user + "email_address: " + info + " "
		info, e3 := buff.ReadString('\n')
		if e3 != nil {
			break
		}
		info = strings.Replace(info, "\n", "", -1)
		info_user = info_user + "phone: " + info + " "
		fmt.Println(info_user)
	}
	return true
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
func meeting_title_exist(title string) bool {
	  fin, err := os.Open("data/account_meeting")
	  if err != nil {
		    panic(err)
	  }
	  buff := bufio.NewReader(fin)
	  flag := false
	  for true {
		    info, e := buff.ReadString('\n')
		    if e != nil {
			      break
		    }
			  info, ee := buff.ReadString('\n')
			  if ee != nil {
				    break
			  }
			  info = strings.Replace(info, "\n", "", -1)
			  if info == title {
				    return true
			  }
		    //skip 5 lines
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
				info, e4 := buff.ReadString('\n')
		    if e4 != nil {
			      break
		    }
				info, e5 := buff.ReadString('\n')
		    if e5 != nil {
			      break
		    }
	}
	return flag
}

func Create_acount(usrname string, password string, mail string, phone string) {
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	//flag = true if usrname has not been used
	flag := check_user(usrname)
	if flag {
		fout, err := os.OpenFile("data/account", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
			fmt.Println("Sorry, can't open the data/account file, the opration 'register' failed!\n")
			reg = "[register] 5.Sorry, can't open the data/account file, the opration 'register' failed!\n"
		} else {
			io.WriteString(fout, usrname+"\n")
			reg = "[register] 6.Write the username into data/account file."
			querylog.Println(reg)
			io.WriteString(fout, password+"\n")
			reg = "[register] 7.Write the password into data/account file"
			querylog.Println(reg)
			io.WriteString(fout, mail+"\n")
			reg = "[register] 8.Write the email_address into data/account file"
			querylog.Println(reg)
			io.WriteString(fout, phone+"\n")
			reg = "[register] 9.Write the phone into data/account file"
			querylog.Println(reg)
			fmt.Println("register success!thanks for your cooperation!")
			reg = "[register] 10.register success!thanks for your cooperation!"
			querylog.Println(reg)
		}
	} else {
			reg = "[register] 11.The username has been used!"
			querylog.Println(reg)
	}
}

func Login(name string, password string) {
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	flag := able_to_login()
	if flag {
		logtoin = "[login] 4.It is able to log in."
		querylog.Println(logtoin)
		check := check_user_password(name, password)
		if check {
			login_writefile(name, password)
			logtoin = "[login] 5.You login success!"
			querylog.Println(logtoin)
		} else {
			fmt.Printf("your name or password is not correct, please try again\n")
			logtoin = "[login] 6.Your name or password is not correct, please try again"
			querylog.Println(logtoin)
		}
	} else {
		fmt.Printf("you have to log out before log in\n")
		logtoin = "[login] 7.You should log out to log in another account or you have logged in."
		querylog.Println(logtoin)
	}
}

func able_to_login() bool {

	//read from document 'user_login'
	fin, err := os.Open("data/user_login")
	if err != nil {
		logtoin = "[login] 3.Can not open the data/user_login file."
		del = "[delete] 3.Can not open the data/user_login file."
		panic(err)
	}
	buff := bufio.NewReader(fin)
	info, e := buff.ReadString('\n')
	info = strings.Replace(info, "\n", "", -1) //delete \n from the string
	if e != nil {
		//no one log in now
		//logtoin = "[login] 4.It is able to log in."
		//querylog.Println(logtoin)
		fin.Close()
		return true
	}
	fin.Close()
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
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	//read from document 'user_login'
	fin, err := os.Open("data/user_login")
	if err != nil {
		logtoout = "[logout] 3.Can not open the data/user_login file."
		querylog.Println(logtoout)
		//del = "[delete] 5.Can not open the data/user_login file.\n"
		panic(err)
	}
	buff := bufio.NewReader(fin)
	info, e := buff.ReadString('\n')
	info = strings.Replace(info, "\n", "", -1) //delete \n from the string
	if e != nil {
		fmt.Println("Sorry, there is no need to log out because no one log in")
		logtoout = "[logout] 4.Sorry, there is no need to log out because no one log in"
		querylog.Println(logtoout)
	} else {
		pass, e := buff.ReadString('\n')
		if e != nil {
			logtoout = "[logout] 5.Some error occured when read the bufio"
			querylog.Println(logtoout)
			//del = "[delete] 6.Some error occured when read the bufio.\n"
			panic(e)
		}
		pass = strings.Replace(pass, "\n", "", -1) //delete \n from the string
		if info == name && password == pass {
			logout_clearfile()
		} else {
			fmt.Println("Sorry, the name or password is not correct, please try again!")
			logtoout = "[logout] 6.Sorry, the name or password is not correct, please try again!"
			querylog.Println(logtoout)
			//del = "[delete] 7.Sorry, the name or password is not correct!"
		}
	}
	fin.Close()
}

func logout_clearfile() {
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	fout, err := os.OpenFile("data/user_login", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		logtoout = "[logout] 7.Can not open the data/user_login file\n"
		querylog.Println(logtoout)
		logtoout = "[logout] 8.Logout failed\n"
		querylog.Println(logtoout)
		//del = "[delete] 8.Can not open the data/user_login file\n[logout] 8.Logout failed\n"
		fmt.Println("Sorry, can't open the user_login file, the opration 'logout' failed!\n")
		panic(err)
	} else {
		io.WriteString(fout, "")
		//del = "[delete] 9.Logout sucess!\n"
		logtoout = "[logout] 9.Clear the data/user_login file!"
		querylog.Println(logtoout)
		logtoout = "[logout] 10.Logout sucess!"
		querylog.Println(logtoout)
		fmt.Println("log out success!")
	}
	fout.Close()
}

func able_to_createMeeting(name string, title string, participators []string, startDate string, endDate string) bool {
    var s_date model.Date
    var e_date model.Date
    s_date = model.StringToDate(startDate)
    e_date = model.StringToDate(endDate)

    if ((!s_date.IsValid()) || (!e_date.IsValid())) {
        //fmt.Println(startDate)
				//fmt.Println(endDate)
				fmt.Println("Sorry,the date is invalid!")
        return false
    }

    if (s_date.NotLessThan(e_date)) {
        fmt.Println("Sorry,the startDate can't more than the endDate!")
        return false
    }

    // don't need queryuser to check whether the sponsor is log in

		if (meeting_title_exist(title)) {
			  fmt.Println("Sorry,the meeting title is exist!")
				return false
		}

    for i := 0; i < len(participators); i++ {
        for j := i + 1; j < len(participators); j++ {
            if (participators[i] == participators[j]) {
                fmt.Println("Sorry,the participator can't be the same!")
                return false
            }
        }
    }

    for _, p := range participators {
        if (name == p) {
            fmt.Println("Sorry,the sponsor can't be the participator!")
            return false
        }
    }

    if (len(participators) == 0) {
        fmt.Println("Sorry,there must be at least one participator!")
        return false
    }

    return true
}

func Create_meeting(title string, participators []string, startDate string, endDate string) {
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
			 log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	if able_to_login() == true {
		querylog.Println("[createMeeting] 5.Create meeting failed because you do not log in.")
		fmt.Println("Sorry, please log in")
		} else {
			fin, _ := os.Open("data/user_login")
			buff := bufio.NewReader(fin)
			CurrentUser, _ := buff.ReadString('\n')
			CurrentUser = strings.Replace(CurrentUser, "\n", "", -1) //get current user

			if able_to_createMeeting(CurrentUser, title, participators, startDate, endDate) {
				fout, err := os.OpenFile("data/account_meeting", os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					querylog.Println("[createMeeting] 6.Create meeting failed because can not open the data/account_meeting file.")
					panic(err)
					fmt.Println("Sorry, can't open the account_meeting file, the opration 'createMeeting' failed!\n")
				} else {
					io.WriteString(fout, CurrentUser+"\n")
					io.WriteString(fout, title+"\n")
					for _, people := range participators {
						  io.WriteString(fout, people+" ")
					}
					io.WriteString(fout, "\n")
					io.WriteString(fout, startDate+"\n")
					io.WriteString(fout, endDate+"\n")
					fmt.Println("[createMeeting] 7.Create meeting successfully")
					querylog.Println("[createMeeting] 7.Create meeting sucessfully")
				}
			}
		}
	}

func Query_meeting(startDate string, endDate string) {
    var tmp int = 1
		var s2 model.Date
		var end2 model.Date
		s2 = model.StringToDate(startDate)
		end2 = model.StringToDate(endDate)
		file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
				 log.Fatalln("Fail to create testquery.log file!")
		}
		querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
		if able_to_login() == true {
			  querylog.Println("[QueryMeeting] 3.Query meeting failed because you do not log in.")
		    fmt.Println("Sorry, please log in")
		} else {
			  fin1, _ := os.Open("data/user_login")
			  buff1 := bufio.NewReader(fin1)
			  CurrentUser, _ := buff1.ReadString('\n')
			  CurrentUser = strings.Replace(CurrentUser, "\n", "", -1) //get current user

        fin, err := os.Open("data/account_meeting")
				if err != nil {
					  querylog.Println("[QueryMeeting] 4.Query meeting failed because can not open the data/account_meeting file.")
					  panic(err)
				}
				buff := bufio.NewReader(fin)
				fmt.Println("Sponsor Title Participators StartDate EndDate")
				for true {
					    sponsor, e1 := buff.ReadString('\n')
							if e1 != nil {
								  break
							}
							sponsor = strings.Replace(sponsor, "\n", "", -1)
              // find current user's meeting account
							if sponsor == CurrentUser {
                  // memory title
									title, e2 := buff.ReadString('\n')
									if e2 != nil {
										  break
									}
									title = strings.Replace(title, "\n", "", -1)
                  // get all participators (don't need to check one by one)
									participators, e3 := buff.ReadString('\n')
									if e3 != nil {
										  break
									}
									participators = strings.Replace(participators, "\n", "", -1)

									// get start date
									s_date, e4 := buff.ReadString('\n')
									if e4 != nil {
										  break
									}
									s_date = strings.Replace(s_date, "\n", "", -1)
									var s1 model.Date
							    //var s2 model.Date
							    s1 = model.StringToDate(s_date)
							   // s2 = model.StringToDate(startDate)
									if s1.IsLessThan(s2) == true {
                      tmp = 0
									}

									// get end date
									e_date, e5 := buff.ReadString('\n')
									if e5 != nil {
										  break
									}
									e_date = strings.Replace(e_date, "\n", "", -1)

									var end1 model.Date
							    //var end2 model.Date
							    end1 = model.StringToDate(e_date)
							    //end2 = model.StringToDate(endDate)
									//fmt.Println(end1.Year, end1.Month, end1.Day, end1.Hour, end1.Minute)
									//fmt.Println(end2.Year, end2.Month, end2.Day, end2.Hour, end2.Minute)

									if end1.IsMoreThan(end2) == true {
                      tmp = 0
									}
                  if tmp != 0 {
										  fmt.Println(sponsor, title, participators, s_date, e_date)
									}
							}
				}
				fin2, err2 := os.Open("data/account_meeting")
				if err2 != nil {
						panic(err2)
				}
				buff2 := bufio.NewReader(fin2)
				for true {
              // memory sponsor
							sponsor1, e1 := buff2.ReadString('\n')
							if e1 != nil {
									break
							}
							sponsor1 = strings.Replace(sponsor1, "\n", "", -1)
							// memory title
							title1, e2 := buff2.ReadString('\n')
							if e2 != nil {
									break
							}
							title1 = strings.Replace(title1, "\n", "", -1)
							// get all participators (need to check if current user in them)
							participators1, e3 := buff2.ReadString('\n')
							if e3 != nil {
									break
							}
							participators1 = strings.Replace(participators1, "\n", "", -1)
              tt := strings.Contains(participators1, CurrentUser)
							//fmt.Println(strings.Contains(participators1, CurrentUser))

								  // get start date
								  s_date1, e4 := buff2.ReadString('\n')
								  if e4 != nil {
										  break
								  }
								  s_date1 = strings.Replace(s_date1, "\n", "", -1)

									var s11 model.Date
								  var s21 model.Date
								  s11 = model.StringToDate(s_date1)
								  s21 = model.StringToDate(startDate)
								  if s11.IsLessThan(s21) == true {
										  tmp = 0
										  //fmt.Println("in s")
										  //break
								  }
                  // get end date
								  e_date1, e5 := buff2.ReadString('\n')
								  if e5 != nil {
										  break
								  }
								  e_date1 = strings.Replace(e_date1, "\n", "", -1)

								  var end1 model.Date
								  var end2 model.Date
								  end1 = model.StringToDate(e_date1)
								  end2 = model.StringToDate(endDate)

								  if end1.IsMoreThan(end2) == true {
										  tmp = 0
										  fmt.Println("in e")
										  break
								  }

								  if tmp != 0 && tt == true {
										  fmt.Println(sponsor1, title1, participators1, s_date1, e_date1)
											querylog.Println("[QueryMeeting] 5.Query meeting sucessfully.")
								  }

				}
				if tmp == 0 {
					  fmt.Println("Sorry,there don't have any meeting account!")
				}
		}
}
func Delete_user(name,password string) bool{
	file,err := os.OpenFile("testquery.log",os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to create testquery.log file!")
	}
	querylog := log.New(file,"",log.LstdFlags|log.Llongfile)
	if able_to_login()==true {
		fmt.Println("Please login!")
		del = "[delete] 4.You should log in."
		querylog.Println(del)
		return false
	}
	var userlist []model.User
	querylog.Println("[delete] 14.First to log out")
	logout_readfile(name,password)
	if able_to_login()==false {
		fmt.Println("You can not delete others' account!")
		return false
	}
	fout,err := os.Open("data/account")
	if err != nil {
		del = "[delete] 10.Can not open the data/account file to delete the user."
		querylog.Println(del)
		panic(err)
	}
	buff := bufio.NewReader(fout)
	flag := false
	var count = 0
	var usr model.User
	for true {
		flag = false
		//read the first line of id
		info, e := buff.ReadString('\n')
		if e != nil { //end of the account
			break
		}
		info = strings.Replace(info, "\n", "", -1) //delete \n from the string
		//if the id alredy exist
		if info == name {
			flag = true
		}
		//skip 3 lines of email and mobilephone number
		info1, e1 := buff.ReadString('\n')
		if e1 != nil {
			break
		}
		info1 = strings.Replace(info1, "\n", "", -1)
		info2, e2 := buff.ReadString('\n')
		if e2 != nil {
			break
		}
		info2 = strings.Replace(info2, "\n", "", -1)
		info3, e3 := buff.ReadString('\n')
		if e3 != nil {
			break
		}
		info3 = strings.Replace(info3, "\n", "", -1)
		if flag != true {
			usr.Name = info
			usr.Password = info1
			usr.Email = info2
			usr.Phone = info3
			userlist = append(userlist,usr)
			count++
		} else {
			del = "[delete] 11.Find the information of user which will be deleted from data/account file."
			querylog.Println(del)
		}
  }
	//

	fout1, err := os.OpenFile("data/account", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		del =  "[delete] 12.can not open data/account file to write the updated information."
		querylog.Println(del)
		panic(err)
		fmt.Println("Sorry, can't open the data/account file, the opration 'delete user' failed!\n")
	} else {
		for i:=0;i<count;i++ {
			io.WriteString(fout1, userlist[i].Name+"\n"+userlist[i].Password+"\n"+userlist[i].Email+"\n"+userlist[i].Phone+"\n")
			//fmt.Println(userlist[i].Name,"\n",userlist[i].Password,"\n",userlist[i].Email,"\n"+userlist[i].Phone,"\n")
		}
		del = "[delete] 13.delete user sucess!"
		querylog.Println(del)
		fmt.Println("delete user success!")
	}
	fout1.Close()
	return true
}
func GetReg() Reg{
	return reg
}
func GetLogtoin() Logtoin{
	return logtoin
}
func GetLogtoout() Logtoout{
	return logtoout
}
func GetDelete() Delete{
	return del
}
