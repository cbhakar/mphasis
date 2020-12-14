package store

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var (
		dBase *sql.DB
		err   error
	)
	dBase, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", dBase)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	db = s.DB
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetImage() {
	var (
		id   = 1
		name = "test-name"
		time = time.Now().Format("2006-01-02 15:04:05")
	)
	mockedRow := sqlmock.NewRows([]string{"image_id", "image_name", "created_at"}).AddRow(id, name, time)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "images" WHERE (image_id= $1)`)).
		WithArgs(id).
		WillReturnRows(mockedRow)
	res, err := GetImages(QueryDetails{
		ImageId: id,
	})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal([]Image{{
		ImageID:   id,
		ImageName: name,
		CreatedAt: time,
	}}, res))
}
func (s *Suite) TestAddImage() {
	var (
		name = "test-name"
		time = time.Now().Format("2006-01-02 15:04:05")
	)
	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "images" ("image_name","created_at") VALUES ($1,$2) RETURNING "images"."image_id"`)).
		WithArgs(name, time).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	err := AddImage(Image{
		ImageName: name,
		CreatedAt: time,
	})
	require.NoError(s.T(), err)
}
