Tech Stack
1. DB : mysql
2. GORM => interact with DB
3. JSON Marshall, Unmarshall
4. Project Structure
5. Gorilla mux

CMD => Main.go
PKG :
 - config
    - app.go
 - controllers
    - book-controller
 - models
    - book.go
 - routes
    - bookstore-route
 - utils
    - utils.go
  

routes:
- /book/      : GET     :   getbooks
- /book/      : POST    :   createBook
- /book/{id}  : GET     :   getBook
- /book/{id}  : PUT     :   updateBook
- /book/{id}  : DELETE  :   deleteBook