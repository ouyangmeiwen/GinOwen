package serviceinit

import (
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
)

func CusSyncDatabase() error {
	// Syncing Abpauditlog model data
	var AbpauditlogData []model.Abpauditlog
	if err := global.OWEN_DBList["from"].Find(&AbpauditlogData).Error; err != nil {
		//return err
	}
	for _, record := range AbpauditlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpbackgroundjob model data
	var AbpbackgroundjobData []model.Abpbackgroundjob
	if err := global.OWEN_DBList["from"].Find(&AbpbackgroundjobData).Error; err != nil {
		//return err
	}
	for _, record := range AbpbackgroundjobData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpedition model data
	var AbpeditionData []model.Abpedition
	if err := global.OWEN_DBList["from"].Find(&AbpeditionData).Error; err != nil {
		//return err
	}
	for _, record := range AbpeditionData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitychange model data
	var AbpentitychangeData []model.Abpentitychange
	if err := global.OWEN_DBList["from"].Find(&AbpentitychangeData).Error; err != nil {
		//return err
	}
	for _, record := range AbpentitychangeData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitychangeset model data
	var AbpentitychangesetData []model.Abpentitychangeset
	if err := global.OWEN_DBList["from"].Find(&AbpentitychangesetData).Error; err != nil {
		//return err
	}
	for _, record := range AbpentitychangesetData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpentitypropertychange model data
	var AbpentitypropertychangeData []model.Abpentitypropertychange
	if err := global.OWEN_DBList["from"].Find(&AbpentitypropertychangeData).Error; err != nil {
		//return err
	}
	for _, record := range AbpentitypropertychangeData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpfeature model data
	var AbpfeatureData []model.Abpfeature
	if err := global.OWEN_DBList["from"].Find(&AbpfeatureData).Error; err != nil {
		//return err
	}
	for _, record := range AbpfeatureData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abplanguage model data
	var AbplanguageData []model.Abplanguage
	if err := global.OWEN_DBList["from"].Find(&AbplanguageData).Error; err != nil {
		//return err
	}
	for _, record := range AbplanguageData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abplanguagetext model data
	var AbplanguagetextData []model.Abplanguagetext
	if err := global.OWEN_DBList["from"].Find(&AbplanguagetextData).Error; err != nil {
		//return err
	}
	for _, record := range AbplanguagetextData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpnotification model data
	var AbpnotificationData []model.Abpnotification
	if err := global.OWEN_DBList["from"].Find(&AbpnotificationData).Error; err != nil {
		//return err
	}
	for _, record := range AbpnotificationData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpnotificationsubscription model data
	var AbpnotificationsubscriptionData []model.Abpnotificationsubscription
	if err := global.OWEN_DBList["from"].Find(&AbpnotificationsubscriptionData).Error; err != nil {
		//return err
	}
	for _, record := range AbpnotificationsubscriptionData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abporganizationunitrole model data
	var AbporganizationunitroleData []model.Abporganizationunitrole
	if err := global.OWEN_DBList["from"].Find(&AbporganizationunitroleData).Error; err != nil {
		//return err
	}
	for _, record := range AbporganizationunitroleData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abporganizationunit model data
	var AbporganizationunitData []model.Abporganizationunit
	if err := global.OWEN_DBList["from"].Find(&AbporganizationunitData).Error; err != nil {
		//return err
	}
	for _, record := range AbporganizationunitData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abppermission model data
	var AbppermissionData []model.Abppermission
	if err := global.OWEN_DBList["from"].Find(&AbppermissionData).Error; err != nil {
		//return err
	}
	for _, record := range AbppermissionData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abppersistedgrant model data
	var AbppersistedgrantData []model.Abppersistedgrant
	if err := global.OWEN_DBList["from"].Find(&AbppersistedgrantData).Error; err != nil {
		//return err
	}
	for _, record := range AbppersistedgrantData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abproleclaim model data
	var AbproleclaimData []model.Abproleclaim
	if err := global.OWEN_DBList["from"].Find(&AbproleclaimData).Error; err != nil {
		//return err
	}
	for _, record := range AbproleclaimData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abprole model data
	var AbproleData []model.Abprole
	if err := global.OWEN_DBList["from"].Find(&AbproleData).Error; err != nil {
		//return err
	}
	for _, record := range AbproleData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpsetting model data
	var AbpsettingData []model.Abpsetting
	if err := global.OWEN_DBList["from"].Find(&AbpsettingData).Error; err != nil {
		//return err
	}
	for _, record := range AbpsettingData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abptenantnotification model data
	var AbptenantnotificationData []model.Abptenantnotification
	if err := global.OWEN_DBList["from"].Find(&AbptenantnotificationData).Error; err != nil {
		//return err
	}
	for _, record := range AbptenantnotificationData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abptenant model data
	var AbptenantData []model.Abptenant
	if err := global.OWEN_DBList["from"].Find(&AbptenantData).Error; err != nil {
		//return err
	}
	for _, record := range AbptenantData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuseraccount model data
	var AbpuseraccountData []model.Abpuseraccount
	if err := global.OWEN_DBList["from"].Find(&AbpuseraccountData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuseraccountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserclaim model data
	var AbpuserclaimData []model.Abpuserclaim
	if err := global.OWEN_DBList["from"].Find(&AbpuserclaimData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserclaimData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserloginattempt model data
	var AbpuserloginattemptData []model.Abpuserloginattempt
	if err := global.OWEN_DBList["from"].Find(&AbpuserloginattemptData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserloginattemptData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserlogin model data
	var AbpuserloginData []model.Abpuserlogin
	if err := global.OWEN_DBList["from"].Find(&AbpuserloginData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserloginData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpusernotification model data
	var AbpusernotificationData []model.Abpusernotification
	if err := global.OWEN_DBList["from"].Find(&AbpusernotificationData).Error; err != nil {
		//return err
	}
	for _, record := range AbpusernotificationData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserorganizationunit model data
	var AbpuserorganizationunitData []model.Abpuserorganizationunit
	if err := global.OWEN_DBList["from"].Find(&AbpuserorganizationunitData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserorganizationunitData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuserrole model data
	var AbpuserroleData []model.Abpuserrole
	if err := global.OWEN_DBList["from"].Find(&AbpuserroleData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserroleData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpuser model data
	var AbpuserData []model.Abpuser
	if err := global.OWEN_DBList["from"].Find(&AbpuserData).Error; err != nil {
		//return err
	}
	for _, record := range AbpuserData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Abpusertoken model data
	var AbpusertokenData []model.Abpusertoken
	if err := global.OWEN_DBList["from"].Find(&AbpusertokenData).Error; err != nil {
		//return err
	}
	for _, record := range AbpusertokenData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appaliuser model data
	var AppaliuserData []model.Appaliuser
	if err := global.OWEN_DBList["from"].Find(&AppaliuserData).Error; err != nil {
		//return err
	}
	for _, record := range AppaliuserData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appapprovalinfo model data
	var AppapprovalinfoData []model.Appapprovalinfo
	if err := global.OWEN_DBList["from"].Find(&AppapprovalinfoData).Error; err != nil {
		//return err
	}
	for _, record := range AppapprovalinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appapprovaltemplate model data
	var AppapprovaltemplateData []model.Appapprovaltemplate
	if err := global.OWEN_DBList["from"].Find(&AppapprovaltemplateData).Error; err != nil {
		//return err
	}
	for _, record := range AppapprovaltemplateData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appbinaryobject model data
	var AppbinaryobjectData []model.Appbinaryobject
	if err := global.OWEN_DBList["from"].Find(&AppbinaryobjectData).Error; err != nil {
		//return err
	}
	for _, record := range AppbinaryobjectData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appbookorder model data
	var AppbookorderData []model.Appbookorder
	if err := global.OWEN_DBList["from"].Find(&AppbookorderData).Error; err != nil {
		//return err
	}
	for _, record := range AppbookorderData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appchatmessage model data
	var AppchatmessageData []model.Appchatmessage
	if err := global.OWEN_DBList["from"].Find(&AppchatmessageData).Error; err != nil {
		//return err
	}
	for _, record := range AppchatmessageData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appcreditloginorder model data
	var AppcreditloginorderData []model.Appcreditloginorder
	if err := global.OWEN_DBList["from"].Find(&AppcreditloginorderData).Error; err != nil {
		//return err
	}
	for _, record := range AppcreditloginorderData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appfriendship model data
	var AppfriendshipData []model.Appfriendship
	if err := global.OWEN_DBList["from"].Find(&AppfriendshipData).Error; err != nil {
		//return err
	}
	for _, record := range AppfriendshipData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appinvoice model data
	var AppinvoiceData []model.Appinvoice
	if err := global.OWEN_DBList["from"].Find(&AppinvoiceData).Error; err != nil {
		//return err
	}
	for _, record := range AppinvoiceData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appitemlocked model data
	var AppitemlockedData []model.Appitemlocked
	if err := global.OWEN_DBList["from"].Find(&AppitemlockedData).Error; err != nil {
		//return err
	}
	for _, record := range AppitemlockedData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appmessageboard model data
	var AppmessageboardData []model.Appmessageboard
	if err := global.OWEN_DBList["from"].Find(&AppmessageboardData).Error; err != nil {
		//return err
	}
	for _, record := range AppmessageboardData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appnotificationlog model data
	var AppnotificationlogData []model.Appnotificationlog
	if err := global.OWEN_DBList["from"].Find(&AppnotificationlogData).Error; err != nil {
		//return err
	}
	for _, record := range AppnotificationlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Apppayorder model data
	var ApppayorderData []model.Apppayorder
	if err := global.OWEN_DBList["from"].Find(&ApppayorderData).Error; err != nil {
		//return err
	}
	for _, record := range ApppayorderData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Apppickupcode model data
	var ApppickupcodeData []model.Apppickupcode
	if err := global.OWEN_DBList["from"].Find(&ApppickupcodeData).Error; err != nil {
		//return err
	}
	for _, record := range ApppickupcodeData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appqrcode model data
	var AppqrcodeData []model.Appqrcode
	if err := global.OWEN_DBList["from"].Find(&AppqrcodeData).Error; err != nil {
		//return err
	}
	for _, record := range AppqrcodeData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Apprecommendinfo model data
	var ApprecommendinfoData []model.Apprecommendinfo
	if err := global.OWEN_DBList["from"].Find(&ApprecommendinfoData).Error; err != nil {
		//return err
	}
	for _, record := range ApprecommendinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appsubbookinfo model data
	var AppsubbookinfoData []model.Appsubbookinfo
	if err := global.OWEN_DBList["from"].Find(&AppsubbookinfoData).Error; err != nil {
		//return err
	}
	for _, record := range AppsubbookinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appsubscriptionpayment model data
	var AppsubscriptionpaymentData []model.Appsubscriptionpayment
	if err := global.OWEN_DBList["from"].Find(&AppsubscriptionpaymentData).Error; err != nil {
		//return err
	}
	for _, record := range AppsubscriptionpaymentData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appusercard model data
	var AppusercardData []model.Appusercard
	if err := global.OWEN_DBList["from"].Find(&AppusercardData).Error; err != nil {
		//return err
	}
	for _, record := range AppusercardData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Appweuser model data
	var AppweuserData []model.Appweuser
	if err := global.OWEN_DBList["from"].Find(&AppweuserData).Error; err != nil {
		//return err
	}
	for _, record := range AppweuserData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasbusinesscount model data
	var DasbusinesscountData []model.Dasbusinesscount
	if err := global.OWEN_DBList["from"].Find(&DasbusinesscountData).Error; err != nil {
		//return err
	}
	for _, record := range DasbusinesscountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dascirculatecount model data
	var DascirculatecountData []model.Dascirculatecount
	if err := global.OWEN_DBList["from"].Find(&DascirculatecountData).Error; err != nil {
		//return err
	}
	for _, record := range DascirculatecountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasdatabaselink model data
	var DasdatabaselinkData []model.Dasdatabaselink
	if err := global.OWEN_DBList["from"].Find(&DasdatabaselinkData).Error; err != nil {
		//return err
	}
	for _, record := range DasdatabaselinkData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasdatasource model data
	var DasdatasourceData []model.Dasdatasource
	if err := global.OWEN_DBList["from"].Find(&DasdatasourceData).Error; err != nil {
		//return err
	}
	for _, record := range DasdatasourceData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasfeecount model data
	var DasfeecountData []model.Dasfeecount
	if err := global.OWEN_DBList["from"].Find(&DasfeecountData).Error; err != nil {
		//return err
	}
	for _, record := range DasfeecountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Daspatronlogcount model data
	var DaspatronlogcountData []model.Daspatronlogcount
	if err := global.OWEN_DBList["from"].Find(&DaspatronlogcountData).Error; err != nil {
		//return err
	}
	for _, record := range DaspatronlogcountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasperformance model data
	var DasperformanceData []model.Dasperformance
	if err := global.OWEN_DBList["from"].Find(&DasperformanceData).Error; err != nil {
		//return err
	}
	for _, record := range DasperformanceData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dassecuritygatecount model data
	var DassecuritygatecountData []model.Dassecuritygatecount
	if err := global.OWEN_DBList["from"].Find(&DassecuritygatecountData).Error; err != nil {
		//return err
	}
	for _, record := range DassecuritygatecountData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasvisitpage model data
	var DasvisitpageData []model.Dasvisitpage
	if err := global.OWEN_DBList["from"].Find(&DasvisitpageData).Error; err != nil {
		//return err
	}
	for _, record := range DasvisitpageData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Dasvisittrend model data
	var DasvisittrendData []model.Dasvisittrend
	if err := global.OWEN_DBList["from"].Find(&DasvisittrendData).Error; err != nil {
		//return err
	}
	for _, record := range DasvisittrendData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Efmigrationshistory model data
	var EfmigrationshistoryData []model.Efmigrationshistory
	if err := global.OWEN_DBList["from"].Find(&EfmigrationshistoryData).Error; err != nil {
		//return err
	}
	for _, record := range EfmigrationshistoryData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireaggregatedcounter model data
	var HangfireaggregatedcounterData []model.Hangfireaggregatedcounter
	if err := global.OWEN_DBList["from"].Find(&HangfireaggregatedcounterData).Error; err != nil {
		//return err
	}
	for _, record := range HangfireaggregatedcounterData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirecounter model data
	var HangfirecounterData []model.Hangfirecounter
	if err := global.OWEN_DBList["from"].Find(&HangfirecounterData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirecounterData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfiredistributedlock model data
	var HangfiredistributedlockData []model.Hangfiredistributedlock
	if err := global.OWEN_DBList["from"].Find(&HangfiredistributedlockData).Error; err != nil {
		//return err
	}
	for _, record := range HangfiredistributedlockData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirehash model data
	var HangfirehashData []model.Hangfirehash
	if err := global.OWEN_DBList["from"].Find(&HangfirehashData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirehashData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejob model data
	var HangfirejobData []model.Hangfirejob
	if err := global.OWEN_DBList["from"].Find(&HangfirejobData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirejobData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobparameter model data
	var HangfirejobparameterData []model.Hangfirejobparameter
	if err := global.OWEN_DBList["from"].Find(&HangfirejobparameterData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirejobparameterData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobqueue model data
	var HangfirejobqueueData []model.Hangfirejobqueue
	if err := global.OWEN_DBList["from"].Find(&HangfirejobqueueData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirejobqueueData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirejobstate model data
	var HangfirejobstateData []model.Hangfirejobstate
	if err := global.OWEN_DBList["from"].Find(&HangfirejobstateData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirejobstateData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirelist model data
	var HangfirelistData []model.Hangfirelist
	if err := global.OWEN_DBList["from"].Find(&HangfirelistData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirelistData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireserver model data
	var HangfireserverData []model.Hangfireserver
	if err := global.OWEN_DBList["from"].Find(&HangfireserverData).Error; err != nil {
		//return err
	}
	for _, record := range HangfireserverData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfireset model data
	var HangfiresetData []model.Hangfireset
	if err := global.OWEN_DBList["from"].Find(&HangfiresetData).Error; err != nil {
		//return err
	}
	for _, record := range HangfiresetData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Hangfirestate model data
	var HangfirestateData []model.Hangfirestate
	if err := global.OWEN_DBList["from"].Find(&HangfirestateData).Error; err != nil {
		//return err
	}
	for _, record := range HangfirestateData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpcommandlog model data
	var LcpcommandlogData []model.Lcpcommandlog
	if err := global.OWEN_DBList["from"].Find(&LcpcommandlogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpcommandlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpconfig model data
	var LcpconfigData []model.Lcpconfig
	if err := global.OWEN_DBList["from"].Find(&LcpconfigData).Error; err != nil {
		//return err
	}
	for _, record := range LcpconfigData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpmaintainlog model data
	var LcpmaintainlogData []model.Lcpmaintainlog
	if err := global.OWEN_DBList["from"].Find(&LcpmaintainlogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpmaintainlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpproduct model data
	var LcpproductData []model.Lcpproduct
	if err := global.OWEN_DBList["from"].Find(&LcpproductData).Error; err != nil {
		//return err
	}
	for _, record := range LcpproductData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcprfidantenna model data
	var LcprfidantennaData []model.Lcprfidantenna
	if err := global.OWEN_DBList["from"].Find(&LcprfidantennaData).Error; err != nil {
		//return err
	}
	for _, record := range LcprfidantennaData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcprfidreader model data
	var LcprfidreaderData []model.Lcprfidreader
	if err := global.OWEN_DBList["from"].Find(&LcprfidreaderData).Error; err != nil {
		//return err
	}
	for _, record := range LcprfidreaderData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatebookaccesslog model data
	var LcpsecuritygatebookaccesslogData []model.Lcpsecuritygatebookaccesslog
	if err := global.OWEN_DBList["from"].Find(&LcpsecuritygatebookaccesslogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpsecuritygatebookaccesslogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatebookdailyaccess model data
	var LcpsecuritygatebookdailyaccessData []model.Lcpsecuritygatebookdailyaccess
	if err := global.OWEN_DBList["from"].Find(&LcpsecuritygatebookdailyaccessData).Error; err != nil {
		//return err
	}
	for _, record := range LcpsecuritygatebookdailyaccessData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygateitemlog model data
	var LcpsecuritygateitemlogData []model.Lcpsecuritygateitemlog
	if err := global.OWEN_DBList["from"].Find(&LcpsecuritygateitemlogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpsecuritygateitemlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpsecuritygatepatronlog model data
	var LcpsecuritygatepatronlogData []model.Lcpsecuritygatepatronlog
	if err := global.OWEN_DBList["from"].Find(&LcpsecuritygatepatronlogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpsecuritygatepatronlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpserialport model data
	var LcpserialportData []model.Lcpserialport
	if err := global.OWEN_DBList["from"].Find(&LcpserialportData).Error; err != nil {
		//return err
	}
	for _, record := range LcpserialportData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpserialportext model data
	var LcpserialportextData []model.Lcpserialportext
	if err := global.OWEN_DBList["from"].Find(&LcpserialportextData).Error; err != nil {
		//return err
	}
	for _, record := range LcpserialportextData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpservice model data
	var LcpserviceData []model.Lcpservice
	if err := global.OWEN_DBList["from"].Find(&LcpserviceData).Error; err != nil {
		//return err
	}
	for _, record := range LcpserviceData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminal model data
	var LcpterminalData []model.Lcpterminal
	if err := global.OWEN_DBList["from"].Find(&LcpterminalData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaladvertisement model data
	var LcpterminaladvertisementData []model.Lcpterminaladvertisement
	if err := global.OWEN_DBList["from"].Find(&LcpterminaladvertisementData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminaladvertisementData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalbox model data
	var LcpterminalboxData []model.Lcpterminalbox
	if err := global.OWEN_DBList["from"].Find(&LcpterminalboxData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalboxData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalboxitem model data
	var LcpterminalboxitemData []model.Lcpterminalboxitem
	if err := global.OWEN_DBList["from"].Find(&LcpterminalboxitemData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalboxitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaldevice model data
	var LcpterminaldeviceData []model.Lcpterminaldevice
	if err := global.OWEN_DBList["from"].Find(&LcpterminaldeviceData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminaldeviceData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminaldevicelog model data
	var LcpterminaldevicelogData []model.Lcpterminaldevicelog
	if err := global.OWEN_DBList["from"].Find(&LcpterminaldevicelogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminaldevicelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminallog model data
	var LcpterminallogData []model.Lcpterminallog
	if err := global.OWEN_DBList["from"].Find(&LcpterminallogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminallogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalpermission model data
	var LcpterminalpermissionData []model.Lcpterminalpermission
	if err := global.OWEN_DBList["from"].Find(&LcpterminalpermissionData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalpermissionData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelf model data
	var LcpterminalshelfData []model.Lcpterminalshelf
	if err := global.OWEN_DBList["from"].Find(&LcpterminalshelfData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalshelfData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelfitem model data
	var LcpterminalshelfitemData []model.Lcpterminalshelfitem
	if err := global.OWEN_DBList["from"].Find(&LcpterminalshelfitemData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalshelfitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpterminalshelflog model data
	var LcpterminalshelflogData []model.Lcpterminalshelflog
	if err := global.OWEN_DBList["from"].Find(&LcpterminalshelflogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpterminalshelflogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpupgradelog model data
	var LcpupgradelogData []model.Lcpupgradelog
	if err := global.OWEN_DBList["from"].Find(&LcpupgradelogData).Error; err != nil {
		//return err
	}
	for _, record := range LcpupgradelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Lcpversion model data
	var LcpversionData []model.Lcpversion
	if err := global.OWEN_DBList["from"].Find(&LcpversionData).Error; err != nil {
		//return err
	}
	for _, record := range LcpversionData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfo model data
	var LibailibrarainbaseinfoData []model.Libailibrarainbaseinfo
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainbaseinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainbaseinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfoitem model data
	var LibailibrarainbaseinfoitemData []model.Libailibrarainbaseinfoitem
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainbaseinfoitemData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainbaseinfoitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainbaseinfoprofile model data
	var LibailibrarainbaseinfoprofileData []model.Libailibrarainbaseinfoprofile
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainbaseinfoprofileData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainbaseinfoprofileData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainknowledgefileinfo model data
	var LibailibrarainknowledgefileinfoData []model.Libailibrarainknowledgefileinfo
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainknowledgefileinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainknowledgefileinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainquestionmetric model data
	var LibailibrarainquestionmetricData []model.Libailibrarainquestionmetric
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainquestionmetricData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainquestionmetricData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libailibrarainsessionmetric model data
	var LibailibrarainsessionmetricData []model.Libailibrarainsessionmetric
	if err := global.OWEN_DBList["from"].Find(&LibailibrarainsessionmetricData).Error; err != nil {
		//return err
	}
	for _, record := range LibailibrarainsessionmetricData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libainirobotinfo model data
	var LibainirobotinfoData []model.Libainirobotinfo
	if err := global.OWEN_DBList["from"].Find(&LibainirobotinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibainirobotinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchinfo model data
	var LibbatchinfoData []model.Libbatchinfo
	if err := global.OWEN_DBList["from"].Find(&LibbatchinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibbatchinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchoperateindex model data
	var LibbatchoperateindexData []model.Libbatchoperateindex
	if err := global.OWEN_DBList["from"].Find(&LibbatchoperateindexData).Error; err != nil {
		//return err
	}
	for _, record := range LibbatchoperateindexData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbatchoperatelog model data
	var LibbatchoperatelogData []model.Libbatchoperatelog
	if err := global.OWEN_DBList["from"].Find(&LibbatchoperatelogData).Error; err != nil {
		//return err
	}
	for _, record := range LibbatchoperatelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libbookinfo model data
	var LibbookinfoData []model.Libbookinfo
	if err := global.OWEN_DBList["from"].Find(&LibbookinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibbookinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libcirculatelog model data
	var LibcirculatelogData []model.Libcirculatelog
	if err := global.OWEN_DBList["from"].Find(&LibcirculatelogData).Error; err != nil {
		//return err
	}
	for _, record := range LibcirculatelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing LibcirculatelogBak model data
	var LibcirculatelogBakData []model.LibcirculatelogBak
	if err := global.OWEN_DBList["from"].Find(&LibcirculatelogBakData).Error; err != nil {
		//return err
	}
	for _, record := range LibcirculatelogBakData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libfeedback model data
	var LibfeedbackData []model.Libfeedback
	if err := global.OWEN_DBList["from"].Find(&LibfeedbackData).Error; err != nil {
		//return err
	}
	for _, record := range LibfeedbackData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libfeelog model data
	var LibfeelogData []model.Libfeelog
	if err := global.OWEN_DBList["from"].Find(&LibfeelogData).Error; err != nil {
		//return err
	}
	for _, record := range LibfeelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorystat model data
	var LibinventorystatData []model.Libinventorystat
	if err := global.OWEN_DBList["from"].Find(&LibinventorystatData).Error; err != nil {
		//return err
	}
	for _, record := range LibinventorystatData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorytask model data
	var LibinventorytaskData []model.Libinventorytask
	if err := global.OWEN_DBList["from"].Find(&LibinventorytaskData).Error; err != nil {
		//return err
	}
	for _, record := range LibinventorytaskData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventorywork model data
	var LibinventoryworkData []model.Libinventorywork
	if err := global.OWEN_DBList["from"].Find(&LibinventoryworkData).Error; err != nil {
		//return err
	}
	for _, record := range LibinventoryworkData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventoryworkdetail model data
	var LibinventoryworkdetailData []model.Libinventoryworkdetail
	if err := global.OWEN_DBList["from"].Find(&LibinventoryworkdetailData).Error; err != nil {
		//return err
	}
	for _, record := range LibinventoryworkdetailData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libinventoryworklog model data
	var LibinventoryworklogData []model.Libinventoryworklog
	if err := global.OWEN_DBList["from"].Find(&LibinventoryworklogData).Error; err != nil {
		//return err
	}
	for _, record := range LibinventoryworklogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitem model data
	var LibitemData []model.Libitem
	if err := global.OWEN_DBList["from"].Find(&LibitemData).Error; err != nil {
		//return err
	}
	for _, record := range LibitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing LibitemCopy model data
	var LibitemCopyData []model.LibitemCopy
	if err := global.OWEN_DBList["from"].Find(&LibitemCopyData).Error; err != nil {
		//return err
	}
	for _, record := range LibitemCopyData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libiteminventoryinfo model data
	var LibiteminventoryinfoData []model.Libiteminventoryinfo
	if err := global.OWEN_DBList["from"].Find(&LibiteminventoryinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibiteminventoryinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libiteminventorylog model data
	var LibiteminventorylogData []model.Libiteminventorylog
	if err := global.OWEN_DBList["from"].Find(&LibiteminventorylogData).Error; err != nil {
		//return err
	}
	for _, record := range LibiteminventorylogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemlocinfo model data
	var LibitemlocinfoData []model.Libitemlocinfo
	if err := global.OWEN_DBList["from"].Find(&LibitemlocinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibitemlocinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemoperateindexlog model data
	var LibitemoperateindexlogData []model.Libitemoperateindexlog
	if err := global.OWEN_DBList["from"].Find(&LibitemoperateindexlogData).Error; err != nil {
		//return err
	}
	for _, record := range LibitemoperateindexlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libitemoperatelog model data
	var LibitemoperatelogData []model.Libitemoperatelog
	if err := global.OWEN_DBList["from"].Find(&LibitemoperatelogData).Error; err != nil {
		//return err
	}
	for _, record := range LibitemoperatelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libjournalinfo model data
	var LibjournalinfoData []model.Libjournalinfo
	if err := global.OWEN_DBList["from"].Find(&LibjournalinfoData).Error; err != nil {
		//return err
	}
	for _, record := range LibjournalinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblabel model data
	var LiblabelData []model.Liblabel
	if err := global.OWEN_DBList["from"].Find(&LiblabelData).Error; err != nil {
		//return err
	}
	for _, record := range LiblabelData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblabeloperatelog model data
	var LiblabeloperatelogData []model.Liblabeloperatelog
	if err := global.OWEN_DBList["from"].Find(&LiblabeloperatelogData).Error; err != nil {
		//return err
	}
	for _, record := range LiblabeloperatelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblayer model data
	var LiblayerData []model.Liblayer
	if err := global.OWEN_DBList["from"].Find(&LiblayerData).Error; err != nil {
		//return err
	}
	for _, record := range LiblayerData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Liblayerindexupdatelog model data
	var LiblayerindexupdatelogData []model.Liblayerindexupdatelog
	if err := global.OWEN_DBList["from"].Find(&LiblayerindexupdatelogData).Error; err != nil {
		//return err
	}
	for _, record := range LiblayerindexupdatelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libnotificationlog model data
	var LibnotificationlogData []model.Libnotificationlog
	if err := global.OWEN_DBList["from"].Find(&LibnotificationlogData).Error; err != nil {
		//return err
	}
	for _, record := range LibnotificationlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpartonreservation model data
	var LibpartonreservationData []model.Libpartonreservation
	if err := global.OWEN_DBList["from"].Find(&LibpartonreservationData).Error; err != nil {
		//return err
	}
	for _, record := range LibpartonreservationData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatron model data
	var LibpatronData []model.Libpatron
	if err := global.OWEN_DBList["from"].Find(&LibpatronData).Error; err != nil {
		//return err
	}
	for _, record := range LibpatronData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatronitem model data
	var LibpatronitemData []model.Libpatronitem
	if err := global.OWEN_DBList["from"].Find(&LibpatronitemData).Error; err != nil {
		//return err
	}
	for _, record := range LibpatronitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpatronlog model data
	var LibpatronlogData []model.Libpatronlog
	if err := global.OWEN_DBList["from"].Find(&LibpatronlogData).Error; err != nil {
		//return err
	}
	for _, record := range LibpatronlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpointsclearing model data
	var LibpointsclearingData []model.Libpointsclearing
	if err := global.OWEN_DBList["from"].Find(&LibpointsclearingData).Error; err != nil {
		//return err
	}
	for _, record := range LibpointsclearingData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libpointslog model data
	var LibpointslogData []model.Libpointslog
	if err := global.OWEN_DBList["from"].Find(&LibpointslogData).Error; err != nil {
		//return err
	}
	for _, record := range LibpointslogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidantlayer model data
	var LibrfidantlayerData []model.Librfidantlayer
	if err := global.OWEN_DBList["from"].Find(&LibrfidantlayerData).Error; err != nil {
		//return err
	}
	for _, record := range LibrfidantlayerData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidscandetaillog model data
	var LibrfidscandetaillogData []model.Librfidscandetaillog
	if err := global.OWEN_DBList["from"].Find(&LibrfidscandetaillogData).Error; err != nil {
		//return err
	}
	for _, record := range LibrfidscandetaillogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Librfidscanlog model data
	var LibrfidscanlogData []model.Librfidscanlog
	if err := global.OWEN_DBList["from"].Find(&LibrfidscanlogData).Error; err != nil {
		//return err
	}
	for _, record := range LibrfidscanlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Librow model data
	var LibrowData []model.Librow
	if err := global.OWEN_DBList["from"].Find(&LibrowData).Error; err != nil {
		//return err
	}
	for _, record := range LibrowData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Librowcatalog model data
	var LibrowcatalogData []model.Librowcatalog
	if err := global.OWEN_DBList["from"].Find(&LibrowcatalogData).Error; err != nil {
		//return err
	}
	for _, record := range LibrowcatalogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libscanitemlog model data
	var LibscanitemlogData []model.Libscanitemlog
	if err := global.OWEN_DBList["from"].Find(&LibscanitemlogData).Error; err != nil {
		//return err
	}
	for _, record := range LibscanitemlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libshelf model data
	var LibshelfData []model.Libshelf
	if err := global.OWEN_DBList["from"].Find(&LibshelfData).Error; err != nil {
		//return err
	}
	for _, record := range LibshelfData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libstruct model data
	var LibstructData []model.Libstruct
	if err := global.OWEN_DBList["from"].Find(&LibstructData).Error; err != nil {
		//return err
	}
	for _, record := range LibstructData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtagtobarcodelog model data
	var LibtagtobarcodelogData []model.Libtagtobarcodelog
	if err := global.OWEN_DBList["from"].Find(&LibtagtobarcodelogData).Error; err != nil {
		//return err
	}
	for _, record := range LibtagtobarcodelogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtaskitem model data
	var LibtaskitemData []model.Libtaskitem
	if err := global.OWEN_DBList["from"].Find(&LibtaskitemData).Error; err != nil {
		//return err
	}
	for _, record := range LibtaskitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libtaskpackage model data
	var LibtaskpackageData []model.Libtaskpackage
	if err := global.OWEN_DBList["from"].Find(&LibtaskpackageData).Error; err != nil {
		//return err
	}
	for _, record := range LibtaskpackageData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Libupdatefirstbooklog model data
	var LibupdatefirstbooklogData []model.Libupdatefirstbooklog
	if err := global.OWEN_DBList["from"].Find(&LibupdatefirstbooklogData).Error; err != nil {
		//return err
	}
	for _, record := range LibupdatefirstbooklogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Misactivity model data
	var MisactivityData []model.Misactivity
	if err := global.OWEN_DBList["from"].Find(&MisactivityData).Error; err != nil {
		//return err
	}
	for _, record := range MisactivityData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Mismediainfo model data
	var MismediainfoData []model.Mismediainfo
	if err := global.OWEN_DBList["from"].Find(&MismediainfoData).Error; err != nil {
		//return err
	}
	for _, record := range MismediainfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Misnews model data
	var MisnewsData []model.Misnews
	if err := global.OWEN_DBList["from"].Find(&MisnewsData).Error; err != nil {
		//return err
	}
	for _, record := range MisnewsData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Mistemplate model data
	var MistemplateData []model.Mistemplate
	if err := global.OWEN_DBList["from"].Find(&MistemplateData).Error; err != nil {
		//return err
	}
	for _, record := range MistemplateData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Rescatalog model data
	var RescatalogData []model.Rescatalog
	if err := global.OWEN_DBList["from"].Find(&RescatalogData).Error; err != nil {
		//return err
	}
	for _, record := range RescatalogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Rescipinfo model data
	var RescipinfoData []model.Rescipinfo
	if err := global.OWEN_DBList["from"].Find(&RescipinfoData).Error; err != nil {
		//return err
	}
	for _, record := range RescipinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Resfourcorner model data
	var ResfourcornerData []model.Resfourcorner
	if err := global.OWEN_DBList["from"].Find(&ResfourcornerData).Error; err != nil {
		//return err
	}
	for _, record := range ResfourcornerData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Resjournalinfo model data
	var ResjournalinfoData []model.Resjournalinfo
	if err := global.OWEN_DBList["from"].Find(&ResjournalinfoData).Error; err != nil {
		//return err
	}
	for _, record := range ResjournalinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Resnotfound model data
	var ResnotfoundData []model.Resnotfound
	if err := global.OWEN_DBList["from"].Find(&ResnotfoundData).Error; err != nil {
		//return err
	}
	for _, record := range ResnotfoundData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Respublisherinfo model data
	var RespublisherinfoData []model.Respublisherinfo
	if err := global.OWEN_DBList["from"].Find(&RespublisherinfoData).Error; err != nil {
		//return err
	}
	for _, record := range RespublisherinfoData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Ssbackgroundjob model data
	var SsbackgroundjobData []model.Ssbackgroundjob
	if err := global.OWEN_DBList["from"].Find(&SsbackgroundjobData).Error; err != nil {
		//return err
	}
	for _, record := range SsbackgroundjobData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysattachment model data
	var SysattachmentData []model.Sysattachment
	if err := global.OWEN_DBList["from"].Find(&SysattachmentData).Error; err != nil {
		//return err
	}
	for _, record := range SysattachmentData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditacslog model data
	var SysauditacslogData []model.Sysauditacslog
	if err := global.OWEN_DBList["from"].Find(&SysauditacslogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditacslogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditapilog model data
	var SysauditapilogData []model.Sysauditapilog
	if err := global.OWEN_DBList["from"].Find(&SysauditapilogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditapilogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditapplog model data
	var SysauditapplogData []model.Sysauditapplog
	if err := global.OWEN_DBList["from"].Find(&SysauditapplogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditapplogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditlinklog model data
	var SysauditlinklogData []model.Sysauditlinklog
	if err := global.OWEN_DBList["from"].Find(&SysauditlinklogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditlinklogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditlmslog model data
	var SysauditlmslogData []model.Sysauditlmslog
	if err := global.OWEN_DBList["from"].Find(&SysauditlmslogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditlmslogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysauditsslog model data
	var SysauditsslogData []model.Sysauditsslog
	if err := global.OWEN_DBList["from"].Find(&SysauditsslogData).Error; err != nil {
		//return err
	}
	for _, record := range SysauditsslogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysblocklist model data
	var SysblocklistData []model.Sysblocklist
	if err := global.OWEN_DBList["from"].Find(&SysblocklistData).Error; err != nil {
		//return err
	}
	for _, record := range SysblocklistData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbookblocklist model data
	var SysbookblocklistData []model.Sysbookblocklist
	if err := global.OWEN_DBList["from"].Find(&SysbookblocklistData).Error; err != nil {
		//return err
	}
	for _, record := range SysbookblocklistData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbooknumlib model data
	var SysbooknumlibData []model.Sysbooknumlib
	if err := global.OWEN_DBList["from"].Find(&SysbooknumlibData).Error; err != nil {
		//return err
	}
	for _, record := range SysbooknumlibData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysbooknumset model data
	var SysbooknumsetData []model.Sysbooknumset
	if err := global.OWEN_DBList["from"].Find(&SysbooknumsetData).Error; err != nil {
		//return err
	}
	for _, record := range SysbooknumsetData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscardconfig model data
	var SyscardconfigData []model.Syscardconfig
	if err := global.OWEN_DBList["from"].Find(&SyscardconfigData).Error; err != nil {
		//return err
	}
	for _, record := range SyscardconfigData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscarddevdtl model data
	var SyscarddevdtlData []model.Syscarddevdtl
	if err := global.OWEN_DBList["from"].Find(&SyscarddevdtlData).Error; err != nil {
		//return err
	}
	for _, record := range SyscarddevdtlData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscardtype model data
	var SyscardtypeData []model.Syscardtype
	if err := global.OWEN_DBList["from"].Find(&SyscardtypeData).Error; err != nil {
		//return err
	}
	for _, record := range SyscardtypeData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscoderule model data
	var SyscoderuleData []model.Syscoderule
	if err := global.OWEN_DBList["from"].Find(&SyscoderuleData).Error; err != nil {
		//return err
	}
	for _, record := range SyscoderuleData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syscoderuleseed model data
	var SyscoderuleseedData []model.Syscoderuleseed
	if err := global.OWEN_DBList["from"].Find(&SyscoderuleseedData).Error; err != nil {
		//return err
	}
	for _, record := range SyscoderuleseedData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysconfigbase model data
	var SysconfigbaseData []model.Sysconfigbase
	if err := global.OWEN_DBList["from"].Find(&SysconfigbaseData).Error; err != nil {
		//return err
	}
	for _, record := range SysconfigbaseData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysconfiglog model data
	var SysconfiglogData []model.Sysconfiglog
	if err := global.OWEN_DBList["from"].Find(&SysconfiglogData).Error; err != nil {
		//return err
	}
	for _, record := range SysconfiglogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdataitem model data
	var SysdataitemData []model.Sysdataitem
	if err := global.OWEN_DBList["from"].Find(&SysdataitemData).Error; err != nil {
		//return err
	}
	for _, record := range SysdataitemData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdataitemdetail model data
	var SysdataitemdetailData []model.Sysdataitemdetail
	if err := global.OWEN_DBList["from"].Find(&SysdataitemdetailData).Error; err != nil {
		//return err
	}
	for _, record := range SysdataitemdetailData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdatalog model data
	var SysdatalogData []model.Sysdatalog
	if err := global.OWEN_DBList["from"].Find(&SysdatalogData).Error; err != nil {
		//return err
	}
	for _, record := range SysdatalogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysdepartment model data
	var SysdepartmentData []model.Sysdepartment
	if err := global.OWEN_DBList["from"].Find(&SysdepartmentData).Error; err != nil {
		//return err
	}
	for _, record := range SysdepartmentData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumfield model data
	var SysenumfieldData []model.Sysenumfield
	if err := global.OWEN_DBList["from"].Find(&SysenumfieldData).Error; err != nil {
		//return err
	}
	for _, record := range SysenumfieldData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumvalue model data
	var SysenumvalueData []model.Sysenumvalue
	if err := global.OWEN_DBList["from"].Find(&SysenumvalueData).Error; err != nil {
		//return err
	}
	for _, record := range SysenumvalueData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysenumvalue2 model data
	var Sysenumvalue2Data []model.Sysenumvalue2
	if err := global.OWEN_DBList["from"].Find(&Sysenumvalue2Data).Error; err != nil {
		//return err
	}
	for _, record := range Sysenumvalue2Data {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceoffineoperationlog model data
	var SysfaceoffineoperationlogData []model.Sysfaceoffineoperationlog
	if err := global.OWEN_DBList["from"].Find(&SysfaceoffineoperationlogData).Error; err != nil {
		//return err
	}
	for _, record := range SysfaceoffineoperationlogData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflinefeature model data
	var SysfaceofflinefeatureData []model.Sysfaceofflinefeature
	if err := global.OWEN_DBList["from"].Find(&SysfaceofflinefeatureData).Error; err != nil {
		//return err
	}
	for _, record := range SysfaceofflinefeatureData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflinegroup model data
	var SysfaceofflinegroupData []model.Sysfaceofflinegroup
	if err := global.OWEN_DBList["from"].Find(&SysfaceofflinegroupData).Error; err != nil {
		//return err
	}
	for _, record := range SysfaceofflinegroupData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysfaceofflineuser model data
	var SysfaceofflineuserData []model.Sysfaceofflineuser
	if err := global.OWEN_DBList["from"].Find(&SysfaceofflineuserData).Error; err != nil {
		//return err
	}
	for _, record := range SysfaceofflineuserData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslanguage model data
	var SyslanguageData []model.Syslanguage
	if err := global.OWEN_DBList["from"].Find(&SyslanguageData).Error; err != nil {
		//return err
	}
	for _, record := range SyslanguageData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslayertran model data
	var SyslayertranData []model.Syslayertran
	if err := global.OWEN_DBList["from"].Find(&SyslayertranData).Error; err != nil {
		//return err
	}
	for _, record := range SyslayertranData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Syslocation model data
	var SyslocationData []model.Syslocation
	if err := global.OWEN_DBList["from"].Find(&SyslocationData).Error; err != nil {
		//return err
	}
	for _, record := range SyslocationData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Sysmenu model data
	var SysmenuData []model.Sysmenu
	if err := global.OWEN_DBList["from"].Find(&SysmenuData).Error; err != nil {
		//return err
	}
	for _, record := range SysmenuData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing SysmenuCopy model data
	var SysmenuCopyData []model.SysmenuCopy
	if err := global.OWEN_DBList["from"].Find(&SysmenuCopyData).Error; err != nil {
		//return err
	}
	for _, record := range SysmenuCopyData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Systasklist model data
	var SystasklistData []model.Systasklist
	if err := global.OWEN_DBList["from"].Find(&SystasklistData).Error; err != nil {
		//return err
	}
	for _, record := range SystasklistData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Systenantconfig model data
	var SystenantconfigData []model.Systenantconfig
	if err := global.OWEN_DBList["from"].Find(&SystenantconfigData).Error; err != nil {
		//return err
	}
	for _, record := range SystenantconfigData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	// Syncing Systenantextend model data
	var SystenantextendData []model.Systenantextend
	if err := global.OWEN_DBList["from"].Find(&SystenantextendData).Error; err != nil {
		//return err
	}
	for _, record := range SystenantextendData {
		if err := global.OWEN_DBList["to"].Create(&record).Error; err != nil {
			//return err
		}
	}

	return nil
}
