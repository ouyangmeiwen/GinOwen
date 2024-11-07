package services

import "GINOWEN/repositories"

type Repositories struct {
	orderRepo repositories.GormOrderRepository
	userRepo  repositories.GormUserRepository
}

var RepoApp = new(Repositories)
