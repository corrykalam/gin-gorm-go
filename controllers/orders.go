package controllers

import (
	"errors"
	"fmt"
	"pratice-sesi8/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controllers interface {
	CreateOrder(c *gin.Context)
	GetAllOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

type ControllersStruct struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) Controllers {
	return &ControllersStruct{
		DB: db,
	}
}

func (g *ControllersStruct) CreateOrder(c *gin.Context) {
	req := models.AddOrder{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	order := models.Orders{
		CostumerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}
	result := g.DB.Create(&order)
	if result.Error != nil {
		fmt.Println("error found: ", result.Error)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	for _, items := range req.Items {
		item := models.Items{
			ItemCode:    items.ItemCode,
			Description: items.Description,
			Quantity:    items.Quantity,
			OrderID:     order.OrderID,
		}
		err := g.DB.Create(&item).Error
		if err != nil {
			fmt.Println("error found: ", err)
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "succesful to create new data",
	})
}

func (g *ControllersStruct) GetAllOrder(c *gin.Context) {
	result := []models.Orders{}
	err := g.DB.Find(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("error found: ", err)
		c.JSON(404, gin.H{
			"message": "data empty",
		})
		return
	}
	if err != nil {
		fmt.Println("error found: ", err)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func (g *ControllersStruct) UpdateOrder(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	req := models.AddOrder{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	order := models.Orders{
		CostumerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}
	result := g.DB.Model(&models.Orders{}).Where("order_id = ?", id).Updates(&order)
	if result.Error != nil {
		fmt.Println("error found: ", result.Error)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	for _, items := range req.Items {
		item := models.Items{
			ItemID:      items.ItemID,
			ItemCode:    items.ItemCode,
			Description: items.Description,
			Quantity:    items.Quantity,
		}
		fmt.Println(item)
		err := g.DB.Model(&models.Items{}).Where("item_id = ?", items.ItemID).Updates(&item).Error
		if err != nil {
			fmt.Println("error found: ", err)
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "succesful to update data",
	})
}

func (g *ControllersStruct) DeleteOrder(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result := models.Orders{}
	err = g.DB.Where("order_id = ?", id).Delete(&result).Error
	if err != nil {
		fmt.Println("error found: ", err)
		c.JSON(500, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(200, gin.H{
		"message": "succesful to delete data",
	})
}
