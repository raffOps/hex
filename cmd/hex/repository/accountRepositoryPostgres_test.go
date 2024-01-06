package repository

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestAccountRepositoryPostgres_Save(t *testing.T) {
	type fields struct {
		conn         *gorm.DB
		customerRepo domain.CustomerRepository
	}
	type args struct {
		account domain.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Account
		wantErr bool
	}{
		{
			name: "should save an account",
			fields: fields{
				conn:         GetPostgresConnection(),
				customerRepo: NewCustomerRepositoryPostgres(GetPostgresConnection()),
			},
			args: args{
				account: domain.Account{
					AccountId:   "99",
					CustomerId:  "1",
					OpeningDate: time.Now(),
					AccountType: "saving",
					Amount:      0,
					Status:      true,
				},
			},
			want: &domain.Account{
				AccountId:   "1001",
				CustomerId:  "1001",
				OpeningDate: time.Now(),
				AccountType: "saving",
				Amount:      0,
				Status:      true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AccountRepositoryPostgres{
				conn:         tt.fields.conn,
				customerRepo: tt.fields.customerRepo,
			}
			got, err := s.Save(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}
