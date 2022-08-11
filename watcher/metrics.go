package watcher

import "math"

func CalcDiscomfortIndex(temp float64, hum float64) float64 {
	idx := 0.81*temp + 0.01*hum*(0.99*temp-14.3) + 46.3
	return math.Round(idx*10) / 10
}
