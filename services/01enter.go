package services

import "GINOWEN/services/manager"

type Managers struct {
	frymanager manager.FlyReadAppManager
}

var ManagerGroup = new(Managers)
