package rpc

import (
	"github.com/Darkera524/WinTraceTool/model"
	"github.com/Darkera524/WinTraceServer/g"
)

type Trace int

func (t *Trace) Providers(args model.RequestModel, reply *model.ProvidersResp) error {
	providers := make(map[string][]string)
	metricMap := make(map[string]string)
	hostname := args.Hostname
	tracelist := g.GetConfig().Trace
	for _,v := range tracelist {
		hostlist := v.Hostname
		for _,inshost := range hostlist {
			if hostname == inshost {
				providers[v.Guid] = v.Tags
				metricMap[v.Guid] = v.Name
			}
		}
	}
	reply.Providers = providers
	reply.MetricMap = metricMap

	return nil
}
