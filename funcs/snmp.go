// Copyright 2018 Steven Lee <geekerlw.gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"log"

	"github.com/geekerlw/falcon-agent/g"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/soniah/gosnmp"
)

func snmpGet(addr string, oids []string) (map[string]interface{}, error) {
	gosnmp.Default.Target = addr
	ret := make(map[string]interface{})

	errc := gosnmp.Default.Connect()
	if errc != nil {
		return nil, errc
	}
	defer gosnmp.Default.Conn.Close()

	res, errg := gosnmp.Default.Get(oids)
	if errg != nil {
		return nil, errg
	}

	for _, v := range res.Variables {

		switch v.Type {
		case gosnmp.OctetString:
			ret[v.Name] = string(v.Value.([]byte))
		default:
			ret[v.Name] = gosnmp.ToBigInt(v.Value)
		}
	}

	return ret, nil
}

func SnmpMetrics() (L []*model.MetricValue) {
	reportOids := g.ReportOids()
	sz := len(reportOids)
	if sz == 0 {
		return
	}

	for tags, m := range reportOids {
		var addr string
		var oids []string

		for key, val := range m {
			if key == 0 {
				addr = val
			} else if key == 1 {
				oids = append(oids, val)
			}
		}

		res, err := snmpGet(addr, oids)
		if err != nil {
			log.Printf("failed to get oid, err: %v\n", err)
			continue
		}

		for _, v := range res {
			L = append(L, GaugeValue("snmp.get", v, tags))
		}
	}

	return
}
