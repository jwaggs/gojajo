import os
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from sqlalchemy.orm import relationship
from flask import Flask, request, jsonify
from flask_marshmallow import Marshmallow

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = f'postgres://fyparbfgmnlvnt:{os.getenv("DB_PASS")}@ec2-174-129-253-162.compute-1.amazonaws.com:5432/d93afugd4i967s'
db = SQLAlchemy(app)
migrate = Migrate(app, db)
ma = Marshmallow(app)


class Team(db.Model):
    """
    The root model that users, banks, and bins relate to
    """
    __tablename__ = 'team'
    id = db.Column(db.Integer, primary_key=True)
    users = relationship('User', back_populates='team')
    # banks = relationship('Bank', back_populates='team')
    # major_bins = relationship('MajorBin', back_populates='team')
    # incomes = relationship('Income', back_populates='team')

    name = db.Column(db.String(128))

    def __repr__(self):
        return f'{self.id}: {self.name}'


class User(db.Model):
    """
    An individual user
    """
    __tablename__ = 'user'
    id = db.Column(db.Integer, primary_key=True)
    team_id = db.Column(db.Integer, db.ForeignKey('team.id'))
    team = relationship("Team", back_populates="users")
    name = db.Column(db.String(128))


# class Bank(db.Model):
#     """
#     A Banking institution or logical grouping of transactional accounts
#     """
#     __tablename__ = 'bank'
#     id = db.Column(db.Integer, primary_key=True)
#     team_id = db.Column(db.Integer, db.ForeignKey('team.id'))
#     bank_accounts = relationship('BankAccount', back_populates='bank')
#
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#
#
# class BankAccount(db.Model):
#     """
#     Encapsulates common data for a Bank's Accounts
#     """
#     __tablename__ = 'bank_account'
#     id = db.Column(db.Integer, primary_key=True)
#     bank_id = db.Column(db.Integer, db.ForeignKey('bank.id'))
#
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#
#
# class Transaction(db.Model):
#     """
#     Encapsulates common data between BankAccount credits and debits
#     """
#     __tablename__ = 'transaction'
#     id = db.Column(db.Integer, primary_key=True)
#     bank_account_id = db.Column(db.Integer, db.ForeignKey('bank_account.id'))
#     minor_bin_id = db.Column(db.Integer, db.ForeignKey('minor_bin.id'))
#     # TODO: add status such as pending/confirmed
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#     date_of_transaction = db.Column(db.DateTime)
#     date_created = db.Column(db.DateTime)  # TODO: auto timestamp
#     date_updated = db.Column(db.DateTime)
#
#
# class MajorBin(db.Model):
#     """
#     Allows for grouping of MinorBins together into a parent bin
#     """
#     __tablename__ = 'major_bin'
#     id = db.Column(db.Integer, primary_key=True)
#     team_id = db.Column(db.Integer, db.ForeignKey('team.id'))
#     # TODO: color
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#
#
# class IncomeAllocation(db.Model):
#     """
#     Joins the MinorBin and Income tables to track how much of an income is divided into each bin
#     """
#     __tablename__ = 'income_allocation'
#     income_id = db.Column(db.Integer, db.ForeignKey('minor_bin.id'), primary_key=True)
#     minor_bin_id = db.Column(db.Integer, db.ForeignKey('income.id'), primary_key=True)
#
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#     amount = db.Column(db.Float)
#     interval_unit = db.Column(db.Integer)
#     interval_type = db.Column(db.String(128))  # TODO: change interval_type to enum
#
#
# class MinorBin(db.Model):
#     """
#     Most granular form of bin tracking
#     """
#     __tablename__ = 'minor_bin'
#     id = db.Column(db.Integer, primary_key=True)
#     major_bin_id = db.Column(db.Integer, db.ForeignKey('major_bin.id'))
#     incomes = relationship('Income', secondary='income_allocation', back_populates='minor_bins')
#
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#     amount = db.Column(db.Float)
#
#
# class Income(db.Model):
#     """
#     Tracks each type of income for a Team
#     """
#     __tablename__ = 'income'
#     id = db.Column(db.Integer, primary_key=True)
#     team_id = db.Column(db.Integer, db.ForeignKey('team.id'))
#     minor_bins = relationship('MinorBin', secondary='income_allocation', back_populates='incomes')
#
#     name = db.Column(db.String(128))
#     description = db.Column(db.String(128))
#     amount = db.Column(db.Float)
#     interval_unit = db.Column(db.Integer)
#     interval_type = db.Column(db.String(128))  # TODO: change interval_type to enum


class TeamSchema(ma.Schema):
    class Meta:
        # Fields to expose
        fields = ("id", "name")


@app.route("/teams", methods=['GET', 'POST'])
def get_all():
    team_schema = TeamSchema()
    teams_schema = TeamSchema(many=True)
    if request.method == 'GET':
        print('GET REQUEST HIT@!')
        teams = Team.query.all()
        print(teams)
        result = teams_schema.dump(teams)
        print(result)
        return jsonify(result)
    elif request.method == 'POST':
        import random
        t = Team()
        t.id = random.randint(-1000, 1000)
        t.name = f'User: {t.id}'
        print(t)
        db.session.add(t)
        db.session.commit()
        result = team_schema.dump(t)
        print(result)
        return jsonify(result)


@app.route('/')
def index():
    return 'Index Page'
