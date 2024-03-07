# GLOBAL HITSS - TESTE DE CONHECIMENTO TÉCNICO

Objetivo:
O objetivo desta tarefa é avaliar a capacidade do programador em desenvolver uma API
RESTful CRUD em Golang, utilizando um banco de dados Postgres e a RabbitMQ para
processamento assíncrono de inserções de registros relacionados a dados pessoais de clientes.

Tarefa:
Desenvolva uma API RESTful CRUD em Golang que seja capaz de receber requisições HTTP e
realizar operações CRUD em um banco de dados Postgres. Os registros a serem inseridos na
base de dados serão enviados para uma fila na RabbitMQ e posteriormente processados por
um worker. Os dados são relacionados a dados pessoais de clientes e devem seguir boas
práticas de privacidade.
Os dados dos clientes a serem manipulados são: nome, sobrenome, contato, endereço, data
de nascimento e CPF.

Requisitos:
 A API deve permitir a criação, leitura, atualização e remoção de registros no banco de
dados Postgres.
 As inserções de registros devem ser enviadas para uma fila na RabbitMQ.
 Um worker deve estar ouvindo a fila na RabbitMQ e inserindo os registros no banco de
dados Postgres.
 O código deve seguir boas práticas de programação em Golang e seguir o estilo de
código definido pela comunidade.
 Deve ser fornecida uma documentação detalhada das APIs, descrevendo cada
endpoint e suas funcionalidades.
 O código deve ser hospedado em um repositório Git público e compartilhado com o
avaliador.
 Deve ser fornecido um arquivo README explicando como executar a aplicação e testar
as APIs.

 Deve ser fornecido um arquivo com as especificações das mensagens da fila na
RabbitMQ.
 Os dados pessoais de clientes devem ser protegidos seguindo as boas práticas de
privacidade.
 A API deve ser escalável e preparada para lidar com um grande volume de dados e
requisições.

Critérios de Avaliação:
 Qualidade do código: clareza, modularidade, organização, padronização e utilização
adequada das ferramentas e bibliotecas.
 Conhecimento do Golang: uso correto das estruturas de dados e recursos da
linguagem.
 Conhecimento do Postgres: conhecimento dos conceitos avançados de banco de
dados e habilidade em otimização de consultas e gerenciamento de transações.
 Conhecimento de RabbitMQ: compreensão avançada de como a fila funciona e
habilidade em utilizar técnicas de otimização para processamento assíncrono eficiente.
 Conhecimento de privacidade: capacidade de proteger dados pessoais de clientes
seguindo as boas práticas de privacidade.
 Escalabilidade: habilidade em projetar e implementar uma API escalável capaz de lidar
com grande volume de dados e requisições.

Nota: Esta tarefa é uma avaliação de habilidades e não um tutorial. Não será fornecido
nenhum material de estudo ou suporte técnico. O programador deve utilizar seus
conhecimentos e pesquisar para completar a tarefa.