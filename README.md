First attempt at a Go application:
 - Call urls, getting the status code and ensuring the sites don't go to sleep (to
   be called from a cron job)

 Future plans
 - Modify the application to concurrently call the urls
 - Postgres integration (to get the urls from database)