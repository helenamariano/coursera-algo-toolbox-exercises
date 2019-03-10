package main

import (
	"reflect"
	"testing"
)

func Test_getValuesForSum(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"987654321", args{sum: 987654321}, 44443, false},
		{"12345678910", args{sum: 12345678910}, 157134, false},
		{"2673516735757", args{sum: 2673516735757}, 2312364, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getValuesForSum(tt.args.sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("getValuesForSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValuesForSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
