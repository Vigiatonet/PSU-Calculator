package migrations

import (
	"database/sql"

	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_01() {
	db := db.GetDB()
	createTables(db)

}

func createTables(db *gorm.DB) {
	tables := []interface{}{}
	tables = addTablesIfNotExists(&models.User{}, db, tables)
	tables = addTablesIfNotExists(&models.Role{}, db, tables)
	tables = addTablesIfNotExists(&models.UserRole{}, db, tables)
	tables = addTablesIfNotExists(&models.OpticalDrive{}, db, tables)
	tables = addTablesIfNotExists(&models.HardDrive{}, db, tables)
	tables = addTablesIfNotExists(&models.Ram{}, db, tables)
	tables = addTablesIfNotExists(&models.GpuBrand{}, db, tables)
	tables = addTablesIfNotExists(&models.Graphic{}, db, tables)
	tables = addTablesIfNotExists(&models.Motherboard{}, db, tables)
	tables = addTablesIfNotExists(&models.CpuBrand{}, db, tables)
	tables = addTablesIfNotExists(&models.Cpu{}, db, tables)
	tables = addTablesIfNotExists(&models.Ssd{}, db, tables)

	err := db.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(err, logging.Postgres, logging.Insert, "cant add tables", nil)
		panic(err)
	}
	createDefaultInfo(db)
}

func addTablesIfNotExists(model interface{}, db *gorm.DB, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(&model) {
		tables = append(tables, model)
	}
	return tables
}

func createRoleIfNotExists(db *gorm.DB, r *models.Role) {
	exists := 0
	db.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exists)
	if exists == 0 {
		db.Create(&r)
	}
}

func createDefaultInfo(db *gorm.DB) {
	admin := models.Role{Name: "admin"}
	createRoleIfNotExists(db, &admin)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(db, &defaultRole)

	u := models.User{
		FirstName: "test",
		LastName:  sql.NullString{Valid: true, String: "test"},
		Username:  constants.AdminRoleName,
		Email:     sql.NullString{Valid: true, String: "test@test.com"},
		Enable:    true,
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte("a123"), bcrypt.MinCost)
	u.Password = string(hashedPass)
	createAdmin(db, &u, admin.ID)
}

func createAdmin(db *gorm.DB, usr *models.User, roleID int) {
	exists := 0
	db.Model(&models.User{}).Select("1").Where("username = ?", usr.Username).First(&exists)
	if exists == 0 {
		db.Create(&usr)
		userRole := models.UserRole{UserId: usr.ID, RoleId: roleID}
		db.Create(&userRole)
	}
}
