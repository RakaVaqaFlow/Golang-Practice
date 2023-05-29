package internal

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RegUserCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_user",
	Help: "New user was created",
})

var RegTaskCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_task",
	Help: "New task was created",
})

var UpdatedUserCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "updated_user",
	Help: "User that was updated",
})

var UpdatedTaskCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "updated_task",
	Help: "Task that was updated",
})

var DeletedUserCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "deleted_user",
	Help: "User that was deleted",
})

var DeletedTaskCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "deleted_task",
	Help: "Task that was deleted",
})

func init() {
	// Metrics have to be registered to be exposed:

	// Registration metrics
	prometheus.MustRegister(RegUserCounter)
	prometheus.MustRegister(RegTaskCounter)

	// Update metrics
	prometheus.MustRegister(UpdatedUserCounter)
	prometheus.MustRegister(UpdatedTaskCounter)

	// Deletion metrics
	prometheus.MustRegister(DeletedUserCounter)
	prometheus.MustRegister(DeletedTaskCounter)
}
