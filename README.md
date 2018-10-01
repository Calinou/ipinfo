# ipinfo

A command-line utility to retrieve your current IP address and related
information, or display information related to an IP address (such as
the country and Internet service provider).

## Usage

### Examples

#### Looking up your current IP address

```text
$ ipinfo

    IP address   ... # Results will be based on your IP address
      Hostname   ...
          City   ...
        Region   ...
       Country   ...
      Location   ...
   Postal code   ...
  Organization   ...
```

#### Looking up an arbitrary IP address

```text
$ ipinfo 8.8.8.8

    IP address   8.8.8.8
      Hostname   google-public-dns-a.google.com
          City   Mountain View
        Region   California
       Country   US
      Location   37.3860,-122.0840
   Postal code   94035
  Organization   AS15169 Google LLC
```

### Reference

```text
NAME:
   ipinfo - Retrieve information about IP addresses

USAGE:
   ipinfo [IP address]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## License

Copyright © 2018 Hugo Locurcio and contributors

Unless otherwise specified, files in this repository are licensed under the
MIT license, see [LICENSE.md](LICENSE.md) for more information.
