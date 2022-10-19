package domain

import (
	"Banking/errs"
	"Banking/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows // We don't use a sql.Rows object and iterate through it anymore
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, dob, city, zipcode, status from customers"
		err = d.client.Select(&customers, findAllSql) // sqlx queries and marshals at the same time
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, dob, city, zipcode, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status) // sqlx queries and marshals at the same time
		//rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// We don't need to use the sqlx StructScan anymore since the marshalling has been done by sqlx.Select
	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customers: " + err.Error())
	//	return nil, errs.NewUnexpectedError("Unexpected database error")
	//}

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.DOB, &c.City, &c.ZipCode, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers: " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

// Rows -> sqlx.Select QueryRow -> sqlx.Get
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, dob, city, zipcode, status from customers where customer_id = ?"
	//row := d.client.QueryRow(customerSql, id) // This still works as an sqlx.DB object, but we can use Get
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.DOB, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// mysql driver code from https://github.com/go-sql-driver/mysql
	//client, err := sql.Open("mysql", "root:pass123@tcp(localhost:3306)/banking")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	// client, err := sqlx.Open("mysql", "root:pass123@tcp(localhost:3306)/banking")
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
