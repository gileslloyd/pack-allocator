'use strict';

const BaseController = require('./BaseController');

class SiteController extends BaseController {

    #siteRepo;

    constructor(siteRepository) {
        super();
        this.#siteRepo = siteRepository;
    }
    
    all = (req, res) => {

        this.#siteRepo.getAll().then((sites) => {
            this.successResponse(res, sites);
        }).catch((error) => {
            this.errorResponse(res, error);
        });
    }
}

module.exports = SiteController;
