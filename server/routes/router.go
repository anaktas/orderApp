package routes

import (
	"net/http"
	"orderApp/server/controllers"
	"orderApp/server/controllers/admin"
	"orderApp/server/controllers/bar"
	"orderApp/server/controllers/cashier"
	"orderApp/server/controllers/waiter"
)

func Routes() {
	http.HandleFunc("/", controllers.ErrorNotFound)
	// Admin routes
	http.HandleFunc("/admin", admin.Index)
	// Bar routes
	http.HandleFunc("/bar", bar.Index)
	// Cashier routes
	http.HandleFunc("/cashier", cashier.Index)
	// Waiter routes
	http.HandleFunc("/waiter", waiter.Index)
}
