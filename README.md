<h1 align="center">Advent of Code gen tool</h1>
<p align="center">A cli tool for generating AoC solutions</p>

<p align="center">

<a style="text-decoration: none" href="https://github.com/KarolosLykos/advent-of-code-gen/actions?query=workflow%3AGo+branch%3Amain">
<img src="https://img.shields.io/github/actions/workflow/status/KarolosLykos/advent-of-code-gen/go.yml?style=for-the-badge" alt="Build Status">
</a>
    
<a style="text-decoration: none" href="go.mod">
<img src="https://img.shields.io/badge/Go-v1.17-blue?style=for-the-badge" alt="Go version">
</a>

<br />

<a style="text-decoration: none" href="https://github.com/KarolosLykos/advent-of-code-gen/stargazers">
<img src="https://img.shields.io/github/stars/KarolosLykos/advent-of-code-gen?style=for-the-badge" alt="Stars">
</a>

<a style="text-decoration: none" href="https://github.com/KarolosLykos/advent-of-code-gen/fork">
<img src="https://img.shields.io/github/forks/KarolosLykos/advent-of-code-gen?style=for-the-badge" alt="Forks">
</a>

<a style="text-decoration: none" href="https://github.com/KarolosLykos/advent-of-code-gen/issues">
<img src="https://img.shields.io/github/issues/KarolosLykos/advent-of-code-gen?style=for-the-badge" alt="Issues">
</a>


<br>
<br>

<p align="center">
    <a style="text-decoration: none" href="https://github.com/KarolosLykos/advent-of-code-gen/releases">
        <img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
    </a>
</p>

## Overview

`aoc` is a CLI tool written in Go to streamline your **Advent of Code** workflow.  
It helps you:

- Initialize a Go project structure for AoC puzzles
- Store and manage your Advent of Code session cookie
- Automatically download puzzle input files
- Generate ready-to-edit Go templates for each puzzle day/part

---

## Installation

```bash
go install github.com/KarolosLykos/advent-of-code/cmd/aoc@latest
```

## Usage

### 1. Initialize a Project

```bash
mkdir my-aoc
cd my-aoc
aoc init
```

This creates:

- Project directory
- aoc.yaml configuration file

### 1. Set Your Session Cookie
Copy your session cookie from Advent of Code after logging in:
```bash
aoc session -v "your_session_cookie_value"
```
This enables automatic puzzle input downloads.

### 3. Generate Puzzle Templates
Generate today’s puzzle:
```bash
aoc gen
```

Generate a specific day/year:
```bash
aoc gen -y 2022 -d 7
```

Each generated solution includes:
- A main function ready to run the solutions
- Helper functions for reading AoC inputs

```bash
projectDir/
└── {2022}/day{07}/
  ├── input.txt
  └── main.go
```

## Commands Reference
Command	Description
| Command                      | Description                                  |
| ---------------------------- | -------------------------------------------- |
| `aoc init`                   | Initialize project folder, Go module & config        |
| `aoc session -v <cookie>`    | Save AoC session cookie                      |
| `aoc gen`                    | Generate template & download input for today |
| `aoc gen -y <year> -d <day>` | Generate template for a specific puzzle      |

## Examples
```bash
# Initialize project
aoc init

# Save session cookie
aoc session -v "MY_SESSION_COOKIE"

# Generate today's puzzle
aoc gen

# Generate a specific puzzle
aoc gen -y 2023 -d 5

# Run a day's code
go run {projectDir}/2023/day05/main.go
```
## Contributing

Contributions are welcome! You can:
- Add helpers for input parsing

Open an issue or submit a pull request with your changes.

## Notes & Security

The session cookie grants access to your AoC inputs. Do not commit it to a public repository.