# Explicações técnicas

## Arquitetura

Relação de dependências:

```txt
                 +---+         +--------+
      +--------->|gin| +------>|platform+------+
      |          +-+-+ |       +---+----+      |
   +--+-+          |^  |           |           |
   |main|          ||  |           |           |
   +--+-+          v|  |           v           v
      |        +----+--+-+     +-------+    +------+
      +------->|ginrouter+---->|usecase+--->|domain|
               +----+----+     +-------+    +------+
                    |              ^           ^
                    |              |           |
                    |            +-+-+         |
                    +----------->|web+---------+
                                 +---+
```

- `main` é a execução da aplicação
- `gin` biblioteca externa
- `ginrouter` implementação do `gin`
- `web` interpretação das requisições e respostas http com a biblioteca padrão do go
- `platform` implementação das interfaces de `usecase`
- `usecase` struct com método específico da regra de negócio e suas dependências como interfaces
- `domain` structs e seus comportamentos que refletem entidades das regras de negócio

### Conclusões

- **Prós**:
    - `domain` e `usecase` não dependem de implementações de `platform`
    - `web` não depende do router utilizado em `ginrouter`
- **Contras**:

## Documentação

### Web API

Utilizando o [swaggo](https://github.com/swaggo/swag) "implemento" somente em `ginrouter`. Evitando "sujar" `web`.

<!-- ### Package API-->

<!-- ## Testes -->

## Versionamento de UI vs API

Normalmente vai fazer mais sentido o projeto em Go não conhecer o projeto (arquivos fontes, bibliotecas e etc) da SPA, e portanto serem projetos separados. Nessa poc ambos estão juntos por praticidade.

A separação pode ser feita de duas formas:

- Adicionar no `.gitignore` a pasta `web/app` e codar ali dentro utilizando o backend real **(recomendado)**
- Ou criar o projeto em qualquer outro diretório fora do projeto e utilizar um servidor falso

## Referências

- https://golang.org/doc/effective_go
- https://blog.golang.org/package-names
- https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
- https://github.com/golang-standards/project-layout
- https://github.com/eminetto/pos-web-go
- https://www.youtube.com/watch?v=LOn1GUsjOF4
- https://www.alexedwards.net/blog/organising-database-access
