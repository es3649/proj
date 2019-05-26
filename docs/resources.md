# Resources

this is where the resources are stored. Later they will go in `~/.proj_config`

I'm thinking that inside of `~/.proj_config` we should have files which tell us which files to create and to copy over from the resources for that language.
Perhaps a manifest.json for each language.

``` json
{
    "folders": [
        "bin",
        "cmd",
        "data",
        "src",
        "tools"
    ],
    "copyFiles" : [
        {
            "name":"Makefile",
            "mode":644
        },
        {
            "name":"path.sh",
            "mode":755
        },
        {
            "name":"tools/graphCompiler.py",
            "mode":755
        }
    ]
}
```

Then those files will be retrieved and copied over upon new file creation. 
Each `copyFile.name` should have a path relative to the location of the `manifest.json`, and `copyFile.mode` is an octal signifying the permissions the destination file shall have.

We might have some other settings, like default text headers we will put in new class files.
We should probably store some `git` settings, or have a way to fetch those.

## Config Requirements by Command
| Command | Needs global | Needs Local |
|---|---|---
| init | yes | (created here) |
| class | maybe | yes|
| mv | no | yes |
| version | no | no |