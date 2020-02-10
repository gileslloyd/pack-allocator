'use strict';

const mongoose = require('mongoose');
const Site = require('../../Domain/Site/Site');

const connection = 'mongodb://db:27017/speed-dash';

const connectDb = () => {
  return mongoose.connect(connection);
};

module.exports = connectDb;
