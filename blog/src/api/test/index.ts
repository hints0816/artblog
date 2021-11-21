import { api } from '../../boot/axios';
import { ArticleInfo } from './article.model';

export const login = async (params: any) => {
    return await api.post('/api/login/users', params);
};

export const getProfileMe = async () => {
    return await api.get('/api/blog/getme');
};

export const listArticle = async (params: any) => {
    console.log(params)
    return await api.get('/api/blog/articlelist');
};

export const getArticle = async (id: number) => {
    return await api.get(`/api/blog/article/${id}`);
};

export const addArticle = async (articleParams: ArticleInfo) => {
    return await api.post('/api/blog/article', articleParams);
};