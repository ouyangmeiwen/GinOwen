package config

import (
	"GINOWEN/global"
	"GINOWEN/model"
	"GINOWEN/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	//"github.com/dzwvip/oracle"
	//_ "github.com/godror/godror"
	"github.com/glebarez/sqlite"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func LoadConfig() global.YarmConfig {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	var config global.YarmConfig
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return config
}

func InitDB() *gorm.DB {
	config := global.OWEN_CONFIG
	var dbErr error

	// 配置 Gorm 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer 输出日志
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
		},
	)

	switch config.DB.Type {
	case "mysql":
		DB, dbErr = gorm.Open(mysql.Open(config.DB.Mysql), &gorm.Config{Logger: newLogger})
	case "postgres":
		DB, dbErr = gorm.Open(postgres.Open(config.DB.Postgres), &gorm.Config{Logger: newLogger})
	case "sqlite":
		DB, dbErr = gorm.Open(sqlite.Open(config.DB.Sqlite), &gorm.Config{Logger: newLogger})
	case "sqlserver":
		DB, dbErr = gorm.Open(sqlserver.Open(config.DB.Mssql), &gorm.Config{Logger: newLogger})
	case "oracle":
		// oracleConfig := oracle.Config{
		// 	DSN:               config.DB.Oracle, // DSN data source name
		// 	DefaultStringSize: 191,              // string 类型字段的默认长度
		// }
		// DB, dbErr = gorm.Open(oracle.New(oracleConfig), &gorm.Config{Logger: newLogger})
	default:
		log.Fatalf("Unsupported database type: %s", config.DB.Type)
	}

	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	// 配置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)                                      // 最大连接数
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)                                      // 最大空闲连接数
	sqlDB.SetConnMaxLifetime((time.Duration(config.DB.ConnMaxLifetime)) * time.Minute) // 连接的最大生命周期

	return DB
}
func InitRedis() {
	redisCfg := global.OWEN_CONFIG.Redis
	if len(redisCfg.Addr) <= 0 {
		return
	}
	var client redis.UniversalClient
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		global.OWEN_REDIS = client
	}
}

