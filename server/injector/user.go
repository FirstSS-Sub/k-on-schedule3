package injector

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/config"
	"github.com/FirstSS-Sub/k-on-schedule2/server/domain/repository"
	"github.com/FirstSS-Sub/k-on-schedule2/server/infrastructure"
	"github.com/FirstSS-Sub/k-on-schedule2/server/interfaces/handler"
	"github.com/FirstSS-Sub/k-on-schedule2/server/usecase"
)

func InjectUserRepository() repository.UserRepository {
	return infrastructure.NewUserRepository(config.NewDB())
}

func InjectUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(InjectUserRepository())
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}
