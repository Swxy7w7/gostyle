<h1 align="center">GoStyle</h1>
<br>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-beta-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/xtekky/gostyle/blob/main/README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="https://github.com/xtekky/gostyle" target="_blank">
    <img alt="Maintenance" src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" />
  </a>
  <a href="https://github.com/xtekky/gostyle/blob/main/LICENSE" target="_blank">
    <img alt="License: EPL-2.0" src="https://img.shields.io/github/license/billythegoat356/pystyle" />
  </a>
</p>

> **GoStyle** is a golang library to make very beautiful TUI designs.
> <br>
> Inspired by **pystyle**,

## Example of usage
I am working on implementing all functions of pystyle in the file are the currently translated ones

```go
func main() {
	fmt.Println(Color(blue, "Hello World", reset))
}
```

![image](https://user-images.githubusercontent.com/98614666/185834904-d015d890-3973-4ad5-987b-21aacaf0338e.png)
-----------------------------------------
```go
func main() {
	fmt.Println(Symbol("+", blue, purple, blue, "Hello World"))
}
```

![image](https://user-images.githubusercontent.com/98614666/185837133-b932c161-0d26-40ff-b2b1-2b124bea933f.png)
