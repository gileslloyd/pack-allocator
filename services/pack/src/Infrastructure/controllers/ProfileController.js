'use strict';

const BaseController = require('./BaseController');

class ProfileController extends BaseController {

    #siteRepo;

    #profiler;

    constructor(siteRepository, profileService) {
        super();
        this.#siteRepo = siteRepository;
        this.#profiler = profileService;
    }
    
    profileAll = (req, res) => {

        this.#siteRepo.getAll().then((sites) => {
            this.runProfile(sites, res);
        }).catch((error) => {
            this.errorResponse(res, error);
        });
    }

    profileSite = (req, res) => {
        this.#siteRepo.getByID(req.params.siteId).then((site) => {
            this.runProfile([site], res);
        }).catch((error) => {
            this.errorResponse(res, error);
        });
    }

    runProfile(sites, res) {
        this.#profiler.profileAll(sites)
            .then((response) => {
                this.successResponse(res, response);
            }).catch((error) => {
                this.errorResponse(res, error);
            });
    }
}

module.exports = ProfileController;
