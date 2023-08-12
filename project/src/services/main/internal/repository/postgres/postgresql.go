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

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Db  *pgx.Conn
	cfg *configs.Config
}

func NewRepository(cfg *configs.Config) (*Repository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	log.Println("database successfully connected")
	return &Repository{Db: conn, cfg: cfg}, nil
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

func (r *Repository) CloseConnection() {
	r.Db.Close(context.Background())
}

func (r *Repository) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {

	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	return &usecases.GetPingResult{
		Data: model.Ping{
			Base: base.Base{
				ID:        1,
				Uuid:      string(newUUID),
				Version:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Message: "Pong Message",
		},
	}, nil
}

func (r *Repository) AddUser(ctx *gin.Context, data *usecases.AddUserParams) (*usecases.AddUserResult, error) {
	var conn *pgx.Conn
	err := r.Db.Ping(ctx)
	if err != nil {
		conn, err = r.OpenDbConnection()
		if err != nil {
			return nil, err
		}
	} else {
		conn = r.Db
	}

	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
		r.CloseConnection()
	}()

	var lastInsertedId int64 = 0
	err = tx.QueryRow(ctx,
		"insert into users (uuid, name, email, password) values ($1, $2, $3, $4) returning id",
		string(newUUID),
		data.Name,
		data.Email,
		string(hashPass),
	).Scan(&lastInsertedId)

	if err != nil {
		return nil, err
	}

	return &usecases.AddUserResult{
		Id: int(lastInsertedId),
	}, nil
}
