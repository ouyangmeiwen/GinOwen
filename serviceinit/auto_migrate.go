package serviceinit

import (
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
	"log"
)

func (s AutoService) CusAutoMigrate() {
	DB := global.OWEN_DBList["to"]
	var err error
	err = DB.AutoMigrate(&model.Abpauditlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpauditlog")
		return
	}
	err = DB.AutoMigrate(&model.Abpbackgroundjob{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpbackgroundjob")
		return
	}
	err = DB.AutoMigrate(&model.Abpedition{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpedition")
		return
	}
	err = DB.AutoMigrate(&model.Abpentitychange{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpentitychange")
		return
	}
	err = DB.AutoMigrate(&model.Abpentitychangeset{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpentitychangeset")
		return
	}
	err = DB.AutoMigrate(&model.Abpentitypropertychange{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpentitypropertychange")
		return
	}
	err = DB.AutoMigrate(&model.Abpfeature{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpfeature")
		return
	}
	err = DB.AutoMigrate(&model.Abplanguage{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abplanguage")
		return
	}
	err = DB.AutoMigrate(&model.Abplanguagetext{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abplanguagetext")
		return
	}
	err = DB.AutoMigrate(&model.Abpnotification{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpnotification")
		return
	}
	err = DB.AutoMigrate(&model.Abpnotificationsubscription{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpnotificationsubscription")
		return
	}
	err = DB.AutoMigrate(&model.Abporganizationunitrole{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abporganizationunitrole")
		return
	}
	err = DB.AutoMigrate(&model.Abporganizationunit{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abporganizationunit")
		return
	}
	err = DB.AutoMigrate(&model.Abppermission{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abppermission")
		return
	}
	err = DB.AutoMigrate(&model.Abppersistedgrant{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abppersistedgrant")
		return
	}
	err = DB.AutoMigrate(&model.Abproleclaim{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abproleclaim")
		return
	}
	err = DB.AutoMigrate(&model.Abprole{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abprole")
		return
	}
	err = DB.AutoMigrate(&model.Abpsetting{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpsetting")
		return
	}
	err = DB.AutoMigrate(&model.Abptenantnotification{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abptenantnotification")
		return
	}
	err = DB.AutoMigrate(&model.Abptenant{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abptenant")
		return
	}
	err = DB.AutoMigrate(&model.Abpuseraccount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuseraccount")
		return
	}
	err = DB.AutoMigrate(&model.Abpuserclaim{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuserclaim")
		return
	}
	err = DB.AutoMigrate(&model.Abpuserloginattempt{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuserloginattempt")
		return
	}
	err = DB.AutoMigrate(&model.Abpuserlogin{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuserlogin")
		return
	}
	err = DB.AutoMigrate(&model.Abpusernotification{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpusernotification")
		return
	}
	err = DB.AutoMigrate(&model.Abpuserorganizationunit{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuserorganizationunit")
		return
	}
	err = DB.AutoMigrate(&model.Abpuserrole{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuserrole")
		return
	}
	err = DB.AutoMigrate(&model.Abpuser{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpuser")
		return
	}
	err = DB.AutoMigrate(&model.Abpusertoken{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Abpusertoken")
		return
	}
	err = DB.AutoMigrate(&model.Appaliuser{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appaliuser")
		return
	}
	err = DB.AutoMigrate(&model.Appapprovalinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appapprovalinfo")
		return
	}
	err = DB.AutoMigrate(&model.Appapprovaltemplate{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appapprovaltemplate")
		return
	}
	err = DB.AutoMigrate(&model.Appbinaryobject{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appbinaryobject")
		return
	}
	err = DB.AutoMigrate(&model.Appbookorder{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appbookorder")
		return
	}
	err = DB.AutoMigrate(&model.Appchatmessage{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appchatmessage")
		return
	}
	err = DB.AutoMigrate(&model.Appcreditloginorder{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appcreditloginorder")
		return
	}
	err = DB.AutoMigrate(&model.Appfriendship{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appfriendship")
		return
	}
	err = DB.AutoMigrate(&model.Appinvoice{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appinvoice")
		return
	}
	err = DB.AutoMigrate(&model.Appitemlocked{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appitemlocked")
		return
	}
	err = DB.AutoMigrate(&model.Appmessageboard{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appmessageboard")
		return
	}
	err = DB.AutoMigrate(&model.Appnotificationlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appnotificationlog")
		return
	}
	err = DB.AutoMigrate(&model.Apppayorder{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Apppayorder")
		return
	}
	err = DB.AutoMigrate(&model.Apppickupcode{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Apppickupcode")
		return
	}
	err = DB.AutoMigrate(&model.Appqrcode{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appqrcode")
		return
	}
	err = DB.AutoMigrate(&model.Apprecommendinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Apprecommendinfo")
		return
	}
	err = DB.AutoMigrate(&model.Appsubbookinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appsubbookinfo")
		return
	}
	err = DB.AutoMigrate(&model.Appsubscriptionpayment{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appsubscriptionpayment")
		return
	}
	err = DB.AutoMigrate(&model.Appusercard{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appusercard")
		return
	}
	err = DB.AutoMigrate(&model.Appweuser{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Appweuser")
		return
	}
	err = DB.AutoMigrate(&model.Dasbusinesscount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasbusinesscount")
		return
	}
	err = DB.AutoMigrate(&model.Dascirculatecount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dascirculatecount")
		return
	}
	err = DB.AutoMigrate(&model.Dasdatabaselink{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasdatabaselink")
		return
	}
	err = DB.AutoMigrate(&model.Dasdatasource{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasdatasource")
		return
	}
	err = DB.AutoMigrate(&model.Dasfeecount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasfeecount")
		return
	}
	err = DB.AutoMigrate(&model.Daspatronlogcount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Daspatronlogcount")
		return
	}
	err = DB.AutoMigrate(&model.Dasperformance{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasperformance")
		return
	}
	err = DB.AutoMigrate(&model.Dassecuritygatecount{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dassecuritygatecount")
		return
	}
	err = DB.AutoMigrate(&model.Dasvisitpage{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasvisitpage")
		return
	}
	err = DB.AutoMigrate(&model.Dasvisittrend{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Dasvisittrend")
		return
	}
	err = DB.AutoMigrate(&model.Efmigrationshistory{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Efmigrationshistory")
		return
	}
	err = DB.AutoMigrate(&model.Hangfireaggregatedcounter{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfireaggregatedcounter")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirecounter{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirecounter")
		return
	}
	err = DB.AutoMigrate(&model.Hangfiredistributedlock{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfiredistributedlock")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirehash{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirehash")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirejob{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirejob")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirejobparameter{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirejobparameter")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirejobqueue{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirejobqueue")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirejobstate{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirejobstate")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirelist{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirelist")
		return
	}
	err = DB.AutoMigrate(&model.Hangfireserver{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfireserver")
		return
	}
	err = DB.AutoMigrate(&model.Hangfireset{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfireset")
		return
	}
	err = DB.AutoMigrate(&model.Hangfirestate{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Hangfirestate")
		return
	}
	err = DB.AutoMigrate(&model.Lcpcommandlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpcommandlog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpconfig{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpconfig")
		return
	}
	err = DB.AutoMigrate(&model.Lcpmaintainlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpmaintainlog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpproduct{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpproduct")
		return
	}
	err = DB.AutoMigrate(&model.Lcprfidantenna{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcprfidantenna")
		return
	}
	err = DB.AutoMigrate(&model.Lcprfidreader{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcprfidreader")
		return
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookaccesslog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpsecuritygatebookaccesslog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatebookdailyaccess{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpsecuritygatebookdailyaccess")
		return
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygateitemlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpsecuritygateitemlog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpsecuritygatepatronlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpsecuritygatepatronlog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpserialport{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpserialport")
		return
	}
	err = DB.AutoMigrate(&model.Lcpserialportext{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpserialportext")
		return
	}
	err = DB.AutoMigrate(&model.Lcpservice{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpservice")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminal{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminal")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminaladvertisement{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminaladvertisement")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalbox{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalbox")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalboxitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalboxitem")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminaldevice{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminaldevice")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminaldevicelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminaldevicelog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminallog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminallog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalpermission{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalpermission")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelf{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalshelf")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelfitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalshelfitem")
		return
	}
	err = DB.AutoMigrate(&model.Lcpterminalshelflog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpterminalshelflog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpupgradelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpupgradelog")
		return
	}
	err = DB.AutoMigrate(&model.Lcpversion{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Lcpversion")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainbaseinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainbaseinfoitem")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainbaseinfoprofile{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainbaseinfoprofile")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainknowledgefileinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainknowledgefileinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainquestionmetric{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainquestionmetric")
		return
	}
	err = DB.AutoMigrate(&model.Libailibrarainsessionmetric{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libailibrarainsessionmetric")
		return
	}
	err = DB.AutoMigrate(&model.Libainirobotinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libainirobotinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libbatchinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libbatchinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libbatchoperateindex{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libbatchoperateindex")
		return
	}
	err = DB.AutoMigrate(&model.Libbatchoperatelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libbatchoperatelog")
		return
	}
	err = DB.AutoMigrate(&model.Libbookinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libbookinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libcirculatelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libcirculatelog")
		return
	}
	err = DB.AutoMigrate(&model.LibcirculatelogBak{})

	if err != nil {
		log.Fatalf("Failed to migrate table: LibcirculatelogBak")
		return
	}
	err = DB.AutoMigrate(&model.Libfeedback{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libfeedback")
		return
	}
	err = DB.AutoMigrate(&model.Libfeelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libfeelog")
		return
	}
	err = DB.AutoMigrate(&model.Libinventorystat{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libinventorystat")
		return
	}
	err = DB.AutoMigrate(&model.Libinventorytask{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libinventorytask")
		return
	}
	err = DB.AutoMigrate(&model.Libinventorywork{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libinventorywork")
		return
	}
	err = DB.AutoMigrate(&model.Libinventoryworkdetail{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libinventoryworkdetail")
		return
	}
	err = DB.AutoMigrate(&model.Libinventoryworklog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libinventoryworklog")
		return
	}
	err = DB.AutoMigrate(&model.Libitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libitem")
		return
	}
	err = DB.AutoMigrate(&model.LibitemCopy{})

	if err != nil {
		log.Fatalf("Failed to migrate table: LibitemCopy")
		return
	}
	err = DB.AutoMigrate(&model.Libiteminventoryinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libiteminventoryinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libiteminventorylog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libiteminventorylog")
		return
	}
	err = DB.AutoMigrate(&model.Libitemlocinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libitemlocinfo")
		return
	}
	err = DB.AutoMigrate(&model.Libitemoperateindexlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libitemoperateindexlog")
		return
	}
	err = DB.AutoMigrate(&model.Libitemoperatelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libitemoperatelog")
		return
	}
	err = DB.AutoMigrate(&model.Libjournalinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libjournalinfo")
		return
	}
	err = DB.AutoMigrate(&model.Liblabel{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Liblabel")
		return
	}
	err = DB.AutoMigrate(&model.Liblabeloperatelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Liblabeloperatelog")
		return
	}
	err = DB.AutoMigrate(&model.Liblayer{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Liblayer")
		return
	}
	err = DB.AutoMigrate(&model.Liblayerindexupdatelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Liblayerindexupdatelog")
		return
	}
	err = DB.AutoMigrate(&model.Libnotificationlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libnotificationlog")
		return
	}
	err = DB.AutoMigrate(&model.Libpartonreservation{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpartonreservation")
		return
	}
	err = DB.AutoMigrate(&model.Libpatron{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpatron")
		return
	}
	err = DB.AutoMigrate(&model.Libpatronitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpatronitem")
		return
	}
	err = DB.AutoMigrate(&model.Libpatronlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpatronlog")
		return
	}
	err = DB.AutoMigrate(&model.Libpointsclearing{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpointsclearing")
		return
	}
	err = DB.AutoMigrate(&model.Libpointslog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libpointslog")
		return
	}
	err = DB.AutoMigrate(&model.Librfidantlayer{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Librfidantlayer")
		return
	}
	err = DB.AutoMigrate(&model.Librfidscandetaillog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Librfidscandetaillog")
		return
	}
	err = DB.AutoMigrate(&model.Librfidscanlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Librfidscanlog")
		return
	}
	err = DB.AutoMigrate(&model.Librow{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Librow")
		return
	}
	err = DB.AutoMigrate(&model.Librowcatalog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Librowcatalog")
		return
	}
	err = DB.AutoMigrate(&model.Libscanitemlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libscanitemlog")
		return
	}
	err = DB.AutoMigrate(&model.Libshelf{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libshelf")
		return
	}
	err = DB.AutoMigrate(&model.Libstruct{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libstruct")
		return
	}
	err = DB.AutoMigrate(&model.Libtagtobarcodelog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libtagtobarcodelog")
		return
	}
	err = DB.AutoMigrate(&model.Libtaskitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libtaskitem")
		return
	}
	err = DB.AutoMigrate(&model.Libtaskpackage{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libtaskpackage")
		return
	}
	err = DB.AutoMigrate(&model.Libupdatefirstbooklog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Libupdatefirstbooklog")
		return
	}
	err = DB.AutoMigrate(&model.Misactivity{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Misactivity")
		return
	}
	err = DB.AutoMigrate(&model.Mismediainfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Mismediainfo")
		return
	}
	err = DB.AutoMigrate(&model.Misnews{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Misnews")
		return
	}
	err = DB.AutoMigrate(&model.Mistemplate{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Mistemplate")
		return
	}
	err = DB.AutoMigrate(&model.Rescatalog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Rescatalog")
		return
	}
	err = DB.AutoMigrate(&model.Rescipinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Rescipinfo")
		return
	}
	err = DB.AutoMigrate(&model.Resfourcorner{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Resfourcorner")
		return
	}
	err = DB.AutoMigrate(&model.Resjournalinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Resjournalinfo")
		return
	}
	err = DB.AutoMigrate(&model.Resnotfound{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Resnotfound")
		return
	}
	err = DB.AutoMigrate(&model.Respublisherinfo{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Respublisherinfo")
		return
	}
	err = DB.AutoMigrate(&model.Ssbackgroundjob{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Ssbackgroundjob")
		return
	}
	err = DB.AutoMigrate(&model.Sysattachment{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysattachment")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditacslog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditacslog")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditapilog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditapilog")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditapplog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditapplog")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditlinklog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditlinklog")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditlmslog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditlmslog")
		return
	}
	err = DB.AutoMigrate(&model.Sysauditsslog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysauditsslog")
		return
	}
	err = DB.AutoMigrate(&model.Sysblocklist{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysblocklist")
		return
	}
	err = DB.AutoMigrate(&model.Sysbookblocklist{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysbookblocklist")
		return
	}
	err = DB.AutoMigrate(&model.Sysbooknumlib{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysbooknumlib")
		return
	}
	err = DB.AutoMigrate(&model.Sysbooknumset{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysbooknumset")
		return
	}
	err = DB.AutoMigrate(&model.Syscardconfig{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syscardconfig")
		return
	}
	err = DB.AutoMigrate(&model.Syscarddevdtl{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syscarddevdtl")
		return
	}
	err = DB.AutoMigrate(&model.Syscardtype{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syscardtype")
		return
	}
	err = DB.AutoMigrate(&model.Syscoderule{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syscoderule")
		return
	}
	err = DB.AutoMigrate(&model.Syscoderuleseed{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syscoderuleseed")
		return
	}
	err = DB.AutoMigrate(&model.Sysconfigbase{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysconfigbase")
		return
	}
	err = DB.AutoMigrate(&model.Sysconfiglog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysconfiglog")
		return
	}
	err = DB.AutoMigrate(&model.Sysdataitem{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysdataitem")
		return
	}
	err = DB.AutoMigrate(&model.Sysdataitemdetail{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysdataitemdetail")
		return
	}
	err = DB.AutoMigrate(&model.Sysdatalog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysdatalog")
		return
	}
	err = DB.AutoMigrate(&model.Sysdepartment{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysdepartment")
		return
	}
	err = DB.AutoMigrate(&model.Sysenumfield{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysenumfield")
		return
	}
	err = DB.AutoMigrate(&model.Sysenumvalue{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysenumvalue")
		return
	}
	err = DB.AutoMigrate(&model.Sysenumvalue2{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysenumvalue2")
		return
	}
	err = DB.AutoMigrate(&model.Sysfaceoffineoperationlog{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysfaceoffineoperationlog")
		return
	}
	err = DB.AutoMigrate(&model.Sysfaceofflinefeature{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysfaceofflinefeature")
		return
	}
	err = DB.AutoMigrate(&model.Sysfaceofflinegroup{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysfaceofflinegroup")
		return
	}
	err = DB.AutoMigrate(&model.Sysfaceofflineuser{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysfaceofflineuser")
		return
	}
	err = DB.AutoMigrate(&model.Syslanguage{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syslanguage")
		return
	}
	err = DB.AutoMigrate(&model.Syslayertran{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syslayertran")
		return
	}
	err = DB.AutoMigrate(&model.Syslocation{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Syslocation")
		return
	}
	err = DB.AutoMigrate(&model.Sysmenu{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Sysmenu")
		return
	}
	err = DB.AutoMigrate(&model.SysmenuCopy{})

	if err != nil {
		log.Fatalf("Failed to migrate table: SysmenuCopy")
		return
	}
	err = DB.AutoMigrate(&model.Systasklist{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Systasklist")
		return
	}
	err = DB.AutoMigrate(&model.Systenantconfig{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Systenantconfig")
		return
	}
	err = DB.AutoMigrate(&model.Systenantextend{})

	if err != nil {
		log.Fatalf("Failed to migrate table: Systenantextend")
		return
	}

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return
	}
}
