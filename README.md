# proj

**A command line tool for workspace management**

## Contents
* [Commands](#commands)
  * [`init`](#init)
    * [`cpp`](#init-cpp)
    * [`java`](#init-java)
  * [`class`](#class)
* [Internal](#internal)

# Commands

[jump to top](#proj)

## `init`

Initializes a new repository, use on of the subcommands to specify which language the 
project will be written in.

Creates a README.md for the project, and other files specific to the project's language.

#### `init` Subcommands

* [`cpp`](#init-cpp)
* [`java`](#init-java)

#### `init` Flags

|Flag|Short|Description|
|---|---|---|
|`--git`|`-g`|Initializes a git repository with the project. Creates git's default `.gitignore` for the specified language.|

[jump to top](#proj)

## `init cpp`

Initialzes a new c++ project. Creates a [Makefile](path/to/cpp-Makefile), a [path modifier](path/to/path-cpp.sh), as well as `src`, `bin`, `cmd`, and `data` directories.

[jump to top](#proj)

## `init java`

[jump to top](#proj)

## `class`

Creates a new class or class files with the given name for the language of the project

[jump to top](#proj)

# Internal

`proj` maintains project details in a `.proj/` folder at the root of the project.

`~/.proj_config` maintains write protected copies of Makefiles, class files, and other scripts which are copied into directories.

[jump to top](#proj)