package main

import (
	"dpw-latihan/database"
	"dpw-latihan/model"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	db := database.InitDB()

	// Middleware to set CORS headers
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(corsMiddleware)

	//Get List User
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Datauser
		var response []byte
		_ = db.Table("datausers").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create User
	r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Datauser
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Delete user
	r.Delete("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract user ID from the request URL
	userID := chi.URLParam(r, "id")
	// Check if the user ID is provided
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}
	// Delete the user with the specified ID
	result := db.Table("datausers").Where("id = ?", userID).Delete(&model.Datauser{})
	if result.Error != nil {
		log.Print("Error deleting user:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
	})

	// Get Detail user based on ID
	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract user ID from the URL parameters
	userID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Datauser
	result := db.Table("datausers").Where("id = ?", userID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Update user
	r.Put("/user/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Datauser
	var data []byte
	var err error
	// Extract user ID from the URL parameters
	userID := chi.URLParam(request, "id")
	// Check if the user with the given ID exists
	var existingUser model.Datauser
	result := db.Table("datausers").Where("id = ?", userID).First(&existingUser)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Banner not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingUser).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update user"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingUser)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})


	//Get List Banner
	r.Get("/banner", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Databanner
		var response []byte
		_ = db.Table("databanners").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create Banner
	r.Post("/banner", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Databanner
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})
	                       

	// Delete banner
	r.Delete("/banner/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract user ID from the request URL
	bannerID := chi.URLParam(r, "id")
	// Check if the user ID is provided
	if bannerID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Banner ID is required"))
		return
	}
	// Delete the user with the specified ID
	result := db.Table("databanners").Where("id = ?", bannerID).Delete(&model.Databanner{})
	if result.Error != nil {
		log.Print("Error deleting banner:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Banner deleted successfully"))
	})

	// Get Detail banner based on ID
	r.Get("/banner/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract banner ID from the URL parameters
	bannerID := chi.URLParam(r, "id")
	// Query the database for the specific banner ID
	var data model.Databanner
	result := db.Table("databanners").Where("id = ?", bannerID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Update Banner
	r.Put("/banner/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Databanner
	var data []byte
	var err error
	// Extract user ID from the URL parameters
	bannerID := chi.URLParam(request, "id")
	// Check if the user with the given ID exists
	var existingBanner model.Databanner
	result := db.Table("databanners").Where("id = ?", bannerID).First(&existingBanner)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Banner not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingBanner).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update user"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingBanner)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List Notification
	r.Get("/notification", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Datanotification
		var response []byte
		_ = db.Table("datanotifications").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create notification
	r.Post("/notification", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Datanotification
	var data []byte
	var err error

	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.Write([]byte("terjadi error diaplikasi kita"))
		return
	}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.Write([]byte("terjadi error diaplikasi kita"))
		return
	}
	_ = db.Create(&payload)
	data, err = json.Marshal(payload)

	writer.Write(data)
	})

	// Get Detail notification based on ID
	r.Get("/notification/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract notification ID from the URL parameters
	notificationID := chi.URLParam(r, "id")
	// Query the database for the specific notification ID
	var data model.Datanotification
	result := db.Table("datanotifications").Where("id = ?", notificationID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete Notification
	r.Delete("/notification/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract user ID from the request URL
	notificationID := chi.URLParam(r, "id")
	// Check if the user ID is provided
	if notificationID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Notification ID is required"))
		return
	}
	// Delete the notification with the specified ID
	result := db.Table("datanotifications").Where("id = ?", notificationID).Delete(&model.Datanotification{})
	if result.Error != nil {
		log.Print("Error deleting notification:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("notification deleted successfully"))
	})

	r.Put("/notification/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Datanotification
	var data []byte
	var err error
	// Extract notification ID from the URL parameters
	notificationID := chi.URLParam(request, "id")
	// Check if the notification with the given ID exists
	var existingNotification model.Datanotification
	result := db.Table("datanotifications").Where("id = ?", notificationID).First(&existingNotification)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Notification not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing notification data with the new payload
	result = db.Model(&existingNotification).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update product"))
		return
	}
	// Return the updated notification data as the response
	data, err = json.Marshal(existingNotification)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List Product
	r.Get("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Dataproduct
		var response []byte
		_ = db.Table("dataproducts").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create Product
	r.Post("/product", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Dataproduct
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Get Detail Product based on ID
	r.Get("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract product ID from the URL parameters
	productID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Dataproduct
	result := db.Table("dataproducts").Where("id = ?", productID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete product
	r.Delete("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract user ID from the request URL
	productID := chi.URLParam(r, "id")
	// Check if the user ID is provided
	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Product ID is required"))
		return
	}
	// Delete the user with the specified ID
	result := db.Table("dataproducts").Where("id = ?", productID).Delete(&model.Dataproduct{})
	if result.Error != nil {
		log.Print("Error deleting product:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
	})

	r.Put("/product/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Dataproduct
	var data []byte
	var err error
	// Extract product ID from the URL parameters
	productID := chi.URLParam(request, "id")
	// Check if the product with the given ID exists
	var existingProduct model.Dataproduct
	result := db.Table("dataproducts").Where("id = ?", productID).First(&existingProduct)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Product not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing Product data with the new payload
	result = db.Model(&existingProduct).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update product"))
		return
	}
	// Return the updated Product data as the response
	data, err = json.Marshal(existingProduct)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List plannogram
	r.Get("/plannogram", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Dataplannogram
		var response []byte
		_ = db.Table("dataplannograms").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create plannogram
	r.Post("/plannogram", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Dataplannogram
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Get Detail plannogram based on ID
	r.Get("/plannogram/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract plannogram ID from the URL parameters
	plannogramID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Dataplannogram
	result := db.Table("dataplannograms").Where("id = ?", plannogramID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete plannogram
	r.Delete("/plannogram/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract plannogram ID from the request URL
	plannogramID := chi.URLParam(r, "id")
	// Check if the plannogram ID is provided
	if plannogramID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Plannogram ID is required"))
		return
	}
	// Delete the plannogram with the specified ID
	result := db.Table("dataplannograms").Where("id = ?", plannogramID).Delete(&model.Dataplannogram{})
	if result.Error != nil {
		log.Print("Error deleting plannogram:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
	})

	r.Put("/plannogram/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Dataplannogram
	var data []byte
	var err error
	// Extract plannogram ID from the URL parameters
	plannogramID := chi.URLParam(request, "id")
	// Check if the plannogram with the given ID exists
	var existingPlannogram model.Dataplannogram
	result := db.Table("dataplannograms").Where("id = ?", plannogramID).First(&existingPlannogram)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Plannogram not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingPlannogram).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update product"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingPlannogram)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List voucher
	r.Get("/voucher", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Datavoucher
		var response []byte
		_ = db.Table("datavouchers").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create voucher
	r.Post("/voucher", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Datavoucher
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Get Detail voucher based on ID
	r.Get("/voucher/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract voucher ID from the URL parameters
	voucherID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Datavoucher
	result := db.Table("datavouchers").Where("id = ?", voucherID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete voucher
	r.Delete("/voucher/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract voucher ID from the request URL
	voucherID := chi.URLParam(r, "id")
	// Check if the voucher ID is provided
	if voucherID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Voucher ID is required"))
		return
	}
	// Delete the voucher with the specified ID
	result := db.Table("datavouchers").Where("id = ?", voucherID).Delete(&model.Datavoucher{})
	if result.Error != nil {
		log.Print("Error deleting voucher:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Voucher deleted successfully"))
	})

	r.Put("/voucher/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Datavoucher
	var data []byte
	var err error
	// Extract voucher ID from the URL parameters
	voucherID := chi.URLParam(request, "id")
	// Check if the voucher with the given ID exists
	var existingVoucher model.Datavoucher
	result := db.Table("datavouchers").Where("id = ?", voucherID).First(&existingVoucher)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Voucher not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingVoucher).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update product"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingVoucher)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List article
	r.Get("/article", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Dataarticle
		var response []byte
		_ = db.Table("dataarticles").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create article
	r.Post("/article", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Dataarticle
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Get Detail article based on ID
	r.Get("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract article ID from the URL parameters
	articleID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Dataarticle
	result := db.Table("dataarticles").Where("id = ?", articleID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete article
	r.Delete("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract article ID from the request URL
	articleID := chi.URLParam(r, "id")
	// Check if the article ID is provided
	if articleID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("article ID is required"))
		return
	}
	// Delete the voucher with the specified ID
	result := db.Table("dataarticles").Where("id = ?", articleID).Delete(&model.Dataarticle{})
	if result.Error != nil {
		log.Print("Error deleting article:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Article deleted successfully"))
	})

	r.Put("/article/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Dataarticle
	var data []byte
	var err error
	// Extract article ID from the URL parameters
	articleID := chi.URLParam(request, "id")
	// Check if the article with the given ID exists
	var existingArticle model.Dataarticle
	result := db.Table("dataarticles").Where("id = ?", articleID).First(&existingArticle)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Article not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingArticle).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update product"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingArticle)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	//Get List deliverymethod
	r.Get("/delivery-method", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []model.Datadeliverymethod
		var response []byte
		_ = db.Table("datadeliverymethods").Find(&data)

		response, _ = json.Marshal(data)

		w.Write(response)
	})

	//Create article
	r.Post("/delivery-method", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var payload model.Datadeliverymethod
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi error diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		writer.Write(data)
	})

	// Get Detail deliverymethod based on ID
	r.Get("/delivery-method/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract deliverymethod ID from the URL parameters
	deliverymethodID := chi.URLParam(r, "id")
	// Query the database for the specific user ID
	var data model.Datadeliverymethod
	result := db.Table("datadeliverymethods").Where("id = ?", deliverymethodID).First(&data)
	if result.Error != nil {
		log.Print("Error querying database:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Print("Error marshalling JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	})

	// Delete deliverymethod
	r.Delete("/delivery-method/{id}", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract deliverymethod ID from the request URL
	deliverymethodID := chi.URLParam(r, "id")
	// Check if the deliverymethod ID is provided
	if deliverymethodID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delivery method ID is required"))
		return
	}
	// Delete the deliverymethod with the specified ID
	result := db.Table("datadeliverymethods").Where("id = ?", deliverymethodID).Delete(&model.Datadeliverymethod{})
	if result.Error != nil {
		log.Print("Error deleting delivery method:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delivery method deleted successfully"))
	})

	r.Put("/delivery-method/{id}", func(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var payload model.Datadeliverymethod
	var data []byte
	var err error
	// Extract deliverymethod ID from the URL parameters
	deliverymethodID := chi.URLParam(request, "id")
	// Check if the deliverymethod with the given ID exists
	var existingDeliveryMethod model.Datadeliverymethod
	result := db.Table("datadeliverymethods").Where("id = ?", deliverymethodID).First(&existingDeliveryMethod)
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Delivery Method not found"))
		return
	}
	// Read and parse the request body
	data, err = io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	// Unmarshal the request body into the payload struct
	err = json.Unmarshal(data, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid request payload"))
		return
	}
	// Update the existing user data with the new payload
	result = db.Model(&existingDeliveryMethod).Updates(payload)
	if result.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to update Delivery Method"))
		return
	}
	// Return the updated user data as the response
	data, err = json.Marshal(existingDeliveryMethod)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Write(data)
	})

	http.ListenAndServe(":3001", r)

}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
