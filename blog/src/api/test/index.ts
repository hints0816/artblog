import { api } from '../../boot/axios';
import { ArticleInfo } from './article.model';

export const login = async (params: any) => {
    return await api.post('/api/login/users', params);
};

export const adduser = async (params: any) => {
    return await api.post('/api/user/add', params);
};

export const getProfile = async (id: number) => {
    return await api.get(`/api/blog/profile/${id}`);
};

export const getProfileMe = async () => {
    return await api.get('/api/blog/getme');
};

export const listArticle = async (params: any) => {
    return await api.get('/api/blog/articlelist',{params:params});
};

export const getArticle = async (id: number) => {
    return await api.get(`/api/blog/article/${id}`);
};

export const addArticle = async (articleParams: ArticleInfo) => {
    return await api.post('/api/blog/article', articleParams);
};

export const listComment = async (id: number, params: any) => {
    return await api.get(`/api/blog/comment/list/${id}`,{params:params});
};

export const addComment = async (params: any) => {
    return await api.post('/api/blog/comment',params);
};

export const digg = async (params: any) => {
    return await api.post('/api/blog/comment/digg',params);
};

export const undigg = async (params: any) => {
    return await api.post('/api/blog/comment/undigg',params);
};

export const delArticle = async (id: number) => {
    return await api.delete(`/api/blog/article/${id}`);
};