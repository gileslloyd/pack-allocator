'use strict';

const ProfileResult = require('./ProfileResult');
// Ideally the API key below would be in a secret store
// rather than an environment variable
const baseUrl = `https://www.googleapis.com/pagespeedonline/v5/runPagespeed?key=${process.env.GOOGLE_API_KEY}&url=`;
const axios = require('axios');

class ProfileService {
  profileAll(sites) {
    const results = Promise.all(
      sites.map(this.profileSite),
    );

    return results;
  }

  async profileSite(site) {
    try {
      const response = await axios.get(baseUrl + site.url);

      return new ProfileResult(
        site,
        response.status,
        response.data.lighthouseResult.categories.performance.score * 100,
      );
    } catch (error) {
      return new ProfileResult(
        site,
        error.response.status,
        'N/A',
        error,
      );
    }
  }
}

module.exports = ProfileService;
