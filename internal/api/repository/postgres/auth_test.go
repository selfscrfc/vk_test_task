package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_SignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(login string)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		login         string
		wantErr       bool
	}{
		{
			name:  "default",
			login: "login",
			mockBehaviour: func(login string) {
				rows := sqlmock.NewRows([]string{"user_id", "password", "is_admin"}).AddRow("", "", "")
				mock.ExpectQuery(`select user_id, password, is_admin`).
					WithArgs(login).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "no login",
			login: "",
			mockBehaviour: func(login string) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.login)

			_, err := r.SignIn(testCase.login)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				if err = mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
				assert.NoError(t, err)
			}

		})
	}
}

func TestRepository_SignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(login, hashPasword, userId string)

	type args struct {
		login        string
		hashPassword string
		userId       string
	}

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantErr       bool
	}{
		{
			name: "default",
			args: args{
				login:        "login",
				hashPassword: "hp",
				userId:       "userid",
			},
			mockBehaviour: func(login, hashPasword, userId string) {
				rows := sqlmock.NewRows([]string{"exists"}).AddRow("")
				mock.ExpectQuery(`select exists`).
					WithArgs(login).WillReturnRows(rows)

				mock.ExpectExec("insert into \"user\"").
					WithArgs(userId, login, hashPasword, false).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "no login",
			args: args{
				login:        "",
				hashPassword: "hp",
				userId:       "userid",
			},
			mockBehaviour: func(login, hashPasword, userId string) {
			},
			wantErr: true,
		},
		{
			name: "no pass",
			args: args{
				login:        "login",
				hashPassword: "",
				userId:       "userid",
			},
			mockBehaviour: func(login, hashPasword, userId string) {
			},
			wantErr: true,
		},
		{
			name: "no userid",
			args: args{
				login:        "login",
				hashPassword: "hp",
				userId:       "",
			},
			mockBehaviour: func(login, hashPasword, userId string) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args.login, testCase.args.hashPassword, testCase.args.userId)

			err = r.SignUp(testCase.args.login, testCase.args.hashPassword, testCase.args.userId)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				if err = mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
				assert.NoError(t, err)
			}

		})
	}
}
