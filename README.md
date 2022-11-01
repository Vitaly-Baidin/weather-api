# Run app
```bash
make compose-up
```

# Example requests
| request                                                  | info                                   |
|----------------------------------------------------------|----------------------------------------|
| localhost:8080/docs                                      | swagger doc                            |
| localhost:8080/city/{country}/{city}/summary             | return summary info about city         |
| localhost:8080/city/{country}/{name}/weather/{timestamp} | return full info about weather in city |

```bash
    curl localhost:8080/city/ru/bratsk/summary
    curl localhost:8080/city/ru/bratsk/weather/1667692800    
```