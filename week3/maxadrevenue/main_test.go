package main

import "testing"

func Test_maxAdRevenue(t *testing.T) {
	type args struct {
		profitPerClicks []int
		clicksPerDays   []int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"1", args{profitPerClicks: []int{23}, clicksPerDays: []int{39}}, 897, false},
		{"3", args{profitPerClicks: []int{1, 3, -5}, clicksPerDays: []int{-2, 4, 1}}, 23, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxAdRevenue(tt.args.profitPerClicks, tt.args.clicksPerDays)
			if (err != nil) != tt.wantErr {
				t.Errorf("maxAdRevenue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("maxAdRevenue() = %v, want %v", got, tt.want)
			}
		})
	}
}
