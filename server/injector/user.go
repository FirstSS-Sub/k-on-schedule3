package injector

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/config"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"github.com/FirstSS-Sub/k-on-schedule3/server/infrastructure"
	"github.com/FirstSS-Sub/k-on-schedule3/server/interfaces/handler"
	"github.com/FirstSS-Sub/k-on-schedule3/server/usecase"
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
