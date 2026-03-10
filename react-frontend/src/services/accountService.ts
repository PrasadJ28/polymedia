import axios from 'axios';
import config from '../config';

export interface AccountDetails {
  userid: number;
  username: string;
  email: string;
  firstname: string;
  lastname: string;
  password: string;
}

//const API_URL = 'http://localhost:80/api';
const API_URL = config.apiBaseUrl

export const createAccount = async (details: any) => {
  // This matches the POST endpoint we discussed earlier
  const response = await axios.post(`${API_URL}/create-account`, details);
  return response.data;
};
