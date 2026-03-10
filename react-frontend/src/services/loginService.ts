import axios from 'axios';
import config from '../config';

// Define the shape of the response data here
export interface UserData {
  id: number;
  name: string;
  email: string;
}

//const API_URL = 'http://localhost:80/api';
const API_URL = config.apiBaseUrl

// Just a pure function that returns a Promise
export const loginService = async (credentials: any) => {
  // This matches the POST endpoint we discussed earlier
  const response = await axios.post(`${API_URL}/login`, credentials);
  return response.data;
};
