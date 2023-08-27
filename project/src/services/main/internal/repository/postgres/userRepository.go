package postgres

import (
	"fmt"
	"math"
	"project/src/pkg/utils"
	"project/src/services/main/domain/model"
	"project/src/services/main/domain/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) GetUser(ctx *gin.Context, data *usecases.GetUserParams) (*usecases.GetUserResult, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer r.CloseConnection(ctx)
	var u model.User
	var queryParam any
	var whereQuery string = ""
	if data.Email != nil {
		whereQuery = " and u.email = $1 "
		queryParam = *data.Email
	} else if data.Id != nil && data.Email == nil && data.Uuid == nil {
		whereQuery = " and u.id = $1 "
		queryParam = *data.Id
	} else if data.Uuid != nil && data.Id == nil && data.Email == nil {
		whereQuery = " and u.uuid = $1 "
		queryParam = *data.Uuid
	}
	query := fmt.Sprintf(
		`
      select u.id, u.version, u.uuid, u.name,
      u.email, u.created_at, u.updated_at
      from public.users u
      where 1=1
      %s
      limit 1
    `, whereQuery)

	row := r.Db.QueryRow(ctx, query,
		queryParam,
	)
	err = row.Scan(
		&u.Base.ID,
		&u.Base.Version,
		&u.Base.Uuid,
		&u.Name,
		&u.Email,
		&u.Base.CreatedAt,
		&u.Base.UpdatedAt,
	)
	if err != nil {
		return nil, utils.NewNotFoundError()
	}
	return &usecases.GetUserResult{
		Data: &u,
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
		return nil, utils.NewNotFoundError()
	}
	// * incorrect password
	err = utils.ValidatePassword(data.Password, u.Password)
	if err != nil {
		return nil, utils.NewWrongPasswordError()
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
			err = tx.Rollback(ctx)
			if err != nil {
				utils.FatalResult("", err)
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				utils.FatalResult("", err)
			}
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
			err = tx.Rollback(ctx)
			if err != nil {
				utils.FatalResult("", err)
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				utils.FatalResult("", err)
			}
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

func (r *Repository) GetUsers(ctx *gin.Context, params *usecases.GetUsersParams) (*usecases.GetUsersResult, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer r.CloseConnection(ctx)

	items := []*model.User{}
	result := usecases.GetUsersResult{
		CurrentPage:       0,
		TotalPages:        0,
		TotalItems:        0,
		TotalItemsPerPage: params.Limit,
		Items:             items,
	}
	queryCount, query, enableWhere, detail := "", "", "", ""
	if *params.Status == "enabled" {
		enableWhere = ` and u.deleted_at is null `
	} else {
		enableWhere = ` and u.deleted_at is not null `
	}

	queryCount = `select count(*) as total
    from public.users u
    where 1=1` + enableWhere
	query = `select u.id, u.uuid, u.name, u.email,
    u.created_at, u.updated_at, u.deleted_at
    from public.users u
    where 1=1` + enableWhere +
		`order by u.id desc limit $1 offset $2`

	totalItems := 0
	row := r.Db.QueryRow(ctx, queryCount)
	err = row.Scan(&totalItems)
	if err != nil {
		detail = err.Error()
		return nil, utils.NewServerError(&detail)
	}
	result.TotalItems = uint32(totalItems)
	var offset uint32
	if params.Page == 1 {
		offset = 0
	} else {
		offset = (params.Page - 1) * params.Limit
	}

	rows, err := r.Db.Query(ctx, query, params.Limit, offset)
	if err != nil {
		detail = err.Error()
		return nil, utils.NewServerError(&detail)
	}

	for rows.Next() {
		var u model.User
		if err = rows.Scan(
			&u.Base.ID,
			&u.Base.Uuid,
			&u.Name,
			&u.Email,
			&u.Base.CreatedAt,
			&u.Base.UpdatedAt,
			&u.Base.DeletedAt,
		); err != nil {
			detail = err.Error()
			return nil, utils.NewServerError(&detail)
		}
		result.Items = append(result.Items, &u)
	}

	result.CurrentPage = params.Page
	result.TotalPages = uint32(math.Ceil(float64(totalItems) / float64(params.Limit)))

	return &result, nil
}
