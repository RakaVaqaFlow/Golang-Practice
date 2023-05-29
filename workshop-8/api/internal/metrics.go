package internal

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RegCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_todos",
	Help: "New todo was created",
})

var DeletedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "deleted_todos",
	Help: "Todo that was deleted",
})

func init() {
	prometheus.MustRegister(RegCounter)
	prometheus.MustRegister(DeletedCounter)
}
