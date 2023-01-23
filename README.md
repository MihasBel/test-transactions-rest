# test-transactions-rest
A test task for the implementation of a transaction system. It accepts requests to deposit or withdraw money and enters them into the database. It also gives transaction data from the database when requesting information on any transaction.
## [main repo](https://github.com/MihasBel/test-transactions)
## Architecture schema in [miro](https://miro.com/app/board/uXjVPvwzV6U=/?share_link_id=797824750985)
## REST part 

This part of test-transactions task responsible for consuming HTTP REST requests. According to task uses "schema first" approach. Created Restful API according to docs 
[docs folder](https://github.com/MihasBel/test-transactions-rest/tree/main/api/docs)

- API accepts requests to create transaction and send it to kafka
- API accepts requests to get info about created transaction
- For speed up requests and dont miss transaction uses Redis as cache