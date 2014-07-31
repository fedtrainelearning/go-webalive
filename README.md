First attempt at a Go application:
 - Call urls concurrently, getting the status code and ensuring the sites don't go to sleep (to
   be called from a cron job)
 - Using YAML for configuration (to store the urls)

Config:

    urls:
        - http://someurl.com/example
        - http://anotherurl.com/another

Call with:

    webalive --file=<path_to_yaml_config_file>

Future plans:
 - Postgres integration (to get the urls from database)