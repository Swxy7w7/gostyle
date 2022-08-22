package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
	"unicode"
)

var (
	black_to_white  = []string{"m;m;m"}
	black_to_red    = []string{"m;0;0"}
	black_to_green  = []string{"0;m;0"}
	black_to_blue   = []string{"0;0;m"}
	white_to_black  = []string{"n;n;n"}
	white_to_red    = []string{"255;n;n"}
	white_to_green  = []string{"n;255;n"}
	white_to_blue   = []string{"n;n;255"}
	red_to_black    = []string{"n;0;0"}
	red_to_white    = []string{"255;m;m"}
	red_to_yellow   = []string{"255;m;0"}
	red_to_purple   = []string{"255;0;m"}
	green_to_black  = []string{"0;n;0"}
	green_to_white  = []string{"m;255;m"}
	green_to_yellow = []string{"m;255;0"}
	green_to_cyan   = []string{"0;255;m"}
	blue_to_black   = []string{"0;0;n"}
	blue_to_white   = []string{"m;m;255"}
	blue_to_cyan    = []string{"0;m;255"}
	blue_to_purple  = []string{"m;0;255"}
	yellow_to_red   = []string{"255;n;0"}
	yellow_to_green = []string{"n;255;0"}
	purple_to_red   = []string{"255;0;n"}
	purple_to_blue  = []string{"n;0;255"}
	cyan_to_green   = []string{"0;255;n"}
	cyan_to_blue    = []string{"0;n;255"}
	red             = _start("255;0;0")
	green           = _start("0;255;0")
	blue            = _start("0;0;255")
	white           = _start("255;255;255")
	black           = _start("0;0;0")
	gray            = _start("150;150;150")
	yellow          = _start("255;255;0")
	purple          = _start("255;0;255")
	cyan            = _start("0;255;255")
	orange          = _start("255;150;0")
	pink            = _start("255;0;150")
	turquoise       = _start("0;150;255")
	light_gray      = _start("200;200;200")
	dark_gray       = _start("100;100;100")
	light_red       = _start("255;100;100")
	light_green     = _start("100;255;100")
	light_blue      = _start("100;100;255")
	dark_red        = _start("100;0;0")
	dark_green      = _start("0;100;0")
	dark_blue       = _start("0;0;100")
	reset           = white
)

// _MakeColords

func _makeansi(col string, text string) string {
	return func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("\u001b[38;2;{{.col}}m{{.text}}\u001b[38;2;255;255;255m")).
			Execute(&buf, map[string]interface{}{"col": col, "text": text})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}()
}

func _rmansi(col string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(strings.ReplaceAll(col, "\u001b[38;2;", ""), "m", ""),
			"50m",
			"",
		),
		"\u001b[38",
		"",
	)
}

func _start(color string) string {
	return func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("\u001b[38;2;{{.color}}m")).
			Execute(&buf, map[string]interface{}{"color": color})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}()
}

func _end() string {
	return "\u001b[38;2;255;255;255m"
}

func _maketext(color string, text string, end string) string {
	if len(end) != 0 {
		end = _end()
	} else {
		end = ""
	}
	return color + text + string(end)
}

func _getspaces(text string) int {
	return len(text) - len(strings.TrimLeftFunc(text, unicode.IsSpace))
}

// Colors

func StaticRGB(r int, g int, b int) string {
	return _start(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("{{.r}};{{.g}};{{.b}}")).
			Execute(&buf, map[string]interface{}{"r": r, "g": g, "b": b})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}())
}

func StaticMIX(colors []interface{}, start bool) string {
	rgb := [][]int{}
	for _, col := range colors {
		col := _rmansi(col)
		col = strings.Split(col, ";")
		r := int(col[0])
		g := int(col[1])
		b := int(col[2])
		rgb = append(rgb, []int{r, g, b})
	}
	r_list := [][]int{}
	for _, rgb := range rgb {
		r_list = append(r_list, rgb[0])
	}
	r := round(func() (s []int) {
		for _, e := range r_list {
			s = append(s, e...)
		}
		return
	}() / len(r_list))
	g_list := [][]int{}
	for _, rgb := range rgb {
		g_list = append(g_list, rgb[1])
	}
	g := round(func() (s []int) {
		for _, e := range g_list {
			s = append(s, e...)
		}
		return
	}() / len(g))
	b_list := [][]int{}
	for _, rgb := range rgb {
		b_list = append(b_list, rgb[2])
	}
	b := round(func() (s []int) {
		for _, e := range b_list {
			s = append(s, e...)
		}
		return
	}() / len(b_list))
	rgb = func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("{{.r}};{{.g}};{{.b}}")).
			Execute(&buf, map[string]interface{}{"r": r, "g": g, "b": b})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}()
	return func() interface{} {
		if start {
			return _start(rgb)
		}
		return rgb
	}()
}

