package model

type User struct{
    Name string
    Password string
    Email string
    Phone string
}

func (user *User) userInit(t_name string, t_password string, t_email string, t_phone string) {
    user.Name = t_name
    user.Password = t_password
    user.Email = t_email
    user.Phone = t_phone
}

func (user User) GetName() string {
    return user.Name
}

func (user *User) SetName(t_name string) {
    user.Name = t_name
}

func (user User) GetPassword() string {
    return user.Password
}

func (user *User) SetPassword(t_password string) {
    user.Password = t_password
}

func (user User) GetEmail() string {
    return user.Email
}

func (user *User) SetEmail(t_email string) {
    user.Email = t_email
}

func (user User) GetPhone() string {
    return user.Phone
}

func (user *User) SetPhone(t_phone string) {
    user.Phone = t_phone
}
