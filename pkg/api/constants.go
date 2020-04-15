package api


const (
	PODS string = "/pods"
	RUN string = "/run"
	EXEC string = "/exec"
	ATTACH string = "/attach"
	PORT_FORWARD string = "/portForward"
	CONTAINER_LOGS string = "/containerLogs"
	RUNNING_PODS string = "/runningpods"
	LOGS string = "/pods"
	METRICS string = "/metrics"
	METRICS_CADVISOR string = "/metrics/cadvisor"
	METRICS_PROBES string = "/metrics/probes"
	METRICS_RESOURCE string = "/metrics/resource"
	STATS string = "/stats"
	STATS_CONTAINER string = "/stats/container"
	STATS_SUMMARY string = "/stats/summary"
	SPEC string = "/spec"
	CONFIGZ string = "/configz"
	HEALTHZ string = "/healthz"
	DEBUG string = "/debug/pprof"
	DEBUG_FLAGS string = "/debug/flags/v"
	STD_IN_OUT_TTY string = "input=1&output=1&tty=1"
	CRI string = "/cri/exec"
)
