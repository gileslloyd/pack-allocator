'use strict';

const mongoose = require('mongoose');

const siteSchema = new mongoose.Schema({
  title: {
    type: String,
  },
  url: {
    type: String,
  },
});

const Site = mongoose.model('Site', siteSchema);

module.exports = Site;
