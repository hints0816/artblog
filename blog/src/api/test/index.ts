import { api } from '../../boot/axios';
import { ArticleInfo } from './article.model';

export const login = async (params: any) => {
    return await api.post('/api/login/users', params);
};

export const checkToken = async (params: any) => {
    return await api.get('/api/login/checktoken', {params:params});
};

export const sendvalidator = async (params: any) => {
    return await api.post('/api/login/sendvalidator', params);
};

export const signup = async (params: any) => {
    return await api.post('/api/blog/signup', params);
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

export const listArtList = async (params: any) => {
    return await api.get('/api/blog/artlist',{params:params});
};

export const listCategory = async (params: any) => {
    return await api.get('/api/blog/category',{params:params});
};

export const updateCategory = async (params: any) => {
    return await api.put('/api/blog/category', params);
};

export const listMeArticle = async (params: any) => {
    return await api.get('/api/blog/me/articlelist',{params:params});
};

export const getArticle = async (id: number) => {
    return await api.get(`/api/blog/article/${id}`);
};

export const getImgContent = async (id: number) => {
    return await api.get(`/api/blog/artimg/${id}`);
};

export const addArticle = async (articleParams: ArticleInfo) => {
    return await api.post('/api/blog/article', articleParams);
};

export const listComment = async (id: number, params: any) => {
    return await api.get(`/api/blog/comment/list/${id}`,{params:params});
};

export const listImgComment = async (id: number, params: any) => {
    return await api.get(`/api/blog/imgcomment/list/${id}`,{params:params});
};

export const addComment = async (params: any) => {
    return await api.post('/api/blog/comment',params);
};

export const addImgComment = async (params: any) => {
    return await api.post('/api/blog/imgcomment',params);
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

export const uploadImage = async (params: any) => {
    const config = {
        headers: {'Content-Type': 'multipart/form-data'}
    }
    return await api.post('/api/blog/upload', params, config);
};

export const uploadAvatarImage = async (params: any) => {
    const config = {
        headers: {'Content-Type': 'multipart/form-data'}
    }
    return await api.post('/api/blog/avatar/upload', params, config);
};

export const editEmoji = async (params: any) => {
    return await api.post('/api/blog/emoji', params);
};

export const updateProfile = async (params: any) => {
    return await api.put('/api/blog/profile', params);
};

export const logout = async () => {
    return await api.get('/api/blog/logout');
};

export const deleteCategory = async (id: number) => {
    return await api.delete(`/api/blog/category/${id}`);
};

export const postArt = async (params: any) => {
    const config = {
        headers: {'Content-Type': 'multipart/form-data'}
    }
    return await api.post('/api/blog/postart', params, config);
};