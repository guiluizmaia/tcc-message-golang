Para rodar o projeto deve-se ter docker e docker-compose

É possível fazer dois testes

Comum (Teste em que o user e o server estejam no ar em todo momento):
    Comentar (Caso esteja descomentado) seguinte código no arquivo docker-compose.yaml:
    # depends_on:
    #   - user

    Descomentar (Caso esteja comentado):
    depends_on:
      - server

    Comando 
    - docker-compose up --build

    O container de User irá printar algo parecido com isso
    2022/07/12 22:26:34 Amount of requests: 70000
    2022/07/12 22:26:34 Bytes of one request: 376
    2022/07/12 22:26:34 Time of all requests: 22.089965705s
    2022/07/12 22:26:34 Errors: 0
    2022/07/12 22:26:34 Success: 70000

Resiliência (Teste em que o user inicia a fazer requisições antes do server estar no ar):
    Comentar (Caso esteja descomentado) seguinte código no arquivo docker-compose.yaml:
    depends_on:
      - user

    Descomentar (Caso esteja comentado):
    # depends_on:
    #   - server

    Comando 
    - docker-compose up --build

    O container de User irá printar algo parecido com isso
    2022/07/12 22:26:34 Amount of requests: 70000
    2022/07/12 22:26:34 Bytes of one request: 376
    2022/07/12 22:26:34 Time of all requests: 22.089965705s
    2022/07/12 22:26:34 Errors: 388
    2022/07/12 22:26:34 Success: 69612

Amount of requests = Quantidade de requisições feitas
Bytes of one request = Tamanho em bytes de uma requicição
Time of all requests = Tempo para ter sido feitas todas as requisições
Errors = Requisições que deram erros
Success = Requisições que deram sucesso