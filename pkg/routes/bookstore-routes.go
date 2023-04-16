 package routes
 
 import(
	 "github.com/gorilla/mux"
	 "go-bookstore/pkg/controllers"
 )

 var RegisterBookStoreRoutes = func (router *mux.Router)  {
	 router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	 router.HandleFunc("/book/", controllers.GetBook).Method("GET")
	 router.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	 router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Method("PUT")
	 router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Method("DELETE")
	 
 }