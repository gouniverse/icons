# Icons <a href="https://github.com/gouniverse/icons" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

The following icons are included:

- Bootstrap 1.9.0 

  website: https://icons.getbootstrap.com/
  
  source: https://github.com/twbs/icons
  
  license: MIT
  
- Boxicon 2.1.2

  website: https://boxicons.com/
  
  source: https://github.com/atisawd/boxicons 
  
  license: MIT
  
## Installation

```
go get github.com/gouniverse/icons
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
