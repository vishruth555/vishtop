package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"golang.org/x/term"
)

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var cpuUsage, memUsage, netUsage []float64
	var lastNetIn, lastNetOut uint64

	pollMs := flag.Int("poll", 500, "Polling rate in milliseconds")
	flag.Parse()

	width, _, err := term.GetSize(0)
	width -= 7
	if err != nil || width < 10 {
		width = 50 // fallback if unable to get terminal width
	}

	// Get initial network counters
	netStats, _ := net.IOCounters(false)
	if len(netStats) > 0 {
		lastNetIn = netStats[0].BytesRecv
		lastNetOut = netStats[0].BytesSent
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ticker := time.NewTicker(time.Duration(*pollMs) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-c:
			fmt.Println("\nExiting...")
			return
		case <-ticker.C:
			// --- CPU ---
			cpuPercent, _ := cpu.Percent(0, false) // total CPU usage
			if len(cpuPercent) > 0 {
				cpuUsage = append(cpuUsage, cpuPercent[0])
			}

			// --- Memory ---
			vm, _ := mem.VirtualMemory()
			memUsage = append(memUsage, vm.UsedPercent)

			// --- Network (bytes/sec) ---
			netStats, _ := net.IOCounters(false)
			if len(netStats) > 0 {
				in := netStats[0].BytesRecv
				out := netStats[0].BytesSent
				deltaIn := float64(in - lastNetIn)
				deltaOut := float64(out - lastNetOut)
				lastNetIn, lastNetOut = in, out

				// Just track total throughput (in + out, KB/s)
				netUsage = append(netUsage, (deltaIn+deltaOut)/1024.0)
			}

			// Keep sliding window of last 50 points
			if len(cpuUsage) > width {
				cpuUsage = cpuUsage[1:]
				memUsage = memUsage[1:]
				netUsage = netUsage[1:]
			}

			// --- Render graphs ---
			cpuGraph := asciigraph.Plot(cpuUsage, asciigraph.Height(10), asciigraph.Caption("CPU Usage (%)"))
			memGraph := asciigraph.Plot(memUsage, asciigraph.Height(10), asciigraph.Caption("Memory Usage (%)"))
			netGraph := asciigraph.Plot(netUsage, asciigraph.Height(10), asciigraph.Caption("Network Throughput (KB/s)"))

			clearTerminal()
			fmt.Println(cpuGraph)
			fmt.Println(memGraph)
			fmt.Println(netGraph)
		}
	}
}