func Symbol(symbol string, col string, col_left_right string, left string, right string) string {

	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		col_left_right, left, symbol, right, col_left_right, col, col_left_right,
	)
}

var col = [2]interface{}{list, str}

var dynamic_colors = []interface{}{
	black_to_white,
	black_to_red,
	black_to_green,
	black_to_blue,
	white_to_black,
	white_to_red,
	white_to_green,
	white_to_blue,
	red_to_black,
	red_to_white,
	red_to_yellow,
	red_to_purple,
	green_to_black,
	green_to_white,
	green_to_yellow,
	green_to_cyan,
	blue_to_black,
	blue_to_white,
	blue_to_cyan,
	blue_to_purple,
	yellow_to_red,
	yellow_to_green,
	purple_to_red,
	purple_to_blue,
	cyan_to_green,
	cyan_to_blue,
}

var red_to_blue interface{}
var red_to_green interface{}
var green_to_blue interface{}
var green_to_red interface{}
var blue_to_red interface{}
var blue_to_green interface{}
var rainbow interface{}
var static_colors []interface{}
var all_colors []interface{}

func init() {
	for _, color := range dynamic_colors {
		_col := 20
		reversed_col := 220
		dbl_col := 20
		dbl_reversed_col := 220
		content := color[0]
		func(s *interface{}, i int) interface{} {
			popped := (*s)[i]
			*s = append((*s)[:i], (*s)[i+1:]...)
			return popped
		}(&color, 0)
		for _ := 0; _ < 12; _++ {
			if func() int {
				for i, v := range content {
					if v == "m" {
						return i
					}
				}
				return -1
			}() != -1 {
				result := strings.ReplaceAll(content, "m", fmt.Sprintf("%v", _col))
				color = append(color, result)
			} else if func() int {
				for i, v := range content {
					if v == "n" {
						return i
					}
				}
				return -1
			}() != -1 {
				result := strings.ReplaceAll(content, "n", fmt.Sprintf("%v", reversed_col))
				color = append(color, result)
			}
			_col += 20
			reversed_col -= 20
		}
		for _ := 0; _ < 12; _++ {
			if func() int {
				for i, v := range content {
					if v == "m" {
						return i
					}
				}
				return -1
			}() != -1 {
				result := strings.ReplaceAll(content, "m", fmt.Sprintf("%v", dbl_reversed_col))
				color = append(color, result)
			} else if func() int {
				for i, v := range content {
					if v == "n" {
						return i
					}
				}
				return -1
			}() != -1 {
				result := strings.ReplaceAll(content, "n", fmt.Sprintf("%v", dbl_col))
				color = append(color, result)
			}
			dbl_col += 20
			dbl_reversed_col -= 20
		}
	}
	red_to_blue = _makergbcol(red_to_purple, purple_to_blue)
	red_to_green = _makergbcol(red_to_yellow, yellow_to_green)
	green_to_blue = _makergbcol(green_to_cyan, cyan_to_blue)
	green_to_red = _makergbcol(green_to_yellow, yellow_to_red)
	blue_to_red = _makergbcol(blue_to_purple, purple_to_red)
	blue_to_green = _makergbcol(blue_to_cyan, cyan_to_green)
	rainbow = _makerainbow(red_to_green, green_to_blue, blue_to_red)
	for _, _col := range [6]interface{}{red_to_blue, red_to_green, green_to_blue, green_to_red, blue_to_red, blue_to_green} {
		dynamic_colors = append(dynamic_colors, _col)
	}
	dynamic_colors = append(dynamic_colors, rainbow)
	static_colors = []interface{}{
		red,
		green,
		blue,
		white,
		black,
		gray,
		yellow,
		purple,
		cyan,
		orange,
		pink,
		turquoise,
		light_gray,
		dark_gray,
		light_red,
		light_green,
		light_blue,
		dark_red,
		dark_green,
		dark_blue,
		reset,
	}
	all_colors = func() (elts []interface{}) {
		for _, color := range dynamic_colors {
			elts = append(elts, color)
		}
		return
	}()
	for _, color := range static_colors {
		all_colors = append(all_colors, color)
	}
}

