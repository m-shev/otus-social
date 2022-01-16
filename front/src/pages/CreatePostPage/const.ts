import {FieldItem} from '../../components/Field/types';
import {CreatePostForm} from '../../types/post';

export type CreatePostField = Pick<CreatePostForm, 'content' | 'imageLink'>;

export const initialValues: CreatePostField = {
    content: '',
    imageLink: '',
};

export const fieldList: FieldItem<CreatePostField>[] = [
    {
        id: 'content',
        title: 'Текст поста *',
        type: 'textarea',
        required: true,
    },
    {
        id: 'imageLink',
        title: 'Ссылка на картинку',
        type: 'input',
        required: false,
    },
];
