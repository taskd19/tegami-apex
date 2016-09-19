# tegami-apex

Tegami-apex is a program to run Tegami on AWS Lambda.

## Usage
1. Set up [APEX](https://github.com/apex/apex)
2. Run command `apex init`.
3. Remove a sample function from functions directory.
4. Add environment variables for Tegami to "environment" in project.json.
```json
{
   "environment": {
    "TEGAMI_TIMEFRAME": "TIMEFRAME_TO_CHECK_EVENTS",
    "TEGAMI_TW_CONSUMER_KEY": "TWITTER_CONSUMER_KEY",
    "TEGAMI_TW_CONSUMER_SECRET": "TWITTER_CONSUMER_KEY_SECRET",
    "TEGAMI_TW_ACCESS_TOKEN": "TWITTER_ACCESS_TOKEN",
    "TEGAMI_TW_ACCESS_TOKEN_SECRET": "TWITTER_ACCESS_TOKEN_SECRET"
  }
}
```
5. Run command `apex deploy`.

## License
MIT
