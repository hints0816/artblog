import { api } from '../../boot/axios';
import { ArticleInfo } from './article.model';

export const login = async (params: any) => {
    return await api.post('/api/login/users', params);
};

export const adduser = async (params: any) => {
    return await api.post('/api/user/add', params);
};

export const getProfileMe = async () => {
    return await api.get('/api/blog/getme');
};

export const listArticle = async (params: any) => {
    const pagenum: string = params.pagenum
    const pagesize: string = params.pagesize
    return await api.get('/api/blog/articlelist?pagenum='+pagenum+'&pagesize='+pagesize);
};

export const getArticle = async (id: number) => {
    return await api.get(`/api/blog/article/${id}`);
};

export const addArticle = async (articleParams: ArticleInfo) => {
    return await api.post('/api/blog/article', articleParams);
};