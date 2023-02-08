package main

import (
  	"github.com/gin-gonic/gin"
	"restapi/controllers"
	"restapi/models"
	"restapi/middlewares"
)

func main() {

	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/buku-judulbuku",controllers.GetBukuJudulBuku)
	protected.GET("/buku-namapenulis",controllers.GetBukuNamaPenulis)
	protected.GET("/buku-kategori",controllers.GetBukuKategori)

	r.Run(":8080")

}