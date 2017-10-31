# design report for Agenda-golang
----
(commands that has been completed)
## **register**
>$ go run main.go register -u username -p password -m email_address -n phone_number

register a new account if the username has not been registered. if you don't input the information it will get its default type "default_XXX".

## **login**
>$ go run main.go login -n username -p password

if someone else log in now you will get the information that "you have to log out before log in".
if you haven't registered or the name and password is not correct, you will get the information that "your name or password is not correct, please try again".
if you don't input the information it will get its default type "default_XXX".
if you log in, you will get the information that "log in success".

## **logout**
>$go run main.go logout -n username -p password

if you didn't login first, you will get the information that "Sorry, there is no need to log out because no one log in".
if you haven't registered or the name and password is not correct, you will get the information that "your name or password is not correct, please try again".
if you don't input the information it will get its default type "default_XXX".
if you log out, you will get the information that "log out success".

## **createMeeting**
>$go run main.go createMeeting -t title -p participators -s startDate -e endDate

if you didn't log in,you will get an information that "Sorry,please log in".
if your startDate or endDate invalid,you will get an information that "Sorry,the date is invalid".
if your startDate not less than endDate,you will get an information that "Sorry,the startDate can't more than the endDate".
if your title is exist,you will get an information that "Sorry,the meeting title is exist".
if the participator list has two same participator,you will get an information that "Sorry,the participator can't be the same".
if the participator list include yourself,you will get an information that "Sorry,the sponsor can't be the participator".
if you pass all the check, you will get an information that "Create meeting successfully".

## **queryMeeting**
>$go run main.go queryMeeting -s startDate -e endDate

if you didn't log in,you will get an information that "Sorry,please log in".
if there didn't have any meeting satisfied the date,you will get an information that "Sorry,there don't have any meeting account".
if there exist a meeting that the current user as sponsor or as paticipator,this meeting will be print.
