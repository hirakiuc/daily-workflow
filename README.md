# daily-workflow

![Go](https://github.com/hirakiuc/daily-workflow/workflows/Go/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/hirakiuc/daily-workflow/branch/master/graph/badge.svg)](https://codecov.io/gh/hirakiuc/daily-workflow)

Daily works

- daily report

## Configuration

default path: `~/.config/wf/config.toml`

```
[common]
root = "~/Desktop/workspace/tool"
editor = "nvim"
finder = "ag"
finderOpts = "--vimgrep"
chooser = "/usr/local/bin/peco"

[daily]
path = "daily"
```

## Directory Structure

- {root}/daily/yyyy/mm/dd.md
