# Slack Status Update CLI
A CLI for updating your Slack Status using a preset list of statuses.

Please note that this is still WIP. I'm labelling it as an alpha pre release. But I don't want to wait without shipping, so whatever I have will be whatever goes.

In the current iteration of the CLI, the CLI can read a Slack Token and set a status based off a predefined set of statuses that you keep in a text file. The process of using it is currently quite involved but that'll become easier over time.

Please open up any issues you find on Github. I'm actively developing the project and committing to keep it improving.

You can follow transparent development of the project [on my little engineering log site](http://eng.adnanissadeen.com/projects/slack-status-updater-cli.html)

## How to use it

### Setup steps:
* Download the binary for your relevant OS from [here](https://github.com/kiriappeee/slack-status-updater-cli/releases). Currently available only for 64 bit Mac OS and Linux

* Move your binary to somewhere in your `$PATH`. I personally keep mine in `/usr/local/bin` which I've changed ownership of to myself to avoid using `sudo`. Run `slackstatus -h` to see if things work correctly. You should see this:

```
NAME:
   Slack Status Updater - Painless status updating via your Terminal

USAGE:
   slackstatus [global options] command [command options] [arguments...]

VERSION:
   0.0.1 alpha super pre release

COMMANDS:
     set      Sets a status from a predefined list of statuses
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

* Create a folder called `$HOME/.config/ssucli`

  `mkdir -p $HOME/.config/ssucli`

* Then create a file called `statuses.yaml` and `tokenconfig` inside the `ssucli` folder you just created.

* You can use the following template to get started with adding statuses:

```yaml
- statusName: lunch
  emoji: bowl_with_spoon
  statusText: Having lunch

- statusName: resting
  emoji: bath
  statusText: Resting

- statusName: deepwork
  emoji: hammer_and_wrench
  statusText: In Deep work mode

- statusName: driving
  emoji: car
  statusText: On the road

- statusName: mosque
  emoji: pray
  statusText: At mosque

- statusName: ping
  emoji: bell
  statusText: Working and available for pings :)

- statusName: gym
  emoji: weight_lifter
  statusText: Out getting exercise at the gym!

- statusName: meeting
  emoji: ''
  statusText: In a meeting

- statusName: happy
  emoji: grin
  statusText: ''
```

_Note that the last two statuses don't have all the attributes filled in. The statuses yaml spec allows either emoji or statusText to be empty **but not both**._

* Get your personal Slack token from [the legacy tokens page](https://api.slack.com/custom-integrations/legacy-tokens) and paste it into the `$HOME/.config/ssucli/tokenconfig` file

* You can now run `slackstatus set ping` to test changing your status. If it works correctly you should see:

```
Status post result: Status set for ping
```

* The effect of the above will be to set your status emoji to bell and your status message to "Working and available for pings :)". 

## Adding new statuses.

To add a new status, simply open `$HOME/.config/ssucli/statuses.yaml` and add a new status in the following format:

```yaml
- statusName: name-of-the-status-you-use-when-calling-from-terminal
  emoji: the_emoji_you_want
  statusText: The status message you want to associate with this.
```

That's it :).

## Next steps:

Some immediate next steps I have in mind center around CRUD operations of statuses and setting up a basic `statuses.yaml` file as a placeholder when doing the first run.
