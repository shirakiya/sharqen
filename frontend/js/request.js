import axios from 'axios';
import urljoin from 'url-join';


function _getURL() {
  return 'http://0.0.0.0:1323/';
}

function _getFullURL(endpoint) {
  const url = _getURL();

  return urljoin(url, endpoint);
}

export function get(endpoint, params) {
  const url = _getFullURL(endpoint);

  if (! params) {
    params = {};
  }

  return axios.get(url, {
    params: params,
  });
}

export function post(endpoint, params) {
  const url = _getFullURL(endpoint);

  return axios.post(url, params);
}
