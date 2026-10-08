package main

import "fmt"

func main() {
	var (
		healthyCnt   int
		warningCnt   int
		crashloopCnt int
	)
	const total = 10

	for i := 1; i <= total; i++ {
		restarts := (i * 3) % 7

		var status string
		switch {
		case restarts == 0:
			status = "Healthy"
			healthyCnt++
		case restarts < 3:
			status = "Warning"
			warningCnt++
		default:
			status = "CrashLoop"
			crashloopCnt++
		}
		fmt.Printf("pod-%d: restarts=%d %s\n", i, restarts, status)
	}
	fmt.Printf("Healthy=%d Warning=%d CrashLoop=%d\n", healthyCnt, warningCnt, crashloopCnt)

	if ratio := float64(healthyCnt+warningCnt) / float64(total); ratio >= 0.8 {
		fmt.Printf("cluster OK (%.1f%% healthy)\n", ratio*100)
	} else {
		fmt.Printf("cluster DEGRADED (%.1f%% healthy)\n", ratio*100)
	}

	for ready, step := 0, 0; ready < total; {
		ready += 3
		step++

		fmt.Printf("step %d: ready=%d\n", step, ready)
	}
}
