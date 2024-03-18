package api_usecase

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func (u UseCase) CreateActor(params api_models.CreateActorParams) (string, error) {
	if params.Birth.Sub(time.Now()) > 0 {
		return "", fmt.Errorf("usecase error: invalid actor birth")
	}
	if len(params.Name) < 1 {
		return "", fmt.Errorf("usecase error: invalid actor name length")
	}
	if !(params.Name[0] <= 'Z' && params.Name[0] >= 'A') {
		return "", fmt.Errorf("usecase error: invalid actor name")
	}
	if params.Sex != common.ACTOR_SEX_MALE && params.Sex != common.ACTOR_SEX_FEMALE {
		return "", fmt.Errorf("usecase error: invalid actor sex")
	}

	actorId, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("usecase error: %s", err.Error())
	}
	params.ActorId = actorId.String()

	err = u.db.CreateActor(params)
	if err != nil {
		return "", fmt.Errorf("usecase error: %s", err.Error())
	}

	return params.ActorId, nil
}

func (u UseCase) GetActors() (api_models.GetActorsResponse, error) {
	response, err := u.db.GetActors()
	if err != nil {
		return api_models.GetActorsResponse{}, fmt.Errorf("usecase error: %s", err.Error())
	}
	return response, nil
}

func (u UseCase) UpdateActor(params api_models.UpdateActorParams) error {
	if params.Birth.Sub(time.Now()) > 0 {
		return fmt.Errorf("usecase error: invalid actor birth")
	}
	if len(params.Name) < 1 {
		return fmt.Errorf("usecase error: invalid actor name length")
	}
	if !(params.Name[0] <= 'Z' && params.Name[0] >= 'A') {
		return fmt.Errorf("usecase error: invalid actor name")
	}
	if params.Sex != common.ACTOR_SEX_MALE && params.Sex != common.ACTOR_SEX_FEMALE {
		return fmt.Errorf("usecase error: invalid actor sex")
	}
	if params.ActorId == "" {
		return fmt.Errorf("usecase error: invalid id")
	}

	err := u.db.UpdateActor(params)
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	return nil
}

func (u UseCase) DeleteActor(params api_models.DeleteActorParams) error {
	if params.ActorId == "" {
		return fmt.Errorf("usecase error: invalid actor id")
	}

	err := u.db.DeleteActor(params.ActorId)
	if err != nil {
		return fmt.Errorf("usecase error: %s", err.Error())
	}

	return nil
}
