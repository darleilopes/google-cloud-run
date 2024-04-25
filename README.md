Temperatura por CEP

# Levantando o serviço

- Execute o comando: `make prepare`
- Ele gera dois arquivos: `env.dev.json` em `env.prod.json`, em um real cenário, poderia usar diferentes valores

- No caso vamos testar prod `make run-prod`

# Testando localmente

Execute: `curl -s http://localhost:8080/temperature?cep=05025-000`

# Deploy GCP Cloud Run

- Configurar as seguintes variaveis:
    - `ENV_PROJECT_ID`: seu projectID
    - `ENV_ARG`: pode ser `prod` ou `dev`.
- Locar usando `gcloud auth login`
- executar `gcloud config set project ${ENV_PROJECT_ID}`
- depois `deploy.sh` como o exemplo abaixo:
  ```shell
  ENV_PROJECT_ID=generic-3232322 ENV_ARG=prod ./deploy.sh
  ```
- For this project, we have this URL to test using `GET` method
  ```shell
  curl https://temperature-uz6atyp8yq-uc.a.run.app/temperature?cep=05025-000 | jq
  ```

# Para executar os tests
```shell
API_URL=http://localhost:8080 go test -v ./tests/integration
```