# GOcalculatorAPI
hosted on [render.com](https://render.com/)
## Endpoints:
- ***/add***: takes as arguments `x` and `y` two numbers and returns the result of `x+y`
- ***/subtract***: takes as arguments `x` and `y` two numbers and returns the result of `x-y`
- ***/multiply***: takes as arguments `x` and `y` two numbers and returns the result of `x*y`
- ***/divide***: takes as arguments `x` and `y` two numbers and returns the result of `x/y`
- ***/calculate***: takes one argument `expression` that represents a mathematical expression and returns the result
  - example: expression=`4*2+2/2` returns `9`
## Note:
Mathematical operators: `+, -, *, /` must be URL encoded as follows:
| Operator | URL Encoded |
| :------: | :---------: |
| +        | %2B         |
| -        | %2D         |
| *        | %2A         |
| /        | %2F         |
