package api

import (
	"context"
	"log"

	"github.com/Gemba-Kaizen/menu-service/internal/models"
	repository "github.com/Gemba-Kaizen/menu-service/internal/repository/menu"
	pb "github.com/Gemba-Kaizen/menu-service/pkg/pb"
	"google.golang.org/grpc/codes"
)

type MenuHandler struct {
	pb.UnimplementedMenuServiceServer
	MenuRepo *repository.MenuRepository
}

func (h *MenuHandler) CreateFoodItem(ctx context.Context, req *pb.CreateFoodRequest) (*pb.CreateFoodResponse, error) {
	// convert proto message to model obj
	foodItem := &models.FoodItem{
		MerchantId:  req.GetFoodItem().GetMerchantId(),
		Name:        req.GetFoodItem().GetName(),
		Price:       req.GetFoodItem().GetPrice(),
		Description: req.GetFoodItem().GetDescription(),
	}

	if err := h.MenuRepo.CreateFoodItem(foodItem); err != nil {
		return &pb.CreateFoodResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	// Return the ID of the newly created record
	return &pb.CreateFoodResponse{
		Status: int64(codes.OK),
	}, nil
}

func (h *MenuHandler) DeleteFoodItem(ctx context.Context, req *pb.DeleteFoodRequest) (*pb.DeleteFoodResponse, error) {
	foodId := req.GetId()

	if err := h.MenuRepo.DeleteFoodItemById(foodId); err != nil {
		return &pb.DeleteFoodResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteFoodResponse{
		Status: 200,
	}, nil
}

func (h *MenuHandler) GetFoodItems(ctx context.Context, req *pb.GetFoodRequest) (*pb.GetFoodResponse, error) {
	merchantId := req.GetMerchantId()

	foodItems, err := h.MenuRepo.GetFoodItemsByMerchantId(merchantId)
	if err != nil{
		return &pb.GetFoodResponse{
      Status: 400,
      Error:  err.Error(),
    }, nil
	}

	// convert food model to pb.FoodItem
	pbFoodItems := make([]*pb.FoodItem, len(foodItems))
	for i, item := range foodItems {
		pbFoodItems[i] = &pb.FoodItem{
			Id:          item.Id,
			MerchantId:  item.MerchantId,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		}
	}

	log.Print(pbFoodItems)
	return &pb.GetFoodResponse{
    Status: 200,
    FoodItems: pbFoodItems,
  }, nil
}

// TODO update endpoint
