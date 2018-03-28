# Slack Status Update CLI
A CLI for updating your Slack Status using a preset list of statuses.

Please note that this is still WIP. I'm labelling it as an alpha pre release. But I don't want to wait without shipping, so whatever I have will be whatever goes.

In the current iteration of the CLI, the CLI can read a Slack Token and set a status based off a predefined set of statuses that you keep in a text file. The process of using it is currently improving and isn't as involved as it was when it started but it still needs some manual text editor work along the way.

Please open up any issues you find on Github. I'm actively developing the project and committing to keep it improving.

You can follow transparent development of the project [on my little engineering log site](http://eng.adnanissadeen.com/projects/slack-status-updater-cli.html)

## How to use it

### Setup steps:
* Download the binary for your relevant OS from [here](https://github.com/kiriappeee/slack-status-updater-cli/releases). Currently available only for 64 bit Mac OS and Linux

* Get your personal Slack token from [the legacy tokens page](https://api.slack.com/custom-integrations/legacy-tokens)

* Move your binary to somewhere in your `$PATH`. I personally keep mine in `/usr/local/bin/slackstatus` which I've changed ownership of to myself to avoid using `sudo`. Run `slackstatus -h` to see if things work correctly.

If you are running this for the first time you'll be prompted to paste in your slack personal token. The token will be recorded in a newly created file in `~/.config/ssucli/tokenconfig`. At the same time, some default statuses will be created for you in `~/.config/ssucli/statuses.yaml`. You can manually edit these files at any point. 

Once the above is complete, you should see this:

```
NAME:
   Slack Status Updater - Painless status updating via your Terminal

USAGE:
   slack-status-updater-cli [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
     set      Sets a status from a predefined list of statuses
     list     Lists available statuses and their information
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value  Slack Legacy/API token to use with the CLI [/root/.config/ssucli/tokenconfig]
   --help, -h     show help
   --version, -v  print the version

```

* Check the statuses that have already been added for you by running `slackstatus list`. That should give you this:

```
NAME             EMOJI                   STATUS TEXT 
lunch            bowl_with_spoon         Having lunch 
resting          bath                    Resting 
deepwork         hammer_and_wrench       In Deep work mode 
driving          car                     On the road 
ping             bell                    Working and available for pings :) 
gym              weight_lifter           Out getting exercise at the gym! 
meeting                                  In a meeting 
happy            grin                     

```

_Note that the last two statuses don't have all the attributes filled in. The statuses spec allows either emoji or statusText to be empty **but not both**._

* You can now run `slackstatus set ping` to test changing your status. If it works correctly you should see:

```
Status post result: Status set for ping
```

* The effect of the above will be to set your status emoji to bell and your status message to "Working and available for pings :)". 
* You can alternatively use `slackstatus ping` to achieve the same thing.
* Check out `~/.config/ssucli/statuses.yaml` to see how the statuses exist in the yaml file.

## Adding new statuses.

To add a new status, simply open `$HOME/.config/ssucli/statuses.yaml` and add a new status in the following format:

```yaml
- statusName: name-of-the-status-you-use-when-calling-from-terminal
  emoji: the_emoji_you_want
  statusText: The status message you want to associate with this.
```

That's it :). Remember, either `emoji` or `statusText` can be blank, but not both.

## Next steps:

The immediate next steps I have in mind center around CRUD operations on statuses.
