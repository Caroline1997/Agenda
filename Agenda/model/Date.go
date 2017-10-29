// package

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
        }
        else return true
    }
    if date.Month == 4 || date.Month == 6 || date.Month == 9 || date.Month == 11 {
        if date.Day >= 31 || date.Day <= 0 {
            return false
        }
        else return true
    }
    if ((date.Year % 4 == 0 && date.Year % 100 != 0) || date.Year % 400 == 0) {
        if date.Month == 2 {
            if date.Day >= 30 || date.Day <= 0 {
                return false
            }
            else return true
        }
    }
    else {
        if date.Month == 2 {
            if date.Day >= 29 || date.Day <= 0 {
                return false
            }
            else return true
        }
    }
}

func StringToDate(t_dateString string) Date {
    var temp string
    var t Date
    var isValid bool = 1
    temp = t_dateString
    if temp[4] != '-' || temp[7] != '-' || temp[10] != '/' || temp[13] != ':' {
        isValid = 0
    }
    if temp[0] > '9' || temp[0] < '0' || temp[1] > '9' || temp[1] < '0' || temp[2] > '9' || temp[2] < '0' ||
       temp[3] > '9' || temp[3] < '0' || temp[5] > '9' || temp[5] < '0' || temp[6] > '9' || temp[6] < '0' ||
       temp[8] > '9' || temp[8] < '0' || temp[9] > '9' || temp[9] < '0' || temp[11] > '9' || temp [11] < '0' ||
       temp[12] > '9' || temp[12] < '0' || temp[14] > '9' || temp[14] < '0' || temp[15] > '9' || temp[15] < '0' {
        isValid = 0
    }
    if isValid == 0 {
        t.Year = 0
        t.Month = 0
        t.Day = 0
        t.Hout = 0
        t.Minute = 0
        return t
    } else if isValid == 1 {
        t.Year = (temp[0] - 48) * 1000
        t.Year = t.Year + (temp[1] - 48) * 100
        t.Year = t.Year + (temp[2] - 48) * 10
        t.Year = t.Year + temp[3] - 48

        t.Month = (temp[5] - 48) * 10
        t.Month = t.Month + temp[6] - 48

        t.Day = (temp[8] - 48) * 10
        t.Day = t.Day + temp[9] - 48

        t.Hour = (temp[11] - 48) * 10
        t.Hour = t.Hour + temp[12] - 48

        t.Minute = (temp[14] - 48) * 10
        t.Minute = t.Minute + temp[15] - 48

        return t
    }
}

func DateToString(t_date Date) string {
    var temp string = "0000000000000000"
    if IsValid(t_date) == 0 {
        temp = "0000-00-00/00:00"
        return temp
    } else if IsValid(t_date) == 1 {
        temp[0] = (t_date.Year / 1000) + 48
        temp[1] = ((t_date.Year - 1000 * (temp[0] - 48)) / 100 ) + 48
        temp[2] = ((t_date.Year - 1000 * (temp[0] - 48) - 100 * (temp[1] - 48)) / 10) + 48
        temp[3] = (t_date.Year - 1000 * (temp[0] - 48) - 100 * (temp[1] - 48) - 10 * (temp[2] - 48)) + 48
        temp[4] = '-'
        if t_date.Month < 10 {
            temp[5] = '0'
            temp[6] = t_date.Month + 48
        } else {
            temp[5] = '1'
            temp[6] = t_date.Month - 10 + 48
        }
        temp[7] = '-'
        if t_date.Day < 10 {
            temp[8] = '0'
            temp[9] = t_date.Day + 48
        } else {
            temp[8] = (t_date.Day / 10) + 48
            temp[9] = (t_date.Day - 10 * (temp[8] - 48)) + 48
        }
        temp[10] = '/'
        if t_date.Hour < 10 {
            temp[11] = '0'
            temp[12] = t_date.Hour +48
        } else {
            temp[11] = (t_date.Hour / 10) + 48
            temp[12] = (t_date.Hour - 10 * (temp[11] - 48)) + 48
        }
        temp[13] = ':'
        if t_date.Minute < 10 {
            temp[14] = '0'
            temp[15] = t_date.Minute + 48
        } else {
            temp[14] = (t_date.Minute / 10) + 48
            temp[15] = (t_date.Minute - 10 * (temp[14] - 48)) + 48
        }
        return temp
    }
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
       date.Hour = t_date.Hour &&
       date.Minute == t_date.Minute {
        return true
    } else return false
}

func (date *Date) IsMoreThan(t_date Date) bool {
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
                if date.Hour > t_date.Day {
                    return true
                } else if date.Hour < t_date.Hour {
                    return false
                } else {
                    if date.Minute > t_date.Minute {
                        return true
                    } else return false
                }
            }
        }
    }
}

func (date *Date) IsLessThan(t_date Date) bool {
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
                if date.Hour < t_date.Day {
                    return true
                } else if date.Hour > t_date.Hour {
                    return false
                } else {
                    if date.Minute < t_date.Minute {
                        return true
                    } else return false
                }
            }
        }
    }
}

func (date *Date) NotLessThan(t_date Date) bool {
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
                if date.Hour > t_date.Day {
                    return true
                } else if date.Hour < t_date.Hour {
                    return false
                } else {
                    if date.Minute >= t_date.Minute {
                        return true
                    } else return false
                }
            }
        }
    }
}

func (date *Date) IsMoreThan(t_date Date) bool {
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
                if date.Hour < t_date.Day {
                    return true
                } else if date.Hour > t_date.Hour {
                    return false
                } else {
                    if date.Minute <= t_date.Minute {
                        return true
                    } else return false
                }
            }
        }
    }
}
