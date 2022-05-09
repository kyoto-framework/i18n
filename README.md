
# i18n

Minimalistic internationalizing helper.

## Installation

- Install package

  ```bash
  go get github.com/kyoto-framework/i18n
  ```

- Parse translation files from specified directory

  ```go
  i18n.Parse("./i18n")
  ```

- (Optional) Attach template functions

  ```go
    // If you're using kyoto framewrok
    func FuncMap() template.FuncMap {
        return render.ComposeFuncMap(
            i18n.FuncMap(),    // i18n functions
            ...
        )
    }

    // If you're using i18n separately
    func AttachI18N(fmap template.FuncMap) {
        for k, v := range i18n.FuncMap() {
            fmap[k] = v
        }
    }
  ```

## Usage

First, you'll have to create translation files.  
This module uses structured `.yaml` to define such files.
We call this files "pages".

Your pages should be structured as follows:

```yaml
<language-code-1>:
    <group>:
        <key>: <value>

<language-code-2>:
    <group>:
        <key>: <value>
```

Please note, file names also matter.
It will be used as a translation path, as well as a group and key.

To use specified translations, you have `i18n.TranslateStatic(lang, page, group, key)` Go function
and `{{ translateStatic lang page group key }}` template function.

Example:

```go
lang := "en"
content := i18n.TranslateStatic(lang, "index", "home", "content")
```

This module also provides a way to extract translations from dynamic containers (Go maps and structs).
To do this, you can use `i18n.TranslateDynamic(lang, container, field)` Go function
and `{{ translateDynamic lang container field }}` template function, where `container` is a Go map or struct.
Field name will be chosen based on the language provided.
In case of default language ("en" by default), value will be fetched according to the field name.
In case of non-default language, value will be fetched according to the field name + language code (f.e. `ContentRU`, `Content_ru` or `content_ru`).

Example:

```go
data := map[string]string{"Content": "Hello, world!", "ContentES": "Hola, mundo!"}
contenten := i18n.TranslateDynamic("en", data, "Content") // "Hello, world!"
contentes := i18n.TranslateDynamic("es", data, "Content") // "Hola, mundo!"
```

## Motivation

- Almost every solution we found have really overcomplicated usage for our simple needs (take a needed string from a file, accoring to provided language)
- Most of solutions are not providing any template functions
- There are no functions to dynamically take translated value from a struct or map
