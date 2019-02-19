# Internal

The internal organization of proj is as follows:

`proj` maintains project details in a `.proj/` folder at the root of the project. This `.proj` is arranged in the following way:

```
- | .proj
  `- | 
     | projData.json
```

When `proj` is installed, it creates a `~/.proj_config` folder where a number of resources and settings are stored.

`~/.proj_config` maintains write protected copies of Makefiles, class files, and other scripts which are copied into directories. See [resources.md](./resources.md) for more info on resources. These can be modified at will by the user to change the behavior of the tool.



[back](../README.md)