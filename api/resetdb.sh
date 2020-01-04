psql postgres://fyparbfgmnlvnt:"${DB_PASS}"@ec2-174-129-253-162.compute-1.amazonaws.com:5432/d93afugd4i967s << EOF
  DROP SCHEMA public CASCADE;
  CREATE SCHEMA public;
  GRANT ALL ON SCHEMA public TO postgres;
  GRANT ALL ON SCHEMA public TO public;
  GRANT ALL ON SCHEMA public TO fyparbfgmnlvnt;
EOF

echo 'initializing database'
flask db init
echo 'creating migrations'
flask db migrate
echo 'applying migrations'
flask db upgrade