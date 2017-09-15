package iland

const (
	apiBaseURL = "https://api.ilandcloud.com/ecs"
	accessURL  = "https://console.ilandcloud.com/auth/realms/iland-core/protocol/openid-connect/token"
	refreshURL = "https://console.ilandcloud.com/auth/realms/iland-core/protocol/openid-connect/token"

	TaskStatusSuccess       = "success"
	TaskStatusRunning       = "running"
	TaskStatusError         = "error"
	TaskStatusQueued        = "queued"
	TaskStatusCancelled     = "cancelled"
	TaskStatusWaitingOnUser = "waiting-on-user-input"
	TaskStatusUnknown       = "unknown"

	perfGroupCPU     = "cpu"
	perfGroupMemory  = "mem"
	perfGroupNetwork = "net"
	perfGroupDisk    = "disk"

	perfTypeAverage = "average"
	perfTypeSum     = "summation"
	perfTypeLatest  = "latest"

	PerfIntervalSecond = "second"
	PerfIntervalMinute = "minute"
	PerfIntervalHour   = "hour"
	PerfIntervalDay    = "day"
	PerfIntervalWeek   = "week"
	PerfIntervalMonth  = "month"
)

var PerfCPUUsageAvg = PerfMetric{
	Group: perfGroupCPU,
	Name:  "usage",
	Type:  perfTypeAverage,
}

var PerfCPUUsageMhzAvg = PerfMetric{
	Group: perfGroupCPU,
	Name:  "usagemhz",
	Type:  perfTypeAverage,
}

var PerfCPUReadySum = PerfMetric{
	Group: perfGroupCPU,
	Name:  "ready",
	Type:  perfTypeSum,
}

var PerfMemoryActiveAvg = PerfMetric{
	Group: perfGroupMemory,
	Name:  "active",
	Type:  perfTypeAverage,
}

var PerfMemoryConsumedAvg = PerfMetric{
	Group: perfGroupMemory,
	Name:  "consumed",
	Type:  perfTypeAverage,
}

var PerfMemoryBalloonedAvg = PerfMetric{
	Group: perfGroupMemory,
	Name:  "vmmemctrl",
	Type:  perfTypeAverage,
}

var PerfMemorySwappedAvg = PerfMetric{
	Group: perfGroupMemory,
	Name:  "swapped",
	Type:  perfTypeAverage,
}

var PerfNetworkUsageAvg = PerfMetric{
	Group: perfGroupNetwork,
	Name:  "usage",
	Type:  perfTypeAverage,
}

var PerfNetworkReceivedAvg = PerfMetric{
	Group: perfGroupNetwork,
	Name:  "received",
	Type:  perfTypeAverage,
}

var PerfNetworkTransmittedAvg = PerfMetric{
	Group: perfGroupNetwork,
	Name:  "transmitted",
	Type:  perfTypeAverage,
}

var PerfDiskReadAvg = PerfMetric{
	Group: perfGroupDisk,
	Name:  "read",
	Type:  perfTypeAverage,
}

var PerfDiskWriteAvg = PerfMetric{
	Group: perfGroupDisk,
	Name:  "write",
	Type:  perfTypeAverage,
}

var PerfDiskMaxLatency = PerfMetric{
	Group: perfGroupDisk,
	Name:  "maxtotallatency",
	Type:  perfTypeLatest,
}

var PerfDiskUsageAvg = PerfMetric{
	Group: perfGroupDisk,
	Name:  "usage",
	Type:  perfTypeAverage,
}

var PerfDiskProvisioned = PerfMetric{
	Group: perfGroupDisk,
	Name:  "provisioned",
	Type:  perfTypeLatest,
}

var PerfDiskUsed = PerfMetric{
	Group: perfGroupDisk,
	Name:  "used",
	Type:  perfTypeLatest,
}

var PerfDiskWriteIOPSAvg = PerfMetric{
	Group: perfGroupDisk,
	Name:  "numberwriteaverage",
	Type:  perfTypeAverage,
}

var PerfDiskReadIOPSAvg = PerfMetric{
	Group: perfGroupDisk,
	Name:  "numberreadaverage",
	Type:  perfTypeAverage,
}
