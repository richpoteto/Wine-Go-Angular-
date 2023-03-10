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
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
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
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_tr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
    </a>
    <br>
    <div dir="rtl">
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
    </div>
</h1>


<img src="../templates/vue.webapp/src/assets/logo.svg" align="left" width="230px" alt="goxygen logo">


<div dir="rtl">

**?????????? ???????????????? ???? ???????????? Web ?????????????? Go ??-Angular, React ???? Vue.**

???????? Goxygen ?????? ?????????? ?????? ?????????? ???????? ???????????? ??????.
?????? ?????????? ?????? ???????? ?????????? ???? ???? ???????????????????????? ?????????? ????????????.
?????? ???????????? ?????? ???????????? ???????? ???? ?????????????? ?????????? ????????.
Goxygen ????'?????? back-end ???????? Go, ???????? ???????? ???? ?????????????? ???? ??-front-end, ???????? Dockerfile ???????? ???????????? ???????????? ???????? docker-compose ???????????? ???????? ?????????????? ???????????? ?????????????? (production).
</div>

<div dir="rtl">
<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>???????????????????? ????????????</b></td>
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
        <td>???????? ????????????</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>
</div>

<div dir="rtl">

## Requirements
?????????? ???????????? Go 1.11 ?????????? ?????????? ????????.

The `GO111MODULE` environment variable has to be set to `auto`
for the generation logic to work. It is a default for Go
versions up to 1.15. For Go 1.16, you need to set it explicitly:
```
export GO111MODULE=auto
```

## ?????? ????????????
</div>

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
<div dir="rtl">

?????????? ???? ?????????? ???????????? ?????????????? <span dir="ltr">`my-app`</span>.  
</div>

<div dir="rtl">

???????????? ????????, ???????? ?????????? ??-React ??-MongoDB.  
?????? ???????????? ?????????? front-end framework ?????????? ???????????? ?????????? ?????????????? ???????????? <span dir="ltr">`--frontend`</span> ??-<span dir="ltr">`--db`</span> ????????????.  
???????? ?????????? ???? ?????????? ???????????? ?????????? Vue ??-PostgreSQL:
</div>

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

<div dir="rtl">

?????? <span dir="ltr">`--frontend`</span> ???????? ?????????????????? <span dir="ltr">`angular`</span>, <span dir="ltr">`react`</span> ??-<span dir="ltr">`vue`</span>.
?????? <span dir="ltr">`--db`</span> ???????? ?????????????????? <span dir="ltr">`mongo`</span>, <span dir="ltr">`mysql`</span> ??-<span dir="ltr">`postgres`</span>.
</div>

<div dir="rtl">

?????????????? ?????????? ???????? ???????????? ???? <span dir="ltr">`docker-compose`</span>:  
</div>

```sh
cd my-app
docker-compose up
```

<div dir="rtl">

???????? ????-build ????????????, ???????????? ???????? ??-http://localhost:8080.
</div>

<div dir="rtl">

?????? ???????????? ?????????? ?????????? ???????????? ???? ?????? ?????????? ???? ?????????????? ?????????? ?????????? ??-README ??????.  
</div>

![Showcase](showcase.gif)

<div dir="rtl">

## ???????? ???? ???????????? ????'???????? (?????????? React/MongoDB)
</div>


    my-app
    ????????? server                   # Go ???????? ????????????
    ???   ????????? db                   # MongoDB ????????????
    ???   ????????? model                # domain ????????????????
    ???   ????????? web                  # REST APIs, web ??????
    ???   ????????? server.go            # ?????????? ???????????? ???? ????????
    ???   ????????? go.mod               # ???? ???????? dependencies
    ????????? webapp                    
    ???   ????????? public               # index.html-????????????????, ?????????? ???????????? ??
    ???   ????????? src                       
    ???   ???   ????????? App.js           # React ?????????? ???????????? ????
    ???   ???   ????????? App.css          # App ?????????????? ???????????????? ???????? ????????
    ???   ???   ????????? index.js         # ?????????? ???????????? ???? ????????????          
    ???   ???   ????????? index.css        # ?????????????? ????????????????
    ???   ????????? package.json         # front end-???? ?? dependencies
    ???   ????????? .env.development     # dev-???????????? ?? API-?????????? ???????????? ???????? ???? ??
    ???   ????????? .env.production      # prod-???????????? ?? API-?????????? ???????????? ???????? ???? ??
    ????????? Dockerfile               # ???????? front end-???? back end-???? ?? build
    ????????? docker-compose.yml       # prod ?????????? ?????????? ??????????
    ????????? docker-compose-dev.yml   # ?????????? ?????????? ?????????? MongoDB ??????????
    ????????? init-db.js               # ???? ?????????? ?????????? MongoDB-???????? ?????????? ??
    ????????? .dockerignore            # Docker builds ?????????? ?????????? ?????? ???????????? ?????? ??????
    ????????? .gitignore
    ????????? README.md                # ?????????? repo-?????????? ???????? ???????????? ??

<div dir="rtl">

?????????? ?????? ???????? unit tests ???? ?????????? ???????? ???????? ???????????? ?????? ?????????? ??????????.
</div>

<div dir="rtl">

## ???????????? - Dependencies
</div>

<div dir="rtl">

Goxygen ????'?????? ?????? ?????????? ???? ???????????? ?????????? ?????????? ???????? ???????????? ???????????? ???????? ????????????. ??????, ?????? ???? ?????????? ???????????? ???????? ???????????? ?????????????? ????????.
?????? ?????????? ???? ??-driver ???????? ???????? ???????????? ??-back end ????-[axios](https://github.com/axios/axios) ?????????????????? ???????????? React ???? Vue. ?????????????????? ???????????? Angular ?????? ?????????? ???? ?????????????? ???????????????? ???? Angular.
</div>

<div dir="rtl">

## ?????? ??????????
</div>

<div dir="rtl">

???? ?????????? ?????? ???? ?????? ?????? ?????????? ???????? ???????? ???? ?????????????? [open an issue](https://github.com/Shpota/goxygen/issues) ???????????? ???????? ???????? ?????????? ????????????. ?????? ???????????? ???? ?????????? ?????????????? ?????????????? Pull Request. ???????? Fork ??-repository, ???????? ??????????????, ???????? ?????? ???????? pull ???????????? ?????????? ???????? ??????????. ???? ?????? ???? [Gitter chat](https://gitter.im/goxygen/community) ???? ?????? ???????? ?????? ????????????????.
</div>

<div dir="rtl">

### ??????????????
</div>

<div dir="rtl">

?????????? ???? Goxygen ???????? ???? ?????? [Egon Elbre](https://twitter.com/egonelbre).
</div>
