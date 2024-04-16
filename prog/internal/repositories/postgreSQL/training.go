package postgreSQL

import (
	"context"
	"database/sql"
	"prog/internal/models"
	"prog/internal/repositories"
	"prog/pkg/errors/repositoriesErrors"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type TrainingPostgreSQL struct {
	ID                 uint64    `db:"training_id"`
	CoachID            uint64    `db:"coach_id"`
	HallID             uint64    `db:"hall_id"`
	DirectionID        uint64    `db:"direction_id"`
	Name               string    `db:"name"`
	DateTime           time.Time `db:"date_time"`
	PlacesNum          uint64    `db:"places_num"`
	AvailablePlacesNum uint64    `db:"available_places_num"`
	AcceptableAge      uint16    `db:"acceptable_age"`
}

type TrainingPostgreSQLRepository struct {
	db *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewTrainingPostgreSQLRepository(db *sqlx.DB) repositories.TrainingRepository {
	return &TrainingPostgreSQLRepository{db: db, txResolver: trmsqlx.DefaultCtxGetter}
}

func (t *TrainingPostgreSQLRepository) Create(ctx context.Context, training *models.Training) error {
	query := `insert into trainings(coach_id, hall_id, direction_id, name, date_time, places_num, available_places_num, acceptable_age) values($1, $2, $3, $4, $5, $6, $7, $8) returning training_id;`

	err := t.txResolver.DefaultTrOrDB(ctx, t.db).QueryRowxContext(ctx, query, training.CoachID, training.HallID, training.DirectionID, training.Name, training.DateTime, training.PlacesNum, training.AvailablePlacesNum, training.AcceptableAge).Scan(&training.ID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TrainingPostgreSQLRepository) Delete(ctx context.Context, id uint64) error {
	query := `delete from trainings where training_id=$1 returning training_id;`

	err := t.txResolver.DefaultTrOrDB(ctx, t.db).QueryRowxContext(ctx, query, id).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TrainingPostgreSQLRepository) GetByID(ctx context.Context, id uint64) (*models.Training, error) {
	query := `select * from trainings where training_id=$1;`

	trainingDB := &TrainingPostgreSQL{}
	err := t.txResolver.DefaultTrOrDB(ctx, t.db).GetContext(ctx, trainingDB, query, id)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	trainingModels := &models.Training{}
	err = copier.Copy(trainingModels, trainingDB)
	if err != nil {
		return nil, err
	}

	return trainingModels, nil
}

func (t *TrainingPostgreSQLRepository) GetAllByClient(ctx context.Context, id uint64) ([]models.Training, error) {
	query := `select * from trainings where training_id in (select training_id from clients_trainings where client_id=$1);`

	trainingDB := []TrainingPostgreSQL{}
	err := t.txResolver.DefaultTrOrDB(ctx, t.db).SelectContext(ctx, &trainingDB, query, id)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	trainingModels := []models.Training{}
	for i := range trainingDB {
		training := models.Training{}
		err = copier.Copy(&training, &trainingDB[i])
		if err != nil {
			return nil, err
		}

		trainingModels = append(trainingModels, training)
	}

	return trainingModels, nil
}

func (t *TrainingPostgreSQLRepository) GetAllByCoachOnDate(ctx context.Context, id uint64, date time.Time) ([]models.Training, error) {
	query := `select * from trainings where coach_id=$1 and date_time::date=$2::date;`

	trainingDB := []TrainingPostgreSQL{}
	err := t.txResolver.DefaultTrOrDB(ctx, t.db).SelectContext(ctx, &trainingDB, query, id, date)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	trainingModels := []models.Training{}
	for i := range trainingDB {
		training := models.Training{}
		err = copier.Copy(&training, &trainingDB[i])
		if err != nil {
			return nil, err
		}

		trainingModels = append(trainingModels, training)
	}

	return trainingModels, nil
}

func (t *TrainingPostgreSQLRepository) GetAllByDateTime(ctx context.Context, dateTime time.Time) ([]models.Training, error) {
	query := `select * from trainings where date_time=$1;`

	trainingDB := []TrainingPostgreSQL{}
	err := t.txResolver.DefaultTrOrDB(ctx, t.db).SelectContext(ctx, &trainingDB, query, dateTime)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	trainingModels := []models.Training{}
	for i := range trainingDB {
		training := models.Training{}
		err = copier.Copy(&training, &trainingDB[i])
		if err != nil {
			return nil, err
		}

		trainingModels = append(trainingModels, training)
	}

	return trainingModels, nil
}

func (t *TrainingPostgreSQLRepository) GetAllBetweenDateTime(ctx context.Context, start time.Time, end time.Time) ([]models.Training, error) {
	query := `select * from trainings where date_time between $1 and $2;`

	trainingDB := []TrainingPostgreSQL{}
	err := t.txResolver.DefaultTrOrDB(ctx, t.db).SelectContext(ctx, &trainingDB, query, start, end)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	trainingModels := []models.Training{}
	for i := range trainingDB {
		training := models.Training{}
		err = copier.Copy(&training, &trainingDB[i])
		if err != nil {
			return nil, err
		}

		trainingModels = append(trainingModels, training)
	}

	return trainingModels, nil
}

func (t *TrainingPostgreSQLRepository) ReduceAvailablePlacesNum(ctx context.Context, id uint64) error {
	query := `update trainings set available_places_num = available_places_num - 1 where training_id=$1 returning training_id;`

	err := t.txResolver.DefaultTrOrDB(ctx, t.db).QueryRowxContext(ctx, query, id).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TrainingPostgreSQLRepository) IncreaseAvailablePlacesNum(ctx context.Context, id uint64) error {
	query := `update trainings set available_places_num = available_places_num + 1 where training_id=$1 returning training_id;`

	err := t.txResolver.DefaultTrOrDB(ctx, t.db).QueryRowxContext(ctx, query, id).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
