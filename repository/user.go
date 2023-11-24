package repository

import (
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.User
	First(conditions []valueobject.Condition) *entity.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Find(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.User {
	var users []*entity.User
	limit, offset, order := parseClause(clause)
	query := r.db
	log.Println(conditions)
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&users)
	return users
}

func (r *userRepository) First(conditions []valueobject.Condition) *entity.User {
	var user *entity.User
	query := r.db
	for _, condition := range conditions {
		sql := fmt.Sprintf("%s %s $1", condition.Field, condition.Operation)
		query.Where(sql, condition.Value)
	}
	query.First(&user)
	return user
}
