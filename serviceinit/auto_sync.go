package serviceinit

import (
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
)

func CusSyncDatabase() error {
	batchSize := 1000 // 每批次同步的数据量
	var offset int    // 用于分页
	// Syncing Abpauditlog model data

	for {
		var AbpauditlogData []model.Abpauditlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpauditlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpauditlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpauditlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpbackgroundjob model data

	for {
		var AbpbackgroundjobData []model.Abpbackgroundjob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpbackgroundjobData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpbackgroundjobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpbackgroundjobData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpedition model data

	for {
		var AbpeditionData []model.Abpedition
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpeditionData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpeditionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpeditionData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitychange model data

	for {
		var AbpentitychangeData []model.Abpentitychange
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitychangeData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitychangeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitychangeData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitychangeset model data

	for {
		var AbpentitychangesetData []model.Abpentitychangeset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitychangesetData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitychangesetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitychangesetData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitypropertychange model data

	for {
		var AbpentitypropertychangeData []model.Abpentitypropertychange
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpentitypropertychangeData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpentitypropertychangeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpentitypropertychangeData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpfeature model data

	for {
		var AbpfeatureData []model.Abpfeature
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpfeatureData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpfeatureData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpfeatureData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abplanguage model data

	for {
		var AbplanguageData []model.Abplanguage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbplanguageData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbplanguageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbplanguageData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abplanguagetext model data

	for {
		var AbplanguagetextData []model.Abplanguagetext
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbplanguagetextData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbplanguagetextData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbplanguagetextData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpnotification model data

	for {
		var AbpnotificationData []model.Abpnotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpnotificationData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpnotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpnotificationData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpnotificationsubscription model data

	for {
		var AbpnotificationsubscriptionData []model.Abpnotificationsubscription
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpnotificationsubscriptionData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpnotificationsubscriptionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpnotificationsubscriptionData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abporganizationunitrole model data

	for {
		var AbporganizationunitroleData []model.Abporganizationunitrole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbporganizationunitroleData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbporganizationunitroleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbporganizationunitroleData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abporganizationunit model data

	for {
		var AbporganizationunitData []model.Abporganizationunit
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbporganizationunitData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbporganizationunitData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbporganizationunitData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abppermission model data

	for {
		var AbppermissionData []model.Abppermission
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbppermissionData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbppermissionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbppermissionData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abppersistedgrant model data

	for {
		var AbppersistedgrantData []model.Abppersistedgrant
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbppersistedgrantData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbppersistedgrantData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbppersistedgrantData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abproleclaim model data

	for {
		var AbproleclaimData []model.Abproleclaim
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbproleclaimData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbproleclaimData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbproleclaimData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abprole model data

	for {
		var AbproleData []model.Abprole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbproleData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbproleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbproleData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpsetting model data

	for {
		var AbpsettingData []model.Abpsetting
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpsettingData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpsettingData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpsettingData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abptenantnotification model data

	for {
		var AbptenantnotificationData []model.Abptenantnotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbptenantnotificationData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbptenantnotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbptenantnotificationData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abptenant model data

	for {
		var AbptenantData []model.Abptenant
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbptenantData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbptenantData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbptenantData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuseraccount model data

	for {
		var AbpuseraccountData []model.Abpuseraccount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuseraccountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuseraccountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuseraccountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserclaim model data

	for {
		var AbpuserclaimData []model.Abpuserclaim
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserclaimData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserclaimData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserclaimData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserloginattempt model data

	for {
		var AbpuserloginattemptData []model.Abpuserloginattempt
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserloginattemptData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserloginattemptData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserloginattemptData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserlogin model data

	for {
		var AbpuserloginData []model.Abpuserlogin
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserloginData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserloginData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserloginData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpusernotification model data

	for {
		var AbpusernotificationData []model.Abpusernotification
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpusernotificationData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpusernotificationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpusernotificationData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserorganizationunit model data

	for {
		var AbpuserorganizationunitData []model.Abpuserorganizationunit
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserorganizationunitData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserorganizationunitData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserorganizationunitData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserrole model data

	for {
		var AbpuserroleData []model.Abpuserrole
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserroleData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserroleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserroleData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuser model data

	for {
		var AbpuserData []model.Abpuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpuserData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpuserData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpusertoken model data

	for {
		var AbpusertokenData []model.Abpusertoken
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AbpusertokenData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AbpusertokenData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AbpusertokenData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appaliuser model data

	for {
		var AppaliuserData []model.Appaliuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppaliuserData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppaliuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppaliuserData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appapprovalinfo model data

	for {
		var AppapprovalinfoData []model.Appapprovalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppapprovalinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppapprovalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppapprovalinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appapprovaltemplate model data

	for {
		var AppapprovaltemplateData []model.Appapprovaltemplate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppapprovaltemplateData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppapprovaltemplateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppapprovaltemplateData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appbinaryobject model data

	for {
		var AppbinaryobjectData []model.Appbinaryobject
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppbinaryobjectData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppbinaryobjectData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppbinaryobjectData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appbookorder model data

	for {
		var AppbookorderData []model.Appbookorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppbookorderData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppbookorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppbookorderData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appchatmessage model data

	for {
		var AppchatmessageData []model.Appchatmessage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppchatmessageData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppchatmessageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppchatmessageData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appcreditloginorder model data

	for {
		var AppcreditloginorderData []model.Appcreditloginorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppcreditloginorderData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppcreditloginorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppcreditloginorderData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appfriendship model data

	for {
		var AppfriendshipData []model.Appfriendship
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppfriendshipData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppfriendshipData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppfriendshipData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appinvoice model data

	for {
		var AppinvoiceData []model.Appinvoice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppinvoiceData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppinvoiceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppinvoiceData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appitemlocked model data

	for {
		var AppitemlockedData []model.Appitemlocked
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppitemlockedData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppitemlockedData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppitemlockedData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appmessageboard model data

	for {
		var AppmessageboardData []model.Appmessageboard
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppmessageboardData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppmessageboardData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppmessageboardData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appnotificationlog model data

	for {
		var AppnotificationlogData []model.Appnotificationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppnotificationlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppnotificationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppnotificationlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Apppayorder model data

	for {
		var ApppayorderData []model.Apppayorder
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApppayorderData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ApppayorderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApppayorderData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Apppickupcode model data

	for {
		var ApppickupcodeData []model.Apppickupcode
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApppickupcodeData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ApppickupcodeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApppickupcodeData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appqrcode model data

	for {
		var AppqrcodeData []model.Appqrcode
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppqrcodeData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppqrcodeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppqrcodeData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Apprecommendinfo model data

	for {
		var ApprecommendinfoData []model.Apprecommendinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ApprecommendinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ApprecommendinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ApprecommendinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appsubbookinfo model data

	for {
		var AppsubbookinfoData []model.Appsubbookinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppsubbookinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppsubbookinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppsubbookinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appsubscriptionpayment model data

	for {
		var AppsubscriptionpaymentData []model.Appsubscriptionpayment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppsubscriptionpaymentData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppsubscriptionpaymentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppsubscriptionpaymentData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appusercard model data

	for {
		var AppusercardData []model.Appusercard
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppusercardData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppusercardData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppusercardData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Appweuser model data

	for {
		var AppweuserData []model.Appweuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&AppweuserData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(AppweuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(AppweuserData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasbusinesscount model data

	for {
		var DasbusinesscountData []model.Dasbusinesscount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasbusinesscountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasbusinesscountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasbusinesscountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dascirculatecount model data

	for {
		var DascirculatecountData []model.Dascirculatecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DascirculatecountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DascirculatecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DascirculatecountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasdatabaselink model data

	for {
		var DasdatabaselinkData []model.Dasdatabaselink
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasdatabaselinkData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasdatabaselinkData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasdatabaselinkData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasdatasource model data

	for {
		var DasdatasourceData []model.Dasdatasource
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasdatasourceData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasdatasourceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasdatasourceData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasfeecount model data

	for {
		var DasfeecountData []model.Dasfeecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasfeecountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasfeecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasfeecountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Daspatronlogcount model data

	for {
		var DaspatronlogcountData []model.Daspatronlogcount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DaspatronlogcountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DaspatronlogcountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DaspatronlogcountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasperformance model data

	for {
		var DasperformanceData []model.Dasperformance
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasperformanceData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasperformanceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasperformanceData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dassecuritygatecount model data

	for {
		var DassecuritygatecountData []model.Dassecuritygatecount
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DassecuritygatecountData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DassecuritygatecountData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DassecuritygatecountData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasvisitpage model data

	for {
		var DasvisitpageData []model.Dasvisitpage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasvisitpageData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasvisitpageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasvisitpageData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasvisittrend model data

	for {
		var DasvisittrendData []model.Dasvisittrend
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&DasvisittrendData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(DasvisittrendData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(DasvisittrendData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Efmigrationshistory model data

	for {
		var EfmigrationshistoryData []model.Efmigrationshistory
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&EfmigrationshistoryData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(EfmigrationshistoryData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(EfmigrationshistoryData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireaggregatedcounter model data

	for {
		var HangfireaggregatedcounterData []model.Hangfireaggregatedcounter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfireaggregatedcounterData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfireaggregatedcounterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfireaggregatedcounterData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirecounter model data

	for {
		var HangfirecounterData []model.Hangfirecounter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirecounterData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirecounterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirecounterData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfiredistributedlock model data

	for {
		var HangfiredistributedlockData []model.Hangfiredistributedlock
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfiredistributedlockData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfiredistributedlockData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfiredistributedlockData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirehash model data

	for {
		var HangfirehashData []model.Hangfirehash
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirehashData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirehashData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirehashData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejob model data

	for {
		var HangfirejobData []model.Hangfirejob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobparameter model data

	for {
		var HangfirejobparameterData []model.Hangfirejobparameter
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobparameterData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobparameterData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobparameterData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobqueue model data

	for {
		var HangfirejobqueueData []model.Hangfirejobqueue
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobqueueData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobqueueData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobqueueData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobstate model data

	for {
		var HangfirejobstateData []model.Hangfirejobstate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirejobstateData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirejobstateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirejobstateData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirelist model data

	for {
		var HangfirelistData []model.Hangfirelist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirelistData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirelistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirelistData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireserver model data

	for {
		var HangfireserverData []model.Hangfireserver
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfireserverData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfireserverData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfireserverData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireset model data

	for {
		var HangfiresetData []model.Hangfireset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfiresetData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfiresetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfiresetData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirestate model data

	for {
		var HangfirestateData []model.Hangfirestate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&HangfirestateData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(HangfirestateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(HangfirestateData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpcommandlog model data

	for {
		var LcpcommandlogData []model.Lcpcommandlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpcommandlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpcommandlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpcommandlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpconfig model data

	for {
		var LcpconfigData []model.Lcpconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpconfigData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpconfigData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpmaintainlog model data

	for {
		var LcpmaintainlogData []model.Lcpmaintainlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpmaintainlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpmaintainlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpmaintainlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpproduct model data

	for {
		var LcpproductData []model.Lcpproduct
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpproductData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpproductData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpproductData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcprfidantenna model data

	for {
		var LcprfidantennaData []model.Lcprfidantenna
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcprfidantennaData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcprfidantennaData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcprfidantennaData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcprfidreader model data

	for {
		var LcprfidreaderData []model.Lcprfidreader
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcprfidreaderData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcprfidreaderData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcprfidreaderData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatebookaccesslog model data

	for {
		var LcpsecuritygatebookaccesslogData []model.Lcpsecuritygatebookaccesslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatebookaccesslogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatebookaccesslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatebookaccesslogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatebookdailyaccess model data

	for {
		var LcpsecuritygatebookdailyaccessData []model.Lcpsecuritygatebookdailyaccess
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatebookdailyaccessData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatebookdailyaccessData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatebookdailyaccessData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygateitemlog model data

	for {
		var LcpsecuritygateitemlogData []model.Lcpsecuritygateitemlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygateitemlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygateitemlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygateitemlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatepatronlog model data

	for {
		var LcpsecuritygatepatronlogData []model.Lcpsecuritygatepatronlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpsecuritygatepatronlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpsecuritygatepatronlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpsecuritygatepatronlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpserialport model data

	for {
		var LcpserialportData []model.Lcpserialport
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserialportData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserialportData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserialportData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpserialportext model data

	for {
		var LcpserialportextData []model.Lcpserialportext
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserialportextData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserialportextData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserialportextData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpservice model data

	for {
		var LcpserviceData []model.Lcpservice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpserviceData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpserviceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpserviceData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminal model data

	for {
		var LcpterminalData []model.Lcpterminal
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaladvertisement model data

	for {
		var LcpterminaladvertisementData []model.Lcpterminaladvertisement
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaladvertisementData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaladvertisementData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaladvertisementData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalbox model data

	for {
		var LcpterminalboxData []model.Lcpterminalbox
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalboxData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalboxData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalboxData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalboxitem model data

	for {
		var LcpterminalboxitemData []model.Lcpterminalboxitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalboxitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalboxitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalboxitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaldevice model data

	for {
		var LcpterminaldeviceData []model.Lcpterminaldevice
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaldeviceData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaldeviceData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaldeviceData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaldevicelog model data

	for {
		var LcpterminaldevicelogData []model.Lcpterminaldevicelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminaldevicelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminaldevicelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminaldevicelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminallog model data

	for {
		var LcpterminallogData []model.Lcpterminallog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminallogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminallogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminallogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalpermission model data

	for {
		var LcpterminalpermissionData []model.Lcpterminalpermission
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalpermissionData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalpermissionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalpermissionData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelf model data

	for {
		var LcpterminalshelfData []model.Lcpterminalshelf
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelfData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelfData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelfData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelfitem model data

	for {
		var LcpterminalshelfitemData []model.Lcpterminalshelfitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelfitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelfitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelfitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelflog model data

	for {
		var LcpterminalshelflogData []model.Lcpterminalshelflog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpterminalshelflogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpterminalshelflogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpterminalshelflogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpupgradelog model data

	for {
		var LcpupgradelogData []model.Lcpupgradelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpupgradelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpupgradelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpupgradelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpversion model data

	for {
		var LcpversionData []model.Lcpversion
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LcpversionData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LcpversionData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LcpversionData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfo model data

	for {
		var LibailibrarainbaseinfoData []model.Libailibrarainbaseinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfoitem model data

	for {
		var LibailibrarainbaseinfoitemData []model.Libailibrarainbaseinfoitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfoprofile model data

	for {
		var LibailibrarainbaseinfoprofileData []model.Libailibrarainbaseinfoprofile
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainbaseinfoprofileData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainbaseinfoprofileData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainbaseinfoprofileData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainknowledgefileinfo model data

	for {
		var LibailibrarainknowledgefileinfoData []model.Libailibrarainknowledgefileinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainknowledgefileinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainknowledgefileinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainknowledgefileinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainquestionmetric model data

	for {
		var LibailibrarainquestionmetricData []model.Libailibrarainquestionmetric
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainquestionmetricData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainquestionmetricData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainquestionmetricData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainsessionmetric model data

	for {
		var LibailibrarainsessionmetricData []model.Libailibrarainsessionmetric
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibailibrarainsessionmetricData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibailibrarainsessionmetricData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibailibrarainsessionmetricData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libainirobotinfo model data

	for {
		var LibainirobotinfoData []model.Libainirobotinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibainirobotinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibainirobotinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibainirobotinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchinfo model data

	for {
		var LibbatchinfoData []model.Libbatchinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchoperateindex model data

	for {
		var LibbatchoperateindexData []model.Libbatchoperateindex
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchoperateindexData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchoperateindexData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchoperateindexData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchoperatelog model data

	for {
		var LibbatchoperatelogData []model.Libbatchoperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbatchoperatelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbatchoperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbatchoperatelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbookinfo model data

	for {
		var LibbookinfoData []model.Libbookinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibbookinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibbookinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibbookinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libcirculatelog model data

	for {
		var LibcirculatelogData []model.Libcirculatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibcirculatelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibcirculatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibcirculatelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing LibcirculatelogBak model data

	for {
		var LibcirculatelogBakData []model.LibcirculatelogBak
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibcirculatelogBakData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibcirculatelogBakData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibcirculatelogBakData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libfeedback model data

	for {
		var LibfeedbackData []model.Libfeedback
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibfeedbackData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibfeedbackData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibfeedbackData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libfeelog model data

	for {
		var LibfeelogData []model.Libfeelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibfeelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibfeelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibfeelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorystat model data

	for {
		var LibinventorystatData []model.Libinventorystat
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventorystatData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventorystatData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventorystatData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorytask model data

	for {
		var LibinventorytaskData []model.Libinventorytask
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventorytaskData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventorytaskData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventorytaskData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorywork model data

	for {
		var LibinventoryworkData []model.Libinventorywork
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworkData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworkData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworkData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventoryworkdetail model data

	for {
		var LibinventoryworkdetailData []model.Libinventoryworkdetail
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworkdetailData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworkdetailData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworkdetailData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventoryworklog model data

	for {
		var LibinventoryworklogData []model.Libinventoryworklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibinventoryworklogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibinventoryworklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibinventoryworklogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitem model data

	for {
		var LibitemData []model.Libitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing LibitemCopy model data

	for {
		var LibitemCopyData []model.LibitemCopy
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemCopyData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemCopyData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemCopyData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libiteminventoryinfo model data

	for {
		var LibiteminventoryinfoData []model.Libiteminventoryinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibiteminventoryinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibiteminventoryinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibiteminventoryinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libiteminventorylog model data

	for {
		var LibiteminventorylogData []model.Libiteminventorylog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibiteminventorylogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibiteminventorylogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibiteminventorylogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemlocinfo model data

	for {
		var LibitemlocinfoData []model.Libitemlocinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemlocinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemlocinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemlocinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemoperateindexlog model data

	for {
		var LibitemoperateindexlogData []model.Libitemoperateindexlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemoperateindexlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemoperateindexlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemoperateindexlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemoperatelog model data

	for {
		var LibitemoperatelogData []model.Libitemoperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibitemoperatelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibitemoperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibitemoperatelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libjournalinfo model data

	for {
		var LibjournalinfoData []model.Libjournalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibjournalinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibjournalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibjournalinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblabel model data

	for {
		var LiblabelData []model.Liblabel
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblabelData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblabelData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblabelData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblabeloperatelog model data

	for {
		var LiblabeloperatelogData []model.Liblabeloperatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblabeloperatelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblabeloperatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblabeloperatelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblayer model data

	for {
		var LiblayerData []model.Liblayer
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblayerData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblayerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblayerData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblayerindexupdatelog model data

	for {
		var LiblayerindexupdatelogData []model.Liblayerindexupdatelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LiblayerindexupdatelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LiblayerindexupdatelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LiblayerindexupdatelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libnotificationlog model data

	for {
		var LibnotificationlogData []model.Libnotificationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibnotificationlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibnotificationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibnotificationlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpartonreservation model data

	for {
		var LibpartonreservationData []model.Libpartonreservation
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpartonreservationData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpartonreservationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpartonreservationData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatron model data

	for {
		var LibpatronData []model.Libpatron
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatronitem model data

	for {
		var LibpatronitemData []model.Libpatronitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatronlog model data

	for {
		var LibpatronlogData []model.Libpatronlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpatronlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpatronlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpatronlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpointsclearing model data

	for {
		var LibpointsclearingData []model.Libpointsclearing
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpointsclearingData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpointsclearingData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpointsclearingData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpointslog model data

	for {
		var LibpointslogData []model.Libpointslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibpointslogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibpointslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibpointslogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidantlayer model data

	for {
		var LibrfidantlayerData []model.Librfidantlayer
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidantlayerData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidantlayerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidantlayerData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidscandetaillog model data

	for {
		var LibrfidscandetaillogData []model.Librfidscandetaillog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidscandetaillogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidscandetaillogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidscandetaillogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidscanlog model data

	for {
		var LibrfidscanlogData []model.Librfidscanlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrfidscanlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrfidscanlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrfidscanlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Librow model data

	for {
		var LibrowData []model.Librow
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrowData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrowData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrowData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Librowcatalog model data

	for {
		var LibrowcatalogData []model.Librowcatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibrowcatalogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibrowcatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibrowcatalogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libscanitemlog model data

	for {
		var LibscanitemlogData []model.Libscanitemlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibscanitemlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibscanitemlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibscanitemlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libshelf model data

	for {
		var LibshelfData []model.Libshelf
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibshelfData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibshelfData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibshelfData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libstruct model data

	for {
		var LibstructData []model.Libstruct
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibstructData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibstructData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibstructData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtagtobarcodelog model data

	for {
		var LibtagtobarcodelogData []model.Libtagtobarcodelog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtagtobarcodelogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtagtobarcodelogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtagtobarcodelogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtaskitem model data

	for {
		var LibtaskitemData []model.Libtaskitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtaskitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtaskitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtaskitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtaskpackage model data

	for {
		var LibtaskpackageData []model.Libtaskpackage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibtaskpackageData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibtaskpackageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibtaskpackageData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Libupdatefirstbooklog model data

	for {
		var LibupdatefirstbooklogData []model.Libupdatefirstbooklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&LibupdatefirstbooklogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(LibupdatefirstbooklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(LibupdatefirstbooklogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Misactivity model data

	for {
		var MisactivityData []model.Misactivity
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MisactivityData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(MisactivityData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MisactivityData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Mismediainfo model data

	for {
		var MismediainfoData []model.Mismediainfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MismediainfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(MismediainfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MismediainfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Misnews model data

	for {
		var MisnewsData []model.Misnews
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MisnewsData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(MisnewsData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MisnewsData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Mistemplate model data

	for {
		var MistemplateData []model.Mistemplate
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&MistemplateData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(MistemplateData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(MistemplateData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Rescatalog model data

	for {
		var RescatalogData []model.Rescatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RescatalogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(RescatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RescatalogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Rescipinfo model data

	for {
		var RescipinfoData []model.Rescipinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RescipinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(RescipinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RescipinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Resfourcorner model data

	for {
		var ResfourcornerData []model.Resfourcorner
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResfourcornerData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ResfourcornerData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResfourcornerData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Resjournalinfo model data

	for {
		var ResjournalinfoData []model.Resjournalinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResjournalinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ResjournalinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResjournalinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Resnotfound model data

	for {
		var ResnotfoundData []model.Resnotfound
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&ResnotfoundData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(ResnotfoundData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(ResnotfoundData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Respublisherinfo model data

	for {
		var RespublisherinfoData []model.Respublisherinfo
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&RespublisherinfoData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(RespublisherinfoData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(RespublisherinfoData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Ssbackgroundjob model data

	for {
		var SsbackgroundjobData []model.Ssbackgroundjob
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SsbackgroundjobData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SsbackgroundjobData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SsbackgroundjobData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysattachment model data

	for {
		var SysattachmentData []model.Sysattachment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysattachmentData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysattachmentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysattachmentData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditacslog model data

	for {
		var SysauditacslogData []model.Sysauditacslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditacslogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditacslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditacslogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditapilog model data

	for {
		var SysauditapilogData []model.Sysauditapilog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditapilogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditapilogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditapilogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditapplog model data

	for {
		var SysauditapplogData []model.Sysauditapplog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditapplogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditapplogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditapplogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditlinklog model data

	for {
		var SysauditlinklogData []model.Sysauditlinklog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditlinklogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditlinklogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditlinklogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditlmslog model data

	for {
		var SysauditlmslogData []model.Sysauditlmslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditlmslogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditlmslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditlmslogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditsslog model data

	for {
		var SysauditsslogData []model.Sysauditsslog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysauditsslogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysauditsslogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysauditsslogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysblocklist model data

	for {
		var SysblocklistData []model.Sysblocklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysblocklistData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysblocklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysblocklistData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbookblocklist model data

	for {
		var SysbookblocklistData []model.Sysbookblocklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbookblocklistData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbookblocklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbookblocklistData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbooknumlib model data

	for {
		var SysbooknumlibData []model.Sysbooknumlib
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbooknumlibData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbooknumlibData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbooknumlibData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbooknumset model data

	for {
		var SysbooknumsetData []model.Sysbooknumset
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysbooknumsetData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysbooknumsetData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysbooknumsetData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscardconfig model data

	for {
		var SyscardconfigData []model.Syscardconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscardconfigData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscardconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscardconfigData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscarddevdtl model data

	for {
		var SyscarddevdtlData []model.Syscarddevdtl
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscarddevdtlData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscarddevdtlData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscarddevdtlData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscardtype model data

	for {
		var SyscardtypeData []model.Syscardtype
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscardtypeData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscardtypeData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscardtypeData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscoderule model data

	for {
		var SyscoderuleData []model.Syscoderule
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscoderuleData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscoderuleData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscoderuleData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscoderuleseed model data

	for {
		var SyscoderuleseedData []model.Syscoderuleseed
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyscoderuleseedData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyscoderuleseedData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyscoderuleseedData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysconfigbase model data

	for {
		var SysconfigbaseData []model.Sysconfigbase
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysconfigbaseData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysconfigbaseData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysconfigbaseData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysconfiglog model data

	for {
		var SysconfiglogData []model.Sysconfiglog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysconfiglogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysconfiglogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysconfiglogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdataitem model data

	for {
		var SysdataitemData []model.Sysdataitem
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdataitemData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdataitemData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdataitemData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdataitemdetail model data

	for {
		var SysdataitemdetailData []model.Sysdataitemdetail
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdataitemdetailData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdataitemdetailData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdataitemdetailData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdatalog model data

	for {
		var SysdatalogData []model.Sysdatalog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdatalogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdatalogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdatalogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdepartment model data

	for {
		var SysdepartmentData []model.Sysdepartment
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysdepartmentData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysdepartmentData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysdepartmentData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumfield model data

	for {
		var SysenumfieldData []model.Sysenumfield
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysenumfieldData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysenumfieldData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysenumfieldData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumvalue model data

	for {
		var SysenumvalueData []model.Sysenumvalue
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysenumvalueData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysenumvalueData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysenumvalueData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumvalue2 model data

	for {
		var Sysenumvalue2Data []model.Sysenumvalue2
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&Sysenumvalue2Data).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(Sysenumvalue2Data) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(Sysenumvalue2Data, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceoffineoperationlog model data

	for {
		var SysfaceoffineoperationlogData []model.Sysfaceoffineoperationlog
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceoffineoperationlogData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceoffineoperationlogData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceoffineoperationlogData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflinefeature model data

	for {
		var SysfaceofflinefeatureData []model.Sysfaceofflinefeature
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflinefeatureData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflinefeatureData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflinefeatureData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflinegroup model data

	for {
		var SysfaceofflinegroupData []model.Sysfaceofflinegroup
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflinegroupData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflinegroupData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflinegroupData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflineuser model data

	for {
		var SysfaceofflineuserData []model.Sysfaceofflineuser
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysfaceofflineuserData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysfaceofflineuserData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysfaceofflineuserData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslanguage model data

	for {
		var SyslanguageData []model.Syslanguage
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslanguageData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslanguageData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslanguageData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslayertran model data

	for {
		var SyslayertranData []model.Syslayertran
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslayertranData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslayertranData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslayertranData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslocation model data

	for {
		var SyslocationData []model.Syslocation
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SyslocationData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SyslocationData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SyslocationData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysmenu model data

	for {
		var SysmenuData []model.Sysmenu
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysmenuData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysmenuData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysmenuData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing SysmenuCopy model data

	for {
		var SysmenuCopyData []model.SysmenuCopy
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SysmenuCopyData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SysmenuCopyData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SysmenuCopyData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Systasklist model data

	for {
		var SystasklistData []model.Systasklist
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystasklistData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SystasklistData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystasklistData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Systenantconfig model data

	for {
		var SystenantconfigData []model.Systenantconfig
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystenantconfigData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SystenantconfigData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystenantconfigData, batchSize).Error; err != nil {
			//return err
		}
	}

	// Syncing Systenantextend model data

	for {
		var SystenantextendData []model.Systenantextend
		// 分页查询数据，不依赖排序字段
		if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&SystenantextendData).Error; err != nil {
			//return err
		}

		// 如果没有更多数据，退出循环
		if len(SystenantextendData) == 0 {
			break
		}

		// 更新 offset
		offset += batchSize

		// 将数据插入到 db2
		if err := global.OWEN_DBList["to"].CreateInBatches(SystenantextendData, batchSize).Error; err != nil {
			//return err
		}
	}

	return nil
}
