package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type chart struct {
	Source        string              `json:"source"`
	LabelInterval string              `json:"labelInterval"`
	DecimalPlaces string              `json:"decimalPlaces"`
	Unit          string              `json:"unit"`
	XAxisLabel    string              `json:"xAxisLabel"`
	AspectRatio   string              `json:"aspectRatio"`
	ChartType     string              `json:"chartType"`
	Data          []map[string]string `json:"data"`
}

type highchartsOptions struct {
	Viewport           string                              `json:"viewport,omitempty"`
	Chart              highchartsOptionsChart              `json:"chart,omitempty"`
	Colors             []string                            `json:"colors,omitempty"`
	Series             []highchartsOptionsSeries           `json:"series,omitempty"` // TODO
	Title              highchartsOptionsTitle              `json:"title,omitempty"`
	Subtitle           highchartsOptionsTitle              `json:"subtitle,omitempty"`
	YAxis              highchartsOptionsYAxis              `json:"yAxis,omitempty"`
	XAxis              highchartsOptionsXAxis              `json:"xAxis,omitempty"`
	Legend             highchartsOptionsLegend             `json:"legend,omitempty"`
	PlotOptions        highchartsOptionsPlotOptions        `json:"plotOptions,omitempty"`
	AnnotationsOptions highchartsOptionsAnnotationsOptions `json:"annotationsOptions,omitempty"`
	Annotations        []interface{}                       `json:"annotations,omitempty"`
	AspectRatio        float64                             `json:"aspectRatio"`
	Tooltip            highchartsOptionsTooltip            `json:"tooltip"`
	Credits            highchartsOptionsCredits            `json:"credits"`
	Labels             highchartsOptionsLabels             `json:"labels"`
	Exporting          highchartsOptionsExporting          `json:"exporting"`
}

type highchartsOptionsSeries struct {
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}

type highchartsOptionsTitle struct {
	Text string `json:"text,omitempty"`
	Y    int    `json:"y,omitempty"`
}

type highchartsOptionsYAxis struct {
	PlotBands  []interface{}               `json:"plotBands,omitempty"` // TODO
	PlotLines  []interface{}               `json:"plotLines,omitempty"` // TODO
	Title      highchartsOptionsAxisTitle  `json:"title,omitempty"`
	Labels     highchartsOptionsAxisLabels `json:"labels,omitempty"`
	GridZIndex int                         `json:"gridZIndex,omitempty"`
	Opposite   bool                        `json:"opposite"`
}

type highchartsOptionsAxisTitle struct {
	Text     string                 `json:"text,omitempty"`
	UseHTML  bool                   `json:"useHTML"`
	Align    string                 `json:"align,omitempty"`
	Offset   int                    `json:"offset"`
	Y        int                    `json:"y"`
	Rotation int                    `json:"rotation"`
	Style    map[string]interface{} `json:"style"`
}

type highchartsOptionsAxisLabels struct {
	Format  string `json:"format,omitempty"`
	UseHTML bool   `json:"useHTML"`
}

type highchartsOptionsXAxis struct {
	PlotBands         []interface{}               `json:"plotBands,omitempty"` // TODO
	PlotLines         []interface{}               `json:"plotLines,omitempty"` // TODO
	Categories        []string                    `json:"categories,omitempty"`
	TickInterval      int                         `json:"tickInterval,omitempty"`
	Title             highchartsOptionsAxisTitle  `json:"title,omitempty"`
	Labels            highchartsOptionsAxisLabels `json:"labels,omitempty"`
	Opposite          bool                        `json:"opposite"`
	TickmarkPlacement string                      `json:"tickmarkPlacement"`
}

type highchartsOptionsChart struct {
	Height       int                         `json:"height,omitempty"`
	Width        int                         `json:"width,omitempty"`
	MarginRight  int                         `json:"marginRight,omitempty"`
	MarginTop    int                         `json:"marginTop,omitempty"`
	MarginBottom int                         `json:"marginBottom,omitempty"`
	SpacingTop   int                         `json:"spacingTop,omitempty"`
	Offset       int                         `json:"offset,omitempty"`
	Type         string                      `json:"type"`
	Style        highchartsOptionsChartStyle `json:"style,omitempty"`
}

