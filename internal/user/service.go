package user

import (
	"test-chat/internal/db/postgres"
	"test-chat/internal/entity"
)

type UserService struct {
	db *postgres.Database
}

func NewUserService() *UserService {
	return &UserService{db: postgres.GetDatabase()}
}

func (rs *UserService) CreateUser(name string) (entity.User, error) {
	user := entity.User{
		ID:       0,
		Username: name,
	}

	result := rs.db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (rs *UserService) GetUser(id uint) (entity.User, error) {
	var user entity.User

	result := rs.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (rs *UserService) GetUserByRoomID(RoomID uint) (entity.User, error) {
	var users []entity.User

	result := rs.db.Model(&entity.Room{}).Where("id = ?", RoomID).Association("Users").Find(&users)

	if result.Error != nil {
		return entity.User{}, nil
	}

	return users[0], nil
}
