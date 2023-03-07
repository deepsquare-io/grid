package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/inputs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var (
	host    net.IP
	port    int
	apiURL  string
	rootCmd = &cobra.Command{
		Use:   "submer-pod-exporter",
		Short: "Prometheus exporter for Submer smart pod.",
		Run: func(cmd *cobra.Command, args []string) {
			ticker := time.NewTicker(time.Second)
			go recordMetrics(cmd.Context(), ticker)

			log.Printf("Serving at %s:%d\n", host, port)
			http.Handle("/metrics", promhttp.Handler())
			if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
				panic(err)
			}
		},
	}
	temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "temperature",
		Help:      "The temperature of the smartpod (°C)",
	})
	consumption = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "consumption",
		Help:      "The consumption of the smartpod (kW)",
	})
	dissipation = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "dissipation",
		Help:      "The dissipation of the smartpod (kW)",
	})
	setpoint = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "setpoint",
		Help:      "The setpoint of the smartpod (°C)",
	})
	mpue = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "mpue",
		Help:      "The mPUE of the smartpod",
	})
	pump1rpm = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "pump1rpm",
		Help:      "The pump1rpm of the smartpod (rotations per minute)",
	})
	pump2rpm = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "pump2rpm",
		Help:      "The pump2rpm of the smartpod (rotations per minute)",
	})
	cti = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "coolant_temperature_in",
		Help:      "The temperature of the coolant going in the heat exchanger of the smartpod (°C)",
	})
	cto = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "coolant_temperature_out",
		Help:      "The temperature of the coolant going out of the heat exchanger of the smartpod (°C)",
	})
	cf = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "coolant_flow",
		Help:      "The flow of the coolant in the heat exchanger of the smartpod (m3/h)",
	})
	wti = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "water_temperature_in",
		Help:      "The temperature of the water going in the heat exchanger of the smartpod (°C)",
	})
	wto = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "water_temperature_out",
		Help:      "The temperature of the water going out of the heat exchanger of the smartpod (°C)",
	})
	wf = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "submer",
		Subsystem: "smartpod",
		Name:      "water_flow",
		Help:      "The flow of the water in the heat exchanger of the smartpod (m3/h)",
	})
)

func recordMetrics(ctx context.Context, ticker *time.Ticker) {
	for {
		func() {
			req, err := http.NewRequest("GET", apiURL, nil)
			if err != nil {
				log.Printf("%+v\n", err)
				setZero()
				return
			}
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			req = req.WithContext(ctx)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("%+v\n", err)
				setZero()
				return
			}
			defer resp.Body.Close()

			data := inputs.RealTime{}
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				log.Printf("%+v\n", err)
				setZero()
				return
			}

			log.Printf("Summary: %+v\n", data.Data)
			temperature.Set(data.Data.Temperature)
			consumption.Set(data.Data.Consumption)
			dissipation.Set(data.Data.Dissipation)
			setpoint.Set(data.Data.Setpoint)
			mpue.Set(data.Data.Mpue)
			pump1rpm.Set(data.Data.Pump1RPM)
			pump2rpm.Set(data.Data.Pump2RPM)
			cti.Set(data.Data.CTI)
			cto.Set(data.Data.CTO)
			cf.Set(data.Data.CF)
			wti.Set(data.Data.WTI)
			wto.Set(data.Data.WTO)
			wf.Set(data.Data.WF)
		}()

		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("Error: %+v\n", err)
				setZero()
			}
			return
		case <-ticker.C:
		}
	}
}

func setZero() {
	temperature.Set(0)
	consumption.Set(0)
	dissipation.Set(0)
	setpoint.Set(0)
	mpue.Set(0)
	pump1rpm.Set(0)
	pump2rpm.Set(0)
	cti.Set(0)
	cto.Set(0)
	cf.Set(0)
	wti.Set(0)
	wto.Set(0)
	wf.Set(0)
}

func init() {
	rootCmd.PersistentFlags().IPVar(&host, "host", net.IPv4zero, "listening host")
	rootCmd.PersistentFlags().IntVar(&port, "port", 3000, "listening port")
	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", "http://localhost/api/realTime", "Submer ssmartpod API URL")
}

func Execute() error {
	return rootCmd.Execute()
}
