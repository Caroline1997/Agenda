# design report for Agenda-golang
----
(commands that has been completed)
## **register**
>$ go run main.go register -u username -p password -m email_address -n phone_number

register a new account if the username has not been registered. if you don't input the infomation it will get its default type "default_XXX"

## **login**
>$ go run main.go login -n username -p password

if someone else log in now you will get the infomation that "you have to log out before log in"
if you haven't registered or the name and password is not correct, you will get the information that "your name or password is not correct, please try again"
if you don't input the infomation it will get its default type "default_XXX"
if you log in, you will get the infomation that "log in success"

## **logout**
>$go run main.go logout -n username -p password

if you didn't login first, you will get the infomation that "Sorry, there is no need to log out because no one log in"
if you haven't registered or the name and password is not correct, you will get the information that "your name or password is not correct, please try again"
if you don't input the infomation it will get its default type "default_XXX"
if you log out, you will get the infomation that "log out success"
