<h1 align="center" class="header">Robin</h1>
<p align="center" class="desc">
  A terminal-friendly, dead-simple file server. Written in Go.
</p>

## Table of Content
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Contributions](#contributions)
- [Credits](#credits)
- [License](#license)

## Introduction
[![asciicast](https://asciinema.org/a/508141.svg)](https://asciinema.org/a/508141)

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

Robin is a terminal-friendly, dead-simple file server. Written in Go.

## Features

* Serve static files
* Terminal-friendly, elegant interface (Returns plain text tables)
* `curl`-friendly
* Easily redirect to directories and read files with `/path`

## Installation

```sh
hg clone https://hg.sr.ht/~tsukii/robin # Clone the repo
cd robin
go mod tidy
go build
```

## Contributions

Contributions are welcome. If you found any bugs or want to requrest a feature, please open an issue. 

## Credits

This software uses [echo](https://echo.labstack.com) web framework
and [tablewriter](https://github.com/olekukonko/tablewriter) for displaying beautiful file tables

Thanks [README Templates](https://www.readme-templates.com) and [GitPoint's README](https://github.com/gitpoint/git-point#readme) for the template. This project uses GitPoint's README template.

Thanks [this TOC generator](https://ecotrust-canada.github.io/markdown-toc/) for the Table Of Content.

## License

MIT
