# Kapow!

Why is it named Kapow!. You can pronounce it and it is short and
meaningless... like every good UNIX command ;-)

TODO: Too hard to jump 

## Rationale

The **web** and the **shell** are two different beasts, both packed with
history.

There are some concepts in HTTP and the shell that **resemble each other**.

```
                 +------------------------+-------------------------+
                 | HTTP                   | SHELL                   |
  +--------------+------------------------+-------------------------+
  | Input        | POST form-encoding     | Command line parameters |
  | Parameters   | GET parameters         | Environment variables   |
  |              | Headers                |                         |
  |              | Serialized body (JSON) |                         |
  +--------------+------------------------+-------------------------+
  | Data Streams | Response/Request Body  | Stdin/Stdout/Stderr     |
  |              | Websocket              | Input/Output files      |
  |              | Uploaded files         |                         |
  +--------------+------------------------+-------------------------+
  | Control      | Status codes           | Signals                 |
  |              | HTTP Methods           | Exit Codes              |
  +--------------+------------------------+-------------------------+
```

Any tool designed to give an HTTP interface to an existing shell command
**must map concepts of boths**. For example:

- "GET parameters" to "Command line parameters"
- "Headers" to "Environment variables"
- "Stdout" to "Response body"

Kapow! is not opinionated about the different ways you can map both worlds.
Instead it provides a concise language used to express the mapping and a set
of common defaults.


### Why not tool...?

All the alternatives we found are **rigid** about how they match between HTTP
and shell concepts.

* [shell2http](https://github.com/msoap/shell2http): HTTP-server to execute
  shell commands. Designed for development, prototyping or remote control.
Settings through two command line arguments, path and shell command.
* [websocketd](https://github.com/joewalnes/websocketd): Turn any program that
  uses STDIN/STDOUT into a WebSocket server. Like inetd, but for WebSockets.
* [webhook](https://github.com/adnanh/webhook): webhook is a lightweight
  incoming webhook server to run shell commands.
* [gotty](https://github.com/yudai/gotty): GoTTY is a simple command line tool
  that turns your CLI tools into web applications. (For interactive commands
only)
* [shell-microservice-exposer](https://github.com/jaimevalero/shell-microservice-exposer):
  Expose your own scripts as a cool microservice API dockerizing it.

Tools with a rigid matching **can't evade** *[impedance
mismatch](https://haacked.com/archive/2004/06/15/impedance-mismatch.aspx/)*.
Resulting is an easy-to-use software, convenient in some scenarios but
incapable in others.


### Why not my good-old programming language...?

* Boilerplate
* Custom code = More bugs
* Security issues (Command injection, etc)
* Dependency on developers
* **"A programming language is low level when its programs require attention to
  the irrelevant"** *Alan Perlis*

Kapow! aims to be halfway from one of the mentioned tools and a general
programming language. A limited scripting language. Think of *awk*, 20 years
later, for HTTP.


TODO: Definition
TODO: Intro to Architecture


# API
## Spec'ing the endpoints

## Usage Example

## Test Suite Notes

The test suite is located on [blebleble] directory.
You can run it...


# Framework
## Commands
Any compliant implementation of Kapow! must provide these commands:

### `kapow`
This implements the server, yaddayadda

#### Example

### `kroute`
TODISCUSS: maybe consider using `kapow route` instead

#### Example

### `request`

#### Example

### `response`

#### Example

## Full-fledged example  (TODO: express it more simply)


## Test Suite Notes

# Server


## Test Suite Notes