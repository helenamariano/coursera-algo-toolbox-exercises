package main

import "testing"

func Test_getMinNumberOfCoins(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"2", args{amount: 2}, 2, false},
		{"28", args{amount: 28}, 6, false},
		{"1", args{amount: 1}, 1, false},
		{"10000", args{amount: 10000}, 1000, false},
		{"0", args{amount: 0}, 0, true},
		{"10001", args{amount: 10001}, 0, true},
		{"-1", args{amount: -1}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMinNumberOfCoins(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMinNumberOfCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMinNumberOfCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
