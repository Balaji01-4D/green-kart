package shop

import (
	"shop-near-u/internal/models"
	"shop-near-u/internal/utils"

	"github.com/restayway/gogis"
)

type Service struct {
	repository *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) RegisterShop(registerDTO *ShopRegisterDTORequest) (*models.Farmer, error) {

	password, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		return nil, err
	}
	shop := &models.Farmer{
		Name:      registerDTO.Name,
		Password:  password,
		Email:     registerDTO.Email,
		Mobile:    registerDTO.Mobile,
		Address:   registerDTO.Address,
		Latitude:  registerDTO.Latitude,
		Longitude: registerDTO.Longitude,
		Location: gogis.Point{
			Lng: registerDTO.Longitude,
			Lat: registerDTO.Latitude,
		},
	}

	if err := s.repository.Create(shop); err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *Service) AuthenticateShop(request *ShopLoginDTORequest) (*models.Farmer, error) {
	shop, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if shop == nil {
		return nil, nil
	}

	if err := utils.CheckPasswordHash(request.Password, shop.Password); err != nil {
		return nil, nil
	}

	return shop, nil
}

func (s *Service) GetShopByID(shopID uint) (*models.Farmer, error) {
	return s.repository.FindByID(shopID)

}

func (s *Service) GetNearbyFarmers(lat float64, lon float64, radius float64, limit int) ([]NearByShopsDTORespone, error) {
	farmers, err := s.repository.FindNearbyShops(lat, lon, radius, limit)
	if err != nil {
		return nil, err
	}

	return farmers, nil
}

func (s *Service) SubscribeFarmer(userID uint, farmerID uint) (uint, error) {
	subscriberCount, err := s.repository.SubscribeFarmer(farmerID, userID)
	if err != nil {
		return 0, err
	}
	return subscriberCount, nil
}

func (s *Service) UpdateFarmerStatus(farmerID uint, status bool) error {
	return s.repository.UpdateFarmerStatus(farmerID, status)
}

func (s *Service) UnsubscribeFarmer(userID uint, farmerID uint) (uint, error) {
	subscriberCount, err := s.repository.UnsubscribeShop(farmerID, userID)
	if err != nil {
		return 0, err
	}
	return subscriberCount, nil
}

func (s *Service) GetFarmerDetails(shopID uint, userID uint) (*models.Farmer, bool, error) {
	shop, isSubscribed, err := s.repository.GetShopDetails(shopID, userID)
	if err != nil {
		return nil, false, err
	}
	return shop, isSubscribed, nil
}

func (s *Service) GetUserSubscribedShops(userID uint) ([]models.Farmer, error) {
	return s.repository.GetUserSubscribedShops(userID)
}
