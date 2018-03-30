# charge-notifier

Helps you keep all your integrated batteries healthy by keeping them between set charge thresholds (Default: 0.2 & 0.8) by giving you a notification when they're not every time it's run. 

Made with Linux in mind, but theoretically cross-platform compatible. 

Because it doesn't run as a daemon, *external scheduling is needed. *

I personally use it together with [Jobber](https://dshearer.github.io/jobber/) to run once a minute. 

## Installation
(The way I use it)
`go build charge-notifier.go && sudo cp charge-notifier /usr/bin/charge-notifier && charge-notifier`

(build, install, test)

[Install Jobber](https://dshearer.github.io/jobber/doc/v1.3/#deployment)

Add to $HOME/.jobber: 
```
[jobs]
- name: ChargeNotifier
  cmd: export DISPLAY=:0.0 && /usr/bin/charge-notifier  # shell command to execute
  time: '10 * * * * *'  # SEC MIN HOUR MONTH_DAY MONTH WEEK_DAY.
  onError: Continue  # what to do when the job has an error: Stop, Backoff, or     Continue
  notifyOnError: false  # whether to call notifyProgram when the job has an error
  notifyOnFailure: true  # whether to call notifyProgram when the job stops due to   errors
  notifyOnSuccess: false  # whether to call notifyProgram when the job succeeds

```

`export DISPLAY=:0.0` is important for notify-send to reach the user's display when it's not run in the home shell. 

### Background
This tool was made to scratch a personal itch to keep my laptop's battery health (because it makes liberal use of the safety margins & after more elaborate attempts failed) and to learn a bit about Go and Git. 

Full charging generally puts higher strain on Lithium-Ion batteries. This is why it is [recommended to be avoided for electric vehicles](https://electrek.co/2017/09/01/tesla-battery-expert-recommends-daily-battery-pack-charging/) and many car manufacturers include tools to prevent it. Notebook manufacturers also do this on their more expensive lines. I have found out that many charge controllers from Texas Instruments [include a feature to rewrite the maximum charge voltage](http://www.ti.com/product/BQ24707/datasheet/detailed_description#SLUSA788847) via SMBus, however while trying around noticed that my chip probably doesn't - and I am not willing to risk much on a [chinese-made laptop without a meaningful warranty](https://xiaomi-mi.com/notebooks/xiaomi-mi-notebook-air-133-classic-edition-i5-8gb256gb-silver/). 
