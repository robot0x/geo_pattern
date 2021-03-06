// Package svg provide methods to work effortlessly with SVG.
package svg

import (
	"fmt"
	"reflect"
)

// SVG struct, SVG contains elementry attributes like
// svg string, width, height
type SVG struct {
	svg_string    string
	width, height int
}

// Set_width sets SVG object's width
func (s *SVG) Set_width(w int) {
	s.width = w
}

// Set_height sets SVG object's height
func (s *SVG) Set_height(h int) {
	s.height = h
}

// header returns string representing SVG object's header(staring) part
func (s *SVG) header() string {
	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' width='%v' height='%v'>", s.width, s.height)
}

// footer returns string representing SVG object's footer(ending) part
func (s *SVG) footer() string {
	return "</svg>"
}

// Str returns string representing whole SVG object
func (s *SVG) Str() string {
	return s.header() + s.svg_string + s.footer()
}

// Rect adds a rectangle element to SVG object
func (s *SVG) Rect(x, y, w, h interface{}, args map[string]interface{}) {
	rect_str := fmt.Sprintf("<rect x='%v' y='%v' width='%v' height='%v' %s />", x, y, w, h, s.Write_args(args))
	s.svg_string += rect_str
}

// Circle adds a circle element to SVG object
func (s *SVG) Circle(cx, cy, r interface{}, args map[string]interface{}) {
	circle_str := fmt.Sprintf("<circle cx='%v' cy='%v' r='%v' %s />", cx, cy, r, s.Write_args(args))
	s.svg_string += circle_str
}

// Path adds a path element to SVG object
func (s *SVG) Path(str string, args map[string]interface{}) {
	path_str := fmt.Sprintf("<path d='%s' %s />", str, s.Write_args(args))
	s.svg_string += path_str
}

// Polyline adds a polyline element to SVG object
func (s *SVG) Polyline(str string, args map[string]interface{}) {
	polyline_str := fmt.Sprintf("<polyline points='%s' %s />", str, s.Write_args(args))
	s.svg_string += polyline_str
}

// Group adds a group element to SVG object.
//
// It groups optionally provided elements together.
func (s *SVG) Group(elements [2]string, args map[string]interface{}) {
	s.svg_string += fmt.Sprintf("<g %s>", s.Write_args(args))
	s.svg_string += elements[0] + elements[1]
	s.svg_string += "</g>"
}

// Write_args adds additional attributes to a SVG elements.
//
// It parses provides 'map' arguments to add attributes to SVG element.
func (s *SVG) Write_args(args map[string]interface{}) string {
	str := ""

	for k, v := range args {
		obj_type := fmt.Sprintf("%s", reflect.TypeOf(v))

		switch obj_type {
		case "string":
			str += fmt.Sprintf("%s='%s' ", k, v)
		case "int":
			str += fmt.Sprintf("%s='%v' ", k, v)
		case "float64":
			str += fmt.Sprintf("%s='%v' ", k, v)
		default:
			{
				str += fmt.Sprintf("%s='", k)
				for K, V := range v.(map[string]string) {
					str += fmt.Sprintf("%s:%s;", K, V)
				}
				str += "' "
			}
		}
	}

	return str
}
