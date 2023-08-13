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

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Db           *pgx.Conn
	cfg          *configs.Config
	TokenManager auth.TokenFactory
}

func NewRepository(cfg *configs.Config, tm auth.TokenFactory) (*Repository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	log.Println("database successfully connected")
	return &Repository{Db: conn, cfg: cfg, TokenManager: tm}, nil
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

func (r *Repository) Login(ctx *gin.Context, data *usecases.LoginParams) (*usecases.LoginResult, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer r.CloseConnection(ctx)

	var u model.User
	row := r.Db.QueryRow(ctx,
		` select u.id, u.name, u.email, u.password, u.created_at
      from public.users u
      where 1=1
      and u.email = $1
      and u.deleted_at is null
      limit 1
    `,
		data.Email,
	)
	err = row.Scan(
		&u.Base.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Base.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	// * incorrect password
	err = utils.ValidatePassword(data.Password, u.Password)
	if err != nil {
		return nil, err
	}

	token, err := r.TokenManager.GenerateToken(u.Email, r.cfg.API.TokenDuration)
	if err != nil {
		return nil, err
	}

	return &usecases.LoginResult{
		Token: token,
	}, nil
}

func (r *Repository) UpdateUser(ctx *gin.Context, data *usecases.UpdateUserParams) (*usecases.UpdateUserResult, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := r.Db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
		r.CloseConnection(ctx)
	}()

	var hashPass []byte
	if len(*data.Password) > 0 {
		hashPass, err = bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
	}
	_, err = tx.Exec(ctx,
		`update public.users
    set name = COALESCE($1, name), email = COALESCE($2, email),
    password = COALESCE($3, password), updated_at = COALESCE($4, updated_at)
    where id = $5`,
		data.Name,
		data.Email,
		hashPass,
		time.Now(),
		data.Id,
	)
	if err != nil {
		return nil, err
	}

	return &usecases.UpdateUserResult{
		Id: data.Id,
	}, nil
}

func (r *Repository) AddUser(ctx *gin.Context, data *usecases.AddUserParams) (*usecases.AddUserResult, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	tx, err := r.Db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
		r.CloseConnection(ctx)
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
