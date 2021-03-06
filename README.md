# charge-notifier

Helps you keep all your integrated batteries healthy by keeping them between set charge thresholds (Default: 0.2 & 0.8) by giving you a notification when they're not every time it's run. Auto-suspends (default: at 0.05) using `systemctl suspend`. 

## Installation
(The way I use it)

Build & copy to the bin folder

```
go build charge-notifier.go
cp charge-notifier $HOME/.local/bin/
```

Deploy as a systemd user service

```
cp charge-notifier.service $HOME/.config/systemd/user/
systemctl --user enable --now charge-notifier
```

### Background
This tool was made to scratch a personal itch to keep my laptop's battery health (because it makes liberal use of the safety margins & after more elaborate attempts failed) and to learn a bit about Go and Git. 

Full charging generally puts higher strain on Lithium-Ion batteries. This is why it is [recommended to be avoided for electric vehicles](https://electrek.co/2017/09/01/tesla-battery-expert-recommends-daily-battery-pack-charging/) and many car manufacturers include tools to prevent it. Notebook manufacturers also do this on their more expensive lines. I have found out that many charge controllers from Texas Instruments [include a feature to rewrite the maximum charge voltage](http://www.ti.com/product/BQ24707/datasheet/detailed_description#SLUSA788847) via SMBus, however while trying around noticed that my chip probably doesn't - and I am not willing to risk much on a [chinese-made laptop without a meaningful warranty](https://xiaomi-mi.com/notebooks/xiaomi-mi-notebook-air-133-classic-edition-i5-8gb256gb-silver/). 
