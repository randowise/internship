package controllers

import (
	"example/APIIIDO0/database"
	"example/APIIIDO0/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//LIST
func ShowPros(c *gin.Context) {
	db := database.GetDatabase()

	var pros []models.Pro
	err := db.Find(&pros).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list pro's: " + err.Error(),
		})
		return
	}
	c.JSON(200, pros)
}

func ShowPro(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()

	var pro models.Pro
	err = db.First(&pro, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find pro: " + err.Error(),
		})
		return
	}
	c.JSON(200, pro)
}

//CREATE
func CreatePro(c *gin.Context) {
	db := database.GetDatabase()
	var pro models.Pro

	err := c.ShouldBindJSON(&pro)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&pro).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create pro: " + err.Error(),
		})
		return
	}

	c.JSON(200, pro)

}

//DELETE
// func DeletePro (c *gin.Context){
// 	id:= c.Param("id")
// 	newid, err := strconv.Atoi(id)

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "ID has to be integer",
// 		})
// 		return
// 	}
// 	db := database.GetDatabase()

// 	err = db.Delete(&models.Pro{},newid).Error

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot delete book: " + err.Error(),
// 		})
// 		return
// 	}
// }

//UPDATE
// func EditPro(c *gin.Context) {
// 	db := database.GetDatabase()

// 	var pro []models.Pro

// 	err := c.ShouldBindJSON(&pro)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot bind JSON" + err.Error(),
// 		})
// 		c.JSON(200, pro)
// 	}

// err = db.Save(&pro).Error
// if err != nil{
// 	c.JSON(400, gin.H{
// 		"error": "cannot create Pro: " + err.Error(),

// 	})
// 	return
// }
// c.JSON(200, pro)
// }
