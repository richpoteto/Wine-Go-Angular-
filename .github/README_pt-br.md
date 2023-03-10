<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_pt-br.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_by.md">
        <img height="20px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/.github/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_tr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/github/v/tag/shpota/goxygen?color=green&label=version">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>

<img src="../templates/vue.webapp/src/assets/logo.svg" align="right" width="230px" alt="goxygen logo">

**Crie um projeto Web com Go e Angular/React/Vue em segundos.**

Goxygen visa poupar seu tempo ao montar um novo projeto. Ele 
cria o esqueleto de uma aplica????o com todas as configura????es 
prontas para voc??. Voc?? pode come??ar a implementar sua l??gica 
de neg??cios imediatamente. Goxygen gera o c??digo back end Go, 
conecta ele com os componentes front end, fornece o Dockerfile
para a aplica????o e cria o arquivo docker-compose para rapidamente
executar em ambientes de desenvolvimento e produ????o.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Supported Technologies</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">Front End</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>Back End</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>Database</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Requirements
Voc?? precisa ter Go 1.11 ou mais recente na sua m??quina.

The `GO111MODULE` environment variable has to be set to `auto`
for the generation logic to work. It is a default for Go
versions up to 1.15. For Go 1.16, you need to set it explicitly:
```
export GO111MODULE=auto
```

## Como utilizar
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
Isso gera um projeto no diret??rio `my-app`.

By default, it will use React and MongoDB. You can select
a different front end framework and a database using
`--frontend` and `--db` flags. For instance, this command
will create a project with Vue and PostgreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

The `--frontend` flag accepts `angular`, `react` and `vue`.
The `--db` flag accepts `mongo`, `mysql` and `postgres`.

O projeto gerado est?? pronto para ser executado com `docker-compose`:
```sh
cd my-app
docker-compose up
```

Ap??s o build ser conclu??do, a aplica????o est?? dispon??vel em 
http://localhost:8080. 

Voc?? pode encontrar mais detalhes de como trabalhar com o projeto
gerado em seu arquivo README.

![Showcase](showcase.gif)

## Estrutura do projeto gerado (React/MongoDB example)

    my-app
    ????????? server                   # Go project files
    ???   ????????? db                   # MongoDB communications
    ???   ????????? model                # domain objects
    ???   ????????? web                  # REST APIs, web server
    ???   ????????? server.go            # O ponto inicial do server
    ???   ????????? go.mod               # depend??ncias do server
    ????????? webapp                    
    ???   ????????? public               # ??cones, arquivos est??ticos e index.html
    ???   ????????? src                       
    ???   ???   ????????? App.js           # O principal componente React
    ???   ???   ????????? App.css          # Estiliza????o do componente App
    ???   ???   ????????? index.js         # Ponto de entrada da aplica????o          
    ???   ???   ????????? index.css        # global styles
    ???   ????????? package.json         # depend??ncias do front end
    ???   ????????? .env.development     # API endpoint para ambiente de desenvolvimento
    ???   ????????? .env.production      # API endpoint para ambiente de produ????o
    ????????? Dockerfile               # faz o build do back end e o front end juntos
    ????????? docker-compose.yml       # deploy em abiente de produ????o
    ????????? docker-compose-dev.yml   # executa um MongoDB local para fins de desenvolvimento
    ????????? init-db.js               # cria um MongoDB collection com dados de teste
    ????????? .dockerignore            # especifica arquivos ignorados no Docker build
    ????????? .gitignore
    ????????? README.md                # guia sobre como usar o reposit??rio gerado

Arquivos como testes de unidade ou sample components n??o est??o inclu??dos 
aqui pela simplicidade.

## Depend??ncias

Goxygen gera uma estrutura b??sica de um projeto e n??o o for??a usar um 
conjunto espec??fico de ferramentas. E por isso n??o traz depend??ncias 
desnecess??rias para o seu projeto. Ele usa apenas
database driver no back end
e [axios](https://github.com/axios/axios) nos projetos React e Vue.
Projetos Angular usam apenas bibliotecas espec??ficas do Angular.

## Como contribuir

Se voc?? encontrou um bug ou tem alguma id??ia de como melhorar o projeto,
[abra uma issue](https://github.com/Shpota/goxygen/issues)
e n??s iremos consertar o mais r??pido poss??vel. Voc?? tamb??m pode enviar
suas altera????es via Pull Request. D?? um fork no reposit??rio, fa??a
altera????es, nos envie um pull request que n??s iremos analisar em breve. 
N??s tamb??m temos um
[Gitter chat](https://gitter.im/goxygen/community) onde discutimos todas
as altera????es.

### Cr??ditos
O logo do Goxygen foi criado por [Egon Elbre](https://twitter.com/egonelbre).
