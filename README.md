# Maparo

## Installation

You can use pre-compiled releases compiled for Linux, Mac and Windows.

## Design Principles

- Modules can be used in parallel with other modules. Modules can be initiated
  several times. This allows more complex campaigns like *Realtime Response Under
  Load (RRUL)* like FLENT is doing.
- *Maparo* return either a JSON set or human formatted data. *Maparo* do
  not draw or visualize any data. To visualize data you can use Python
  scripts using the powerful and flexible matplotlib library.

## Development

### Prepare Development Environment

For Debian based systems install go suite:

```
sudo apt-get install golang-go
```
