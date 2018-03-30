# charge-notifier

Helps you keep all your integrated batteries healthy by keeping them between set charge thresholds (Default: 0.2 & 0.8) by giving you a notification when they're not every time it's run. 

Made with Linux in mind, but theoretically cross-platform compatible. 

Because it doesn't run as a daemon, *external scheduling is needed. *

I personally use it together with [Jobber](https://dshearer.github.io/jobber/) to run once a minute. 

##Installation
(The way I use it)
'go build charge-notifier.go && sudo cp charge-notifier /usr/bin/charge-notifier && charge-notifier'
(build, install, test)

Add to $HOME/.jobber: 
'''
[jobs]
- name: ChargeNotifier
  cmd: export DISPLAY=:0.0 && /usr/bin/charge-notifier  # shell command to execute
  time: '10 * * * * *'  # SEC MIN HOUR MONTH_DAY MONTH WEEK_DAY.
  onError: Continue  # what to do when the job has an error: Stop, Backoff, or Continue
  notifyOnError: true  # whether to call notifyProgram when the job has an error
  notifyOnFailure: true  # whether to call notifyProgram when the job stops due to errors
  notifyOnSuccess: true  # whether to call notifyProgram when the job succeeds
'''
'export DISPLAY=:0.0' is important for notify-send to reach the user's display when it's not run in the home shell. 


###Background story
This tool was made to scratch a personal itch to keep my laptop's battery health (because it makes liberal use of the safety margins & after more elaborate attempts failed) and to learn a bit about Go and Git. 
