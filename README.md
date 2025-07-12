<img src="./assets/elflogo.svg" width="100px"/>

ELF (Encrypt Local Configuration File) is a light weight library and CLI to handle sensitive data encryption e.x (connection credentials, passwords, keys) on your local system. Main use is for processes that communicate with external secure services and require an API key to do it.

`ELF IS UNDER ACTIVE DEVELOPMENT AND DESIGN. THINGS MAY BREAK IN COMING UPDATES`


## QUICK START

Create local environment. This will create the following directory structure

- .elf (at home directory)
  - elf.db

```sh
$ elf init
```


If you want to delete the local environment run the following command. Be aware this is a destructive command and all data will be lost.

```sh
$ elf torch
```


Enter REPL mode:
Executing `elf` without any arguments or flags will start in you in REPL mode.

```sh
$ elf
```


## Help Command

```sh
$ elf help

```


Output
```sh
Commands
  db      SQLITE database
  exit    Exits REPL
  init    initialized elf environment
  torch   Cleans up the environment this action is finite and not reversible
  help    Displays help information
  clear   Clears terminal
  create  Creates a resource
```
