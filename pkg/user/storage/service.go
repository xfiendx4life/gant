package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xfiendx4life/gant/pkg/models"
)

type Postgres struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

// Creates new connection with config
func New(ctx context.Context, uri string) (UserStorage, error) {
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("can't create config for postgres %s", err)
	}
	//TODO: read config from some config
	config.MaxConnIdleTime = time.Second * 5
	config.MaxConnLifetime = time.Second * 30
	config.MaxConns = 10
	config.MinConns = 5
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("can't create connection pool %s", err)
	}
	return &Postgres{
		pool: pool,
		ctx:  ctx,
	}, nil
}

func (p *Postgres) Create(user *models.User) (id string, err error) {
	select {
	case <-p.ctx.Done():
		err = fmt.Errorf("done with context")
		return
	default:
		r := p.pool.QueryRow(p.ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
			user.Name, user.Email, user.Password)
		r.Scan(&user.ID)
		return user.ID.String(), nil
	}
}

func (p *Postgres) Get(email string) (*models.User, error) {
	return nil, nil
}

func (p *Postgres) Delete(id uuid.UUID) error {
	return nil
}

func (p *Postgres) Edit(id uuid.UUID) error {
	return nil
}
