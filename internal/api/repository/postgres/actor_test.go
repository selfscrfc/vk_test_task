package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	api_models "vk_test_task/internal/api/models"
)

func TestRepository_CreateActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.CreateActorParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.CreateActorParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.CreateActorParams{
				ActorId: "id1",
				Name:    "name1",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
				mock.ExpectExec(`insert into actor`).
					WithArgs(params.ActorId, params.Name, params.Sex, params.Birth)
			},
			wantErr: false,
		},
		{
			name: "longname",
			args: api_models.CreateActorParams{
				ActorId: "id1",
				Name:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "no actor_id",
			args: api_models.CreateActorParams{
				ActorId: "",
				Name:    "name1",
				Sex:     1,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
		{
			name: "invalid sex",
			args: api_models.CreateActorParams{
				ActorId: "id1",
				Name:    "name1",
				Sex:     999,
				Birth:   time.Now(),
			},
			mockBehaviour: func(params api_models.CreateActorParams) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.CreateActor(testCase.args)

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

func TestRepository_GetActors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	t.Run("default", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "sex", "birth", "films"}).AddRow("", "", "", "", "")
		mock.ExpectQuery(`select actor.*, array_agg`).WithoutArgs().WillReturnRows(rows)

		_, err = r.GetActors()

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestRepository_DeleteActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.DeleteActorParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.DeleteActorParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.DeleteActorParams{
				ActorId: "id1",
			},
			mockBehaviour: func(params api_models.DeleteActorParams) {
				mock.ExpectBegin()

				mock.ExpectExec("delete from film_actor").
					WithArgs(params.ActorId).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("delete from actor").
					WithArgs(params.ActorId).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "no actor_id",
			args: api_models.DeleteActorParams{
				ActorId: "",
			},
			mockBehaviour: func(params api_models.DeleteActorParams) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.DeleteActor(testCase.args.ActorId)

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

func TestRepository_UpdateActor(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.UpdateActorParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.UpdateActorParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.UpdateActorParams{
				ActorId: "id1",
				Name:    "name",
				Sex:     1,
				Birth:   time.Unix(202020, 0),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: false,
		},
		{
			name: "default",
			args: api_models.UpdateActorParams{
				ActorId: "id1",
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: false,
		},
		{
			name: "no actor_id",
			args: api_models.UpdateActorParams{
				ActorId: "",
				Name:    "name",
				Sex:     1,
				Birth:   time.Unix(202020, 0),
			},
			mockBehaviour: func(params api_models.UpdateActorParams) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.UpdateActor(testCase.args)

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
