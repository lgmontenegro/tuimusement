# TUI Musement Assessment Challenge

Based on this [Gist](https://gist.github.com/hpatoio/3aeea8159fb9046a2feba75d39a8d21e "link to assessment description") I created the following solution.

## Installation

The repository already contains the external packages downloaded in the `vendors`' folder. Despite this, here are the steps to compile the CLI both in Windows and in Linux operational system.

After cloned or downloaded and unzipped the zip file, go to the folder and enter the following instrtuctions:

1. `go mod vendor`
2. `go mod tidy`
3. `go build .`

After that, a file called `tuimusement`(Linux)/`tuimusement.exe`(Windows) will be created inside this folder.

Just execute this file to use the CLI.

> Pay attention to `config.json` file! It must be in the same folder as the CLI application.

## Configuration (`config.json`)

**tuiamusement** CLI application reads the `config.json` file located at the same folder as the application.

`config.json` is defined as below:

```{
    "tui_musement_api" : "string",
    "weather_api" : "string",
    "weather_key": "string",
}
```

Where:
+ tui_musement_api is the endpoint URL to musement City API
+ weather_api is the endpoint URL to weather API
+ weather_key is the key to access weather API

## Yes, it has a `makefile`

To make it easier to compile and run the CLI, a makefile was created and is used in 2 steps:

+ Compile the source code:

```
make compile
```

+ run the client

```
make run
```

> Warning: `makefile`/`make` runs on Linux. To have `makefile` running into Windows you have to look after **`make` for Windows** that can be installed using [chocolatey](https://chocolatey.org/ "the Windows package manager") (the Windows package manager) or downloading it [here](http://gnuwin32.sourceforge.net/packages/make.htm "make for Windows")

> Warning: `makefile` uses `docker` to compile, so it's mandatory to have Docker installed on your operating system. For more information on how to install Docker, please, take a look [here](https://docs.docker.com/engine/install/).

## About this CLI

### 3rd Parties Packages

At this project, I use [Cobra](https://github.com/spf13/cobra#readme "Cobra") and [Viper](https://github.com/spf13/viper#readme "Viper") packages due to the facility to construct CLI apps and read configuration file.

Because of my lack of time I prefer not to use the flags and arguments that Cobra and Viper give to the application.

### CLI architecture designed

After that, I decided to implement the following structure (inside internal):

+ application: where the code is called and executed;
+ bootstrap: a kind of middleware, where I create the application objects via configuration values;
+ domain: like a model, represents the data that is used by CLI after retrieving it from the respective APIs;
+ service: service has the abstraction to connect into the webservers that provide data to CLI application. It uses a contract that implements an inversion of responsability due to the characteristcs of each webservice;
+ task: Implements the service according to the referred API.

### ToDo

+ a more flexible bootstrap;
+ better API error handler;
+ code annotations to a better understanding;
+ unit tests;