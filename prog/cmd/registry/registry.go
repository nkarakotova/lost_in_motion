package registry

import (
	"os"
	config "prog/config"
	transaction_manager_implementation "prog/internal/managers/implementation"
	transaction_manager "prog/internal/managers"
	at_manager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"prog/internal/repositories"
	postgres_repo "prog/internal/repositories/postgreSQL"
	"prog/internal/services"
	servicesImplementation "prog/internal/services/implementation"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

type AppServiceFields struct {
	ClientService services.ClientService
	CoachService services.CoachService
	DirectionService services.DirectionService
	HallService services.HallService
	SubscriptionService services.SubscriptionService
	TrainingService services.TrainingService
}

type AppRepositoryFields struct {
	ClientRepository repositories.ClientRepository
	CoachRepository repositories.CoachRepository
	DirectionRepository repositories.DirectionRepository
	HallRepository repositories.HallRepository
	SubscriptionRepository repositories.SubscriptionRepository
	TrainingRepository repositories.TrainingRepository
}

type AppManagerFields struct {
	TransactionManager transaction_manager.TransactionManager
}

type App struct {
	Config       config.Config
	Repositories *AppRepositoryFields
	Managers     *AppManagerFields
	Services     *AppServiceFields
	Logger       *log.Logger
}

func (a *App) initRepositories(fields *postgres_repo.PostgresRepositoryFields) *AppRepositoryFields {
	f := &AppRepositoryFields{
		ClientRepository: postgres_repo.CreateClientPostgreSQLRepository(fields),
		CoachRepository: postgres_repo.CreateCoachPostgreSQLRepository(fields),
		DirectionRepository: postgres_repo.CreateDirectionPostgreSQLRepository(fields),
		HallRepository: postgres_repo.CreateHallPostgreSQLRepository(fields),
		SubscriptionRepository: postgres_repo.CreateSubscriptionPostgreSQLRepository(fields),
		TrainingRepository: postgres_repo.CreateTrainingPostgreSQLRepository(fields),
	}

	a.Logger.Info("Success initialization of repositories")
	return f
}

func (a *App) initManagers(manager *at_manager.Manager) *AppManagerFields {
	f := &AppManagerFields{
		TransactionManager: transaction_manager_implementation.NewTransactionManagerImplementation(manager),
	}

	a.Logger.Info("Success initialization of repositories")
	return f
}

func (a *App) initServices(r *AppRepositoryFields, m *AppManagerFields) *AppServiceFields {
	f := &AppServiceFields{
		ClientService: servicesImplementation.NewClientServiceImplementation(r.ClientRepository, r.TrainingRepository, r.DirectionRepository, r.SubscriptionRepository, m.TransactionManager, a.Logger),
		CoachService: servicesImplementation.NewCoachServiceImplementation(r.CoachRepository, r.TrainingRepository, a.Logger),
		DirectionService: servicesImplementation.NewDirectionServiceImplementation(r.DirectionRepository, a.Logger),
		HallService: servicesImplementation.NewHallServiceImplementation(r.HallRepository, r.TrainingRepository, a.Logger),
		SubscriptionService: servicesImplementation.NewSubscriptionServiceImplementation(r.SubscriptionRepository, a.Logger),
		TrainingService: servicesImplementation.NewTrainingServiceImplementation(r.TrainingRepository, r.ClientRepository, r.SubscriptionRepository, r.HallRepository, m.TransactionManager, a.Logger),
	}

	a.Logger.Info("Success initialization of services")
	return f
}

func (a *App) Init() error {
	a.initLogger()

	fields, err := postgres_repo.CreatePostgresRepositoryFields(a.Config.Postgres, a.Logger)
	if err != nil {
		a.Logger.Fatal("Error create postgres repository fields", "err", err)
		return err
	}

	dbx := sqlx.NewDb(fields.DB, "pgx")
	manager, err := at_manager.New(trmsqlx.NewDefaultFactory(dbx))
	if err != nil {
		return err
	}

	a.Repositories = a.initRepositories(fields)
	a.Managers = a.initManagers(manager)
	a.Services = a.initServices(a.Repositories, a.Managers)

	return nil
}

func (a *App) initLogger() {
	f, err := os.OpenFile(a.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger := log.New(f)

	log.SetFormatter(log.LogfmtFormatter)
	Logger.SetReportTimestamp(true)
	Logger.SetReportCaller(true)

	if a.Config.LogLevel == "debug" {
		Logger.SetLevel(log.DebugLevel)
	} else if a.Config.LogLevel == "info" {
		Logger.SetLevel(log.InfoLevel)
	} else {
		log.Fatal("Error log level")
	}

	Logger.Print("\n")
	Logger.Info("Success initialization of new Logger!")

	a.Logger = Logger
}

func (a *App) Run() error {
	err := a.Init()

	if err != nil {
		a.Logger.Error("Error init app", "err", err)
		return err
	}

	return nil
}
