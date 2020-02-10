'use strict';

const Site = require('./Site');
const SiteNotFoundError = require('./SiteNotFoundError');

class SiteRepository {
  async getAll() {
    return await Site.find();
  }

  async getByID(id) {
    try {
      let site = await Site.findById(id);

      if (site === null) {
        throw new SiteNotFoundError(id);
      }

      return site;
    } catch (error) {
      throw new SiteNotFoundError(id);
    }
  }
}

module.exports = SiteRepository;
