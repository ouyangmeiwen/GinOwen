package serviceinit

import (
	"log"
	"gorm.io/gorm"
	"GINOWEN/extenddb/model"
)

func CusAutoMigrate(DB *gorm.DB) {
	var err error
	err = DB.AutoMigrate(&model.Abpauditlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpauditlog")
	}
	err = DB.AutoMigrate(&model.Abpbackgroundjob{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpbackgroundjob")
	}
	err = DB.AutoMigrate(&model.Abpedition{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpedition")
	}
	err = DB.AutoMigrate(&model.Abpentitychange{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpentitychange")
	}
	err = DB.AutoMigrate(&model.Abpentitychangeset{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpentitychangeset")
	}
	err = DB.AutoMigrate(&model.Abpentitypropertychange{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpentitypropertychange")
	}
	err = DB.AutoMigrate(&model.Abpfeature{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpfeature")
	}
	err = DB.AutoMigrate(&model.Abplanguage{})

	if err != nil {
		log.Printf("Failed to migrate table: Abplanguage")
	}
	err = DB.AutoMigrate(&model.Abplanguagetext{})

	if err != nil {
		log.Printf("Failed to migrate table: Abplanguagetext")
	}
	err = DB.AutoMigrate(&model.Abpnotification{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpnotification")
	}
	err = DB.AutoMigrate(&model.Abpnotificationsubscription{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpnotificationsubscription")
	}
	err = DB.AutoMigrate(&model.Abporganizationunitrole{})

	if err != nil {
		log.Printf("Failed to migrate table: Abporganizationunitrole")
	}
	err = DB.AutoMigrate(&model.Abporganizationunit{})

	if err != nil {
		log.Printf("Failed to migrate table: Abporganizationunit")
	}
	err = DB.AutoMigrate(&model.Abppermission{})

	if err != nil {
		log.Printf("Failed to migrate table: Abppermission")
	}
	err = DB.AutoMigrate(&model.Abppersistedgrant{})

	if err != nil {
		log.Printf("Failed to migrate table: Abppersistedgrant")
	}
	err = DB.AutoMigrate(&model.Abproleclaim{})

	if err != nil {
		log.Printf("Failed to migrate table: Abproleclaim")
	}
	err = DB.AutoMigrate(&model.Abprole{})

	if err != nil {
		log.Printf("Failed to migrate table: Abprole")
	}
	err = DB.AutoMigrate(&model.Abpsetting{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpsetting")
	}
	err = DB.AutoMigrate(&model.Abptenantnotification{})

	if err != nil {
		log.Printf("Failed to migrate table: Abptenantnotification")
	}
	err = DB.AutoMigrate(&model.Abptenant{})

	if err != nil {
		log.Printf("Failed to migrate table: Abptenant")
	}
	err = DB.AutoMigrate(&model.Abpuseraccount{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuseraccount")
	}
	err = DB.AutoMigrate(&model.Abpuserclaim{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuserclaim")
	}
	err = DB.AutoMigrate(&model.Abpuserloginattempt{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuserloginattempt")
	}
	err = DB.AutoMigrate(&model.Abpuserlogin{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuserlogin")
	}
	err = DB.AutoMigrate(&model.Abpusernotification{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpusernotification")
	}
	err = DB.AutoMigrate(&model.Abpuserorganizationunit{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuserorganizationunit")
	}
	err = DB.AutoMigrate(&model.Abpuserrole{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuserrole")
	}
	err = DB.AutoMigrate(&model.Abpuser{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpuser")
	}
	err = DB.AutoMigrate(&model.Abpusertoken{})

	if err != nil {
		log.Printf("Failed to migrate table: Abpusertoken")
	}
	err = DB.AutoMigrate(&model.Appaliuser{})

	if err != nil {
		log.Printf("Failed to migrate table: Appaliuser")
	}
	err = DB.AutoMigrate(&model.Appapprovalinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Appapprovalinfo")
	}
	err = DB.AutoMigrate(&model.Appapprovaltemplate{})

	if err != nil {
		log.Printf("Failed to migrate table: Appapprovaltemplate")
	}
	err = DB.AutoMigrate(&model.Appbinaryobject{})

	if err != nil {
		log.Printf("Failed to migrate table: Appbinaryobject")
	}
	err = DB.AutoMigrate(&model.Appbookorder{})

	if err != nil {
		log.Printf("Failed to migrate table: Appbookorder")
	}
	err = DB.AutoMigrate(&model.Appchatmessage{})

	if err != nil {
		log.Printf("Failed to migrate table: Appchatmessage")
	}
	err = DB.AutoMigrate(&model.Appcreditloginorder{})

	if err != nil {
		log.Printf("Failed to migrate table: Appcreditloginorder")
	}
	err = DB.AutoMigrate(&model.Appfriendship{})

	if err != nil {
		log.Printf("Failed to migrate table: Appfriendship")
	}
	err = DB.AutoMigrate(&model.Appinvoice{})

	if err != nil {
		log.Printf("Failed to migrate table: Appinvoice")
	}
	err = DB.AutoMigrate(&model.Appitemlocked{})

	if err != nil {
		log.Printf("Failed to migrate table: Appitemlocked")
	}
	err = DB.AutoMigrate(&model.Appmessageboard{})

	if err != nil {
		log.Printf("Failed to migrate table: Appmessageboard")
	}
	err = DB.AutoMigrate(&model.Appnotificationlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Appnotificationlog")
	}
	err = DB.AutoMigrate(&model.Apppayorder{})

	if err != nil {
		log.Printf("Failed to migrate table: Apppayorder")
	}
	err = DB.AutoMigrate(&model.Apppickupcode{})

	if err != nil {
		log.Printf("Failed to migrate table: Apppickupcode")
	}
	err = DB.AutoMigrate(&model.Appqrcode{})

	if err != nil {
		log.Printf("Failed to migrate table: Appqrcode")
	}
	err = DB.AutoMigrate(&model.Apprecommendinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Apprecommendinfo")
	}
	err = DB.AutoMigrate(&model.Appsubbookinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Appsubbookinfo")
	}
	err = DB.AutoMigrate(&model.Appsubscriptionpayment{})

	if err != nil {
		log.Printf("Failed to migrate table: Appsubscriptionpayment")
	}
	err = DB.AutoMigrate(&model.Appusercard{})

	if err != nil {
		log.Printf("Failed to migrate table: Appusercard")
	}
	err = DB.AutoMigrate(&model.Appweuser{})

	if err != nil {
		log.Printf("Failed to migrate table: Appweuser")
	}
	err = DB.AutoMigrate(&model.Dasbusinesscount{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasbusinesscount")
	}
	err = DB.AutoMigrate(&model.Dascirculatecount{})

	if err != nil {
		log.Printf("Failed to migrate table: Dascirculatecount")
	}
	err = DB.AutoMigrate(&model.Dasdatabaselink{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasdatabaselink")
	}
	err = DB.AutoMigrate(&model.Dasdatasource{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasdatasource")
	}
	err = DB.AutoMigrate(&model.Dasfeecount{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasfeecount")
	}
	err = DB.AutoMigrate(&model.Daspatronlogcount{})

	if err != nil {
		log.Printf("Failed to migrate table: Daspatronlogcount")
	}
	err = DB.AutoMigrate(&model.Dasperformance{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasperformance")
	}
	err = DB.AutoMigrate(&model.Dassecuritygatecount{})

	if err != nil {
		log.Printf("Failed to migrate table: Dassecuritygatecount")
	}
	err = DB.AutoMigrate(&model.Dasvisitpage{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasvisitpage")
	}
	err = DB.AutoMigrate(&model.Dasvisittrend{})

	if err != nil {
		log.Printf("Failed to migrate table: Dasvisittrend")
	}
	err = DB.AutoMigrate(&model.Efmigrationshistory{})

	if err != nil {
		log.Printf("Failed to migrate table: Efmigrationshistory")
	}
	err = DB.AutoMigrate(&model.Hangfireaggregatedcounter{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfireaggregatedcounter")
	}
	err = DB.AutoMigrate(&model.Hangfirecounter{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirecounter")
	}
	err = DB.AutoMigrate(&model.Hangfiredistributedlock{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfiredistributedlock")
	}
	err = DB.AutoMigrate(&model.Hangfirehash{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirehash")
	}
	err = DB.AutoMigrate(&model.Hangfirejob{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirejob")
	}
	err = DB.AutoMigrate(&model.Hangfirejobparameter{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirejobparameter")
	}
	err = DB.AutoMigrate(&model.Hangfirejobqueue{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirejobqueue")
	}
	err = DB.AutoMigrate(&model.Hangfirejobstate{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirejobstate")
	}
	err = DB.AutoMigrate(&model.Hangfirelist{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirelist")
	}
	err = DB.AutoMigrate(&model.Hangfireserver{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfireserver")
	}
	err = DB.AutoMigrate(&model.Hangfireset{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfireset")
	}
	err = DB.AutoMigrate(&model.Hangfirestate{})

	if err != nil {
		log.Printf("Failed to migrate table: Hangfirestate")
	}
	err = DB.AutoMigrate(&model.Lcpcommandlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpcommandlog")
	}
	err = DB.AutoMigrate(&model.Lcpconfig{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpconfig")
	}
	err = DB.AutoMigrate(&model.Lcpmaintainlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpmaintainlog")
	}
	err = DB.AutoMigrate(&model.Lcpproduct{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpproduct")
	}
	err = DB.AutoMigrate(&model.Lcprfidantenna{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcprfidantenna")
	}
	err = DB.AutoMigrate(&model.Lcprfidreader{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcprfidreader")
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookaccesslog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpsecuritygatebookaccesslog")
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookdailyaccess{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpsecuritygatebookdailyaccess")
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygateitemlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpsecuritygateitemlog")
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatepatronlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpsecuritygatepatronlog")
	}
	err = DB.AutoMigrate(&model.Lcpserialport{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpserialport")
	}
	err = DB.AutoMigrate(&model.Lcpserialportext{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpserialportext")
	}
	err = DB.AutoMigrate(&model.Lcpservice{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpservice")
	}
	err = DB.AutoMigrate(&model.Lcpterminal{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminal")
	}
	err = DB.AutoMigrate(&model.Lcpterminaladvertisement{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminaladvertisement")
	}
	err = DB.AutoMigrate(&model.Lcpterminalbox{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalbox")
	}
	err = DB.AutoMigrate(&model.Lcpterminalboxitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalboxitem")
	}
	err = DB.AutoMigrate(&model.Lcpterminaldevice{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminaldevice")
	}
	err = DB.AutoMigrate(&model.Lcpterminaldevicelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminaldevicelog")
	}
	err = DB.AutoMigrate(&model.Lcpterminallog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminallog")
	}
	err = DB.AutoMigrate(&model.Lcpterminalpermission{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalpermission")
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelf{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalshelf")
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelfitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalshelfitem")
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelflog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpterminalshelflog")
	}
	err = DB.AutoMigrate(&model.Lcpupgradelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpupgradelog")
	}
	err = DB.AutoMigrate(&model.Lcpversion{})

	if err != nil {
		log.Printf("Failed to migrate table: Lcpversion")
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainbaseinfo")
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainbaseinfoitem")
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoprofile{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainbaseinfoprofile")
	}
	err = DB.AutoMigrate(&model.Libailibrarainknowledgefileinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainknowledgefileinfo")
	}
	err = DB.AutoMigrate(&model.Libailibrarainquestionmetric{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainquestionmetric")
	}
	err = DB.AutoMigrate(&model.Libailibrarainsessionmetric{})

	if err != nil {
		log.Printf("Failed to migrate table: Libailibrarainsessionmetric")
	}
	err = DB.AutoMigrate(&model.Libainirobotinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libainirobotinfo")
	}
	err = DB.AutoMigrate(&model.Libbatchinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libbatchinfo")
	}
	err = DB.AutoMigrate(&model.Libbatchoperateindex{})

	if err != nil {
		log.Printf("Failed to migrate table: Libbatchoperateindex")
	}
	err = DB.AutoMigrate(&model.Libbatchoperatelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libbatchoperatelog")
	}
	err = DB.AutoMigrate(&model.Libbookinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libbookinfo")
	}
	err = DB.AutoMigrate(&model.Libcirculatelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libcirculatelog")
	}
	err = DB.AutoMigrate(&model.LibcirculatelogBak{})

	if err != nil {
		log.Printf("Failed to migrate table: LibcirculatelogBak")
	}
	err = DB.AutoMigrate(&model.Libfeedback{})

	if err != nil {
		log.Printf("Failed to migrate table: Libfeedback")
	}
	err = DB.AutoMigrate(&model.Libfeelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libfeelog")
	}
	err = DB.AutoMigrate(&model.Libinventorystat{})

	if err != nil {
		log.Printf("Failed to migrate table: Libinventorystat")
	}
	err = DB.AutoMigrate(&model.Libinventorytask{})

	if err != nil {
		log.Printf("Failed to migrate table: Libinventorytask")
	}
	err = DB.AutoMigrate(&model.Libinventorywork{})

	if err != nil {
		log.Printf("Failed to migrate table: Libinventorywork")
	}
	err = DB.AutoMigrate(&model.Libinventoryworkdetail{})

	if err != nil {
		log.Printf("Failed to migrate table: Libinventoryworkdetail")
	}
	err = DB.AutoMigrate(&model.Libinventoryworklog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libinventoryworklog")
	}
	err = DB.AutoMigrate(&model.Libitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Libitem")
	}
	err = DB.AutoMigrate(&model.LibitemCopy{})

	if err != nil {
		log.Printf("Failed to migrate table: LibitemCopy")
	}
	err = DB.AutoMigrate(&model.Libiteminventoryinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libiteminventoryinfo")
	}
	err = DB.AutoMigrate(&model.Libiteminventorylog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libiteminventorylog")
	}
	err = DB.AutoMigrate(&model.Libitemlocinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libitemlocinfo")
	}
	err = DB.AutoMigrate(&model.Libitemoperateindexlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libitemoperateindexlog")
	}
	err = DB.AutoMigrate(&model.Libitemoperatelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libitemoperatelog")
	}
	err = DB.AutoMigrate(&model.Libjournalinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Libjournalinfo")
	}
	err = DB.AutoMigrate(&model.Liblabel{})

	if err != nil {
		log.Printf("Failed to migrate table: Liblabel")
	}
	err = DB.AutoMigrate(&model.Liblabeloperatelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Liblabeloperatelog")
	}
	err = DB.AutoMigrate(&model.Liblayer{})

	if err != nil {
		log.Printf("Failed to migrate table: Liblayer")
	}
	err = DB.AutoMigrate(&model.Liblayerindexupdatelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Liblayerindexupdatelog")
	}
	err = DB.AutoMigrate(&model.Libnotificationlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libnotificationlog")
	}
	err = DB.AutoMigrate(&model.Libpartonreservation{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpartonreservation")
	}
	err = DB.AutoMigrate(&model.Libpatron{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpatron")
	}
	err = DB.AutoMigrate(&model.Libpatronitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpatronitem")
	}
	err = DB.AutoMigrate(&model.Libpatronlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpatronlog")
	}
	err = DB.AutoMigrate(&model.Libpointsclearing{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpointsclearing")
	}
	err = DB.AutoMigrate(&model.Libpointslog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libpointslog")
	}
	err = DB.AutoMigrate(&model.Librfidantlayer{})

	if err != nil {
		log.Printf("Failed to migrate table: Librfidantlayer")
	}
	err = DB.AutoMigrate(&model.Librfidscandetaillog{})

	if err != nil {
		log.Printf("Failed to migrate table: Librfidscandetaillog")
	}
	err = DB.AutoMigrate(&model.Librfidscanlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Librfidscanlog")
	}
	err = DB.AutoMigrate(&model.Librow{})

	if err != nil {
		log.Printf("Failed to migrate table: Librow")
	}
	err = DB.AutoMigrate(&model.Librowcatalog{})

	if err != nil {
		log.Printf("Failed to migrate table: Librowcatalog")
	}
	err = DB.AutoMigrate(&model.Libscanitemlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libscanitemlog")
	}
	err = DB.AutoMigrate(&model.Libshelf{})

	if err != nil {
		log.Printf("Failed to migrate table: Libshelf")
	}
	err = DB.AutoMigrate(&model.Libstruct{})

	if err != nil {
		log.Printf("Failed to migrate table: Libstruct")
	}
	err = DB.AutoMigrate(&model.Libtagtobarcodelog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libtagtobarcodelog")
	}
	err = DB.AutoMigrate(&model.Libtaskitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Libtaskitem")
	}
	err = DB.AutoMigrate(&model.Libtaskpackage{})

	if err != nil {
		log.Printf("Failed to migrate table: Libtaskpackage")
	}
	err = DB.AutoMigrate(&model.Libupdatefirstbooklog{})

	if err != nil {
		log.Printf("Failed to migrate table: Libupdatefirstbooklog")
	}
	err = DB.AutoMigrate(&model.Misactivity{})

	if err != nil {
		log.Printf("Failed to migrate table: Misactivity")
	}
	err = DB.AutoMigrate(&model.Mismediainfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Mismediainfo")
	}
	err = DB.AutoMigrate(&model.Misnews{})

	if err != nil {
		log.Printf("Failed to migrate table: Misnews")
	}
	err = DB.AutoMigrate(&model.Mistemplate{})

	if err != nil {
		log.Printf("Failed to migrate table: Mistemplate")
	}
	err = DB.AutoMigrate(&model.Rescatalog{})

	if err != nil {
		log.Printf("Failed to migrate table: Rescatalog")
	}
	err = DB.AutoMigrate(&model.Rescipinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Rescipinfo")
	}
	err = DB.AutoMigrate(&model.Resfourcorner{})

	if err != nil {
		log.Printf("Failed to migrate table: Resfourcorner")
	}
	err = DB.AutoMigrate(&model.Resjournalinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Resjournalinfo")
	}
	err = DB.AutoMigrate(&model.Resnotfound{})

	if err != nil {
		log.Printf("Failed to migrate table: Resnotfound")
	}
	err = DB.AutoMigrate(&model.Respublisherinfo{})

	if err != nil {
		log.Printf("Failed to migrate table: Respublisherinfo")
	}
	err = DB.AutoMigrate(&model.Ssbackgroundjob{})

	if err != nil {
		log.Printf("Failed to migrate table: Ssbackgroundjob")
	}
	err = DB.AutoMigrate(&model.Sysattachment{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysattachment")
	}
	err = DB.AutoMigrate(&model.Sysauditacslog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditacslog")
	}
	err = DB.AutoMigrate(&model.Sysauditapilog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditapilog")
	}
	err = DB.AutoMigrate(&model.Sysauditapplog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditapplog")
	}
	err = DB.AutoMigrate(&model.Sysauditlinklog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditlinklog")
	}
	err = DB.AutoMigrate(&model.Sysauditlmslog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditlmslog")
	}
	err = DB.AutoMigrate(&model.Sysauditsslog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysauditsslog")
	}
	err = DB.AutoMigrate(&model.Sysblocklist{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysblocklist")
	}
	err = DB.AutoMigrate(&model.Sysbookblocklist{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysbookblocklist")
	}
	err = DB.AutoMigrate(&model.Sysbooknumlib{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysbooknumlib")
	}
	err = DB.AutoMigrate(&model.Sysbooknumset{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysbooknumset")
	}
	err = DB.AutoMigrate(&model.Syscardconfig{})

	if err != nil {
		log.Printf("Failed to migrate table: Syscardconfig")
	}
	err = DB.AutoMigrate(&model.Syscarddevdtl{})

	if err != nil {
		log.Printf("Failed to migrate table: Syscarddevdtl")
	}
	err = DB.AutoMigrate(&model.Syscardtype{})

	if err != nil {
		log.Printf("Failed to migrate table: Syscardtype")
	}
	err = DB.AutoMigrate(&model.Syscoderule{})

	if err != nil {
		log.Printf("Failed to migrate table: Syscoderule")
	}
	err = DB.AutoMigrate(&model.Syscoderuleseed{})

	if err != nil {
		log.Printf("Failed to migrate table: Syscoderuleseed")
	}
	err = DB.AutoMigrate(&model.Sysconfigbase{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysconfigbase")
	}
	err = DB.AutoMigrate(&model.Sysconfiglog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysconfiglog")
	}
	err = DB.AutoMigrate(&model.Sysdataitem{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysdataitem")
	}
	err = DB.AutoMigrate(&model.Sysdataitemdetail{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysdataitemdetail")
	}
	err = DB.AutoMigrate(&model.Sysdatalog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysdatalog")
	}
	err = DB.AutoMigrate(&model.Sysdepartment{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysdepartment")
	}
	err = DB.AutoMigrate(&model.Sysenumfield{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysenumfield")
	}
	err = DB.AutoMigrate(&model.Sysenumvalue{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysenumvalue")
	}
	err = DB.AutoMigrate(&model.Sysenumvalue2{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysenumvalue2")
	}
	err = DB.AutoMigrate(&model.Sysfaceoffineoperationlog{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysfaceoffineoperationlog")
	}
	err = DB.AutoMigrate(&model.Sysfaceofflinefeature{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysfaceofflinefeature")
	}
	err = DB.AutoMigrate(&model.Sysfaceofflinegroup{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysfaceofflinegroup")
	}
	err = DB.AutoMigrate(&model.Sysfaceofflineuser{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysfaceofflineuser")
	}
	err = DB.AutoMigrate(&model.Syslanguage{})

	if err != nil {
		log.Printf("Failed to migrate table: Syslanguage")
	}
	err = DB.AutoMigrate(&model.Syslayertran{})

	if err != nil {
		log.Printf("Failed to migrate table: Syslayertran")
	}
	err = DB.AutoMigrate(&model.Syslocation{})

	if err != nil {
		log.Printf("Failed to migrate table: Syslocation")
	}
	err = DB.AutoMigrate(&model.Sysmenu{})

	if err != nil {
		log.Printf("Failed to migrate table: Sysmenu")
	}
	err = DB.AutoMigrate(&model.SysmenuCopy{})

	if err != nil {
		log.Printf("Failed to migrate table: SysmenuCopy")
	}
	err = DB.AutoMigrate(&model.Systasklist{})

	if err != nil {
		log.Printf("Failed to migrate table: Systasklist")
	}
	err = DB.AutoMigrate(&model.Systenantconfig{})

	if err != nil {
		log.Printf("Failed to migrate table: Systenantconfig")
	}
	err = DB.AutoMigrate(&model.Systenantextend{})

	if err != nil {
		log.Printf("Failed to migrate table: Systenantextend")
	}

	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
	}
}
