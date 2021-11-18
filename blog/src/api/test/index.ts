import { api } from '../../boot/axios';

export const getArticle = async (id: number) => {
    return await api.get(`/blog/acticle/${id}`);
};