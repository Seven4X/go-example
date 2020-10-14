package sqlmock

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

/*
call to Query 'INSERT INTO "topic" ("name","tags","create_by","score","agree","disagree","create_time","update_time") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"' with args [{Name: Ordinal:1 Value:t1} {Name: Ordinal:2 Value:} {Name: Ordinal:3 Value:} {Name: Ordinal:4 Value:0} {Name: Ordinal:5 Value:0} {Name: Ordinal:6 Value:0} {Name: Ordinal:7 Value:2020-10-14 17:42:00} {Name: Ordinal:8 Value:2020-10-14 17:42:00}], was not expected, next expectation is: ExpectedExec => expecting Exec or ExecContext which:
  - matches sql: 'INSERT INTO "topic" ("name","tags","create_by","score","agree","disagree","create_time","update_time") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"'
  - is without arguments
  - should return Result having:
      LastInsertId: 15
      RowsAffected: 1
*/
// a successful case
func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO \"topic\"").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStatsNoTx(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a failing test case
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
