# plugin-framework

This repository defines what fields all plugins must have and the file tree plugins should follow.

## File system

Aurum Plugins should have the following files and directories. (Note that `json/` is only needed if you have default config file(s) you need to import.)

```
.
├── README.md
├── build
├── json/
│   └── *.json
└── src/
```

### README.md

The README.md should list the names of all the commands defined within the plugin to prevent accidental command overlap.

### build

`build` should compile the main.go and its dependancies into a single .so file in the `bin/`. See `./build` for more info/a template. (Note if you try to build `./...`, your main.go will fail if you don't use `buildmode=plugin` as you shouldn't have a `func main()`.)

### json/

The `./json/` directory should store json config templates. All .json files within this directory will be moved into the bot's `dat/cfg/$PLUGINNAME/`.

Note that the name of your .so will be considered the name of your plugin for json. So plugin.so will have its json files moved to "dat/cfg/plugin/"

### src/

This is where all the code should go. Anything in here is fair game as long as your `build` compiles it all into a single .so.

## main.go

To hook up into the bot correctly, your command must at least define a map `Commands` of type `map[string]*foundation.Command` (foundation pulling from [github.com/aurumbot/lib/foundation](https://github.com/aurumbot/lib/blob/master/foundation/foundation.go)).

If your plugin requires autonomous and external handlers such as checking for deleted messages, package foundation also exports a `Session` of `*discordgo.Session`; call AddHandler from that in the init.

