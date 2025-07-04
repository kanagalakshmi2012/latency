package main

import (
	"fmt"
	"time"
)

func setupHosts(nodes int) map[string]struct{} {
	hosts := make(map[string]struct{})
	for i := 1; i <= nodes; i++ {
		hostname := fmt.Sprintf("node%d.cluster.local", i)
		hosts[hostname] = struct{}{}
	}
	return hosts
}

func hostLookup(hosts map[string]struct{}, hostname string) (bool, float64) {
	start := time.Now()
	_, found := hosts[hostname]
	elapsed := time.Since(start).Seconds() * 1000
	return found, elapsed
}

func testCluster(nodes int) {
	fmt.Printf("Cluster size: %d nodes\n", nodes)
	hosts := setupHosts(nodes)
	testHostnames := []string{"node1.cluster.local", fmt.Sprintf("node%d.cluster.local", nodes+1)}

	for _, h := range testHostnames {
		found, duration := hostLookup(hosts, h)
		if found {
			fmt.Printf("Lookup for %s: Found, Time taken = %.3f ms\n", h, duration)
		} else {
			fmt.Printf("Lookup for %s: Not found, Time taken = %.3f ms\n", h, duration)
		}
	}
	fmt.Println()
}

func main() {
	clusterSizes := []int{3, 5, 7, 9, 11}
	for _, size := range clusterSizes {
		testCluster(size)
	}
}
