import axios from 'axios';

// the base of the backend URL
const base = process.env.REACT_APP_BACKEND_URL;

const GetAxiosInstance = token => {
  const instance = axios.create({
    baseURL: base + '/api/v1', // example: http://localhost:8080/api/v1
    // add the token for authentication
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  return instance;
};

// export base url and function to get axios instance
export { base, GetAxiosInstance };
