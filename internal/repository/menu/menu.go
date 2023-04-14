package repository

import (
	"github.com/Gemba-Kaizen/menu-service/internal/db"
	"github.com/Gemba-Kaizen/menu-service/internal/models"
)

type MenuRepository struct {
	H *db.Handler
}

func (r *MenuRepository) CreateFoodItem(foodItem *models.FoodItem) error{
	if result := r.H.DB.Create(foodItem); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MenuRepository) DeleteFoodItemById(id int64) error {
	if result := r.H.DB.Delete(&models.FoodItem{}, id); result.Error!= nil {
		return result.Error
	}
	return nil
}

func (r *MenuRepository) GetFoodItemsByMerchantId(merchantId int64) ([]models.FoodItem, error) {
  var foodItems []models.FoodItem
  if result := r.H.DB.Where("merchant_id = ?", merchantId).Find(&foodItems); result.Error != nil {
    return nil, result.Error
  }
  return foodItems, nil
}

func (r *MenuRepository) UpdateFoodItem(foodItem *models.FoodItem) error {
  if result := r.H.DB.Save(foodItem); result.Error != nil {
    return result.Error
  }
  return nil
}