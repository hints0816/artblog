import { api } from '../../boot/axios';

export const addArt = async (param) => {
    return await api.post('/api/v1/addcomment', param);
};

export const getCart = async (param) => {
    return await api.post('/generate/test1', param);
};

export const getCart1 = async (param) => {
    return await api.get('/generate/test', param);
};