# stash-hook-router
This is a quick hack to work around some of the shortcomings of Stash and the existing plugins.
This project adds basic plugin functionality using [pull-request-notifier-for-stash](https://github.com/tomasbjerre/pull-request-notifier-for-stash). stash-hook-router will run the script in `STASH_SCRIPT_DIR` matching the project name _ repo name

Example:

- Project: Foobar
- Repository name: Buzz
- Script directory: /var/lib/scripts
- Result: `/var/lib/scripts/foobar_buzz <json with all the data>`


### Installation

Configure the plugin to post to your URL running stash-hook-router. ADD this to `POST content:`
```{ "PULL_REQUEST_ID": "${PULL_REQUEST_ID}", "PULL_REQUEST_ACTION": "${PULL_REQUEST_ACTION}", "PULL_REQUEST_VERSION": "${PULL_REQUEST_VERSION}", "PULL_REQUEST_COMMENT_TEXT": "${PULL_REQUEST_COMMENT_TEXT}", "PULL_REQUEST_USER_DISPLAY_NAME": "${PULL_REQUEST_USER_DISPLAY_NAME}", "PULL_REQUEST_USER_EMAIL_ADDRESS": "${PULL_REQUEST_USER_EMAIL_ADDRESS}", "PULL_REQUEST_USER_ID": "${PULL_REQUEST_USER_ID}", "PULL_REQUEST_USER_NAME": "${PULL_REQUEST_USER_NAME}", "PULL_REQUEST_USER_SLUG": "${PULL_REQUEST_USER_SLUG}", "PULL_REQUEST_AUTHOR_DISPLAY_NAME": "${PULL_REQUEST_AUTHOR_DISPLAY_NAME}", "PULL_REQUEST_AUTHOR_EMAIL": "${PULL_REQUEST_AUTHOR_EMAIL}", "PULL_REQUEST_AUTHOR_ID": "${PULL_REQUEST_AUTHOR_ID}", "PULL_REQUEST_AUTHOR_NAME": "${PULL_REQUEST_AUTHOR_NAME}", "PULL_REQUEST_AUTHOR_SLUG": "${PULL_REQUEST_AUTHOR_SLUG}", "PULL_REQUEST_FROM_SSH_CLONE_URL": "${PULL_REQUEST_FROM_SSH_CLONE_URL}", "PULL_REQUEST_FROM_HTTP_CLONE_URL": "${PULL_REQUEST_FROM_HTTP_CLONE_URL}", "PULL_REQUEST_FROM_HASH": "${PULL_REQUEST_FROM_HASH}", "PULL_REQUEST_FROM_ID": "${PULL_REQUEST_FROM_ID}", "PULL_REQUEST_FROM_BRANCH": "${PULL_REQUEST_FROM_BRANCH}", "PULL_REQUEST_FROM_REPO_ID": "${PULL_REQUEST_FROM_REPO_ID}", "PULL_REQUEST_FROM_REPO_NAME": "${PULL_REQUEST_FROM_REPO_NAME}", "PULL_REQUEST_FROM_REPO_PROJECT_ID": "${PULL_REQUEST_FROM_REPO_PROJECT_ID}", "PULL_REQUEST_FROM_REPO_PROJECT_KEY": "${PULL_REQUEST_FROM_REPO_PROJECT_KEY}", "PULL_REQUEST_FROM_REPO_SLUG": "${PULL_REQUEST_FROM_REPO_SLUG}", "PULL_REQUEST_TO_SSH_CLONE_URL": "${PULL_REQUEST_TO_SSH_CLONE_URL}", "PULL_REQUEST_TO_HTTP_CLONE_URL": "${PULL_REQUEST_TO_HTTP_CLONE_URL}", "PULL_REQUEST_TO_HASH": "${PULL_REQUEST_TO_HASH}", "PULL_REQUEST_TO_ID": "${PULL_REQUEST_TO_ID}", "PULL_REQUEST_TO_BRANCH": "${PULL_REQUEST_TO_BRANCH}", "PULL_REQUEST_TO_REPO_ID": "${PULL_REQUEST_TO_REPO_ID}", "PULL_REQUEST_TO_REPO_NAME": "${PULL_REQUEST_TO_REPO_NAME}", "PULL_REQUEST_TO_REPO_PROJECT_ID": "${PULL_REQUEST_TO_REPO_PROJECT_ID}", "PULL_REQUEST_TO_REPO_PROJECT_KEY": "${PULL_REQUEST_TO_REPO_PROJECT_KEY}", "PULL_REQUEST_TO_REPO_SLUG": "${PULL_REQUEST_TO_REPO_SLUG}" }```

Handle the JSON data as you like in your script(s)