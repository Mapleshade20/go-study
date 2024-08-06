## Go naming

Use interior capital letters instead of underscores.
eg. parseRequestLine <- parse_request_line

## Declarations

```go
const boilingF = 212.0  // package-level
func main() {
  var f = boilingF  // function-level
}
```

### Variables

formula: `var name type = expression`

`var b, f, s = true, 2.3, "four"`
`var boiling float64 = 100`
`var names []string`

short declaration: `i := 100`

tuple assignment: `i, j = j, i`

## Pointers

```go
x := 1
pointer := &x
*p = 2
fmt.Println(x)  // 2
```

zero value for a pointer is `nil`

It's safe for a function to return the address of a local variable, though its variable name is no longer accessible outside.

```go
func incr(p *int) int {
  *p++
  return *p
}
count := 3
result := incr(&count)
```

Pointers are key to the `flag` package. `flag.Bool(...)` or `flag.String(...)` returns a **pointer**, which must be used as `*x` in the post-process. `flag.Parse()` must be called before the flags are used. Non-flag arguments are accessed with `flag.Args()`

## Comparability

Only variables of the **same type** can be compared.

## Type Declarations

`type name underlying-type`. For example `type Celsius float64`

Declaring new types helps preventing mismatching-type calculations, and also makes it possible to define new methods for values of such types.

For example:
```go
func (c Celsius) String() string { return fmt.Sprintf("%g degrees", c)}
c := Celsius(56.0)
fmt.Printf("%s", c)
```
