class SiteNotFoundError extends Error {

    #code = 404;

    constructor(id) {
        super(`No site found with ID: ${id}`);

        this.name = this.constructor.name;
    }

    getCode() {
        return this.#code;
    }
}

module.exports = SiteNotFoundError;
