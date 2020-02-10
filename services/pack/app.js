'use strict';

const MessageListener = require('./src/Infrastructure/message/MessageListener');

const app = new MessageListener();

module.exports.app = app;
