package nagiosfoundation

import (
	"fmt"
	"time"
		
	"github.com/shirou/gopsutil/host"
	"github.com/ncr-devops-platform/nagiosfoundation/lib/pkg/nagiosformatters"
)

//Just a tad bit of error handling

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

//getHostUptime simply returns the host uptime as a uint64

func getHostUptime() uint64 {
	hostStat, err := host.Info()
	dealwithErr(err)
	return hostStat.Uptime
}

// CheckUptime gathers information about the host uptime.
func CheckUptime(checkType, warning, critical string, metricName string) (string, int) {
	
	const checkName = "Checkuptime"
	
	var msg string
	var retcode int
	
	uptime := getHostUptime()
	
	warnParse, _ := time.ParseDuration(warning)
	critParse, _ := time.ParseDuration(critical)
	
	warningSecs := warnParse / time.Nanosecond
	criticalSecs := critParse / time.Nanosecond
	
	msg, retcode = nagiosformatters.GreaterFormatNagiosCheck(checkName, float64(uptime), float64(warningSecs), float64(criticalSecs), metricName)
	
	return msg, retcode
}