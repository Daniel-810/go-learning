package main

import "fmt"

func main() {
	const namespace = "production"
	var serviceName string = "api-server"
	replicas, cpuPerPod, healthy := 3, 0.5, true

	totalCPU := float64(replicas) * cpuPerPod

	var (
		a int
		b float64
		c string
		d bool
	)

	fmt.Printf("[%v] %v\n", namespace, serviceName)
	fmt.Printf("replicas=%v (%T), cpuPerPod=%v (%T), healthy=%v (%T)\n", replicas, replicas, cpuPerPod, cpuPerPod, healthy, healthy)
	fmt.Printf("totalCPU=%.2f cores\n", totalCPU)
	fmt.Printf("zero values: %v %v %q %v\n", a, b, c, d)
}
