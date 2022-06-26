# Department News Board

A service where users can add, view, search, and download news made by <b>Calvert Tanudihardjo (NRP: C14190033)</b> for the Service Oriented Architecture course at Petra Christian University.

Made using **Go programming language** and **MySQL database** with the following main libraries and frameworks:
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [smapping](https://github.com/mashingan/smapping)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)

...and implemented clean architecture, session, also [JWT (JSON Web Tokens)](https://jwt.io/) for authentications.

API Documentation 🚀 [Click here](https://documenter.getpostman.com/view/18705948/UzBsGPoy#538f8522-1e79-4e7d-abbe-cf02f0d62460)

## Services
- Register an account (name, email, and password)
- Login (email and password)
- Logout
- Get all news (+ archive each news if the news has been more than 1 month since it was added)
- Get news by ID
- Download news file
- Upload news (title, author, status, and file) - login and authorization required
- Edit news (title, author, status, and file) - login and authorization required
- Delete news - login and authorization required
