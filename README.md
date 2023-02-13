# Library-Management-System
First, create a config.ini under the config folder
in the config.ini file, insert the following 


[mysql]
MysqlHost = "your db host"
MysqlPort = "your db port"
MysqlUser = "your db user name"
MysqlPass = "your db password"
MysqlDB = "your database name"


Library Management System Written in Go+Gorm+Mysql+Http+Gorilla Framework 
This system include registering user and authenticated user can perform specific functions such as category CRUD, book CRUD etc. 
User can make a booking on a specific books and the booked books status will change to unavailable after the successful booking. 

