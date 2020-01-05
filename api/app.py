import os
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate

# from sqlalchemy import Table, Column, Integer, ForeignKey
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base
import sqlalchemy as sa

Base = declarative_base()


app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = f'postgres://fyparbfgmnlvnt:{os.getenv("DB_PASS")}@ec2-174-129-253-162.compute-1.amazonaws.com:5432/d93afugd4i967s'

db = SQLAlchemy(app)
migrate = Migrate(app, db)


class Team(Base):
    __tablename__ = 'team'
    id = sa.Column(sa.Integer, primary_key=True)
    users = relationship('User', back_populates='team')
    banks = relationship('Bank', back_populates='team')
    major_bins = relationship('MajorBin', back_populates='team')
    incomes = relationship('Income', back_populates='team')

    name = sa.Column(sa.String(128))


class User(Base):
    __tablename__ = 'user'
    id = sa.Column(sa.Integer, primary_key=True)
    team_id = sa.Column(sa.Integer, sa.ForeignKey('team.id'))

    name = sa.Column(sa.String(128))


class Bank(Base):
    __tablename__ = 'bank'
    id = sa.Column(sa.Integer, primary_key=True)
    team_id = sa.Column(sa.Integer, sa.ForeignKey('team.id'))
    bank_accounts = relationship('BankAccount', back_populates='bank')

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))


class BankAccount(Base):
    __tablename__ = 'bank_account'
    id = sa.Column(sa.Integer, primary_key=True)
    bank_id = sa.Column(sa.Integer, sa.ForeignKey('bank.id'))

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))


class Transaction(Base):
    # TODO: add status such as pending/confirmed
    __tablename__ = 'transaction'
    id = sa.Column(sa.Integer, primary_key=True)
    bank_account_id = sa.Column(sa.Integer, sa.ForeignKey('bank_account.id'))
    minor_bin_id = sa.Column(sa.Integer, sa.ForeignKey('minor_bin.id'))

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))
    date_of_transaction = sa.Column(sa.DateTime)
    date_created = sa.Column(sa.DateTime)
    date_updated = sa.Column(sa.DateTime)


class MajorBin(Base):
    # TODO: color
    __tablename__ = 'major_bin'
    id = sa.Column(sa.Integer, primary_key=True)
    team_id = sa.Column(sa.Integer, sa.ForeignKey('team.id'))

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))


class IncomeAllocation(Base):
    """
    Joins the MinorBin and Income tables to track how much of an income is divided into each bin.
    """
    __tablename__ = 'income_allocation'
    income_id = sa.Column(sa.Integer, sa.ForeignKey('minor_bin.id'), primary_key=True)
    minor_bin_id = sa.Column(sa.Integer, sa.ForeignKey('income.id'), primary_key=True)
# income_allocation = sa.Table('income_allocation', Base.metadata,
#     sa.Column('income_id', sa.Integer, sa.ForeignKey('income.id')),
#     sa.Column('minor_bin_id', sa.Integer, sa.ForeignKey('minor_bin.id')),
#     sa.Column('amount', sa.Integer),
#     sa.Column('interval_unit', sa.Integer),
#     sa.Column('interval_type', sa.Column(sa.String(128)))  # TODO: change interval_type to enum
# )


class MinorBin(Base):
    """
    Most granular form of bin tracking
    """
    __tablename__ = 'minor_bin'
    id = sa.Column(sa.Integer, primary_key=True)
    major_bin_id = sa.Column(sa.Integer, sa.ForeignKey('major_bin.id'))
    incomes = relationship('Income', secondary='income_allocation', back_populates='minor_bins')
    # incomes = relationship('Income', secondary=income_allocation, back_populates="minor_bins")

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))
    amount = sa.Column(sa.Float)


class Income(Base):
    """
    Tracks each type of income for a Team
    """
    __tablename__ = 'income'
    id = sa.Column(sa.Integer, primary_key=True)
    team_id = sa.Column(sa.Integer, sa.ForeignKey('team.id'))
    minor_bins = relationship('MinorBin', secondary='income_allocation', back_populates='incomes')
    # minor_bins = relationship('MinorBin', secondary=income_allocation_table, back_populates="incomes")

    name = sa.Column(sa.String(128))
    description = sa.Column(sa.String(128))
    amount = sa.Column(sa.Float)
    interval_unit = sa.Column(sa.Integer)
    interval_type = sa.Column(sa.String(128))  # TODO: change interval_type to enum


@app.route('/')
def index():
    return 'Index Page'
