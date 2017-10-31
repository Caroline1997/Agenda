package model

import (
    "strconv"
    "fmt"
)

type Date struct {
    Year int
    Month int
    Day int
    Hour int
    Minute int
}

func (date *Date) dateInit(t_year int, t_month int, t_day int, t_hour int, t_minute int) {
    date.Year = t_year
    date.Month = t_month
    date.Day = t_day
    date.Hour = t_hour
    date.Minute = t_minute
}

func (date Date) getYear() int {
    return date.Year
}

func (date *Date) setYear(t_year int) {
    date.Year = t_year
}

func (date Date) getMonth() int {
    return date.Month
}

func (date *Date) setMonth(t_month int) {
    date.Month = t_month
}

func (date Date) getDay() int {
    return date.Day
}

func (date *Date) setDay(t_day int) {
    date.Day = t_day
}

func (date Date) getHour() int {
    return date.Hour
}

func (date *Date) setHour(t_hour int) {
    date.Hour = t_hour
}

func (date Date) getMinute() int {
    return date.Minute
}

func (date *Date) setMinute(t_minute int) {
    date.Minute = t_minute
}

func (date Date) IsValid() bool {
    if date.Year < 1000 || date.Year > 9999 {
        return false
    }
    if date.Month >= 13 || date.Month < 1 {
        return false
    }
    if date.Hour >= 24 || date.Hour < 0 {
        return false
    }
    if date.Minute >= 60 || date.Minute < 0 {
        return false
    }
    if date.Month == 1 || date.Month == 3 || date.Month == 5 || date.Month == 7 || date.Month == 8 || date.Month == 10 || date.Month == 12 {
        if date.Day >= 32 || date.Day <= 0 {
            return false
        } else {
            return true
        }
    }
    if date.Month == 4 || date.Month == 6 || date.Month == 9 || date.Month == 11 {
        if date.Day >= 31 || date.Day <= 0 {
            return false
        } else {
            return true
        }
    }
    if ((date.Year % 4 == 0 && date.Year % 100 != 0) || date.Year % 400 == 0) {
        if date.Month == 2 {
            if date.Day >= 30 || date.Day <= 0 {
                return false
            } else {
                return true
            }
        }
    }else {
        if date.Month == 2 {
            if date.Day >= 29 || date.Day <= 0 {
                return false
            } else {
                return true
            }
        }
    }
    return true
}

func StringToInt(s string) int {
    result, error := strconv.Atoi(s)
    if error != nil {
        fmt.Println("fail")
    }
    return result
}

func StringToDate(t_dateString string) Date {
    var t Date
    var isValid bool = true
    if t_dateString[4] != '-' || t_dateString[7] != '-' || t_dateString[10] != '/' || t_dateString[13] != ':' || len(t_dateString) != 16 {
        isValid = false
    }
    if t_dateString[0] > '9' || t_dateString[0] < '0' || t_dateString[1] > '9' || t_dateString[1] < '0' || t_dateString[2] > '9' || t_dateString[2] < '0' ||
       t_dateString[3] > '9' || t_dateString[3] < '0' || t_dateString[5] > '9' || t_dateString[5] < '0' || t_dateString[6] > '9' || t_dateString[6] < '0' ||
       t_dateString[8] > '9' ||t_dateString[8] < '0' || t_dateString[9] > '9' || t_dateString[9] < '0' || t_dateString[11] > '9' || t_dateString[11] < '0' ||
       t_dateString[12] > '9' || t_dateString[12] < '0' || t_dateString[14] > '9' || t_dateString[14] < '0' || t_dateString[15] > '9' || t_dateString[15] < '0' {
        isValid = false
    }
    if isValid == false {
        t.Year = 0
        t.Month = 0
        t.Day = 0
        t.Hour = 0
        t.Minute = 0
        return t
    } else if isValid == true {
      t.setYear(StringToInt(t_dateString[0:4]))
      t.setMonth(StringToInt(t_dateString[5:7]))
      t.setDay(StringToInt(t_dateString[8:10]))
      t.setHour(StringToInt(t_dateString[11:13]))
      t.setMinute(StringToInt(t_dateString[14:16]))
      return t
    }
    return t
}

