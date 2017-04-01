
Initial work was based on the idea / implementation of the golang
`github.com/kelseyhightower/envconfig` package


The example/types.go has a sample structure and sample environment
variable bash file which can be tested and run with the following
commands


```
    go get github.com/davidwalter0/envflagstructconfig

    cd ${GOPATH}/src/github.com/davidwalter0/envflagstructconfig/example
    . environment ; go run main.go types.go

```

The current implementation dumps the initialized structure using the
sourced environment variables in the file `environment` merged with
the text from struct tags like: `default:"text"`

The application environment variable prefix is passed to
initialization with the structure to initialize.

The prefix allows override of the environment variables prefix.

Each environment variable is prefixed with the prefix like so

For prefix app and struct member CC the env var will be uppercased and
underscored: `APP_CC`

The command line flag it will be lower cased and hyphenated: `-cc`

But CaC will be `APP_CA_C` and the flag will be -ca-c

Dynamically altering the prefix for an app instance can be done with
by setting the environment variable for the application prefix

Add code similar to the following before calling

`envflagstructconfig.Initialize(prefix,&structure)`

```
	var prefix = os.Getenv("APP_OVERRIDE_PREFIX")
	if len(prefix) == 0 {
		prefix = "myapp"
	}

    envflagstructconfig.Initialize(prefix,&structure)
```

Set the appropriate environment variable and run

```
    export APP_OVERRIDE_PREFIX=ANOTHER_PREFIX
```

---

Tags parsed

- name:"nameOverride" which will override existing env variable name
  for this member with the camel case underscore insertion to separate
  words and hyphenation of the lower case of the camel case words
  - environment variable
    `export APP_NAME_OVERRIDE="Tom,Jerry"`
  - flag name
    `-name-override`

- usage:"help text for variable"
- short:"abbrev"
- name:"nameoverride"
- default:"initial value(s) for var conforming to language rules"
  - defaults for slices are comma separated lists
  - defaults for maps are comma separated lists of key:value colon
    separated pairs

Partial example from example/types.go

```
    // Specification example test struct
    type Specification struct {
        Debug bool `name:"Debug" short:"d" default:"false" usage:"enable debug mode"`

```

    
---

Limitations, bugs:

The flag implementation for slice is limited to comma separated text
conversion to a slice of strings

- "a,b,c" -> []string -> ["a","b","c"]

The map implementation is restricted to the struct default tag or
environment variables. 

Flag to map aren't working.
