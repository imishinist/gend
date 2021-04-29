# gend

This is a template extension tool for generating random strings.
The template will be from the standard Go package.

For more information, please refer to the following URL.

- https://golang.org/pkg/text/template/

# Functions inside templates

| function name | detail | note |
| --- | --- | --- |
| join | This will join string slices | [strings.Join](https://golang.org/pkg/strings/#Join) |
| trimjoin | This will trim the string and join slices | |
| trim | This will trim whitespaces | [strings.TrimSpace](https://golang.org/pkg/strings/#TrimSpace) |

# Examples

```yaml
rules:
  - key: test1
    value:
      # generate a fixed value
      static: "value1"
    # generator defines a rule to generate a string from key and value
    generator:
      # value is an array, so we need to join it.
      templates: |
        {{- .key }}:{{ join .values "," }}
  - key: test2
    value:
      # generate 1 or 2 or 3 or 4 or 5
      allowed: [1, 2, 3, 4, 5]
    # define values count
    length:
      static: 3
    generator:
      templates: |
        {{- .key }}:{{ join .values "," }}
  - key: test3
    value:
      # generate [1, 1000] value with step 10
      # step is optional
      # range: [from, to, step]
      range: [1, 1000, 10]
    length:
      # It is determined by the following ratio
      occurrence: [0, 1, 2]
    generator:
      templates: |
        {{- .key }}:{{ join .values "," }}
  - key: test4
    value:
      generator:
        bash: date
    generator:
      # reference with $key, $values(joined with ",") variables.
      bash: |
        echo "$key:$values"
generator:
  templates: |
    {{ trimjoin .items "," }}
```

# License

MIT

# Author

Taisuke Miyazaki (a.k.a imishinist)

