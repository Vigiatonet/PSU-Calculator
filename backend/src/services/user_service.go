package services

import (
	"context"
	"database/sql"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"github.com/Vigiatonet/PSU-Calculator/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB     *gorm.DB
	Logger logging.Logger
	Token  *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	Tk := NewTokenService(cfg)
	return &UserService{
		DB:     db,
		Logger: logger,
		Token:  Tk,
	}
}

func (us *UserService) existsByEmail(userEmail string) (bool, error) {
	var exists bool
	err := us.DB.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", userEmail).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	err := us.DB.Model(&models.User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) getDefaultRole() (roleId int, err error) {
	if err := us.DB.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).
		Error; err != nil {
		return -1, err
	}
	return roleId, nil
}

func (us *UserService) RegisterByUsername(req *dto.RegisterByUsername) error {
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  sql.NullString{Valid: true, String: req.LastName},
		Username:  req.Username,
		Email:     sql.NullString{Valid: true, String: req.Email},
	}
	exists, err := us.existsByEmail(req.Email)
	if err == nil && exists {
		return &service_errors.ServiceError{EndUserMsg: service_errors.EmailExists}
	} else if err != nil {
		return err
	}
	exists, err = us.existsByUsername(req.Username)

	if err == nil && exists {
		return &service_errors.ServiceError{EndUserMsg: service_errors.EmailExists}
	} else if err != nil {
		return err
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bs)
	roleId, err := us.getDefaultRole()
	if err != nil {
		return nil
	}
	tx := us.DB.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(err, logging.Postgres, logging.Insert, "can't create userRole", nil)
		return err
	}
	userRole := &models.UserRole{
		UserId: user.ID,
		RoleId: roleId,
	}
	err = tx.Create(&userRole).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(err, logging.Postgres, logging.Insert, "can't create userRole", nil)
		return err
	}
	tx.Commit()
	return nil
}

func (us *UserService) LoginByUsername(req *dto.LoginByUsername) (*dto.TokenDetail, error) {
	var user models.User
	db := us.DB.Begin()
	err := db.Model(&models.User{}).Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		us.Logger.Error(err, logging.Postgres, logging.Get, "can't Get user", nil)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.WrongPassword}
	}

	tDTO := &dto.TokenDTO{
		UserId:   user.ID,
		Username: user.Username,
	}
	if len(user.UserRoles) > 0 {
		for _, r := range user.UserRoles {
			tDTO.Roles = append(tDTO.Roles, r.Role.Name)
		}
	}
	tk, err := us.Token.GenerateToken(tDTO)
	if err != nil {
		return nil, err
	}
	return tk, err
}

func (us *UserService) ShowUser(ctx context.Context) (*dto.UserResponse, error) {
	id := int64(ctx.Value(constants.UserIdKey).(float64))
	var user models.User
	db := us.DB.Begin()
	err := db.Model(&user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	res := &dto.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email.String,
		FirstName: user.FirstName,
		LastName:  user.LastName.String,
		Activated: user.Enable,
	}

	return res, nil
}
