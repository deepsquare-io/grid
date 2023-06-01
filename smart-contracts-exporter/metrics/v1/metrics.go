package metricsv1

import (
	"context"
	"os"
	"strings"
	sync "sync"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/logger"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/utils/metric"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/utils/singleton"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const jobsBufferSize = 10

var (
	oldJobs = make([]*struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	}, 0, jobsBufferSize)
	oldJobsMutex  sync.RWMutex
	latestJobChan = make(chan *struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	})
)

func AddJob(job *struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metascheduler.JobDefinition
	Valid            bool
	Cost             metascheduler.JobCost
	Time             metascheduler.JobTime
	JobName          [32]byte
	HasCancelRequest bool
}) {
	oldJobsMutex.Lock()
	defer oldJobsMutex.Unlock()
	if len(oldJobs) == jobsBufferSize {
		oldJobs = oldJobs[1:]
	}
	oldJobs = append(oldJobs, job)

	// Pass latest jobs if possible
	select {
	case latestJobChan <- job: // Send on channel if there is a listener
	default: // Otherwise do nothing
	}
}

func WatchLatest(ctx context.Context) <-chan *struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metascheduler.JobDefinition
	Valid            bool
	Cost             metascheduler.JobCost
	Time             metascheduler.JobTime
	JobName          [32]byte
	HasCancelRequest bool
} {
	out := make(chan *struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	}, 10)

	// Pass old jobs
	oldJobsMutex.RLock()
	for _, job := range oldJobs {
		out <- job
	}
	oldJobsMutex.RUnlock()

	// Pass new jobs
	go func() {
		for {
			select {
			case job := <-latestJobChan:
				out <- job
			case <-ctx.Done():
				close(out)
				return
			}
		}
	}()

	return out
}

var (
	dbVersion                 string
	metaschedulerAddress      string
	checkpointFile            string
	mapTotalNumberOfJobs      = make(map[string]prometheus.Counter)
	TotalNumberOfJobs         func(key string) prometheus.Counter
	mapTotalJobsPending       = make(map[string]prometheus.Gauge)
	TotalJobsPending          func(key string) prometheus.Gauge
	mapTotalJobsMetaScheduled = make(map[string]prometheus.Gauge)
	TotalJobsMetaScheduled    func(key string) prometheus.Gauge
	mapTotalJobsScheduled     = make(map[string]prometheus.Gauge)
	TotalJobsScheduled        func(key string) prometheus.Gauge
	mapTotalJobsRunning       = make(map[string]prometheus.Gauge)
	TotalJobsRunning          func(key string) prometheus.Gauge
	mapTotalJobsCancelled     = make(map[string]prometheus.Gauge)
	TotalJobsCancelled        func(key string) prometheus.Gauge
	mapTotalJobsFinished      = make(map[string]prometheus.Gauge)
	TotalJobsFinished         func(key string) prometheus.Gauge
	mapTotalJobsFailed        = make(map[string]prometheus.Gauge)
	TotalJobsFailed           func(key string) prometheus.Gauge
	mapTotalJobsOutOfCredits  = make(map[string]prometheus.Gauge)
	TotalJobsOutOfCredits     func(key string) prometheus.Gauge
	mapTotalCreditSpent       = make(map[string]prometheus.Counter)
	TotalCreditSpent          func(key string) prometheus.Counter
	mapTotalGPUTime           = make(map[string]prometheus.Counter)
	TotalGPUTime              func(key string) prometheus.Counter
	mapTotalCPUTime           = make(map[string]prometheus.Counter)
	TotalCPUTime              func(key string) prometheus.Counter
	mapTotalJobDuration       = make(map[string]prometheus.Counter)
	TotalJobDuration          func(key string) prometheus.Counter
	LastBlockWatched          prometheus.Gauge
	mapTotalJobRefused        = make(map[string]prometheus.Gauge)
	TotalJobRefused           func(key string) prometheus.Gauge
	mapTotalBilledTooMuch     = make(map[string]prometheus.Counter)
	TotalBilledTooMuch        func(key string) prometheus.Counter
)

