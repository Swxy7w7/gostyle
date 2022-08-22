package gostyle

func colorate(color string) string {
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
