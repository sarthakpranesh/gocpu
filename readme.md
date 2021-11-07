# gocpu

A simple cli tool to handle and watch your CPU.

<br />

## Usage
```txt
Usage gocpu [subcommand] [flags]
subcommand:
    watch   -   see the realtime cpu frequency, updated at 2 seconds by default, 
                can be changed using the "--int" flag.

    turbo   -   sets the turbo on/off depending on the "--enable" flag.

    govern  -   lets you interactively select the cpu governor out the all the 
                available governors.

flags:
    --int   -   used with watch subcommand, value is treated a number of seconds 
                between refreshes
    --enable-   used with turbo subcommand, if passed turbo mode will be enabled 
                else turbo mode will be disabled
```

<br />

## Issues
If you find any bugs/improvements/feature requests please open them [here](https://github.com/sarthakpranesh/gocpu/issues)
