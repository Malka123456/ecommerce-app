package service

import (
	"ecommerce-app/internal/domain"
	"ecommerce-app/internal/dto"
	"log"
)

type UserService struct {

}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	log.Println(input)
	return "this_is_my_token_as_of_now",nil

}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	//perform some db operation
	//business logic
	return nil, nil
}

func (s UserService) Login(input any) (string, error) {

	return "",nil

}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil

}

func (s UserService) VerifyCode(id uint, code int)  error {

	return nil

}

func (s UserService) CreateProfile(id uint, input any)  error {

	return nil

}

func (s UserService) GetProfile(id uint) (*domain.User, error)  {

	return nil, nil

}

func (s UserService) UpdateProfile(id uint, input any)  error {

	return nil

}

func (s UserService) BecomeSeller(id uint, input any)  (string, error) {

	return "", nil

}

func (s UserService) FindCart(id uint)  ([]interface{}, error) {

	return nil, nil

}

func (s UserService) CreateCart(input any, u domain.User)  ([]interface{}, error) {

	return nil, nil

}

func (s UserService) CreateOrder(u domain.User)  (int, error) {

	return 0, nil

}

func (s UserService) GetOrders(u domain.User)  ([]interface{}, error) {

	return nil, nil

}

func (s UserService) GetOrderByTd(id uint, uId uint)  ([]interface{}, error) {

	return nil, nil

}




