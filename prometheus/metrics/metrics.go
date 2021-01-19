/*  metric 1 : Average Buys -> Number of Buy request ? count :: then with PromSQL we can add the rate ?? 
	metric 2 : Buys by categories 
	metric 3 : Buys by products
*/   

package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promauto"
  )
  //metric 1
  var (
	buyHttpRequest = promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_buy_http_request_total",
			Help: "The total number of buy events",
	}), []string{"path"})
)