func AutoMigrateDB() {
	// 自动迁移数据库结构
	var err error
	err = DB.AutoMigrate(&models.OwenUser{})
	err = DB.AutoMigrate(&models.OwenRole{})
	err = DB.AutoMigrate(&models.OwenAuditLog{})

	//err = DB.AutoMigrate(&model.Efmigrationshistory{})
	err = DB.AutoMigrate(&model.Abpauditlog{})
	err = DB.AutoMigrate(&model.Abpbackgroundjob{})
	err = DB.AutoMigrate(&model.Abpedition{})
	err = DB.AutoMigrate(&model.Abpentitychange{})
	err = DB.AutoMigrate(&model.Abpentitychangeset{})
	err = DB.AutoMigrate(&model.Abpentitypropertychange{})
	err = DB.AutoMigrate(&model.Abpfeature{})
	err = DB.AutoMigrate(&model.Abplanguage{})
	err = DB.AutoMigrate(&model.Abplanguagetext{})
	err = DB.AutoMigrate(&model.Abpnotification{})
	err = DB.AutoMigrate(&model.Abpnotificationsubscription{})
	err = DB.AutoMigrate(&model.Abporganizationunitrole{})
	err = DB.AutoMigrate(&model.Abporganizationunit{})
	err = DB.AutoMigrate(&model.Abppermission{})
	err = DB.AutoMigrate(&model.Abppersistedgrant{})
	err = DB.AutoMigrate(&model.Abproleclaim{})
	err = DB.AutoMigrate(&model.Abprole{})
	err = DB.AutoMigrate(&model.Abpsetting{})
	err = DB.AutoMigrate(&model.Abptenantnotification{})
	err = DB.AutoMigrate(&model.Abptenant{})
	err = DB.AutoMigrate(&model.Abpuseraccount{})
	err = DB.AutoMigrate(&model.Abpuserclaim{})
	err = DB.AutoMigrate(&model.Abpuserloginattempt{})
	err = DB.AutoMigrate(&model.Abpuserlogin{})
	err = DB.AutoMigrate(&model.Abpusernotification{})
	err = DB.AutoMigrate(&model.Abpuserorganizationunit{})
	err = DB.AutoMigrate(&model.Abpuserrole{})
	err = DB.AutoMigrate(&model.Abpuser{})
	err = DB.AutoMigrate(&model.Abpusertoken{})
	err = DB.AutoMigrate(&model.Appaliuser{})
	err = DB.AutoMigrate(&model.Appapprovalinfo{})
	err = DB.AutoMigrate(&model.Appapprovaltemplate{})
	err = DB.AutoMigrate(&model.Appbinaryobject{})
	err = DB.AutoMigrate(&model.Appbookorder{})
	err = DB.AutoMigrate(&model.Appchatmessage{})
	err = DB.AutoMigrate(&model.Appcreditloginorder{})
	err = DB.AutoMigrate(&model.Appfriendship{})
	err = DB.AutoMigrate(&model.Appinvoice{})
	err = DB.AutoMigrate(&model.Appitemlocked{})
	err = DB.AutoMigrate(&model.Appmessageboard{})
	err = DB.AutoMigrate(&model.Appnotificationlog{})
	err = DB.AutoMigrate(&model.Apppayorder{})
	err = DB.AutoMigrate(&model.Apppickupcode{})
	err = DB.AutoMigrate(&model.Appqrcode{})
	err = DB.AutoMigrate(&model.Apprecommendinfo{})
	err = DB.AutoMigrate(&model.Appsubbookinfo{})
	err = DB.AutoMigrate(&model.Appsubscriptionpayment{})
	err = DB.AutoMigrate(&model.Appusercard{})
	err = DB.AutoMigrate(&model.Appweuser{})
	err = DB.AutoMigrate(&model.Dasbusinesscount{})
	err = DB.AutoMigrate(&model.Dascirculatecount{})
	err = DB.AutoMigrate(&model.Dasdatabaselink{})
	err = DB.AutoMigrate(&model.Dasdatasource{})
	err = DB.AutoMigrate(&model.Dasfeecount{})
	err = DB.AutoMigrate(&model.Daspatronlogcount{})
	err = DB.AutoMigrate(&model.Dasperformance{})
	err = DB.AutoMigrate(&model.Dassecuritygatecount{})
	err = DB.AutoMigrate(&model.Dasvisitpage{})
	err = DB.AutoMigrate(&model.Dasvisittrend{})
	err = DB.AutoMigrate(&model.Hangfireaggregatedcounter{})
	err = DB.AutoMigrate(&model.Hangfirecounter{})
	err = DB.AutoMigrate(&model.Hangfiredistributedlock{})
	err = DB.AutoMigrate(&model.Hangfirehash{})
	err = DB.AutoMigrate(&model.Hangfirejob{})
	err = DB.AutoMigrate(&model.Hangfirejobparameter{})
	err = DB.AutoMigrate(&model.Hangfirejobqueue{})
	err = DB.AutoMigrate(&model.Hangfirejobstate{})
	err = DB.AutoMigrate(&model.Hangfirelist{})
	err = DB.AutoMigrate(&model.Hangfireserver{})
	err = DB.AutoMigrate(&model.Hangfireset{})
	err = DB.AutoMigrate(&model.Hangfirestate{})
	err = DB.AutoMigrate(&model.Lcpcommandlog{})
	err = DB.AutoMigrate(&model.Lcpconfig{})
	err = DB.AutoMigrate(&model.Lcpmaintainlog{})
	err = DB.AutoMigrate(&model.Lcpproduct{})
	err = DB.AutoMigrate(&model.Lcprfidantenna{})
	err = DB.AutoMigrate(&model.Lcprfidreader{})
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookaccesslog{})
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookdailyaccess{})
	err = DB.AutoMigrate(&model.Lcpsecuritygateitemlog{})
	err = DB.AutoMigrate(&model.Lcpsecuritygatepatronlog{})
	err = DB.AutoMigrate(&model.Lcpserialport{})
	err = DB.AutoMigrate(&model.Lcpserialportext{})
	err = DB.AutoMigrate(&model.Lcpservice{})
	err = DB.AutoMigrate(&model.Lcpterminal{})
	err = DB.AutoMigrate(&model.Lcpterminaladvertisement{})
	err = DB.AutoMigrate(&model.Lcpterminalbox{})
	err = DB.AutoMigrate(&model.Lcpterminalboxitem{})
	err = DB.AutoMigrate(&model.Lcpterminaldevice{})
	err = DB.AutoMigrate(&model.Lcpterminaldevicelog{})
	err = DB.AutoMigrate(&model.Lcpterminallog{})
	err = DB.AutoMigrate(&model.Lcpterminalpermission{})
	err = DB.AutoMigrate(&model.Lcpterminalshelf{})
	err = DB.AutoMigrate(&model.Lcpterminalshelfitem{})
	err = DB.AutoMigrate(&model.Lcpterminalshelflog{})
	err = DB.AutoMigrate(&model.Lcpupgradelog{})
	err = DB.AutoMigrate(&model.Lcpversion{})
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfo{})
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoitem{})
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoprofile{})
	err = DB.AutoMigrate(&model.Libailibrarainknowledgefileinfo{})
	err = DB.AutoMigrate(&model.Libailibrarainquestionmetric{})
	err = DB.AutoMigrate(&model.Libailibrarainsessionmetric{})
	err = DB.AutoMigrate(&model.Libainirobotinfo{})
	err = DB.AutoMigrate(&model.Libbatchinfo{})
	err = DB.AutoMigrate(&model.Libbatchoperateindex{})
	err = DB.AutoMigrate(&model.Libbatchoperatelog{})
	err = DB.AutoMigrate(&model.Libbookinfo{})
	err = DB.AutoMigrate(&model.Libcirculatelog{})
	err = DB.AutoMigrate(&model.LibcirculatelogBak{})
	err = DB.AutoMigrate(&model.Libfeedback{})
	err = DB.AutoMigrate(&model.Libfeelog{})
	err = DB.AutoMigrate(&model.Libinventorystat{})
	err = DB.AutoMigrate(&model.Libinventorytask{})
	err = DB.AutoMigrate(&model.Libinventorywork{})
	err = DB.AutoMigrate(&model.Libinventoryworkdetail{})
	err = DB.AutoMigrate(&model.Libinventoryworklog{})
	err = DB.AutoMigrate(&model.Libitem{})
	err = DB.AutoMigrate(&model.LibitemCopy{})
	err = DB.AutoMigrate(&model.Libiteminventoryinfo{})
	err = DB.AutoMigrate(&model.Libiteminventorylog{})
	err = DB.AutoMigrate(&model.Libitemlocinfo{})
	err = DB.AutoMigrate(&model.Libitemoperateindexlog{})
	err = DB.AutoMigrate(&model.Libitemoperatelog{})
	err = DB.AutoMigrate(&model.Libjournalinfo{})
	err = DB.AutoMigrate(&model.Liblabel{})
	err = DB.AutoMigrate(&model.Liblabeloperatelog{})
	err = DB.AutoMigrate(&model.Liblayer{})
	err = DB.AutoMigrate(&model.Liblayerindexupdatelog{})
	err = DB.AutoMigrate(&model.Libnotificationlog{})
	err = DB.AutoMigrate(&model.Libpartonreservation{})
	err = DB.AutoMigrate(&model.Libpatron{})
	err = DB.AutoMigrate(&model.Libpatronitem{})
	err = DB.AutoMigrate(&model.Libpatronlog{})
	err = DB.AutoMigrate(&model.Libpointsclearing{})
	err = DB.AutoMigrate(&model.Libpointslog{})
	err = DB.AutoMigrate(&model.Librfidantlayer{})
	err = DB.AutoMigrate(&model.Librfidscandetaillog{})
	err = DB.AutoMigrate(&model.Librfidscanlog{})
	err = DB.AutoMigrate(&model.Librow{})
	err = DB.AutoMigrate(&model.Librowcatalog{})
	err = DB.AutoMigrate(&model.Libscanitemlog{})
	err = DB.AutoMigrate(&model.Libshelf{})
	err = DB.AutoMigrate(&model.Libstruct{})
	err = DB.AutoMigrate(&model.Libtagtobarcodelog{})
	err = DB.AutoMigrate(&model.Libtaskitem{})
	err = DB.AutoMigrate(&model.Libtaskpackage{})
	err = DB.AutoMigrate(&model.Libupdatefirstbooklog{})
	err = DB.AutoMigrate(&model.Misactivity{})
	err = DB.AutoMigrate(&model.Mismediainfo{})
	err = DB.AutoMigrate(&model.Misnews{})
	err = DB.AutoMigrate(&model.Mistemplate{})
	err = DB.AutoMigrate(&model.Rescatalog{})
	err = DB.AutoMigrate(&model.Rescipinfo{})
	err = DB.AutoMigrate(&model.Resfourcorner{})
	err = DB.AutoMigrate(&model.Resjournalinfo{})
	err = DB.AutoMigrate(&model.Resnotfound{})
	err = DB.AutoMigrate(&model.Respublisherinfo{})
	err = DB.AutoMigrate(&model.Ssbackgroundjob{})
	err = DB.AutoMigrate(&model.Sysattachment{})
	err = DB.AutoMigrate(&model.Sysauditacslog{})
	err = DB.AutoMigrate(&model.Sysauditapilog{})
	err = DB.AutoMigrate(&model.Sysauditapplog{})
	err = DB.AutoMigrate(&model.Sysauditlinklog{})
	err = DB.AutoMigrate(&model.Sysauditlmslog{})
	err = DB.AutoMigrate(&model.Sysauditsslog{})
	err = DB.AutoMigrate(&model.Sysblocklist{})
	err = DB.AutoMigrate(&model.Sysbookblocklist{})
	err = DB.AutoMigrate(&model.Sysbooknumlib{})
	err = DB.AutoMigrate(&model.Sysbooknumset{})
	err = DB.AutoMigrate(&model.Syscardconfig{})
	err = DB.AutoMigrate(&model.Syscarddevdtl{})
	err = DB.AutoMigrate(&model.Syscardtype{})
	err = DB.AutoMigrate(&model.Syscoderule{})
	err = DB.AutoMigrate(&model.Syscoderuleseed{})
	err = DB.AutoMigrate(&model.Sysconfigbase{})
	err = DB.AutoMigrate(&model.Sysconfiglog{})
	err = DB.AutoMigrate(&model.Sysdataitem{})
	err = DB.AutoMigrate(&model.Sysdataitemdetail{})
	err = DB.AutoMigrate(&model.Sysdatalog{})
	err = DB.AutoMigrate(&model.Sysdepartment{})
	err = DB.AutoMigrate(&model.Sysenumfield{})
	err = DB.AutoMigrate(&model.Sysenumvalue{})
	err = DB.AutoMigrate(&model.Sysenumvalue2{})
	err = DB.AutoMigrate(&model.Sysfaceoffineoperationlog{})
	err = DB.AutoMigrate(&model.Sysfaceofflinefeature{})
	err = DB.AutoMigrate(&model.Sysfaceofflinegroup{})
	err = DB.AutoMigrate(&model.Sysfaceofflineuser{})
	err = DB.AutoMigrate(&model.Syslanguage{})
	err = DB.AutoMigrate(&model.Syslayertran{})
	err = DB.AutoMigrate(&model.Syslocation{})
	err = DB.AutoMigrate(&model.Sysmenu{})
	err = DB.AutoMigrate(&model.SysmenuCopy{})
	err = DB.AutoMigrate(&model.Systasklist{})
	err = DB.AutoMigrate(&model.Systenantconfig{})
	err = DB.AutoMigrate(&model.Systenantextend{})

	// err = DB.AutoMigrate(&models.TestOrder{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	addDefaultData()
}
func addDefaultData() {

	// 检查并添加默认角色数据
	if err := DB.First(&models.OwenRole{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		defaultRoles := []models.OwenRole{
			{Name: "Admin", Permissions: "all"},
			{Name: "User", Permissions: "sysauditlmslog"},
		}
		for _, role := range defaultRoles {
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Failed to insert default role %s: %v", role.Name, err)
			}
		}
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// 检查并添加默认用户数据
	if err := DB.First(&models.OwenUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		defaultUsers := []models.OwenUser{
			{Username: "admin", Password: string(hashedPassword), RoleID: 1}, // 示例用户
			{Username: "user", Password: string(hashedPassword), RoleID: 2},
		}
		for _, user := range defaultUsers {
			if err := DB.Create(&user).Error; err != nil {
				log.Printf("Failed to insert default user %s: %v", user.Username, err)
			}
		}
	}

}

// InitMongoDB 连接 MongoDB
func InitMongoDB() (*mongo.Client, error) {
	config := global.OWEN_CONFIG
	if len(config.MongoDB.URI) <= 0 {
		return nil, errors.New("入参不能空")
	}
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	if config.MongoDB.Username != "" && config.MongoDB.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.MongoDB.Username,
			Password: config.MongoDB.Password,
		})
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	global.OWEN_MONGO = client
	return client, nil
}
