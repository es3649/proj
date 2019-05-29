# proj

**A command line tool for workspace management**

## Contents
* [Commands](#commands)
  * [`init`](#init)
    * [`cpp`](#init-cpp)
    * [`go`](#init-go)
    * [`java`](#init-java)
    * [`python`](#init-python)
  * [`add`](#add)
  * [`class`](#class)
  * [`mv`](#mv
* [Internal](docs/internal.md)

# Commands

[jump to top](#proj)

## `init`

Initializes a new repository, use one of the subcommands to specify which language the 
project will be written in.

Creates a README.md for the project, and other files specific to the project's language.

#### `init` Subcommands

* [`cpp`](#init-cpp)
* [`go`](#init-go)
* [`java`](#init-java)
* [`python`](#init-python)

#### `init` Flags

|Flag|Short|Description|
|---|---|---|
|`--git`|`-g`|Initializes a git repository with the project. Creates git's default `.gitignore` for the specified language.|

[jump to top](#proj)

### `init cpp`

Initializes a new c++ project. Creates a [Makefile](path/to/cpp-Makefile), a [path modifier](path/to/path-cpp.sh), as well as `src`, `bin`, `cmd`, and `data` directories.

[jump to top](#proj)

### `init go`

[jump to top](#proj)

### `init java`

[jump to top](#proj)

### `init python`

[jump to top](#proj)

### `init web`

[jump to top](#proj)

## `add`

TODO is this a good idea? It might help with `mv`, but it might be extraneous, and we can determine mv behavior by file extension.

Adds language support to a project. 

[jump to top](#proj)

## `class`

Creates a new class or class files with the given name for the language of the project

[jump to top](#proj)

## `mv`

As would be expected, `mv` moves a file or directory. Based on the language specifications, this move also updates packages, includes/imports and references between files.

# Internal

`proj` maintains project details in a `.proj/` folder at the root of the project. This `.proj` is arranged in the following way:

```
- | .proj
  `- | projData.json
```

`~/.proj_config` maintains write protected copies of Makefiles, class files, and other scripts which are copied into directories.

[jump to top](#proj)
