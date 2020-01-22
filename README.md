# daily-workflow

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
chooserOpts = "--promot '>'"

[daily]
path = "daily"
```

## Directory Structure

- {root}/daily/yyyy/mm/dd.md