func Init(msAddress string, file string, version string) {
	metaschedulerAddress = msAddress
	checkpointFile = file
	dbVersion = version
	TotalNumberOfJobs = singleton.Map(mapTotalNumberOfJobs, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "total",
			Help:      "Total number of jobs",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsPending = singleton.Map(mapTotalJobsPending, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "pending_total",
			Help:      "Total jobs pending",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsMetaScheduled = singleton.Map(
		mapTotalJobsMetaScheduled,
		func(key string) prometheus.Gauge {
			return promauto.NewGauge(prometheus.GaugeOpts{
				Namespace: "metascheduler",
				Subsystem: "jobs",
				Name:      "metascheduled_total",
				Help:      "Total jobs meta-scheduled",
				ConstLabels: prometheus.Labels{
					"wallet_address":        key,
					"metascheduler_address": msAddress,
				},
			})
		},
	)
	TotalJobsScheduled = singleton.Map(mapTotalJobsScheduled, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "scheduled_total",
			Help:      "Total jobs scheduled",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsRunning = singleton.Map(mapTotalJobsRunning, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "running_total",
			Help:      "Total jobs running",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsCancelled = singleton.Map(mapTotalJobsCancelled, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "cancelled_total",
			Help:      "Total jobs cancelled",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsFinished = singleton.Map(mapTotalJobsFinished, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "finished_total",
			Help:      "Total jobs finished",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsFailed = singleton.Map(mapTotalJobsFailed, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "failed_total",
			Help:      "Total jobs failed",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobsOutOfCredits = singleton.Map(
		mapTotalJobsOutOfCredits,
		func(key string) prometheus.Gauge {
			return promauto.NewGauge(prometheus.GaugeOpts{
				Namespace: "metascheduler",
				Subsystem: "jobs",
				Name:      "out_of_credits_total",
				Help:      "Total jobs out of credits",
				ConstLabels: prometheus.Labels{
					"wallet_address":        key,
					"metascheduler_address": msAddress,
				},
			})
		},
	)
	TotalCreditSpent = singleton.Map(mapTotalCreditSpent, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "credits",
			Name:      "total",
			Help:      "Total credits spent on jobs",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalGPUTime = singleton.Map(mapTotalGPUTime, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "gpu",
			Name:      "total_minutes",
			Help:      "Total GPU minutes",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalCPUTime = singleton.Map(mapTotalCPUTime, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "cpu",
			Name:      "total_minutes",
			Help:      "Total CPU minutes",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalJobDuration = singleton.Map(mapTotalJobDuration, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "duration_total_minutes",
			Help:      "Total job minutes",
			ConstLabels: prometheus.Labels{
				"wallet_address":        key,
				"metascheduler_address": msAddress,
			},
		})
	})
	LastBlockWatched = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "metascheduler",
		Subsystem: "eth",
		Name:      "block_last",
		Help:      "Last block watched",
		ConstLabels: prometheus.Labels{
			"metascheduler_address": msAddress,
		},
	})
	TotalJobRefused = singleton.Map(mapTotalJobRefused, func(key string) prometheus.Gauge {
		return promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "refused_total",
			Help:      "Total jobs refused",
			ConstLabels: prometheus.Labels{
				"provider_address":      key,
				"metascheduler_address": msAddress,
			},
		})
	})
	TotalBilledTooMuch = singleton.Map(mapTotalBilledTooMuch, func(key string) prometheus.Counter {
		return promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "metascheduler",
			Subsystem: "jobs",
			Name:      "billed_too_much",
			Help:      "Total billed too much",
			ConstLabels: prometheus.Labels{
				"provider_address":      key,
				"metascheduler_address": msAddress,
			},
		})
	})
}

func Save() error {
	protoOldJobs := make([]*Job, 0, len(oldJobs))
	for _, j := range oldJobs {
		protoOldJobs = append(protoOldJobs, MetaToProtoJob(j))
	}
	db := &DB{
		Metrics: &Metrics{
			MetaschedulerAddress:   metaschedulerAddress,
			TotalNumberOfJobs:      dumpCounterMap(mapTotalNumberOfJobs),
			TotalJobsPending:       dumpGaugeMap(mapTotalJobsPending),
			TotalJobsMetaScheduled: dumpGaugeMap(mapTotalJobsMetaScheduled),
			TotalJobsScheduled:     dumpGaugeMap(mapTotalJobsScheduled),
			TotalJobsRunning:       dumpGaugeMap(mapTotalJobsRunning),
			TotalJobsCancelled:     dumpGaugeMap(mapTotalJobsCancelled),
			TotalJobsFinished:      dumpGaugeMap(mapTotalJobsFinished),
			TotalJobsFailed:        dumpGaugeMap(mapTotalJobsFailed),
			TotalJobsOutOfCredits:  dumpGaugeMap(mapTotalJobsOutOfCredits),
			TotalCreditSpent:       dumpCounterMap(mapTotalCreditSpent),
			TotalGpuTime:           dumpCounterMap(mapTotalGPUTime),
			TotalCpuTime:           dumpCounterMap(mapTotalCPUTime),
			TotalJobDuration:       dumpCounterMap(mapTotalJobDuration),
			LastBlockWatched:       metric.GetGaugeValue(LastBlockWatched),
			TotalJobRefused:        dumpGaugeMap(mapTotalJobRefused),
			TotalBilledTooMuch:     dumpCounterMap(mapTotalBilledTooMuch),
		},
		OldJobs: protoOldJobs,
		Version: dbVersion,
	}

	logger.I.Debug("save", zap.Any("metrics", db.Metrics))

	data, err := proto.Marshal(db)
	if err != nil {
		return err
	}

	return os.WriteFile(checkpointFile, data, 0644)
}

func Load() error {
	if _, err := os.Stat(checkpointFile); os.IsNotExist(err) {
		return nil
	}
	logger.I.Info("loading checkpoint...")

	db := &DB{}

	data, err := os.ReadFile(checkpointFile)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, db); err != nil {
		return err
	}

	if db.Version != dbVersion {
		logger.I.Warn("db version is different, will not load checkpoint")
		return nil
	}

	if strings.Contains(db.Version, "dev") {
		logger.I.Warn("db version contains dev, will not load checkpoint")
		return nil
	}

	if db.Metrics.MetaschedulerAddress != metaschedulerAddress {
		logger.I.Warn("contract address is different, will not load checkpoint")
		return nil
	}

	loadMap(db.Metrics.TotalNumberOfJobs, TotalNumberOfJobs)
	loadMap(db.Metrics.TotalJobsPending, TotalJobsPending)
	loadMap(db.Metrics.TotalJobsMetaScheduled, TotalJobsMetaScheduled)
	loadMap(db.Metrics.TotalJobsScheduled, TotalJobsScheduled)
	loadMap(db.Metrics.TotalJobsRunning, TotalJobsRunning)
	loadMap(db.Metrics.TotalJobsCancelled, TotalJobsCancelled)
	loadMap(db.Metrics.TotalJobsFinished, TotalJobsFinished)
	loadMap(db.Metrics.TotalJobsFailed, TotalJobsFailed)
	loadMap(db.Metrics.TotalJobsOutOfCredits, TotalJobsOutOfCredits)
	loadMap(db.Metrics.TotalCreditSpent, TotalCreditSpent)
	loadMap(db.Metrics.TotalGpuTime, TotalGPUTime)
	loadMap(db.Metrics.TotalCpuTime, TotalCPUTime)
	loadMap(db.Metrics.TotalJobDuration, TotalJobDuration)
	LastBlockWatched.Add(db.Metrics.LastBlockWatched)
	loadMap(db.Metrics.TotalJobRefused, TotalJobRefused)
	loadMap(db.Metrics.TotalBilledTooMuch, TotalBilledTooMuch)
	for _, j := range db.OldJobs {
		AddJob(ProtoToMetaJob(j))
	}

	logger.I.Debug("loaded", zap.Any("metrics", db.Metrics))

	return nil
}

func dumpGaugeMap(m map[string]prometheus.Gauge) (metrics []*LabeledMetric) {
	metrics = make([]*LabeledMetric, 0, len(m))
	for key, collector := range m {
		value := metric.GetGaugeValue(collector)
		metrics = append(metrics, &LabeledMetric{
			Key:   key,
			Value: value,
		})
	}
	return metrics
}

func dumpCounterMap(m map[string]prometheus.Counter) (metrics []*LabeledMetric) {
	metrics = make([]*LabeledMetric, 0, len(m))
	for key, collector := range m {
		value := metric.GetCounterValue(collector)
		metrics = append(metrics, &LabeledMetric{
			Key:   key,
			Value: value,
		})
	}
	return metrics
}

func loadMap[T prometheus.Counter](metrics []*LabeledMetric, getter func(key string) T) {
	for _, metric := range metrics {
		getter(metric.GetKey()).Add(metric.GetValue())
	}
}
