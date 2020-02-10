'use strict';

const amqp = require('amqplib/callback_api');

class MessageListener {

    #connection;

    #channel;

    #queue;

    #exchanges = [];

    #messageHandler;

    constructor(messageHandler) {
        this.#messageHandler = messageHandler;

        this.connect();
    }

    connect() {
        const self = this;

        amqp.connect('amqp://guest:guest@rabbit:5672?heartbeat=60', function(err, conn) {
            if (err) {
                console.error("[AMQP]", err.message);
                return setTimeout(start, 1000);
            }
            conn.on("error", function(err) {
                if (err.message !== "Connection closing") {
                    console.error("[AMQP] conn error", err.message);
                }
            });
            conn.on("close", function() {
                console.error("[AMQP] reconnecting");
                return setTimeout(start, 1000);
            });
            console.log("[AMQP] connected");

            self.onConnected();
        });
    }

    onConnected(connection) {
        this.#connection = connection;
        this.startClient();
    }

    startClient() {
        const self = this;
        this.#connection.createChannel(function(err, ch) {
            if (closeOnErr(err)) return;
            ch.on("error", function(err) {
                console.error("[AMQP] channel error", err.message);
            });
            ch.on("close", function() {
                console.log("[AMQP] channel closed");
            });

            ch.prefetch(10);
            ch.assertQueue("gs-pack", { durable: true }, function(err, _ok) {
                if (closeOnErr(err)) return;
                ch.consume("gs-pack", self.processMsg, { noAck: false });
                console.log("Worker is started");
            });

            self.#channel = ch;
        });
    }

    processMsg(msg) {
        self.handle(msg, function(ok) {
            try {
                if (ok)
                    ch.ack(msg);
                else
                    ch.reject(msg, true);
            } catch (e) {
                closeOnErr(e);
            }
        });
    }

    handle(message, callback) {
        this.#messageHandler.processMessage(msg.content.toString());
        callback(true);
    }
}

module.exports = MessageListener;