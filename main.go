package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type CLI struct {
	Teams       []string      `help:"All of the teams to emit metrics for" env:"TEAMS"`
	MaxSize     uint          `help:"The largest that the metric can grow" default:"100000"`
	MaxIncrease int           `help:"The highest that the metric can be increased by" default:"10000"`
	Interval    time.Duration `help:"The interval at which the metric increases" default:"5s"`
	ServerName  string        `help:"The name of the server" optional:""`
}

var (
	teamVals map[string]uint64 = map[string]uint64{}

	processMemoryConsumptionGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "process_memory_total_bytes",
		Help: "The total amount of memory being used by this process",
	}, []string{"server", "team"})
)

func recordMetrics(cli *CLI) {
	for {
		for _, team := range cli.Teams {
			toIncrease := rand.Intn(cli.MaxIncrease)
			if teamVals[team]+uint64(toIncrease) > uint64(cli.MaxSize) {
				teamVals[team] = 0
			}

			teamVals[team] += uint64(toIncrease)
			processMemoryConsumptionGauge.WithLabelValues(cli.ServerName, team).Set(float64(teamVals[team]))
		}

		time.Sleep(cli.Interval)
	}
}

func metricsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Hardcoded metrics for demonstration
	metrics := map[string]int{
		"cpu_usage":    75, // 75% CPU usage
		"memory_usage": 40, // 40% Memory usage
}

for label, value := range metrics {
	var message string
	if value > 50 {
		message = fmt.Sprintf("High %s: %d%%. Consider taking action.\n", label, value)
	} else {
		message = fmt.Sprintf("Normal %s: %d%%. System is stable.\n", label, value)
	}
	// Write each message to the response
	fmt.Fprint(w, message)
}
}

func DiscordWebhook(webhookURL string, message string) error {
	content := map[string]string{"content": message}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("discord webhook error: %s", string(body))
	}

	return nil
}

func main() {
	// Temporary names for illustration. Replace with actual URLs and paths.
	metricsURL := "https://example.com/metrics"
	discordWebhookURL := os.Getenv("https://discord.com/api/webhooks/1222659972874371222/STTaFd0r8uTnfAeskLv_6xgIT41dm-SbmvnxkgIuT_ZZZLP0UDkavgBmJm60qJNTNSIU") // Set this in your environment or replace with the actual URL

	// Read metrics
	metric, err := ReadMetrics(metricsURL)
	if err != nil {
		fmt.Println("Error reading metrics:", err)
		return
	}

	// Send an alert to Discord
	message := fmt.Sprintf("Alert: Metric value is %d", metric.Value)
	if err := DiscordWebhook(discordWebhookURL, message); err != nil {
		fmt.Println("Error sending alert to Discord:", err)
		return
	}

	fmt.Println("Alert sent to Discord successfully.")
}

//function to record metrics
func recordMetrics(cli *CLI) {
	for {
		for _, team := range cli.Teams {
			toIncrease := rand.Intn(cli.MaxIncrease)
			if teamVals[team]+uint64(toIncrease) > uint64(cli.MaxSize) {
				teamVals[team] = 0
			}

			teamVals[team] += uint64(toIncrease)
			processMemoryConsumptionGauge.WithLabelValues(cli.ServerName, team).Set(float64(teamVals[team]))
		}

		time.Sleep(cli.Interval)
	}
}

//function to send metrics to discord
func sendMetricsToDiscord(cli *CLI) {
	for {
		for _, team := range cli.Teams {
			toIncrease := rand.Intn(cli.MaxIncrease)
			if teamVals[team]+uint64(toIncrease) > uint64(cli.MaxSize) {
				teamVals[team] = 0
			}

			teamVals[team] += uint64(toIncrease)
			processMemoryConsumptionGauge.WithLabelValues(cli.ServerName, team).Set(float64(teamVals[team]))
		}

		time.Sleep(cli.Interval)
	}
}