// Colorate

func Color(color string, text string, end string) string {
	return _maketext(color, text, end)
}

func Error(text string, color string, end bool, spaces bool, enter bool, wait int) string {
	var _var string
	content := _maketext(color, strings.Repeat("\n", func() int {
		if spaces {
			return 1
		}
		return 0
	}())+text, end)
	if enter {
		_var = func(msg string) string {
			fmt.Print(msg)
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			return strings.ReplaceAll(text, "\n", "")
		}(content)
	} else {
		fmt.Println(content)
		_var = nil
	}
	if &wait == (*int)(&true) {
		os.Exit(0)
	} else if &wait != (*int)(&false) {
		_sleep(wait)
	}
	return _var
}
func Vertical(color []interface {
}, text string, speed int, start int, stop int, cut int, fill bool) string {
	color := color[cut:]
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(text)
	result := ""
	nstart := 0
	color_n := 0
	for _, lin := range lines {
		colorR := string(color[color_n])
		if fill {
			result += strings.Repeat(" ", int(_getspaces(lin))) + strings.Join(func() func() <-chan interface {
			} {
				wait := make(chan struct {
				})
				yield := make(chan interface {
				})
				go func() {
					defer close(yield)
					<-wait
					for _, x := range strings.TrimSpace(lin) {
						yield <- _makeansi(colorR, x)
						<-wait
					}
				}()
				return func() <-chan interface {
				} {
					wait <- struct {
					}{}
					return yield
				}
			}(), "") + "\n"
		} else {
			result += strings.Repeat(" ", int(_getspaces(lin))) + _makeansi(colorR, strings.TrimSpace(lin)) + "\n"
		}
		if nstart != start {
			nstart += 1
			continue
		}
		if len(strings.TrimRightFunc(lin, unicode.IsSpace)) != 0 {
			if stop == 0 && color_n+speed < len(color) || stop != 0 && color_n+speed < stop {
				color_n += speed
			} else if stop == 0 {
				color_n = 0
			} else {
				color_n = stop
			}
		}
	}
	return strings.TrimRightFunc(result, unicode.IsSpace)
}
func Horizontal(color []interface {
}, text string, speed int, cut int) string {
	color := color[cut:]
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(text)
	result := ""
	for _, lin := range lines {
		carac := func() {
			elts := []interface {
			}{}
			for _, elt := range lin {
				elts = append(elts, elt)
			}
			elts
		}()
		color_n := 0
		for _, car := range carac {
			colorR := string(color[color_n])
			result += strings.Repeat(" ", int(_getspaces(car))) + _makeansi(colorR, strings.TrimSpace(car))
			if color_n+speed < len(color) {
				color_n += speed
			} else {
				color_n = 0
			}
		}
		result += "\n"
	}
	return strings.TrimRightFunc(result, unicode.IsSpace)
}
func Diagonal(color []interface {
}, text string, speed int, cut int) string {
	color := color[cut:]
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(text)
	result := ""
	color_n := 0
	for _, lin := range lines {
		carac := func() {
			elts := []interface {
			}{}
			for _, elt := range lin {
				elts = append(elts, elt)
			}
			elts
		}()
		for _, car := range carac {
			colorR := string(color[color_n])
			result += strings.Repeat(" ", int(_getspaces(car))) + _makeansi(colorR, strings.TrimSpace(car))
			if color_n+speed < len(color) {
				color_n += speed
			} else {
				color_n = 1
			}
		}
		result += "\n"
	}
	return strings.TrimRightFunc(result, unicode.IsSpace)
}
func DiagonalBackwards(color []interface {
}, text string, speed int, cut int) string {
	color := color[cut:]
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(text)
	result := ""
	resultL := ""
	color_n := 0
	for _, lin := range lines {
		carac := func() {
			elts := []interface {
			}{}
			for _, elt := range lin {
				elts = append(elts, elt)
			}
			elts
		}()
		carac.reverse()
		resultL = ""
		for _, car := range carac {
			colorR := string(color[color_n])
			resultL = strings.Repeat(" ", int(_getspaces(car))) + _makeansi(colorR, strings.TrimSpace(car)) + resultL
			if color_n+speed < len(color) {
				color_n += speed
			} else {
				color_n = 0
			}
		}
		result = result + "\n" + resultL
	}
	return strings.TrimSpace(result)
}
func Format(text string, second_chars []interface {
}, mode interface {
}, principal_col Colors.col, second_col string) {
	var ctext interface {
	}
	if reflect.DeepEqual(mode, Colorate.Vertical) {
		ctext = mode(principal_col, text, true)
	} else {
		ctext = mode(principal_col, text)
	}
	ntext := ""
	for _, x := range ctext {
		if func() int {
			for i, v := range second_chars {
				if v == x {
					return i
				}
			}
			return -1
		}() != -1 {
			x := Colorate.Color(second_col, x)
		}
		ntext += x
	}
	ntext
}

