'use_strict';

class MessagePayload {

    #payload;

    constructor(payload)  {
        this.#payload = payload;
    }

    getPayload() {
        return this.#payload;
    }

    get(key, def = null) {
        let value = this.#payload;

        key.split('.').forEach(function(part) {
            if (!value[part]) {
                return def;
            }

            value = value[part];
        });

        return value;
    }
}

module.exports = MessagePayload;