package repository

import (
	"context"
	"strconv"
	"test/pkg/db"
	"test/pkg/request"
)

type IUserRepository interface {
	Create(ctx context.Context, user User) error
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c request.Conditions) ([]User, error)
	Count() int64
}

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user User) error {
	return repo.Database.WithContext(ctx).Create(&user).Error
}

func (repo *UserRepository) GetByID(ctx context.Context, id string) (User, error) {
	var user User
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}
	err = repo.Database.WithContext(ctx).First(&user, "id = ?", idInt).Error
	return user, err
}

func (repo *UserRepository) Update(ctx context.Context, user User) error {
	return repo.Database.WithContext(ctx).Save(&user).Error
}

func (repo *UserRepository) Delete(ctx context.Context, id string) error {
	return repo.Database.WithContext(ctx).Delete(&User{}, "id = ?", id).Error
}

func (repo *UserRepository) Count() int64 {
	var count int64
	repo.Database.Model(&User{}).Count(&count)
	return count
}

func (repo *UserRepository) List(ctx context.Context, c request.Conditions) ([]User, error) {
	var users []User

	result := repo.Database.
		WithContext(ctx).
		Table("users").
		Order("id desc").
		Limit(c.Limit).
		Offset(c.Offset).
		Scan(&users)

	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}
