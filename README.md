# Aegis Lokr

Aplicativo desktop para organizar senhas localmente, com foco em privacidade e controle dos dados. O projeto está em desenvolvimento: a versão atual ainda é um scaffold do Wails e **não implementa armazenamento de senhas nem criptografia**.

## Status

| Recurso | Situação |
| --- | --- |
| Aplicativo desktop com Wails | Scaffold inicial |
| Interface em React e TypeScript | Tela inicial |
| Cofre local de senhas | Planejado |
| Grupos e subgrupos | Planejado |
| Criptografia e senha mestra | Planejado |
| Importação e exportação CSV | Planejado |
| Histórico de alterações e arquivamento | Planejado |

As funcionalidades planejadas incluem armazenamento local, organização hierárquica de credenciais, importação/exportação CSV, histórico de alterações e arquivamento. A sincronização com a nuvem está fora do escopo da primeira versão.

> **Atenção:** não use a versão atual para guardar credenciais reais. Ela ainda não oferece um cofre seguro.

## Tecnologias

- **Desktop:** Wails v3
- **Backend:** Go
- **Interface:** React, TypeScript e Vite
- **Estilos:** Tailwind CSS

## Requisitos para desenvolvimento

- Go 1.26 ou superior
- Node.js e npm
- [Task](https://taskfile.dev/installation/)
- CLI do [Wails v3](https://v3.wails.io/)
- Dependências da WebView exigidas pelo Wails para o seu sistema operacional

Consulte a documentação do Wails para instruções de instalação da CLI e dos requisitos específicos de cada plataforma.

## Executar em desenvolvimento

Na raiz do repositório, instale os pré-requisitos acima e execute:

```sh
task dev
```

O comando inicia o frontend e o aplicativo desktop em modo de desenvolvimento.

## Compilar

Para gerar uma build de produção para o sistema operacional atual:

```sh
task build
```

Para empacotar a aplicação:

```sh
task package
```

Os comandos de build e empacotamento dependem das ferramentas e bibliotecas exigidas pelo Wails na plataforma de destino. Consulte também as tarefas disponíveis em [`Taskfile.yml`](./Taskfile.yml).

## Estrutura do projeto

```text
.
├── build/                 # Configurações e recursos de build por plataforma
├── frontend/              # Interface React, TypeScript e Vite
├── main.go                # Inicialização do aplicativo Wails
├── greetservice.go        # Serviço de exemplo do scaffold
├── go.mod                 # Dependências Go
└── Taskfile.yml           # Tarefas de desenvolvimento, build e empacotamento
```

## Segurança e privacidade

A meta do projeto é manter os dados localmente e funcionar offline. Esses são objetivos da versão planejada, não garantias da implementação atual. Até que o armazenamento protegido e a criptografia sejam implementados e validados, não armazene senhas neste aplicativo.