type highchartsOptionsLegend struct {
	VerticalAlign    string                            `json:"verticalAlign,omitempty"`
	Y                int                               `json:"y,omitempty"`
	X                int                               `json:"x,omitempty"`
	UseHTML          bool                              `json:"useHTML"`
	Enabled          bool                              `json:"enabled"`
	Align            string                            `json:"align,omitempty"`
	ItemMarginBottom int                               `json:"itemMarginBottom,omitempty"`
	ItemStyle        highchartsOptionsLegendItemStyle  `json:"itemStyle,omitempty"`
	Navigation       highchartsOptionsLegendNavigation `json:"navigation,omitempty"`
}

type highchartsOptionsChartStyle struct {
	FontFamily string `json:"fontFamily,omitempty"`
}

type highchartsOptionsLegendItemStyle struct {
	FontWeight string `json:"fontWeight,omitempty"`
	Color      string `json:"color,omitempty"`
}

type highchartsOptionsLegendNavigation struct {
	Enabled bool `json:"enabled"`
}

type highchartsOptionsPlotOptions struct {
	Series highchartsOptionsPlotOptionsSeries `json:"series,omitempty"`
	Line   highchartsOptionsPlotOptionsLine   `json:"line,omitempty"`
	Area   highchartsOptionsPlotOptionsArea   `json:"area,omitempty"`
}

type highchartsOptionsPlotOptionsSeries struct {
	BorderWidth  int                                      `json:"borderWidth"`
	Animation    bool                                     `json:"animation"`
	PointPadding int                                      `json:"pointPadding"`
	GroupPadding float64                                  `json:"groupPadding"`
	States       highchartsOptionsPlotOptionsSeriesStates `json:"states"`
}

type highchartsOptionsPlotOptionsSeriesStates struct {
	Hover highchartsOptionsPlotOptionsSeriesStatesHover `json:"hover,omitempty"`
}

type highchartsOptionsPlotOptionsSeriesStatesHover struct {
	Enabled bool `json:"enabled"`
}

type highchartsOptionsPlotOptionsLine struct {
	LineWidth int                                    `json:"lineWidth"`
	Marker    highchartsOptionsPlotOptionsLineMarker `json:"marker"`
}

type highchartsOptionsPlotOptionsLineMarker struct {
	Enabled bool   `json:"enabled"`
	Radius  int    `json:"radius"`
	Symbol  string `json:"symbol"`
}

type highchartsOptionsPlotOptionsArea struct {
	Stacking string `json:"stacking"`
}

type highchartsOptionsAnnotationsOptions struct {
	EnabledButtons bool `json:"enabledButtons"`
}

type highchartsOptionsTooltip struct {
	ValueDecimals   int                           `json:"valueDecimals"`
	Shared          bool                          `json:"shared"`
	UseHTML         bool                          `json:"useHTML"`
	Style           highchartsOptionsTooltipStyle `json:"style"`
	BackgroundColor string                        `json:"backgroundColor"`
	BorderWidth     int                           `json:"borderWidth"`
	Padding         int                           `json:"padding"`
	BorderRadius    int                           `json:"borderRadius"`
	BorderColor     string                        `json:"borderColor"`
	Shadow          bool                          `json:"shadow"`
}

type highchartsOptionsTooltipStyle struct {
	Padding int `json:"padding"`
}

type highchartsOptionsCredits struct {
	Enabled bool `json:"enabled"`
}

type highchartsOptionsLabels struct {
	UseHTML bool `json:"useHTML"`
}

type highchartsOptionsExporting struct {
	Enabled bool `json:"enabled"`
}

