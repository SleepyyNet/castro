---
name: Functions
---

# Functions

Castro provides a list of already defined functions you can use on your templates:

- [url](#url)
- [vocation](#vocation)
- [serverName](#servername)
- [serverMotd](#serverMotd)
- [nl2br](#nl2br)
- [urlEncode](#urlencode)
- [isDev](#isdev)
- [str2html](#str2html)
- [str2url](#str2url)
- [menuPages](#menuPages)

# url

Creates an absolute URL for Castro. You should use this when creating links inside Castro.

```html
<a href="{{ url "subtopic" "login" }}></a> 
```

You can pass as many strings as you want to create the URL.

# vocation

Returns the name of the given vocation identifier.

```html
<p>{{ vocation 12 }}</p>
```

# serverName 

Returns the `config.lua` server name.

```html
{{ serverName }}
```

# serverMotd

Returns the `config.lua` server motd.

```html
{{ serverMotd }}
```

# nl2br

Similar to the `PHP` function. It converts all newlines to `<br>`. Useful for textareas for example.

```html
{{ nl2br .text }}
```

# urlEncode

Encodes the given value so its safe to use inside an URL.

```html
<a href="/test?name={{ urlEncode .name }}">Test</a>
```

# isDev

Checks if Castro runs on development mode.

```html
{{ if isDev }}
<p>Development</p>
{{ end }}
```

# str2html

Converts the given string to HTML output. By default Castro sanitizes all HTML output for variables.

```html
{{ str2html .text }}
```

# str2url

Converts the given string to a safe URL output. By default Castro sanitizes URL values to prevent any kind of attacks.

```html
<a href="/test/{{ str2url "hello }}">
```

# menuPages

Returns an array of menu pages, usually you call this function to render your template menu pages.

```html
{{ range $index, $element := menuPages }}
    <li>Menu page number {{ $index }}</li>
{{ end }}
```

This function returns an array of menu pages, each page has the following attributes you can use:

- [Name](#name)
- [Link](#link)
- [Dropdown](#dropdown)
- [DropdownItems](#dropdownitems)

## Name

Returns the menu item name.

```html
{{ range $index, $element := menuPages }}
    <li>{{ $element.Name }}</li>
{{ end }}
```

## Link

Returns the menu item url.

```html
{{ range $index, $element := menuPages }}
    <li><a href="{{ url $element.Link }}">{{ $element.Name }}</a></li>
{{ end }}
```

## Dropdown

Boolean field that returns true or false depending if the page is a dropdown.

```html
{{ range $index, $element := menuPages }}
    {{ if not $element.Dropdown }}
        <li><a href="{{ url $element.Link }}">{{ $element.Name }}</a></li>
    {{ end }}
{{ end }}
```

## DropdownItems

Returns an array of all the items inside a dropdown.

```html
{{ range $index, $element := menuPages }}
    {{ if $element.Dropdown }}
        {{ range $base, $item := $element.DropdownItems }}
            Dropdown {{ $index }} - Sub-item number {{ $base }}
        {{ end }}
    {{ end }}
{{ end }}
```

This function returns an array of sub-items with the following attributes:

- [Link](#link) (same as a non-dropdown item)
- [Text](#text)

### Text

Returns the name of the dropdown item.

```html
{{ range $index, $element := menuPages }}
    {{ if $element.Dropdown }}
        {{ range $base, $item := $element.DropdownItems }}
            <li><a href="{{ url $item.Link }}">{{ $item.Text }}</a></li>
        {{ end }}
    {{ end }}
{{ end }}
```