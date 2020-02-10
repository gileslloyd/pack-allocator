'use_strict';

class MessageHandler {
    #routes = [];

    add(route, controller, method) {
        this.#routes[route] = [new controller, method];
    }

    processMessage(message) {
        const payload = JSON.parse(message);
        const route = this.getRoute(payload);
        let response = null;

        if (route) {
            const controller = route[0];
            const method = route[1];

            response = (new controller()).method(
                new MessagePayload(payload['payload'] ?? [])
            );
        }

        return response;
    }

    getRoute(payload) {
        let route = false;

        if (payload.role) {
            const route_name = `${payload['role']}.${payload['cmd'] ?? 'execute'}`;

            if(this.#routes[route_name]) {
                route = this.#routes[route_name];
            }
        }

        return route;
    }
}

module.exports = MessageHandler;