var defaultChart = highchartsOptions{
	Viewport: "lg",
	Chart: highchartsOptionsChart{
		Height:       392,
		Width:        700,
		MarginRight:  35,
		MarginTop:    150,
		MarginBottom: 150,
		SpacingTop:   16,
		Offset:       150,
		Style: highchartsOptionsChartStyle{
			FontFamily: `"Open Sans", Tahoma, Verdana, Arial`,
		},
	},
	Colors: []string{
		"rgba(39,71,150,0.9)",
		"rgba(245,148,47,0.9)",
		"rgba(231,63,64,0.9)",
		"rgba(123,202,226,0.9)",
		"rgba(151,151,150,0.9)",
		"rgba(233,225,23,0.9)",
		"rgba(116,182,48,0.9)",
		"rgba(103,71,150,0.9)",
		"rgba(189,91,158,0.9)",
	},
	Series: []highchartsOptionsSeries{},
	Title: highchartsOptionsTitle{
		Text: "Figure 5: UK goods exports to the EU and non-EU areas, percentage of total UK goods exports, current prices, 1999 to 2015",
		Y:    30,
	},
	Subtitle: highchartsOptionsTitle{
		Text: "",
		Y:    80,
	},
	YAxis: highchartsOptionsYAxis{
		PlotBands: nil,
		PlotLines: nil,
		Title: highchartsOptionsAxisTitle{
			Text:    "%",
			UseHTML: true,
			Align:   "high",
		},
		Labels: highchartsOptionsAxisLabels{
			Format:  "{value:,.f}",
			UseHTML: true,
		},
		GridZIndex: -200,
		Opposite:   false,
	},
	XAxis: highchartsOptionsXAxis{
		PlotBands: nil,
		PlotLines: nil,
		Categories: []string{
			"1999", "2000", "2001", "2002", "2003",
			"2004", "2005", "2006", "2007", "2008",
			"2009", "2010", "2011", "2012", "2013",
			"2014", "2015",
		},
		TickInterval: 2,
		Labels: highchartsOptionsAxisLabels{
			UseHTML: true,
		},
		Title: highchartsOptionsAxisTitle{
			UseHTML: true,
			Text:    "",
		},
		Opposite: false,
	},
	Legend: highchartsOptionsLegend{
		VerticalAlign:    "top",
		Y:                582,
		X:                30,
		UseHTML:          true,
		Enabled:          true,
		Align:            "left",
		ItemMarginBottom: 8,
		ItemStyle: highchartsOptionsLegendItemStyle{
			FontWeight: "normal",
			Color:      "rgb(112,112,112)",
		},
		Navigation: highchartsOptionsLegendNavigation{
			Enabled: false,
		},
	},
	PlotOptions: highchartsOptionsPlotOptions{
		Series: highchartsOptionsPlotOptionsSeries{
			BorderWidth:  0,
			Animation:    false,
			PointPadding: 0,
			GroupPadding: 0.1,
			States: highchartsOptionsPlotOptionsSeriesStates{
				Hover: highchartsOptionsPlotOptionsSeriesStatesHover{
					Enabled: false,
				},
			},
		},
		Line: highchartsOptionsPlotOptionsLine{
			LineWidth: 2,
			Marker: highchartsOptionsPlotOptionsLineMarker{
				Enabled: false,
				Radius:  2,
				Symbol:  "circle",
			},
		},
		Area: highchartsOptionsPlotOptionsArea{
			Stacking: "normal",
		},
	},
	AnnotationsOptions: highchartsOptionsAnnotationsOptions{
		EnabledButtons: false,
	},
	Annotations: nil,
	AspectRatio: 0.56,
	Tooltip: highchartsOptionsTooltip{
		ValueDecimals: 1,
		Shared:        true,
		UseHTML:       true,
		Style: highchartsOptionsTooltipStyle{
			Padding: 0,
		},
		BackgroundColor: "rgba(208,210,211,1)",
		BorderWidth:     0,
		Padding:         0,
		BorderRadius:    0,
		BorderColor:     "rgba(255, 255, 255, 0)",
		Shadow:          false,
	},
	Credits: highchartsOptionsCredits{
		Enabled: false,
	},
	Labels: highchartsOptionsLabels{
		UseHTML: true,
	},
	Exporting: highchartsOptionsExporting{
		Enabled: false,
	},
}

func main() {
	b, err := ioutil.ReadFile("7aded033.json")
	if err != nil {
		panic(err)
	}

	var c chart
	if err = json.Unmarshal(b, &c); err != nil {
		panic(err)
	}

	//log.Println(c)

	var o = func() highchartsOptions { return defaultChart }()
	o.YAxis.Title.Align = "high"
	o.YAxis.Title.Offset = 0
	o.YAxis.Title.Y = -4
	o.YAxis.Title.Rotation = 360
	o.YAxis.Title.Style = map[string]interface{}{
		"left":  "0px",
		"right": "0px",
	}
	o.Series = []highchartsOptionsSeries{
		{
			Name: "EU",
			Data: []float64{61.2, 60.1, 60.7, 61.9, 59.6, 58.9, 58.0, 62.2, 57.6, 56.1, 54.8, 53.7, 53.6, 50.0, 49.9, 49.7, 46.9},
		},
		{
			Name: "Non-EU",
			Data: []float64{38.8, 39.9, 39.3, 38.1, 40.4, 41.1, 42.0, 37.8, 42.4, 43.9, 45.2, 46.3, 46.4, 50.0, 50.1, 50.3, 53.1},
		},
	}
	o.Chart.Type = "line"
	//o.XAxis.TickmarkPlacement = "on"

	ob, err := json.Marshal(&o)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(ob))
}
