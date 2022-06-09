package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xfiendx4life/gant/pkg/models"
)

type Postgres struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

type doneWithContext struct {
	status string
}

func (d *doneWithContext) Error() string {
	return d.status
}

var simpleContextError = &doneWithContext{
	status: "done with context",
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
		err = simpleContextError
		return
	default:
		r := p.pool.QueryRow(p.ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
			user.Name, user.Email, user.Password)
		r.Scan(&user.ID)
		if user.ID == uuid.Nil {
			return "", fmt.Errorf("can't get id, maybe the row is already exists")
		}
		return user.ID.String(), nil
	}
}

func (p *Postgres) Get(email string) (*models.User, error) {
	select {
	case <-p.ctx.Done():
		return nil, simpleContextError
	default:
		r := p.pool.QueryRow(p.ctx, "SELECT * FROM users WHERE email=$1", email)
		user := models.User{}
		err := r.Scan(&user.ID, &user.Name, &user.Password, &user.Email)
		if user.Name == "" || err != nil {
			return nil, fmt.Errorf("error occured while getting data from db: %s", err)
		}
		return &user, nil
	}
}

func (p *Postgres) Delete(email string) error {
	select {
	case <-p.ctx.Done():
		return simpleContextError
	default:
		_, err := p.pool.Exec(p.ctx, "DELETE FROM users WHERE email=$1", email)
		if err != nil {
			return fmt.Errorf("can't delete row with email %s: %s", email, err)
		}
		return nil
	}
}

func (p *Postgres) Edit(id uuid.UUID, data map[string]string) error {
	select {
	case <-p.ctx.Done():
		return simpleContextError
	default:
		unpacked := make([]string, 0)
		for k, v := range data {
			unpacked = append(unpacked, fmt.Sprintf("%s='%s'", k, v))
		}
		prep := "UPDATE users SET " + strings.Join(unpacked, ", ")
		_, err := p.pool.Exec(p.ctx, prep+"WHERE id=$1", [16]byte(id))
		if err != nil {
			return fmt.Errorf("can't change user with id %s: %s", id, err)
		}
		return nil
	}
}
