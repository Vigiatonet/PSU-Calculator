package services

import (
	"time"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/data/cache"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"github.com/Vigiatonet/PSU-Calculator/pkg/service_errors"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	Cfg    *config.Config
	Logger logging.Logger
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		Cfg:    cfg,
		Logger: logger,
	}
}

func (ts *TokenService) GenerateToken(req *dto.TokenDTO) (*dto.TokenDetail, error) {
	tkDetail := &dto.TokenDetail{}
	tkDetail.AccessTokenExpireTime = time.Now().Add(ts.Cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	tkDetail.RefreshTokenExpireTime = time.Now().Add(ts.Cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[constants.UserIdKey] = req.UserId
	accessTokenClaims[constants.UserNameKey] = req.Username
	accessTokenClaims[constants.RolesKey] = req.Roles
	accessTokenClaims[constants.ExpKey] = tkDetail.AccessTokenExpireTime
	accessTokenClaims[constants.AccessType] = true

	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	var err error
	tkDetail.AccessToken, err = acToken.SignedString([]byte(ts.Cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[constants.UserIdKey] = req.UserId
	refreshTokenClaims[constants.ExpKey] = tkDetail.RefreshTokenExpireTime
	refreshTokenClaims[constants.RefreshType] = true

	rtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tkDetail.RefreshToken, err = rtToken.SignedString([]byte(ts.Cfg.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return tkDetail, nil
}

func (ts *TokenService) validateToken(token string) (*jwt.Token, error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenInvalid}
		}
		return []byte(ts.Cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	blackListed := isBlackList(token)
	if blackListed {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenInvalid}
	}
	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {

		if _, ok := claims[constants.ExpKey].(float64); !ok {
			return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenInvalid}
		}

		expTime := time.Unix(int64(claims[constants.ExpKey].(float64)), 0)
		nowTime := time.Now()
		if nowTime.After(expTime) {
			return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenExpired}
		}
	}
	return tk, nil
}

func isBlackList(Token string) bool {
	rds := cache.GetRedis()
	_, err := cache.Get[bool](Token, rds)
	return err == nil
}

func AddToBlacklist(token string, ttl time.Duration) error {
	rds := cache.GetRedis()

	err := cache.Set(token, true, ttl, rds)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TokenService) GetClaims(token string) (map[string]interface{}, error) {
	claimMap := map[string]interface{}{}
	verification, err := ts.validateToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verification.Claims.(jwt.MapClaims)
	if ok && verification.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMsg: service_errors.ClaimNotFound}
}

func (ts *TokenService) ValidateRefreshToken(req *dto.RefreshTokenDTO) (*dto.TokenDetail, error) {
	tk, err := jwt.Parse(req.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(ts.Cfg.JWT.RefreshSecret), nil
	})
	if err != nil {
		return nil, &service_errors.ServiceError{
			EndUserMsg: service_errors.TokenInvalid,
			Err:        err,
		}
	}
	isBlackListed := isBlackList(req.RefreshToken)
	claimMap := tk.Claims.(jwt.MapClaims)
	if _, ok := claimMap[constants.RefreshType]; !ok {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.NotRefreshToken}

	}

	if isBlackListed || !tk.Valid {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenInvalid}
	}

	if _, ok := claimMap[constants.ExpKey].(float64); !ok {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenInvalid}
	}
	tokenExpTime := time.Unix(int64(claimMap[constants.ExpKey].(float64)), 0)
	timeNow := time.Now()
	if timeNow.After(tokenExpTime) {
		return nil, &service_errors.ServiceError{EndUserMsg: service_errors.TokenExpired}
	}
	AddToBlacklist(req.RefreshToken, time.Minute*ts.Cfg.JWT.RefreshTokenExpireDuration)

	userId := claimMap[constants.UserIdKey]
	var user models.User
	tx := db.GetDB().Begin()
	err = tx.Model(&models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	tDto := &dto.TokenDTO{
		UserId:   user.ID,
		Username: user.Username,
	}
	if len(user.UserRoles) > 0 {
		for _, r := range user.UserRoles {
			tDto.Roles = append(tDto.Roles, r.Role.Name)
		}
	}

	accessToken, err := ts.GenerateToken(tDto)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
