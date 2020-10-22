# Upload Travis Artifacts to Telegram

#### 1) Define some ENV variables
```
TG_BOT_TOKEN=<your token>
TG_BOT_USERS=<target chat ids separated with commas>
```

#### 2) Use in your CI pipeline
```
curl https://github.com/sokket/tgTravisArtifactUploader/releases/download/1.0/tcn > ./tcn
chmod +x ./tcn && ./tcn <file to upload>
```
