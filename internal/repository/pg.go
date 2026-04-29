package repository

import (
	"context"
	"database/sql"
	"errors"
	"lentera/internal/model"
	"time"
)

var (
	ErrEmployeeNotFound       = errors.New("employee not found")
	ErrEmployeeCheckInAlready = errors.New("employee already check in")
)

type PgRepo struct {
	DB *sql.DB
}

func (pg *PgRepo) CheckIn(ctx context.Context, req model.AttendaceRequest) (uint, error) {

	var id uint
	err := pg.DB.QueryRowContext(
		ctx,
		`SELECT id FROM attendances WHERE check_in::date = current_date AND employee_id = $1`,
		req.EmployeeId).Scan(&id)

	if err == nil {
		return 0, ErrEmployeeCheckInAlready
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	err = pg.DB.QueryRowContext(ctx,
		`SELECT id FROM employees WHERE id = $1`, req.EmployeeId).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, ErrEmployeeNotFound
	}

	if err != nil {
		return 0, err
	}

	isLate := time.Now().After(time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		9,
		0,
		0,
		0,
		time.UTC))

	var status string
	if isLate {
		status = "LATE"
	} else {
		status = "ON_TIME"
	}

	err = pg.DB.QueryRowContext(
		ctx,
		`INSERT INTO attendances(employee_id, check_in, status) VALUES ($1, $2, $3) RETURNING id`,
		req.EmployeeId,
		time.Now().UTC(),
		status,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil

}
