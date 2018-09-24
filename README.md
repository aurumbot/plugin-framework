# plugin-framework

This repository defines what fields all plugins must have and the file tree plugins should follow.

## File system

Aurum Plugins should have the following files and directories. (Note that `json/` is only needed if you have a default config file you need to import.)

```
.
├── README.md
├── build
├── json
│   └── *.json
└── src
    └── main
        └── main.go
```

### README.md

The README.md should at the minimum list all the commands defined within the plugin to prevent accidental command overlap.

### Build

build should compile the main.go and all its dependancies into a single .so file in bin/ and error check before hand. See `./build` for more info.

### json/

The `./json/` directory should store json config templates. All .json files within this directory will be moved into the bot's `dat/cfg/$PLUGINNAME/`.

### src/

This is where all the code should go