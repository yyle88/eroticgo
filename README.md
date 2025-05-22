[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/eroticgo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/eroticgo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/eroticgo)](https://pkg.go.dev/github.com/yyle88/eroticgo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/eroticgo/master.svg)](https://coveralls.io/github/yyle88/eroticgo?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/eroticgo.svg)](https://github.com/yyle88/eroticgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/eroticgo)](https://goreportcard.com/report/github.com/yyle88/eroticgo)

# eroticgo

**eroticgo** is a simple Go package for adding color to your console screen.

## Installation

```bash
go get github.com/yyle88/eroticgo
```

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/yyle88/eroticgo"
)

func main() {
	fmt.Println(eroticgo.RED.Sprint("print RED text"))
	fmt.Println(eroticgo.GREEN.Sprint("print GREEN text"))
	fmt.Println(eroticgo.BLUE.Sprint("print BLUE text"))
}
```

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo in GitHub.
2. Create a feature branch (`git checkout -b feature/xxx`).
3. Commit changes (`git commit -m "Add feature xxx"`).
4. Push to the branch (`git push origin feature/xxx`).
5. Open a pull request.

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with `eroticgo`!** 🎉

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/eroticgo.svg?variant=adaptive)](https://starchart.cc/yyle88/eroticgo)
