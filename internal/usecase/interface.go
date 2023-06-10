package usecase

import (
	"context"

	"github.com/kirakulakov/go-clean-arch/internal/entity"
)

type (
	Translation interface {
		Translation(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}

	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	TranslationWebApi interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)
