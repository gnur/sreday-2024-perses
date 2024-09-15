package main

import (
	"flag"

	"github.com/perses/perses/go-sdk"
	"github.com/perses/perses/go-sdk/dashboard"
	"github.com/perses/perses/go-sdk/panel"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	"github.com/perses/perses/go-sdk/prometheus/query"

	timeSeriesPanel "github.com/perses/perses/go-sdk/panel/time-series"
	labelValuesVar "github.com/perses/perses/go-sdk/prometheus/variable/label-values"
	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
)

func main() {
	flag.Parse()
	exec := sdk.NewExec()
	builder, buildErr := dashboard.New("GoSDKexample",
		dashboard.ProjectName("sreday2024"),

		dashboard.AddVariable("stack",
			listVar.List(
				labelValuesVar.PrometheusLabelValues("instance",
					labelValuesVar.Matchers(`{job="blackbox"}`),
				),
				listVar.DisplayName("stack?"),
			),
		),

		dashboard.AddPanelGroup("Certificate expiry",
			panelgroup.PanelsPerLine(3),
			panelgroup.AddPanel("date of expiry",
				timeSeriesPanel.Chart(),
				panel.AddQuery(
					query.PromQL("probe_ssl_earliest_cert_expiry{instance=\"$stack\"}"),
				),
			),
		),
	)
	exec.BuildDashboard(builder, buildErr)
}
