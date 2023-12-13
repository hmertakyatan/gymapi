import axios from 'axios';

export const logoutUser = async () => {
  try {
    const response = await api.get('/auth/logout');
    window.location.href = '/login';
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.message);
  }
};

const api = axios.create({
  baseURL: '/api',
  withCredentials: true,
});
let refresh = false;

api.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    if (error.response.status === 401 && !refresh) {
      refresh = true;
      try {
        const response = await api.get('/auth/refresh');
        if (response.status === 200) {
          api.defaults.headers.common['Authorization'] = `Bearer ${response.data['access_token']}`;
          refresh = false;
          window.location.reload();
          return api(error.config); 
        }
      } catch (refreshError) {
        refresh = false; // Token yenileme başarısız, refresh'i sıfırla
        await logoutUser();
        return Promise.reject(error); // Hata durumunu ileterek işlemi sonlandır
      }
    }
    return Promise.reject(error);
  }
);

export const fetchCustomers = async () => {
  try {
    const response = await api.get('customer'); // Müşterileri getiren API endpoint'i
    return response.data;
  } catch (error) {
    throw new Error(error.message);
  }
};


export const createMembership = async (membershipData) => {
  try {
    const response = await api.post('membership', membershipData);
    return response.data;
  } catch (error) {
    throw new Error(error.response.data.message);
  }
};

export const fetchMembershipTypes = async (membershipTypeData) => {
  try {
    const response = await api.get('membership-type', membershipTypeData);
    return response.data;
  } catch (error) {
    throw new Error(error.message);
  }
};

export const fetchMemberships = async (membershipData) => {
  try {
    const response = await api.get('membership', membershipData);
    return response.data;
    } catch (error) {
    throw new Error(error.message);
  }
};

export default api;
