package postgreSQL

import (
	"context"
	"database/sql"
	"prog/internal/models"
	"prog/internal/repositories"
	"prog/pkg/errors/repositoriesErrors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"github.com/jinzhu/copier"
)

type ClientPostgreSQL struct {
	ID             uint64        `db:"client_id"`
	SubscriptionID sql.NullInt64 `db:"subscription_id"`
	Name           string        `db:"name"`
	Telephone      string        `db:"telephone"`
	Mail           string        `db:"mail"`
	Password       string        `db:"password"`
	Age            uint16        `db:"age"`
	Gender         models.Gender `db:"gender"`
   }

type ClientPostgreSQLRepository struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewClientPostgreSQLRepository(db *sqlx.DB) repositories.ClientRepository {
	return &ClientPostgreSQLRepository{db: db, txResolver: trmsqlx.DefaultCtxGetter}
}

func (c *ClientPostgreSQLRepository) Create(ctx context.Context, client *models.Client) error {
	var err error

	if client.SubscriptionID == 0 {
		query := `insert into clients(name, telephone, mail, password, age, gender) values($1, $2, $3, $4, $5, $6) returning client_id;`
		err = c.txResolver.DefaultTrOrDB(ctx, c.db).
			QueryRowxContext(ctx, query, client.Name, client.Telephone, client.Mail, client.Password, client.Age, client.Gender).
			Scan(&client.ID)
	} else {
		query := `insert into clients(subscription_id, name, telephone, mail, password, age, gender) values($1, $2, $3, $4, $5, $6, $7) returning client_id;`
		err = c.txResolver.DefaultTrOrDB(ctx, c.db).
			QueryRowxContext(ctx, query, client.SubscriptionID, client.Name, client.Telephone, client.Mail, client.Password, client.Age, client.Gender).
			Scan(&client.ID)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientPostgreSQLRepository) GetByID(ctx context.Context, id uint64) (*models.Client, error) {
	query := `select * from clients where client_id = $1;`

	clientDB := &ClientPostgreSQL{}
	err := c.txResolver.DefaultTrOrDB(ctx, c.db).GetContext(ctx, clientDB, query, id)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	if !clientDB.SubscriptionID.Valid {
		clientDB.SubscriptionID = sql.NullInt64{Int64: 0, Valid: true}
   }

	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientDB)
	if err != nil {
	 	return nil, err
	}

	return clientModels, nil
}

func (c *ClientPostgreSQLRepository) GetByTelephone(ctx context.Context, telephone string) (*models.Client, error) {
	query := `select * from clients where telephone = $1;`
   
	clientDB := &ClientPostgreSQL{}
	err := c.txResolver.DefaultTrOrDB(ctx, c.db).GetContext(ctx, clientDB, query, telephone)
	if err == sql.ErrNoRows {
	 	return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
	 	return nil, err
	}
   
	if !clientDB.SubscriptionID.Valid {
	 	clientDB.SubscriptionID = sql.NullInt64{Int64: 0, Valid: true}
	}
   
	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientDB)
	if err != nil {
	 	return nil, err
	}
   
	return clientModels, nil
   }

func (c *ClientPostgreSQLRepository) GetByTraining(ctx context.Context, id uint64) ([]models.Client, error) {
	query := `select * from clients where client_id in (select client_id from clients_trainings where training_id=$1);`

	clientDB := []ClientPostgreSQL{}
	err := c.txResolver.DefaultTrOrDB(ctx, c.db).SelectContext(ctx, &clientDB, query, id)
	if err == sql.ErrNoRows {
		return nil, repositoriesErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, err
	}

	clientModels := []models.Client{}
	for i := range clientDB {
		client := models.Client{}
		err = copier.Copy(&client, &clientDB[i])
		if err != nil {
			return nil, err
		}

		clientModels = append(clientModels, client)
	}

	return clientModels, nil
}

func (c *ClientPostgreSQLRepository) СreateAssignment(ctx context.Context, clientID, trainingID uint64) error {
	query := `insert into clients_trainings(client_id, training_id) values($1, $2) returning client_id;`

	err := c.txResolver.DefaultTrOrDB(ctx, c.db).QueryRowxContext(ctx, query, clientID, trainingID).Scan(&clientID)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientPostgreSQLRepository) DeleteAssignment(ctx context.Context, clientID, trainingID uint64) error {
	query := `delete from clients_trainings where client_id=$1 and training_id=$2 returning client_id;`

	err := c.txResolver.DefaultTrOrDB(ctx, c.db).QueryRowxContext(ctx, query, clientID, trainingID).Scan(&clientID)
	if err != nil {
		return err
	}

	return nil
}
