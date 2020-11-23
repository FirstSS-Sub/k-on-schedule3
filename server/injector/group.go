package injector

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/config"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"github.com/FirstSS-Sub/k-on-schedule3/server/infrastructure"
	"github.com/FirstSS-Sub/k-on-schedule3/server/interfaces/handler"
	"github.com/FirstSS-Sub/k-on-schedule3/server/usecase"
)

func InjectGroupRepository() repository.GroupRepository {
	return infrastructure.NewGroupRepository(config.NewDB())
}

func InjectGroupUsecase() usecase.GroupUsecase {
	return usecase.NewGroupUsecase(InjectGroupRepository())
}

func InjectGroupHandler() handler.GroupHandler {
	return handler.NewGroupHandler(InjectGroupUsecase())
}
