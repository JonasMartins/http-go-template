package postgres

import (
	"context"
	"log"
	"project/src/pkg/utils"
	"project/src/services/main/configs"
	"project/src/services/main/domain/model"
	"project/src/services/main/domain/usecases"
	"time"

	base "project/src/pkg/model"
	auth "project/src/services/main/internal/handler/auth"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Repository struct {
	Db           *pgx.Conn
	cfg          *configs.Config
	TokenManager auth.TokenFactory
}

func NewRepository(cfg *configs.Config, tm auth.TokenFactory, u *utils.Utils) (*Repository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	log.Println("database successfully connected")
	migrationsDirectory, err := u.GetFilePath(&[]string{"src", "services", "main", "internal",
		"repository", "postgres", "migrations"})
	if err != nil {
		return nil, err
	}
	path := "file://" + *migrationsDirectory
	m, err := migrate.New(path, cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && err.Error() != utils.ErrMigrationNoChange.Error() {
		return nil, err
	}

	return &Repository{Db: conn, cfg: cfg, TokenManager: tm}, nil
}

func (r *Repository) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &usecases.GetPingResult{
		Data: model.Ping{
			Base: base.Base{
				ID:        1,
				Uuid:      string(newUUID),
				Version:   1,
				CreatedAt: now,
				UpdatedAt: &now,
			},
			Message: "Pong Message",
		},
	}, nil
}
func (r *Repository) GetConnection(ctx *gin.Context) error {
	err := r.Db.Ping(ctx)
	if err != nil {
		r.Db, err = r.OpenDbConnection()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) TestCoonection() error {
	conn, err := pgx.Connect(context.Background(), r.cfg.DB.Conn)
	if err != nil {
		utils.FatalResult("error connecting database", err)
	}
	defer conn.Close(context.Background())
	return nil
}
func (r *Repository) OpenDbConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), r.cfg.DB.Conn)
	if err != nil {
		utils.FatalResult("error connecting database", err)
	}
	return conn, nil
}

func (r *Repository) CloseConnection(ctx *gin.Context) {
	r.Db.Close(ctx)
}
