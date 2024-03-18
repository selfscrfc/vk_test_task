package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	api_models "vk_test_task/internal/api/models"
	"vk_test_task/internal/common"
)

func TestRepository_CreateFilm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.CreateFilmParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.CreateFilmParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.CreateFilmParams{
				FilmId:      "id",
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
				mock.ExpectBegin()

				mock.ExpectExec("insert into film").
					WithArgs(params.Name, params.Description, params.ReleaseDate, params.Rate, params.FilmId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("insert into film_actor").
					WithArgs(params.FilmId, params.Actors[0], params.Actors[1]).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "no film_id",
			args: api_models.CreateFilmParams{
				FilmId:      "",
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
			},
			wantErr: true,
		},
		{
			name: "no actors",
			args: api_models.CreateFilmParams{
				FilmId:      "id",
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{},
			},
			mockBehaviour: func(params api_models.CreateFilmParams) {
				mock.ExpectBegin()

				mock.ExpectExec("insert into film").
					WithArgs(params.Name, params.Description, params.ReleaseDate, params.Rate, params.FilmId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.CreateFilm(testCase.args)

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

func TestRepository_GetFilms(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.GetFilmsParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.GetFilmsParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.GetFilmsParams{
				SortBy:      common.SORT_FILM_BY_NAME,
				IsAscending: common.SORT_FILM_ASC,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", "", "", "", "")

				mock.ExpectQuery("select film.*, array_agg").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "no sort by",
			args: api_models.GetFilmsParams{
				SortBy:      0,
				IsAscending: common.SORT_FILM_ASC,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", "", "", "", "")

				mock.ExpectQuery("select film.*, array_agg").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "no is asc",
			args: api_models.GetFilmsParams{
				SortBy:      common.SORT_FILM_BY_RELEASE_DATE,
				IsAscending: 0,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", "", "", "", "")

				mock.ExpectQuery("select film.*, array_agg").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "no sort by and no is asc",
			args: api_models.GetFilmsParams{
				SortBy:      0,
				IsAscending: 0,
			},
			mockBehaviour: func(params api_models.GetFilmsParams) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", time.Now(), 10, "", pq.StringArray{})

				mock.ExpectQuery("").WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			_, err = r.GetFilms(testCase.args)

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

func TestRepository_DeleteFilm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.DeleteFilmParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.DeleteFilmParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.DeleteFilmParams{
				FilmId: "id1",
			},
			mockBehaviour: func(params api_models.DeleteFilmParams) {
				mock.ExpectBegin()

				mock.ExpectExec("delete from film_actor").
					WithArgs(params.FilmId).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("delete from film").
					WithArgs(params.FilmId).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "no film_id",
			args: api_models.DeleteFilmParams{
				FilmId: "",
			},
			mockBehaviour: func(params api_models.DeleteFilmParams) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.DeleteFilm(testCase.args.FilmId)

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

func TestRepository_SearchFilmByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(name string)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		fName         string
		wantErr       bool
	}{
		{
			name:  "default",
			fName: "film1",
			mockBehaviour: func(name string) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", time.Now(), 10, "", pq.StringArray{})

				mock.ExpectQuery("").WithArgs("%" + name + "%").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "no fname",
			fName: "",
			mockBehaviour: func(name string) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.fName)

			_, err = r.SearchFilmByName(testCase.fName)

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

func TestRepository_SearchFilmByActorName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(name string)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		fName         string
		wantErr       bool
	}{
		{
			name:  "default",
			fName: "film1",
			mockBehaviour: func(name string) {
				rows := sqlmock.NewRows([]string{"name", "description", "date_released", "rate", "id", "actors"}).
					AddRow("", "", time.Now(), 10, "", pq.StringArray{})

				mock.ExpectQuery("").WithArgs("%" + name + "%").WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "no fname",
			fName: "",
			mockBehaviour: func(name string) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.fName)

			_, err = r.SearchFilmByActorName(testCase.fName)

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

func TestRepository_UpdateFilm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	r := Repository{db: sqlx.NewDb(db, "pgx")}

	type mockBehaviour func(params api_models.UpdateFilmParams)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          api_models.UpdateFilmParams
		wantErr       bool
	}{
		{
			name: "default",
			args: api_models.UpdateFilmParams{
				FilmId: "id1",
				Name: "name" +
					"",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
				mock.ExpectBegin()
			},
			wantErr: false,
		},
		{
			name: "no actors",
			args: api_models.UpdateFilmParams{
				FilmId:      "id1",
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
				mock.ExpectBegin()
			},
			wantErr: false,
		},
		{
			name: "no filmId",
			args: api_models.UpdateFilmParams{
				FilmId:      "",
				Name:        "name",
				Description: "desc",
				ReleaseDate: time.Now(),
				Rate:        10,
				Actors:      []string{"id1", "id2"},
			},
			mockBehaviour: func(params api_models.UpdateFilmParams) {
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)

			err = r.UpdateFilm(testCase.args)

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
