package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"/push"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

// func Test_push(t *testing.T) {
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			push(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_pop(t *testing.T) {
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			pop(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_dbConn(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *sql.DB
// 	}{

// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := dbConn(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("dbConn() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
