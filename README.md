# Icons

The following icons are included:

- Bootstrap 1.13.1 

  website: https://icons.getbootstrap.com/
  
  source: https://github.com/twbs/icons
  
  license: MIT
  
- Boxicon 2.1.4

  website: https://boxicons.com/
  
  source: https://github.com/atisawd/boxicons 
  
  license: MIT
  
## Installation

```
go get github.com/dracory/icons
```

## Usage


- Use as HB Tag

```go
iconTag := Icon("bi-globe", 120, 120, "white")
```

- Use as HTML

```go
iconHtml := Icon("bi-globe", 120, 120, "white").ToHTML()
```

- Use as String (DEPRECATED)

```go
log.Println(icons.BootstrapMusicNote)
```

## To Do

https://github.com/astrit/css.gg
