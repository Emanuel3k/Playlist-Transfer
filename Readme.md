# playlist-transfer

O **playlist-transfer** é uma aplicação que permite transferir playlists entre diferentes plataformas de música, como YouTube, Spotify, Apple Music, entre outras. O objetivo é facilitar a migração e o compartilhamento de playlists, tornando o processo simples e rápido para o usuário.

---

## Funcionalidades

- Criação de conta e autenticação de usuários
- Integração inicial com o Spotify
- Estrutura preparada para integração com outras plataformas (YouTube, Apple Music, etc.)

---

## Tecnologias

- Go (Golang)
- CHI (framework web)
- PostgreSQL (banco de dados)
- Docker (containerização)
- Makefile (automação de tarefas)
- Migrate (migração de banco de dados)
- air (recarregamento automático em desenvolvimento)
- Mockery e testify (testes)
- REST API
- Integração com serviços de música via API

---

## Como rodar o projeto

1. **Clone o repositório:**
    ```bash
    git clone https://github.com/Emanuel3k/Playlist-Transfer
    cd playlist-transfer
    ```

2. **Instale as dependências:**
    ```bash
    go mod tidy
    ```

3. **Configure as variáveis de ambiente** (exemplo: credenciais do Spotify).

4. **Inicie o Docker:**
    ```bash
    docker compose up -d
    ```
   > Certifique-se de que o Docker está instalado e configurado corretamente em sua máquina.

5. **Execute as migrações do banco de dados:**
    ```bash
    make migration-up
    ```

6. **Inicie a aplicação:**
    ```bash
    air
    ```

---

## Rotas já implementadas

---

## Rotas a implementar

---

> *Este projeto está em desenvolvimento e novas funcionalidades serão adicionadas em breve.*