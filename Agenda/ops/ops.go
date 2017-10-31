package ops

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"Agenda/model"
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
	if able_to_login() == true {
		fmt.Println("Sorry, please log in")
		} else {
			fin, _ := os.Open("data/user_login")
			buff := bufio.NewReader(fin)
			CurrentUser, _ := buff.ReadString('\n')
			CurrentUser = strings.Replace(CurrentUser, "\n", "", -1) //get current user

			if able_to_createMeeting(CurrentUser, title, participators, startDate, endDate) {
				fout, err := os.OpenFile("data/account_meeting", os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
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
					fmt.Println("Create meeting successfully")
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
		if able_to_login() == true {
		    fmt.Println("Sorry, please log in")
		} else {
			  fin1, _ := os.Open("data/user_login")
			  buff1 := bufio.NewReader(fin1)
			  CurrentUser, _ := buff1.ReadString('\n')
			  CurrentUser = strings.Replace(CurrentUser, "\n", "", -1) //get current user

        fin, err := os.Open("data/account_meeting")
				if err != nil {
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
								  }

				}
				if tmp == 0 {
					  fmt.Println("Sorry,there don't have any meeting account!")
				}
		}
}
