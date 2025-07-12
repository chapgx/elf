# ELF

<img src="./assets/elflogo.svg" width="100px"/>

ELF (Encrypt Local Configuration File) is a light weight library and CLI to handle sensitive data encryption e.x (connection credentials, passwords, keys) on your local system. Main use is for processes that communicate with external secure services and require an API key to do it.


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


