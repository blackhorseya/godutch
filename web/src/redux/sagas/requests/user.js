import axios from 'axios';
import qs from 'qs';

const url = process.env.REACT_APP_API_URL || '';

export function requestLogin(email, password) {
  return axios.request({
    method: 'post',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    data: qs.stringify({
      email: email,
      password: password,
    }),
    url: `${url}/api/v1/auth/login`,
  });
}

export function requestSignup(email, password, name) {
  return axios.request({
    method: 'post',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    data: qs.stringify({
      email: email,
      password: password,
      name: name,
    }),
    url: `${url}/api/v1/auth/signup`,
  });
}

export function requestGetMe(token) {
  return axios.request({
    method: 'get',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    url: `${url}/api/v1/users/me`,
  });
}

export function requestLogout(token) {
  return axios.request({
    method: 'delete',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    url: `${url}/api/v1/auth/logout`,
  });
}
