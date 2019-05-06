## Simple GoLang service template: 


**Start up**

    Simple Template accepts an argument for the path to its config file directory. 
    
**Config file**

    Simple Template uses Viper to read config file
    Config File defaults to YAML but can be changed to JSON.
    
**Logging**

    Simple Template uses Logrus for logging 
    Simple Template looks in the application config file for a 'log_path'.
    No log rotation specified. Use as desired. 
    
**Work**

    Simple Templates' main work entry point is 'DoWork()'
    
**Testing**

    See main_test.go
  
**Building**
```
go build    
```
    
**Usage example:**
```bash
./simple_template .
```
