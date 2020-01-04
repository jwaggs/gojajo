import os
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate

from sqlalchemy import Table, Column, Integer, ForeignKey
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = f'postgres://fyparbfgmnlvnt:{os.getenv("DB_PASS")}@ec2-174-129-253-162.compute-1.amazonaws.com:5432/d93afugd4i967s'

db = SQLAlchemy(app)
migrate = Migrate(app, db)


class Team(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(128))
    

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(128))


@app.route('/')
def index():
    return 'Index Page'


@app.route('/transactions')
def get_transactions():
    x = [
        'transaction 1',
        'transaction 2'
    ]
    return x
