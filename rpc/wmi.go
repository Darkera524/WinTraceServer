package rpc

import (
"github.com/Darkera524/WinTraceTool/model"
"github.com/Darkera524/WinTraceServer/g"
)

type Wmi int

func (t *Wmi) Infos(args model.RequestModel, reply *model.WmiResp) error {
	//providers := make(map[string][]string)
	//metricMap := make(map[string]string)
	hostname := args.Hostname
	wmilist := g.GetConfig().Wmi
	for _,v := range wmilist {
		hostlist := v.Hostname
		for _,inshost := range hostlist {
			if hostname == inshost {
				//providers[v.Guid] = v.Tags
				reply.WmiList = v.WmiList
				break
			}
		}
	}
	//reply.Providers = providers

	return nil
}
