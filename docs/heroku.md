## heroku helpers

#### logs
`heroku logs --tail -a piggy-police`

#### postgres
[heroku psql cli docs](https://devcenter.heroku.com/articles/heroku-postgresql#using-the-cli)

[local postgres setup](https://devcenter.heroku.com/articles/heroku-postgresql#local-setup)

info
`watch heroku pg:info -a piggy-police`

psql
`heroku pg:psql -a piggy-police`

creds
`heroku pg:credentials:url DATABASE -a piggy-police`

rotate creds
`heroku pg:credentials:rotate DATABASE -a piggy-police`

drop & recreate db
`heroku pg:reset DATABASE -a piggy-police`