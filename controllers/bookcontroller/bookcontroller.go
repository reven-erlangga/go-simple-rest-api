package bookcontroller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/reven-erlangga/go-simple-rest-api/initializers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	ctx := context.Background()
	var books []models.Book

	result, err := initializers.RDB.Get(ctx, "fetch:books").Result()
	if err == redis.Nil {
		initializers.DB.Find(&books)
		data, _ := json.Marshal(books)

		err := initializers.RDB.Set(ctx, "fetch:books", data, viper.GetDuration("REDIS_CACHE_DURATION")*time.Minute).Err()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	} else {
		err = json.Unmarshal([]byte(result), &books)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"books": &books})

}

func Show(c *gin.Context)  {
	id := c.Param("id")
	var book models.Book

	if err := initializers.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Sorry, book not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func Create(c *gin.Context)  {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// err := c.SaveUploadedFile(book.ImageCover, "assets/books/images/" + book.ImageCover.Filename)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, "unknown error")
	// 	return
	// }

	initializers.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func Update(c *gin.Context)  {
	id := c.Param("id")
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (initializers.DB.Model(&book).Where("id = ?", id).Updates(&book).RowsAffected == 0) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Sorry, your data not update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated!", "book": book})
}

func Delete(c *gin.Context)  {
	var book models.Book

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if initializers.DB.Delete(&book, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Sorry, we cannot delete this item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully to delete!"})
}