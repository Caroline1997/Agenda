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

func (user User) getName() string {
    return user.Name
}

func (user *User) setName(t_name string) {
    user.Name = t_name
}

func (user User) getPassword() string {
    return user.Password
}

func (user *User) setPassword(t_password string) {
    user.Password = t_password
}

func (user User) getEmail() string {
    return user.Email
}

func (user *User) setEmail(t_email string) {
    user.Email = t_email
}

func (user User) getPhone() string {
    return user.Phone
}

func (user *User) setPhone(t_phone string) {
    user.Phone = t_phone
}