func IntToString(a int) string {
    var r_string string
    r_string = strconv.Itoa(a)
    return r_string
}

func DateToString(t_date Date) string {
    var temp string = ""
    if t_date.IsValid() == false {
        return "0000-00-00/00:00"
    } else if t_date.IsValid() == true {
        temp += IntToString(t_date.Year) + "-"
        if t_date.Month < 10 {
            temp += "0"
        }
        temp += IntToString(t_date.Month) + "-"
        if t_date.Day < 10 {
            temp += "0"
        }
        temp += IntToString(t_date.Day) + "/"
        if t_date.Hour < 10 {
            temp += "0"
        }
        temp += IntToString(t_date.Hour) + ":"
        if t_date.Minute < 10 {
            temp += "0"
        }
        temp += IntToString(t_date.Minute)
        return temp
    }
    return temp
}

func (date *Date) AssignData(t_date Date) {
    date.Year = t_date.Year
    date.Month = t_date.Month
    date.Day = t_date.Day
    date.Hour = t_date.Hour
    date.Minute = t_date.Minute
}

func (date *Date) IsEqual(t_date Date) bool {
    if date.Year == t_date.Year &&
       date.Month == t_date.Month &&
       date.Day == t_date.Day &&
       date.Hour == t_date.Hour &&
       date.Minute == t_date.Minute {
        return true
    } else {
        return false
    }
}

func (date Date) IsMoreThan(t_date Date) bool {
    if date.Year > t_date.Year {
        return true
    } else if date.Year < t_date.Year {
        return false
    } else {
        if date.Month > t_date.Month {
            return true
        } else if date.Month < t_date.Month {
            return false
        } else {
            if date.Day > t_date.Day {
                return true
            } else if date.Day < t_date.Day {
                return false
            } else {
                if date.Hour > t_date.Hour {
                    return true
                } else if date.Hour < t_date.Hour {
                    return false
                } else {
                    if date.Minute > t_date.Minute {
                        return true
                    } else {
                        return false
                    }
                }
            }
        }
    }
}

func (date Date) IsLessThan(t_date Date) bool {
    if date.Year < t_date.Year {
        return true
    } else if date.Year > t_date.Year {
        return false
    } else {
        if date.Month < t_date.Month {
            return true
        } else if date.Month > t_date.Month {
            return false
        } else {
            if date.Day < t_date.Day {
                return true
            } else if date.Day > t_date.Day {
                return false
            } else {
                if date.Hour < t_date.Hour {
                    return true
                } else if date.Hour > t_date.Hour {
                    return false
                } else {
                    if date.Minute < t_date.Minute {
                        return true
                    } else {
                        return false
                    }
                }
            }
        }
    }
}

func (date Date) NotLessThan(t_date Date) bool {
    if date.Year > t_date.Year {
        return true
    } else if date.Year < t_date.Year {
        return false
    } else {
        if date.Month > t_date.Month {
            return true
        } else if date.Month < t_date.Month {
            return false
        } else {
            if date.Day > t_date.Day {
                return true
            } else if date.Day < t_date.Day {
                return false
            } else {
                if date.Hour > t_date.Hour {
                    return true
                } else if date.Hour < t_date.Hour {
                    return false
                } else {
                    if date.Minute >= t_date.Minute {
                        return true
                    } else {
                        return false
                    }
                }
            }
        }
    }
}

func (date Date) NotMoreThan(t_date Date) bool {
    if date.Year < t_date.Year {
        return true
    } else if date.Year > t_date.Year {
        return false
    } else {
        if date.Month < t_date.Month {
            return true
        } else if date.Month > t_date.Month {
            return false
        } else {
            if date.Day < t_date.Day {
                return true
            } else if date.Day > t_date.Day {
                return false
            } else {
                if date.Hour < t_date.Hour {
                    return true
                } else if date.Hour > t_date.Hour {
                    return false
                } else {
                    if date.Minute <= t_date.Minute {
                        return true
                    } else {
                        return false
                    }
                }
            }
        }
    }
}
