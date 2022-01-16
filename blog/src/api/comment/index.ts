import { api } from '../../boot/axios';

export const listComment = async (id: number, params: any) => {
    return await api.get(`/api/blog/comment/list/${id}`,{params:params});
};

export const listSettingComment = async (params: any) => {
    return await api.get('/api/blog/setting/comment',{params:params});
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

export const delComment = async (id: number) => {
    return await api.delete(`/api/blog/comment/${id}`);
};