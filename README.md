# Dashboard para meu homelab

Dashboard para meu seabian muito irado, onde mostra serviços que tenho no meu homelab além de metricas do sistema.

## Dependencias
- [Mattix-agent](https://github.com/Mattix-Monitoring/Mattix-agent)
- [Templ](https://github.com/a-h/templ)
- [HTMX](https://github.com/bigskysoftware/htmx)
- [Air](https://github.com/air-verse/air)

Baixando dependencias:
```bash
go get .
```

Rodando o projeto:
```bash
go tool air
```

# Todo

- [X] Desenhar layout
- [X] Pegar metricas do mattix
- [X] Criar componentes sistema
- [X] Criar serviços (Precisa usar socket do docker)
- [X] Aplicar htmx nos componentes
- [ ] Estilizar

