- go build 
- go build -ldflags "-s -w"
- using env: export GIN_MODE = release 
- using code: gin.SetMode(gin.ReleaseMode)



# Deployment on Heroku 
1. Create a Pipeline by github 
2. Create Procfile
3. Create app in Heroku - https://dashboard.heroku.com/apps 
4. Create app and pipe line 
5. Back to sourcecode & revise server.go by editing port configuration
6. go to heroku and click setting tab and add variable GOVERSION

--- View logs in heroku ---
1. https://devcenter.heroku.com/articles/heroku-cli
2. heroku logs --tail --app stock-golang