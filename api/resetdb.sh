psql postgres://fyparbfgmnlvnt:"${DB_PASS}"@ec2-174-129-253-162.compute-1.amazonaws.com:5432/d93afugd4i967s << EOF
  DROP SCHEMA public CASCADE;
  CREATE SCHEMA public;
  GRANT ALL ON SCHEMA public TO postgres;
  GRANT ALL ON SCHEMA public TO public;
  GRANT ALL ON SCHEMA public TO fyparbfgmnlvnt;
EOF

echo 'deleting migrations'
rm -rf migrations
echo 'initializing migrations'
flask db init
echo 'generating migrations'
flask db migrate
echo 'applying migrations'
flask db upgrade

  # insert into team (id, name) values (1, 'james-account');
  # insert into team (id, name) values (2, 'jon-account');