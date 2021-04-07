package repository

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mjhjj/Tyrannosaurus/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSelectAllPlaces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewPlaceSQLite(db)
	type args struct {
		places []domain.Place
	}
	type mockBehavior func()

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		want         args
		wantErr      bool
	}{
		{
			name: "OK",
			want: args{[]domain.Place{{"1", "1", "1", "1", "1", "1", "1", "1"}, {"2", "2", "2", "2", "2", "2", "2", "2"}}},
			mockBehavior: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "x", "y", "name", "address", "about", "bio", "link"}).
					AddRow(1, "1", "1", "1", "1", "1", "1", "1").
					AddRow(2, "2", "2", "2", "2", "2", "2", "2")

				mock.ExpectQuery("SELECT id, x, y, name, address, about, bio, link FROM places;").WillReturnRows(rows)

				mock.ExpectCommit()

			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := r.SelectAllPlaces()
			log.Println(got)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
