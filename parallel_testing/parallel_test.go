package main

import (
	"testing"
	"time"
)

func Test_Slow1(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Name     string
		Duration time.Duration
	}{
		{Name: "Case 1", Duration: time.Second * 1},
		{Name: "Case 2", Duration: time.Second * 1},
		{Name: "Case 3", Duration: time.Second * 1},
		{Name: "Case 4", Duration: time.Second * 1},
		{Name: "Case 5", Duration: time.Second * 1},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			t.Logf("%s sleeping..", tc.Name)
			sleepFn(tc.Duration)
			t.Logf("%s slept", tc.Name)
		})
	}
}
func sleepFn(duration time.Duration) {
	time.Sleep(duration)
}
