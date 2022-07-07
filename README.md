# Slack Age Bot

A simple Slack Bot that calculates age, given the year of birth (yob).

## How to Use
1. Add the Slack Bot to your channel

2. Type ```@age-bot my dob is <YYYY-MM-DD>```

## Setup
1. Clone the repo ```git clone <repo>```

2. Install dependencies ```go get "github.com/joho/godotenv"``` and ```go get "github.com/shomali11/slacker"``` in root directory

3. Visit https://api.slack.com/apps to configure Slack Bot 

    - Enable Socket Mode and obtain ```SLACK_APP_TOKEN```

    - Install app to Slack Workspace to obtain ```SLACK_BOT_TOKEN```
    
Your ```.env``` file should be formatted as follows:

```
SLACK_BOT_TOKEN=
SLACK_APP_TOKEN=
```

4. Run the commands ```go build``` and then ```source .env && go run main.go```
