# YAML to Go
This application translates the adventure.yaml file into a structure that can be consumed directly by the open-adventure application.

Aims to duplicate the functionality of
[make_dungeon.py](https://gitlab.com/esr/open-adventure/blob/master/make_dungeon.py).

## Building

First, build the yaml2go binary.

    PS D:\open-adventure\cmd\yaml2go> go build

The run it from the root.

    PS D:\open-adventure> .\cmd\yaml2go\yaml2go.exe
    info: INPUT                          == "adventure.yaml"
    info: OUTPUT_PATH                    == "./state/"
    info: completed successfully

It will create files in the `./state/` directory.