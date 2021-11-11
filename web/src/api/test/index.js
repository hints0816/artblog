import { api } from '../../boot/axios';

export const getCart = async () => {
    return await api.get('/123/123', null);
};