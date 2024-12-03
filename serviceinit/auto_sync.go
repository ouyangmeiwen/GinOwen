package serviceinit

import (
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
	"log"
)

func CusSyncDatabase() error {
	batchSize := 500 // 每批次同步的数据量
	var offset int   // 用于分页

	// Syncing Abpauditlog model data
	offset = 0
	for {
		var AbpauditlogData []model.Abpauditlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpauditlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpauditlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpauditlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpauditlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpauditlog")
			return err
		}
	}

	// Syncing Abpbackgroundjob model data
	offset = 0
	for {
		var AbpbackgroundjobData []model.Abpbackgroundjob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpbackgroundjobData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpbackgroundjob")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpbackgroundjobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpbackgroundjobData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpbackgroundjob")
			return err
		}
	}

	// Syncing Abpedition model data
	offset = 0
	for {
		var AbpeditionData []model.Abpedition
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpeditionData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpedition")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpeditionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpeditionData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpedition")
			return err
		}
	}

	// Syncing Abpentitychange model data
	offset = 0
	for {
		var AbpentitychangeData []model.Abpentitychange
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitychangeData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpentitychange")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitychangeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitychangeData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpentitychange")
			return err
		}
	}

	// Syncing Abpentitychangeset model data
	offset = 0
	for {
		var AbpentitychangesetData []model.Abpentitychangeset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitychangesetData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpentitychangeset")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitychangesetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitychangesetData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpentitychangeset")
			return err
		}
	}

	// Syncing Abpentitypropertychange model data
	offset = 0
	for {
		var AbpentitypropertychangeData []model.Abpentitypropertychange
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitypropertychangeData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpentitypropertychange")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitypropertychangeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitypropertychangeData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpentitypropertychange")
			return err
		}
	}

	// Syncing Abpfeature model data
	offset = 0
	for {
		var AbpfeatureData []model.Abpfeature
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpfeatureData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpfeature")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpfeatureData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpfeatureData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpfeature")
			return err
		}
	}

	// Syncing Abplanguage model data
	offset = 0
	for {
		var AbplanguageData []model.Abplanguage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbplanguageData).Error; err != nil {
			log.Fatalf("Failed to query table: Abplanguage")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbplanguageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbplanguageData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abplanguage")
			return err
		}
	}

	// Syncing Abplanguagetext model data
	offset = 0
	for {
		var AbplanguagetextData []model.Abplanguagetext
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbplanguagetextData).Error; err != nil {
			log.Fatalf("Failed to query table: Abplanguagetext")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbplanguagetextData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbplanguagetextData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abplanguagetext")
			return err
		}
	}

	// Syncing Abpnotification model data
	offset = 0
	for {
		var AbpnotificationData []model.Abpnotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpnotificationData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpnotification")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpnotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpnotificationData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpnotification")
			return err
		}
	}

	// Syncing Abpnotificationsubscription model data
	offset = 0
	for {
		var AbpnotificationsubscriptionData []model.Abpnotificationsubscription
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpnotificationsubscriptionData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpnotificationsubscription")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpnotificationsubscriptionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpnotificationsubscriptionData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpnotificationsubscription")
			return err
		}
	}

	// Syncing Abporganizationunitrole model data
	offset = 0
	for {
		var AbporganizationunitroleData []model.Abporganizationunitrole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbporganizationunitroleData).Error; err != nil {
			log.Fatalf("Failed to query table: Abporganizationunitrole")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbporganizationunitroleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbporganizationunitroleData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abporganizationunitrole")
			return err
		}
	}

	// Syncing Abporganizationunit model data
	offset = 0
	for {
		var AbporganizationunitData []model.Abporganizationunit
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbporganizationunitData).Error; err != nil {
			log.Fatalf("Failed to query table: Abporganizationunit")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbporganizationunitData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbporganizationunitData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abporganizationunit")
			return err
		}
	}

	// Syncing Abppermission model data
	offset = 0
	for {
		var AbppermissionData []model.Abppermission
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbppermissionData).Error; err != nil {
			log.Fatalf("Failed to query table: Abppermission")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbppermissionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbppermissionData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abppermission")
			return err
		}
	}

	// Syncing Abppersistedgrant model data
	offset = 0
	for {
		var AbppersistedgrantData []model.Abppersistedgrant
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbppersistedgrantData).Error; err != nil {
			log.Fatalf("Failed to query table: Abppersistedgrant")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbppersistedgrantData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbppersistedgrantData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abppersistedgrant")
			return err
		}
	}

	// Syncing Abproleclaim model data
	offset = 0
	for {
		var AbproleclaimData []model.Abproleclaim
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbproleclaimData).Error; err != nil {
			log.Fatalf("Failed to query table: Abproleclaim")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbproleclaimData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbproleclaimData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abproleclaim")
			return err
		}
	}

	// Syncing Abprole model data
	offset = 0
	for {
		var AbproleData []model.Abprole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbproleData).Error; err != nil {
			log.Fatalf("Failed to query table: Abprole")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbproleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbproleData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abprole")
			return err
		}
	}

	// Syncing Abpsetting model data
	offset = 0
	for {
		var AbpsettingData []model.Abpsetting
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpsettingData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpsetting")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpsettingData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpsettingData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpsetting")
			return err
		}
	}

	// Syncing Abptenantnotification model data
	offset = 0
	for {
		var AbptenantnotificationData []model.Abptenantnotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbptenantnotificationData).Error; err != nil {
			log.Fatalf("Failed to query table: Abptenantnotification")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbptenantnotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbptenantnotificationData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abptenantnotification")
			return err
		}
	}

	// Syncing Abptenant model data
	offset = 0
	for {
		var AbptenantData []model.Abptenant
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbptenantData).Error; err != nil {
			log.Fatalf("Failed to query table: Abptenant")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbptenantData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbptenantData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abptenant")
			return err
		}
	}

	// Syncing Abpuseraccount model data
	offset = 0
	for {
		var AbpuseraccountData []model.Abpuseraccount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuseraccountData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuseraccount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuseraccountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuseraccountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuseraccount")
			return err
		}
	}

	// Syncing Abpuserclaim model data
	offset = 0
	for {
		var AbpuserclaimData []model.Abpuserclaim
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserclaimData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuserclaim")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserclaimData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserclaimData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuserclaim")
			return err
		}
	}

	// Syncing Abpuserloginattempt model data
	offset = 0
	for {
		var AbpuserloginattemptData []model.Abpuserloginattempt
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserloginattemptData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuserloginattempt")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserloginattemptData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserloginattemptData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuserloginattempt")
			return err
		}
	}

	// Syncing Abpuserlogin model data
	offset = 0
	for {
		var AbpuserloginData []model.Abpuserlogin
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserloginData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuserlogin")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserloginData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserloginData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuserlogin")
			return err
		}
	}

	// Syncing Abpusernotification model data
	offset = 0
	for {
		var AbpusernotificationData []model.Abpusernotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpusernotificationData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpusernotification")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpusernotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpusernotificationData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpusernotification")
			return err
		}
	}

	// Syncing Abpuserorganizationunit model data
	offset = 0
	for {
		var AbpuserorganizationunitData []model.Abpuserorganizationunit
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserorganizationunitData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuserorganizationunit")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserorganizationunitData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserorganizationunitData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuserorganizationunit")
			return err
		}
	}

	// Syncing Abpuserrole model data
	offset = 0
	for {
		var AbpuserroleData []model.Abpuserrole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserroleData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuserrole")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserroleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserroleData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuserrole")
			return err
		}
	}

	// Syncing Abpuser model data
	offset = 0
	for {
		var AbpuserData []model.Abpuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpuser")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpuser")
			return err
		}
	}

	// Syncing Abpusertoken model data
	offset = 0
	for {
		var AbpusertokenData []model.Abpusertoken
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpusertokenData).Error; err != nil {
			log.Fatalf("Failed to query table: Abpusertoken")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpusertokenData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpusertokenData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Abpusertoken")
			return err
		}
	}

	// Syncing Appaliuser model data
	offset = 0
	for {
		var AppaliuserData []model.Appaliuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppaliuserData).Error; err != nil {
			log.Fatalf("Failed to query table: Appaliuser")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppaliuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppaliuserData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appaliuser")
			return err
		}
	}

	// Syncing Appapprovalinfo model data
	offset = 0
	for {
		var AppapprovalinfoData []model.Appapprovalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppapprovalinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Appapprovalinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppapprovalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppapprovalinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appapprovalinfo")
			return err
		}
	}

	// Syncing Appapprovaltemplate model data
	offset = 0
	for {
		var AppapprovaltemplateData []model.Appapprovaltemplate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppapprovaltemplateData).Error; err != nil {
			log.Fatalf("Failed to query table: Appapprovaltemplate")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppapprovaltemplateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppapprovaltemplateData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appapprovaltemplate")
			return err
		}
	}

	// Syncing Appbinaryobject model data
	offset = 0
	for {
		var AppbinaryobjectData []model.Appbinaryobject
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppbinaryobjectData).Error; err != nil {
			log.Fatalf("Failed to query table: Appbinaryobject")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppbinaryobjectData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppbinaryobjectData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appbinaryobject")
			return err
		}
	}

	// Syncing Appbookorder model data
	offset = 0
	for {
		var AppbookorderData []model.Appbookorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppbookorderData).Error; err != nil {
			log.Fatalf("Failed to query table: Appbookorder")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppbookorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppbookorderData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appbookorder")
			return err
		}
	}

	// Syncing Appchatmessage model data
	offset = 0
	for {
		var AppchatmessageData []model.Appchatmessage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppchatmessageData).Error; err != nil {
			log.Fatalf("Failed to query table: Appchatmessage")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppchatmessageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppchatmessageData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appchatmessage")
			return err
		}
	}

	// Syncing Appcreditloginorder model data
	offset = 0
	for {
		var AppcreditloginorderData []model.Appcreditloginorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppcreditloginorderData).Error; err != nil {
			log.Fatalf("Failed to query table: Appcreditloginorder")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppcreditloginorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppcreditloginorderData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appcreditloginorder")
			return err
		}
	}

	// Syncing Appfriendship model data
	offset = 0
	for {
		var AppfriendshipData []model.Appfriendship
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppfriendshipData).Error; err != nil {
			log.Fatalf("Failed to query table: Appfriendship")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppfriendshipData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppfriendshipData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appfriendship")
			return err
		}
	}

	// Syncing Appinvoice model data
	offset = 0
	for {
		var AppinvoiceData []model.Appinvoice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppinvoiceData).Error; err != nil {
			log.Fatalf("Failed to query table: Appinvoice")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppinvoiceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppinvoiceData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appinvoice")
			return err
		}
	}

	// Syncing Appitemlocked model data
	offset = 0
	for {
		var AppitemlockedData []model.Appitemlocked
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppitemlockedData).Error; err != nil {
			log.Fatalf("Failed to query table: Appitemlocked")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppitemlockedData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppitemlockedData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appitemlocked")
			return err
		}
	}

	// Syncing Appmessageboard model data
	offset = 0
	for {
		var AppmessageboardData []model.Appmessageboard
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppmessageboardData).Error; err != nil {
			log.Fatalf("Failed to query table: Appmessageboard")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppmessageboardData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppmessageboardData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appmessageboard")
			return err
		}
	}

	// Syncing Appnotificationlog model data
	offset = 0
	for {
		var AppnotificationlogData []model.Appnotificationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppnotificationlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Appnotificationlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppnotificationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppnotificationlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appnotificationlog")
			return err
		}
	}

	// Syncing Apppayorder model data
	offset = 0
	for {
		var ApppayorderData []model.Apppayorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApppayorderData).Error; err != nil {
			log.Fatalf("Failed to query table: Apppayorder")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ApppayorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApppayorderData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Apppayorder")
			return err
		}
	}

	// Syncing Apppickupcode model data
	offset = 0
	for {
		var ApppickupcodeData []model.Apppickupcode
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApppickupcodeData).Error; err != nil {
			log.Fatalf("Failed to query table: Apppickupcode")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ApppickupcodeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApppickupcodeData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Apppickupcode")
			return err
		}
	}

	// Syncing Appqrcode model data
	offset = 0
	for {
		var AppqrcodeData []model.Appqrcode
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppqrcodeData).Error; err != nil {
			log.Fatalf("Failed to query table: Appqrcode")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppqrcodeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppqrcodeData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appqrcode")
			return err
		}
	}

	// Syncing Apprecommendinfo model data
	offset = 0
	for {
		var ApprecommendinfoData []model.Apprecommendinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApprecommendinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Apprecommendinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ApprecommendinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApprecommendinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Apprecommendinfo")
			return err
		}
	}

	// Syncing Appsubbookinfo model data
	offset = 0
	for {
		var AppsubbookinfoData []model.Appsubbookinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppsubbookinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Appsubbookinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppsubbookinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppsubbookinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appsubbookinfo")
			return err
		}
	}

	// Syncing Appsubscriptionpayment model data
	offset = 0
	for {
		var AppsubscriptionpaymentData []model.Appsubscriptionpayment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppsubscriptionpaymentData).Error; err != nil {
			log.Fatalf("Failed to query table: Appsubscriptionpayment")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppsubscriptionpaymentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppsubscriptionpaymentData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appsubscriptionpayment")
			return err
		}
	}

	// Syncing Appusercard model data
	offset = 0
	for {
		var AppusercardData []model.Appusercard
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppusercardData).Error; err != nil {
			log.Fatalf("Failed to query table: Appusercard")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppusercardData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppusercardData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appusercard")
			return err
		}
	}

	// Syncing Appweuser model data
	offset = 0
	for {
		var AppweuserData []model.Appweuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppweuserData).Error; err != nil {
			log.Fatalf("Failed to query table: Appweuser")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(AppweuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppweuserData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Appweuser")
			return err
		}
	}

	// Syncing Dasbusinesscount model data
	offset = 0
	for {
		var DasbusinesscountData []model.Dasbusinesscount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasbusinesscountData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasbusinesscount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasbusinesscountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasbusinesscountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasbusinesscount")
			return err
		}
	}

	// Syncing Dascirculatecount model data
	offset = 0
	for {
		var DascirculatecountData []model.Dascirculatecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DascirculatecountData).Error; err != nil {
			log.Fatalf("Failed to query table: Dascirculatecount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DascirculatecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DascirculatecountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dascirculatecount")
			return err
		}
	}

	// Syncing Dasdatabaselink model data
	offset = 0
	for {
		var DasdatabaselinkData []model.Dasdatabaselink
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasdatabaselinkData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasdatabaselink")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasdatabaselinkData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasdatabaselinkData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasdatabaselink")
			return err
		}
	}

	// Syncing Dasdatasource model data
	offset = 0
	for {
		var DasdatasourceData []model.Dasdatasource
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasdatasourceData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasdatasource")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasdatasourceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasdatasourceData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasdatasource")
			return err
		}
	}

	// Syncing Dasfeecount model data
	offset = 0
	for {
		var DasfeecountData []model.Dasfeecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasfeecountData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasfeecount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasfeecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasfeecountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasfeecount")
			return err
		}
	}

	// Syncing Daspatronlogcount model data
	offset = 0
	for {
		var DaspatronlogcountData []model.Daspatronlogcount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DaspatronlogcountData).Error; err != nil {
			log.Fatalf("Failed to query table: Daspatronlogcount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DaspatronlogcountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DaspatronlogcountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Daspatronlogcount")
			return err
		}
	}

	// Syncing Dasperformance model data
	offset = 0
	for {
		var DasperformanceData []model.Dasperformance
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasperformanceData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasperformance")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasperformanceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasperformanceData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasperformance")
			return err
		}
	}

	// Syncing Dassecuritygatecount model data
	offset = 0
	for {
		var DassecuritygatecountData []model.Dassecuritygatecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DassecuritygatecountData).Error; err != nil {
			log.Fatalf("Failed to query table: Dassecuritygatecount")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DassecuritygatecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DassecuritygatecountData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dassecuritygatecount")
			return err
		}
	}

	// Syncing Dasvisitpage model data
	offset = 0
	for {
		var DasvisitpageData []model.Dasvisitpage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasvisitpageData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasvisitpage")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasvisitpageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasvisitpageData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasvisitpage")
			return err
		}
	}

	// Syncing Dasvisittrend model data
	offset = 0
	for {
		var DasvisittrendData []model.Dasvisittrend
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasvisittrendData).Error; err != nil {
			log.Fatalf("Failed to query table: Dasvisittrend")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(DasvisittrendData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasvisittrendData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Dasvisittrend")
			return err
		}
	}

	// Syncing Efmigrationshistory model data
	offset = 0
	for {
		var EfmigrationshistoryData []model.Efmigrationshistory
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&EfmigrationshistoryData).Error; err != nil {
			log.Fatalf("Failed to query table: Efmigrationshistory")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(EfmigrationshistoryData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(EfmigrationshistoryData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Efmigrationshistory")
			return err
		}
	}

	// Syncing Hangfireaggregatedcounter model data
	offset = 0
	for {
		var HangfireaggregatedcounterData []model.Hangfireaggregatedcounter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfireaggregatedcounterData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfireaggregatedcounter")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfireaggregatedcounterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfireaggregatedcounterData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfireaggregatedcounter")
			return err
		}
	}

	// Syncing Hangfirecounter model data
	offset = 0
	for {
		var HangfirecounterData []model.Hangfirecounter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirecounterData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirecounter")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirecounterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirecounterData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirecounter")
			return err
		}
	}

	// Syncing Hangfiredistributedlock model data
	offset = 0
	for {
		var HangfiredistributedlockData []model.Hangfiredistributedlock
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfiredistributedlockData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfiredistributedlock")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfiredistributedlockData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfiredistributedlockData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfiredistributedlock")
			return err
		}
	}

	// Syncing Hangfirehash model data
	offset = 0
	for {
		var HangfirehashData []model.Hangfirehash
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirehashData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirehash")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirehashData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirehashData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirehash")
			return err
		}
	}

	// Syncing Hangfirejob model data
	offset = 0
	for {
		var HangfirejobData []model.Hangfirejob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirejob")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirejob")
			return err
		}
	}

	// Syncing Hangfirejobparameter model data
	offset = 0
	for {
		var HangfirejobparameterData []model.Hangfirejobparameter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobparameterData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirejobparameter")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobparameterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobparameterData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirejobparameter")
			return err
		}
	}

	// Syncing Hangfirejobqueue model data
	offset = 0
	for {
		var HangfirejobqueueData []model.Hangfirejobqueue
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobqueueData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirejobqueue")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobqueueData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobqueueData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirejobqueue")
			return err
		}
	}

	// Syncing Hangfirejobstate model data
	offset = 0
	for {
		var HangfirejobstateData []model.Hangfirejobstate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobstateData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirejobstate")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobstateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobstateData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirejobstate")
			return err
		}
	}

	// Syncing Hangfirelist model data
	offset = 0
	for {
		var HangfirelistData []model.Hangfirelist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirelistData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirelist")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirelistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirelistData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirelist")
			return err
		}
	}

	// Syncing Hangfireserver model data
	offset = 0
	for {
		var HangfireserverData []model.Hangfireserver
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfireserverData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfireserver")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfireserverData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfireserverData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfireserver")
			return err
		}
	}

	// Syncing Hangfireset model data
	offset = 0
	for {
		var HangfiresetData []model.Hangfireset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfiresetData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfireset")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfiresetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfiresetData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfireset")
			return err
		}
	}

	// Syncing Hangfirestate model data
	offset = 0
	for {
		var HangfirestateData []model.Hangfirestate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirestateData).Error; err != nil {
			log.Fatalf("Failed to query table: Hangfirestate")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirestateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirestateData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Hangfirestate")
			return err
		}
	}

	// Syncing Lcpcommandlog model data
	offset = 0
	for {
		var LcpcommandlogData []model.Lcpcommandlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpcommandlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpcommandlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpcommandlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpcommandlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpcommandlog")
			return err
		}
	}

	// Syncing Lcpconfig model data
	offset = 0
	for {
		var LcpconfigData []model.Lcpconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpconfigData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpconfig")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpconfigData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpconfig")
			return err
		}
	}

	// Syncing Lcpmaintainlog model data
	offset = 0
	for {
		var LcpmaintainlogData []model.Lcpmaintainlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpmaintainlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpmaintainlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpmaintainlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpmaintainlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpmaintainlog")
			return err
		}
	}

	// Syncing Lcpproduct model data
	offset = 0
	for {
		var LcpproductData []model.Lcpproduct
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpproductData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpproduct")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpproductData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpproductData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpproduct")
			return err
		}
	}

	// Syncing Lcprfidantenna model data
	offset = 0
	for {
		var LcprfidantennaData []model.Lcprfidantenna
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcprfidantennaData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcprfidantenna")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcprfidantennaData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcprfidantennaData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcprfidantenna")
			return err
		}
	}

	// Syncing Lcprfidreader model data
	offset = 0
	for {
		var LcprfidreaderData []model.Lcprfidreader
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcprfidreaderData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcprfidreader")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcprfidreaderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcprfidreaderData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcprfidreader")
			return err
		}
	}

	// Syncing Lcpsecuritygatebookaccesslog model data
	offset = 0
	for {
		var LcpsecuritygatebookaccesslogData []model.Lcpsecuritygatebookaccesslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatebookaccesslogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpsecuritygatebookaccesslog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatebookaccesslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatebookaccesslogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpsecuritygatebookaccesslog")
			return err
		}
	}

	// Syncing Lcpsecuritygatebookdailyaccess model data
	offset = 0
	for {
		var LcpsecuritygatebookdailyaccessData []model.Lcpsecuritygatebookdailyaccess
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatebookdailyaccessData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpsecuritygatebookdailyaccess")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatebookdailyaccessData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatebookdailyaccessData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpsecuritygatebookdailyaccess")
			return err
		}
	}

	// Syncing Lcpsecuritygateitemlog model data
	offset = 0
	for {
		var LcpsecuritygateitemlogData []model.Lcpsecuritygateitemlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygateitemlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpsecuritygateitemlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygateitemlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygateitemlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpsecuritygateitemlog")
			return err
		}
	}

	// Syncing Lcpsecuritygatepatronlog model data
	offset = 0
	for {
		var LcpsecuritygatepatronlogData []model.Lcpsecuritygatepatronlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatepatronlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpsecuritygatepatronlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatepatronlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatepatronlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpsecuritygatepatronlog")
			return err
		}
	}

	// Syncing Lcpserialport model data
	offset = 0
	for {
		var LcpserialportData []model.Lcpserialport
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserialportData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpserialport")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserialportData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserialportData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpserialport")
			return err
		}
	}

	// Syncing Lcpserialportext model data
	offset = 0
	for {
		var LcpserialportextData []model.Lcpserialportext
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserialportextData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpserialportext")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserialportextData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserialportextData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpserialportext")
			return err
		}
	}

	// Syncing Lcpservice model data
	offset = 0
	for {
		var LcpserviceData []model.Lcpservice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserviceData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpservice")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserviceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserviceData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpservice")
			return err
		}
	}

	// Syncing Lcpterminal model data
	offset = 0
	for {
		var LcpterminalData []model.Lcpterminal
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminal")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminal")
			return err
		}
	}

	// Syncing Lcpterminaladvertisement model data
	offset = 0
	for {
		var LcpterminaladvertisementData []model.Lcpterminaladvertisement
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaladvertisementData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminaladvertisement")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaladvertisementData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaladvertisementData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminaladvertisement")
			return err
		}
	}

	// Syncing Lcpterminalbox model data
	offset = 0
	for {
		var LcpterminalboxData []model.Lcpterminalbox
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalboxData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalbox")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalboxData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalboxData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalbox")
			return err
		}
	}

	// Syncing Lcpterminalboxitem model data
	offset = 0
	for {
		var LcpterminalboxitemData []model.Lcpterminalboxitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalboxitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalboxitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalboxitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalboxitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalboxitem")
			return err
		}
	}

	// Syncing Lcpterminaldevice model data
	offset = 0
	for {
		var LcpterminaldeviceData []model.Lcpterminaldevice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaldeviceData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminaldevice")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaldeviceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaldeviceData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminaldevice")
			return err
		}
	}

	// Syncing Lcpterminaldevicelog model data
	offset = 0
	for {
		var LcpterminaldevicelogData []model.Lcpterminaldevicelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaldevicelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminaldevicelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaldevicelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaldevicelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminaldevicelog")
			return err
		}
	}

	// Syncing Lcpterminallog model data
	offset = 0
	for {
		var LcpterminallogData []model.Lcpterminallog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminallogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminallog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminallogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminallogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminallog")
			return err
		}
	}

	// Syncing Lcpterminalpermission model data
	offset = 0
	for {
		var LcpterminalpermissionData []model.Lcpterminalpermission
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalpermissionData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalpermission")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalpermissionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalpermissionData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalpermission")
			return err
		}
	}

	// Syncing Lcpterminalshelf model data
	offset = 0
	for {
		var LcpterminalshelfData []model.Lcpterminalshelf
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelfData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalshelf")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelfData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelfData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalshelf")
			return err
		}
	}

	// Syncing Lcpterminalshelfitem model data
	offset = 0
	for {
		var LcpterminalshelfitemData []model.Lcpterminalshelfitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelfitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalshelfitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelfitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelfitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalshelfitem")
			return err
		}
	}

	// Syncing Lcpterminalshelflog model data
	offset = 0
	for {
		var LcpterminalshelflogData []model.Lcpterminalshelflog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelflogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpterminalshelflog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelflogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelflogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpterminalshelflog")
			return err
		}
	}

	// Syncing Lcpupgradelog model data
	offset = 0
	for {
		var LcpupgradelogData []model.Lcpupgradelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpupgradelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpupgradelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpupgradelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpupgradelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpupgradelog")
			return err
		}
	}

	// Syncing Lcpversion model data
	offset = 0
	for {
		var LcpversionData []model.Lcpversion
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpversionData).Error; err != nil {
			log.Fatalf("Failed to query table: Lcpversion")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpversionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpversionData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Lcpversion")
			return err
		}
	}

	// Syncing Libailibrarainbaseinfo model data
	offset = 0
	for {
		var LibailibrarainbaseinfoData []model.Libailibrarainbaseinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainbaseinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainbaseinfo")
			return err
		}
	}

	// Syncing Libailibrarainbaseinfoitem model data
	offset = 0
	for {
		var LibailibrarainbaseinfoitemData []model.Libailibrarainbaseinfoitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainbaseinfoitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainbaseinfoitem")
			return err
		}
	}

	// Syncing Libailibrarainbaseinfoprofile model data
	offset = 0
	for {
		var LibailibrarainbaseinfoprofileData []model.Libailibrarainbaseinfoprofile
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoprofileData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainbaseinfoprofile")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoprofileData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoprofileData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainbaseinfoprofile")
			return err
		}
	}

	// Syncing Libailibrarainknowledgefileinfo model data
	offset = 0
	for {
		var LibailibrarainknowledgefileinfoData []model.Libailibrarainknowledgefileinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainknowledgefileinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainknowledgefileinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainknowledgefileinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainknowledgefileinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainknowledgefileinfo")
			return err
		}
	}

	// Syncing Libailibrarainquestionmetric model data
	offset = 0
	for {
		var LibailibrarainquestionmetricData []model.Libailibrarainquestionmetric
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainquestionmetricData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainquestionmetric")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainquestionmetricData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainquestionmetricData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainquestionmetric")
			return err
		}
	}

	// Syncing Libailibrarainsessionmetric model data
	offset = 0
	for {
		var LibailibrarainsessionmetricData []model.Libailibrarainsessionmetric
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainsessionmetricData).Error; err != nil {
			log.Fatalf("Failed to query table: Libailibrarainsessionmetric")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainsessionmetricData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainsessionmetricData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libailibrarainsessionmetric")
			return err
		}
	}

	// Syncing Libainirobotinfo model data
	offset = 0
	for {
		var LibainirobotinfoData []model.Libainirobotinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibainirobotinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libainirobotinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibainirobotinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibainirobotinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libainirobotinfo")
			return err
		}
	}

	// Syncing Libbatchinfo model data
	offset = 0
	for {
		var LibbatchinfoData []model.Libbatchinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libbatchinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libbatchinfo")
			return err
		}
	}

	// Syncing Libbatchoperateindex model data
	offset = 0
	for {
		var LibbatchoperateindexData []model.Libbatchoperateindex
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchoperateindexData).Error; err != nil {
			log.Fatalf("Failed to query table: Libbatchoperateindex")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchoperateindexData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchoperateindexData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libbatchoperateindex")
			return err
		}
	}

	// Syncing Libbatchoperatelog model data
	offset = 0
	for {
		var LibbatchoperatelogData []model.Libbatchoperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchoperatelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libbatchoperatelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchoperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchoperatelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libbatchoperatelog")
			return err
		}
	}

	// Syncing Libbookinfo model data
	offset = 0
	for {
		var LibbookinfoData []model.Libbookinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbookinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libbookinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbookinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbookinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libbookinfo")
			return err
		}
	}

	// Syncing Libcirculatelog model data
	offset = 0
	for {
		var LibcirculatelogData []model.Libcirculatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibcirculatelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libcirculatelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibcirculatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibcirculatelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libcirculatelog")
			return err
		}
	}

	// Syncing LibcirculatelogBak model data
	offset = 0
	for {
		var LibcirculatelogBakData []model.LibcirculatelogBak
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibcirculatelogBakData).Error; err != nil {
			log.Fatalf("Failed to query table: LibcirculatelogBak")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibcirculatelogBakData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibcirculatelogBakData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: LibcirculatelogBak")
			return err
		}
	}

	// Syncing Libfeedback model data
	offset = 0
	for {
		var LibfeedbackData []model.Libfeedback
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibfeedbackData).Error; err != nil {
			log.Fatalf("Failed to query table: Libfeedback")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibfeedbackData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibfeedbackData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libfeedback")
			return err
		}
	}

	// Syncing Libfeelog model data
	offset = 0
	for {
		var LibfeelogData []model.Libfeelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibfeelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libfeelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibfeelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibfeelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libfeelog")
			return err
		}
	}

	// Syncing Libinventorystat model data
	offset = 0
	for {
		var LibinventorystatData []model.Libinventorystat
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventorystatData).Error; err != nil {
			log.Fatalf("Failed to query table: Libinventorystat")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventorystatData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventorystatData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libinventorystat")
			return err
		}
	}

	// Syncing Libinventorytask model data
	offset = 0
	for {
		var LibinventorytaskData []model.Libinventorytask
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventorytaskData).Error; err != nil {
			log.Fatalf("Failed to query table: Libinventorytask")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventorytaskData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventorytaskData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libinventorytask")
			return err
		}
	}

	// Syncing Libinventorywork model data
	offset = 0
	for {
		var LibinventoryworkData []model.Libinventorywork
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworkData).Error; err != nil {
			log.Fatalf("Failed to query table: Libinventorywork")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworkData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworkData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libinventorywork")
			return err
		}
	}

	// Syncing Libinventoryworkdetail model data
	offset = 0
	for {
		var LibinventoryworkdetailData []model.Libinventoryworkdetail
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworkdetailData).Error; err != nil {
			log.Fatalf("Failed to query table: Libinventoryworkdetail")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworkdetailData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworkdetailData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libinventoryworkdetail")
			return err
		}
	}

	// Syncing Libinventoryworklog model data
	offset = 0
	for {
		var LibinventoryworklogData []model.Libinventoryworklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworklogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libinventoryworklog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworklogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libinventoryworklog")
			return err
		}
	}

	// Syncing Libitem model data
	offset = 0
	for {
		var LibitemData []model.Libitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Libitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libitem")
			return err
		}
	}

	// Syncing LibitemCopy model data
	offset = 0
	for {
		var LibitemCopyData []model.LibitemCopy
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemCopyData).Error; err != nil {
			log.Fatalf("Failed to query table: LibitemCopy")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemCopyData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemCopyData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: LibitemCopy")
			return err
		}
	}

	// Syncing Libiteminventoryinfo model data
	offset = 0
	for {
		var LibiteminventoryinfoData []model.Libiteminventoryinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibiteminventoryinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libiteminventoryinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibiteminventoryinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibiteminventoryinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libiteminventoryinfo")
			return err
		}
	}

	// Syncing Libiteminventorylog model data
	offset = 0
	for {
		var LibiteminventorylogData []model.Libiteminventorylog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibiteminventorylogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libiteminventorylog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibiteminventorylogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibiteminventorylogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libiteminventorylog")
			return err
		}
	}

	// Syncing Libitemlocinfo model data
	offset = 0
	for {
		var LibitemlocinfoData []model.Libitemlocinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemlocinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libitemlocinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemlocinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemlocinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libitemlocinfo")
			return err
		}
	}

	// Syncing Libitemoperateindexlog model data
	offset = 0
	for {
		var LibitemoperateindexlogData []model.Libitemoperateindexlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemoperateindexlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libitemoperateindexlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemoperateindexlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemoperateindexlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libitemoperateindexlog")
			return err
		}
	}

	// Syncing Libitemoperatelog model data
	offset = 0
	for {
		var LibitemoperatelogData []model.Libitemoperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemoperatelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libitemoperatelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemoperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemoperatelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libitemoperatelog")
			return err
		}
	}

	// Syncing Libjournalinfo model data
	offset = 0
	for {
		var LibjournalinfoData []model.Libjournalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibjournalinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Libjournalinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibjournalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibjournalinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libjournalinfo")
			return err
		}
	}

	// Syncing Liblabel model data
	offset = 0
	for {
		var LiblabelData []model.Liblabel
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblabelData).Error; err != nil {
			log.Fatalf("Failed to query table: Liblabel")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblabelData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblabelData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Liblabel")
			return err
		}
	}

	// Syncing Liblabeloperatelog model data
	offset = 0
	for {
		var LiblabeloperatelogData []model.Liblabeloperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblabeloperatelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Liblabeloperatelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblabeloperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblabeloperatelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Liblabeloperatelog")
			return err
		}
	}

	// Syncing Liblayer model data
	offset = 0
	for {
		var LiblayerData []model.Liblayer
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblayerData).Error; err != nil {
			log.Fatalf("Failed to query table: Liblayer")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblayerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblayerData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Liblayer")
			return err
		}
	}

	// Syncing Liblayerindexupdatelog model data
	offset = 0
	for {
		var LiblayerindexupdatelogData []model.Liblayerindexupdatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblayerindexupdatelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Liblayerindexupdatelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblayerindexupdatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblayerindexupdatelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Liblayerindexupdatelog")
			return err
		}
	}

	// Syncing Libnotificationlog model data
	offset = 0
	for {
		var LibnotificationlogData []model.Libnotificationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibnotificationlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libnotificationlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibnotificationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibnotificationlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libnotificationlog")
			return err
		}
	}

	// Syncing Libpartonreservation model data
	offset = 0
	for {
		var LibpartonreservationData []model.Libpartonreservation
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpartonreservationData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpartonreservation")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpartonreservationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpartonreservationData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpartonreservation")
			return err
		}
	}

	// Syncing Libpatron model data
	offset = 0
	for {
		var LibpatronData []model.Libpatron
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpatron")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpatron")
			return err
		}
	}

	// Syncing Libpatronitem model data
	offset = 0
	for {
		var LibpatronitemData []model.Libpatronitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpatronitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpatronitem")
			return err
		}
	}

	// Syncing Libpatronlog model data
	offset = 0
	for {
		var LibpatronlogData []model.Libpatronlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpatronlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpatronlog")
			return err
		}
	}

	// Syncing Libpointsclearing model data
	offset = 0
	for {
		var LibpointsclearingData []model.Libpointsclearing
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpointsclearingData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpointsclearing")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpointsclearingData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpointsclearingData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpointsclearing")
			return err
		}
	}

	// Syncing Libpointslog model data
	offset = 0
	for {
		var LibpointslogData []model.Libpointslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpointslogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libpointslog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpointslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpointslogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libpointslog")
			return err
		}
	}

	// Syncing Librfidantlayer model data
	offset = 0
	for {
		var LibrfidantlayerData []model.Librfidantlayer
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidantlayerData).Error; err != nil {
			log.Fatalf("Failed to query table: Librfidantlayer")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidantlayerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidantlayerData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Librfidantlayer")
			return err
		}
	}

	// Syncing Librfidscandetaillog model data
	offset = 0
	for {
		var LibrfidscandetaillogData []model.Librfidscandetaillog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidscandetaillogData).Error; err != nil {
			log.Fatalf("Failed to query table: Librfidscandetaillog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidscandetaillogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidscandetaillogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Librfidscandetaillog")
			return err
		}
	}

	// Syncing Librfidscanlog model data
	offset = 0
	for {
		var LibrfidscanlogData []model.Librfidscanlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidscanlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Librfidscanlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidscanlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidscanlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Librfidscanlog")
			return err
		}
	}

	// Syncing Librow model data
	offset = 0
	for {
		var LibrowData []model.Librow
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrowData).Error; err != nil {
			log.Fatalf("Failed to query table: Librow")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrowData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrowData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Librow")
			return err
		}
	}

	// Syncing Librowcatalog model data
	offset = 0
	for {
		var LibrowcatalogData []model.Librowcatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrowcatalogData).Error; err != nil {
			log.Fatalf("Failed to query table: Librowcatalog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrowcatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrowcatalogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Librowcatalog")
			return err
		}
	}

	// Syncing Libscanitemlog model data
	offset = 0
	for {
		var LibscanitemlogData []model.Libscanitemlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibscanitemlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libscanitemlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibscanitemlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibscanitemlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libscanitemlog")
			return err
		}
	}

	// Syncing Libshelf model data
	offset = 0
	for {
		var LibshelfData []model.Libshelf
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibshelfData).Error; err != nil {
			log.Fatalf("Failed to query table: Libshelf")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibshelfData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibshelfData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libshelf")
			return err
		}
	}

	// Syncing Libstruct model data
	offset = 0
	for {
		var LibstructData []model.Libstruct
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibstructData).Error; err != nil {
			log.Fatalf("Failed to query table: Libstruct")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibstructData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibstructData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libstruct")
			return err
		}
	}

	// Syncing Libtagtobarcodelog model data
	offset = 0
	for {
		var LibtagtobarcodelogData []model.Libtagtobarcodelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtagtobarcodelogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libtagtobarcodelog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtagtobarcodelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtagtobarcodelogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libtagtobarcodelog")
			return err
		}
	}

	// Syncing Libtaskitem model data
	offset = 0
	for {
		var LibtaskitemData []model.Libtaskitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtaskitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Libtaskitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtaskitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtaskitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libtaskitem")
			return err
		}
	}

	// Syncing Libtaskpackage model data
	offset = 0
	for {
		var LibtaskpackageData []model.Libtaskpackage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtaskpackageData).Error; err != nil {
			log.Fatalf("Failed to query table: Libtaskpackage")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtaskpackageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtaskpackageData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libtaskpackage")
			return err
		}
	}

	// Syncing Libupdatefirstbooklog model data
	offset = 0
	for {
		var LibupdatefirstbooklogData []model.Libupdatefirstbooklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibupdatefirstbooklogData).Error; err != nil {
			log.Fatalf("Failed to query table: Libupdatefirstbooklog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(LibupdatefirstbooklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibupdatefirstbooklogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Libupdatefirstbooklog")
			return err
		}
	}

	// Syncing Misactivity model data
	offset = 0
	for {
		var MisactivityData []model.Misactivity
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MisactivityData).Error; err != nil {
			log.Fatalf("Failed to query table: Misactivity")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(MisactivityData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MisactivityData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Misactivity")
			return err
		}
	}

	// Syncing Mismediainfo model data
	offset = 0
	for {
		var MismediainfoData []model.Mismediainfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MismediainfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Mismediainfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(MismediainfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MismediainfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Mismediainfo")
			return err
		}
	}

	// Syncing Misnews model data
	offset = 0
	for {
		var MisnewsData []model.Misnews
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MisnewsData).Error; err != nil {
			log.Fatalf("Failed to query table: Misnews")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(MisnewsData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MisnewsData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Misnews")
			return err
		}
	}

	// Syncing Mistemplate model data
	offset = 0
	for {
		var MistemplateData []model.Mistemplate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MistemplateData).Error; err != nil {
			log.Fatalf("Failed to query table: Mistemplate")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(MistemplateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MistemplateData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Mistemplate")
			return err
		}
	}

	// Syncing Rescatalog model data
	offset = 0
	for {
		var RescatalogData []model.Rescatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RescatalogData).Error; err != nil {
			log.Fatalf("Failed to query table: Rescatalog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(RescatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RescatalogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Rescatalog")
			return err
		}
	}

	// Syncing Rescipinfo model data
	offset = 0
	for {
		var RescipinfoData []model.Rescipinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RescipinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Rescipinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(RescipinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RescipinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Rescipinfo")
			return err
		}
	}

	// Syncing Resfourcorner model data
	offset = 0
	for {
		var ResfourcornerData []model.Resfourcorner
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResfourcornerData).Error; err != nil {
			log.Fatalf("Failed to query table: Resfourcorner")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ResfourcornerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResfourcornerData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Resfourcorner")
			return err
		}
	}

	// Syncing Resjournalinfo model data
	offset = 0
	for {
		var ResjournalinfoData []model.Resjournalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResjournalinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Resjournalinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ResjournalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResjournalinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Resjournalinfo")
			return err
		}
	}

	// Syncing Resnotfound model data
	offset = 0
	for {
		var ResnotfoundData []model.Resnotfound
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResnotfoundData).Error; err != nil {
			log.Fatalf("Failed to query table: Resnotfound")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(ResnotfoundData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResnotfoundData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Resnotfound")
			return err
		}
	}

	// Syncing Respublisherinfo model data
	offset = 0
	for {
		var RespublisherinfoData []model.Respublisherinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RespublisherinfoData).Error; err != nil {
			log.Fatalf("Failed to query table: Respublisherinfo")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(RespublisherinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RespublisherinfoData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Respublisherinfo")
			return err
		}
	}

	// Syncing Ssbackgroundjob model data
	offset = 0
	for {
		var SsbackgroundjobData []model.Ssbackgroundjob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SsbackgroundjobData).Error; err != nil {
			log.Fatalf("Failed to query table: Ssbackgroundjob")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SsbackgroundjobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SsbackgroundjobData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Ssbackgroundjob")
			return err
		}
	}

	// Syncing Sysattachment model data
	offset = 0
	for {
		var SysattachmentData []model.Sysattachment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysattachmentData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysattachment")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysattachmentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysattachmentData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysattachment")
			return err
		}
	}

	// Syncing Sysauditacslog model data
	offset = 0
	for {
		var SysauditacslogData []model.Sysauditacslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditacslogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditacslog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditacslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditacslogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditacslog")
			return err
		}
	}

	// Syncing Sysauditapilog model data
	offset = 0
	for {
		var SysauditapilogData []model.Sysauditapilog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditapilogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditapilog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditapilogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditapilogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditapilog")
			return err
		}
	}

	// Syncing Sysauditapplog model data
	offset = 0
	for {
		var SysauditapplogData []model.Sysauditapplog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditapplogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditapplog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditapplogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditapplogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditapplog")
			return err
		}
	}

	// Syncing Sysauditlinklog model data
	offset = 0
	for {
		var SysauditlinklogData []model.Sysauditlinklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditlinklogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditlinklog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditlinklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditlinklogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditlinklog")
			return err
		}
	}

	// Syncing Sysauditlmslog model data
	offset = 0
	for {
		var SysauditlmslogData []model.Sysauditlmslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditlmslogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditlmslog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditlmslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditlmslogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditlmslog")
			return err
		}
	}

	// Syncing Sysauditsslog model data
	offset = 0
	for {
		var SysauditsslogData []model.Sysauditsslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditsslogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysauditsslog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditsslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditsslogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysauditsslog")
			return err
		}
	}

	// Syncing Sysblocklist model data
	offset = 0
	for {
		var SysblocklistData []model.Sysblocklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysblocklistData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysblocklist")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysblocklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysblocklistData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysblocklist")
			return err
		}
	}

	// Syncing Sysbookblocklist model data
	offset = 0
	for {
		var SysbookblocklistData []model.Sysbookblocklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbookblocklistData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysbookblocklist")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbookblocklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbookblocklistData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysbookblocklist")
			return err
		}
	}

	// Syncing Sysbooknumlib model data
	offset = 0
	for {
		var SysbooknumlibData []model.Sysbooknumlib
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbooknumlibData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysbooknumlib")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbooknumlibData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbooknumlibData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysbooknumlib")
			return err
		}
	}

	// Syncing Sysbooknumset model data
	offset = 0
	for {
		var SysbooknumsetData []model.Sysbooknumset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbooknumsetData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysbooknumset")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbooknumsetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbooknumsetData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysbooknumset")
			return err
		}
	}

	// Syncing Syscardconfig model data
	offset = 0
	for {
		var SyscardconfigData []model.Syscardconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscardconfigData).Error; err != nil {
			log.Fatalf("Failed to query table: Syscardconfig")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscardconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscardconfigData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syscardconfig")
			return err
		}
	}

	// Syncing Syscarddevdtl model data
	offset = 0
	for {
		var SyscarddevdtlData []model.Syscarddevdtl
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscarddevdtlData).Error; err != nil {
			log.Fatalf("Failed to query table: Syscarddevdtl")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscarddevdtlData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscarddevdtlData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syscarddevdtl")
			return err
		}
	}

	// Syncing Syscardtype model data
	offset = 0
	for {
		var SyscardtypeData []model.Syscardtype
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscardtypeData).Error; err != nil {
			log.Fatalf("Failed to query table: Syscardtype")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscardtypeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscardtypeData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syscardtype")
			return err
		}
	}

	// Syncing Syscoderule model data
	offset = 0
	for {
		var SyscoderuleData []model.Syscoderule
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscoderuleData).Error; err != nil {
			log.Fatalf("Failed to query table: Syscoderule")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscoderuleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscoderuleData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syscoderule")
			return err
		}
	}

	// Syncing Syscoderuleseed model data
	offset = 0
	for {
		var SyscoderuleseedData []model.Syscoderuleseed
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscoderuleseedData).Error; err != nil {
			log.Fatalf("Failed to query table: Syscoderuleseed")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscoderuleseedData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscoderuleseedData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syscoderuleseed")
			return err
		}
	}

	// Syncing Sysconfigbase model data
	offset = 0
	for {
		var SysconfigbaseData []model.Sysconfigbase
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysconfigbaseData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysconfigbase")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysconfigbaseData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysconfigbaseData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysconfigbase")
			return err
		}
	}

	// Syncing Sysconfiglog model data
	offset = 0
	for {
		var SysconfiglogData []model.Sysconfiglog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysconfiglogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysconfiglog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysconfiglogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysconfiglogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysconfiglog")
			return err
		}
	}

	// Syncing Sysdataitem model data
	offset = 0
	for {
		var SysdataitemData []model.Sysdataitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdataitemData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysdataitem")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdataitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdataitemData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysdataitem")
			return err
		}
	}

	// Syncing Sysdataitemdetail model data
	offset = 0
	for {
		var SysdataitemdetailData []model.Sysdataitemdetail
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdataitemdetailData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysdataitemdetail")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdataitemdetailData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdataitemdetailData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysdataitemdetail")
			return err
		}
	}

	// Syncing Sysdatalog model data
	offset = 0
	for {
		var SysdatalogData []model.Sysdatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdatalogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysdatalog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdatalogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysdatalog")
			return err
		}
	}

	// Syncing Sysdepartment model data
	offset = 0
	for {
		var SysdepartmentData []model.Sysdepartment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdepartmentData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysdepartment")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdepartmentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdepartmentData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysdepartment")
			return err
		}
	}

	// Syncing Sysenumfield model data
	offset = 0
	for {
		var SysenumfieldData []model.Sysenumfield
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysenumfieldData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysenumfield")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysenumfieldData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysenumfieldData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysenumfield")
			return err
		}
	}

	// Syncing Sysenumvalue model data
	offset = 0
	for {
		var SysenumvalueData []model.Sysenumvalue
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysenumvalueData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysenumvalue")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysenumvalueData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysenumvalueData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysenumvalue")
			return err
		}
	}

	// Syncing Sysenumvalue2 model data
	offset = 0
	for {
		var Sysenumvalue2Data []model.Sysenumvalue2
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&Sysenumvalue2Data).Error; err != nil {
			log.Fatalf("Failed to query table: Sysenumvalue2")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(Sysenumvalue2Data) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(Sysenumvalue2Data, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysenumvalue2")
			return err
		}
	}

	// Syncing Sysfaceoffineoperationlog model data
	offset = 0
	for {
		var SysfaceoffineoperationlogData []model.Sysfaceoffineoperationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceoffineoperationlogData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysfaceoffineoperationlog")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceoffineoperationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceoffineoperationlogData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysfaceoffineoperationlog")
			return err
		}
	}

	// Syncing Sysfaceofflinefeature model data
	offset = 0
	for {
		var SysfaceofflinefeatureData []model.Sysfaceofflinefeature
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflinefeatureData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysfaceofflinefeature")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflinefeatureData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflinefeatureData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysfaceofflinefeature")
			return err
		}
	}

	// Syncing Sysfaceofflinegroup model data
	offset = 0
	for {
		var SysfaceofflinegroupData []model.Sysfaceofflinegroup
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflinegroupData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysfaceofflinegroup")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflinegroupData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflinegroupData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysfaceofflinegroup")
			return err
		}
	}

	// Syncing Sysfaceofflineuser model data
	offset = 0
	for {
		var SysfaceofflineuserData []model.Sysfaceofflineuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflineuserData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysfaceofflineuser")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflineuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflineuserData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysfaceofflineuser")
			return err
		}
	}

	// Syncing Syslanguage model data
	offset = 0
	for {
		var SyslanguageData []model.Syslanguage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslanguageData).Error; err != nil {
			log.Fatalf("Failed to query table: Syslanguage")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslanguageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslanguageData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syslanguage")
			return err
		}
	}

	// Syncing Syslayertran model data
	offset = 0
	for {
		var SyslayertranData []model.Syslayertran
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslayertranData).Error; err != nil {
			log.Fatalf("Failed to query table: Syslayertran")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslayertranData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslayertranData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syslayertran")
			return err
		}
	}

	// Syncing Syslocation model data
	offset = 0
	for {
		var SyslocationData []model.Syslocation
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslocationData).Error; err != nil {
			log.Fatalf("Failed to query table: Syslocation")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslocationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslocationData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Syslocation")
			return err
		}
	}

	// Syncing Sysmenu model data
	offset = 0
	for {
		var SysmenuData []model.Sysmenu
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysmenuData).Error; err != nil {
			log.Fatalf("Failed to query table: Sysmenu")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysmenuData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysmenuData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Sysmenu")
			return err
		}
	}

	// Syncing SysmenuCopy model data
	offset = 0
	for {
		var SysmenuCopyData []model.SysmenuCopy
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysmenuCopyData).Error; err != nil {
			log.Fatalf("Failed to query table: SysmenuCopy")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SysmenuCopyData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysmenuCopyData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: SysmenuCopy")
			return err
		}
	}

	// Syncing Systasklist model data
	offset = 0
	for {
		var SystasklistData []model.Systasklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystasklistData).Error; err != nil {
			log.Fatalf("Failed to query table: Systasklist")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SystasklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystasklistData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Systasklist")
			return err
		}
	}

	// Syncing Systenantconfig model data
	offset = 0
	for {
		var SystenantconfigData []model.Systenantconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystenantconfigData).Error; err != nil {
			log.Fatalf("Failed to query table: Systenantconfig")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SystenantconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystenantconfigData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Systenantconfig")
			return err
		}
	}

	// Syncing Systenantextend model data
	offset = 0
	for {
		var SystenantextendData []model.Systenantextend
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystenantextendData).Error; err != nil {
			log.Fatalf("Failed to query table: Systenantextend")
			return err
		}

		// 如果没有更多数据，退出循环
		if len(SystenantextendData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystenantextendData, batchSize).Error; err != nil {
			log.Fatalf("Failed to insert table: Systenantextend")
			return err
		}
	}

	return nil
}
