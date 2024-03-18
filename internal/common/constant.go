package common

const (
	ACTOR_NAME_MAXSIZE = 128
	ACTOR_SEX_MALE     = 1
	ACTOR_SEX_FEMALE   = 2

	FILM_NAME_MAXSIZE        = 150
	FILM_NAME_MINSIZE        = 1
	FILM_DESCRIPTION_MAXSIZE = 1000

	LOGIN_MAXSIZE    = 100
	LOGIN_MINSIZE    = 5
	PASSWORD_MAXSIZE = 100
	PASSWORD_MINSIZE = 5

	SORT_FILM_BY_NAME         = 1
	SORT_FILM_BY_RATE         = 2
	SORT_FILM_BY_RELEASE_DATE = 3
	SORT_FILM_ASC             = 1
	SORT_FILM_DESC            = 2

	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)
