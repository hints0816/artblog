import { api } from '../../boot/axios';
import { ActicleInfo } from './acticle.model';

export const listArticle = async (params: any) => {
    console.log(params)
    return await api.get('/api/blog/list');
};

export const getArticle = async (id: number) => {
    return await api.get(`/api/blog/article/${id}`);
};

export const addArticle = async (acticleParams: ActicleInfo) => {
    return await api.post('/api/blog/article', acticleParams);
};