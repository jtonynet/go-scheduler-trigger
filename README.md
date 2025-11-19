<a id="header"></a>


<center>
<a href="#header">
    <img src="./docs/assets/images/layout/header.png" alt="gopher azuis torcedores" />
</a>
</center>

<!-- 
    icons by:
    https://devicon.dev/
    https://simpleicons.org/
-->
[<img src="./docs/assets/images/icons/go.svg" width="25px" height="25px" alt="Go Logo" title="Go">](https://go.dev/) [<img src="./docs/assets/images/icons/gin.svg" width="25px" height="25px" alt="Gin Logo" title="Gin">](https://gin-gonic.com/) [<img src="./docs/assets/images/icons/redis.svg" width="25px" height="25px" alt="Redis Logo" title="Redis">](https://redis.io/) [<img src="./docs/assets/images/icons/docker.svg" width="25px" height="25px" alt="Docker Logo" title="Docker">](https://www.docker.com/) [<img src="./docs/assets/images/icons/ubuntu.svg" width="25px" height="25px Logo" title="Ubuntu" alt="Ubuntu" />](https://ubuntu.com/) [<img src="./docs/assets/images/icons/dotenv.svg" width="25px" height="25px" alt="Viper DotEnv Logo" title="Viper DotEnv">](https://github.com/spf13/viper) [<img src="./docs/assets/images/icons/github.svg" width="25px" height="25px" alt="GitHub Logo" title="GitHub">](https://github.com/jtonynet) [<img src="./docs/assets/images/icons/visualstudiocode.svg" width="25px" height="25px" alt="VsCode Logo" title="VsCode">](https://code.visualstudio.com/) [<img src="./docs/assets/images/icons/mermaidjs.svg" width="25px" height="25px" alt="MermaidJS Logo" title="MermaidJS">](https://mermaid.js.org/) [<img src="./docs/assets/images/icons/rabbitmq.svg" width="25px" height="25px" alt="RabbitMQ Logo" title="RabbitMQ">](https://rabbitmq.com/) [<img src="./docs/assets/images/icons/mailhog.png" width="40px" height="30px" alt="MailHog Logo" title="MailHog">](https://github.com/mailhog/MailHog)

<!-- 
[<img src="./docs/assets/images/icons/swagger.svg" width="25px" height="25px" alt="Swagger Logo" title="Swagger">](https://swagger.io/) [<img src="./docs/assets/images/icons/githubactions.svg" width="25px" height="25px" alt="GithubActions Logo" title="GithubActions">](https://docs.github.com/en/actions) 

[<img src="./docs/assets/images/icons/miro.svg" width="25px" height="25px" alt="Miro Logo" title="Miro">](https://https://miro.com/)
-->

 <!--[![Badge GitHubActions](https://github.com/jtonynet/go-pique-nique/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/jtonynet/go-pique-nique/actions) --> 

[![Go Version](https://img.shields.io/badge/GO-1.23.2-blue?logo=go&logoColor=white)](https://go.dev/)

## 🕸️ Redes

[![linkedin](https://img.shields.io/badge/Linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/jos%C3%A9-r-99896a39/) [![gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:learningenuity@gmail.com)

---

## 📁 O Projeto

<a id="index"></a>
### ⤴️ Índice


__[Go Scheduler Trigger](#header)__<br/>
  1.  ⤴️ [Índice](#index)
  2.  📖 [Sobre](#about)
  3.  💻 [Rodando o Projeto](#run)
      - 🌐 [Ambiente](#environment)
      - 🐋 [Conteinerizado](#run-containerized)
      - ✍️ [Endpoints e Uso](#run-use)
  4.  🔢 [Versões](#versions)
  5.  📊 [Diagramas](#diagrams)
      - 📈 [ER](#diagrams-erchart)
  6.  🤖 [Uso de IA](#ia)
  7.  🏁 [Conclusão](#conclusion)

<hr/>

<a id="about"></a>
### 📖 Sobre

`go-scheduler-trigger` é um estudo para agendador reativo desenvolvido em `Go`, projetado para disparar mensagens e notificações quase em tempo real, sem depender de cron jobs ou polling contínuo.

O projeto resolve o problema de sistemas que precisam enviar alertas, emails ou executar tarefas temporizadas com precisão e eficiência, eliminando a complexidade e o overhead de soluções tradicionais baseadas em agendamento recorrente.

Ao utilizar uma abordagem baseada em eventos, ele garante que cada mensagem seja acionada exatamente no momento previsto, proporcionando uma solução simples, confiável e altamente escalável para notificações e tarefas programadas.


<br/>

[⤴️ de volta ao índice](#index)

---

<a id="run"></a>
### 💻 Rodando o Projeto

<a id="environment"></a>
#### 🌐 Ambiente

`Docker` e `Docker Compose` são necessários para rodar a aplicação de forma containerizada, e é fortemente recomendado utilizá-los para rodar o banco de dados e demais dependências localmente. Siga as instruções abaixo caso não tenha esses softwares instalados em sua máquina:

- &nbsp;<img src='./docs/assets/images/icons/docker.svg' width='13' alt='Github do' title='Github do'>&nbsp;[Instalando Docker](https://docs.docker.com/engine/install/)
- &nbsp;<img src='./docs/assets/images/icons/docker.svg' width='13' alt='Github do' title='Github do'>&nbsp;[Instalando Docker Compose](https://docs.docker.com/compose/install/)

<br/>
<div align="center">. . . . . . . . . . . . . . . . . . . . . . . . . . . .</div>
<br/>

<a id="run-containerized"></a>
#### 🐋 Containerizado 

<br/>
<div align="center">. . . . . . . . . . . . . . . . . . . . . . . . . . . .</div>
<br/>


<a id="run-use"></a>
#### ✍️ Endpoints e Uso


<br/>

[⤴️ de volta ao índice](#index)

---

<a id="diagrams"></a>
### 📊 Diagramas

<a id="diagrams-erchart"></a>
#### 📈 ER

**TODO**

<div align="center">


</div>

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="versions"></a>
### 🔢 Versões

As tags de versões estão sendo criadas manualmente a medida que o projeto avança. Cada tarefa é desenvolvida em uma branch a parte (Branch Based, [feature branch](https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow)) e quando finalizadas é gerada tag e mergeadas em main.

Para obter mais informações, consulte o [Histórico de Versões](./CHANGELOG.md).

<br/>

[⤴️ de volta ao Index](#index)

---

<a id="ia"></a>
### 🤖 Uso de IA

A figura do cabeçalho nesta página foi criada com a ajuda de inteligência artificial e um mínimo de retoques e construção no Gimp [<img src="./docs/assets/images/icons/gimp.svg" width="30" height="30 " title="Gimp" alt="Gimp Logo" />](https://www.gimp.org/)

__Os seguintes prompts foram usados para criação no  [Bing IA:](https://www.bing.com/images/create/)__

<details>
  <summary><b>Gopher ocupado com agendamentos</b></summary>
"gopher azul, simbolo da linguagem golang o mais proximo possivel do mascote, olhando para um relogio de pulso na sua mao e na outra um calendario de compromissos, estilo cartoon, historia em quadrinhos, fundo branco chapado para facilitar remoção"<b>(sic)</b>
</details>

<br/>

IA também é utilizada em minhas pesquisas e estudos como ferramenta de apoio; no entanto,  __artes e desenvolvimento são, acima de tudo, atividades criativas humanas. Valorize as pessoas!__

Contrate artistas para projetos comerciais ou mais elaborados e aprenda a ser engenhoso!

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="conclusion"></a>
### 🏁 Conclusão

**TODO**

<br/>

Este desafio me permite consolidar conhecimentos e identificar pontos cegos para aprimoramento. Continuarei trabalhando para evoluir o projeto e expandir minhas habilidades.

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="footer"></a>

<br/>

>  _"Lifelong Learning & Prosper"_
> <br/> 
>  _Mr. Spock, maybe_   🖖🏾🚀

<div align="center">
<a href="#footer">
<img src="./docs/assets/images/layout/footer.png" />
</a>
</div>
