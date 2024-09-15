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
				labelValuesVar.PrometheusLabelValues("paas",
					labelValuesVar.Matchers(`{job="blackbox"}`),
				),
				listVar.DisplayName("stack?"),
			),
		),

		dashboard.AddPanelGroup("Resource usage",
			panelgroup.PanelsPerLine(3),
			panelgroup.AddPanel("Container memory",
				timeSeriesPanel.Chart(),
				panel.AddQuery(
					query.PromQL("prob_ssl_earliest_cert_expiry{instance=\"$stack\"}"),
				),
			),
		),
	)
	exec.BuildDashboard(builder, buildErr)
}