// Banner

func Box(
	content string,
	up_left string,
	up_right string,
	down_left string,
	down_right string,
	left_line string,
	up_line string,
	right_line string,
	down_line string,
) string {
	l := 0
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(content)
	for _, a := range lines {
		if len(a) > l {
			l = len(a)
		}
	}
	if l%2 == 1 {
		l += 1
	}
	box := up_left + strings.Repeat(up_line, l) + up_right + "\n"
	for _, line := range lines {
		box += left_line + " " + line + strings.Repeat(" ", l-len(line)) + " " + right_line + "\n"
	}
	box += down_left + strings.Repeat(down_line, l) + down_right + "\n"
	return box
}

func SimpleCube(content string) string {
	l := 0
	lines := func(s string) (lines []string) {
		sc := bufio.NewScanner(strings.NewReader(s))
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return
	}(content)
	for _, a := range lines {
		if len(a) > l {
			l = len(a)
		}
	}
	if l%2 == 1 {
		l += 1
	}
	box := "__" + strings.Repeat("_", l) + "__\n"
	box += "| " + strings.Repeat(
		" ",
		int(float64(l)/float64(2)),
	) + strings.Repeat(
		" ",
		int(float64(l)/float64(2)),
	) + " |\n"
	for _, line := range lines {
		box += "| " + line + strings.Repeat(" ", l-len(line)) + " |\n"
	}
	box += "|_" + strings.Repeat("_", l) + "_|\n"
	return box
}

func DoubleCube(content string) string {
	return Box(content, "╔═", "═╗", "╚═", "═╝", "║", "═", "║", "═")
}

func Arrow(icon string, size int, number int, direction interface{}) string {
	var i int
	var line interface{}
	spaces := strings.Repeat(" ", size+1)
	_arrow := ""
	structure := [2]interface{}{size + 2, []int{size * 2, size * 2}}
	count := 0
	if reflect.DeepEqual(direction, "right") {
		for i = 0; i < structure[1][0]; i++ {
			line = strings.Repeat(icon, int(structure[0]))
			_arrow += strings.Repeat(
				" ",
				count,
			) + strings.Join(
				func(repeated []interface{}, n int) (result []interface{}) {
					for i := 0; i < n; i++ {
						result = append(result, repeated...)
					}
					return result
				}([]interface{}{line}, number),
				spaces,
			) + "\n"
			count += 2
		}
		for i = 0; i < structure[1][0]+1; i++ {
			line = strings.Repeat(icon, int(structure[0]))
			_arrow += strings.Repeat(
				" ",
				count,
			) + strings.Join(
				func(repeated []interface{}, n int) (result []interface{}) {
					for i := 0; i < n; i++ {
						result = append(result, repeated...)
					}
					return result
				}([]interface{}{line}, number),
				spaces,
			) + "\n"
			count -= 2
		}
	} else if reflect.DeepEqual(direction, "left") {
		for i = 0; i < structure[1][0]; i++ {
			count += 2
		}
		for i = 0; i < structure[1][0]; i++ {
			line = strings.Repeat(icon, int(structure[0]))
			_arrow += strings.Repeat(" ", count) + strings.Join(func(repeated []interface{}, n int) (result []interface{}) {
				for i := 0; i < n; i++ {
					result = append(result, repeated...)
				}
				return result
			}([]interface{}{line}, number), spaces) + "\n"
			count -= 2
		}
		for i = 0; i < structure[1][0]+1; i++ {
			line = strings.Repeat(icon, int(structure[0]))
			_arrow += strings.Repeat(" ", count) + strings.Join(func(repeated []interface{}, n int) (result []interface{}) {
				for i := 0; i < n; i++ {
					result = append(result, repeated...)
				}
				return result
			}([]interface{}{line}, number), spaces) + "\n"
			count += 2
		}
	}
	return _arrow
}
