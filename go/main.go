package main

import (
	// "Project-Desain-Pemrograman-Web-main/model"

	"encoding/json"
	"io"
	"latihan-golang/database"
	"latihan-golang/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func main() {

	r := chi.NewRouter()
	db := database.InitDB()

	r.Get("/article/add.vue", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "aplication/json")

		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.File("add.vue")
		})
		r.GET("/image", func(c *gin.Context) {
			if currentImage == nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			currentImage.Write(c.Writer)
		})
		r.GET("/thumbnail", func(c *gin.Context) {
			if currentImage == nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)
			if err != nil {
				panic(err)
			}
			t.Write(c.Writer)
		})
		r.POST("/upload", func(c *gin.Context) {
			img, err := imageupload.Process(c.Request, "file")
			if err != nil {
				panic(err)
			}
			currentImage = img
			c.Redirect(http.StatusMovedPermanently, "/")
		})

		_ = chi.URLParam(request, "id")
		var data model.Article

		result := db.Table("Article").Where("Article_name= ?", data.Article_name).First(&data)

		if result.Error != nil {
			log.Print("Error querying database:", result.Error)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(data)
		if err != nil {
			log.Print("Error querying database:", result.Error)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Write(response)
	})
	r.Delete("/list-Article/{id}", func(w http.ResponseWriter, r *http.Request) {
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
		result := db.Table("datausers").Where("id = ?", userID).Delete(&model.Article{})
		if result.Error != nil {
			log.Print("Error deleting user:", result.Error)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted successfully"))
	})
	r.Post("/Article", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "aplication/json")

		var payload model.Article
		var data []byte
		var err error

		data, err = io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("terjadi eror diaplikasi kita"))
			return
		}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			writer.Write([]byte("terjadi eror diaplikasi kita"))
			return
		}
		_ = db.Create(&payload)
		data, err = json.Marshal(payload)

		// db.Table("users").Create(&payload)
		writer.WriteHeader(201)
		writer.Write(data)
		if err != nil {
			writer.Write([]byte("terjadi eror diaplikasi kita"))
			return
		}
		// if err != nil {
		// 	writer.Write([]byte("terjadi eror diaplikasi kita"))
		// 	return
		// }
		// _ = db.Create(&payload)
		// data, err = json.Marshal(payload)

		// writer.Write([]byte("ini dari handler post"))
	})
	http.ListenAndServe(":8081", r)
}
