package theclock

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type Clock struct {
	HoursHand   HoursHand
	MinutesHand MinutesHand
	SecondsHand SecondsHand
}

type Point struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

type HoursHand Point
type MinutesHand Point
type SecondsHand Point

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

func BuildClock(w io.Writer, tm time.Time) Clock {
	fmt.Fprint(w, svgStart)
	fmt.Fprint(w, svgBezel)
	fmt.Fprintf(w, svgPoints)
	pt := GetPoint(tm)
	fmt.Fprintf(w, `<line x1="%f" y1="%f" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:7px;"/>`, pt.SecondsHand.X1, pt.SecondsHand.Y1, pt.SecondsHand.X2, pt.SecondsHand.Y2)
	fmt.Fprintf(w, `<line x1="%f" y1="%f" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, pt.MinutesHand.X1, pt.MinutesHand.Y1, pt.MinutesHand.X2, pt.MinutesHand.Y2)
	fmt.Fprintf(w, `<line x1="%f" y1="%f" x2="%f" y2="%f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, pt.HoursHand.X1, pt.HoursHand.Y1, pt.HoursHand.X2, pt.HoursHand.Y2)
	fmt.Fprint(w, svgEnd)
	return pt
}

func InRadians(currentValue float64, totalDivisions int) float64 {
	return math.Pi / ((float64(totalDivisions) / 2) / currentValue)
}

func GetUnitPoint(tm time.Time) Clock {
	secAngle := InRadians(float64(tm.Second()), 60)
	secHand := SecondsHand{0, 0, math.Sin(secAngle), -math.Cos(secAngle)}
	minAngle := InRadians(float64(tm.Minute()), 60) + secAngle/60
	minHand := MinutesHand{0, 0, math.Sin(minAngle), -math.Cos(minAngle)}
	hourAngle := InRadians(float64(tm.Hour()), 12) + minAngle/60
	hourHand := HoursHand{0, 0, math.Sin(hourAngle), -math.Cos(hourAngle)}
	return Clock{
		HoursHand:   hourHand,
		MinutesHand: minHand,
		SecondsHand: secHand,
	}
}

const (
	StartPosX     = 150
	StartPosY     = 150
	MinuteHandLen = 90
	HourHandLen   = 50
	SecondHandLen = 90
)

func GetPoint(tm time.Time) Clock {
	pt := GetUnitPoint(tm)
	pt.HoursHand.X1 += StartPosX
	pt.HoursHand.Y1 += StartPosY
	pt.MinutesHand.X1 += StartPosX
	pt.MinutesHand.Y1 += StartPosY
	pt.SecondsHand.X1 += StartPosX
	pt.SecondsHand.Y1 += StartPosY

	pt.HoursHand.X2 = pt.HoursHand.X2*HourHandLen + StartPosX
	pt.HoursHand.Y2 = pt.HoursHand.Y2*HourHandLen + StartPosY
	pt.MinutesHand.X2 = pt.MinutesHand.X2*MinuteHandLen + StartPosX
	pt.MinutesHand.Y2 = pt.MinutesHand.Y2*MinuteHandLen + StartPosY
	pt.SecondsHand.X2 = pt.SecondsHand.X2*SecondHandLen + StartPosX
	pt.SecondsHand.Y2 = pt.SecondsHand.Y2*SecondHandLen + StartPosY

	return pt
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const svgBezel = `<circle cx="150" cy="150" r="110" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgPoints = `<circle cx="150" cy="45" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="45" cy="150" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="98" cy="59" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="98" cy="241" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="200" cy="58" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="200" cy="242" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="59" cy="98" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="59" cy="202" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="241" cy="98" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="241" cy="202" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="150" cy="255" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>
<circle cx="255" cy="150" r="2" style="fill:#fff;stroke:#000;stroke-width:1px;"/>`
const svgEnd = `</svg>`
