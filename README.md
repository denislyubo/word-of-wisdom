## What kind of proof of work was chosen?

It was chosen algorithm based on hashing random string obtained from server. The client task is to add nonce number to server
string until sha256 of random string and nonce concatenation will not contain required number of zeros at the beginning of hash.
Server timeout is 15s. In case client finds nonce for given string earlier it will receive quote in response.

## How to run client and server in docker?

Use cli command: make ```make run-app``` or ```docker compose up```

## Какой используется алгоритм proof of work?

Был выбран алгоритм на основе хэширования строки, генерируемой сервером случайно. Задача клиента, используя алгоритм 
хэширования sha256, добавлять к строке последовательный номер nonce пока хэш от заданной конструкции не будет содержать 
заданное количество нулей в начале.
Таймаут сервера 15 сек, только, если клиент находит nonce раньше, он получает quote в ответе.


## Как запустить клиент и сервер?

С помощью команды make ```make run-app``` или ```docker compose up```