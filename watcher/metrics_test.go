package watcher

import "testing"

func TestCalcDiscomfortIndex(t *testing.T) {
	type args struct {
		temp float64
		hum  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"basic", args{29, 70}, 79.9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcDiscomfortIndex(tt.args.temp, tt.args.hum); got != tt.want {
				t.Errorf("DiscomfortIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
