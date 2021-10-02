import axios from 'axios';

const url = process.env.REACT_APP_API_URL || '';

export function requestListActivities(page, size) {
  return axios.request({
    method: 'get',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`,
    },
    url: `${url}/api/v1/activities?page=${page}&size=${size}`,
  });
}
