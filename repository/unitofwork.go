package repository

import "context"

type UnitOfWork interface {
	Run(ctx context.Context, runner func(c context.Context) error) error
}
