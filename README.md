# trafilea
API that categorizes and stores numbers based on their divisibility by 3 and 5


- GET Number
> curl --location 'localhost:8080/api/trafilea/number/75'

- GET ALL Numbers
> curl --location 'localhost:8080/api/trafilea/number/all'
 
- ADD Numbers
> curl --location --request POST 'localhost:8080/api/trafilea/number/10'
