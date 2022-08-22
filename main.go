package main

import (
	"bufio"
	"bytes"
	"fmt"
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

func Symbol(symbol string, col string, col_left_right string, left string, right string) string {

	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		col_left_right, left, symbol, right, col_left_right, col, col_left_right,
	)
}

func Color(color string, text string, end string) string {
	return _maketext(color, text, end)
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

func main() {
	fmt.Println(DoubleCube("Hello World"))
}
