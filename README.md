# 🐹 golang-studies-2026

Meu caderno de estudos de Go — saindo do .NET e mergulhando de cabeça no ecossistema Golang.

A ideia não é apenas aprender sintaxe: quero construir uma base sólida de **Go idiomático, concorrência, testes, arquitetura e sistemas backend em produção**.

## 🧭 Estrutura

O código atual continua simples e experimental. Conforme os estudos evoluem, a organização segue estes eixos:

```text
.
├── .github/              # Instruções e prompts para o estudo com IA
├── docs/                 # Roadmap e documentação de estudos
├── praticas/             # Práticas e exercícios já organizados
├── *.go                  # Exercícios iniciais e experimentos
├── go.mod
└── go.work
```

Os exercícios iniciais permanecem na raiz por enquanto para preservar o histórico e evitar transformar um caderno de estudos em uma arquitetura artificial. A organização por módulos/pastas será feita conforme cada assunto ganhar massa suficiente.

## 📚 Roadmap

Veja o [roadmap de estudos](./docs/roadmap.md).

1. **Fundamentos** — tipos, funções, fluxo, pointers, structs e packages
2. **Data structures** — arrays, slices, maps e interfaces
3. **Errors & testing** — errors, testes, benchmarks e race detector
4. **Backend Go** — HTTP, context, JSON, banco e observabilidade
5. **Concurrency** — goroutines, channels, sincronização e backpressure
6. **Production Go** — profiling, memória, GC e resiliência
7. **Architecture** — packages, dependency inversion e Clean Architecture pragmática

## 🤖 Estudo com IA

Este repositório também funciona como um **Go Study Lab**.

As instruções em `.github/` configuram a IA para agir como uma parceira de estudos: fazer perguntas, testar meu entendimento, propor experimentos e revisar meu raciocínio em vez de simplesmente gerar código.

Modos disponíveis:

- **Study** — aprender um conceito com perguntas e experimentos
- **Quiz** — active recall, uma pergunta por vez
- **Debug** — investigação guiada de bugs
- **Review** — code review pedagógico
- **Interview** — simulação de entrevista técnica

## 🆚 .NET → Go

| Conceito .NET | Equivalente aproximado em Go |
|---|---|
| `class` | `struct` + métodos |
| `interface` | `interface` (implementação implícita) |
| `IEnumerable` | `slice` / `channel` dependendo do caso |
| `async/await` | goroutines + channels / I/O concorrente |
| `try/catch` | retorno explícito de `error` |
| `null` | zero values + ponteiros |
| `namespace` | `package` |
| `NuGet` | `go get` / `go.mod` |

Essas equivalências são apenas um mapa mental inicial, não regras de tradução entre as linguagens.

## 📖 Recursos

- [Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Playground](https://go.dev/play/)
- [Go Documentation](https://go.dev/doc/)

## 🧠 Filosofia

- Um conceito por vez.
- Código pequeno e executável.
- Comentários explicam **por quê**, não apenas **o quê**.
- Comparações com C# quando ajudam a formar o modelo mental.
- Preferência pela standard library.
- Medir antes de otimizar.
- Entender antes de abstrair.
- Registrar erros, dúvidas e experimentos — inclusive os que deram errado.

<p align="center">
  <img src="https://go.dev/images/gophers/motorcycle.svg" width="120" alt="Gopher" />
  <br/>
  <i>Aprender Go entendendo o que está acontecendo por baixo.</i>
</p>
